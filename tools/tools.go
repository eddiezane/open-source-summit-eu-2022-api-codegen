//go:build tools
// +build tools

package tools

import (
	_ "github.com/deepmap/oapi-codegen/cmd/oapi-codegen"
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
