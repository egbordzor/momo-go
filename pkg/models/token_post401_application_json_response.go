package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TokenPost401ApplicationJSONResponse token post401 application Json response
//
// swagger:model TokenPost401ApplicationJsonResponse
type TokenPost401ApplicationJSONResponse struct {

	// An error code.
	Error string `json:"error,omitempty"`
}

// Validate validates this token post401 application Json response
func (m *TokenPost401ApplicationJSONResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TokenPost401ApplicationJSONResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenPost401ApplicationJSONResponse) UnmarshalBinary(b []byte) error {
	var res TokenPost401ApplicationJSONResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
