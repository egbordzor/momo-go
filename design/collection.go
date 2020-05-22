package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("collection", func() {

	Error("internal_error", ErrorResult, "check log for information")
	Error("bad_request", ErrorResult, "invalid data was sent in the request")

	HTTP(func() {
		Path("/collection")
	})

	// Used to authorize and authenticate other end-points of the API.
	// a) Provider system requests an access token using the API Key and API user as authentication.
	// POST/token (Authorization: Basic)
	// b) Wallet platform authenticates credentials and responds with the access token
	// Response (Access Token and Validity time)
	// c) Provider system will use the access token for any request that is sent to Wallet Platform.
	// e.g. POST /requesttopay (Authentication: Bearer)
	// Response(HTTP: 202: Accepted)
	Method("create", func() {
		Description("Creates an Access Token.")
		Security(BasicAuth)
		Payload(func() {
			Description("Credentials used to authenticate to retrieve JWT token")
			Attribute("Ocp-Apim-Subscription-Key", String, "Subscription Key.")
			Username("APIKey", String, "API Key")
			Password("APISecret", String, "API Secret")
			Required("APIKey", "APISecret")
		})
		Result(OAuthTokenResponse)
		Result(OAuthTokenError)
		HTTP(func() {
			POST("/token")
			Headers(func() {
				Header("Ocp-Apim-Subscription-Key: Ocp-Apim-Subscription-Key", String, func() {
					Description("Subscription key which provides access to this API")
				})
				Required("Ocp-Apim-Subscription-Key")
			})
			Response(StatusOK)
			Response("internal_error", StatusInternalServerError)
		})
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
				Header("Authorization: Authorization", func() {
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

	// Operation is used  to check if an account holder is registered and active in the system.
	Method("get", func() {
		Description("Checks if an account holder is registered and active in the system")

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
			Params(func() {

				//  Allowed values [msisdn, email, party_code].
				//  accountHolderId should explicitly be in small letters.
				Param("accountHolderIdType", String, func() {
					Description("Specifies the type of the party ID")
				})

				// Validated according to the party ID type (case Sensitive).
				// msisdn - Mobile Number validated according to ITU-T E.164.
				// Validated with IsMSISDN email - Validated to be a valid e-mail format.
				// Validated with IsEmail party_code - UUID of the party.
				// Validated with IsUuid
				Param("accountHolderId", String, func() {
					Description("The party number.")
				})
				Required("accountHolderIdType", "accountHolderId")
			})
			Response(StatusOK)
			Response("bad_request", StatusBadRequest)
		})
	})

	// The payer will be asked to authorize the payment.
	// The transaction will be executed once the payer has authorized the payment.
	// The requesttopay will be in status PENDING until the transaction is authorized
	// or declined by the payer or it is timed out by the system.
	// Status of the transaction can be validated by using the GET /requesttopay/ <resourceId>
	Method("post", func() {
		Description("Request a payment from a consumer (Payer).")

		// Use JWT to auth requests to this endpoint.
		// Enforce presence of "api:read" scope in JWT claims.
		Security(JWTAuth, func() {
			Scope("api:read")
		})
		Payload(RequestToPay)
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
			POST("/v1_0/requesttopay")
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

	Method("get", func() {
		Description("Get the status of a request to pay.")

		// Use JWT to auth requests to this endpoint.
		// Enforce presence of "api:read" scope in JWT claims.
		Security(JWTAuth, func() {
			Scope("api:read")
		})
		Payload(func() {
			Token("Authorization", String, " Authorization header used for Basic authentication")
			Attribute("X-Target-Environment", String, "Identifier of the EWP system.")
			Attribute("Ocp-Apim-Subscription-Key", String, "Subscription Key.")
			Required("X-Target-Environment", "Ocp-Apim-Subscription-Key")
		})
		Result(RequestToPayResult)
		Result(ErrorReason)
		HTTP(func() {
			GET("/v1_0/requesttopay/{referenceId}")
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
			Params(func() {

				// Reference id  used when creating the request to pay.
				// X-Reference-Id that was passed in the post is
				// used as reference to the request.
				Param("referenceId", String, func() {
					Description(" UUID of transaction to get result")
				})
			})

			// Note that a  failed request to pay will be returned with this status too.
			// The 'status' of the RequestToPayResult can be used to determine the outcome of the request.
			// The 'reason' field can be used to retrieve a cause in case of failure.
			Response(StatusOK)
			Response("bad_request", StatusBadRequest)
		})
	})
})
