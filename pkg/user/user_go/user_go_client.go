package user_go

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user go API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user go API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetApiuser(params *GetApiuserParams, authInfo runtime.ClientAuthInfoWriter) (*GetApiuserOK, error)

	PostApiuser(params *PostApiuserParams, authInfo runtime.ClientAuthInfoWriter) (*PostApiuserCreated, error)

	PostApiuserApikey(params *PostApiuserApikeyParams, authInfo runtime.ClientAuthInfoWriter) (*PostApiuserApikeyCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetApiuser v1s 0 apiuser x reference Id g e t

  Used to get API user information.
*/
func (a *Client) GetApiuser(params *GetApiuserParams, authInfo runtime.ClientAuthInfoWriter) (*GetApiuserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApiuserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-v1_0-apiuser",
		Method:             "GET",
		PathPattern:        "/v1_0/apiuser/{X-Reference-Id}",
		ProducesMediaTypes: []string{"Requested resource was not found"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetApiuserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApiuserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-v1_0-apiuser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostApiuser apiusers p o s t

  Used to create an API user in the sandbox target environment.
*/
func (a *Client) PostApiuser(params *PostApiuserParams, authInfo runtime.ClientAuthInfoWriter) (*PostApiuserCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostApiuserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "post-v1_0-apiuser",
		Method:             "POST",
		PathPattern:        "/v1_0/apiuser",
		ProducesMediaTypes: []string{"ReferenceId already in use", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostApiuserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostApiuserCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post-v1_0-apiuser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostApiuserApikey v1s 0 apiuser x reference Id apikey p o s t

  Used to create an API key for an API user in the sandbox target environment.
*/
func (a *Client) PostApiuserApikey(params *PostApiuserApikeyParams, authInfo runtime.ClientAuthInfoWriter) (*PostApiuserApikeyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostApiuserApikeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "post-v1_0-apiuser-apikey",
		Method:             "POST",
		PathPattern:        "/v1_0/apiuser/{X-Reference-Id}/apikey",
		ProducesMediaTypes: []string{"Requested resource was not found", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostApiuserApikeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostApiuserApikeyCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post-v1_0-apiuser-apikey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
