# Compile stage
FROM golang:1.12-alpine AS build-env

RUN apk add --no-cache git
RUN mkdir -p /app
ADD . /app
WORKDIR /app



ENV CGO_ENABLED 0
# The -gcflags "all=-N -l" flag helps us get a better debug experience
# RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -gcflags "all=-N -l" -o /server cmd/server/main.go

# Compile Delve
WORKDIR $GOPATH
RUN go get github.com/derekparker/delve/cmd/dlv
 
# Final stage
FROM alpine:3.7
 
# Port 8080 belongs to our application, 2345 belongs to Delve
EXPOSE 8080 2345
 
# Allow delve to run on Alpine based containers.
RUN apk add --no-cache libc6-compat
 
WORKDIR /
 
COPY --from=build-env /server /
COPY --from=build-env /go/bin/dlv /
 
# Run delve
# CMD ["/dlv", "--listen=:2345", "--headless=true", "--api-version=2", "exec", "/server"]
CMD ["/dlv", "--listen=:2345", "--headless=true", "exec", "/server"]