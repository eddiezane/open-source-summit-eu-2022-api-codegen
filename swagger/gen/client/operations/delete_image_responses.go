// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteImageReader is a Reader for the DeleteImage structure.
type DeleteImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteImageOK creates a DeleteImageOK with default headers values
func NewDeleteImageOK() *DeleteImageOK {
	return &DeleteImageOK{}
}

/*
DeleteImageOK describes a response with status code 200, with default header values.

OK
*/
type DeleteImageOK struct {
}

// IsSuccess returns true when this delete image o k response has a 2xx status code
func (o *DeleteImageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete image o k response has a 3xx status code
func (o *DeleteImageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete image o k response has a 4xx status code
func (o *DeleteImageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete image o k response has a 5xx status code
func (o *DeleteImageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete image o k response a status code equal to that given
func (o *DeleteImageOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteImageOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/images/{id}][%d] deleteImageOK ", 200)
}

func (o *DeleteImageOK) String() string {
	return fmt.Sprintf("[DELETE /v1/images/{id}][%d] deleteImageOK ", 200)
}

func (o *DeleteImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}