package disbursements_go

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new disbursements go API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for disbursements go API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAccountBalance(params *GetAccountBalanceParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountBalanceOK, error)

	GetAccountholderAccountholderidtypeAccountholderidActive(params *GetAccountholderAccountholderidtypeAccountholderidActiveParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountholderAccountholderidtypeAccountholderidActiveOK, error)

	TokenPOST(params *TokenPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*TokenPOSTOK, error)

	TransferPOST(params *TransferPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*TransferPOSTAccepted, error)

	TransferReferenceIDGET(params *TransferReferenceIDGETParams, authInfo runtime.ClientAuthInfoWriter) (*TransferReferenceIDGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAccountBalance v1s 0 account balance g e t

  Get the balance of the account.
*/
func (a *Client) GetAccountBalance(params *GetAccountBalanceParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountBalanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountBalanceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-v1_0-account-balance",
		Method:             "GET",
		PathPattern:        "/v1_0/account/balance",
		ProducesMediaTypes: []string{"Incorrect target environment", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountBalanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountBalanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-v1_0-account-balance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAccountholderAccountholderidtypeAccountholderidActive v1s 0 accountholder account holder Id type account holder Id active g e t

  Operation is used  to check if an account holder is registered and active in the system.
*/
func (a *Client) GetAccountholderAccountholderidtypeAccountholderidActive(params *GetAccountholderAccountholderidtypeAccountholderidActiveParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountholderAccountholderidtypeAccountholderidActiveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountholderAccountholderidtypeAccountholderidActiveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-v1_0-accountholder-accountholderidtype-accountholderid-active",
		Method:             "GET",
		PathPattern:        "/v1_0/accountholder/{accountHolderIdType}/{accountHolderId}/active",
		ProducesMediaTypes: []string{"Incorrect target environment"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountholderAccountholderidtypeAccountholderidActiveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountholderAccountholderidtypeAccountholderidActiveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-v1_0-accountholder-accountholderidtype-accountholderid-active: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TokenPOST tokens p o s t

  This operation is used to create an access token which can then be used to authorize and authenticate towards the other end-points of the API.
*/
func (a *Client) TokenPOST(params *TokenPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*TokenPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTokenPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "token-POST",
		Method:             "POST",
		PathPattern:        "/token/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TokenPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TokenPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for token-POST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TransferPOST transfers p o s t

  Transfer operation is used to transfer an amount from the owner’s account to a payee account.<br> Status of the transaction can be validated by using the GET /transfer/\{referenceId\}
*/
func (a *Client) TransferPOST(params *TransferPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*TransferPOSTAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTransferPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "transfer-POST",
		Method:             "POST",
		PathPattern:        "/v1_0/transfer",
		ProducesMediaTypes: []string{"Incorrect currency for target environment", "ReferenceId already in use", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TransferPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TransferPOSTAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for transfer-POST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TransferReferenceIDGET transfers reference Id g e t

  This operation is used to get the status of a transfer. X-Reference-Id that was passed in the post is used as reference to the request.
*/
func (a *Client) TransferReferenceIDGET(params *TransferReferenceIDGETParams, authInfo runtime.ClientAuthInfoWriter) (*TransferReferenceIDGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTransferReferenceIDGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "transfer-referenceId-GET",
		Method:             "GET",
		PathPattern:        "/v1_0/transfer/{referenceId}",
		ProducesMediaTypes: []string{"API user insufficient balance", "Payer limit breached", "Successful transfer", "Transfer not found", "Unspecified internal error", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TransferReferenceIDGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TransferReferenceIDGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for transfer-referenceId-GET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
