package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Get Balance
// Get balance request is used to check the balance on the default account connected to the API User.
// The following is the sequence flow for get balance use case.
// 1. The partner will send a GET /account/balance request
// 2. Wallet platform will respond with the available balance on the API user account.
var Balance = ResultType("Balance", func() {
	Description("The available balance of the account")
	TypeName("Balance")
	ContentType("application/json")

	Attributes(func() {
		Attribute("availableBalance", String, "The available balance of the account")
		Attribute("currency", String, "ISO4217 Currency")
	})

	View("default", func() {
		Attribute("availableBalance")
		Attribute("currency")
	})
})
