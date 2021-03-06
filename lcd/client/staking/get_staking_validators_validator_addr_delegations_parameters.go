// Code generated by go-swagger; DO NOT EDIT.

package staking

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

// NewGetStakingValidatorsValidatorAddrDelegationsParams creates a new GetStakingValidatorsValidatorAddrDelegationsParams object
// with the default values initialized.
func NewGetStakingValidatorsValidatorAddrDelegationsParams() *GetStakingValidatorsValidatorAddrDelegationsParams {
	var ()
	return &GetStakingValidatorsValidatorAddrDelegationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStakingValidatorsValidatorAddrDelegationsParamsWithTimeout creates a new GetStakingValidatorsValidatorAddrDelegationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStakingValidatorsValidatorAddrDelegationsParamsWithTimeout(timeout time.Duration) *GetStakingValidatorsValidatorAddrDelegationsParams {
	var ()
	return &GetStakingValidatorsValidatorAddrDelegationsParams{

		timeout: timeout,
	}
}

// NewGetStakingValidatorsValidatorAddrDelegationsParamsWithContext creates a new GetStakingValidatorsValidatorAddrDelegationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStakingValidatorsValidatorAddrDelegationsParamsWithContext(ctx context.Context) *GetStakingValidatorsValidatorAddrDelegationsParams {
	var ()
	return &GetStakingValidatorsValidatorAddrDelegationsParams{

		Context: ctx,
	}
}

// NewGetStakingValidatorsValidatorAddrDelegationsParamsWithHTTPClient creates a new GetStakingValidatorsValidatorAddrDelegationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStakingValidatorsValidatorAddrDelegationsParamsWithHTTPClient(client *http.Client) *GetStakingValidatorsValidatorAddrDelegationsParams {
	var ()
	return &GetStakingValidatorsValidatorAddrDelegationsParams{
		HTTPClient: client,
	}
}

/*GetStakingValidatorsValidatorAddrDelegationsParams contains all the parameters to send to the API endpoint
for the get staking validators validator addr delegations operation typically these are written to a http.Request
*/
type GetStakingValidatorsValidatorAddrDelegationsParams struct {

	/*ValidatorAddr
	  Bech32 OperatorAddress of validator

	*/
	ValidatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get staking validators validator addr delegations params
func (o *GetStakingValidatorsValidatorAddrDelegationsParams) WithTimeout(timeout time.Duration) *GetStakingValidatorsValidatorAddrDelegationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get staking validators validator addr delegations params
func (o *GetStakingValidatorsValidatorAddrDelegationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get staking validators validator addr delegations params
func (o *GetStakingValidatorsValidatorAddrDelegationsParams) WithContext(ctx context.Context) *GetStakingValidatorsValidatorAddrDelegationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get staking validators validator addr delegations params
func (o *GetStakingValidatorsValidatorAddrDelegationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get staking validators validator addr delegations params
func (o *GetStakingValidatorsValidatorAddrDelegationsParams) WithHTTPClient(client *http.Client) *GetStakingValidatorsValidatorAddrDelegationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get staking validators validator addr delegations params
func (o *GetStakingValidatorsValidatorAddrDelegationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValidatorAddr adds the validatorAddr to the get staking validators validator addr delegations params
func (o *GetStakingValidatorsValidatorAddrDelegationsParams) WithValidatorAddr(validatorAddr string) *GetStakingValidatorsValidatorAddrDelegationsParams {
	o.SetValidatorAddr(validatorAddr)
	return o
}

// SetValidatorAddr adds the validatorAddr to the get staking validators validator addr delegations params
func (o *GetStakingValidatorsValidatorAddrDelegationsParams) SetValidatorAddr(validatorAddr string) {
	o.ValidatorAddr = validatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *GetStakingValidatorsValidatorAddrDelegationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
