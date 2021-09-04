#!/bin/bash

#
# The code is automatically generated by the Goland.
# Copyright © Aquarian-Age. All Rights Reserved.
# Licensed under MIT
#

# set -v
# go get -u github.com/ying32/liblclbinres
app="rjqm-govcl"
version="0.0.6" 
v=$(go version)
gov=${v:1-20}
commit=$(git rev-parse HEAD)
time=$(date +"%Y-%m-%d %H:%M:%S")
path="liangzi.local/qx/cmd"

flags="
-X $path.Version=$version
-X '$path.GoVersion=$gov'
-X '$path.BuildTime=$time'
-X $path.GitCommit=$commit
-s -w"
go build -ldflags "$flags" -trimpath -tags tempdll -o /home/xuan/rpm/BUILD/$app-$version/$app .
