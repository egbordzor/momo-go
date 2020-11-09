package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaymentServerURL payment server Url
//
// swagger:model PaymentServerUrl
type PaymentServerURL struct {

	// The payment server URL
	APIKey string `json:"apiKey,omitempty"`
}

// Validate validates this payment server Url
func (m *PaymentServerURL) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentServerURL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentServerURL) UnmarshalBinary(b []byte) error {
	var res PaymentServerURL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
