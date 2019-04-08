# FROM golang:latest 
# RUN mkdir /app 
# ADD . /app/ 
# WORKDIR /app 
# RUN go build -o main . 
# CMD ["/app/main"]

FROM golang:alpine AS build-env
WORKDIR /app
COPY main .

FROM scratch
LABEL maintainer="Adam Plánský <adamplansky@gmail.com>"
ENV PORT 5000
EXPOSE 5000
ENV DATABASE_URL=""
COPY --from=build-env /app/main /
CMD ["/main"]
