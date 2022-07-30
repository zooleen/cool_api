FROM golang:1.16-alpine

WORKDIR /app

RUN go mod tidy
RUN go build main.go -o cool_api

COPY . /

RUN /cool_api