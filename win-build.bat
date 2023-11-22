@echo off
set CGO_ENABLED=1 
set GOOS=windows
set GOARCH=amd64
go build -o ./bin/windows/cp2server.exe