// Code generated by go-swagger; DO NOT EDIT.

package wasm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/terra-project/mantle-sdk/lcd/models"
)

// GetWasmContractsContractAddressReader is a Reader for the GetWasmContractsContractAddress structure.
type GetWasmContractsContractAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWasmContractsContractAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWasmContractsContractAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWasmContractsContractAddressOK creates a GetWasmContractsContractAddressOK with default headers values
func NewGetWasmContractsContractAddressOK() *GetWasmContractsContractAddressOK {
	return &GetWasmContractsContractAddressOK{}
}

/*GetWasmContractsContractAddressOK handles this case with default header values.

OK
*/
type GetWasmContractsContractAddressOK struct {
	Payload *GetWasmContractsContractAddressOKBody
}

func (o *GetWasmContractsContractAddressOK) Error() string {
	return fmt.Sprintf("[GET /wasm/contracts/{contractAddress}][%d] getWasmContractsContractAddressOK  %+v", 200, o.Payload)
}

func (o *GetWasmContractsContractAddressOK) GetPayload() *GetWasmContractsContractAddressOKBody {
	return o.Payload
}

func (o *GetWasmContractsContractAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWasmContractsContractAddressOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWasmContractsContractAddressOKBody get wasm contracts contract address o k body
swagger:model GetWasmContractsContractAddressOKBody
*/
type GetWasmContractsContractAddressOKBody struct {
	Error string `json:"error"`

	// height
	Height string `json:"height,omitempty"`

	// result
	Result *models.ContractInfo `json:"result,omitempty"`
}

// Validate validates this get wasm contracts contract address o k body
func (o *GetWasmContractsContractAddressOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWasmContractsContractAddressOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWasmContractsContractAddressOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWasmContractsContractAddressOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWasmContractsContractAddressOKBody) UnmarshalBinary(b []byte) error {
	var res GetWasmContractsContractAddressOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
