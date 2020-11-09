package remittance_go

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewTokenPOSTParams creates a new TokenPOSTParams object
// with the default values initialized.
func NewTokenPOSTParams() *TokenPOSTParams {
	var ()
	return &TokenPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTokenPOSTParamsWithTimeout creates a new TokenPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTokenPOSTParamsWithTimeout(timeout time.Duration) *TokenPOSTParams {
	var ()
	return &TokenPOSTParams{

		timeout: timeout,
	}
}

// NewTokenPOSTParamsWithContext creates a new TokenPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewTokenPOSTParamsWithContext(ctx context.Context) *TokenPOSTParams {
	var ()
	return &TokenPOSTParams{

		Context: ctx,
	}
}

// NewTokenPOSTParamsWithHTTPClient creates a new TokenPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTokenPOSTParamsWithHTTPClient(client *http.Client) *TokenPOSTParams {
	var ()
	return &TokenPOSTParams{
		HTTPClient: client,
	}
}

/*TokenPOSTParams contains all the parameters to send to the API endpoint
for the token p o s t operation typically these are written to a http.Request
*/
type TokenPOSTParams struct {

	/*Authorization
	  Basic authentication header containing API user ID and API key. Should be sent in as B64 encoded.

	*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the token p o s t params
func (o *TokenPOSTParams) WithTimeout(timeout time.Duration) *TokenPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the token p o s t params
func (o *TokenPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the token p o s t params
func (o *TokenPOSTParams) WithContext(ctx context.Context) *TokenPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the token p o s t params
func (o *TokenPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the token p o s t params
func (o *TokenPOSTParams) WithHTTPClient(client *http.Client) *TokenPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the token p o s t params
func (o *TokenPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the token p o s t params
func (o *TokenPOSTParams) WithAuthorization(authorization string) *TokenPOSTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the token p o s t params
func (o *TokenPOSTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *TokenPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
