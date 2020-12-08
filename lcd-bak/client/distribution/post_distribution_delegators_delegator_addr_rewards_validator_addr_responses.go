// Code generated by go-swagger; DO NOT EDIT.

package distribution

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

// PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrReader is a Reader for the PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddr structure.
type PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrOK creates a PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrOK with default headers values
func NewPostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrOK() *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrOK {
	return &PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrOK{}
}

/*PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrOK handles this case with default header values.

OK
*/
type PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrOK struct {
	Payload *models.BroadcastTxCommitResult
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrOK) Error() string {
	return fmt.Sprintf("[POST /distribution/delegators/{delegatorAddr}/rewards/{validatorAddr}][%d] postDistributionDelegatorsDelegatorAddrRewardsValidatorAddrOK  %+v", 200, o.Payload)
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrOK) GetPayload() *models.BroadcastTxCommitResult {
	return o.Payload
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BroadcastTxCommitResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBadRequest creates a PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBadRequest with default headers values
func NewPostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBadRequest() *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBadRequest {
	return &PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBadRequest{}
}

/*PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBadRequest handles this case with default header values.

Invalid delegator address or delegation body
*/
type PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBadRequest struct {
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBadRequest) Error() string {
	return fmt.Sprintf("[POST /distribution/delegators/{delegatorAddr}/rewards/{validatorAddr}][%d] postDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBadRequest ", 400)
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrUnauthorized creates a PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrUnauthorized with default headers values
func NewPostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrUnauthorized() *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrUnauthorized {
	return &PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrUnauthorized{}
}

/*PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrUnauthorized handles this case with default header values.

Key password is wrong
*/
type PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrUnauthorized struct {
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrUnauthorized) Error() string {
	return fmt.Sprintf("[POST /distribution/delegators/{delegatorAddr}/rewards/{validatorAddr}][%d] postDistributionDelegatorsDelegatorAddrRewardsValidatorAddrUnauthorized ", 401)
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrInternalServerError creates a PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrInternalServerError with default headers values
func NewPostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrInternalServerError() *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrInternalServerError {
	return &PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrInternalServerError{}
}

/*PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrInternalServerError struct {
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrInternalServerError) Error() string {
	return fmt.Sprintf("[POST /distribution/delegators/{delegatorAddr}/rewards/{validatorAddr}][%d] postDistributionDelegatorsDelegatorAddrRewardsValidatorAddrInternalServerError ", 500)
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBody post distribution delegators delegator addr rewards validator addr body
swagger:model PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBody
*/
type PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBody struct {

	// base req
	BaseReq *models.BaseReq `json:"base_req,omitempty"`
}

// Validate validates this post distribution delegators delegator addr rewards validator addr body
func (o *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBody) validateBaseReq(formats strfmt.Registry) error {

	if swag.IsZero(o.BaseReq) { // not required
		return nil
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Withdraw request body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBody) UnmarshalBinary(b []byte) error {
	var res PostDistributionDelegatorsDelegatorAddrRewardsValidatorAddrBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
