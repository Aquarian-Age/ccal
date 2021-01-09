#!/bin/bash

go build -o ~/ccal/web/ccal-web-dev -ldflags="-s -w" main.go home.go  yiji.go names.go jz60.go jianchu.go tongshu.go other.go

GOOS=windows GOARCH=amd64 go build -o ~/ccal/web/ccal-web-dev.exe -ldflags="-s -w" main.go home.go  yiji.go names.go jz60.go jianchu.go tongshu.go other.go
