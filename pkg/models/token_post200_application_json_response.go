package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TokenPost200ApplicationJSONResponse token post200 application Json response
//
// swagger:model TokenPost200ApplicationJsonResponse
type TokenPost200ApplicationJSONResponse struct {

	// A JWT token which can be used to authrize against the other API end-points. The format of the token follows the JWT standard format (see jwt.io for an example). This is the token that should be sent in in the Authorization header when calling the other API end-points.
	AccessToken string `json:"access_token,omitempty"`

	// The validity time in seconds of the token.
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// The token type.
	TokenType string `json:"token_type,omitempty"`
}

// Validate validates this token post200 application Json response
func (m *TokenPost200ApplicationJSONResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TokenPost200ApplicationJSONResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenPost200ApplicationJSONResponse) UnmarshalBinary(b []byte) error {
	var res TokenPost200ApplicationJSONResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
