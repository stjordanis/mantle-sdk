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

// NewGetDistributionValidatorsValidatorAddrParams creates a new GetDistributionValidatorsValidatorAddrParams object
// with the default values initialized.
func NewGetDistributionValidatorsValidatorAddrParams() *GetDistributionValidatorsValidatorAddrParams {
	var ()
	return &GetDistributionValidatorsValidatorAddrParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDistributionValidatorsValidatorAddrParamsWithTimeout creates a new GetDistributionValidatorsValidatorAddrParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDistributionValidatorsValidatorAddrParamsWithTimeout(timeout time.Duration) *GetDistributionValidatorsValidatorAddrParams {
	var ()
	return &GetDistributionValidatorsValidatorAddrParams{

		timeout: timeout,
	}
}

// NewGetDistributionValidatorsValidatorAddrParamsWithContext creates a new GetDistributionValidatorsValidatorAddrParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDistributionValidatorsValidatorAddrParamsWithContext(ctx context.Context) *GetDistributionValidatorsValidatorAddrParams {
	var ()
	return &GetDistributionValidatorsValidatorAddrParams{

		Context: ctx,
	}
}

// NewGetDistributionValidatorsValidatorAddrParamsWithHTTPClient creates a new GetDistributionValidatorsValidatorAddrParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDistributionValidatorsValidatorAddrParamsWithHTTPClient(client *http.Client) *GetDistributionValidatorsValidatorAddrParams {
	var ()
	return &GetDistributionValidatorsValidatorAddrParams{
		HTTPClient: client,
	}
}

/*GetDistributionValidatorsValidatorAddrParams contains all the parameters to send to the API endpoint
for the get distribution validators validator addr operation typically these are written to a http.Request
*/
type GetDistributionValidatorsValidatorAddrParams struct {

	/*ValidatorAddr
	  Bech32 OperatorAddress of validator

	*/
	ValidatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get distribution validators validator addr params
func (o *GetDistributionValidatorsValidatorAddrParams) WithTimeout(timeout time.Duration) *GetDistributionValidatorsValidatorAddrParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get distribution validators validator addr params
func (o *GetDistributionValidatorsValidatorAddrParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get distribution validators validator addr params
func (o *GetDistributionValidatorsValidatorAddrParams) WithContext(ctx context.Context) *GetDistributionValidatorsValidatorAddrParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get distribution validators validator addr params
func (o *GetDistributionValidatorsValidatorAddrParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get distribution validators validator addr params
func (o *GetDistributionValidatorsValidatorAddrParams) WithHTTPClient(client *http.Client) *GetDistributionValidatorsValidatorAddrParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get distribution validators validator addr params
func (o *GetDistributionValidatorsValidatorAddrParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValidatorAddr adds the validatorAddr to the get distribution validators validator addr params
func (o *GetDistributionValidatorsValidatorAddrParams) WithValidatorAddr(validatorAddr string) *GetDistributionValidatorsValidatorAddrParams {
	o.SetValidatorAddr(validatorAddr)
	return o
}

// SetValidatorAddr adds the validatorAddr to the get distribution validators validator addr params
func (o *GetDistributionValidatorsValidatorAddrParams) SetValidatorAddr(validatorAddr string) {
	o.ValidatorAddr = validatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *GetDistributionValidatorsValidatorAddrParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param validatorAddr
	if err := r.SetPathParam("validatorAddr", o.ValidatorAddr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
