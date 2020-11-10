/*
 * Sandbox User Provisioning
 *
 * Partner Gateway sandbox provisioning API document
 *
 * API version: 1.0
 *
 */

package user

// ApiUser The create API user information
type ApiUser struct {
	// The provider callback host
	ProviderCallbackHost string `json:"providerCallbackHost,omitempty"`
}
