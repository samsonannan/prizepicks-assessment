# Build the application from source
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container's working directory
COPY . .

# Build the Go application inside the container
RUN go build -o app

EXPOSE 8080

# Set the command to run the binary when the container starts
CMD ["./app"]