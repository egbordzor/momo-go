package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Transfer
// Transfer is used for transferring money from the provider account to a customer.
//
// The below sequence gives an overview of the flow of the transfer use case.
// 1. The Provider sends a POST /transfer request to Wallet platform.
// 2. Wallet platform will directly respond to indicate that the request is received and will be processed.
// 3. Wallet platform will authorize the request to ensure that the transfer is allowed.
// The funds will be transferred from the provider account to the Payee account provided in the transfer request.
// 4. Callback will be sent if a callback URL was provided in the POST request.
// The callback is sent when the request has reach a final state (SUCCESSFUL, FAILED).
// 5. The Provider can use the GET request to validate the status of the transfer.
var Transfer = Type("Transfer", func() {
	Description("Transfer")

	Attribute("amount", String, func() {
		Description("Amount that will be debited from the payer account.")
	})
	Attribute("currency", String, func() {
		Description("ISO4217 Currency")
	})

	// External id is used for reconciliation.
	// The external id will be included in transaction history report.
	// External id is not required to be unique.
	Attribute("externalId", String, func() {
		Description("External id is used as a reference to the transaction.")
	})
	Attribute("payee", Party)
	Attribute("payerMessage", String, func() {
		Description("Message that will be written in the payer transaction history message field.")
	})
	Attribute("payeeNote", String, func() {
		Description("Message that will be written in the payee transaction history note field.")
	})
})

var TransferResult = ResultType("TransferResult", func() {
	Description("Transfer Result")
	TypeName("TransferResult")
	ContentType("application/json")

	Attributes(func() {
		Attribute("amount", String, func() {
			Description(" Amount that will be debited from the payer account.")
		})
		Attribute("currency", String, func() {
			Description("ISO4217 Currency")
		})

		// Used to connect to the specific financial transaction made in the account
		Attribute("financialTransactionId", String, func() {
			Description("Financial transactionIdd from mobile money manager")
		})

		// External id is used for reconciliation.
		// The external id will be included in transaction history report.
		// External id is not required to be unique.
		Attribute("externalId", String, func() {
			Description("External id is used as a reference to the transaction")
		})
		Attribute("payee", Party)
		Attribute("payerMessage", String, func() {
			Description("Message that will be written in the payer transaction history message field.")
		})
		Attribute("payeeNote", String, func() {
			Description("Message that will be written in the payee transaction history note field.")
		})
		Attribute("status", String, func() {
			Description("Status")
			Enum("PENDING", "SUCCESSFUL", "FAILED")
		})
		Attribute("reason", ErrorReason)
	})

	View("default", func() {
		Attribute("amount")
		Attribute("currency")
		Attribute("financialTransactionId")
		Attribute("externalId")
		Attribute("payee")
		Attribute("payeeNote")
		Attribute("status")
		Attribute("reason")
	})
})
