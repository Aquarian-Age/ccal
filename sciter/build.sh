#!/bin/bash

go build -o ~/ccal/sciter/ccal-sciter-dev -ldflags="-s -w" ccal.go
cp ./ccal.html ~/ccal/sciter/ccal.html
