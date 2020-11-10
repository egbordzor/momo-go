/*
 * Collection
 *
 * Partner Gateway API document
 *
 * API version: 1.0
 *
 */

package collection

// RequestToPay struct for RequestToPay
type RequestToPay struct {
	// Amount that will be debited from the payer account.
	Amount string `json:"amount,omitempty"`
	// ISO4217 Currency
	Currency string `json:"currency,omitempty"`
	// External id is used as a reference to the transaction. External id is used for reconciliation. The external id will be included in transaction history report. <br>External id is not required to be unique.
	ExternalId string `json:"externalId,omitempty"`
	Payer      Party  `json:"payer,omitempty"`
	// Message that will be written in the payer transaction history message field.
	PayerMessage string `json:"payerMessage,omitempty"`
	// Message that will be written in the payee transaction history note field.
	PayeeNote string `json:"payeeNote,omitempty"`
}
