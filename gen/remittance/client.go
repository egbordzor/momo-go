// Code generated by goa v3.1.3, DO NOT EDIT.
//
// remittance client
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package remittance

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Client is the "remittance" service client.
type Client struct {
	NewTokenEndpoint              endpoint.Endpoint
	GetBalanceEndpoint            endpoint.Endpoint
	RetrieveAccountHolderEndpoint endpoint.Endpoint
	TransferEndpoint              endpoint.Endpoint
	TransferStatusEndpoint        endpoint.Endpoint
}

// NewClient initializes a "remittance" service client given the endpoints.
func NewClient(newToken, getBalance, retrieveAccountHolder, transfer, transferStatus endpoint.Endpoint) *Client {
	return &Client{
		NewTokenEndpoint:              newToken,
		GetBalanceEndpoint:            getBalance,
		RetrieveAccountHolderEndpoint: retrieveAccountHolder,
		TransferEndpoint:              transfer,
		TransferStatusEndpoint:        transferStatus,
	}
}

// NewToken calls the "NewToken" endpoint of the "remittance" service.
// NewToken may return the following errors:
//	- "unauthorized" (type *TokenPost401ApplicationJSONResponse): Unauthorized
//	- "internal_error" (type *goa.ServiceError): Error
//	- error: internal error
func (c *Client) NewToken(ctx context.Context, p string) (res *TokenPost200ApplicationJSONResponse, err error) {
	var ires interface{}
	ires, err = c.NewTokenEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*TokenPost200ApplicationJSONResponse), nil
}

// GetBalance calls the "GetBalance" endpoint of the "remittance" service.
// GetBalance may return the following errors:
//	- "bad_request" (type *goa.ServiceError): Bad request, e.g. invalid data was sent in the request.
//	- "internal_error" (type *goa.ServiceError): Internal error. The returned response contains details.
//	- error: internal error
func (c *Client) GetBalance(ctx context.Context, p string) (res *Balance, err error) {
	var ires interface{}
	ires, err = c.GetBalanceEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Balance), nil
}

// RetrieveAccountHolder calls the "RetrieveAccountHolder" endpoint of the
// "remittance" service.
// RetrieveAccountHolder may return the following errors:
//	- "bad_request" (type *goa.ServiceError): Bad request, e.g. invalid data was sent in the request.
//	- "internal_error" (type *goa.ServiceError): Internal error. The returned response contains details.
//	- error: internal error
func (c *Client) RetrieveAccountHolder(ctx context.Context, p *RetrieveAccountHolderPayload) (res string, err error) {
	var ires interface{}
	ires, err = c.RetrieveAccountHolderEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Transfer calls the "Transfer" endpoint of the "remittance" service.
// Transfer may return the following errors:
//	- "bad_request" (type *goa.ServiceError): Bad request, e.g. invalid data was sent in the request.
//	- "conflict" (type *goa.ServiceError): Conflict, duplicated reference id
//	- "internal_error" (type *goa.ServiceError): Internal Error.
//	- error: internal error
func (c *Client) Transfer(ctx context.Context, p *Transfer1) (res string, err error) {
	var ires interface{}
	ires, err = c.TransferEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// TransferStatus calls the "TransferStatus" endpoint of the "remittance"
// service.
// TransferStatus may return the following errors:
//	- "bad_request" (type *goa.ServiceError): Bad request, e.g. an incorrectly formatted reference id was provided.
//	- "not_found" (type *goa.ServiceError): Resource not found.
//	- "internal_error" (type *goa.ServiceError): Internal Error. Note that if the retreieved transfer has failed, it will not cause this status to be returned. This status is only returned if the GET request itself fails.
//	- error: internal error
func (c *Client) TransferStatus(ctx context.Context, p *TransferStatusPayload) (res *TransferResult, err error) {
	var ires interface{}
	ires, err = c.TransferStatusEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*TransferResult), nil
}
