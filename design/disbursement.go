package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("disbursement", func() {
	Title("Disbursement")

	Error("bad_request", StatusBadRequest, "invalid data was sent in the request")
	Error("internal_error", StatusInternalServerError, "check log for information")

	HTTP(func() {
		Path("/disbursement")

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
			Attribute("subscription_key", String, func() {
				Description("Subscription key which provides access to this API")
			})
			Username("api_key", String, "API Key", func() {
				Example("user")
			})
			Password("api_secret", String, "API Secret", func() {
				Example("password")
			})
			Required("api_key", "api_secret")
		})
		Result(OAuthTokenResponse)
		HTTP(func() {
			POST("/token")
			Headers(func() {
				Header("subscription_key:Ocp-Apim-Subscription-Key", String, func() {
					Description("Subscription key which provides access to this API")
				})
				Required("Ocp-Apim-Subscription-Key")
			})
			Response(StatusOK)
			Response(OAuthTokenError)
			Response("internal_error")
		})
	})

	// Get the balance of the account.
	Method("show", func() {
		Description("Get the balance of the account")
		Payload(func() {
			Token("access_token", String, func() {
				Description("JWT used for authentication")
				Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
			})
		})
		Result(Balance)
		HTTP(func() {
			GET("/v1_0/account/balance")
			Headers(func() {

				// Format of the header parameter follows the standard
				// for Basic and Bearer.
				// Oauth uses Bearer authentication type where the
				// credential is the received access token.
				Header("access_token:Authorization", String, func() {
					Description("Authorization header for Basic authentication & oauth")
					Pattern("^Bearer [^ ]+$")
					Example("Basic 67ew8n31me")
				})

				// This parameter is used to route the request to the
				// EWP system that will initiate the transaction.
				Header("X-Target-Environment", String, func() {
					Description("The identifier of the EWP system where the transaction shall be processed")
				})
				Header("Ocp-Apim-Subscription-Key", String, func() {
					Description("Subscription key which provides access to this API")
				})
				Required("X-Target-Environment", "Ocp-Apim-Subscription-Key")
			})
			Response(StatusOK)
			Response("bad_request")
			Response(ErrorReason)
		})
	})

	// Operation is used  to check if an account holder is registered and active in the system.
	Method("get", func() {
		Description("Checks if an account holder is registered and active in the system")
		Payload(func() {
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
			Result(String)
			HTTP(func() {
				GET("/v1_0/accountholder/{accountHolderIdType}/{accountHolderId}/active")

				Headers(func() {
					// Format of the header parameter follows the standard for Basic and Bearer.
					// Oauth uses Bearer authentication type where the credential is the received access token.
					Header("Authorization", String, func() {
						Description("Authorization header for Basic authentication & oauth")
						Pattern("^Bearer [^ ]+$")
						Example("Basic 67ew8n31me")
					})

					// This parameter is used to route the request to the EWP system that will initiate the transaction.
					Header("X-Target-Environment", String, func() {
						Description("The identifier of the EWP system where the transaction shall be processed")
					})

					Header("Ocp-Apim-Subscription-Key", String, func() {
						Description("Subscription key which provides access to this API")
					})

					Required("X-Target-Environment", "Ocp-Apim-Subscription-Key")
				})
				Response(StatusOK)
				Response("bad_request")
				Response(ErrorReason)
			})

		})
	})

	// The payer will be asked to authorize the payment.
	// The transaction will be executed once the payer has authorized the payment.
	// The requesttopay will be in status PENDING until the transaction is authorized
	// or declined by the payer or it is timed out by the system.
	// Status of the transaction can be validated by using the GET /requesttopay/ <resourceId>
	Method("post", func() {
		Description("Transfers an amount from the owner’s account to a payee account")
		Payload(Transfer)
		Result(String)
		HTTP(func() {
			POST("/v1_0/transfer")
			Headers(func() {

				// Format of the header parameter follows the standard for Basic and Bearer.
				// Oauth uses Bearer authentication type where the credential is the received access token.
				Header("Authorization", String, func() {
					Description(" Authorization header used for Basic authentication and oauth")
				})
				Header("X-Callback-Url", String, func() {
					Description("URL to the server where the callback should be sent.")
				})

				// This ID is used, for example, validating the status of the request.
				// ‘Universal Unique ID’ for the transaction generated using UUID version 4.
				Header("X-Reference-Id", String, func() {
					Description("Resource ID of the created request to pay transaction.")
				})

				// This parameter is used to route the request to the EWP
				// system that will initiate the transaction.
				Header("X-Target-Environment", String, func() {
					Description("The identifier of the EWP system where the transaction shall be processed.")
				})
				Header("Ocp-Apim-Subscription-Key", String, func() {
					Description("Subscription key which provides access to this API")
				})
				Required("X-Reference-Id", "X-Target-Environment", "Ocp-Apim-Subscription-Key")

			})
			Response(StatusAccepted)
			Response("bad_request")
			Response(ErrorReason)
		})
	})

	// Status of the transaction can validated by using the GET /transfer/\{referenceId\}
	// X-Reference-Id that was passed in the post is used as reference to the request
	Method("get", func() {
		Description("This operation is used to get the status of a transfer.")
		Payload(func() {
			Params(func() {

				// Reference id  used when creating the request to pay.
				// X-Reference-Id that was passed in the post is
				// used as reference to the request.
				Param("referenceId", String, func() {
					Description(" UUID of transaction to get result")
				})

			})
		})
		Result(TransferResult)
		HTTP(func() {
			GET("/v1_0/transfer/{referenceId}")
			Headers(func() {

				// Format of the header parameter follows the standard for Basic and Bearer.
				// Oauth uses Bearer authentication type where the credential is the received access token.
				Header("Authorization", String, func() {
					Description(" Authorization header used for Basic authentication and oauth")
				})

				// This parameter is used to route the request to the EWP
				// system that will initiate the transaction.
				Header("X-Target-Environment", String, func() {
					Description("The identifier of the EWP system where the transaction shall be processed.")
				})
				Header("Ocp-Apim-Subscription-Key", String, func() {
					Description("Subscription key which provides access to this API")
				})
				Required("X-Target-Environment", "Ocp-Apim-Subscription-Key")
			})

			// Note that a  failed request to pay will be returned with this status too.
			// The 'status' of the RequestToPayResult can be used to determine the outcome of the request.
			// The 'reason' field can be used to retrieve a cause in case of failure.
			Response(StatusOK)
			Response("bad_request")
			Response(ErrorReason)
		})
	})
})
