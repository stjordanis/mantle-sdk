// Code generated by go-swagger; DO NOT EDIT.

package oracle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terra-project/mantle/lcd/models"
)

// GetOracleParametersReader is a Reader for the GetOracleParameters structure.
type GetOracleParametersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOracleParametersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOracleParametersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOracleParametersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOracleParametersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOracleParametersOK creates a GetOracleParametersOK with default headers values
func NewGetOracleParametersOK() *GetOracleParametersOK {
	return &GetOracleParametersOK{}
}

/*GetOracleParametersOK handles this case with default header values.

OK
*/
type GetOracleParametersOK struct {
	Payload *models.OracleParams
}

func (o *GetOracleParametersOK) Error() string {
	return fmt.Sprintf("[GET /oracle/parameters][%d] getOracleParametersOK  %+v", 200, o.Payload)
}

func (o *GetOracleParametersOK) GetPayload() *models.OracleParams {
	return o.Payload
}

func (o *GetOracleParametersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OracleParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOracleParametersBadRequest creates a GetOracleParametersBadRequest with default headers values
func NewGetOracleParametersBadRequest() *GetOracleParametersBadRequest {
	return &GetOracleParametersBadRequest{}
}

/*GetOracleParametersBadRequest handles this case with default header values.

Bad Request
*/
type GetOracleParametersBadRequest struct {
}

func (o *GetOracleParametersBadRequest) Error() string {
	return fmt.Sprintf("[GET /oracle/parameters][%d] getOracleParametersBadRequest ", 400)
}

func (o *GetOracleParametersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOracleParametersInternalServerError creates a GetOracleParametersInternalServerError with default headers values
func NewGetOracleParametersInternalServerError() *GetOracleParametersInternalServerError {
	return &GetOracleParametersInternalServerError{}
}

/*GetOracleParametersInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetOracleParametersInternalServerError struct {
}

func (o *GetOracleParametersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /oracle/parameters][%d] getOracleParametersInternalServerError ", 500)
}

func (o *GetOracleParametersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}