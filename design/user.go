package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Partner Gateway sandbox provisioning API document
var _ = Service("user", func() {
	Title("Sandbox User Provisioning")

	Error("bad_request", StatusBadRequest, "invalid data was sent in the request")
	Error("internal_error", StatusInternalServerError, "check log for information")

	HTTP(func() {
		Path("/v1_0/apiuser")
	})

	// Create API User
	// a) The Provider sends a POST {baseURL}/apiuser request to Wallet platform.
	// b) The Provider specifies the UUID Reference ID in the request Header and the subscription Key.
	// c) Reference ID will be used as the User ID for the API user to be created.
	// d) Wallet Platform creates the User and responds with 201
	Method("create", func() {
		Description("Used to create an API user in the sandbox target environment")
		Payload(ApiUser)
		Result(String)
		HTTP(func() {
			POST("/")
			Headers(func() {

				// This ID is used, for example,
				// validating the status of the request.
				// ‘Universal Unique ID’ for the transaction generated using UUID version 4.
				Header("X-Reference-Id", String, func() {
					Description("Resource ID of the created request to pay transaction.")
					Example("c72025f5-5cd1-4630-99e4-8ba4722fad56")
				})

				Header("Ocp-Apim-Subscription-Key", String, func() {
					Description("Subscription key which provides access to this API")
					Example("d484a1f0d34f4301916d0f2c9e9106a2")
				})
				Required("X-Reference-Id", "Ocp-Apim-Subscription-Key")
			})
			Response(StatusCreated)
			Response("bad_request")
			Response("internal_error")
			Response(ErrorReason)
		})
	})

	// Create API Key
	// a) The Provider sends a POST {baseURL}/apiuser/{APIUser}/apikey request to Wallet platform.
	// b) The Provider specifies the API User in the URL and subscription Key in the header.
	// c) Wallet Platform creates the API Key and responds with 201 Created with the newly
	// Created API Key in the Body.
	// d) Provider now has both API User and API Key created.
	Method("create", func() {
		Description("Used to create an API key for an API user in the sandbox target environment.")
		Payload(ApiUserKeyResult)
		Result(String)
		HTTP(func() {
			POST("/{X-Reference-Id}/apikey")
			Params(func() {

				// This ID is used, for example,
				// validating the status of the request.
				// ‘Universal Unique ID’ for the transaction generated using UUID version 4.
				Param("X-Reference-Id", String, func() {
					Description("Resource ID of the created request to pay transaction.")
					Example("c72025f5-5cd1-4630-99e4-8ba4722fad56")
				})
				Required("X-Reference-Id")
			})
			Headers(func() {
				Header("Ocp-Apim-Subscription-Key", String, func() {
					Description("Subscription key which provides access to this API")
					Example("d484a1f0d34f4301916d0f2c9e9106a2")
				})
				Required("Ocp-Apim-Subscription-Key")
			})
			Response(StatusCreated)
			Response("bad_request")
			Response("internal_error")
			Response(ErrorReason)
		})
	})

	Method("show", func() {
		Description("Used to get API user information.")
		Payload(func() {
			Params(func() {

				// This ID is used, for example,
				// validating the status of the request.
				// ‘Universal Unique ID’ for the transaction generated using UUID version 4.
				Param("X-Reference-Id", String, func() {
					Description("Resource ID of the created request to pay transaction.")
					Example("c72025f5-5cd1-4630-99e4-8ba4722fad56")
				})
				Required("X-Reference-Id")
			})
		})
		Result(String)
		HTTP(func() {
			GET("/{X-Reference-Id}")
			Headers(func() {
				Header("Ocp-Apim-Subscription-Key", String, func() {
					Description("Subscription key which provides access to this API")
					Example("d484a1f0d34f4301916d0f2c9e9106a2")
				})
				Required("Ocp-Apim-Subscription-Key")
			})
			Response(StatusOK)
			Response("bad_request")
			Response("internal_error")
			Response(ErrorReason)
		})
	})

	// GET API User Details
	// a) The Provider sends a GET {baseURL}/apiuser/{APIUser} request to Wallet platform.
	// b) The Provider specifies the API User in the URL and subscription Key in the header.
	// c) Wallet Platform responds with 200 Ok and details of the user.
	// d) TargetEnvironment is preconfigured to sandbox in the Sandbox environment,
	// therefore Providers will not have the option of setting it to a different parameter.
	Method("show", func() {
		Description("GET API User Details")
		Payload(func() {
			Params(func() {
				Param("APIUser", String, func() {
					Description("API User.")
					Example("c72025f5-5cd1-4630-99e4-8ba4722fad56")
				})
				Required("APIUser")
			})
		})
		Result(ApiUserResult)
		HTTP(func() {
			GET("/apiuser/{APIUser}")
			Headers(func() {
				Header("Ocp-Apim-Subscription-Key", String, func() {
					Description("Subscription key which provides access to this API")
					Example("d484a1f0d34f4301916d0f2c9e9106a2")
				})
				Required("Ocp-Apim-Subscription-Key")
			})
			Response(StatusOK)
			Response("bad_request")
			Response("internal_error")
			Response(ErrorReason)
		})
	})
})

var ApiUser = Type("ApiUser", func() {
	Description("The create API user information")
	Attribute("providerCallbackHost", String, func() {
		Description("The provider callback host")
	})
})

var ApiUserResult = Type("ApiUserResult", func() {
	Description("The API user information")
	Attribute("providerCallbackHost", String, "The provider callback host")
	Attribute("paymentServerUrl", PaymentServerUrl)
	Attribute("targetEnvironment", TargetEnvironment)
})

var ApiUserKeyResult = Type("ApiUserKeyResult", func() {
	Attribute("apiKey", String, "The created API user key")
})

var PaymentServerUrl = Type("PaymentServerUrl", func() {
	Attribute("apiKey", String, "The payment server URL")
})

var TargetEnvironment = Type("TargetEnvironment", func() {
	Attribute("apiKey", String, "The target environment")
})
