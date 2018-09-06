// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AssetSignature asset signature
// swagger:model assetSignature
type AssetSignature struct {

	// asset
	Asset *Asset `json:"asset,omitempty"`

	// form
	Form *AssetForm `json:"form,omitempty"`
}

// Validate validates this asset signature
func (m *AssetSignature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetSignature) validateAsset(formats strfmt.Registry) error {

	if swag.IsZero(m.Asset) { // not required
		return nil
	}

	if m.Asset != nil {
		if err := m.Asset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("asset")
			}
			return err
		}
	}

	return nil
}

func (m *AssetSignature) validateForm(formats strfmt.Registry) error {

	if swag.IsZero(m.Form) { // not required
		return nil
	}

	if m.Form != nil {
		if err := m.Form.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("form")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetSignature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetSignature) UnmarshalBinary(b []byte) error {
	var res AssetSignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
