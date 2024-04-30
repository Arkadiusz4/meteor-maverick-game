# Use the official Golang image as a base
FROM golang:latest AS build

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files for dependency resolution
COPY go.mod go.sum ./

# Download dependencies (including specific versions)
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Set the working directory to the cmd directory
WORKDIR /app/cmd

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o game .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the executable from the build stage
COPY --from=build /app/cmd/game .

# Expose port 80 to the outside world
EXPOSE 80

# Command to run the executable
CMD ["./game"]
