FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o band-service ./cmd/main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/band-service .

ENV CONFIG_PATH="/app/config.yaml"

EXPOSE 8082

CMD ["./band-service"]
