#!/bin/bash

WS=$PWD

cd web
npm install
npm run build

cd $WS

echo "build..."
go build -o sysmon -ldflags '-linkmode "external" -extldflags "-static"'

echo "pack..."
tar zcvf sysmon.tar.gz sysmon views conf web/dist swagger
