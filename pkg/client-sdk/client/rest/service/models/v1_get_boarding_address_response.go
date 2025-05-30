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

// V1GetBoardingAddressResponse v1 get boarding address response
//
// swagger:model v1GetBoardingAddressResponse
type V1GetBoardingAddressResponse struct {

	// address
	Address string `json:"address,omitempty"`

	// descriptor
	Descriptor string `json:"descriptor,omitempty"`

	// tapscripts
	Tapscripts *V1Tapscripts `json:"tapscripts,omitempty"`
}

// Validate validates this v1 get boarding address response
func (m *V1GetBoardingAddressResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTapscripts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetBoardingAddressResponse) validateTapscripts(formats strfmt.Registry) error {
	if swag.IsZero(m.Tapscripts) { // not required
		return nil
	}

	if m.Tapscripts != nil {
		if err := m.Tapscripts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tapscripts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tapscripts")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 get boarding address response based on the context it is used
func (m *V1GetBoardingAddressResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTapscripts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetBoardingAddressResponse) contextValidateTapscripts(ctx context.Context, formats strfmt.Registry) error {

	if m.Tapscripts != nil {

		if swag.IsZero(m.Tapscripts) { // not required
			return nil
		}

		if err := m.Tapscripts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tapscripts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tapscripts")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1GetBoardingAddressResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GetBoardingAddressResponse) UnmarshalBinary(b []byte) error {
	var res V1GetBoardingAddressResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
