package design

import (
	. "github.com/wondenge/momo-go/momo"
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("validate", func() {

	Error("internal_error", ErrorResult, "check log for information")
	Error("bad_request", ErrorResult, "invalid data was sent in the request")

	HTTP(func() {
		Path("/collection")
	})

	// Operation is used  to check if an account holder is registered and active in the system.
	Method("show", func() {
		Description("Checks if an account holder is registered and active in the system")

		// Use JWT to auth requests to this endpoint.
		// Enforce presence of "api:read" scope in JWT claims.
		Security(JWTAuth, func() {
			Scope("api:read")
		})
		Payload(func() {

			//  Allowed values [msisdn, email, party_code].
			//  accountHolderId should explicitly be in small letters.
			Attribute("accountHolderIdType", String, func() {
				Description("Specifies the type of the party ID")
			})

			// Validated according to the party ID type (case Sensitive).
			// msisdn - Mobile Number validated according to ITU-T E.164.
			// Validated with IsMSISDN email - Validated to be a valid e-mail format.
			// Validated with IsEmail party_code - UUID of the party.
			// Validated with IsUuid
			Attribute("accountHolderId", String, func() {
				Description("The party number.")
			})

			Token("Authorization", String, "Authorization header for Basic authentication.")
			Attribute("X-Target-Environment", String, "Identifier of the EWP system.")
			Attribute("Ocp-Apim-Subscription-Key", String, "Subscription Key")
			Required("accountHolderIdType", "accountHolderId", "X-Target-Environment", "Ocp-Apim-Subscription-Key")
		})
		Result(String)
		Result(ErrorReason)
		HTTP(func() {
			GET("/v1_0/accountholder/{accountHolderIdType}/{accountHolderId}/active")
			Headers(func() {
				// Format of the header parameter follows the standard for Basic and Bearer.
				// Oauth uses Bearer authentication type where the credential is the received access token.
				Header("Authorization: Authorization", String, func() {
					Description("Authorization header for Basic authentication & oauth")
					Pattern("^Bearer [^ ]+$")
					Example("Basic 67ew8n31me")
				})

				// This parameter is used to route the request to the EWP system that will initiate the transaction.
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
