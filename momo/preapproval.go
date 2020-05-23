package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Pre-Approval
// Pre-approval is used to setup an auto debit towards a customer.
// The Partner can request a pre-approval from the customer.
// Once the customer has approved then the partner can debit the customer account
// without authorization from the customer.
//
// The call flow for setting up a pre-approval is like the request to pay use case.
// The following picture describes the sequence for pre-approval.
// 1. The Provider sends a POST /preapproval request to Wallet platform.
// 2. Provider shall inform the customer that pre-approval needs to be approved.
// 3. Customer (Payer) will use the own client to view the pre-approval request.
// Customer can approve or reject the request.
// 4. Callback will be sent if a callback URL was provided in the POST request.
// The callback is sent when the request has reach a final state (Successful, Failed).
// 5. The Provider can use the GET request to validate the status of the pre-approval.
var PreApproval = Type("PreApproval", func() {

	// TODO
	// $ref: '#/components/schemas/Party'
	Attribute("payer", Party)
	Attribute("payerCurrency", String, func() {
		Description("ISO4217 Currency")
	})
	Attribute("payerMessage", String, func() {
		Description("The message that is shown to the approver.")
	})
	Attribute("validityTime", Int, func() {
		Description("The request validity time of the pre-approval")
	})
})

var PreApprovalResult = ResultType("PreApprovalResult", func() {

	// TODO
	// $ref: '#/components/schemas/Party'
	Attribute("payer", Party)
	Attribute("payerCurrency", String, func() {
		Description("ISO4217 Currency")
	})
	Attribute("payerMessage", String, func() {
		Description("The message that is shown to the approver.")
	})
	Attribute("validityTime", Int, func() {
		Description("The request validity time of the pre-approval")
	})
	Attribute("status", String, func() {
		Enum("PENDING", "SUCCESSFUL", "FAILED")
	})
	// TODO
	// // $ref: '#/components/schemas/ErrorReason'
	Attribute("reason", ErrorReason)
})

