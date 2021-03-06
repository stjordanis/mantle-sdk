// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EstimateFeeResp estimate fee resp
//
// swagger:model EstimateFeeResp
type EstimateFeeResp struct {

	// fees
	Fees []*Coin `json:"fees"`

	// gas
	Gas int64 `json:"gas,omitempty"`
}

// Validate validates this estimate fee resp
func (m *EstimateFeeResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFees(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EstimateFeeResp) validateFees(formats strfmt.Registry) error {

	if swag.IsZero(m.Fees) { // not required
		return nil
	}

	for i := 0; i < len(m.Fees); i++ {
		if swag.IsZero(m.Fees[i]) { // not required
			continue
		}

		if m.Fees[i] != nil {
			if err := m.Fees[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EstimateFeeResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EstimateFeeResp) UnmarshalBinary(b []byte) error {
	var res EstimateFeeResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
