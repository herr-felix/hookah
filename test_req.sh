#!/bin/bash

projectName=$1
status=$2
buildName=$3

payload="{
  \"buildPath\":\"./$status\",
  \"projectName\":\"$projectName\",
  \"buildName\": \"$buildName\",
  \"pullHandler\":\"demo\",
  \"pullParams\":\"\",
  \"pushHandler\":\"docker\",
  \"pushParams\":\"testing/$projectName:latest\"
}"

curl -X POST -H "Content-Type: application/json" -d "$payload" http://localhost:8080/build