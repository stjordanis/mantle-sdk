// Code generated by go-swagger; DO NOT EDIT.

package wasm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terra-project/mantle-sdk/lcd/models"
)

// PostWasmContractsContractAddressReader is a Reader for the PostWasmContractsContractAddress structure.
type PostWasmContractsContractAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWasmContractsContractAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWasmContractsContractAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWasmContractsContractAddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWasmContractsContractAddressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWasmContractsContractAddressOK creates a PostWasmContractsContractAddressOK with default headers values
func NewPostWasmContractsContractAddressOK() *PostWasmContractsContractAddressOK {
	return &PostWasmContractsContractAddressOK{}
}

/*PostWasmContractsContractAddressOK handles this case with default header values.

OK
*/
type PostWasmContractsContractAddressOK struct {
	Payload *models.StdTx
}

func (o *PostWasmContractsContractAddressOK) Error() string {
	return fmt.Sprintf("[POST /wasm/contracts/{contractAddress}][%d] postWasmContractsContractAddressOK  %+v", 200, o.Payload)
}

func (o *PostWasmContractsContractAddressOK) GetPayload() *models.StdTx {
	return o.Payload
}

func (o *PostWasmContractsContractAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StdTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWasmContractsContractAddressBadRequest creates a PostWasmContractsContractAddressBadRequest with default headers values
func NewPostWasmContractsContractAddressBadRequest() *PostWasmContractsContractAddressBadRequest {
	return &PostWasmContractsContractAddressBadRequest{}
}

/*PostWasmContractsContractAddressBadRequest handles this case with default header values.

Bad request
*/
type PostWasmContractsContractAddressBadRequest struct {
}

func (o *PostWasmContractsContractAddressBadRequest) Error() string {
	return fmt.Sprintf("[POST /wasm/contracts/{contractAddress}][%d] postWasmContractsContractAddressBadRequest ", 400)
}

func (o *PostWasmContractsContractAddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWasmContractsContractAddressInternalServerError creates a PostWasmContractsContractAddressInternalServerError with default headers values
func NewPostWasmContractsContractAddressInternalServerError() *PostWasmContractsContractAddressInternalServerError {
	return &PostWasmContractsContractAddressInternalServerError{}
}

/*PostWasmContractsContractAddressInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostWasmContractsContractAddressInternalServerError struct {
}

func (o *PostWasmContractsContractAddressInternalServerError) Error() string {
	return fmt.Sprintf("[POST /wasm/contracts/{contractAddress}][%d] postWasmContractsContractAddressInternalServerError ", 500)
}

func (o *PostWasmContractsContractAddressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
