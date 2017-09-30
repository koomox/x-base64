#!/bin/bash
GOPATH=$(pwd)
cd ${GOPATH}
go build ./src/x-base64.go
ldd ./x-base64