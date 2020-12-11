#!/bin/bash

if [ "$1" = "--help" ];
then
  echo "usage: run.sh docker-image schemafiles..."
  exit 0
fi

SCHEMA_DIR=/opt/schemafiles
VOLUMES=""

DOCKER_IMAGE="$1"

echo "${@:2}"

for schema in "${@:2}"
do
  schema_path=$(realpath "$schema")
  schema_name=$(basename "$schema_path")
  VOLUMES+=" -v $schema_path:$SCHEMA_DIR/$schema_name"
done

docker run -it $VOLUMES "$DOCKER_IMAGE"
