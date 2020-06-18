package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var ApiUser = Type("ApiUser", func() {
	Description("The create API user information")

	Attribute("providerCallbackHost", String, func() {
		Description("The provider callback host")
	})
})

var ApiUserResult = ResultType("ApiUserResult", func() {
	Description("The API user information")
	TypeName("ApiUserResult")
	ContentType("application/json")

	Attributes(func() {
		Attribute("providerCallbackHost", String, "The provider callback host")
		Attribute("paymentServerUrl", PaymentServerUrl)
		Attribute("targetEnvironment", TargetEnvironment)
	})
	View("default", func() {
		Attribute("providerCallbackHost")
		Attribute("paymentServerUrl")
		Attribute("targetEnvironment")
	})
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
