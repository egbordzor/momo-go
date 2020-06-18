package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var Party = Type("Party", func() {
	Description("Party identifies a account holder in the wallet platform.")

	Attribute("partyIdType", String, func() {
		Description("PartyIdType")
		Enum("MSISDN", "EMAIL", "PARTY_CODE")
	})

	// Party consists of two parameters, type and partyId.
	// Each type have its own validation of the partyId<br> MSISDN - Mobile Number validated according to ITU-T E.164.
	// Validated with IsMSISDN<br> EMAIL - Validated to be a valid e-mail format.
	// Validated with IsEmail<br> PARTY_CODE - UUID of the party.
	// Validated with IsUuid
	Attribute("partyId", String, func() {
	})
})
