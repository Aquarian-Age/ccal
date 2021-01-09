#!/bin/bash

go build -o ~/ccal/sciter/ccal-sciter-v2 -ldflags="-s -w" main.go
cp ./cal.html ~/ccal/sciter/cal.html
