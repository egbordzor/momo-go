// Code generated by goa v3.1.3, DO NOT EDIT.
//
// User HTTP server types
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package server

import (
	user "github.com/wondenge/momo-go/gen/user"
	userviews "github.com/wondenge/momo-go/gen/user/views"
	goa "goa.design/goa/v3/pkg"
)

// NewUserRequestBody is the type of the "User" service "NewUser" endpoint HTTP
// request body.
type NewUserRequestBody struct {
	// The provider callback host
	ProviderCallbackHost *string `form:"providerCallbackHost,omitempty" json:"providerCallbackHost,omitempty" xml:"providerCallbackHost,omitempty"`
}

// NewKeyRequestBody is the type of the "User" service "NewKey" endpoint HTTP
// request body.
type NewKeyRequestBody struct {
	// Format - UUID.
	XReferenceID *string `form:"X-Reference-Id,omitempty" json:"X-Reference-Id,omitempty" xml:"X-Reference-Id,omitempty"`
}

// GetUserRequestBody is the type of the "User" service "GetUser" endpoint HTTP
// request body.
type GetUserRequestBody struct {
	// Format - UUID.
	XReferenceID *string `form:"X-Reference-Id,omitempty" json:"X-Reference-Id,omitempty" xml:"X-Reference-Id,omitempty"`
}

// NewKeyResponseBody is the type of the "User" service "NewKey" endpoint HTTP
// response body.
type NewKeyResponseBody struct {
	// The created API user key
	APIKey *string `form:"apiKey,omitempty" json:"apiKey,omitempty" xml:"apiKey,omitempty"`
}

// GetUserDetailsResponseBody is the type of the "User" service
// "GetUserDetails" endpoint HTTP response body.
type GetUserDetailsResponseBody struct {
	// The provider callback host
	ProviderCallbackHost *string                        `form:"providerCallbackHost,omitempty" json:"providerCallbackHost,omitempty" xml:"providerCallbackHost,omitempty"`
	PaymentServerURL     *PaymentServerURLResponseBody  `form:"paymentServerUrl,omitempty" json:"paymentServerUrl,omitempty" xml:"paymentServerUrl,omitempty"`
	TargetEnvironment    *TargetEnvironmentResponseBody `form:"targetEnvironment,omitempty" json:"targetEnvironment,omitempty" xml:"targetEnvironment,omitempty"`
}

// NewUserBadRequestResponseBody is the type of the "User" service "NewUser"
// endpoint HTTP response body for the "bad_request" error.
type NewUserBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewUserInternalErrorResponseBody is the type of the "User" service "NewUser"
// endpoint HTTP response body for the "internal_error" error.
type NewUserInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewUserConflictResponseBody is the type of the "User" service "NewUser"
// endpoint HTTP response body for the "conflict" error.
type NewUserConflictResponseBody struct {
	// Code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// NewKeyBadRequestResponseBody is the type of the "User" service "NewKey"
// endpoint HTTP response body for the "bad_request" error.
type NewKeyBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewKeyInternalErrorResponseBody is the type of the "User" service "NewKey"
// endpoint HTTP response body for the "internal_error" error.
type NewKeyInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewKeyNotFoundResponseBody is the type of the "User" service "NewKey"
// endpoint HTTP response body for the "not_found" error.
type NewKeyNotFoundResponseBody struct {
	// Code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// GetUserBadRequestResponseBody is the type of the "User" service "GetUser"
// endpoint HTTP response body for the "bad_request" error.
type GetUserBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetUserNotFoundResponseBody is the type of the "User" service "GetUser"
// endpoint HTTP response body for the "not_found" error.
type GetUserNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetUserInternalErrorResponseBody is the type of the "User" service "GetUser"
// endpoint HTTP response body for the "internal_error" error.
type GetUserInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetUserDetailsBadRequestResponseBody is the type of the "User" service
// "GetUserDetails" endpoint HTTP response body for the "bad_request" error.
type GetUserDetailsBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetUserDetailsInternalErrorResponseBody is the type of the "User" service
// "GetUserDetails" endpoint HTTP response body for the "internal_error" error.
type GetUserDetailsInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PaymentServerURLResponseBody is used to define fields on response body types.
type PaymentServerURLResponseBody struct {
	// The payment server URL
	APIKey *string `form:"apiKey,omitempty" json:"apiKey,omitempty" xml:"apiKey,omitempty"`
}

// TargetEnvironmentResponseBody is used to define fields on response body
// types.
type TargetEnvironmentResponseBody struct {
	// The target environment
	APIKey *string `form:"apiKey,omitempty" json:"apiKey,omitempty" xml:"apiKey,omitempty"`
}

// NewNewKeyResponseBody builds the HTTP response body from the result of the
// "NewKey" endpoint of the "User" service.
func NewNewKeyResponseBody(res *user.APIUserKeyResult) *NewKeyResponseBody {
	body := &NewKeyResponseBody{
		APIKey: res.APIKey,
	}
	return body
}

// NewGetUserDetailsResponseBody builds the HTTP response body from the result
// of the "GetUserDetails" endpoint of the "User" service.
func NewGetUserDetailsResponseBody(res *userviews.APIUserResultView) *GetUserDetailsResponseBody {
	body := &GetUserDetailsResponseBody{
		ProviderCallbackHost: res.ProviderCallbackHost,
	}
	if res.PaymentServerURL != nil {
		body.PaymentServerURL = marshalUserviewsPaymentServerURLViewToPaymentServerURLResponseBody(res.PaymentServerURL)
	}
	if res.TargetEnvironment != nil {
		body.TargetEnvironment = marshalUserviewsTargetEnvironmentViewToTargetEnvironmentResponseBody(res.TargetEnvironment)
	}
	return body
}

// NewNewUserBadRequestResponseBody builds the HTTP response body from the
// result of the "NewUser" endpoint of the "User" service.
func NewNewUserBadRequestResponseBody(res *goa.ServiceError) *NewUserBadRequestResponseBody {
	body := &NewUserBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewNewUserInternalErrorResponseBody builds the HTTP response body from the
// result of the "NewUser" endpoint of the "User" service.
func NewNewUserInternalErrorResponseBody(res *goa.ServiceError) *NewUserInternalErrorResponseBody {
	body := &NewUserInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewNewUserConflictResponseBody builds the HTTP response body from the result
// of the "NewUser" endpoint of the "User" service.
func NewNewUserConflictResponseBody(res *user.ErrorReason) *NewUserConflictResponseBody {
	body := &NewUserConflictResponseBody{
		Code:    res.Code,
		Message: res.Message,
	}
	return body
}

// NewNewKeyBadRequestResponseBody builds the HTTP response body from the
// result of the "NewKey" endpoint of the "User" service.
func NewNewKeyBadRequestResponseBody(res *goa.ServiceError) *NewKeyBadRequestResponseBody {
	body := &NewKeyBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewNewKeyInternalErrorResponseBody builds the HTTP response body from the
// result of the "NewKey" endpoint of the "User" service.
func NewNewKeyInternalErrorResponseBody(res *goa.ServiceError) *NewKeyInternalErrorResponseBody {
	body := &NewKeyInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewNewKeyNotFoundResponseBody builds the HTTP response body from the result
// of the "NewKey" endpoint of the "User" service.
func NewNewKeyNotFoundResponseBody(res *user.ErrorReason) *NewKeyNotFoundResponseBody {
	body := &NewKeyNotFoundResponseBody{
		Code:    res.Code,
		Message: res.Message,
	}
	return body
}

// NewGetUserBadRequestResponseBody builds the HTTP response body from the
// result of the "GetUser" endpoint of the "User" service.
func NewGetUserBadRequestResponseBody(res *goa.ServiceError) *GetUserBadRequestResponseBody {
	body := &GetUserBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetUserNotFoundResponseBody builds the HTTP response body from the result
// of the "GetUser" endpoint of the "User" service.
func NewGetUserNotFoundResponseBody(res *goa.ServiceError) *GetUserNotFoundResponseBody {
	body := &GetUserNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetUserInternalErrorResponseBody builds the HTTP response body from the
// result of the "GetUser" endpoint of the "User" service.
func NewGetUserInternalErrorResponseBody(res *goa.ServiceError) *GetUserInternalErrorResponseBody {
	body := &GetUserInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetUserDetailsBadRequestResponseBody builds the HTTP response body from
// the result of the "GetUserDetails" endpoint of the "User" service.
func NewGetUserDetailsBadRequestResponseBody(res *goa.ServiceError) *GetUserDetailsBadRequestResponseBody {
	body := &GetUserDetailsBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetUserDetailsInternalErrorResponseBody builds the HTTP response body
// from the result of the "GetUserDetails" endpoint of the "User" service.
func NewGetUserDetailsInternalErrorResponseBody(res *goa.ServiceError) *GetUserDetailsInternalErrorResponseBody {
	body := &GetUserDetailsInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewNewUserAPIUser builds a User service NewUser endpoint payload.
func NewNewUserAPIUser(body *NewUserRequestBody) *user.APIUser {
	v := &user.APIUser{
		ProviderCallbackHost: body.ProviderCallbackHost,
	}

	return v
}

// NewNewKeyPayload builds a User service NewKey endpoint payload.
func NewNewKeyPayload(body *NewKeyRequestBody) *user.NewKeyPayload {
	v := &user.NewKeyPayload{
		XReferenceID: *body.XReferenceID,
	}

	return v
}

// NewGetUserPayload builds a User service GetUser endpoint payload.
func NewGetUserPayload(body *GetUserRequestBody) *user.GetUserPayload {
	v := &user.GetUserPayload{
		XReferenceID: *body.XReferenceID,
	}

	return v
}

// ValidateNewKeyRequestBody runs the validations defined on NewKeyRequestBody
func ValidateNewKeyRequestBody(body *NewKeyRequestBody) (err error) {
	if body.XReferenceID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("X-Reference-Id", "body"))
	}
	return
}

// ValidateGetUserRequestBody runs the validations defined on GetUserRequestBody
func ValidateGetUserRequestBody(body *GetUserRequestBody) (err error) {
	if body.XReferenceID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("X-Reference-Id", "body"))
	}
	return
}