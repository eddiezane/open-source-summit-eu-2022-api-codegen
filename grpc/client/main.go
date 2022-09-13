package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	txt2imgv1 "github.com/eddiezane/open-source-summit-eu-2022-api-codegen/grpc/client/txt2img/v1"
)

func main() {
	host := os.Getenv("TXT2IMG_HOST")
	if host == "" {
		panic(errors.New("TXT2IMG_HOST required"))
	}

	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := txt2imgv1.NewTxt2ImgServiceClient(conn)

	gir, err := client.GenerateImage(context.Background(), &txt2imgv1.GenerateImageRequest{Prompt: "a roller disco"})
	if err != nil {
		panic(err)
	}
	fmt.Println(gir.GetImage().GetUrl())

	lir, err := client.ListImages(context.Background(), &txt2imgv1.ListImagesRequest{})
	if err != nil {
		panic(err)
	}

	for _, img := range lir.Images {
		client.DeleteImage(context.Background(), &txt2imgv1.DeleteImageRequest{Id: img.GetId()})
	}
}
