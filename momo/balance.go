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

	Attribute("availableBalance", String, "The available balance of the account")
	Attribute("currency", String, "ISO4217 Currency")
})

var BalanceHeader = Type("BalanceHeader", func() {

	// Format of the header parameter follows the standard for Basic and Bearer.
	// Oauth uses Bearer authentication type where the credential is the received access token.
	Attribute("Authorization", String, func() {
		Description("Authorization header used for Basic authentication and oauth.")
	})

	// This parameter is used to route the request to the EWP system that will initiate the transaction.
	Attribute("X-Target-Environment", String, func() {
		Description("The identifier of the EWP system where the transaction shall be processed.")
	})
})
