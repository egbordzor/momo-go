package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Validate Account Holder
// Validate account holder can be used to do a validation if a customer is
// active and able to receive funds.
// The use case will only validate that the customer is available and active.
// It does not validate that a specific amount can be received.
//
// The sequence for the validate account holder is described below.
// 1. The Partner can send a GET /accountholder request to validate is a customer is active.
// The Partner provides the id of that customer as part of the URL
// 2. Wallet platform will respond with HTTP 200 if the account holder is active.

var AccountHolderHeader = Type("AccountHolderHeader", func() {

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
