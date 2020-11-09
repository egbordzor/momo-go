package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BooleanResult boolean result
//
// swagger:model BooleanResult
type BooleanResult struct {

	// result
	Result bool `json:"result,omitempty"`
}

// Validate validates this boolean result
func (m *BooleanResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BooleanResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BooleanResult) UnmarshalBinary(b []byte) error {
	var res BooleanResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
