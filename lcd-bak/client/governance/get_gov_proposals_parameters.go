// Code generated by go-swagger; DO NOT EDIT.

package governance

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

// NewGetGovProposalsParams creates a new GetGovProposalsParams object
// with the default values initialized.
func NewGetGovProposalsParams() *GetGovProposalsParams {
	var ()
	return &GetGovProposalsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGovProposalsParamsWithTimeout creates a new GetGovProposalsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGovProposalsParamsWithTimeout(timeout time.Duration) *GetGovProposalsParams {
	var ()
	return &GetGovProposalsParams{

		timeout: timeout,
	}
}

// NewGetGovProposalsParamsWithContext creates a new GetGovProposalsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGovProposalsParamsWithContext(ctx context.Context) *GetGovProposalsParams {
	var ()
	return &GetGovProposalsParams{

		Context: ctx,
	}
}

// NewGetGovProposalsParamsWithHTTPClient creates a new GetGovProposalsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGovProposalsParamsWithHTTPClient(client *http.Client) *GetGovProposalsParams {
	var ()
	return &GetGovProposalsParams{
		HTTPClient: client,
	}
}

/*GetGovProposalsParams contains all the parameters to send to the API endpoint
for the get gov proposals operation typically these are written to a http.Request
*/
type GetGovProposalsParams struct {

	/*Depositor
	  depositor address

	*/
	Depositor *string
	/*Status
	  proposal status, valid values can be `"deposit_period"`, `"voting_period"`, `"passed"`, `"rejected"`

	*/
	Status *string
	/*Voter
	  voter address

	*/
	Voter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gov proposals params
func (o *GetGovProposalsParams) WithTimeout(timeout time.Duration) *GetGovProposalsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gov proposals params
func (o *GetGovProposalsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gov proposals params
func (o *GetGovProposalsParams) WithContext(ctx context.Context) *GetGovProposalsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gov proposals params
func (o *GetGovProposalsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gov proposals params
func (o *GetGovProposalsParams) WithHTTPClient(client *http.Client) *GetGovProposalsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gov proposals params
func (o *GetGovProposalsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDepositor adds the depositor to the get gov proposals params
func (o *GetGovProposalsParams) WithDepositor(depositor *string) *GetGovProposalsParams {
	o.SetDepositor(depositor)
	return o
}

// SetDepositor adds the depositor to the get gov proposals params
func (o *GetGovProposalsParams) SetDepositor(depositor *string) {
	o.Depositor = depositor
}

// WithStatus adds the status to the get gov proposals params
func (o *GetGovProposalsParams) WithStatus(status *string) *GetGovProposalsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get gov proposals params
func (o *GetGovProposalsParams) SetStatus(status *string) {
	o.Status = status
}

// WithVoter adds the voter to the get gov proposals params
func (o *GetGovProposalsParams) WithVoter(voter *string) *GetGovProposalsParams {
	o.SetVoter(voter)
	return o
}

// SetVoter adds the voter to the get gov proposals params
func (o *GetGovProposalsParams) SetVoter(voter *string) {
	o.Voter = voter
}

// WriteToRequest writes these params to a swagger request
func (o *GetGovProposalsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Depositor != nil {

		// query param depositor
		var qrDepositor string
		if o.Depositor != nil {
			qrDepositor = *o.Depositor
		}
		qDepositor := qrDepositor
		if qDepositor != "" {
			if err := r.SetQueryParam("depositor", qDepositor); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.Voter != nil {

		// query param voter
		var qrVoter string
		if o.Voter != nil {
			qrVoter = *o.Voter
		}
		qVoter := qrVoter
		if qVoter != "" {
			if err := r.SetQueryParam("voter", qVoter); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
