/*
 * Remittance
 *
 * Partner Gateway API document
 *
 * API version: 1.0
 *
 */

package remittance

// ErrorReason struct for ErrorReason
type ErrorReason struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
