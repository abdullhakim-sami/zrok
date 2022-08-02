// Code generated by go-swagger; DO NOT EDIT.

package rest_model_zrok

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Identity identity
//
// swagger:model identity
type Identity struct {

	// active
	Active bool `json:"active,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`

	// ziti Id
	ZitiID string `json:"zitiId,omitempty"`
}

// Validate validates this identity
func (m *Identity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this identity based on context it is used
func (m *Identity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Identity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Identity) UnmarshalBinary(b []byte) error {
	var res Identity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}