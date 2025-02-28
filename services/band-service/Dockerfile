FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o user-service ./cmd/main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/band-service .

ENV PORT=8081

EXPOSE 8081

CMD ["./band-service"]