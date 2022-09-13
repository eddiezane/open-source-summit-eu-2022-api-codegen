// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/eddiezane/open-source-summit-eu-2022-api-codegen/swagger/gen/models"
)

// GetImageReader is a Reader for the GetImage structure.
type GetImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetImageOK creates a GetImageOK with default headers values
func NewGetImageOK() *GetImageOK {
	return &GetImageOK{}
}

/*
GetImageOK describes a response with status code 200, with default header values.

OK
*/
type GetImageOK struct {
	Payload *models.Image
}

// IsSuccess returns true when this get image o k response has a 2xx status code
func (o *GetImageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get image o k response has a 3xx status code
func (o *GetImageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get image o k response has a 4xx status code
func (o *GetImageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get image o k response has a 5xx status code
func (o *GetImageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get image o k response a status code equal to that given
func (o *GetImageOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetImageOK) Error() string {
	return fmt.Sprintf("[GET /images/{id}][%d] getImageOK  %+v", 200, o.Payload)
}

func (o *GetImageOK) String() string {
	return fmt.Sprintf("[GET /images/{id}][%d] getImageOK  %+v", 200, o.Payload)
}

func (o *GetImageOK) GetPayload() *models.Image {
	return o.Payload
}

func (o *GetImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
