FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

RUN go install github.com/air-verse/air@latest

ENTRYPOINT ["air", "server"]
EXPOSE 3333