#!/bin/bash

# set android SDK_ROOT
# go install gioui.org/cmd/gogio@latest
v="3.3.9"
gogio -target android -icon icon.png  -o /home/xuan/ccal/sjqm/gioui/sjqmUI-$v.apk .
# go build -ldflags="-s -w" -tags nowayland -o /home/xuan/ccal/sjqm/gioui/sjqmUI-JQH-$v-linux
# GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H windowsgui" -tags nowayland -o /home/xuan/ccal/sjqm/gioui/sjqmUI-JQH-$v-win10.exe