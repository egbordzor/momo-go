// Code generated by goa v3.1.3, DO NOT EDIT.
//
// User go-kit HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package client

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/wondenge/momo-go/gen/http/User/client"
	goahttp "goa.design/goa/v3/http"
)

// EncodeNewUserRequest returns a go-kit EncodeRequestFunc suitable for
// encoding User NewUser requests.
func EncodeNewUserRequest(encoder func(*http.Request) goahttp.Encoder) kithttp.EncodeRequestFunc {
	enc := client.EncodeNewUserRequest(encoder)
	return func(_ context.Context, r *http.Request, v interface{}) error {
		return enc(r, v)
	}
}

// DecodeNewUserResponse returns a go-kit DecodeResponseFunc suitable for
// decoding User NewUser responses.
func DecodeNewUserResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeNewUserResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}

// EncodeNewKeyRequest returns a go-kit EncodeRequestFunc suitable for encoding
// User NewKey requests.
func EncodeNewKeyRequest(encoder func(*http.Request) goahttp.Encoder) kithttp.EncodeRequestFunc {
	enc := client.EncodeNewKeyRequest(encoder)
	return func(_ context.Context, r *http.Request, v interface{}) error {
		return enc(r, v)
	}
}

// DecodeNewKeyResponse returns a go-kit DecodeResponseFunc suitable for
// decoding User NewKey responses.
func DecodeNewKeyResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeNewKeyResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}

// EncodeGetUserRequest returns a go-kit EncodeRequestFunc suitable for
// encoding User GetUser requests.
func EncodeGetUserRequest(encoder func(*http.Request) goahttp.Encoder) kithttp.EncodeRequestFunc {
	enc := client.EncodeGetUserRequest(encoder)
	return func(_ context.Context, r *http.Request, v interface{}) error {
		return enc(r, v)
	}
}

// DecodeGetUserResponse returns a go-kit DecodeResponseFunc suitable for
// decoding User GetUser responses.
func DecodeGetUserResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeGetUserResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}

// DecodeGetUserDetailsResponse returns a go-kit DecodeResponseFunc suitable
// for decoding User GetUserDetails responses.
func DecodeGetUserDetailsResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeGetUserDetailsResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}