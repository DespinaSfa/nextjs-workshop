FROM golang:1.22.2

WORKDIR /app

# Install dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the source code
COPY . ./

# Expose the backend server port
EXPOSE 3001

# Start the Go server
CMD ["go", "run", "main.go"]