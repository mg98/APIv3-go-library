// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpdateSender update sender
// swagger:model updateSender
type UpdateSender struct {

	// From Email to update the sender
	Email strfmt.Email `json:"email,omitempty"`

	// ips
	Ips UpdateSenderIps `json:"ips"`

	// From Name to update the sender
	Name string `json:"name,omitempty"`
}

// Validate validates this update sender
func (m *UpdateSender) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateSender) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateSender) UnmarshalBinary(b []byte) error {
	var res UpdateSender
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
