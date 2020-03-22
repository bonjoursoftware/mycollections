#!/bin/sh

echo '=\n====\nBuilding the user interface...'
cd ./ui
./build.sh

echo '\n====\nBuilding the application...'
cd ../
./gradlew clean shadowJar