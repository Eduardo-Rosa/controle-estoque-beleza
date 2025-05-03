# Dockerfile para a aplicação Go

# Etapa de build
FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main .

# Etapa final
FROM debian:bullseye-slim
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]