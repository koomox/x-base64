#!/bin/bash
GOPATH=$(pwd)
cd ${GOPATH}
go build ./src/main.go
ldd ./x-base64