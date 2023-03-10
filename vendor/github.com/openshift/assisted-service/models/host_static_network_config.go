// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HostStaticNetworkConfig host static network config
//
// swagger:model host_static_network_config
type HostStaticNetworkConfig struct {

	// mapping of host macs to logical interfaces used in the network yaml
	MacInterfaceMap MacInterfaceMap `json:"mac_interface_map,omitempty"`

	// yaml string that can be processed by nmstate
	NetworkYaml string `json:"network_yaml,omitempty"`
}

// Validate validates this host static network config
func (m *HostStaticNetworkConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMacInterfaceMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostStaticNetworkConfig) validateMacInterfaceMap(formats strfmt.Registry) error {
	if swag.IsZero(m.MacInterfaceMap) { // not required
		return nil
	}

	if err := m.MacInterfaceMap.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mac_interface_map")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("mac_interface_map")
		}
		return err
	}

	return nil
}

// ContextValidate validate this host static network config based on the context it is used
func (m *HostStaticNetworkConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMacInterfaceMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostStaticNetworkConfig) contextValidateMacInterfaceMap(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MacInterfaceMap.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mac_interface_map")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("mac_interface_map")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostStaticNetworkConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostStaticNetworkConfig) UnmarshalBinary(b []byte) error {
	var res HostStaticNetworkConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
