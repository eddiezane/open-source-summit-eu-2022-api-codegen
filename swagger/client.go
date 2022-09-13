package main

import (
	"errors"
	"fmt"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/eddiezane/open-source-summit-eu-2022-api-codegen/swagger/gen/client"
	"github.com/eddiezane/open-source-summit-eu-2022-api-codegen/swagger/gen/client/operations"
)

func main() {
	host := os.Getenv("TXT2IMG_HOST")
	if host == "" {
		panic(errors.New("TXT2IMG_HOST must be defined"))
	}

	transport := httptransport.New(host, "", nil)

	txt2imgClient := client.New(transport, strfmt.Default)

	// req := operations.NewCreateImageParams()
	// req.Body = operations.CreateImageBody{Prompt: "A big yellow dog"}
	// res, err := txt2imgClient.Operations.CreateImage(req)
	// if err != nil {
	// panic(err)
	// }

	// if !res.IsSuccess() {
	// panic(errors.New("Something went wrong"))
	// }

	// fmt.Println(res.GetPayload())

	res, err := txt2imgClient.Operations.ListImages(nil)
	if err != nil {
		panic(err)
	}
	for _, image := range res.GetPayload().Images {
		req := operations.NewDeleteImageParams()
		req.ID = image.ID
		res, err := txt2imgClient.Operations.DeleteImage(req)
		if err != nil {
			panic(err)
		}
		fmt.Println(res.IsSuccess())
	}
}
