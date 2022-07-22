#!/bin/bash

## GOARM=7 GOOS=linux  GOARCH=arm go build -o watch-arm -ldflags="-s -w" main.go
## GOARM=7 GOOS=linux  GOARCH=arm go build -o gpioNumber-armv7 -ldflags="-s -w" gpioNumber.go 

# GOARM=7 GOOS=linux  GOARCH=arm go build -o led-armv7 -ldflags="-s -w" led.go
# GOARM=7 GOOS=linux  GOARCH=arm go build -o led-armv7 -ldflags="-s -w" led.go 
# GOOS=linux  GOARCH=amd64  go build -o watch-linux -ldflags="-s -w" main.go
# GOARM=7 GOOS=linux  GOARCH=arm  go build -o gpiox-8 -ldflags="-s -w" gpiox.go
GOARM=7 GOOS=linux GOARCH=arm go build -o readInputNumber -ldflags="-w -s" readInpuptNumber.go
