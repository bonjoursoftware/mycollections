#!/usr/bin/env bash
docker build -t ui-builder .
CONTAINER_ID=$(docker create ui-builder)
docker cp "$CONTAINER_ID":/app/bin/mycollections.js ../src/main/resources/static
docker cp "$CONTAINER_ID":/app/package-lock.json .
docker container rm "$CONTAINER_ID"
