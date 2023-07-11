FROM golang:latest

# Working directory
WORKDIR /app

# Copy everything at /app
COPY . /app

# Build the go app
RUN go build -o main ./cmd

# Expose port
EXPOSE 8080

# Define the command to run the app
CMD ["./main"]
