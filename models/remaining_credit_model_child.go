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

// RemainingCreditModelChild Credits remaining for child account
// swagger:model remainingCreditModelChild
type RemainingCreditModelChild struct {

	// Email Credits remaining for child account
	// Required: true
	Email *int64 `json:"email"`

	// SMS Credits remaining for child account
	// Required: true
	Sms *int64 `json:"sms"`
}

// Validate validates this remaining credit model child
func (m *RemainingCreditModelChild) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSms(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemainingCreditModelChild) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *RemainingCreditModelChild) validateSms(formats strfmt.Registry) error {

	if err := validate.Required("sms", "body", m.Sms); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemainingCreditModelChild) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemainingCreditModelChild) UnmarshalBinary(b []byte) error {
	var res RemainingCreditModelChild
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
