#!/bin/sh
set -e;
echo "Apply migrations"
/app/migrate.sh up
echo "Start web server"
/app/apiserver
