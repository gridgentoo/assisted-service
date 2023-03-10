// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MacInterfaceMap mac interface map
//
// swagger:model mac_interface_map
type MacInterfaceMap []*MacInterfaceMapItems0

// Validate validates this mac interface map
func (m MacInterfaceMap) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this mac interface map based on the context it is used
func (m MacInterfaceMap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if m[i] != nil {
			if err := m[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MacInterfaceMapItems0 mac interface map items0
//
// swagger:model MacInterfaceMapItems0
type MacInterfaceMapItems0 struct {

	// nic name used in the yaml, which relates 1:1 to the mac address
	LogicalNicName string `json:"logical_nic_name,omitempty"`

	// mac address present on the host
	// Pattern: ^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$
	MacAddress string `json:"mac_address,omitempty"`
}

// Validate validates this mac interface map items0
func (m *MacInterfaceMapItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMacAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MacInterfaceMapItems0) validateMacAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.MacAddress) { // not required
		return nil
	}

	if err := validate.Pattern("mac_address", "body", m.MacAddress, `^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mac interface map items0 based on context it is used
func (m *MacInterfaceMapItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MacInterfaceMapItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MacInterfaceMapItems0) UnmarshalBinary(b []byte) error {
	var res MacInterfaceMapItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
