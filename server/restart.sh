#!/bin/bash

# Install fswatch (MacOS)
# brew install fswatch

PORT=$1
# Source: https://proto.holochain.org/webserver_auto_restart
# -e: exclude all files from fswatch
# -i: include only *.go files for fswatch
# -I{}: capture the actual line output as no arguments should be passed to go run`
 # Usage: fswatch -e ".*" -i ".go" . | xargs -n1 -I{} './restart.sh' $PORT

# Do not rely on ps as it will get the process that calls go-build but not the actual server listening on the port
# PID=$(ps | grep "go run server.go" | awk '{print $1}')

#-t: outputs only the PID - used specifically for piping to kill a process
PID=$(lsof -t -i:$PORT -sTCP:LISTEN)
echo killing running server at PID $PID
kill $PID
# ps | grep "go run server.go" | awk '{if($1!="") {print "killing process: "$1; system("kill " $1)}}'
echo 'starting server' & go run server.go &
