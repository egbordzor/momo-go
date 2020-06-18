// Code generated by goa v3.1.3, DO NOT EDIT.
//
// User client HTTP transport
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package client

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	goahttp "goa.design/goa/v3/http"
)

// Client lists the User service endpoint HTTP clients.
type Client struct {
	// NewUser Doer is the HTTP client used to make requests to the NewUser
	// endpoint.
	NewUserDoer goahttp.Doer

	// NewKey Doer is the HTTP client used to make requests to the NewKey endpoint.
	NewKeyDoer goahttp.Doer

	// GetUser Doer is the HTTP client used to make requests to the GetUser
	// endpoint.
	GetUserDoer goahttp.Doer

	// GetUserDetails Doer is the HTTP client used to make requests to the
	// GetUserDetails endpoint.
	GetUserDetailsDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the User service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		NewUserDoer:         doer,
		NewKeyDoer:          doer,
		GetUserDoer:         doer,
		GetUserDetailsDoer:  doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// NewUser returns an endpoint that makes HTTP requests to the User service
// NewUser server.
func (c *Client) NewUser() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeNewUserRequest(c.encoder)
		decodeResponse = DecodeNewUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildNewUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.NewUserDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("User", "NewUser", err)
		}
		return decodeResponse(resp)
	}
}

// NewKey returns an endpoint that makes HTTP requests to the User service
// NewKey server.
func (c *Client) NewKey() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeNewKeyRequest(c.encoder)
		decodeResponse = DecodeNewKeyResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildNewKeyRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.NewKeyDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("User", "NewKey", err)
		}
		return decodeResponse(resp)
	}
}

// GetUser returns an endpoint that makes HTTP requests to the User service
// GetUser server.
func (c *Client) GetUser() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeGetUserRequest(c.encoder)
		decodeResponse = DecodeGetUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetUserDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("User", "GetUser", err)
		}
		return decodeResponse(resp)
	}
}

// GetUserDetails returns an endpoint that makes HTTP requests to the User
// service GetUserDetails server.
func (c *Client) GetUserDetails() endpoint.Endpoint {
	var (
		decodeResponse = DecodeGetUserDetailsResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetUserDetailsRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetUserDetailsDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("User", "GetUserDetails", err)
		}
		return decodeResponse(resp)
	}
}