// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetChildInfoAllOf1 get child info all of1
// swagger:model getChildInfoAllOf1
type GetChildInfoAllOf1 struct {

	// api keys
	APIKeys GetChildInfoAllOf1APIKeys `json:"apiKeys"`

	// credits
	Credits *GetChildInfoAllOf1Credits `json:"credits,omitempty"`

	// ips
	Ips GetChildInfoAllOf1Ips `json:"ips"`

	// The encrypted password of child account
	// Required: true
	Password *strfmt.Password `json:"password"`

	// statistics
	Statistics *GetChildInfoAllOf1Statistics `json:"statistics,omitempty"`
}

// Validate validates this get child info all of1
func (m *GetChildInfoAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredits(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatistics(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetChildInfoAllOf1) validateCredits(formats strfmt.Registry) error {

	if swag.IsZero(m.Credits) { // not required
		return nil
	}

	if m.Credits != nil {

		if err := m.Credits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credits")
			}
			return err
		}
	}

	return nil
}

func (m *GetChildInfoAllOf1) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.FormatOf("password", "body", "password", m.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetChildInfoAllOf1) validateStatistics(formats strfmt.Registry) error {

	if swag.IsZero(m.Statistics) { // not required
		return nil
	}

	if m.Statistics != nil {

		if err := m.Statistics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statistics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetChildInfoAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetChildInfoAllOf1) UnmarshalBinary(b []byte) error {
	var res GetChildInfoAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
