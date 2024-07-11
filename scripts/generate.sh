#!/bin/bash
set -e

# Generate gRPC and HTTP gateway code
protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. api/proto/*.proto

# Generate SQL models using XO
xo sql/schema.sql -o internal/domain

# Copy the generated files to the appropriate directories

