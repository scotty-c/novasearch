# Stage 1: Build the Golang binary
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first (for better layer caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /novasearch ./cmd/novasearch

# Stage 2: Create a lightweight image to run the application
FROM alpine:3.18

# Set environment variables
ENV REDIS_ADDR=localhost:6379
ENV AWS_REGION=us-east-1

# Install any required dependencies (if any)
RUN apk --no-cache add ca-certificates

# Copy the Go binary from the builder stage
COPY --from=builder /novasearch /usr/local/bin/novasearch

# Expose the port on which your application listens
EXPOSE 8080

# Command to run the application
CMD ["/usr/local/bin/novasearch"]
