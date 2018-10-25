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

// DeviceAllocationRequest device allocation request
// swagger:model DeviceAllocationRequest
type DeviceAllocationRequest struct {

	// The desired name for this host in the netbox
	// Required: true
	Name *string `json:"name"`

	// The name of the tenant to assign this device to
	// Required: true
	Tenant *string `json:"tenant"`

	// The name of the tenant group to assign this device to
	// Required: true
	TenantGroup *string `json:"tenant_group"`
}

// Validate validates this device allocation request
func (m *DeviceAllocationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceAllocationRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DeviceAllocationRequest) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

func (m *DeviceAllocationRequest) validateTenantGroup(formats strfmt.Registry) error {

	if err := validate.Required("tenant_group", "body", m.TenantGroup); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceAllocationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceAllocationRequest) UnmarshalBinary(b []byte) error {
	var res DeviceAllocationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
