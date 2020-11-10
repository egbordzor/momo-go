/*
 * Sandbox User Provisioning
 *
 * Partner Gateway sandbox provisioning API document
 *
 * API version: 1.0
 *
 */

package user

// ErrorReason struct for ErrorReason
type ErrorReason struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
