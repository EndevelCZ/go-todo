#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go
# /dlv", "--listen=:2345", "--headless=true", "--api-version=2", "exec", "/server"
sleep 1
dlv --headless=true --api-version=2 exec /server