package collection_go

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewRequesttopayReferenceIDGETParams creates a new RequesttopayReferenceIDGETParams object
// with the default values initialized.
func NewRequesttopayReferenceIDGETParams() *RequesttopayReferenceIDGETParams {
	var ()
	return &RequesttopayReferenceIDGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRequesttopayReferenceIDGETParamsWithTimeout creates a new RequesttopayReferenceIDGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRequesttopayReferenceIDGETParamsWithTimeout(timeout time.Duration) *RequesttopayReferenceIDGETParams {
	var ()
	return &RequesttopayReferenceIDGETParams{

		timeout: timeout,
	}
}

// NewRequesttopayReferenceIDGETParamsWithContext creates a new RequesttopayReferenceIDGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewRequesttopayReferenceIDGETParamsWithContext(ctx context.Context) *RequesttopayReferenceIDGETParams {
	var ()
	return &RequesttopayReferenceIDGETParams{

		Context: ctx,
	}
}

// NewRequesttopayReferenceIDGETParamsWithHTTPClient creates a new RequesttopayReferenceIDGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRequesttopayReferenceIDGETParamsWithHTTPClient(client *http.Client) *RequesttopayReferenceIDGETParams {
	var ()
	return &RequesttopayReferenceIDGETParams{
		HTTPClient: client,
	}
}

/*RequesttopayReferenceIDGETParams contains all the parameters to send to the API endpoint
for the requesttopay reference Id g e t operation typically these are written to a http.Request
*/
type RequesttopayReferenceIDGETParams struct {

	/*Authorization
	  Authorization header used for Basic authentication and oauth. Format of the header parameter follows the standard for Basic and Bearer. Oauth uses Bearer authentication type where the credential is the received access token.

	*/
	Authorization *string
	/*XTargetEnvironment
	  The identifier of the EWP system where the transaction shall be processed. This parameter is used to route the request to the EWP system that will initiate the transaction.

	*/
	XTargetEnvironment string
	/*ReferenceID
	  UUID of transaction to get result. Reference id  used when creating the request to pay.

	*/
	ReferenceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the requesttopay reference Id g e t params
func (o *RequesttopayReferenceIDGETParams) WithTimeout(timeout time.Duration) *RequesttopayReferenceIDGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the requesttopay reference Id g e t params
func (o *RequesttopayReferenceIDGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the requesttopay reference Id g e t params
func (o *RequesttopayReferenceIDGETParams) WithContext(ctx context.Context) *RequesttopayReferenceIDGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the requesttopay reference Id g e t params
func (o *RequesttopayReferenceIDGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the requesttopay reference Id g e t params
func (o *RequesttopayReferenceIDGETParams) WithHTTPClient(client *http.Client) *RequesttopayReferenceIDGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the requesttopay reference Id g e t params
func (o *RequesttopayReferenceIDGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the requesttopay reference Id g e t params
func (o *RequesttopayReferenceIDGETParams) WithAuthorization(authorization *string) *RequesttopayReferenceIDGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the requesttopay reference Id g e t params
func (o *RequesttopayReferenceIDGETParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXTargetEnvironment adds the xTargetEnvironment to the requesttopay reference Id g e t params
func (o *RequesttopayReferenceIDGETParams) WithXTargetEnvironment(xTargetEnvironment string) *RequesttopayReferenceIDGETParams {
	o.SetXTargetEnvironment(xTargetEnvironment)
	return o
}

// SetXTargetEnvironment adds the xTargetEnvironment to the requesttopay reference Id g e t params
func (o *RequesttopayReferenceIDGETParams) SetXTargetEnvironment(xTargetEnvironment string) {
	o.XTargetEnvironment = xTargetEnvironment
}

// WithReferenceID adds the referenceID to the requesttopay reference Id g e t params
func (o *RequesttopayReferenceIDGETParams) WithReferenceID(referenceID string) *RequesttopayReferenceIDGETParams {
	o.SetReferenceID(referenceID)
	return o
}

// SetReferenceID adds the referenceId to the requesttopay reference Id g e t params
func (o *RequesttopayReferenceIDGETParams) SetReferenceID(referenceID string) {
	o.ReferenceID = referenceID
}

// WriteToRequest writes these params to a swagger request
func (o *RequesttopayReferenceIDGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authorization != nil {

		// header param Authorization
		if err := r.SetHeaderParam("Authorization", *o.Authorization); err != nil {
			return err
		}

	}

	// header param X-Target-Environment
	if err := r.SetHeaderParam("X-Target-Environment", o.XTargetEnvironment); err != nil {
		return err
	}

	// path param referenceId
	if err := r.SetPathParam("referenceId", o.ReferenceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
