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

// NewTransferReferenceIDGETParams creates a new TransferReferenceIDGETParams object
// with the default values initialized.
func NewTransferReferenceIDGETParams() *TransferReferenceIDGETParams {
	var ()
	return &TransferReferenceIDGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTransferReferenceIDGETParamsWithTimeout creates a new TransferReferenceIDGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTransferReferenceIDGETParamsWithTimeout(timeout time.Duration) *TransferReferenceIDGETParams {
	var ()
	return &TransferReferenceIDGETParams{

		timeout: timeout,
	}
}

// NewTransferReferenceIDGETParamsWithContext creates a new TransferReferenceIDGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewTransferReferenceIDGETParamsWithContext(ctx context.Context) *TransferReferenceIDGETParams {
	var ()
	return &TransferReferenceIDGETParams{

		Context: ctx,
	}
}

// NewTransferReferenceIDGETParamsWithHTTPClient creates a new TransferReferenceIDGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTransferReferenceIDGETParamsWithHTTPClient(client *http.Client) *TransferReferenceIDGETParams {
	var ()
	return &TransferReferenceIDGETParams{
		HTTPClient: client,
	}
}

/*TransferReferenceIDGETParams contains all the parameters to send to the API endpoint
for the transfer reference Id g e t operation typically these are written to a http.Request
*/
type TransferReferenceIDGETParams struct {

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

// WithTimeout adds the timeout to the transfer reference Id g e t params
func (o *TransferReferenceIDGETParams) WithTimeout(timeout time.Duration) *TransferReferenceIDGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the transfer reference Id g e t params
func (o *TransferReferenceIDGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the transfer reference Id g e t params
func (o *TransferReferenceIDGETParams) WithContext(ctx context.Context) *TransferReferenceIDGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the transfer reference Id g e t params
func (o *TransferReferenceIDGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the transfer reference Id g e t params
func (o *TransferReferenceIDGETParams) WithHTTPClient(client *http.Client) *TransferReferenceIDGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the transfer reference Id g e t params
func (o *TransferReferenceIDGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the transfer reference Id g e t params
func (o *TransferReferenceIDGETParams) WithAuthorization(authorization *string) *TransferReferenceIDGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the transfer reference Id g e t params
func (o *TransferReferenceIDGETParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXTargetEnvironment adds the xTargetEnvironment to the transfer reference Id g e t params
func (o *TransferReferenceIDGETParams) WithXTargetEnvironment(xTargetEnvironment string) *TransferReferenceIDGETParams {
	o.SetXTargetEnvironment(xTargetEnvironment)
	return o
}

// SetXTargetEnvironment adds the xTargetEnvironment to the transfer reference Id g e t params
func (o *TransferReferenceIDGETParams) SetXTargetEnvironment(xTargetEnvironment string) {
	o.XTargetEnvironment = xTargetEnvironment
}

// WithReferenceID adds the referenceID to the transfer reference Id g e t params
func (o *TransferReferenceIDGETParams) WithReferenceID(referenceID string) *TransferReferenceIDGETParams {
	o.SetReferenceID(referenceID)
	return o
}

// SetReferenceID adds the referenceId to the transfer reference Id g e t params
func (o *TransferReferenceIDGETParams) SetReferenceID(referenceID string) {
	o.ReferenceID = referenceID
}

// WriteToRequest writes these params to a swagger request
func (o *TransferReferenceIDGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
