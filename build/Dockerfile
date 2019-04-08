# FROM golang:latest 
# RUN mkdir /app 
# ADD . /app/ 
# WORKDIR /app 
# RUN go build -o main . 
# CMD ["/app/main"]

FROM golang:alpine AS build-env
WORKDIR /app
COPY . .
RUN apk update \
    && apk upgrade \
    && apk add git \
    && cd /app \
    && go get ./... 
RUN CGO_ENABLED=0 GOOS=linux go build -o httpserver cmd/server/main.go

FROM scratch
LABEL maintainer="Adam Plánský <adamplansky@gmail.com>"
ENV PORT 5000
EXPOSE 5000
ENV DATABASE_URL=""
COPY --from=build-env /app/httpserver /
CMD ["/httpserver"]