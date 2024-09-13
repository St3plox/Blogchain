# Use the official Golang image with the correct version
FROM golang:1.22.5 AS build

# Set the working directory inside the container
WORKDIR /app/backend/user-service

# Copy the Go modules and build files
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Change to the directory containing the main.go
WORKDIR /app/backend/user-service

# Build the Go app
RUN go build -o myapp -ldflags '-linkmode external -w -extldflags "-static"' ./app/backend/user-service

# Start a new stage from Alpine
FROM alpine:3.19

# Set necessary environment variables
ENV INTERNAL_PORT=3000
ENV EXPOSE_PORT=3000

# Set the working directory inside the container
WORKDIR /app/backend/user-service

# Copy the built executable from the previous stage
COPY --from=build /app/backend/user-service/myapp .

# Copy the keys folder from the build context
COPY zarf/keys ./zarf/keys

# Expose the port the app runs on
EXPOSE $EXPOSE_PORT
EXPOSE 4000

# Command to run the executable
CMD ["./myapp"]
