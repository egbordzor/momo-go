package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIUserKeyResult Api user key result
//
// swagger:model ApiUserKeyResult
type APIUserKeyResult struct {

	// The created API user key
	APIKey string `json:"apiKey,omitempty"`
}

// Validate validates this Api user key result
func (m *APIUserKeyResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIUserKeyResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIUserKeyResult) UnmarshalBinary(b []byte) error {
	var res APIUserKeyResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
