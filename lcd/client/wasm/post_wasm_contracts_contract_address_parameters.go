// Code generated by go-swagger; DO NOT EDIT.

package wasm

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

	"github.com/terra-project/mantle-sdk/lcd/models"
)

// NewPostWasmContractsContractAddressParams creates a new PostWasmContractsContractAddressParams object
// with the default values initialized.
func NewPostWasmContractsContractAddressParams() *PostWasmContractsContractAddressParams {
	var ()
	return &PostWasmContractsContractAddressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWasmContractsContractAddressParamsWithTimeout creates a new PostWasmContractsContractAddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWasmContractsContractAddressParamsWithTimeout(timeout time.Duration) *PostWasmContractsContractAddressParams {
	var ()
	return &PostWasmContractsContractAddressParams{

		timeout: timeout,
	}
}

// NewPostWasmContractsContractAddressParamsWithContext creates a new PostWasmContractsContractAddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWasmContractsContractAddressParamsWithContext(ctx context.Context) *PostWasmContractsContractAddressParams {
	var ()
	return &PostWasmContractsContractAddressParams{

		Context: ctx,
	}
}

// NewPostWasmContractsContractAddressParamsWithHTTPClient creates a new PostWasmContractsContractAddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWasmContractsContractAddressParamsWithHTTPClient(client *http.Client) *PostWasmContractsContractAddressParams {
	var ()
	return &PostWasmContractsContractAddressParams{
		HTTPClient: client,
	}
}

/*PostWasmContractsContractAddressParams contains all the parameters to send to the API endpoint
for the post wasm contracts contract address operation typically these are written to a http.Request
*/
type PostWasmContractsContractAddressParams struct {

	/*ContractAddress
	  contract address you want to execute

	*/
	ContractAddress string
	/*ExecuteContractRequestBody*/
	ExecuteContractRequestBody *models.ExecuteContractReq

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post wasm contracts contract address params
func (o *PostWasmContractsContractAddressParams) WithTimeout(timeout time.Duration) *PostWasmContractsContractAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post wasm contracts contract address params
func (o *PostWasmContractsContractAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post wasm contracts contract address params
func (o *PostWasmContractsContractAddressParams) WithContext(ctx context.Context) *PostWasmContractsContractAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post wasm contracts contract address params
func (o *PostWasmContractsContractAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post wasm contracts contract address params
func (o *PostWasmContractsContractAddressParams) WithHTTPClient(client *http.Client) *PostWasmContractsContractAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post wasm contracts contract address params
func (o *PostWasmContractsContractAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContractAddress adds the contractAddress to the post wasm contracts contract address params
func (o *PostWasmContractsContractAddressParams) WithContractAddress(contractAddress string) *PostWasmContractsContractAddressParams {
	o.SetContractAddress(contractAddress)
	return o
}

// SetContractAddress adds the contractAddress to the post wasm contracts contract address params
func (o *PostWasmContractsContractAddressParams) SetContractAddress(contractAddress string) {
	o.ContractAddress = contractAddress
}

// WithExecuteContractRequestBody adds the executeContractRequestBody to the post wasm contracts contract address params
func (o *PostWasmContractsContractAddressParams) WithExecuteContractRequestBody(executeContractRequestBody *models.ExecuteContractReq) *PostWasmContractsContractAddressParams {
	o.SetExecuteContractRequestBody(executeContractRequestBody)
	return o
}

// SetExecuteContractRequestBody adds the executeContractRequestBody to the post wasm contracts contract address params
func (o *PostWasmContractsContractAddressParams) SetExecuteContractRequestBody(executeContractRequestBody *models.ExecuteContractReq) {
	o.ExecuteContractRequestBody = executeContractRequestBody
}

// WriteToRequest writes these params to a swagger request
func (o *PostWasmContractsContractAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contractAddress
	if err := r.SetPathParam("contractAddress", o.ContractAddress); err != nil {
		return err
	}

	if o.ExecuteContractRequestBody != nil {
		if err := r.SetBodyParam(o.ExecuteContractRequestBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
