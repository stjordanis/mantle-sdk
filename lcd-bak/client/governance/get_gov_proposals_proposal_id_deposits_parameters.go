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

// NewGetGovProposalsProposalIDDepositsParams creates a new GetGovProposalsProposalIDDepositsParams object
// with the default values initialized.
func NewGetGovProposalsProposalIDDepositsParams() *GetGovProposalsProposalIDDepositsParams {
	var ()
	return &GetGovProposalsProposalIDDepositsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGovProposalsProposalIDDepositsParamsWithTimeout creates a new GetGovProposalsProposalIDDepositsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGovProposalsProposalIDDepositsParamsWithTimeout(timeout time.Duration) *GetGovProposalsProposalIDDepositsParams {
	var ()
	return &GetGovProposalsProposalIDDepositsParams{

		timeout: timeout,
	}
}

// NewGetGovProposalsProposalIDDepositsParamsWithContext creates a new GetGovProposalsProposalIDDepositsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGovProposalsProposalIDDepositsParamsWithContext(ctx context.Context) *GetGovProposalsProposalIDDepositsParams {
	var ()
	return &GetGovProposalsProposalIDDepositsParams{

		Context: ctx,
	}
}

// NewGetGovProposalsProposalIDDepositsParamsWithHTTPClient creates a new GetGovProposalsProposalIDDepositsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGovProposalsProposalIDDepositsParamsWithHTTPClient(client *http.Client) *GetGovProposalsProposalIDDepositsParams {
	var ()
	return &GetGovProposalsProposalIDDepositsParams{
		HTTPClient: client,
	}
}

/*GetGovProposalsProposalIDDepositsParams contains all the parameters to send to the API endpoint
for the get gov proposals proposal ID deposits operation typically these are written to a http.Request
*/
type GetGovProposalsProposalIDDepositsParams struct {

	/*ProposalID*/
	ProposalID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gov proposals proposal ID deposits params
func (o *GetGovProposalsProposalIDDepositsParams) WithTimeout(timeout time.Duration) *GetGovProposalsProposalIDDepositsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gov proposals proposal ID deposits params
func (o *GetGovProposalsProposalIDDepositsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gov proposals proposal ID deposits params
func (o *GetGovProposalsProposalIDDepositsParams) WithContext(ctx context.Context) *GetGovProposalsProposalIDDepositsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gov proposals proposal ID deposits params
func (o *GetGovProposalsProposalIDDepositsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gov proposals proposal ID deposits params
func (o *GetGovProposalsProposalIDDepositsParams) WithHTTPClient(client *http.Client) *GetGovProposalsProposalIDDepositsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gov proposals proposal ID deposits params
func (o *GetGovProposalsProposalIDDepositsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProposalID adds the proposalID to the get gov proposals proposal ID deposits params
func (o *GetGovProposalsProposalIDDepositsParams) WithProposalID(proposalID string) *GetGovProposalsProposalIDDepositsParams {
	o.SetProposalID(proposalID)
	return o
}

// SetProposalID adds the proposalId to the get gov proposals proposal ID deposits params
func (o *GetGovProposalsProposalIDDepositsParams) SetProposalID(proposalID string) {
	o.ProposalID = proposalID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGovProposalsProposalIDDepositsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param proposalId
	if err := r.SetPathParam("proposalId", o.ProposalID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
