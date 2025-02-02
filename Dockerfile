# Stage 1: Build
FROM golang:1.23 AS builder

# Disable CGO to produce a fully static binary
ENV CGO_ENABLED=0

# Set the Go proxy to help with module resolution
ENV GOPROXY=https://proxy.golang.org,direct

# Set the working directory
WORKDIR /app

# Copy dependency files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary and place it in the /app/bin directory
RUN go build -o ./bin/go-modules-api main.go

# Stage 2: Final image for execution
FROM alpine:3.17

# Install necessary certificates
RUN apk add --no-cache ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the compiled binary and the .env file into the final image
COPY --from=builder /app/bin/go-modules-api /bin/go-modules-api
COPY --from=builder /app/.env .env

# Default command to run the binary
CMD ["/bin/go-modules-api"]