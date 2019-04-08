#!/usr/bin/env bash
#CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go
go build -gcflags '-N -l' -o server cmd/server/main.go
# docker build --tag todo-go-local-build .
docker-compose -f docker-compose-local-build.yaml up
# rm -f main