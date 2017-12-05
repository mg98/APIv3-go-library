// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetSmsCampaignsCampaignsItemsAllOf1Statistics get sms campaigns campaigns items all of1 statistics
// swagger:model getSmsCampaignsCampaignsItemsAllOf1Statistics
type GetSmsCampaignsCampaignsItemsAllOf1Statistics struct {
	GetSmsCampaignStats
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetSmsCampaignsCampaignsItemsAllOf1Statistics) UnmarshalJSON(raw []byte) error {

	var aO0 GetSmsCampaignStats
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GetSmsCampaignStats = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetSmsCampaignsCampaignsItemsAllOf1Statistics) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.GetSmsCampaignStats)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get sms campaigns campaigns items all of1 statistics
func (m *GetSmsCampaignsCampaignsItemsAllOf1Statistics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.GetSmsCampaignStats.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetSmsCampaignsCampaignsItemsAllOf1Statistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSmsCampaignsCampaignsItemsAllOf1Statistics) UnmarshalBinary(b []byte) error {
	var res GetSmsCampaignsCampaignsItemsAllOf1Statistics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
