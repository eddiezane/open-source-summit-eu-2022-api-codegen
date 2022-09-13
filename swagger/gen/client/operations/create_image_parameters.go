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

// NewCreateImageParams creates a new CreateImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateImageParams() *CreateImageParams {
	return &CreateImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateImageParamsWithTimeout creates a new CreateImageParams object
// with the ability to set a timeout on a request.
func NewCreateImageParamsWithTimeout(timeout time.Duration) *CreateImageParams {
	return &CreateImageParams{
		timeout: timeout,
	}
}

// NewCreateImageParamsWithContext creates a new CreateImageParams object
// with the ability to set a context for a request.
func NewCreateImageParamsWithContext(ctx context.Context) *CreateImageParams {
	return &CreateImageParams{
		Context: ctx,
	}
}

// NewCreateImageParamsWithHTTPClient creates a new CreateImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateImageParamsWithHTTPClient(client *http.Client) *CreateImageParams {
	return &CreateImageParams{
		HTTPClient: client,
	}
}

/*
CreateImageParams contains all the parameters to send to the API endpoint

	for the create image operation.

	Typically these are written to a http.Request.
*/
type CreateImageParams struct {

	// Prompt.
	Prompt *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateImageParams) WithDefaults() *CreateImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create image params
func (o *CreateImageParams) WithTimeout(timeout time.Duration) *CreateImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create image params
func (o *CreateImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create image params
func (o *CreateImageParams) WithContext(ctx context.Context) *CreateImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create image params
func (o *CreateImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create image params
func (o *CreateImageParams) WithHTTPClient(client *http.Client) *CreateImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create image params
func (o *CreateImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPrompt adds the prompt to the create image params
func (o *CreateImageParams) WithPrompt(prompt *string) *CreateImageParams {
	o.SetPrompt(prompt)
	return o
}

// SetPrompt adds the prompt to the create image params
func (o *CreateImageParams) SetPrompt(prompt *string) {
	o.Prompt = prompt
}

// WriteToRequest writes these params to a swagger request
func (o *CreateImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Prompt != nil {

		// form param prompt
		var frPrompt string
		if o.Prompt != nil {
			frPrompt = *o.Prompt
		}
		fPrompt := frPrompt
		if fPrompt != "" {
			if err := r.SetFormParam("prompt", fPrompt); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
