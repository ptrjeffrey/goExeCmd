@echo off
set CGO_ENABLED=0 
set GOOS=linux
set GOARCH=amd64
go build -o ./bin/linux/cp2server