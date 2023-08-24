# Dockerfile for Golang console project
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Download Go modules
# COPY go.mod go.sum ./

# Copy the source code from the current directory to the container
COPY . .

# Build the Go application inside the container
# RUN go mod download
RUN go build -o app

# Set the command to run the executable
CMD ["./app"]
