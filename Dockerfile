# Start from a Golang base image
FROM golang:latest as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code to the working directory
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o notification-subscriber .

# Use a minimal base image for the final build

#FROM gcr.io/distroless/base-debian11 AS server
FROM alpine:latest AS server

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage to the current directory in the container
COPY --from=builder /app/notification-subscriber .

# Copy the .env file to the container
COPY .env .
COPY key.json .

ENV TZ=Asia/Dhaka

# Command to run the application
CMD ["./notification-subscriber"]

EXPOSE 3000