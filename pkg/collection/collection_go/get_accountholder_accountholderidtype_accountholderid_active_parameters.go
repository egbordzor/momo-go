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

// NewGetAccountholderAccountholderidtypeAccountholderidActiveParams creates a new GetAccountholderAccountholderidtypeAccountholderidActiveParams object
// with the default values initialized.
func NewGetAccountholderAccountholderidtypeAccountholderidActiveParams() *GetAccountholderAccountholderidtypeAccountholderidActiveParams {
	var ()
	return &GetAccountholderAccountholderidtypeAccountholderidActiveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountholderAccountholderidtypeAccountholderidActiveParamsWithTimeout creates a new GetAccountholderAccountholderidtypeAccountholderidActiveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountholderAccountholderidtypeAccountholderidActiveParamsWithTimeout(timeout time.Duration) *GetAccountholderAccountholderidtypeAccountholderidActiveParams {
	var ()
	return &GetAccountholderAccountholderidtypeAccountholderidActiveParams{

		timeout: timeout,
	}
}

// NewGetAccountholderAccountholderidtypeAccountholderidActiveParamsWithContext creates a new GetAccountholderAccountholderidtypeAccountholderidActiveParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountholderAccountholderidtypeAccountholderidActiveParamsWithContext(ctx context.Context) *GetAccountholderAccountholderidtypeAccountholderidActiveParams {
	var ()
	return &GetAccountholderAccountholderidtypeAccountholderidActiveParams{

		Context: ctx,
	}
}

// NewGetAccountholderAccountholderidtypeAccountholderidActiveParamsWithHTTPClient creates a new GetAccountholderAccountholderidtypeAccountholderidActiveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountholderAccountholderidtypeAccountholderidActiveParamsWithHTTPClient(client *http.Client) *GetAccountholderAccountholderidtypeAccountholderidActiveParams {
	var ()
	return &GetAccountholderAccountholderidtypeAccountholderidActiveParams{
		HTTPClient: client,
	}
}

/*GetAccountholderAccountholderidtypeAccountholderidActiveParams contains all the parameters to send to the API endpoint
for the get v1 0 accountholder accountholderidtype accountholderid active operation typically these are written to a http.Request
*/
type GetAccountholderAccountholderidtypeAccountholderidActiveParams struct {

	/*Authorization
	  Authorization header used for Basic authentication and oauth. Format of the header parameter follows the standard for Basic and Bearer. Oauth uses Bearer authentication type where the credential is the received access token.

	*/
	Authorization *string
	/*XTargetEnvironment
	  The identifier of the EWP system where the transaction shall be processed. This parameter is used to route the request to the EWP system that will initiate the transaction.

	*/
	XTargetEnvironment string
	/*AccountHolderID
	  The party number. Validated according to the party ID type (case Sensitive). <br> msisdn - Mobile Number validated according to ITU-T E.164. Validated with IsMSISDN<br> email - Validated to be a valid e-mail format. Validated with IsEmail<br> party_code - UUID of the party. Validated with IsUuid

	*/
	AccountHolderID string
	/*AccountHolderIDType
	  Specifies the type of the party ID. Allowed values [msisdn, email, party_code].  <br> accountHolderId should explicitly be in small letters.

	*/
	AccountHolderIDType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 0 accountholder accountholderidtype accountholderid active params
func (o *GetAccountholderAccountholderidtypeAccountholderidActiveParams) WithTimeout(timeout time.Duration) *GetAccountholderAccountholderidtypeAccountholderidActiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 0 accountholder accountholderidtype accountholderid active params
func (o *GetAccountholderAccountholderidtypeAccountholderidActiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 0 accountholder accountholderidtype accountholderid active params
func (o *GetAccountholderAccountholderidtypeAccountholderidActiveParams) WithContext(ctx context.Context) *GetAccountholderAccountholderidtypeAccountholderidActiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 0 accountholder accountholderidtype accountholderid active params
func (o *GetAccountholderAccountholderidtypeAccountholderidActiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 0 accountholder accountholderidtype accountholderid active params
func (o *GetAccountholderAccountholderidtypeAccountholderidActiveParams) WithHTTPClient(client *http.Client) *GetAccountholderAccountholderidtypeAccountholderidActiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 0 accountholder accountholderidtype accountholderid active params
func (o *GetAccountholderAccountholderidtypeAccountholderidActiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get v1 0 accountholder accountholderidtype accountholderid active params
func (o *GetAccountholderAccountholderidtypeAccountholderidActiveParams) WithAuthorization(authorization *string) *GetAccountholderAccountholderidtypeAccountholderidActiveParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get v1 0 accountholder accountholderidtype accountholderid active params
func (o *GetAccountholderAccountholderidtypeAccountholderidActiveParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXTargetEnvironment adds the xTargetEnvironment to the get v1 0 accountholder accountholderidtype accountholderid active params
func (o *GetAccountholderAccountholderidtypeAccountholderidActiveParams) WithXTargetEnvironment(xTargetEnvironment string) *GetAccountholderAccountholderidtypeAccountholderidActiveParams {
	o.SetXTargetEnvironment(xTargetEnvironment)
	return o
}

// SetXTargetEnvironment adds the xTargetEnvironment to the get v1 0 accountholder accountholderidtype accountholderid active params
func (o *GetAccountholderAccountholderidtypeAccountholderidActiveParams) SetXTargetEnvironment(xTargetEnvironment string) {
	o.XTargetEnvironment = xTargetEnvironment
}

// WithAccountHolderID adds the accountHolderID to the get v1 0 accountholder accountholderidtype accountholderid active params
func (o *GetAccountholderAccountholderidtypeAccountholderidActiveParams) WithAccountHolderID(accountHolderID string) *GetAccountholderAccountholderidtypeAccountholderidActiveParams {
	o.SetAccountHolderID(accountHolderID)
	return o
}

// SetAccountHolderID adds the accountHolderId to the get v1 0 accountholder accountholderidtype accountholderid active params
func (o *GetAccountholderAccountholderidtypeAccountholderidActiveParams) SetAccountHolderID(accountHolderID string) {
	o.AccountHolderID = accountHolderID
}

// WithAccountHolderIDType adds the accountHolderIDType to the get v1 0 accountholder accountholderidtype accountholderid active params
func (o *GetAccountholderAccountholderidtypeAccountholderidActiveParams) WithAccountHolderIDType(accountHolderIDType string) *GetAccountholderAccountholderidtypeAccountholderidActiveParams {
	o.SetAccountHolderIDType(accountHolderIDType)
	return o
}

// SetAccountHolderIDType adds the accountHolderIdType to the get v1 0 accountholder accountholderidtype accountholderid active params
func (o *GetAccountholderAccountholderidtypeAccountholderidActiveParams) SetAccountHolderIDType(accountHolderIDType string) {
	o.AccountHolderIDType = accountHolderIDType
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountholderAccountholderidtypeAccountholderidActiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param accountHolderId
	if err := r.SetPathParam("accountHolderId", o.AccountHolderID); err != nil {
		return err
	}

	// path param accountHolderIdType
	if err := r.SetPathParam("accountHolderIdType", o.AccountHolderIDType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
