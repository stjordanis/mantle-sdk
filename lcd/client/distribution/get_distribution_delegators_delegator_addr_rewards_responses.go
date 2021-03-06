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

// GetDistributionDelegatorsDelegatorAddrRewardsReader is a Reader for the GetDistributionDelegatorsDelegatorAddrRewards structure.
type GetDistributionDelegatorsDelegatorAddrRewardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDistributionDelegatorsDelegatorAddrRewardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDistributionDelegatorsDelegatorAddrRewardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDistributionDelegatorsDelegatorAddrRewardsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDistributionDelegatorsDelegatorAddrRewardsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDistributionDelegatorsDelegatorAddrRewardsOK creates a GetDistributionDelegatorsDelegatorAddrRewardsOK with default headers values
func NewGetDistributionDelegatorsDelegatorAddrRewardsOK() *GetDistributionDelegatorsDelegatorAddrRewardsOK {
	return &GetDistributionDelegatorsDelegatorAddrRewardsOK{}
}

/*GetDistributionDelegatorsDelegatorAddrRewardsOK handles this case with default header values.

OK
*/
type GetDistributionDelegatorsDelegatorAddrRewardsOK struct {
	Payload *GetDistributionDelegatorsDelegatorAddrRewardsOKBody
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsOK) Error() string {
	return fmt.Sprintf("[GET /distribution/delegators/{delegatorAddr}/rewards][%d] getDistributionDelegatorsDelegatorAddrRewardsOK  %+v", 200, o.Payload)
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsOK) GetPayload() *GetDistributionDelegatorsDelegatorAddrRewardsOKBody {
	return o.Payload
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDistributionDelegatorsDelegatorAddrRewardsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDistributionDelegatorsDelegatorAddrRewardsBadRequest creates a GetDistributionDelegatorsDelegatorAddrRewardsBadRequest with default headers values
func NewGetDistributionDelegatorsDelegatorAddrRewardsBadRequest() *GetDistributionDelegatorsDelegatorAddrRewardsBadRequest {
	return &GetDistributionDelegatorsDelegatorAddrRewardsBadRequest{}
}

/*GetDistributionDelegatorsDelegatorAddrRewardsBadRequest handles this case with default header values.

Invalid delegator address
*/
type GetDistributionDelegatorsDelegatorAddrRewardsBadRequest struct {
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsBadRequest) Error() string {
	return fmt.Sprintf("[GET /distribution/delegators/{delegatorAddr}/rewards][%d] getDistributionDelegatorsDelegatorAddrRewardsBadRequest ", 400)
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDistributionDelegatorsDelegatorAddrRewardsInternalServerError creates a GetDistributionDelegatorsDelegatorAddrRewardsInternalServerError with default headers values
func NewGetDistributionDelegatorsDelegatorAddrRewardsInternalServerError() *GetDistributionDelegatorsDelegatorAddrRewardsInternalServerError {
	return &GetDistributionDelegatorsDelegatorAddrRewardsInternalServerError{}
}

/*GetDistributionDelegatorsDelegatorAddrRewardsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetDistributionDelegatorsDelegatorAddrRewardsInternalServerError struct {
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /distribution/delegators/{delegatorAddr}/rewards][%d] getDistributionDelegatorsDelegatorAddrRewardsInternalServerError ", 500)
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetDistributionDelegatorsDelegatorAddrRewardsOKBody get distribution delegators delegator addr rewards o k body
swagger:model GetDistributionDelegatorsDelegatorAddrRewardsOKBody
*/
type GetDistributionDelegatorsDelegatorAddrRewardsOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result *models.DelegatorTotalRewards `json:"result,omitempty"`
}

// Validate validates this get distribution delegators delegator addr rewards o k body
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBody) UnmarshalBinary(b []byte) error {
	var res GetDistributionDelegatorsDelegatorAddrRewardsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
