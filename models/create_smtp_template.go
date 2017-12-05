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

// CreateSMTPTemplate create Smtp template
// swagger:model createSmtpTemplate
type CreateSMTPTemplate struct {

	// Absolute url of the attachment (no local file). Extensions allowed xlsx, xls, ods, docx, docm, doc, csv, pdf, txt, gif, jpg, jpeg, png, tif, tiff and rtf
	AttachmentURL string `json:"attachmentUrl,omitempty"`

	// Body of the message (HTML version). The field must have more than 10 characters. REQUIRED if htmlUrl is empty
	HTMLContent string `json:"htmlContent,omitempty"`

	// Url which contents the body of the email message. REQUIRED if htmlContent is empty
	HTMLURL string `json:"htmlUrl,omitempty"`

	// Status of template. isActive = true means template is active and isActive = false means template is inactive
	IsActive bool `json:"isActive,omitempty"`

	// Email on which campaign recipients will be able to reply to
	ReplyTo strfmt.Email `json:"replyTo,omitempty"`

	// sender
	// Required: true
	Sender *CreateSMTPTemplateSender `json:"sender"`

	// Subject of the template
	// Required: true
	Subject *string `json:"subject"`

	// Tag of the template
	Tag string `json:"tag,omitempty"`

	// Name of the template
	// Required: true
	TemplateName *string `json:"templateName"`

	// This is to personalize the «To» Field. If you want to include the first name and last name of your recipient, add [FNAME] [LNAME]. To use the contact attributes here, these must already exist in SendinBlue account
	ToField string `json:"toField,omitempty"`
}

// Validate validates this create Smtp template
func (m *CreateSMTPTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSender(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTemplateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateSMTPTemplate) validateSender(formats strfmt.Registry) error {

	if err := validate.Required("sender", "body", m.Sender); err != nil {
		return err
	}

	if m.Sender != nil {

		if err := m.Sender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sender")
			}
			return err
		}
	}

	return nil
}

func (m *CreateSMTPTemplate) validateSubject(formats strfmt.Registry) error {

	if err := validate.Required("subject", "body", m.Subject); err != nil {
		return err
	}

	return nil
}

func (m *CreateSMTPTemplate) validateTemplateName(formats strfmt.Registry) error {

	if err := validate.Required("templateName", "body", m.TemplateName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateSMTPTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSMTPTemplate) UnmarshalBinary(b []byte) error {
	var res CreateSMTPTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
