package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var ErrorReason = ResultType("ErrorReason", func() {
	Attribute("code", String, func() {
		Enum(
			"PAYEE_NOT_FOUND",
			"PAYER_NOT_FOUND",
			"NOT_ALLOWED",
			"NOT_ALLOWED_TARGET_ENVIRONMENT",
			"INVALID_CALLBACK_URL_HOST",
			"INVALID_CURRENCY",
			"SERVICE_UNAVAILABLE",
			"INTERNAL_PROCESSING_ERROR",
			"NOT_ENOUGH_FUNDS",
			"PAYER_LIMIT_REACHED",
			"PAYEE_NOT_ALLOWED_TO_RECEIVE",
			"PAYMENT_NOT_APPROVED",
			"RESOURCE_NOT_FOUND",
			"APPROVAL_REJECTED",
			"EXPIRED",
			"TRANSACTION_CANCELED",
			"RESOURCE_ALREADY_EXIST",
		)
		Example("PAYER_NOT_FOUND")
	})
	Attribute("message", String, func() {
		Description("message")
		Example("Payee does not exist")
	})
})
