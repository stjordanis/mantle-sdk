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

// InstantiateContractReq instantiate contract req
//
// swagger:model InstantiateContractReq
type InstantiateContractReq struct {

	// base req
	BaseReq *BaseReq `json:"base_req,omitempty"`

	// init coins
	InitCoins []*Coin `json:"init_coins"`

	// json formatted string
	InitMsg string `json:"init_msg,omitempty"`
}

// Validate validates this instantiate contract req
func (m *InstantiateContractReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitCoins(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstantiateContractReq) validateBaseReq(formats strfmt.Registry) error {

	if swag.IsZero(m.BaseReq) { // not required
		return nil
	}

	if m.BaseReq != nil {
		if err := m.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("base_req")
			}
			return err
		}
	}

	return nil
}

func (m *InstantiateContractReq) validateInitCoins(formats strfmt.Registry) error {

	if swag.IsZero(m.InitCoins) { // not required
		return nil
	}

	for i := 0; i < len(m.InitCoins); i++ {
		if swag.IsZero(m.InitCoins[i]) { // not required
			continue
		}

		if m.InitCoins[i] != nil {
			if err := m.InitCoins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("init_coins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstantiateContractReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstantiateContractReq) UnmarshalBinary(b []byte) error {
	var res InstantiateContractReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
