package main

import (
	"errors"
	"fmt"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

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

	res, err := txt2imgClient.Operations.CreateImage(&operations.CreateImageParams{Prompt: swag.String("a big red dog")})
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Payload)
}
