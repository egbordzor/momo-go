// Code generated by goa v3.1.3, DO NOT EDIT.
//
// user HTTP client encoders and decoders
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

	user "github.com/wondenge/momo-go/gen/user"
	userviews "github.com/wondenge/momo-go/gen/user/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildNewUserRequest instantiates a HTTP request object with method and path
// set to call the "user" service "NewUser" endpoint
func (c *Client) BuildNewUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: NewUserUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "NewUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeNewUserRequest returns an encoder for requests sent to the user
// NewUser server.
func EncodeNewUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.APIUser)
		if !ok {
			return goahttp.ErrInvalidType("user", "NewUser", "*user.APIUser", v)
		}
		body := NewNewUserRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("user", "NewUser", err)
		}
		return nil
	}
}

// DecodeNewUserResponse returns a decoder for responses returned by the user
// NewUser endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeNewUserResponse may return the following errors:
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- "conflict" (type *user.ErrorReason): http.StatusConflict
//	- error: internal error
func DecodeNewUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
		case http.StatusCreated:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "NewUser", err)
			}
			return body, nil
		case http.StatusBadRequest:
			var (
				body NewUserBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "NewUser", err)
			}
			err = ValidateNewUserBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "NewUser", err)
			}
			return nil, NewNewUserBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body NewUserInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "NewUser", err)
			}
			err = ValidateNewUserInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "NewUser", err)
			}
			return nil, NewNewUserInternalError(&body)
		case http.StatusConflict:
			var (
				body NewUserConflictResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "NewUser", err)
			}
			err = ValidateNewUserConflictResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "NewUser", err)
			}
			return nil, NewNewUserConflict(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "NewUser", resp.StatusCode, string(body))
		}
	}
}

// BuildNewKeyRequest instantiates a HTTP request object with method and path
// set to call the "user" service "NewKey" endpoint
func (c *Client) BuildNewKeyRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: NewKeyUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "NewKey", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeNewKeyRequest returns an encoder for requests sent to the user NewKey
// server.
func EncodeNewKeyRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.NewKeyPayload)
		if !ok {
			return goahttp.ErrInvalidType("user", "NewKey", "*user.NewKeyPayload", v)
		}
		body := NewNewKeyRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("user", "NewKey", err)
		}
		return nil
	}
}

// DecodeNewKeyResponse returns a decoder for responses returned by the user
// NewKey endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeNewKeyResponse may return the following errors:
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- "not_found" (type *user.ErrorReason): http.StatusNotFound
//	- error: internal error
func DecodeNewKeyResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
		case http.StatusCreated:
			var (
				body NewKeyResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "NewKey", err)
			}
			res := NewNewKeyAPIUserKeyResultCreated(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body NewKeyBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "NewKey", err)
			}
			err = ValidateNewKeyBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "NewKey", err)
			}
			return nil, NewNewKeyBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body NewKeyInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "NewKey", err)
			}
			err = ValidateNewKeyInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "NewKey", err)
			}
			return nil, NewNewKeyInternalError(&body)
		case http.StatusNotFound:
			var (
				body NewKeyNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "NewKey", err)
			}
			err = ValidateNewKeyNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "NewKey", err)
			}
			return nil, NewNewKeyNotFound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "NewKey", resp.StatusCode, string(body))
		}
	}
}

// BuildGetUserRequest instantiates a HTTP request object with method and path
// set to call the "user" service "GetUser" endpoint
func (c *Client) BuildGetUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetUserUserPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "GetUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetUserRequest returns an encoder for requests sent to the user
// GetUser server.
func EncodeGetUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.GetUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("user", "GetUser", "*user.GetUserPayload", v)
		}
		body := NewGetUserRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("user", "GetUser", err)
		}
		return nil
	}
}

// DecodeGetUserResponse returns a decoder for responses returned by the user
// GetUser endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetUserResponse may return the following errors:
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "not_found" (type *goa.ServiceError): http.StatusNotFound
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeGetUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				return nil, goahttp.ErrDecodingError("user", "GetUser", err)
			}
			return body, nil
		case http.StatusBadRequest:
			var (
				body GetUserBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "GetUser", err)
			}
			err = ValidateGetUserBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "GetUser", err)
			}
			return nil, NewGetUserBadRequest(&body)
		case http.StatusNotFound:
			var (
				body GetUserNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "GetUser", err)
			}
			err = ValidateGetUserNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "GetUser", err)
			}
			return nil, NewGetUserNotFound(&body)
		case http.StatusInternalServerError:
			var (
				body GetUserInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "GetUser", err)
			}
			err = ValidateGetUserInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "GetUser", err)
			}
			return nil, NewGetUserInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "GetUser", resp.StatusCode, string(body))
		}
	}
}

// BuildGetUserDetailsRequest instantiates a HTTP request object with method
// and path set to call the "user" service "GetUserDetails" endpoint
func (c *Client) BuildGetUserDetailsRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		aPIUser string
	)
	{
		p, ok := v.(string)
		if !ok {
			return nil, goahttp.ErrInvalidType("user", "GetUserDetails", "string", v)
		}
		aPIUser = p
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetUserDetailsUserPath(aPIUser)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "GetUserDetails", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetUserDetailsResponse returns a decoder for responses returned by the
// user GetUserDetails endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeGetUserDetailsResponse may return the following errors:
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeGetUserDetailsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body GetUserDetailsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "GetUserDetails", err)
			}
			p := NewGetUserDetailsAPIUserResultOK(&body)
			view := "default"
			vres := &userviews.APIUserResult{Projected: p, View: view}
			if err = userviews.ValidateAPIUserResult(vres); err != nil {
				return nil, goahttp.ErrValidationError("user", "GetUserDetails", err)
			}
			res := user.NewAPIUserResult(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body GetUserDetailsBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "GetUserDetails", err)
			}
			err = ValidateGetUserDetailsBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "GetUserDetails", err)
			}
			return nil, NewGetUserDetailsBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body GetUserDetailsInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "GetUserDetails", err)
			}
			err = ValidateGetUserDetailsInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "GetUserDetails", err)
			}
			return nil, NewGetUserDetailsInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "GetUserDetails", resp.StatusCode, string(body))
		}
	}
}

// unmarshalPaymentServerURLResponseBodyToUserviewsPaymentServerURLView builds
// a value of type *userviews.PaymentServerURLView from a value of type
// *PaymentServerURLResponseBody.
func unmarshalPaymentServerURLResponseBodyToUserviewsPaymentServerURLView(v *PaymentServerURLResponseBody) *userviews.PaymentServerURLView {
	if v == nil {
		return nil
	}
	res := &userviews.PaymentServerURLView{
		APIKey: v.APIKey,
	}

	return res
}

// unmarshalTargetEnvironmentResponseBodyToUserviewsTargetEnvironmentView
// builds a value of type *userviews.TargetEnvironmentView from a value of type
// *TargetEnvironmentResponseBody.
func unmarshalTargetEnvironmentResponseBodyToUserviewsTargetEnvironmentView(v *TargetEnvironmentResponseBody) *userviews.TargetEnvironmentView {
	if v == nil {
		return nil
	}
	res := &userviews.TargetEnvironmentView{
		APIKey: v.APIKey,
	}

	return res
}
