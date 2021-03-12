#!/usr/bin/env bash

echo '\n====\nBuilding the user interface...'
pushd ui
./build.sh
popd

echo '\n====\nBuilding the application...'
docker build -t bonjoursoftware/mycollections:local .
