package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Request to Pay
// Request to Pay service is used for requesting a payment from a customer (Payer).
// This can be used by e.g. an online web shop to request a payment for a customer.
// The customer is requested to approve the transaction on the customer client.
// 1. Customer (Payer) have selected product(s) in the merchant web shop and decided to check out.
// Customer select to pay with Mobile Money.
// 2. The provider system collects the account information for the customer
// e.g. mobile number and calculate the total amount of the products.
// 3. The provider system sends a request to pay (POST /requesttopay) operation to Wallet Platform.
// This request includes the amount and customer (Payer) account holder number.
// 4. Wallet Platform will respond with HTTP 202 Accepted to the provider system
// 5. Provider shall inform the customer that a payment needs to be approved,
// by giving information on the merchant web page.
// For example, the merchant could show information that payment is being processed
// and that customer needs to approve using the own client, e.g. USSD, mobile app.
// 6. Wallet Platform will process the request so that the customer can approve the payment.
// The request to pay will be in PENDING state until the customer have approved/Rejected the payment.
// 7. The Customer (Payer) will use his/her own client to review the payment.
// Customer can approve or reject the payment.
// 8. Wallet platform will transfer the funds if the customer approves the payment.
// Status of the payment is updated to SUCCESSFUL or FAILED.
// 9. If a callback URL was provided in the POST /requesttopay then a callback
// will be sent once the request to pay have reached a final state (SUCCESSFUL, FAILED).
// Note the callback will only be sent once. There is no retry.
// 10. GET request can be used for validating the status of the transaction.
// GET is used if the partner system has not requested a callback by providing a
// callback URL or if the callback was not received.

var RequestToPay = Type("RequestToPay", func() {

	Attribute("amount", String, func() {
		Description("Amount that will be debited from the payer account.")
	})
	Attribute("currency", String, func() {
		Description("ISO4217 Currency")
	})

	// External id is used for reconciliation.
	// The external id will be included in transaction history report.
	// <br>External id is not required to be unique
	Attribute("externalId", String, func() {
		Description("External id is used as a reference to the transaction.")
	})
	// TODO
	// $ref: '#/components/schemas/Party'
	Attribute("payer", Party)
	Attribute("payerMessage", String, func() {
		Description("Message that will be written in the payer transaction history message field.")
	})
	Attribute("payeeNote", String, func() {
		Description("Message that will be written in the payee transaction history note field.")
	})
})

var RequestToPayResult = ResultType("RequestToPayResult", func() {
	Attribute("amount", Int, func() {
		Description("Amount that will be debited from the payer account.")
		Example(100)
	})
	Attribute("currency", String, func() {
		Description("ISO4217 Currency")
		Example("UGX")
	})

	// <br> Used to connect to the specific financial transaction made in the account
	Attribute("financialTransactionId", Int, func() {
		Description("Financial transactionIdd from mobile money manager.")
		Example(23503452)
	})
	Attribute("externalId", Int, func() {
		Description("External id provided in the creation of the requestToPay transaction.")
		Example(947354)
	})
	// TODO
	// $ref: '#/components/schemas/Party'
	Attribute("payer", Party)
	Attribute("payerMessage", String, func() {
		Description("Message that will be written in the payer transaction history message field.")
	})
	Attribute("payeeNote", String, func() {
		Description("Message that will be written in the payee transaction history note field.")
	})
	Attribute("status", String, func() {
		Enum("PENDING", "SUCCESSFUL", "FAILED")
		Example("SUCCESSFUL")
	})
	// TODO
	// $ref: '#/components/schemas/ErrorReason'
	Attribute("reason", ErrorReason)
})
