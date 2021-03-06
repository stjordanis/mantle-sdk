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

// NewPostGovProposalsCommunityPoolSpendParams creates a new PostGovProposalsCommunityPoolSpendParams object
// with the default values initialized.
func NewPostGovProposalsCommunityPoolSpendParams() *PostGovProposalsCommunityPoolSpendParams {
	var ()
	return &PostGovProposalsCommunityPoolSpendParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostGovProposalsCommunityPoolSpendParamsWithTimeout creates a new PostGovProposalsCommunityPoolSpendParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostGovProposalsCommunityPoolSpendParamsWithTimeout(timeout time.Duration) *PostGovProposalsCommunityPoolSpendParams {
	var ()
	return &PostGovProposalsCommunityPoolSpendParams{

		timeout: timeout,
	}
}

// NewPostGovProposalsCommunityPoolSpendParamsWithContext creates a new PostGovProposalsCommunityPoolSpendParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostGovProposalsCommunityPoolSpendParamsWithContext(ctx context.Context) *PostGovProposalsCommunityPoolSpendParams {
	var ()
	return &PostGovProposalsCommunityPoolSpendParams{

		Context: ctx,
	}
}

// NewPostGovProposalsCommunityPoolSpendParamsWithHTTPClient creates a new PostGovProposalsCommunityPoolSpendParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostGovProposalsCommunityPoolSpendParamsWithHTTPClient(client *http.Client) *PostGovProposalsCommunityPoolSpendParams {
	var ()
	return &PostGovProposalsCommunityPoolSpendParams{
		HTTPClient: client,
	}
}

/*PostGovProposalsCommunityPoolSpendParams contains all the parameters to send to the API endpoint
for the post gov proposals community pool spend operation typically these are written to a http.Request
*/
type PostGovProposalsCommunityPoolSpendParams struct {

	/*PostProposalBody
	  The community pool spend proposal body that contains recipient and reward info

	*/
	PostProposalBody PostGovProposalsCommunityPoolSpendBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post gov proposals community pool spend params
func (o *PostGovProposalsCommunityPoolSpendParams) WithTimeout(timeout time.Duration) *PostGovProposalsCommunityPoolSpendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post gov proposals community pool spend params
func (o *PostGovProposalsCommunityPoolSpendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post gov proposals community pool spend params
func (o *PostGovProposalsCommunityPoolSpendParams) WithContext(ctx context.Context) *PostGovProposalsCommunityPoolSpendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post gov proposals community pool spend params
func (o *PostGovProposalsCommunityPoolSpendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post gov proposals community pool spend params
func (o *PostGovProposalsCommunityPoolSpendParams) WithHTTPClient(client *http.Client) *PostGovProposalsCommunityPoolSpendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post gov proposals community pool spend params
func (o *PostGovProposalsCommunityPoolSpendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostProposalBody adds the postProposalBody to the post gov proposals community pool spend params
func (o *PostGovProposalsCommunityPoolSpendParams) WithPostProposalBody(postProposalBody PostGovProposalsCommunityPoolSpendBody) *PostGovProposalsCommunityPoolSpendParams {
	o.SetPostProposalBody(postProposalBody)
	return o
}

// SetPostProposalBody adds the postProposalBody to the post gov proposals community pool spend params
func (o *PostGovProposalsCommunityPoolSpendParams) SetPostProposalBody(postProposalBody PostGovProposalsCommunityPoolSpendBody) {
	o.PostProposalBody = postProposalBody
}

// WriteToRequest writes these params to a swagger request
func (o *PostGovProposalsCommunityPoolSpendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.PostProposalBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
