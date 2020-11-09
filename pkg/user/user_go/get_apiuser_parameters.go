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

// NewGetApiuserParams creates a new GetApiuserParams object
// with the default values initialized.
func NewGetApiuserParams() *GetApiuserParams {
	var ()
	return &GetApiuserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetApiuserParamsWithTimeout creates a new GetApiuserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetApiuserParamsWithTimeout(timeout time.Duration) *GetApiuserParams {
	var ()
	return &GetApiuserParams{

		timeout: timeout,
	}
}

// NewGetApiuserParamsWithContext creates a new GetApiuserParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetApiuserParamsWithContext(ctx context.Context) *GetApiuserParams {
	var ()
	return &GetApiuserParams{

		Context: ctx,
	}
}

// NewGetApiuserParamsWithHTTPClient creates a new GetApiuserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetApiuserParamsWithHTTPClient(client *http.Client) *GetApiuserParams {
	var ()
	return &GetApiuserParams{
		HTTPClient: client,
	}
}

/*GetApiuserParams contains all the parameters to send to the API endpoint
for the get v1 0 apiuser operation typically these are written to a http.Request
*/
type GetApiuserParams struct {

	/*XReferenceID
	  Format - UUID. Recource ID for the API user to be created. UUID version 4 is required.

	*/
	XReferenceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 0 apiuser params
func (o *GetApiuserParams) WithTimeout(timeout time.Duration) *GetApiuserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 0 apiuser params
func (o *GetApiuserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 0 apiuser params
func (o *GetApiuserParams) WithContext(ctx context.Context) *GetApiuserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 0 apiuser params
func (o *GetApiuserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 0 apiuser params
func (o *GetApiuserParams) WithHTTPClient(client *http.Client) *GetApiuserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 0 apiuser params
func (o *GetApiuserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXReferenceID adds the xReferenceID to the get v1 0 apiuser params
func (o *GetApiuserParams) WithXReferenceID(xReferenceID string) *GetApiuserParams {
	o.SetXReferenceID(xReferenceID)
	return o
}

// SetXReferenceID adds the xReferenceId to the get v1 0 apiuser params
func (o *GetApiuserParams) SetXReferenceID(xReferenceID string) {
	o.XReferenceID = xReferenceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetApiuserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
