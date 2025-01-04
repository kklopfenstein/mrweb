# Start with the official Golang image
FROM golang:1.18 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build ./cmd/...

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/mrweb .

# Expose the port the application runs on
EXPOSE 80

# Command to run the executable
CMD ["./mrweb"]