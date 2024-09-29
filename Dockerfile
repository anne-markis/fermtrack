FROM golang:1.23-alpine

# Copy your Go application files
COPY . /app

# Set the working directory
WORKDIR /app

# Install dependencies
RUN go mod download

# Build the Go application
RUN go build -o fermtrack .

# Expose the port your HTTP service will listen on
EXPOSE 8080

# Command to run when the container starts
CMD ["/app/fermtrack"]