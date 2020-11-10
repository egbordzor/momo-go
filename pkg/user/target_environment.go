/*
 * Sandbox User Provisioning
 *
 * Partner Gateway sandbox provisioning API document
 *
 * API version: 1.0
 *
 */

package user

// TargetEnvironment struct for TargetEnvironment
type TargetEnvironment struct {
	// The target environment
	ApiKey string `json:"apiKey,omitempty"`
}
