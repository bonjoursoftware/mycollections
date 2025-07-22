#!/usr/bin/env bash

echo 'Building the user interface...'
pushd ui || exit
./build.sh
popd || exit

echo 'Building the application...'
docker build --no-cache -t bonjoursoftware/mycollections:local .
CONTAINER_ID=$(docker create docker.io/bonjoursoftware/mycollections:local)
docker cp "$CONTAINER_ID":mycollections.jar ./mycollections-1.8.2-all.jar
docker container rm "$CONTAINER_ID"
