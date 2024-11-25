FROM golang:1.20

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go modules
RUN go mod tidy

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o consumer .

CMD ["./consumer"]
