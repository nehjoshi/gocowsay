FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o gocowsay ./cmd/gocowsay

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/gocowsay .

ENTRYPOINT ["./gocowsay"]
