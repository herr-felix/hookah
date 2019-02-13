#!/bin/bash

payload='{"buildPath":"./fail","projectName":"demo","pullHandler":"demo","pullParams":"","pushHandler":"","pushParams":""}'

curl -X POST -H "Content-Type: application/json" -d "$payload" http://localhost:8080/build