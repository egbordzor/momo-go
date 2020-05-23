package design

import (
	. "github.com/wondenge/momo-go/momo"
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("token", func() {

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
})
