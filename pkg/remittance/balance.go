/*
 * Remittance
 *
 * Partner Gateway API document
 *
 * API version: 1.0
 *
 */

package remittance

// Balance The available balance of the account
type Balance struct {
	// The available balance of the account
	AvailableBalance string `json:"availableBalance,omitempty"`
	// ISO4217 Currency
	Currency string `json:"currency,omitempty"`
}
