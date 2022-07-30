FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./

RUN go build -o /cool_api

EXPOSE 8080

CMD [ "/cool_api" ]

# ToDo basic multistage for future