package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TargetEnvironment target environment
//
// swagger:model TargetEnvironment
type TargetEnvironment struct {

	// The target environment
	APIKey string `json:"apiKey,omitempty"`
}

// Validate validates this target environment
func (m *TargetEnvironment) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TargetEnvironment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TargetEnvironment) UnmarshalBinary(b []byte) error {
	var res TargetEnvironment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
