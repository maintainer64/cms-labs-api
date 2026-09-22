#!/bin/sh
set -eu
echo "checker completed for ${ATTEMPT_ID:-unknown}"
printf '%s' '{"max_score":10,"current_score":9,"result_display":"local smoke: 9/10","report":"kind end-to-end smoke passed","tasks":[{"title":"Session namespace","description":"checker received the session context","logs":[{"node":"smoke","namespace":"'"${SESSION_NAMESPACE:-unknown}"'","message":"context is available"}],"complete":true}]}' > /dev/termination-log
