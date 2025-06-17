# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application source
COPY . .

# Install necessary dependencies (e.g., certificates)
RUN apk --no-cache add ca-certificates

# Build the Go application statically
RUN go build -ldflags="-w -s -extldflags '-static'" -o myapp ./cmd/api

# Stage 2: Create a minimal runtime container
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy certificate bundle for HTTPS communication
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy built binary and .env file
COPY --from=builder /app/myapp .
COPY --from=builder /app/.env .env

# Expose port 8080
EXPOSE 8080

# Command to run the binary
CMD ["./myapp"]
