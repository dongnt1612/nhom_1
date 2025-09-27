FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy Go module files
COPY go.mod ./

# Download dependencies (if any)
RUN go mod download

# Copy source code
COPY backend/ ./backend/
COPY frontend/ ./frontend/

# Build the Go binary
RUN cd backend && go build -o ../device-management main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the binary and frontend files
COPY --from=builder /app/device-management .
COPY --from=builder /app/frontend/ ./frontend/

# Expose port
EXPOSE 8080

# Start the server
CMD ["./device-management"]
