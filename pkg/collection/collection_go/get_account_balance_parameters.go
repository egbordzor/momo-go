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

// NewGetAccountBalanceParams creates a new GetAccountBalanceParams object
// with the default values initialized.
func NewGetAccountBalanceParams() *GetAccountBalanceParams {
	var ()
	return &GetAccountBalanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountBalanceParamsWithTimeout creates a new GetAccountBalanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountBalanceParamsWithTimeout(timeout time.Duration) *GetAccountBalanceParams {
	var ()
	return &GetAccountBalanceParams{

		timeout: timeout,
	}
}

// NewGetAccountBalanceParamsWithContext creates a new GetAccountBalanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountBalanceParamsWithContext(ctx context.Context) *GetAccountBalanceParams {
	var ()
	return &GetAccountBalanceParams{

		Context: ctx,
	}
}

// NewGetAccountBalanceParamsWithHTTPClient creates a new GetAccountBalanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountBalanceParamsWithHTTPClient(client *http.Client) *GetAccountBalanceParams {
	var ()
	return &GetAccountBalanceParams{
		HTTPClient: client,
	}
}

/*GetAccountBalanceParams contains all the parameters to send to the API endpoint
for the get v1 0 account balance operation typically these are written to a http.Request
*/
type GetAccountBalanceParams struct {

	/*Authorization
	  Authorization header used for Basic authentication and oauth. Format of the header parameter follows the standard for Basic and Bearer. Oauth uses Bearer authentication type where the credential is the received access token.

	*/
	Authorization *string
	/*XTargetEnvironment
	  The identifier of the EWP system where the transaction shall be processed. This parameter is used to route the request to the EWP system that will initiate the transaction.

	*/
	XTargetEnvironment string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 0 account balance params
func (o *GetAccountBalanceParams) WithTimeout(timeout time.Duration) *GetAccountBalanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 0 account balance params
func (o *GetAccountBalanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 0 account balance params
func (o *GetAccountBalanceParams) WithContext(ctx context.Context) *GetAccountBalanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 0 account balance params
func (o *GetAccountBalanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 0 account balance params
func (o *GetAccountBalanceParams) WithHTTPClient(client *http.Client) *GetAccountBalanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 0 account balance params
func (o *GetAccountBalanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get v1 0 account balance params
func (o *GetAccountBalanceParams) WithAuthorization(authorization *string) *GetAccountBalanceParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get v1 0 account balance params
func (o *GetAccountBalanceParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXTargetEnvironment adds the xTargetEnvironment to the get v1 0 account balance params
func (o *GetAccountBalanceParams) WithXTargetEnvironment(xTargetEnvironment string) *GetAccountBalanceParams {
	o.SetXTargetEnvironment(xTargetEnvironment)
	return o
}

// SetXTargetEnvironment adds the xTargetEnvironment to the get v1 0 account balance params
func (o *GetAccountBalanceParams) SetXTargetEnvironment(xTargetEnvironment string) {
	o.XTargetEnvironment = xTargetEnvironment
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountBalanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
