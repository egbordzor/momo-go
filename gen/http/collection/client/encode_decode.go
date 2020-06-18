// Code generated by goa v3.1.3, DO NOT EDIT.
//
// collection HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	collection "github.com/wondenge/momo-go/gen/collection"
	collectionviews "github.com/wondenge/momo-go/gen/collection/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildNewTokenRequest instantiates a HTTP request object with method and path
// set to call the "collection" service "NewToken" endpoint
func (c *Client) BuildNewTokenRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: NewTokenCollectionPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("collection", "NewToken", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeNewTokenRequest returns an encoder for requests sent to the collection
// NewToken server.
func EncodeNewTokenRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(string)
		if !ok {
			return goahttp.ErrInvalidType("collection", "NewToken", "string", v)
		}
		body := p
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("collection", "NewToken", err)
		}
		return nil
	}
}

// DecodeNewTokenResponse returns a decoder for responses returned by the
// collection NewToken endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeNewTokenResponse may return the following errors:
//	- "unauthorized" (type *collection.TokenPost401ApplicationJSONResponse): http.StatusUnauthorized
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeNewTokenResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body NewTokenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "NewToken", err)
			}
			p := NewNewTokenTokenPost200ApplicationJSONResponseOK(&body)
			view := "default"
			vres := &collectionviews.TokenPost200ApplicationJSONResponse{Projected: p, View: view}
			if err = collectionviews.ValidateTokenPost200ApplicationJSONResponse(vres); err != nil {
				return nil, goahttp.ErrValidationError("collection", "NewToken", err)
			}
			res := collection.NewTokenPost200ApplicationJSONResponse(vres)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body NewTokenUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "NewToken", err)
			}
			return nil, NewNewTokenUnauthorized(&body)
		case http.StatusInternalServerError:
			var (
				body NewTokenInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "NewToken", err)
			}
			err = ValidateNewTokenInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("collection", "NewToken", err)
			}
			return nil, NewNewTokenInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("collection", "NewToken", resp.StatusCode, string(body))
		}
	}
}

// BuildGetBalanceRequest instantiates a HTTP request object with method and
// path set to call the "collection" service "GetBalance" endpoint
func (c *Client) BuildGetBalanceRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetBalanceCollectionPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("collection", "GetBalance", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetBalanceRequest returns an encoder for requests sent to the
// collection GetBalance server.
func EncodeGetBalanceRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(string)
		if !ok {
			return goahttp.ErrInvalidType("collection", "GetBalance", "string", v)
		}
		body := p
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("collection", "GetBalance", err)
		}
		return nil
	}
}

// DecodeGetBalanceResponse returns a decoder for responses returned by the
// collection GetBalance endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeGetBalanceResponse may return the following errors:
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeGetBalanceResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetBalanceResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "GetBalance", err)
			}
			p := NewGetBalanceBalanceOK(&body)
			view := "default"
			vres := &collectionviews.Balance{Projected: p, View: view}
			if err = collectionviews.ValidateBalance(vres); err != nil {
				return nil, goahttp.ErrValidationError("collection", "GetBalance", err)
			}
			res := collection.NewBalance(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body GetBalanceBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "GetBalance", err)
			}
			err = ValidateGetBalanceBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("collection", "GetBalance", err)
			}
			return nil, NewGetBalanceBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body GetBalanceInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "GetBalance", err)
			}
			err = ValidateGetBalanceInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("collection", "GetBalance", err)
			}
			return nil, NewGetBalanceInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("collection", "GetBalance", resp.StatusCode, string(body))
		}
	}
}

// BuildRetrieveAccountHolderRequest instantiates a HTTP request object with
// method and path set to call the "collection" service "RetrieveAccountHolder"
// endpoint
func (c *Client) BuildRetrieveAccountHolderRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		accountHolderIDType string
		accountHolderID     string
	)
	{
		p, ok := v.(*collection.RetrieveAccountHolderPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("collection", "RetrieveAccountHolder", "*collection.RetrieveAccountHolderPayload", v)
		}
		accountHolderIDType = p.AccountHolderIDType
		accountHolderID = p.AccountHolderID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RetrieveAccountHolderCollectionPath(accountHolderIDType, accountHolderID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("collection", "RetrieveAccountHolder", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeRetrieveAccountHolderResponse returns a decoder for responses returned
// by the collection RetrieveAccountHolder endpoint. restoreBody controls
// whether the response body should be restored after having been read.
// DecodeRetrieveAccountHolderResponse may return the following errors:
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeRetrieveAccountHolderResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "RetrieveAccountHolder", err)
			}
			return body, nil
		case http.StatusBadRequest:
			var (
				body RetrieveAccountHolderBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "RetrieveAccountHolder", err)
			}
			err = ValidateRetrieveAccountHolderBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("collection", "RetrieveAccountHolder", err)
			}
			return nil, NewRetrieveAccountHolderBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body RetrieveAccountHolderInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "RetrieveAccountHolder", err)
			}
			err = ValidateRetrieveAccountHolderInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("collection", "RetrieveAccountHolder", err)
			}
			return nil, NewRetrieveAccountHolderInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("collection", "RetrieveAccountHolder", resp.StatusCode, string(body))
		}
	}
}

// BuildPaymentRequestRequest instantiates a HTTP request object with method
// and path set to call the "collection" service "PaymentRequest" endpoint
func (c *Client) BuildPaymentRequestRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: PaymentRequestCollectionPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("collection", "PaymentRequest", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodePaymentRequestRequest returns an encoder for requests sent to the
// collection PaymentRequest server.
func EncodePaymentRequestRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*collection.RequestToPay)
		if !ok {
			return goahttp.ErrInvalidType("collection", "PaymentRequest", "*collection.RequestToPay", v)
		}
		body := NewPaymentRequestRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("collection", "PaymentRequest", err)
		}
		return nil
	}
}

// DecodePaymentRequestResponse returns a decoder for responses returned by the
// collection PaymentRequest endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodePaymentRequestResponse may return the following errors:
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "conflict" (type *goa.ServiceError): http.StatusConflict
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodePaymentRequestResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusAccepted:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "PaymentRequest", err)
			}
			return body, nil
		case http.StatusBadRequest:
			var (
				body PaymentRequestBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "PaymentRequest", err)
			}
			err = ValidatePaymentRequestBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("collection", "PaymentRequest", err)
			}
			return nil, NewPaymentRequestBadRequest(&body)
		case http.StatusConflict:
			var (
				body PaymentRequestConflictResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "PaymentRequest", err)
			}
			err = ValidatePaymentRequestConflictResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("collection", "PaymentRequest", err)
			}
			return nil, NewPaymentRequestConflict(&body)
		case http.StatusInternalServerError:
			var (
				body PaymentRequestInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "PaymentRequest", err)
			}
			err = ValidatePaymentRequestInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("collection", "PaymentRequest", err)
			}
			return nil, NewPaymentRequestInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("collection", "PaymentRequest", resp.StatusCode, string(body))
		}
	}
}

// BuildPaymentStatusRequest instantiates a HTTP request object with method and
// path set to call the "collection" service "PaymentStatus" endpoint
func (c *Client) BuildPaymentStatusRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		referenceID string
	)
	{
		p, ok := v.(*collection.PaymentStatusPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("collection", "PaymentStatus", "*collection.PaymentStatusPayload", v)
		}
		referenceID = p.ReferenceID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: PaymentStatusCollectionPath(referenceID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("collection", "PaymentStatus", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodePaymentStatusResponse returns a decoder for responses returned by the
// collection PaymentStatus endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodePaymentStatusResponse may return the following errors:
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "not_found" (type *goa.ServiceError): http.StatusNotFound
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodePaymentStatusResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body PaymentStatusResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "PaymentStatus", err)
			}
			p := NewPaymentStatusRequestToPayResultOK(&body)
			view := "default"
			vres := &collectionviews.RequestToPayResult{Projected: p, View: view}
			if err = collectionviews.ValidateRequestToPayResult(vres); err != nil {
				return nil, goahttp.ErrValidationError("collection", "PaymentStatus", err)
			}
			res := collection.NewRequestToPayResult(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body PaymentStatusBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "PaymentStatus", err)
			}
			err = ValidatePaymentStatusBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("collection", "PaymentStatus", err)
			}
			return nil, NewPaymentStatusBadRequest(&body)
		case http.StatusNotFound:
			var (
				body PaymentStatusNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "PaymentStatus", err)
			}
			err = ValidatePaymentStatusNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("collection", "PaymentStatus", err)
			}
			return nil, NewPaymentStatusNotFound(&body)
		case http.StatusInternalServerError:
			var (
				body PaymentStatusInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("collection", "PaymentStatus", err)
			}
			err = ValidatePaymentStatusInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("collection", "PaymentStatus", err)
			}
			return nil, NewPaymentStatusInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("collection", "PaymentStatus", resp.StatusCode, string(body))
		}
	}
}

// marshalCollectionPartyToPartyRequestBody builds a value of type
// *PartyRequestBody from a value of type *collection.Party.
func marshalCollectionPartyToPartyRequestBody(v *collection.Party) *PartyRequestBody {
	if v == nil {
		return nil
	}
	res := &PartyRequestBody{
		PartyIDType: v.PartyIDType,
		PartyID:     v.PartyID,
	}

	return res
}

// marshalPartyRequestBodyToCollectionParty builds a value of type
// *collection.Party from a value of type *PartyRequestBody.
func marshalPartyRequestBodyToCollectionParty(v *PartyRequestBody) *collection.Party {
	if v == nil {
		return nil
	}
	res := &collection.Party{
		PartyIDType: v.PartyIDType,
		PartyID:     v.PartyID,
	}

	return res
}

// unmarshalPartyResponseBodyToCollectionviewsPartyView builds a value of type
// *collectionviews.PartyView from a value of type *PartyResponseBody.
func unmarshalPartyResponseBodyToCollectionviewsPartyView(v *PartyResponseBody) *collectionviews.PartyView {
	if v == nil {
		return nil
	}
	res := &collectionviews.PartyView{
		PartyIDType: v.PartyIDType,
		PartyID:     v.PartyID,
	}

	return res
}

// unmarshalErrorReasonResponseBodyToCollectionviewsErrorReasonView builds a
// value of type *collectionviews.ErrorReasonView from a value of type
// *ErrorReasonResponseBody.
func unmarshalErrorReasonResponseBodyToCollectionviewsErrorReasonView(v *ErrorReasonResponseBody) *collectionviews.ErrorReasonView {
	if v == nil {
		return nil
	}
	res := &collectionviews.ErrorReasonView{
		Code:    v.Code,
		Message: v.Message,
	}

	return res
}
