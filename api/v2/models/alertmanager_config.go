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

// AlertmanagerConfig alertmanager config
// swagger:model alertmanagerConfig
type AlertmanagerConfig struct {

	// original
	// Required: true
	Original *string `json:"original"`
}

// Validate validates this alertmanager config
func (m *AlertmanagerConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOriginal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertmanagerConfig) validateOriginal(formats strfmt.Registry) error {

	if err := validate.Required("original", "body", m.Original); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertmanagerConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertmanagerConfig) UnmarshalBinary(b []byte) error {
	var res AlertmanagerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
