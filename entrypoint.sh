#!/bin/sh
set -e

/opt/drtest/drtest-gen --target-dir=/opt/drtest/generated /opt/schemafiles/*.avsc

go run /opt/drtest/webserver/cmd/app.go