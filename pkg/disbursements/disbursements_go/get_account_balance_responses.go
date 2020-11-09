package disbursements_go

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wondenge/momo-go/pkg/models"
)

// GetAccountBalanceReader is a Reader for the GetAccountBalance structure.
type GetAccountBalanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountBalanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountBalanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountBalanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountBalanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountBalanceOK creates a GetAccountBalanceOK with default headers values
func NewGetAccountBalanceOK() *GetAccountBalanceOK {
	return &GetAccountBalanceOK{}
}

/*GetAccountBalanceOK handles this case with default header values.

Ok
*/
type GetAccountBalanceOK struct {
	Payload *models.Balance
}

func (o *GetAccountBalanceOK) Error() string {
	return fmt.Sprintf("[GET /v1_0/account/balance][%d] getAccountBalanceOK  %+v", 200, o.Payload)
}

func (o *GetAccountBalanceOK) GetPayload() *models.Balance {
	return o.Payload
}

func (o *GetAccountBalanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Balance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountBalanceBadRequest creates a GetAccountBalanceBadRequest with default headers values
func NewGetAccountBalanceBadRequest() *GetAccountBalanceBadRequest {
	return &GetAccountBalanceBadRequest{}
}

/*GetAccountBalanceBadRequest handles this case with default header values.

Bad request, e.g. invalid data was sent in the request.
*/
type GetAccountBalanceBadRequest struct {
}

func (o *GetAccountBalanceBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1_0/account/balance][%d] getAccountBalanceBadRequest ", 400)
}

func (o *GetAccountBalanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccountBalanceInternalServerError creates a GetAccountBalanceInternalServerError with default headers values
func NewGetAccountBalanceInternalServerError() *GetAccountBalanceInternalServerError {
	return &GetAccountBalanceInternalServerError{}
}

/*GetAccountBalanceInternalServerError handles this case with default header values.

Internal error. The returned response contains details.
*/
type GetAccountBalanceInternalServerError struct {
	Payload *models.ErrorReason
}

func (o *GetAccountBalanceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1_0/account/balance][%d] getAccountBalanceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAccountBalanceInternalServerError) GetPayload() *models.ErrorReason {
	return o.Payload
}

func (o *GetAccountBalanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorReason)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
