// Code generated by goa v3.1.3, DO NOT EDIT.
//
// collection go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package server

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/wondenge/momo-go/gen/http/collection/server"
	goahttp "goa.design/goa/v3/http"
)

// EncodeNewTokenResponse returns a go-kit EncodeResponseFunc suitable for
// encoding collection NewToken responses.
func EncodeNewTokenResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeNewTokenResponse(encoder)
}

// DecodeNewTokenRequest returns a go-kit DecodeRequestFunc suitable for
// decoding collection NewToken requests.
func DecodeNewTokenRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeNewTokenRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeNewTokenError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the collection NewToken endpoint.
func EncodeNewTokenError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeNewTokenError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeGetBalanceResponse returns a go-kit EncodeResponseFunc suitable for
// encoding collection GetBalance responses.
func EncodeGetBalanceResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeGetBalanceResponse(encoder)
}

// DecodeGetBalanceRequest returns a go-kit DecodeRequestFunc suitable for
// decoding collection GetBalance requests.
func DecodeGetBalanceRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeGetBalanceRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeGetBalanceError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the collection GetBalance endpoint.
func EncodeGetBalanceError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeGetBalanceError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeRetrieveAccountHolderResponse returns a go-kit EncodeResponseFunc
// suitable for encoding collection RetrieveAccountHolder responses.
func EncodeRetrieveAccountHolderResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeRetrieveAccountHolderResponse(encoder)
}

// DecodeRetrieveAccountHolderRequest returns a go-kit DecodeRequestFunc
// suitable for decoding collection RetrieveAccountHolder requests.
func DecodeRetrieveAccountHolderRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeRetrieveAccountHolderRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeRetrieveAccountHolderError returns a go-kit EncodeResponseFunc
// suitable for encoding errors returned by the collection
// RetrieveAccountHolder endpoint.
func EncodeRetrieveAccountHolderError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeRetrieveAccountHolderError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodePaymentRequestResponse returns a go-kit EncodeResponseFunc suitable
// for encoding collection PaymentRequest responses.
func EncodePaymentRequestResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodePaymentRequestResponse(encoder)
}

// DecodePaymentRequestRequest returns a go-kit DecodeRequestFunc suitable for
// decoding collection PaymentRequest requests.
func DecodePaymentRequestRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodePaymentRequestRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodePaymentRequestError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the collection PaymentRequest endpoint.
func EncodePaymentRequestError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodePaymentRequestError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodePaymentStatusResponse returns a go-kit EncodeResponseFunc suitable for
// encoding collection PaymentStatus responses.
func EncodePaymentStatusResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodePaymentStatusResponse(encoder)
}

// DecodePaymentStatusRequest returns a go-kit DecodeRequestFunc suitable for
// decoding collection PaymentStatus requests.
func DecodePaymentStatusRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodePaymentStatusRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodePaymentStatusError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the collection PaymentStatus endpoint.
func EncodePaymentStatusError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodePaymentStatusError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}
