# FROM golang:latest 
# RUN mkdir /app 
# ADD . /app/ 
# WORKDIR /app 
# RUN go build -o main . 
# CMD ["/app/main"]

FROM golang:alpine AS build-env
RUN apk add bash ca-certificates git gcc g++ libc-dev 

WORKDIR /go/src/github.com/adamplansky/todo

# Force the go compiler to use modules
ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .
RUN go mod download

FROM build-env AS go_builder
WORKDIR /go/src/github.com/adamplansky/todo
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/server/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o client cmd/client/main.go

FROM alpine as app
LABEL maintainer="Adam Plánský <adamplansky@gmail.com>"
ENV PORT 5000
EXPOSE 5000
ENV DATABASE_URL=""
# WORKDIR /app 
COPY --from=go_builder /go/src/github.com/adamplansky/todo/client /client
COPY --from=go_builder /go/src/github.com/adamplansky/todo/server /server
CMD ["/server", "-grpc"]