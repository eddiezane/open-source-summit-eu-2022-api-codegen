// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteImagesIDReader is a Reader for the DeleteImagesID structure.
type DeleteImagesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteImagesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteImagesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteImagesIDOK creates a DeleteImagesIDOK with default headers values
func NewDeleteImagesIDOK() *DeleteImagesIDOK {
	return &DeleteImagesIDOK{}
}

/*
DeleteImagesIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteImagesIDOK struct {
}

// IsSuccess returns true when this delete images Id o k response has a 2xx status code
func (o *DeleteImagesIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete images Id o k response has a 3xx status code
func (o *DeleteImagesIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete images Id o k response has a 4xx status code
func (o *DeleteImagesIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete images Id o k response has a 5xx status code
func (o *DeleteImagesIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete images Id o k response a status code equal to that given
func (o *DeleteImagesIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteImagesIDOK) Error() string {
	return fmt.Sprintf("[DELETE /images/{id}][%d] deleteImagesIdOK ", 200)
}

func (o *DeleteImagesIDOK) String() string {
	return fmt.Sprintf("[DELETE /images/{id}][%d] deleteImagesIdOK ", 200)
}

func (o *DeleteImagesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
