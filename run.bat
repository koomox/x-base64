@echo off
cd /d %~dp0
SET GOPATH=%~dp0
go build ./src/main.go
pause