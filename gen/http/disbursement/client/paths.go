// Code generated by goa v3.1.3, DO NOT EDIT.
//
// HTTP request path constructors for the Disbursement service.
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package client

import (
	"fmt"
)

// NewTokenDisbursementPath returns the URL path to the Disbursement service NewToken HTTP endpoint.
func NewTokenDisbursementPath() string {
	return "/disbursement/token"
}

// GetBalanceDisbursementPath returns the URL path to the Disbursement service GetBalance HTTP endpoint.
func GetBalanceDisbursementPath() string {
	return "/disbursement/v1_0/account/balance"
}

// RetrieveAccountHolderDisbursementPath returns the URL path to the Disbursement service RetrieveAccountHolder HTTP endpoint.
func RetrieveAccountHolderDisbursementPath(accountHolderIDType string, accountHolderID string) string {
	return fmt.Sprintf("/disbursement/v1_0/accountholder/%v/%v/active", accountHolderIDType, accountHolderID)
}

// TransferDisbursementPath returns the URL path to the Disbursement service Transfer HTTP endpoint.
func TransferDisbursementPath() string {
	return "/disbursement/v1_0/transfer"
}

// TransferStatusDisbursementPath returns the URL path to the Disbursement service TransferStatus HTTP endpoint.
func TransferStatusDisbursementPath(referenceID string) string {
	return fmt.Sprintf("/disbursement/v1_0/transfer/%v", referenceID)
}