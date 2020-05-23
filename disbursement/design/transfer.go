package design

import (
	. "github.com/wondenge/momo-go/momo"
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("transfer", func() {

	Error("internal_error", ErrorResult, "check log for information")
	Error("bad_request", ErrorResult, "invalid data was sent in the request")

	HTTP(func() {
		Path("/disbursement")
	})

	// The payer will be asked to authorize the payment.
	// The transaction will be executed once the payer has authorized the payment.
	// The requesttopay will be in status PENDING until the transaction is authorized
	// or declined by the payer or it is timed out by the system.
	// Status of the transaction can be validated by using the GET /requesttopay/ <resourceId>
	Method("create", func() {
		Description("Transfers an amount from the owner’s account to a payee account")

		// Use JWT to auth requests to this endpoint.
		// Enforce presence of "api:read" scope in JWT claims.
		Security(JWTAuth, func() {
			Scope("api:read")
		})
		Payload(Transfer)
		Payload(func() {
			Token("Authorization", String, " Authorization header used for Basic authentication")
			Attribute("X-Callback-Url", String, "URL to the server where the callback should be sent.")
			Attribute("X-Reference-Id", String, "Resource ID of the created request to pay transaction.")
			Attribute("X-Target-Environment", String, "Identifier of the EWP system.")
			Attribute("Ocp-Apim-Subscription-Key", String, "Subscription Key")
			Required("X-Reference-Id", "X-Target-Environment", "Ocp-Apim-Subscription-Key")
		})
		Result(String)
		Result(ErrorReason)
		HTTP(func() {
			POST("/v1_0/transfer")
			Headers(func() {

				// Format of the header parameter follows the standard for Basic and Bearer.
				// Oauth uses Bearer authentication type where the credential is the received access token.
				Header("Authorization: Authorization", String, func() {
					Description(" Authorization header used for Basic authentication and oauth")
				})
				Header("X-Callback-Url: X-Callback-Url", String, func() {
					Description("URL to the server where the callback should be sent.")
				})

				// This ID is used, for example, validating the status of the request.
				// ‘Universal Unique ID’ for the transaction generated using UUID version 4.
				Header("X-Reference-Id: X-Reference-Id", String, func() {
					Description("Resource ID of the created request to pay transaction.")
				})

				// This parameter is used to route the request to the EWP
				// system that will initiate the transaction.
				Header("X-Target-Environment: X-Target-Environment", String, func() {
					Description("The identifier of the EWP system where the transaction shall be processed.")
				})
				Header("Ocp-Apim-Subscription-Key: Ocp-Apim-Subscription-Key", String, func() {
					Description("Subscription key which provides access to this API")
				})
				Required("X-Reference-Id", "X-Target-Environment", "Ocp-Apim-Subscription-Key")
			})
			Response(StatusAccepted)
			Response("bad_request", StatusBadRequest)
		})
	})

	// Status of the transaction can validated by using the GET /transfer/\{referenceId\}
	// X-Reference-Id that was passed in the post is used as reference to the request
	Method("get", func() {
		Description("This operation is used to get the status of a transfer.")

		// Use JWT to auth requests to this endpoint.
		// Enforce presence of "api:read" scope in JWT claims.
		Security(JWTAuth, func() {
			Scope("api:read")
		})
		Payload(func() {

			// Reference id  used when creating the request to pay.
			// X-Reference-Id that was passed in the post is
			// used as reference to the request.
			Attribute("referenceId", String, func() {
				Description(" UUID of transaction to get result")
			})
			Token("Authorization", String, " Authorization header used for Basic authentication")
			Attribute("X-Target-Environment", String, "Identifier of the EWP system.")
			Attribute("Ocp-Apim-Subscription-Key", String, "Subscription Key.")
			Required("X-Target-Environment", "Ocp-Apim-Subscription-Key")
		})
		Result(TransferResult)
		Result(ErrorReason)
		HTTP(func() {
			GET("/v1_0/transfer/{referenceId}")
			Headers(func() {

				// Format of the header parameter follows the standard for Basic and Bearer.
				// Oauth uses Bearer authentication type where the credential is the received access token.
				Header("Authorization: Authorization", String, func() {
					Description(" Authorization header used for Basic authentication and oauth")
				})

				// This parameter is used to route the request to the EWP
				// system that will initiate the transaction.
				Header("X-Target-Environment: X-Target-Environment", String, func() {
					Description("The identifier of the EWP system where the transaction shall be processed.")
				})
				Header("Ocp-Apim-Subscription-Key: Ocp-Apim-Subscription-Key", String, func() {
					Description("Subscription key which provides access to this API")
				})
				Required("X-Target-Environment", "Ocp-Apim-Subscription-Key")
			})

			// Note that a  failed request to pay will be returned with this status too.
			// The 'status' of the RequestToPayResult can be used to determine the outcome of the request.
			// The 'reason' field can be used to retrieve a cause in case of failure.
			Response(StatusOK)
			Response("bad_request", StatusBadRequest)
		})
	})
})
