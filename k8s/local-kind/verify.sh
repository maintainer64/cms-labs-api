#!/bin/sh
set -eu

context=kind-cms-labs-local
namespace=lab-550e8400-e29b-41d4-a716-446655440000
session_id=550e8400-e29b-41d4-a716-446655440000
base_url=http://127.0.0.1:18080
repo_root=$(CDPATH='' cd -- "$(dirname "$0")/../.." && pwd)
cookie_jar=$(mktemp "${TMPDIR:-/tmp}/cms-labs-cookie.XXXXXX")
trap 'rm -f "$cookie_jar"' EXIT HUP INT TERM

kubectl wait --for=condition=Available deployment/clabgate deployment/front deployment/smoke-mock \
  -n cms-labs-system --context "$context" --timeout=90s >/dev/null
kubectl wait --for=condition=Available deployment/jupyter \
  -n "$namespace" --context "$context" --timeout=90s >/dev/null
kubectl wait --for=jsonpath='{.status.topologyReady}'=true topology/smoke \
  -n "$namespace" --context "$context" --timeout=90s >/dev/null
kubectl get pvc/jupyter service/jupyter -n "$namespace" --context "$context" >/dev/null

desired_replicas=$(kubectl get deployment/clabgate -n cms-labs-system \
  --context "$context" -o jsonpath='{.spec.replicas}')
available_replicas=$(kubectl get deployment/clabgate -n cms-labs-system \
  --context "$context" -o jsonpath='{.status.availableReplicas}')
test "$desired_replicas" = 2
test "$available_replicas" = 2

holder=$(kubectl get lease clabgate-session-reconciler -n cms-labs-system \
  --context "$context" -o jsonpath='{.spec.holderIdentity}')
test -n "$holder"

state=$(curl --noproxy '*' --fail-with-body --silent --show-error "$base_url/api/state")
test "$(printf '%s' "$state" | jq -r '.status')" = active

unauthorized=$(curl --noproxy '*' --silent --show-error --output /dev/null --write-out '%{http_code}' \
  "$base_url/clabgate/workspace/$session_id/")
test "$unauthorized" = 401

token=$(cd "$repo_root/clabgate" && \
  GOCACHE="${TMPDIR:-/tmp}/cms-labs-go-cache" go run ./e2e/token -env-file ../backend/.env.test)
authorization="Authorization: Bearer $token"
rpc_url="$base_url/clabgate/api/v1/rpc"

previous_check_id=$(printf '%s' "$state" | jq -r '.result.check_id // ""')
check_response=$(curl --noproxy '*' --fail-with-body --silent --show-error \
  -H "$authorization" -H 'Content-Type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":\"smoke-check\",\"method\":\"session.check\",\"params\":{\"session_id\":\"$session_id\"}}" \
  "$rpc_url")
job_name=$(printf '%s' "$check_response" | jq -er '.result.job_name')
kubectl wait --for=condition=complete "job/$job_name" -n "$namespace" \
  --context "$context" --timeout=90s >/dev/null

attempt=0
while :; do
  state=$(curl --noproxy '*' --fail-with-body --silent --show-error "$base_url/api/state")
  check_id=$(printf '%s' "$state" | jq -r '.result.check_id // ""')
  if test -n "$check_id" && test "$check_id" != "$previous_check_id"; then
    break
  fi
  attempt=$((attempt + 1))
  if test "$attempt" -ge 30; then
    printf '%s\n' 'checker result was not delivered to CMS within 30 seconds' >&2
    exit 1
  fi
  sleep 1
done

test "$(printf '%s' "$state" | jq -r '.result.current_score')" = 9
test -n "$(printf '%s' "$state" | jq -r '.result.report // ""')"
test -n "$(printf '%s' "$state" | jq -r '.result.logs // ""')"
test "$(printf '%s' "$state" | jq -r '.result.tasks[0].complete')" = true
test "$(printf '%s' "$state" | jq -r '.result.tasks[0].logs[0].message')" = "context is available"

open_response=$(curl --noproxy '*' --fail-with-body --silent --show-error \
  -H "$authorization" -H 'Content-Type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":\"smoke-open\",\"method\":\"session.open\",\"params\":{\"session_id\":\"$session_id\"}}" \
  "$rpc_url")
open_url=$(printf '%s' "$open_response" | jq -er '.result.url')
exchange_status=$(curl --noproxy '*' --silent --show-error --cookie-jar "$cookie_jar" \
  --output /dev/null --write-out '%{http_code}' "$base_url$open_url")
test "$exchange_status" = 303
workspace_cookie=$(awk '$6 == "clabgate_workspace" {print $7}' "$cookie_jar")
test -n "$workspace_cookie"
workspace=$(curl --noproxy '*' --fail-with-body --silent --show-error \
  -H "Cookie: clabgate_workspace=$workspace_cookie" \
  "$base_url/clabgate/workspace/$session_id/")
printf '%s' "$workspace" | grep -q 'workspace ready'

printf 'smoke passed: leader=%s job=%s check_id=%s score=%s\n' \
  "$holder" "$job_name" "$check_id" "$(printf '%s' "$state" | jq -r '.result.result_display')"
