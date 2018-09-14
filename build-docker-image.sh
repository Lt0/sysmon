#!/bin/bash
set -e

./pack.sh
docker build -t sysmon .

if [[ $? -eq 0 ]]; then
	echo "=================== success ==================="
	echo "Start dockerized sysmon:"
	echo "docker run --name sysmon --restart=always -d -p 4096:2048 -v /proc:/hproc --privileged -it sysmon"
	echo "Visit web UI from localhost:4096 in browser"
else 
	echo "=================== failed ==================="
	echo build image failed...
fi
