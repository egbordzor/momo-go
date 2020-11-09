package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Balance The available balance of the account
//
// swagger:model Balance
type Balance struct {

	// The available balance of the account
	AvailableBalance string `json:"availableBalance,omitempty"`

	// ISO4217 Currency
	Currency string `json:"currency,omitempty"`
}

// Validate validates this balance
func (m *Balance) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Balance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Balance) UnmarshalBinary(b []byte) error {
	var res Balance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
