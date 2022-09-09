package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/eddiezane/open-source-summit-eu-2022-api-codegen/grpc/client/txt2img/v1"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("35.188.187.17:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := txt2imgv1.NewTxt2ImgServiceClient(conn)

	wg := sync.WaitGroup{}

	for _, p := range []string{"a big bad wolf", "the end of the world"} {
		wg.Add(1)
		go func(p string) {
			res, err := client.GenerateImage(context.Background(), &txt2imgv1.GenerateImageRequest{Prompt: p})
			if err != nil {
				fmt.Println(err)
			}
			url, err := client.GetImage(context.Background(), &txt2imgv1.GetImageRequest{Id: res.GetId()})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(url)
			wg.Done()
		}(p)
	}

	wg.Wait()
}
