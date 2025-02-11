# 1. Use the official Golang image
FROM golang:1.22.10

# 2. Set the working directory inside the container
WORKDIR /app

# 3. Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./

# 4. Download dependencies
RUN go mod tidy

# 5. Copy the entire project into the container
COPY . .

# 6. Build the Go application (targeting main.go inside `cmd/server/`)
RUN go build -o app ./cmd/server

# 7. Expose the port that your Go app runs on
EXPOSE 8080

# 8. Run the application
CMD ["./app"]
