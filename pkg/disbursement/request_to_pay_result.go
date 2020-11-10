/*
 * Disbursements
 *
 * Partner Gateway API document
 *
 * API version: 1.0
 *
 */

package disbursement

// RequestToPayResult struct for RequestToPayResult
type RequestToPayResult struct {
	// Amount that will be debited from the payer account.
	Amount string `json:"amount,omitempty"`
	// ISO4217 Currency
	Currency string `json:"currency,omitempty"`
	// Financial transactionIdd from mobile money manager.<br> Used to connect to the specific financial transaction made in the account
	FinancialTransactionId string `json:"financialTransactionId,omitempty"`
	// External id provided in the creation of the requestToPay transaction.
	ExternalId string `json:"externalId,omitempty"`
	Payer      Party  `json:"payer,omitempty"`
	// Message that will be written in the payer transaction history message field.
	PayerMessage string `json:"payerMessage,omitempty"`
	// Message that will be written in the payee transaction history note field.
	PayeeNote string      `json:"payeeNote,omitempty"`
	Status    string      `json:"status,omitempty"`
	Reason    ErrorReason `json:"reason,omitempty"`
}
