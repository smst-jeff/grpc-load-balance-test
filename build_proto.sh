#! /bin/sh
# Generates code for protobufs

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pb/ping.proto
