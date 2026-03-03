#!/bin/sh
ENV_DIR=/var/run/secrets/app
set -e;
if [ -d "$ENV_DIR" ]; then
  echo "Copy file env"
  cp "$ENV_DIR/env" /app/.env
fi
echo "Apply migrations"
/app/migrate.sh up
echo "Start web server"
/app/apiserver $*
