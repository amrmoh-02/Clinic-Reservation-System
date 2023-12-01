# Use an official Golang runtime as a parent image
FROM golang:1.21.4

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
ADD  .  .

# Install any needed packages specified in go.mod and go.sum
RUN go mod download

# Build the backend binary
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]