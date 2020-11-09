package collection_go

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new collection go API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for collection go API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAccountBalance(params *GetAccountBalanceParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountBalanceOK, error)

	GetAccountholderAccountholderidtypeAccountholderidActive(params *GetAccountholderAccountholderidtypeAccountholderidActiveParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountholderAccountholderidtypeAccountholderidActiveOK, error)

	RequesttopayPOST(params *RequesttopayPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*RequesttopayPOSTAccepted, error)

	RequesttopayReferenceIDGET(params *RequesttopayReferenceIDGETParams, authInfo runtime.ClientAuthInfoWriter) (*RequesttopayReferenceIDGETOK, error)

	TokenPOST(params *TokenPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*TokenPOSTOK, error)

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
  RequesttopayPOST requesttopays p o s t

  This operation is used to request a payment from a consumer (Payer). The payer will be asked to authorize the payment. The transaction will be executed once the payer has authorized the payment. The requesttopay will be in status PENDING until the transaction is authorized or declined by the payer or it is timed out by the system.
 Status of the transaction can be validated by using the GET /requesttopay/\<resourceId\>
*/
func (a *Client) RequesttopayPOST(params *RequesttopayPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*RequesttopayPOSTAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequesttopayPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "requesttopay-POST",
		Method:             "POST",
		PathPattern:        "/v1_0/requesttopay",
		ProducesMediaTypes: []string{"ReferenceId already in use", "Unspecified internal error", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RequesttopayPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RequesttopayPOSTAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for requesttopay-POST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RequesttopayReferenceIDGET requesttopays reference Id g e t

  This operation is used to get the status of a request to pay. X-Reference-Id that was passed in the post is used as reference to the request.
*/
func (a *Client) RequesttopayReferenceIDGET(params *RequesttopayReferenceIDGETParams, authInfo runtime.ClientAuthInfoWriter) (*RequesttopayReferenceIDGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequesttopayReferenceIDGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "requesttopay-referenceId-GET",
		Method:             "GET",
		PathPattern:        "/v1_0/requesttopay/{referenceId}",
		ProducesMediaTypes: []string{"Payer not found", "Request to pay not found", "Successful request to pay", "Unspecified internal error", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RequesttopayReferenceIDGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RequesttopayReferenceIDGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for requesttopay-referenceId-GET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
