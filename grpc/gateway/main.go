package main

import (
	"context"
	"log"
	"net/http"

	gwpb "github.com/eddiezane/open-source-summit-eu-2022-api-codegen/grpc/gateway/txt2img/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), "localhost:8080", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	gwmux := runtime.NewServeMux()

	err = gwpb.RegisterTxt2ImgServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		panic(err)
	}

	gwServer := &http.Server{
		Addr:    ":8081",
		Handler: gwmux,
	}

	log.Println("starting grpc gateway")
	log.Fatalln(gwServer.ListenAndServe())
}
