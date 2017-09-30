@echo off
SET GOPATH=%~dp0
cd /d %GOPATH%
go build ./src/x-base64.go
pause