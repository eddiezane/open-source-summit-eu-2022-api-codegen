// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetImagesIDParams creates a new GetImagesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetImagesIDParams() *GetImagesIDParams {
	return &GetImagesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetImagesIDParamsWithTimeout creates a new GetImagesIDParams object
// with the ability to set a timeout on a request.
func NewGetImagesIDParamsWithTimeout(timeout time.Duration) *GetImagesIDParams {
	return &GetImagesIDParams{
		timeout: timeout,
	}
}

// NewGetImagesIDParamsWithContext creates a new GetImagesIDParams object
// with the ability to set a context for a request.
func NewGetImagesIDParamsWithContext(ctx context.Context) *GetImagesIDParams {
	return &GetImagesIDParams{
		Context: ctx,
	}
}

// NewGetImagesIDParamsWithHTTPClient creates a new GetImagesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetImagesIDParamsWithHTTPClient(client *http.Client) *GetImagesIDParams {
	return &GetImagesIDParams{
		HTTPClient: client,
	}
}

/*
GetImagesIDParams contains all the parameters to send to the API endpoint

	for the get images id operation.

	Typically these are written to a http.Request.
*/
type GetImagesIDParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get images id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetImagesIDParams) WithDefaults() *GetImagesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get images id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetImagesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get images id params
func (o *GetImagesIDParams) WithTimeout(timeout time.Duration) *GetImagesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get images id params
func (o *GetImagesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get images id params
func (o *GetImagesIDParams) WithContext(ctx context.Context) *GetImagesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get images id params
func (o *GetImagesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get images id params
func (o *GetImagesIDParams) WithHTTPClient(client *http.Client) *GetImagesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get images id params
func (o *GetImagesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get images id params
func (o *GetImagesIDParams) WithID(id string) *GetImagesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get images id params
func (o *GetImagesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetImagesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
