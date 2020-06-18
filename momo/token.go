package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var TokenPost200ApplicationJsonResponse = ResultType("TokenPost200ApplicationJsonResponse", func() {

	Attribute("access_token", String, func() {

		// The format of the token follows the JWT standard format (see jwt.io for an example).
		// This is the token that should be sent in in the Authorization header when calling the other API end-points.
		Description("A JWT token which can be used to authorize against the other API end-points.")
	})
	Attribute("token_type", String, func() {
		Description("The token type.")
	})
	Attribute("expires_in", Int32, func() {
		Description("The validity time in seconds of the token.")
	})

})

var TokenPost401ApplicationJsonResponse = ResultType("TokenPost401ApplicationJsonResponse", func() {
	Attribute("error", String, func() {
		Description("An error code.")
	})
})
