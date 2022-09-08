#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Go
protoc -I protos --go_out=protos --go_opt=paths=source_relative --go-grpc_out=protos --go-grpc_opt=paths=source_relative protos/txt2img/v1/txt2img.proto

# Python
source venv/bin/activate
python -m grpc_tools.protoc -I protos --python_out=protos --grpc_python_out=protos protos/txt2img/v1/txt2img.proto
deactivate
