#!/bin/bash

t=$(date +%Y-%M-%d-%T)
go build -o ~/ccal/sciter/ccal-sciter-v2-$t -ldflags="-s -w" main.go
cp ./cal.html ~/ccal/sciter/cal.html
