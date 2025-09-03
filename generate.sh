#!/bin/bash

# Create gen directory if it doesn't exist
mkdir -p gen

# Generate Go code from protobuf
protoc --go_out=gen --go_opt=paths=source_relative \
    --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
    proto/report.proto

echo "Protobuf code generated in gen/"