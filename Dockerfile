# syntax=docker/dockerfile:1
FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init
RUN CGO_ENABLED=1 GOOS=linux go build -o /imagex

CMD [ "/imagex" ]
