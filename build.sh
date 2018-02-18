#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/main_amd64_windows.exe main.go
