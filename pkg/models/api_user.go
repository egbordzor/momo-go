package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIUser The create API user information
//
// swagger:model ApiUser
type APIUser struct {

	// The provider callback host
	ProviderCallbackHost string `json:"providerCallbackHost,omitempty"`
}

// Validate validates this Api user
func (m *APIUser) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIUser) UnmarshalBinary(b []byte) error {
	var res APIUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
