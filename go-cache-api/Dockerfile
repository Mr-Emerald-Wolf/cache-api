# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local source code into the container's working directory
COPY . .

# Build the Golang application
RUN go build -o app .

# Set the command to run when the container starts
CMD ["./app"]

