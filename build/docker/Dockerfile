FROM golang:latest AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y make

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/search-engine ./cmd

FROM debian:latest

WORKDIR /root/

COPY --from=builder /app/search-engine .

RUN chmod +x /root/search-engine

CMD ["./search-engine"]