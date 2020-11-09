package collection_go

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

// NewRequesttopayPOSTParams creates a new RequesttopayPOSTParams object
// with the default values initialized.
func NewRequesttopayPOSTParams() *RequesttopayPOSTParams {
	var ()
	return &RequesttopayPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRequesttopayPOSTParamsWithTimeout creates a new RequesttopayPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRequesttopayPOSTParamsWithTimeout(timeout time.Duration) *RequesttopayPOSTParams {
	var ()
	return &RequesttopayPOSTParams{

		timeout: timeout,
	}
}

// NewRequesttopayPOSTParamsWithContext creates a new RequesttopayPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewRequesttopayPOSTParamsWithContext(ctx context.Context) *RequesttopayPOSTParams {
	var ()
	return &RequesttopayPOSTParams{

		Context: ctx,
	}
}

// NewRequesttopayPOSTParamsWithHTTPClient creates a new RequesttopayPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRequesttopayPOSTParamsWithHTTPClient(client *http.Client) *RequesttopayPOSTParams {
	var ()
	return &RequesttopayPOSTParams{
		HTTPClient: client,
	}
}

/*RequesttopayPOSTParams contains all the parameters to send to the API endpoint
for the requesttopay p o s t operation typically these are written to a http.Request
*/
type RequesttopayPOSTParams struct {

	/*Authorization
	  Authorization header used for Basic authentication and oauth. Format of the header parameter follows the standard for Basic and Bearer. Oauth uses Bearer authentication type where the credential is the received access token.

	*/
	Authorization *string
	/*XCallbackURL
	  URL to the server where the callback should be sent.

	*/
	XCallbackURL *string
	/*XReferenceID
	  Format - UUID. Recource ID of the created request to pay transaction. This ID is used, for example, validating the status of the request. ‘Universal Unique ID’ for the transaction generated using UUID version 4.

	*/
	XReferenceID string
	/*XTargetEnvironment
	  The identifier of the EWP system where the transaction shall be processed. This parameter is used to route the request to the EWP system that will initiate the transaction.

	*/
	XTargetEnvironment string
	/*RequestToPay*/
	RequestToPay *models.RequestToPay

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the requesttopay p o s t params
func (o *RequesttopayPOSTParams) WithTimeout(timeout time.Duration) *RequesttopayPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the requesttopay p o s t params
func (o *RequesttopayPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the requesttopay p o s t params
func (o *RequesttopayPOSTParams) WithContext(ctx context.Context) *RequesttopayPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the requesttopay p o s t params
func (o *RequesttopayPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the requesttopay p o s t params
func (o *RequesttopayPOSTParams) WithHTTPClient(client *http.Client) *RequesttopayPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the requesttopay p o s t params
func (o *RequesttopayPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the requesttopay p o s t params
func (o *RequesttopayPOSTParams) WithAuthorization(authorization *string) *RequesttopayPOSTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the requesttopay p o s t params
func (o *RequesttopayPOSTParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXCallbackURL adds the xCallbackURL to the requesttopay p o s t params
func (o *RequesttopayPOSTParams) WithXCallbackURL(xCallbackURL *string) *RequesttopayPOSTParams {
	o.SetXCallbackURL(xCallbackURL)
	return o
}

// SetXCallbackURL adds the xCallbackUrl to the requesttopay p o s t params
func (o *RequesttopayPOSTParams) SetXCallbackURL(xCallbackURL *string) {
	o.XCallbackURL = xCallbackURL
}

// WithXReferenceID adds the xReferenceID to the requesttopay p o s t params
func (o *RequesttopayPOSTParams) WithXReferenceID(xReferenceID string) *RequesttopayPOSTParams {
	o.SetXReferenceID(xReferenceID)
	return o
}

// SetXReferenceID adds the xReferenceId to the requesttopay p o s t params
func (o *RequesttopayPOSTParams) SetXReferenceID(xReferenceID string) {
	o.XReferenceID = xReferenceID
}

// WithXTargetEnvironment adds the xTargetEnvironment to the requesttopay p o s t params
func (o *RequesttopayPOSTParams) WithXTargetEnvironment(xTargetEnvironment string) *RequesttopayPOSTParams {
	o.SetXTargetEnvironment(xTargetEnvironment)
	return o
}

// SetXTargetEnvironment adds the xTargetEnvironment to the requesttopay p o s t params
func (o *RequesttopayPOSTParams) SetXTargetEnvironment(xTargetEnvironment string) {
	o.XTargetEnvironment = xTargetEnvironment
}

// WithRequestToPay adds the requestToPay to the requesttopay p o s t params
func (o *RequesttopayPOSTParams) WithRequestToPay(requestToPay *models.RequestToPay) *RequesttopayPOSTParams {
	o.SetRequestToPay(requestToPay)
	return o
}

// SetRequestToPay adds the requestToPay to the requesttopay p o s t params
func (o *RequesttopayPOSTParams) SetRequestToPay(requestToPay *models.RequestToPay) {
	o.RequestToPay = requestToPay
}

// WriteToRequest writes these params to a swagger request
func (o *RequesttopayPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.XCallbackURL != nil {

		// header param X-Callback-Url
		if err := r.SetHeaderParam("X-Callback-Url", *o.XCallbackURL); err != nil {
			return err
		}

	}

	// header param X-Reference-Id
	if err := r.SetHeaderParam("X-Reference-Id", o.XReferenceID); err != nil {
		return err
	}

	// header param X-Target-Environment
	if err := r.SetHeaderParam("X-Target-Environment", o.XTargetEnvironment); err != nil {
		return err
	}

	if o.RequestToPay != nil {
		if err := r.SetBodyParam(o.RequestToPay); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
