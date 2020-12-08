// Code generated by go-swagger; DO NOT EDIT.

package governance

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

// PostGovProposalsParamChangeReader is a Reader for the PostGovProposalsParamChange structure.
type PostGovProposalsParamChangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGovProposalsParamChangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostGovProposalsParamChangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostGovProposalsParamChangeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostGovProposalsParamChangeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGovProposalsParamChangeOK creates a PostGovProposalsParamChangeOK with default headers values
func NewPostGovProposalsParamChangeOK() *PostGovProposalsParamChangeOK {
	return &PostGovProposalsParamChangeOK{}
}

/*PostGovProposalsParamChangeOK handles this case with default header values.

The transaction was successfully generated
*/
type PostGovProposalsParamChangeOK struct {
	Payload *models.StdTx
}

func (o *PostGovProposalsParamChangeOK) Error() string {
	return fmt.Sprintf("[POST /gov/proposals/param_change][%d] postGovProposalsParamChangeOK  %+v", 200, o.Payload)
}

func (o *PostGovProposalsParamChangeOK) GetPayload() *models.StdTx {
	return o.Payload
}

func (o *PostGovProposalsParamChangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StdTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGovProposalsParamChangeBadRequest creates a PostGovProposalsParamChangeBadRequest with default headers values
func NewPostGovProposalsParamChangeBadRequest() *PostGovProposalsParamChangeBadRequest {
	return &PostGovProposalsParamChangeBadRequest{}
}

/*PostGovProposalsParamChangeBadRequest handles this case with default header values.

Invalid proposal body
*/
type PostGovProposalsParamChangeBadRequest struct {
}

func (o *PostGovProposalsParamChangeBadRequest) Error() string {
	return fmt.Sprintf("[POST /gov/proposals/param_change][%d] postGovProposalsParamChangeBadRequest ", 400)
}

func (o *PostGovProposalsParamChangeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostGovProposalsParamChangeInternalServerError creates a PostGovProposalsParamChangeInternalServerError with default headers values
func NewPostGovProposalsParamChangeInternalServerError() *PostGovProposalsParamChangeInternalServerError {
	return &PostGovProposalsParamChangeInternalServerError{}
}

/*PostGovProposalsParamChangeInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostGovProposalsParamChangeInternalServerError struct {
}

func (o *PostGovProposalsParamChangeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /gov/proposals/param_change][%d] postGovProposalsParamChangeInternalServerError ", 500)
}

func (o *PostGovProposalsParamChangeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostGovProposalsParamChangeBody post gov proposals param change body
swagger:model PostGovProposalsParamChangeBody
*/
type PostGovProposalsParamChangeBody struct {

	// base req
	BaseReq *models.BaseReq `json:"base_req,omitempty"`

	// changes
	Changes []*models.ParamChange `json:"changes"`

	// deposit
	Deposit []*models.Coin `json:"deposit"`

	// description
	Description string `json:"description,omitempty"`

	// proposer
	Proposer models.Address `json:"proposer,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this post gov proposals param change body
func (o *PostGovProposalsParamChangeBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDeposit(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProposer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostGovProposalsParamChangeBody) validateBaseReq(formats strfmt.Registry) error {

	if swag.IsZero(o.BaseReq) { // not required
		return nil
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("post_proposal_body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

func (o *PostGovProposalsParamChangeBody) validateChanges(formats strfmt.Registry) error {

	if swag.IsZero(o.Changes) { // not required
		return nil
	}

	for i := 0; i < len(o.Changes); i++ {
		if swag.IsZero(o.Changes[i]) { // not required
			continue
		}

		if o.Changes[i] != nil {
			if err := o.Changes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("post_proposal_body" + "." + "changes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostGovProposalsParamChangeBody) validateDeposit(formats strfmt.Registry) error {

	if swag.IsZero(o.Deposit) { // not required
		return nil
	}

	for i := 0; i < len(o.Deposit); i++ {
		if swag.IsZero(o.Deposit[i]) { // not required
			continue
		}

		if o.Deposit[i] != nil {
			if err := o.Deposit[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("post_proposal_body" + "." + "deposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostGovProposalsParamChangeBody) validateProposer(formats strfmt.Registry) error {

	if swag.IsZero(o.Proposer) { // not required
		return nil
	}

	if err := o.Proposer.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("post_proposal_body" + "." + "proposer")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostGovProposalsParamChangeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostGovProposalsParamChangeBody) UnmarshalBinary(b []byte) error {
	var res PostGovProposalsParamChangeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
