package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Partner Gateway sandbox provisioning API document
var _ = Service("User", func() {

	HTTP(func() {
		Path("/")
	})

	// Create API User
	// a) The Provider sends a POST {baseURL}/apiuser request to Wallet platform.
	// b) The Provider specifies the UUID Reference ID in the request Header and the subscription Key.
	// c) Reference ID will be used as the User ID for the API user to be created.
	// d) Wallet Platform creates the User and responds with 201
	Method("NewUser", func() {
		Description("Used to create an API user in the sandbox target environment.")
		Payload(ApiUser)
		Result(String)

		// 400
		Error("bad_request", ErrorResult, "Bad request, e.g. invalid data was sent in the request.")

		// 409
		Error("conflict", ErrorReason, "Conflict, duplicated reference id")

		// 500
		Error("internal_error", ErrorResult, "Internal error. Check log for information.")

		HTTP(func() {
			POST("/v1_0/apiuser")

			Response(StatusCreated)

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 409
			// RFC 7231, 6.5.8
			Response("conflict", StatusConflict)

			// Status 500
			// RFC 7231, 6.6.1
			Response("internal_error", StatusInternalServerError)
		})
	})

	// Create API Key
	// a) The Provider sends a POST {baseURL}/apiuser/{APIUser}/apikey request to Wallet platform.
	// b) The Provider specifies the API User in the URL and subscription Key in the header.
	// c) Wallet Platform creates the API Key and responds with 201 Created with the newly
	// Created API Key in the Body.
	// d) Provider now has both API User and API Key created.
	Method("NewKey", func() {
		Description("Used to create an API key for an API user in the sandbox target environment.")
		Payload(func() {

			// Resource ID for the API user to be created.
			// UUID version 4 is required
			Attribute("X-Reference-Id", String, func() {
				Description("Format - UUID.")
			})
			Required("X-Reference-Id")
		})
		Result(ApiUserKeyResult)

		// 400
		Error("bad_request", ErrorResult, "Bad request, e.g. invalid data was sent in the request.")

		// 404
		Error("not_found", ErrorReason, "Not found, reference id not found or closed in sandbox")

		// 500
		Error("internal_error", ErrorResult, "Internal error. Check log for information.")

		HTTP(func() {
			POST("/v1_0/apiuser/{X-Reference-Id}/apikey")

			Response(StatusCreated)

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 404
			// RFC 7231, 6.5.4
			Response("not_found", StatusNotFound)

			// Status 500
			// RFC 7231, 6.6.1
			Response("internal_error", StatusInternalServerError)
		})
	})

	Method("GetUser", func() {
		Description("Used to get API user information.")
		Payload(func() {

			// Resource ID for the API user to be created.
			// UUID version 4 is required
			Attribute("X-Reference-Id", String, func() {
				Description("Format - UUID.")
			})
			Required("X-Reference-Id")
		})
		Result(String)

		// 400
		Error("bad_request", ErrorResult, "Bad request, e.g. invalid data was sent in the request.")

		// 404
		Error("not_found", ErrorResult, "Not found, reference id not found or closed in sandbox")

		// 500
		Error("internal_error", ErrorResult, "Internal error. Check log for information.")

		HTTP(func() {
			GET("/v1_0/apiuser/{X-Reference-Id}")

			Response(StatusOK)

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 404
			// RFC 7231, 6.5.4
			Response("not_found", StatusNotFound)

			// Status 500
			// RFC 7231, 6.6.1
			Response("internal_error", StatusInternalServerError)

		})
	})

	// GET API User Details
	// a) The Provider sends a GET {baseURL}/apiuser/{APIUser} request to Wallet platform.
	// b) The Provider specifies the API User in the URL and subscription Key in the header.
	// c) Wallet Platform responds with 200 Ok and details of the user.
	// d) TargetEnvironment is preconfigured to sandbox in the Sandbox environment,
	// therefore Providers will not have the option of setting it to a different parameter.
	Method("GetUserDetails", func() {
		Description("GET API User Details")
		Payload(String)
		Result(ApiUserResult)

		// 400
		Error("bad_request", ErrorResult, "Bad request, e.g. invalid data was sent in the request.")

		// 500
		Error("internal_error", ErrorResult, "Internal error. Check log for information.")

		HTTP(func() {
			GET("/apiuser/{APIUser}")

			Response(StatusOK)
			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 500
			// RFC 7231, 6.6.1
			Response("internal_error", StatusInternalServerError)
		})
	})
})
