// Code generated by go-swagger; DO NOT EDIT.

package oracle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/terra-project/mantle-sdk/lcd/models"
)

// GetOracleDenomsDenomPrevotesValidatorReader is a Reader for the GetOracleDenomsDenomPrevotesValidator structure.
type GetOracleDenomsDenomPrevotesValidatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOracleDenomsDenomPrevotesValidatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOracleDenomsDenomPrevotesValidatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOracleDenomsDenomPrevotesValidatorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOracleDenomsDenomPrevotesValidatorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOracleDenomsDenomPrevotesValidatorOK creates a GetOracleDenomsDenomPrevotesValidatorOK with default headers values
func NewGetOracleDenomsDenomPrevotesValidatorOK() *GetOracleDenomsDenomPrevotesValidatorOK {
	return &GetOracleDenomsDenomPrevotesValidatorOK{}
}

/*GetOracleDenomsDenomPrevotesValidatorOK handles this case with default header values.

OK
*/
type GetOracleDenomsDenomPrevotesValidatorOK struct {
	Payload *GetOracleDenomsDenomPrevotesValidatorOKBody
}

func (o *GetOracleDenomsDenomPrevotesValidatorOK) Error() string {
	return fmt.Sprintf("[GET /oracle/denoms/{denom}/prevotes/{validator}][%d] getOracleDenomsDenomPrevotesValidatorOK  %+v", 200, o.Payload)
}

func (o *GetOracleDenomsDenomPrevotesValidatorOK) GetPayload() *GetOracleDenomsDenomPrevotesValidatorOKBody {
	return o.Payload
}

func (o *GetOracleDenomsDenomPrevotesValidatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOracleDenomsDenomPrevotesValidatorOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOracleDenomsDenomPrevotesValidatorBadRequest creates a GetOracleDenomsDenomPrevotesValidatorBadRequest with default headers values
func NewGetOracleDenomsDenomPrevotesValidatorBadRequest() *GetOracleDenomsDenomPrevotesValidatorBadRequest {
	return &GetOracleDenomsDenomPrevotesValidatorBadRequest{}
}

/*GetOracleDenomsDenomPrevotesValidatorBadRequest handles this case with default header values.

Bad Request
*/
type GetOracleDenomsDenomPrevotesValidatorBadRequest struct {
}

func (o *GetOracleDenomsDenomPrevotesValidatorBadRequest) Error() string {
	return fmt.Sprintf("[GET /oracle/denoms/{denom}/prevotes/{validator}][%d] getOracleDenomsDenomPrevotesValidatorBadRequest ", 400)
}

func (o *GetOracleDenomsDenomPrevotesValidatorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOracleDenomsDenomPrevotesValidatorInternalServerError creates a GetOracleDenomsDenomPrevotesValidatorInternalServerError with default headers values
func NewGetOracleDenomsDenomPrevotesValidatorInternalServerError() *GetOracleDenomsDenomPrevotesValidatorInternalServerError {
	return &GetOracleDenomsDenomPrevotesValidatorInternalServerError{}
}

/*GetOracleDenomsDenomPrevotesValidatorInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetOracleDenomsDenomPrevotesValidatorInternalServerError struct {
}

func (o *GetOracleDenomsDenomPrevotesValidatorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /oracle/denoms/{denom}/prevotes/{validator}][%d] getOracleDenomsDenomPrevotesValidatorInternalServerError ", 500)
}

func (o *GetOracleDenomsDenomPrevotesValidatorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetOracleDenomsDenomPrevotesValidatorOKBody get oracle denoms denom prevotes validator o k body
swagger:model GetOracleDenomsDenomPrevotesValidatorOKBody
*/
type GetOracleDenomsDenomPrevotesValidatorOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result []*models.ExchangeRatePrevote `json:"result"`
}

// Validate validates this get oracle denoms denom prevotes validator o k body
func (o *GetOracleDenomsDenomPrevotesValidatorOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOracleDenomsDenomPrevotesValidatorOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	for i := 0; i < len(o.Result); i++ {
		if swag.IsZero(o.Result[i]) { // not required
			continue
		}

		if o.Result[i] != nil {
			if err := o.Result[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOracleDenomsDenomPrevotesValidatorOK" + "." + "result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOracleDenomsDenomPrevotesValidatorOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOracleDenomsDenomPrevotesValidatorOKBody) UnmarshalBinary(b []byte) error {
	var res GetOracleDenomsDenomPrevotesValidatorOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
