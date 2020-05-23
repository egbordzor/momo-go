package design

import (
	. "github.com/wondenge/momo-go/momo"
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("getbalance", func() {

	Error("internal_error", ErrorResult, "check log for information")
	Error("bad_request", ErrorResult, "invalid data was sent in the request")

	HTTP(func() {
		Path("/remittance")
	})

	// Get the balance of the account.
	Method("show", func() {
		Description("Get the balance of the account")

		// Use JWT to auth requests to this endpoint.
		// Enforce presence of "api:read" scope in JWT claims.
		Security(JWTAuth, func() {
			Scope("api:read")
		})
		Payload(func() {
			Token("Authorization", String, "Authorization header for Basic authentication.")
			Attribute("X-Target-Environment", String, "Identifier of the EWP system.")
			Attribute("Ocp-Apim-Subscription-Key", String, "Subscription Key")
			Required("X-Target-Environment", "Ocp-Apim-Subscription-Key")
		})
		Result(Balance)
		Result(ErrorReason)
		HTTP(func() {
			GET("/v1_0/account/balance")
			Headers(func() {

				// Format of the header parameter follows the standard
				// for Basic and Bearer.
				// Oauth uses Bearer authentication type where the
				// credential is the received access token.
				Header("Authorization: Authorization", String, func() {
					Description("Authorization header for Basic authentication & oauth")
					Pattern("^Bearer [^ ]+$")
					Example("Basic 67ew8n31me")
				})

				// This parameter is used to route the request to the
				// EWP system that will initiate the transaction.
				Header("X-Target-Environment: X-Target-Environment", String, func() {
					Description("The identifier of the EWP system where the transaction shall be processed")
				})
				Header("Ocp-Apim-Subscription-Key: Ocp-Apim-Subscription-Key", String, func() {
					Description("Subscription key which provides access to this API")
				})
				Required("X-Target-Environment", "Ocp-Apim-Subscription-Key")
			})
			Response(StatusOK)
			Response("bad_request", StatusBadRequest)
		})
	})
})
