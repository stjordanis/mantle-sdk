// Code generated by go-swagger; DO NOT EDIT.

package distribution

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

// NewGetDistributionCommunityPoolParams creates a new GetDistributionCommunityPoolParams object
// with the default values initialized.
func NewGetDistributionCommunityPoolParams() *GetDistributionCommunityPoolParams {

	return &GetDistributionCommunityPoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDistributionCommunityPoolParamsWithTimeout creates a new GetDistributionCommunityPoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDistributionCommunityPoolParamsWithTimeout(timeout time.Duration) *GetDistributionCommunityPoolParams {

	return &GetDistributionCommunityPoolParams{

		timeout: timeout,
	}
}

// NewGetDistributionCommunityPoolParamsWithContext creates a new GetDistributionCommunityPoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDistributionCommunityPoolParamsWithContext(ctx context.Context) *GetDistributionCommunityPoolParams {

	return &GetDistributionCommunityPoolParams{

		Context: ctx,
	}
}

// NewGetDistributionCommunityPoolParamsWithHTTPClient creates a new GetDistributionCommunityPoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDistributionCommunityPoolParamsWithHTTPClient(client *http.Client) *GetDistributionCommunityPoolParams {

	return &GetDistributionCommunityPoolParams{
		HTTPClient: client,
	}
}

/*GetDistributionCommunityPoolParams contains all the parameters to send to the API endpoint
for the get distribution community pool operation typically these are written to a http.Request
*/
type GetDistributionCommunityPoolParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get distribution community pool params
func (o *GetDistributionCommunityPoolParams) WithTimeout(timeout time.Duration) *GetDistributionCommunityPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get distribution community pool params
func (o *GetDistributionCommunityPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get distribution community pool params
func (o *GetDistributionCommunityPoolParams) WithContext(ctx context.Context) *GetDistributionCommunityPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get distribution community pool params
func (o *GetDistributionCommunityPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get distribution community pool params
func (o *GetDistributionCommunityPoolParams) WithHTTPClient(client *http.Client) *GetDistributionCommunityPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get distribution community pool params
func (o *GetDistributionCommunityPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDistributionCommunityPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
