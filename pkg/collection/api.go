/*
 * Collection
 *
 * Partner Gateway API document
 *
 * API version: 1.0
 *
 */

package collection

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// MoMoCollection DefaultApi service
type MoMoCollection service

// GetAccountBalanceOpts Optional parameters for the method 'GetAccountBalance'
type GetAccountBalanceOpts struct {
	Authorization optional.String
}

/*
GetAccountBalance /v1_0/account/balance - GET
Get the balance of the account.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xTargetEnvironment The identifier of the EWP system where the transaction shall be processed. This parameter is used to route the request to the EWP system that will initiate the transaction.
 * @param optional nil or *GetAccountBalanceOpts - Optional Parameters:
 * @param "Authorization" (optional.String) -  Authorization header used for Basic authentication and oauth. Format of the header parameter follows the standard for Basic and Bearer. Oauth uses Bearer authentication type where the credential is the received access token.
@return Balance
*/
func (a *MoMoCollection) GetAccountBalance(ctx _context.Context, xTargetEnvironment string, opts *GetAccountBalanceOpts) (Balance, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  Balance
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/v1_0/account/balance"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json", "Incorrect target environment"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	if opts != nil && opts.Authorization.IsSet() {
		headerParams["Authorization"] = parameterToString(opts.Authorization.Value(), "")
	}
	headerParams["X-Target-Environment"] = parameterToString(xTargetEnvironment, "")
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			headerParams["Ocp-Apim-Subscription-Key"] = key
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			queryParams.Add("subscription-key", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, path, httpMethod, postBody, headerParams, queryParams, formParams, formFileName, fileName, fileBytes)
	if err != nil {
		return returnValue, nil, err
	}

	res, err := a.client.callAPI(r)
	if err != nil || res == nil {
		return returnValue, res, err
	}

	body, err := _ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return returnValue, res, err
	}

	if res.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  body,
			error: res.Status,
		}
		if res.StatusCode == 500 {
			var v ErrorReason
			err = a.client.decode(&v, body, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, res, newErr
			}
			newErr.model = v
		}
		return returnValue, res, newErr
	}

	err = a.client.decode(&returnValue, body, res.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  body,
			error: err.Error(),
		}
		return returnValue, res, newErr
	}

	return returnValue, res, nil
}

// GetAccountholderAccountholderidtypeAccountholderidActiveOpts Optional parameters for the method 'GetAccountholderAccountholderidtypeAccountholderidActive'
type GetAccountholderAccountholderidtypeAccountholderidActiveOpts struct {
	Authorization optional.String
}

/*
GetAccountholderAccountholderidtypeAccountholderidActive /v1_0/accountholder/{accountHolderIdType}/{accountHolderId}/active - GET
Operation is used  to check if an account holder is registered and active in the system.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountHolderId The party number. Validated according to the party ID type (case Sensitive). <br> msisdn - Mobile Number validated according to ITU-T E.164. Validated with IsMSISDN<br> email - Validated to be a valid e-mail format. Validated with IsEmail<br> party_code - UUID of the party. Validated with IsUuid
 * @param accountHolderIdType Specifies the type of the party ID. Allowed values [msisdn, email, party_code].  <br> accountHolderId should explicitly be in small letters.
 * @param xTargetEnvironment The identifier of the EWP system where the transaction shall be processed. This parameter is used to route the request to the EWP system that will initiate the transaction.
 * @param optional nil or *GetAccountholderAccountholderidtypeAccountholderidActiveOpts - Optional Parameters:
 * @param "Authorization" (optional.String) -  Authorization header used for Basic authentication and oauth. Format of the header parameter follows the standard for Basic and Bearer. Oauth uses Bearer authentication type where the credential is the received access token.
*/
func (a *MoMoCollection) GetAccountholderAccountholderidtypeAccountholderidActive(ctx _context.Context, accountHolderId string, accountHolderIdType string, xTargetEnvironment string, opts *GetAccountholderAccountholderidtypeAccountholderidActiveOpts) (*_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/v1_0/accountholder/{accountHolderIdType}/{accountHolderId}/active"
	path = strings.Replace(path, "{"+"accountHolderId"+"}", _neturl.QueryEscape(parameterToString(accountHolderId, "")), -1)

	path = strings.Replace(path, "{"+"accountHolderIdType"+"}", _neturl.QueryEscape(parameterToString(accountHolderIdType, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"Incorrect target environment"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	if opts != nil && opts.Authorization.IsSet() {
		headerParams["Authorization"] = parameterToString(opts.Authorization.Value(), "")
	}
	headerParams["X-Target-Environment"] = parameterToString(xTargetEnvironment, "")
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			headerParams["Ocp-Apim-Subscription-Key"] = key
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			queryParams.Add("subscription-key", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, path, httpMethod, postBody, headerParams, queryParams, formParams, formFileName, fileName, fileBytes)
	if err != nil {
		return nil, err
	}

	res, err := a.client.callAPI(r)
	if err != nil || res == nil {
		return res, err
	}

	body, err := _ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return res, err
	}

	if res.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  body,
			error: res.Status,
		}
		if res.StatusCode == 500 {
			var v ErrorReason
			err = a.client.decode(&v, body, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return res, newErr
			}
			newErr.model = v
		}
		return res, newErr
	}

	return res, nil
}

// RequesttopayPOSTOpts Optional parameters for the method 'RequesttopayPOST'
type RequesttopayPOSTOpts struct {
	Authorization optional.String
	XCallbackUrl  optional.String
	RequestToPay  optional.Interface
}

/*
RequesttopayPOST /requesttopay - POST
This operation is used to request a payment from a consumer (Payer). The payer will be asked to authorize the payment. The transaction will be executed once the payer has authorized the payment. The requesttopay will be in status PENDING until the transaction is authorized or declined by the payer or it is timed out by the system.   Status of the transaction can be validated by using the GET /requesttopay/\\&lt;resourceId\\&gt;
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xReferenceId Format - UUID. Recource ID of the created request to pay transaction. This ID is used, for example, validating the status of the request. ‘Universal Unique ID’ for the transaction generated using UUID version 4.
 * @param xTargetEnvironment The identifier of the EWP system where the transaction shall be processed. This parameter is used to route the request to the EWP system that will initiate the transaction.
 * @param optional nil or *RequesttopayPOSTOpts - Optional Parameters:
 * @param "Authorization" (optional.String) -  Authorization header used for Basic authentication and oauth. Format of the header parameter follows the standard for Basic and Bearer. Oauth uses Bearer authentication type where the credential is the received access token.
 * @param "XCallbackUrl" (optional.String) -  URL to the server where the callback should be sent.
 * @param "RequestToPay" (optional.Interface of RequestToPay) -
*/
func (a *MoMoCollection) RequesttopayPOST(ctx _context.Context, xReferenceId string, xTargetEnvironment string, opts *RequesttopayPOSTOpts) (*_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/v1_0/requesttopay"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json", "ReferenceId already in use", "Unspecified internal error"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	if opts != nil && opts.Authorization.IsSet() {
		headerParams["Authorization"] = parameterToString(opts.Authorization.Value(), "")
	}
	if opts != nil && opts.XCallbackUrl.IsSet() {
		headerParams["X-Callback-Url"] = parameterToString(opts.XCallbackUrl.Value(), "")
	}
	headerParams["X-Reference-Id"] = parameterToString(xReferenceId, "")
	headerParams["X-Target-Environment"] = parameterToString(xTargetEnvironment, "")
	// body params
	if opts != nil && opts.RequestToPay.IsSet() {
		optRequestToPay, optRequestToPayok := opts.RequestToPay.Value().(RequestToPay)
		if !optRequestToPayok {
			return nil, reportError("requestToPay should be RequestToPay")
		}
		postBody = &optRequestToPay
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			headerParams["Ocp-Apim-Subscription-Key"] = key
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			queryParams.Add("subscription-key", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, path, httpMethod, postBody, headerParams, queryParams, formParams, formFileName, fileName, fileBytes)
	if err != nil {
		return nil, err
	}

	res, err := a.client.callAPI(r)
	if err != nil || res == nil {
		return res, err
	}

	body, err := _ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return res, err
	}

	if res.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  body,
			error: res.Status,
		}
		if res.StatusCode == 409 {
			var v ErrorReason
			err = a.client.decode(&v, body, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return res, newErr
			}
			newErr.model = v
			return res, newErr
		}
		if res.StatusCode == 500 {
			var v ErrorReason
			err = a.client.decode(&v, body, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return res, newErr
			}
			newErr.model = v
		}
		return res, newErr
	}

	return res, nil
}

// RequesttopayReferenceIdGETOpts Optional parameters for the method 'RequesttopayReferenceIdGET'
type RequesttopayReferenceIdGETOpts struct {
	Authorization optional.String
}

/*
RequesttopayReferenceIdGET /requesttopay/{referenceId} - GET
This operation is used to get the status of a request to pay. X-Reference-Id that was passed in the post is used as reference to the request.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param referenceId UUID of transaction to get result. Reference id  used when creating the request to pay.
 * @param xTargetEnvironment The identifier of the EWP system where the transaction shall be processed. This parameter is used to route the request to the EWP system that will initiate the transaction.
 * @param optional nil or *RequesttopayReferenceIdGETOpts - Optional Parameters:
 * @param "Authorization" (optional.String) -  Authorization header used for Basic authentication and oauth. Format of the header parameter follows the standard for Basic and Bearer. Oauth uses Bearer authentication type where the credential is the received access token.
@return RequestToPayResult
*/
func (a *MoMoCollection) RequesttopayReferenceIdGET(ctx _context.Context, referenceId string, xTargetEnvironment string, opts *RequesttopayReferenceIdGETOpts) (RequestToPayResult, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  RequestToPayResult
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/v1_0/requesttopay/{referenceId}"
	path = strings.Replace(path, "{"+"referenceId"+"}", _neturl.QueryEscape(parameterToString(referenceId, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"Successful request to pay", "Payer not found", "application/json", "Request to pay not found", "Unspecified internal error"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	if opts != nil && opts.Authorization.IsSet() {
		headerParams["Authorization"] = parameterToString(opts.Authorization.Value(), "")
	}
	headerParams["X-Target-Environment"] = parameterToString(xTargetEnvironment, "")
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			headerParams["Ocp-Apim-Subscription-Key"] = key
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			queryParams.Add("subscription-key", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, path, httpMethod, postBody, headerParams, queryParams, formParams, formFileName, fileName, fileBytes)
	if err != nil {
		return returnValue, nil, err
	}

	res, err := a.client.callAPI(r)
	if err != nil || res == nil {
		return returnValue, res, err
	}

	body, err := _ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return returnValue, res, err
	}

	if res.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  body,
			error: res.Status,
		}
		if res.StatusCode == 404 {
			var v ErrorReason
			err = a.client.decode(&v, body, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, res, newErr
			}
			newErr.model = v
			return returnValue, res, newErr
		}
		if res.StatusCode == 500 {
			var v ErrorReason
			err = a.client.decode(&v, body, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, res, newErr
			}
			newErr.model = v
		}
		return returnValue, res, newErr
	}

	err = a.client.decode(&returnValue, body, res.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  body,
			error: err.Error(),
		}
		return returnValue, res, newErr
	}

	return returnValue, res, nil
}

/*
TokenPOST /token - POST
This operation is used to create an access token which can then be used to authorize and authenticate towards the other end-points of the API.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization Basic authentication header containing API user ID and API key. Should be sent in as B64 encoded.
@return TokenPost200ApplicationJsonResponse
*/
func (a *MoMoCollection) TokenPOST(ctx _context.Context, authorization string) (TokenPost200ApplicationJsonResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  TokenPost200ApplicationJsonResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/token/"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	headerParams["Authorization"] = parameterToString(authorization, "")
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			headerParams["Ocp-Apim-Subscription-Key"] = key
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			queryParams.Add("subscription-key", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, path, httpMethod, postBody, headerParams, queryParams, formParams, formFileName, fileName, fileBytes)
	if err != nil {
		return returnValue, nil, err
	}

	res, err := a.client.callAPI(r)
	if err != nil || res == nil {
		return returnValue, res, err
	}

	body, err := _ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return returnValue, res, err
	}

	if res.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  body,
			error: res.Status,
		}
		if res.StatusCode == 401 {
			var v TokenPost401ApplicationJsonResponse
			err = a.client.decode(&v, body, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, res, newErr
			}
			newErr.model = v
			return returnValue, res, newErr
		}
		return returnValue, res, newErr
	}

	err = a.client.decode(&returnValue, body, res.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  body,
			error: err.Error(),
		}
		return returnValue, res, newErr
	}

	return returnValue, res, nil
}
