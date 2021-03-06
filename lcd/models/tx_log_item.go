// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TxLogItem tx log item
//
// swagger:model TxLogItem
type TxLogItem struct {

	// events
	Events TxEvents `json:"events,omitempty"`

	// log
	Log string `json:"log,omitempty"`

	// msg index
	MsgIndex float64 `json:"msg_index,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this tx log item
func (m *TxLogItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TxLogItem) validateEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.Events) { // not required
		return nil
	}

	if err := m.Events.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("events")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TxLogItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TxLogItem) UnmarshalBinary(b []byte) error {
	var res TxLogItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
