/*
 * Remittance
 *
 * Partner Gateway API document
 *
 * API version: 1.0
 *
 */

package remittance

// TokenPost200ApplicationJsonResponse struct for TokenPost200ApplicationJsonResponse
type TokenPost200ApplicationJsonResponse struct {
	// A JWT token which can be used to authrize against the other API end-points. The format of the token follows the JWT standard format (see jwt.io for an example). This is the token that should be sent in in the Authorization header when calling the other API end-points.
	AccessToken string `json:"access_token,omitempty"`
	// The token type.
	TokenType string `json:"token_type,omitempty"`
	// The validity time in seconds of the token.
	ExpiresIn int32 `json:"expires_in,omitempty"`
}
