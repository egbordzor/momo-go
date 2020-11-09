package user_go

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPostApiuserApikeyParams creates a new PostApiuserApikeyParams object
// with the default values initialized.
func NewPostApiuserApikeyParams() *PostApiuserApikeyParams {
	var ()
	return &PostApiuserApikeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostApiuserApikeyParamsWithTimeout creates a new PostApiuserApikeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostApiuserApikeyParamsWithTimeout(timeout time.Duration) *PostApiuserApikeyParams {
	var ()
	return &PostApiuserApikeyParams{

		timeout: timeout,
	}
}

// NewPostApiuserApikeyParamsWithContext creates a new PostApiuserApikeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostApiuserApikeyParamsWithContext(ctx context.Context) *PostApiuserApikeyParams {
	var ()
	return &PostApiuserApikeyParams{

		Context: ctx,
	}
}

// NewPostApiuserApikeyParamsWithHTTPClient creates a new PostApiuserApikeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostApiuserApikeyParamsWithHTTPClient(client *http.Client) *PostApiuserApikeyParams {
	var ()
	return &PostApiuserApikeyParams{
		HTTPClient: client,
	}
}

/*PostApiuserApikeyParams contains all the parameters to send to the API endpoint
for the post v1 0 apiuser apikey operation typically these are written to a http.Request
*/
type PostApiuserApikeyParams struct {

	/*XReferenceID
	  Format - UUID. Recource ID for the API user to be created. UUID version 4 is required.

	*/
	XReferenceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post v1 0 apiuser apikey params
func (o *PostApiuserApikeyParams) WithTimeout(timeout time.Duration) *PostApiuserApikeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 0 apiuser apikey params
func (o *PostApiuserApikeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 0 apiuser apikey params
func (o *PostApiuserApikeyParams) WithContext(ctx context.Context) *PostApiuserApikeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 0 apiuser apikey params
func (o *PostApiuserApikeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 0 apiuser apikey params
func (o *PostApiuserApikeyParams) WithHTTPClient(client *http.Client) *PostApiuserApikeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 0 apiuser apikey params
func (o *PostApiuserApikeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXReferenceID adds the xReferenceID to the post v1 0 apiuser apikey params
func (o *PostApiuserApikeyParams) WithXReferenceID(xReferenceID string) *PostApiuserApikeyParams {
	o.SetXReferenceID(xReferenceID)
	return o
}

// SetXReferenceID adds the xReferenceId to the post v1 0 apiuser apikey params
func (o *PostApiuserApikeyParams) SetXReferenceID(xReferenceID string) {
	o.XReferenceID = xReferenceID
}

// WriteToRequest writes these params to a swagger request
func (o *PostApiuserApikeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param X-Reference-Id
	if err := r.SetPathParam("X-Reference-Id", o.XReferenceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
