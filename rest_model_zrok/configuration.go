// Code generated by go-swagger; DO NOT EDIT.

package rest_model_zrok

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Configuration configuration
//
// swagger:model configuration
type Configuration struct {

	// invite token contact
	InviteTokenContact string `json:"inviteTokenContact,omitempty"`

	// invites open
	InvitesOpen bool `json:"invitesOpen,omitempty"`

	// password requirements
	PasswordRequirements *PasswordRequirements `json:"passwordRequirements,omitempty"`

	// requires invite token
	RequiresInviteToken bool `json:"requiresInviteToken,omitempty"`

	// tou link
	TouLink string `json:"touLink,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this configuration
func (m *Configuration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePasswordRequirements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configuration) validatePasswordRequirements(formats strfmt.Registry) error {
	if swag.IsZero(m.PasswordRequirements) { // not required
		return nil
	}

	if m.PasswordRequirements != nil {
		if err := m.PasswordRequirements.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("passwordRequirements")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("passwordRequirements")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configuration based on the context it is used
func (m *Configuration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePasswordRequirements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configuration) contextValidatePasswordRequirements(ctx context.Context, formats strfmt.Registry) error {

	if m.PasswordRequirements != nil {
		if err := m.PasswordRequirements.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("passwordRequirements")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("passwordRequirements")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configuration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configuration) UnmarshalBinary(b []byte) error {
	var res Configuration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
