# Use Golang image
FROM golang:1.23-alpine as builder

WORKDIR /app

# Copy the Go module files and download the dependencies
#COPY go.mod go.sum ./
#RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o main .

# Use a smaller image for the final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the compiled binary from the builder image
COPY --from=builder /app/main .

# Copy .env file into the container
COPY .env .env

# Expose the port (if necessary)
EXPOSE 8080

# Command to run the application
CMD ["./main"]
