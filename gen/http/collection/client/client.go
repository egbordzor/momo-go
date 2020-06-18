// Code generated by goa v3.1.3, DO NOT EDIT.
//
// Collection client HTTP transport
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

// Client lists the Collection service endpoint HTTP clients.
type Client struct {
	// NewToken Doer is the HTTP client used to make requests to the NewToken
	// endpoint.
	NewTokenDoer goahttp.Doer

	// GetBalance Doer is the HTTP client used to make requests to the GetBalance
	// endpoint.
	GetBalanceDoer goahttp.Doer

	// RetrieveAccountHolder Doer is the HTTP client used to make requests to the
	// RetrieveAccountHolder endpoint.
	RetrieveAccountHolderDoer goahttp.Doer

	// PaymentRequest Doer is the HTTP client used to make requests to the
	// PaymentRequest endpoint.
	PaymentRequestDoer goahttp.Doer

	// PaymentStatus Doer is the HTTP client used to make requests to the
	// PaymentStatus endpoint.
	PaymentStatusDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the Collection service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		NewTokenDoer:              doer,
		GetBalanceDoer:            doer,
		RetrieveAccountHolderDoer: doer,
		PaymentRequestDoer:        doer,
		PaymentStatusDoer:         doer,
		RestoreResponseBody:       restoreBody,
		scheme:                    scheme,
		host:                      host,
		decoder:                   dec,
		encoder:                   enc,
	}
}

// NewToken returns an endpoint that makes HTTP requests to the Collection
// service NewToken server.
func (c *Client) NewToken() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeNewTokenRequest(c.encoder)
		decodeResponse = DecodeNewTokenResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildNewTokenRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.NewTokenDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("Collection", "NewToken", err)
		}
		return decodeResponse(resp)
	}
}

// GetBalance returns an endpoint that makes HTTP requests to the Collection
// service GetBalance server.
func (c *Client) GetBalance() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeGetBalanceRequest(c.encoder)
		decodeResponse = DecodeGetBalanceResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetBalanceRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetBalanceDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("Collection", "GetBalance", err)
		}
		return decodeResponse(resp)
	}
}

// RetrieveAccountHolder returns an endpoint that makes HTTP requests to the
// Collection service RetrieveAccountHolder server.
func (c *Client) RetrieveAccountHolder() endpoint.Endpoint {
	var (
		decodeResponse = DecodeRetrieveAccountHolderResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRetrieveAccountHolderRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RetrieveAccountHolderDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("Collection", "RetrieveAccountHolder", err)
		}
		return decodeResponse(resp)
	}
}

// PaymentRequest returns an endpoint that makes HTTP requests to the
// Collection service PaymentRequest server.
func (c *Client) PaymentRequest() endpoint.Endpoint {
	var (
		encodeRequest  = EncodePaymentRequestRequest(c.encoder)
		decodeResponse = DecodePaymentRequestResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPaymentRequestRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PaymentRequestDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("Collection", "PaymentRequest", err)
		}
		return decodeResponse(resp)
	}
}

// PaymentStatus returns an endpoint that makes HTTP requests to the Collection
// service PaymentStatus server.
func (c *Client) PaymentStatus() endpoint.Endpoint {
	var (
		decodeResponse = DecodePaymentStatusResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPaymentStatusRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PaymentStatusDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("Collection", "PaymentStatus", err)
		}
		return decodeResponse(resp)
	}
}