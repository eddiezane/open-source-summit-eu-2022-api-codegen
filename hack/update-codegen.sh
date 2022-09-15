#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Go
# protoc -I protos --go_out=protos --go_opt=paths=source_relative --go-grpc_out=protos --go-grpc_opt=paths=source_relative protos/txt2img/v1/txt2img.proto

# Python
# source venv/bin/activate
# python -m grpc_tools.protoc -I protos --python_out=protos --grpc_python_out=protos protos/txt2img/v1/txt2img.proto
# deactivate

# Buf
buf generate

swagger generate client -t swagger/gen -f swagger/swagger.yaml

protoc -I grpc/simple/hello --go_out=grpc/simple/hello --go_opt=paths=source_relative --go-grpc_out=grpc/simple/hello --go-grpc_opt=paths=source_relative grpc/simple/hello/hello.proto

swagger generate client -t grpc/gateway/txt2img/v1 -f grpc/gateway/txt2img/v1/txt2img.swagger.json
