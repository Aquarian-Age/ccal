#!/bin/bash
t=$(date +%Y-%m-%d-%H)
go build -o ~/ccal/qxqm/qxGUI-$t -ldflags="-s -w" .
