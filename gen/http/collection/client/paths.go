// Code generated by goa v3.1.3, DO NOT EDIT.
//
// HTTP request path constructors for the Collection service.
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package client

import (
	"fmt"
)

// NewTokenCollectionPath returns the URL path to the Collection service NewToken HTTP endpoint.
func NewTokenCollectionPath() string {
	return "/collection/token"
}

// GetBalanceCollectionPath returns the URL path to the Collection service GetBalance HTTP endpoint.
func GetBalanceCollectionPath() string {
	return "/collection/v1_0/account/balance"
}

// RetrieveAccountHolderCollectionPath returns the URL path to the Collection service RetrieveAccountHolder HTTP endpoint.
func RetrieveAccountHolderCollectionPath(accountHolderIDType string, accountHolderID string) string {
	return fmt.Sprintf("/collection/v1_0/accountholder/%v/%v/active", accountHolderIDType, accountHolderID)
}

// PaymentRequestCollectionPath returns the URL path to the Collection service PaymentRequest HTTP endpoint.
func PaymentRequestCollectionPath() string {
	return "/collection/v1_0/requesttopay"
}

// PaymentStatusCollectionPath returns the URL path to the Collection service PaymentStatus HTTP endpoint.
func PaymentStatusCollectionPath(referenceID string) string {
	return fmt.Sprintf("/collection/v1_0/requesttopay/%v", referenceID)
}