// Code generated by go-swagger; DO NOT EDIT.

package supply

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

// NewGetSupplyTotalDenominationParams creates a new GetSupplyTotalDenominationParams object
// with the default values initialized.
func NewGetSupplyTotalDenominationParams() *GetSupplyTotalDenominationParams {
	var ()
	return &GetSupplyTotalDenominationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSupplyTotalDenominationParamsWithTimeout creates a new GetSupplyTotalDenominationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSupplyTotalDenominationParamsWithTimeout(timeout time.Duration) *GetSupplyTotalDenominationParams {
	var ()
	return &GetSupplyTotalDenominationParams{

		timeout: timeout,
	}
}

// NewGetSupplyTotalDenominationParamsWithContext creates a new GetSupplyTotalDenominationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSupplyTotalDenominationParamsWithContext(ctx context.Context) *GetSupplyTotalDenominationParams {
	var ()
	return &GetSupplyTotalDenominationParams{

		Context: ctx,
	}
}

// NewGetSupplyTotalDenominationParamsWithHTTPClient creates a new GetSupplyTotalDenominationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSupplyTotalDenominationParamsWithHTTPClient(client *http.Client) *GetSupplyTotalDenominationParams {
	var ()
	return &GetSupplyTotalDenominationParams{
		HTTPClient: client,
	}
}

/*GetSupplyTotalDenominationParams contains all the parameters to send to the API endpoint
for the get supply total denomination operation typically these are written to a http.Request
*/
type GetSupplyTotalDenominationParams struct {

	/*Denomination
	  Coin denomination

	*/
	Denomination string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get supply total denomination params
func (o *GetSupplyTotalDenominationParams) WithTimeout(timeout time.Duration) *GetSupplyTotalDenominationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get supply total denomination params
func (o *GetSupplyTotalDenominationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get supply total denomination params
func (o *GetSupplyTotalDenominationParams) WithContext(ctx context.Context) *GetSupplyTotalDenominationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get supply total denomination params
func (o *GetSupplyTotalDenominationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get supply total denomination params
func (o *GetSupplyTotalDenominationParams) WithHTTPClient(client *http.Client) *GetSupplyTotalDenominationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get supply total denomination params
func (o *GetSupplyTotalDenominationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDenomination adds the denomination to the get supply total denomination params
func (o *GetSupplyTotalDenominationParams) WithDenomination(denomination string) *GetSupplyTotalDenominationParams {
	o.SetDenomination(denomination)
	return o
}

// SetDenomination adds the denomination to the get supply total denomination params
func (o *GetSupplyTotalDenominationParams) SetDenomination(denomination string) {
	o.Denomination = denomination
}

// WriteToRequest writes these params to a swagger request
func (o *GetSupplyTotalDenominationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param denomination
	if err := r.SetPathParam("denomination", o.Denomination); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
