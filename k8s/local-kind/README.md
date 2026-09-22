# Local kind smoke environment

This is a disposable self-hosted Kubernetes environment for the complete
Clabgate session lifecycle. It runs two real Clabgate replicas against the
Kubernetes API and the upstream Clabernetes controller. A stateful boundary
mock replaces CMS and GitLab; small protocol-compatible images replace the
production Jupyter and checker images.

The same lifecycle runs for every pull request in
`.github/workflows/e2e.yml` on an ephemeral GitHub-hosted kind cluster. No
production Kubernetes credentials are used.

The fixed attempt is `550e8400-e29b-41d4-a716-446655440000`. The frontend is
published at `http://127.0.0.1:18080`.

## Tested versions

- kind `0.33.0`;
- Kubernetes `1.33.12`;
- Clabernetes chart `0.6.0`, digest
  `sha256:21b22d346de11b11b9764d6ab25e21ab5b9212dfe7acdc5f40e4483bf6677cbd`;
- Docker Desktop on `linux/arm64`.

The same commands work on `amd64` when `GOARCH` and Docker `--platform` are
changed accordingly.

## Bootstrap the cluster

From the repository root:

```bash
kind create cluster \
  --config k8s/local-kind/kind.yaml \
  --image kindest/node:v1.33.12@sha256:3f5c8443c620245e4d355cfe09e96a91ead32ceaa569d3f1ca9edf0cb2fe2ff4

helm upgrade --install clabernetes \
  oci://ghcr.io/srl-labs/clabernetes/clabernetes \
  --version 0.6.0 \
  --namespace c9s \
  --create-namespace \
  --kube-context kind-cms-labs-local

kubectl wait --for=condition=Available deployment \
  --all -n c9s --context kind-cms-labs-local --timeout=5m
```

Do not add Helm `--wait` for chart `0.6.0`: its manager intentionally deletes
the bootstrap `clabernetes-config` ConfigMap after merging it, while Helm keeps
waiting for that object and eventually records a false timeout. Deployment
readiness is checked explicitly above.

If the node inherits a host proxy bound to `127.0.0.1`, containerd cannot use
that address from inside the node container. Configure Docker Desktop with a
proxy reachable as `host.docker.internal`, or set the corresponding systemd
environment on `cms-labs-local-control-plane` before pulling images.

## Build and load local images

The following commands are for Apple Silicon / ARM64:

```bash
mkdir -p clabgate/build clabgate/e2e/mock/build clabgate/e2e/jupyter/build

(cd clabgate && \
  CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o build/clabgate main.go && \
  cp start.sh build/start.sh)
(cd clabgate/e2e/mock && \
  CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o build/smoke-mock .)
(cd clabgate/e2e/jupyter && \
  CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o build/jupyter-smoke .)
(cd nextui-dashboard && ./node_modules/.bin/vite build)

docker build --platform linux/arm64 -t cms-labs/clabgate:local clabgate
docker build --platform linux/arm64 -t cms-labs/smoke-mock:local clabgate/e2e/mock
docker build --platform linux/arm64 -t cms-labs/jupyter-smoke:local clabgate/e2e/jupyter
docker build --platform linux/arm64 -t cms-labs/checker-smoke:local clabgate/e2e/checker
docker build --platform linux/arm64 -t cms-labs/front:local nextui-dashboard

kind load docker-image --name cms-labs-local \
  cms-labs/clabgate:local \
  cms-labs/smoke-mock:local \
  cms-labs/jupyter-smoke:local \
  cms-labs/checker-smoke:local \
  cms-labs/front:local

kubectl apply --context kind-cms-labs-local -f k8s/local-kind/smoke.yaml
```

After rebuilding images in an existing cluster, reset only the disposable
smoke state and restart its deployments:

```bash
./k8s/local-kind/reset-smoke.sh
```

The manifest uses `imagePullPolicy: Never` for top-level smoke components, so
these images cannot accidentally be pulled into a production cluster. The
session controller uses `IfNotPresent`; therefore rebuilds under the same tag
must be followed by `kind load docker-image` and replacement of the old Pod.

## Verify end to end

The verifier uses the repository's existing test RSA key to act as student 42.
It does not persist or print the five-minute JWT, workspace grant or cookie.

```bash
./k8s/local-kind/verify.sh
```

It checks all production-relevant boundaries:

- two Clabgate replicas are Ready and one holder owns Lease
  `clabgate-session-reconciler`;
- namespace `lab-550e8400-e29b-41d4-a716-446655440000`, Jupyter, PVC, Service
  and Clabernetes Topology are Ready;
- mock CMS moves from `pending` to `active` only after readiness;
- direct workspace access without a grant returns HTTP 401;
- `session.open` exchanges a short-lived grant for a scoped cookie and reaches
  the namespace-local Jupyter Service;
- `session.check` creates a Job and CMS receives a new `check_id`, score,
  report and bounded Pod logs.

The successful run on 2026-09-22 returned score `9/10`, report
`kind end-to-end smoke passed`, and checker Pod logs. Clabernetes Topology
`smoke` reported `TopologyReady=True` and `topologyState=running`.

## Inspect and remove

```bash
kubectl get pods -A --context kind-cms-labs-local
kubectl get topology -A --context kind-cms-labs-local
kubectl get lease -n cms-labs-system --context kind-cms-labs-local
curl --noproxy '*' http://127.0.0.1:18080/api/state

kind delete cluster --name cms-labs-local
```

The last command removes the entire disposable cluster, including the session
PVC and all smoke data.
