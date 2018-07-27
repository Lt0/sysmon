#!/bin/bash
set -e

./pack.sh
docker build -t sysmon .
