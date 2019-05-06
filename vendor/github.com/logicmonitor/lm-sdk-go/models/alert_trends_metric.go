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

// AlertTrendsMetric alert trends metric
// swagger:model AlertTrendsMetric
type AlertTrendsMetric struct {

	// item type
	// Required: true
	ItemType *string `json:"itemType"`

	// item val
	// Required: true
	ItemVal *string `json:"itemVal"`
}

// Validate validates this alert trends metric
func (m *AlertTrendsMetric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemVal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertTrendsMetric) validateItemType(formats strfmt.Registry) error {

	if err := validate.Required("itemType", "body", m.ItemType); err != nil {
		return err
	}

	return nil
}

func (m *AlertTrendsMetric) validateItemVal(formats strfmt.Registry) error {

	if err := validate.Required("itemVal", "body", m.ItemVal); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertTrendsMetric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertTrendsMetric) UnmarshalBinary(b []byte) error {
	var res AlertTrendsMetric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
