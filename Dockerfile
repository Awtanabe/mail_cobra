FROM golang:1.23-alpine

WORKDIR /app
RUN apk update && apk add git
COPY . .
RUN go mod download