// Code generated by go-swagger; DO NOT EDIT.

package oracle

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

// NewGetOracleVotersValidatorFeederParams creates a new GetOracleVotersValidatorFeederParams object
// with the default values initialized.
func NewGetOracleVotersValidatorFeederParams() *GetOracleVotersValidatorFeederParams {
	var ()
	return &GetOracleVotersValidatorFeederParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOracleVotersValidatorFeederParamsWithTimeout creates a new GetOracleVotersValidatorFeederParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOracleVotersValidatorFeederParamsWithTimeout(timeout time.Duration) *GetOracleVotersValidatorFeederParams {
	var ()
	return &GetOracleVotersValidatorFeederParams{

		timeout: timeout,
	}
}

// NewGetOracleVotersValidatorFeederParamsWithContext creates a new GetOracleVotersValidatorFeederParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOracleVotersValidatorFeederParamsWithContext(ctx context.Context) *GetOracleVotersValidatorFeederParams {
	var ()
	return &GetOracleVotersValidatorFeederParams{

		Context: ctx,
	}
}

// NewGetOracleVotersValidatorFeederParamsWithHTTPClient creates a new GetOracleVotersValidatorFeederParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOracleVotersValidatorFeederParamsWithHTTPClient(client *http.Client) *GetOracleVotersValidatorFeederParams {
	var ()
	return &GetOracleVotersValidatorFeederParams{
		HTTPClient: client,
	}
}

/*GetOracleVotersValidatorFeederParams contains all the parameters to send to the API endpoint
for the get oracle voters validator feeder operation typically these are written to a http.Request
*/
type GetOracleVotersValidatorFeederParams struct {

	/*Validator
	  Feeder right delegator

	*/
	Validator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get oracle voters validator feeder params
func (o *GetOracleVotersValidatorFeederParams) WithTimeout(timeout time.Duration) *GetOracleVotersValidatorFeederParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get oracle voters validator feeder params
func (o *GetOracleVotersValidatorFeederParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get oracle voters validator feeder params
func (o *GetOracleVotersValidatorFeederParams) WithContext(ctx context.Context) *GetOracleVotersValidatorFeederParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get oracle voters validator feeder params
func (o *GetOracleVotersValidatorFeederParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get oracle voters validator feeder params
func (o *GetOracleVotersValidatorFeederParams) WithHTTPClient(client *http.Client) *GetOracleVotersValidatorFeederParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get oracle voters validator feeder params
func (o *GetOracleVotersValidatorFeederParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValidator adds the validator to the get oracle voters validator feeder params
func (o *GetOracleVotersValidatorFeederParams) WithValidator(validator string) *GetOracleVotersValidatorFeederParams {
	o.SetValidator(validator)
	return o
}

// SetValidator adds the validator to the get oracle voters validator feeder params
func (o *GetOracleVotersValidatorFeederParams) SetValidator(validator string) {
	o.Validator = validator
}

// WriteToRequest writes these params to a swagger request
func (o *GetOracleVotersValidatorFeederParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param validator
	if err := r.SetPathParam("validator", o.Validator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
