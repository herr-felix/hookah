#!/bin/bash

name=$1
status=$2

payload="{
  \"buildPath\":\"./$status\",
  \"projectName\":\"$name\",
  \"pullHandler\":\"demo\",
  \"pullParams\":\"\",
  \"pushHandler\":\"docker\",
  \"pushParams\":\"testing/$name:latest\"
}"

curl -X POST -H "Content-Type: application/json" -d "$payload" http://localhost:8080/build