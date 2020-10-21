// Code generated by go-swagger; DO NOT EDIT.

package staking

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

// GetStakingRedelegationsReader is a Reader for the GetStakingRedelegations structure.
type GetStakingRedelegationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStakingRedelegationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStakingRedelegationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetStakingRedelegationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStakingRedelegationsOK creates a GetStakingRedelegationsOK with default headers values
func NewGetStakingRedelegationsOK() *GetStakingRedelegationsOK {
	return &GetStakingRedelegationsOK{}
}

/*GetStakingRedelegationsOK handles this case with default header values.

OK
*/
type GetStakingRedelegationsOK struct {
	Payload *GetStakingRedelegationsOKBody
}

func (o *GetStakingRedelegationsOK) Error() string {
	return fmt.Sprintf("[GET /staking/redelegations][%d] getStakingRedelegationsOK  %+v", 200, o.Payload)
}

func (o *GetStakingRedelegationsOK) GetPayload() *GetStakingRedelegationsOKBody {
	return o.Payload
}

func (o *GetStakingRedelegationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetStakingRedelegationsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStakingRedelegationsInternalServerError creates a GetStakingRedelegationsInternalServerError with default headers values
func NewGetStakingRedelegationsInternalServerError() *GetStakingRedelegationsInternalServerError {
	return &GetStakingRedelegationsInternalServerError{}
}

/*GetStakingRedelegationsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetStakingRedelegationsInternalServerError struct {
}

func (o *GetStakingRedelegationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /staking/redelegations][%d] getStakingRedelegationsInternalServerError ", 500)
}

func (o *GetStakingRedelegationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetStakingRedelegationsOKBody get staking redelegations o k body
swagger:model GetStakingRedelegationsOKBody
*/
type GetStakingRedelegationsOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result []*models.Redelegation `json:"result"`
}

// Validate validates this get staking redelegations o k body
func (o *GetStakingRedelegationsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStakingRedelegationsOKBody) validateResult(formats strfmt.Registry) error {

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
					return ve.ValidateName("getStakingRedelegationsOK" + "." + "result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStakingRedelegationsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStakingRedelegationsOKBody) UnmarshalBinary(b []byte) error {
	var res GetStakingRedelegationsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
