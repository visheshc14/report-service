FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install protoc and necessary tools
RUN apk add --no-cache protoc bash git

# Install Go protobuf plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

# Ensure protoc can find the Go plugins
ENV PATH="$PATH:$(go env GOPATH)/bin"

COPY go.mod go.sum ./
RUN go mod download

# Copy project files
COPY . .

# Generate protobuf code 
RUN sh generate.sh

# Build the Go binary statically
RUN CGO_ENABLED=0 go build -o report-service ./cmd/server

FROM alpine:3.18

WORKDIR /app

RUN apk add --no-cache libc6-compat

# Copy the Go binary from builder
COPY --from=builder /app/report-service .

# Expose gRPC port
EXPOSE 50051

# Run the application
CMD ["./report-service"]
