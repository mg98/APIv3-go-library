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

// SendSMTPEmailCcItems send Smtp email cc items
// swagger:model sendSmtpEmailCcItems
type SendSMTPEmailCcItems struct {

	// Email address of the recipient in cc
	// Required: true
	Email *strfmt.Email `json:"email"`

	// Name of the recipient in cc
	Name string `json:"name,omitempty"`
}

// Validate validates this send Smtp email cc items
func (m *SendSMTPEmailCcItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendSMTPEmailCcItems) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendSMTPEmailCcItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendSMTPEmailCcItems) UnmarshalBinary(b []byte) error {
	var res SendSMTPEmailCcItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
