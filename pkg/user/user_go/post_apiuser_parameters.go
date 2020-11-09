package user_go

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/wondenge/momo-go/pkg/models"
)

// NewPostApiuserParams creates a new PostApiuserParams object
// with the default values initialized.
func NewPostApiuserParams() *PostApiuserParams {
	var ()
	return &PostApiuserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostApiuserParamsWithTimeout creates a new PostApiuserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostApiuserParamsWithTimeout(timeout time.Duration) *PostApiuserParams {
	var ()
	return &PostApiuserParams{

		timeout: timeout,
	}
}

// NewPostApiuserParamsWithContext creates a new PostApiuserParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostApiuserParamsWithContext(ctx context.Context) *PostApiuserParams {
	var ()
	return &PostApiuserParams{

		Context: ctx,
	}
}

// NewPostApiuserParamsWithHTTPClient creates a new PostApiuserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostApiuserParamsWithHTTPClient(client *http.Client) *PostApiuserParams {
	var ()
	return &PostApiuserParams{
		HTTPClient: client,
	}
}

/*PostApiuserParams contains all the parameters to send to the API endpoint
for the post v1 0 apiuser operation typically these are written to a http.Request
*/
type PostApiuserParams struct {

	/*XReferenceID
	  Format - UUID. Recource ID for the API user to be created. UUID version 4 is required.

	*/
	XReferenceID string
	/*APIUser*/
	APIUser *models.APIUser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post v1 0 apiuser params
func (o *PostApiuserParams) WithTimeout(timeout time.Duration) *PostApiuserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 0 apiuser params
func (o *PostApiuserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 0 apiuser params
func (o *PostApiuserParams) WithContext(ctx context.Context) *PostApiuserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 0 apiuser params
func (o *PostApiuserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 0 apiuser params
func (o *PostApiuserParams) WithHTTPClient(client *http.Client) *PostApiuserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 0 apiuser params
func (o *PostApiuserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXReferenceID adds the xReferenceID to the post v1 0 apiuser params
func (o *PostApiuserParams) WithXReferenceID(xReferenceID string) *PostApiuserParams {
	o.SetXReferenceID(xReferenceID)
	return o
}

// SetXReferenceID adds the xReferenceId to the post v1 0 apiuser params
func (o *PostApiuserParams) SetXReferenceID(xReferenceID string) {
	o.XReferenceID = xReferenceID
}

// WithAPIUser adds the aPIUser to the post v1 0 apiuser params
func (o *PostApiuserParams) WithAPIUser(aPIUser *models.APIUser) *PostApiuserParams {
	o.SetAPIUser(aPIUser)
	return o
}

// SetAPIUser adds the apiUser to the post v1 0 apiuser params
func (o *PostApiuserParams) SetAPIUser(aPIUser *models.APIUser) {
	o.APIUser = aPIUser
}

// WriteToRequest writes these params to a swagger request
func (o *PostApiuserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Reference-Id
	if err := r.SetHeaderParam("X-Reference-Id", o.XReferenceID); err != nil {
		return err
	}

	if o.APIUser != nil {
		if err := r.SetBodyParam(o.APIUser); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
