#!/bin/bash

print_usage () {
  echo "usage: run.sh docker-image schemafiles..."
  exit 1
}

if [ $# -eq 0 ];
then
  print_usage
fi

if [ "$1" = "--help" ];
then
  print_usage
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

docker run -p 8080:8080 -it $VOLUMES "$DOCKER_IMAGE"
