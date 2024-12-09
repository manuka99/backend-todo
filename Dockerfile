# Stage 1: Build the Go application
FROM golang:1.21.0 AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Create a lightweight image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Ensure it has execution permissions
RUN chmod +x ./main

# Expose the port
EXPOSE 8080

# Run the binary
CMD ["./main"]
