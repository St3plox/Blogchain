# Use the official Golang image with the correct version
FROM golang:1.22.5 AS build

# Set the working directory inside the container
WORKDIR /app/backend/notification-service

# Copy the Go modules and build files
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Change to the directory containing the main.go
WORKDIR /app/backend/notification-service

# Build the Go app
RUN go build -o notification-service -ldflags '-linkmode external -w -extldflags "-static"' ./app/backend/notification-service

# Start a new stage from Alpine
FROM alpine:3.19

# Set necessary environment variables
ENV INTERNAL_PORT=3001
ENV EXPOSE_PORT=3001

# Set the working directory inside the container
WORKDIR /app/backend/notification-service

# Copy the built executable from the previous stage
COPY --from=build /app/backend/notification-service .

# Copy the keys folder from the build context
COPY app/backend/notification-service/config.json ./app/backend/notification-service/config.json

# Expose the port the app runs on
EXPOSE $EXPOSE_PORT
EXPOSE 4000

# Command to run the executable
CMD ["./notification-service"]
