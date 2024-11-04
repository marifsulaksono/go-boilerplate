# Use official Golang image for building the app
FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Copy .env file
COPY .env ./.env

# Install certificates and build the application
RUN apk --no-cache add ca-certificates && \
    go build -ldflags="-w -s -extldflags '-static'" -o myapp ./cmd/api

# Build the application
RUN go build -o myapp ./cmd/api

# Start a new minimal container for runtime
FROM alpine:latest

# Set working directory for the runtime container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/myapp .
COPY --from=builder /app/.env .env

# Expose the application's port
EXPOSE 8080

# Run the application
CMD ["./myapp"]
