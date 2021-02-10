#!/bin/bash

# Source: https://proto.holochain.org/webserver_auto_restart
# Usage: fswatch -e ".*" -i ".go" . | xargs -n1 -I{} './restart.sh'

PID=$(ps | grep "go run server.go" | awk '{print $1}')
echo killing running server at PID $PID
kill $PID
# ps | grep "go run server.go" | awk '{if($1!="") {print "killing process: "$1; system("kill " $1)}}'
echo 'starting server' & go run server.go &
