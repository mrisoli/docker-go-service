# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Marcelo Risoli<celorisoli@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . /app

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
