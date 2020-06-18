// Code generated by goa v3.1.3, DO NOT EDIT.
//
// collection HTTP server types
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package server

import (
	collection "github.com/wondenge/momo-go/gen/collection"
	collectionviews "github.com/wondenge/momo-go/gen/collection/views"
	goa "goa.design/goa/v3/pkg"
)

// PaymentRequestRequestBody is the type of the "collection" service
// "PaymentRequest" endpoint HTTP request body.
type PaymentRequestRequestBody struct {
	// Amount that will be debited from the payer account.
	Amount *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// ISO4217 Currency
	Currency *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
	// External id is used as a reference to the transaction.
	ExternalID *string           `form:"externalId,omitempty" json:"externalId,omitempty" xml:"externalId,omitempty"`
	Payer      *PartyRequestBody `form:"payer,omitempty" json:"payer,omitempty" xml:"payer,omitempty"`
	// Message that will be written in the payer transaction history message field.
	PayerMessage *string `form:"payerMessage,omitempty" json:"payerMessage,omitempty" xml:"payerMessage,omitempty"`
	// Message that will be written in the payee transaction history note field.
	PayeeNote *string `form:"payeeNote,omitempty" json:"payeeNote,omitempty" xml:"payeeNote,omitempty"`
}

// NewTokenResponseBody is the type of the "collection" service "NewToken"
// endpoint HTTP response body.
type NewTokenResponseBody struct {
	// A JWT token which can be used to authorize against the other API end-points.
	AccessToken *string `form:"access_token,omitempty" json:"access_token,omitempty" xml:"access_token,omitempty"`
	// The token type.
	TokenType *string `form:"token_type,omitempty" json:"token_type,omitempty" xml:"token_type,omitempty"`
	// The validity time in seconds of the token.
	ExpiresIn *int32 `form:"expires_in,omitempty" json:"expires_in,omitempty" xml:"expires_in,omitempty"`
}

// GetBalanceResponseBody is the type of the "collection" service "GetBalance"
// endpoint HTTP response body.
type GetBalanceResponseBody struct {
	// The available balance of the account
	AvailableBalance *string `form:"availableBalance,omitempty" json:"availableBalance,omitempty" xml:"availableBalance,omitempty"`
	// ISO4217 Currency
	Currency *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
}

// PaymentStatusResponseBody is the type of the "collection" service
// "PaymentStatus" endpoint HTTP response body.
type PaymentStatusResponseBody struct {
	// Amount that will be debited from the payer account.
	Amount *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// ISO4217 Currency
	Currency *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
	// Financial transactionIdd from mobile money manager.
	FinancialTransactionID *string `form:"financialTransactionId,omitempty" json:"financialTransactionId,omitempty" xml:"financialTransactionId,omitempty"`
	// External id provided in the creation of the requestToPay transaction.
	ExternalID *string            `form:"externalId,omitempty" json:"externalId,omitempty" xml:"externalId,omitempty"`
	Payer      *PartyResponseBody `form:"payer,omitempty" json:"payer,omitempty" xml:"payer,omitempty"`
	// Message that will be written in the payer transaction history message field.
	PayerMessage *string `form:"payerMessage,omitempty" json:"payerMessage,omitempty" xml:"payerMessage,omitempty"`
	// Message that will be written in the payee transaction history note field.
	PayeeNote *string `form:"payeeNote,omitempty" json:"payeeNote,omitempty" xml:"payeeNote,omitempty"`
	// Status
	Status *string                  `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	Reason *ErrorReasonResponseBody `form:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
}

// NewTokenUnauthorizedResponseBody is the type of the "collection" service
// "NewToken" endpoint HTTP response body for the "unauthorized" error.
type NewTokenUnauthorizedResponseBody struct {
	// An error code.
	Error *string `form:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
}

// NewTokenInternalErrorResponseBody is the type of the "collection" service
// "NewToken" endpoint HTTP response body for the "internal_error" error.
type NewTokenInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetBalanceBadRequestResponseBody is the type of the "collection" service
// "GetBalance" endpoint HTTP response body for the "bad_request" error.
type GetBalanceBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetBalanceInternalErrorResponseBody is the type of the "collection" service
// "GetBalance" endpoint HTTP response body for the "internal_error" error.
type GetBalanceInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RetrieveAccountHolderBadRequestResponseBody is the type of the "collection"
// service "RetrieveAccountHolder" endpoint HTTP response body for the
// "bad_request" error.
type RetrieveAccountHolderBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RetrieveAccountHolderInternalErrorResponseBody is the type of the
// "collection" service "RetrieveAccountHolder" endpoint HTTP response body for
// the "internal_error" error.
type RetrieveAccountHolderInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PaymentRequestBadRequestResponseBody is the type of the "collection" service
// "PaymentRequest" endpoint HTTP response body for the "bad_request" error.
type PaymentRequestBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PaymentRequestConflictResponseBody is the type of the "collection" service
// "PaymentRequest" endpoint HTTP response body for the "conflict" error.
type PaymentRequestConflictResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PaymentRequestInternalErrorResponseBody is the type of the "collection"
// service "PaymentRequest" endpoint HTTP response body for the
// "internal_error" error.
type PaymentRequestInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PaymentStatusBadRequestResponseBody is the type of the "collection" service
// "PaymentStatus" endpoint HTTP response body for the "bad_request" error.
type PaymentStatusBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PaymentStatusNotFoundResponseBody is the type of the "collection" service
// "PaymentStatus" endpoint HTTP response body for the "not_found" error.
type PaymentStatusNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PaymentStatusInternalErrorResponseBody is the type of the "collection"
// service "PaymentStatus" endpoint HTTP response body for the "internal_error"
// error.
type PaymentStatusInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PartyResponseBody is used to define fields on response body types.
type PartyResponseBody struct {
	// PartyIdType
	PartyIDType *string `form:"partyIdType,omitempty" json:"partyIdType,omitempty" xml:"partyIdType,omitempty"`
	PartyID     *string `form:"partyId,omitempty" json:"partyId,omitempty" xml:"partyId,omitempty"`
}

// ErrorReasonResponseBody is used to define fields on response body types.
type ErrorReasonResponseBody struct {
	// Code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PartyRequestBody is used to define fields on request body types.
type PartyRequestBody struct {
	// PartyIdType
	PartyIDType *string `form:"partyIdType,omitempty" json:"partyIdType,omitempty" xml:"partyIdType,omitempty"`
	PartyID     *string `form:"partyId,omitempty" json:"partyId,omitempty" xml:"partyId,omitempty"`
}

// NewNewTokenResponseBody builds the HTTP response body from the result of the
// "NewToken" endpoint of the "collection" service.
func NewNewTokenResponseBody(res *collectionviews.TokenPost200ApplicationJSONResponseView) *NewTokenResponseBody {
	body := &NewTokenResponseBody{
		AccessToken: res.AccessToken,
		TokenType:   res.TokenType,
		ExpiresIn:   res.ExpiresIn,
	}
	return body
}

// NewGetBalanceResponseBody builds the HTTP response body from the result of
// the "GetBalance" endpoint of the "collection" service.
func NewGetBalanceResponseBody(res *collectionviews.BalanceView) *GetBalanceResponseBody {
	body := &GetBalanceResponseBody{
		AvailableBalance: res.AvailableBalance,
		Currency:         res.Currency,
	}
	return body
}

// NewPaymentStatusResponseBody builds the HTTP response body from the result
// of the "PaymentStatus" endpoint of the "collection" service.
func NewPaymentStatusResponseBody(res *collectionviews.RequestToPayResultView) *PaymentStatusResponseBody {
	body := &PaymentStatusResponseBody{
		Amount:                 res.Amount,
		Currency:               res.Currency,
		FinancialTransactionID: res.FinancialTransactionID,
		ExternalID:             res.ExternalID,
		PayerMessage:           res.PayerMessage,
		PayeeNote:              res.PayeeNote,
		Status:                 res.Status,
	}
	if res.Payer != nil {
		body.Payer = marshalCollectionviewsPartyViewToPartyResponseBody(res.Payer)
	}
	if res.Reason != nil {
		body.Reason = marshalCollectionviewsErrorReasonViewToErrorReasonResponseBody(res.Reason)
	}
	return body
}

// NewNewTokenUnauthorizedResponseBody builds the HTTP response body from the
// result of the "NewToken" endpoint of the "collection" service.
func NewNewTokenUnauthorizedResponseBody(res *collection.TokenPost401ApplicationJSONResponse) *NewTokenUnauthorizedResponseBody {
	body := &NewTokenUnauthorizedResponseBody{
		Error: res.Error,
	}
	return body
}

// NewNewTokenInternalErrorResponseBody builds the HTTP response body from the
// result of the "NewToken" endpoint of the "collection" service.
func NewNewTokenInternalErrorResponseBody(res *goa.ServiceError) *NewTokenInternalErrorResponseBody {
	body := &NewTokenInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetBalanceBadRequestResponseBody builds the HTTP response body from the
// result of the "GetBalance" endpoint of the "collection" service.
func NewGetBalanceBadRequestResponseBody(res *goa.ServiceError) *GetBalanceBadRequestResponseBody {
	body := &GetBalanceBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetBalanceInternalErrorResponseBody builds the HTTP response body from
// the result of the "GetBalance" endpoint of the "collection" service.
func NewGetBalanceInternalErrorResponseBody(res *goa.ServiceError) *GetBalanceInternalErrorResponseBody {
	body := &GetBalanceInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRetrieveAccountHolderBadRequestResponseBody builds the HTTP response body
// from the result of the "RetrieveAccountHolder" endpoint of the "collection"
// service.
func NewRetrieveAccountHolderBadRequestResponseBody(res *goa.ServiceError) *RetrieveAccountHolderBadRequestResponseBody {
	body := &RetrieveAccountHolderBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRetrieveAccountHolderInternalErrorResponseBody builds the HTTP response
// body from the result of the "RetrieveAccountHolder" endpoint of the
// "collection" service.
func NewRetrieveAccountHolderInternalErrorResponseBody(res *goa.ServiceError) *RetrieveAccountHolderInternalErrorResponseBody {
	body := &RetrieveAccountHolderInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPaymentRequestBadRequestResponseBody builds the HTTP response body from
// the result of the "PaymentRequest" endpoint of the "collection" service.
func NewPaymentRequestBadRequestResponseBody(res *goa.ServiceError) *PaymentRequestBadRequestResponseBody {
	body := &PaymentRequestBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPaymentRequestConflictResponseBody builds the HTTP response body from the
// result of the "PaymentRequest" endpoint of the "collection" service.
func NewPaymentRequestConflictResponseBody(res *goa.ServiceError) *PaymentRequestConflictResponseBody {
	body := &PaymentRequestConflictResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPaymentRequestInternalErrorResponseBody builds the HTTP response body
// from the result of the "PaymentRequest" endpoint of the "collection" service.
func NewPaymentRequestInternalErrorResponseBody(res *goa.ServiceError) *PaymentRequestInternalErrorResponseBody {
	body := &PaymentRequestInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPaymentStatusBadRequestResponseBody builds the HTTP response body from
// the result of the "PaymentStatus" endpoint of the "collection" service.
func NewPaymentStatusBadRequestResponseBody(res *goa.ServiceError) *PaymentStatusBadRequestResponseBody {
	body := &PaymentStatusBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPaymentStatusNotFoundResponseBody builds the HTTP response body from the
// result of the "PaymentStatus" endpoint of the "collection" service.
func NewPaymentStatusNotFoundResponseBody(res *goa.ServiceError) *PaymentStatusNotFoundResponseBody {
	body := &PaymentStatusNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPaymentStatusInternalErrorResponseBody builds the HTTP response body from
// the result of the "PaymentStatus" endpoint of the "collection" service.
func NewPaymentStatusInternalErrorResponseBody(res *goa.ServiceError) *PaymentStatusInternalErrorResponseBody {
	body := &PaymentStatusInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRetrieveAccountHolderPayload builds a collection service
// RetrieveAccountHolder endpoint payload.
func NewRetrieveAccountHolderPayload(accountHolderIDType string, accountHolderID string) *collection.RetrieveAccountHolderPayload {
	v := &collection.RetrieveAccountHolderPayload{}
	v.AccountHolderIDType = accountHolderIDType
	v.AccountHolderID = accountHolderID

	return v
}

// NewPaymentRequestRequestToPay builds a collection service PaymentRequest
// endpoint payload.
func NewPaymentRequestRequestToPay(body *PaymentRequestRequestBody) *collection.RequestToPay {
	v := &collection.RequestToPay{
		Amount:       body.Amount,
		Currency:     body.Currency,
		ExternalID:   body.ExternalID,
		PayerMessage: body.PayerMessage,
		PayeeNote:    body.PayeeNote,
	}
	if body.Payer != nil {
		v.Payer = unmarshalPartyRequestBodyToCollectionParty(body.Payer)
	}

	return v
}

// NewPaymentStatusPayload builds a collection service PaymentStatus endpoint
// payload.
func NewPaymentStatusPayload(referenceID string) *collection.PaymentStatusPayload {
	v := &collection.PaymentStatusPayload{}
	v.ReferenceID = referenceID

	return v
}

// ValidatePaymentRequestRequestBody runs the validations defined on
// PaymentRequestRequestBody
func ValidatePaymentRequestRequestBody(body *PaymentRequestRequestBody) (err error) {
	if body.Payer != nil {
		if err2 := ValidatePartyRequestBody(body.Payer); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidatePartyRequestBody runs the validations defined on PartyRequestBody
func ValidatePartyRequestBody(body *PartyRequestBody) (err error) {
	if body.PartyIDType != nil {
		if !(*body.PartyIDType == "MSISDN" || *body.PartyIDType == "EMAIL" || *body.PartyIDType == "PARTY_CODE") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.partyIdType", *body.PartyIDType, []interface{}{"MSISDN", "EMAIL", "PARTY_CODE"}))
		}
	}
	return
}
