# 1. Build Stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy dependency files first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary (named "main")
RUN go build -o main main.go

# 2. Run Stage (Small image)
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env . 
# Note: .env usually isn't copied in prod, but fine for now if you need it.

# Expose port (Cloud usually ignores this but good for documentation)
EXPOSE 8000

# Run the app
CMD ["./main"]