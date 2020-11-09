package disbursements_go

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

// NewTransferPOSTParams creates a new TransferPOSTParams object
// with the default values initialized.
func NewTransferPOSTParams() *TransferPOSTParams {
	var ()
	return &TransferPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTransferPOSTParamsWithTimeout creates a new TransferPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTransferPOSTParamsWithTimeout(timeout time.Duration) *TransferPOSTParams {
	var ()
	return &TransferPOSTParams{

		timeout: timeout,
	}
}

// NewTransferPOSTParamsWithContext creates a new TransferPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewTransferPOSTParamsWithContext(ctx context.Context) *TransferPOSTParams {
	var ()
	return &TransferPOSTParams{

		Context: ctx,
	}
}

// NewTransferPOSTParamsWithHTTPClient creates a new TransferPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTransferPOSTParamsWithHTTPClient(client *http.Client) *TransferPOSTParams {
	var ()
	return &TransferPOSTParams{
		HTTPClient: client,
	}
}

/*TransferPOSTParams contains all the parameters to send to the API endpoint
for the transfer p o s t operation typically these are written to a http.Request
*/
type TransferPOSTParams struct {

	/*Authorization
	  Authorization header used for Basic authentication and oauth. Format of the header parameter follows the standard for Basic and Bearer. Oauth uses Bearer authentication type where the credential is the received access token.

	*/
	Authorization *string
	/*XCallbackURL
	  URL to the server where the callback should be sent.

	*/
	XCallbackURL *string
	/*XReferenceID
	  Format - UUID. Recource ID of the created ‘request-to-pay’ transaction. This ID is used for e.g. validating the status of the request. Universal Unique ID for the transaction generated using UUID version 4.

	*/
	XReferenceID string
	/*XTargetEnvironment
	  The identifier of the EWP system where the transaction shall be processed. This parameter is used to route the request to the EWP system that will initiate the transaction.

	*/
	XTargetEnvironment string
	/*Transfer*/
	Transfer *models.Transfer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the transfer p o s t params
func (o *TransferPOSTParams) WithTimeout(timeout time.Duration) *TransferPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the transfer p o s t params
func (o *TransferPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the transfer p o s t params
func (o *TransferPOSTParams) WithContext(ctx context.Context) *TransferPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the transfer p o s t params
func (o *TransferPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the transfer p o s t params
func (o *TransferPOSTParams) WithHTTPClient(client *http.Client) *TransferPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the transfer p o s t params
func (o *TransferPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the transfer p o s t params
func (o *TransferPOSTParams) WithAuthorization(authorization *string) *TransferPOSTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the transfer p o s t params
func (o *TransferPOSTParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXCallbackURL adds the xCallbackURL to the transfer p o s t params
func (o *TransferPOSTParams) WithXCallbackURL(xCallbackURL *string) *TransferPOSTParams {
	o.SetXCallbackURL(xCallbackURL)
	return o
}

// SetXCallbackURL adds the xCallbackUrl to the transfer p o s t params
func (o *TransferPOSTParams) SetXCallbackURL(xCallbackURL *string) {
	o.XCallbackURL = xCallbackURL
}

// WithXReferenceID adds the xReferenceID to the transfer p o s t params
func (o *TransferPOSTParams) WithXReferenceID(xReferenceID string) *TransferPOSTParams {
	o.SetXReferenceID(xReferenceID)
	return o
}

// SetXReferenceID adds the xReferenceId to the transfer p o s t params
func (o *TransferPOSTParams) SetXReferenceID(xReferenceID string) {
	o.XReferenceID = xReferenceID
}

// WithXTargetEnvironment adds the xTargetEnvironment to the transfer p o s t params
func (o *TransferPOSTParams) WithXTargetEnvironment(xTargetEnvironment string) *TransferPOSTParams {
	o.SetXTargetEnvironment(xTargetEnvironment)
	return o
}

// SetXTargetEnvironment adds the xTargetEnvironment to the transfer p o s t params
func (o *TransferPOSTParams) SetXTargetEnvironment(xTargetEnvironment string) {
	o.XTargetEnvironment = xTargetEnvironment
}

// WithTransfer adds the transfer to the transfer p o s t params
func (o *TransferPOSTParams) WithTransfer(transfer *models.Transfer) *TransferPOSTParams {
	o.SetTransfer(transfer)
	return o
}

// SetTransfer adds the transfer to the transfer p o s t params
func (o *TransferPOSTParams) SetTransfer(transfer *models.Transfer) {
	o.Transfer = transfer
}

// WriteToRequest writes these params to a swagger request
func (o *TransferPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Transfer != nil {
		if err := r.SetBodyParam(o.Transfer); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
