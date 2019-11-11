# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Urmatbek Bozov <ubozov@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o web-app .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./web-app"]

