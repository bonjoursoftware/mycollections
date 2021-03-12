#!/usr/bin/env bash

echo '\n====\nBuilding the user interface...'
pushd ui
./build.sh
popd

echo '\n====\nBuilding the application...'
docker build -t bonjoursoftware/mycollections:local .
CONTAINER_ID=$(docker create bonjoursoftware/mycollections:local)
docker cp $CONTAINER_ID:mycollections.jar ./mycollections-1.2.5-all.jar
docker container rm $CONTAINER_ID
