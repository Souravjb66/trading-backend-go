# Stage 1: Build the binary
FROM golang:1.23-alpine AS builder

# Install git (required for fetching some Go dependencies)
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
# Using CGO_ENABLED=0 ensures the binary is statically linked
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Final lightweight image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy .env or config files if your app needs them
# COPY .env .

# Expose the ports for documentation
EXPOSE 8080
EXPOSE 8081

# Run the binary
CMD ["./main"]