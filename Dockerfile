# Use Go 1.22 with Alpine Linux
FROM golang:1.22-alpine

# Install necessary dependencies
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy everything to the container
COPY . .

# Download dependencies
RUN go mod tidy

# Build the application
RUN go build -o user-auth-service ./cmd/main.go

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./user-auth-service"]
