// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CollectorGroup collector group
// swagger:model CollectorGroup
type CollectorGroup struct {

	// The time at which the group was created, in epoch format
	// Read Only: true
	CreateOn int64 `json:"createOn,omitempty"`

	// The custom properties defined for the Collector group
	CustomProperties []*NameAndValue `json:"customProperties,omitempty"`

	// The description of the Collector Group
	Description string `json:"description,omitempty"`

	// The id of the Collector Group
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// The name of the Collector Group
	// Required: true
	Name *string `json:"name"`

	// The number of Collectors that belong to the group
	// Read Only: true
	NumOfCollectors int32 `json:"numOfCollectors,omitempty"`

	// The permission level of the user that made the API request
	// Read Only: true
	UserPermission string `json:"userPermission,omitempty"`
}

// Validate validates this collector group
func (m *CollectorGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectorGroup) validateCustomProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomProperties); i++ {
		if swag.IsZero(m.CustomProperties[i]) { // not required
			continue
		}

		if m.CustomProperties[i] != nil {
			if err := m.CustomProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CollectorGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CollectorGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectorGroup) UnmarshalBinary(b []byte) error {
	var res CollectorGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
