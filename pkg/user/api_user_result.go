/*
 * Sandbox User Provisioning
 *
 * Partner Gateway sandbox provisioning API document
 *
 * API version: 1.0
 *
 */

package user

// ApiUserResult The API user information
type ApiUserResult struct {
	// The provider callback host
	ProviderCallbackHost string            `json:"providerCallbackHost,omitempty"`
	PaymentServerUrl     PaymentServerUrl  `json:"paymentServerUrl,omitempty"`
	TargetEnvironment    TargetEnvironment `json:"targetEnvironment,omitempty"`
}
