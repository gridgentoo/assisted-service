// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DiskConfigParams disk config params
//
// swagger:model disk-config-params
type DiskConfigParams struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// role
	Role DiskRole `json:"role,omitempty"`
}

// Validate validates this disk config params
func (m *DiskConfigParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiskConfigParams) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DiskConfigParams) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if err := m.Role.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("role")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("role")
		}
		return err
	}

	return nil
}

// ContextValidate validate this disk config params based on the context it is used
func (m *DiskConfigParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiskConfigParams) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Role.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("role")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("role")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DiskConfigParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiskConfigParams) UnmarshalBinary(b []byte) error {
	var res DiskConfigParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
