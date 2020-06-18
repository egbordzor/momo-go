package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Oauth 2.0
// The Open API is using Oauth 2.0 token for authentication of request.
// Client will request an access token using Client Credential Grant according to RFC 6749.
// The token received is according to RFC 6750 Bearer Token.
// The API user and API key are used in the basic authentication header when requesting the access token.
// The API user and key are managed in the Partner GUI for the country where the account is located.
// The Partner can create and manage API user and key from the Partner GUI.
// In the case of the Sandbox, the API Key and API User are managed through the Provisioning API
// The received token has an expiry time.
// The same token can be used for transactions until it expires.
// A new token is requested by using the POST /token service in the same way as the initial token.
// The new token can be requested for before the previous one has expired to avoid authentication
// failure due to expired token.
// The token must be treated as a credential and kept secret.
// The party that have access to the token will be authenticated as the user that requested the token.
// The below sequence describes the flow for requesting a token and using the token in a request.

var OAuthTokenResponse = ResultType("OAuthTokenResponse", func() {

	// The format of the token follows the JWT standard format
	// (see jwt.io for an example).
	// This is the token that should be sent in in the Authorization
	// header when calling the other API  end-points.
	Attribute("access_token", String, func() {
		Description("A JWT token which can be used to authorize against the other API end-points.")
	})
	Attribute("token_type", String, "The token type")
	Attribute("expires_in", Int, "The validity time in seconds of the token.")
})

var OAuthTokenError = Type("OAuthTokenError", func() {
	Attribute("error", String, func() {
		Description("An error code.")
	})
})

// OAuth Token
// OAuth Token is generated from the merchants’ API Key and Secret.
// The API Key and API Secret can be obtained through the provisioning API in Sandbox.

// BasicAuth defines a security scheme using basic authentication. The scheme
// protects the "signin" action used to create JWTs.
var BasicAuth = BasicAuthSecurity("basic", func() {
	Description("Basic authentication used to authenticate security principal during signin")
	Scope("api:read", "Read-only access")
})
