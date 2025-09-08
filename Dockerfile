# Build stage
FROM golang:1.23.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go

# Runtime stage
FROM ubuntu:22.04

WORKDIR /app

# کپی فایل اجرایی Go
COPY --from=builder /app/server .

# کپی پوشه‌ی فرانت برای سرویس‌دهی استاتیک
COPY --from=builder /app/frontend /app/frontend

# ✅ کپی پوشه‌ی docs برای Swagger
COPY --from=builder /app/docs /app/docs

EXPOSE 8080

CMD ["./server"]
