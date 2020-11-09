package collection_go

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wondenge/momo-go/pkg/models"
)

// GetAccountholderAccountholderidtypeAccountholderidActiveReader is a Reader for the GetAccountholderAccountholderidtypeAccountholderidActive structure.
type GetAccountholderAccountholderidtypeAccountholderidActiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountholderAccountholderidtypeAccountholderidActiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountholderAccountholderidtypeAccountholderidActiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountholderAccountholderidtypeAccountholderidActiveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountholderAccountholderidtypeAccountholderidActiveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountholderAccountholderidtypeAccountholderidActiveOK creates a GetAccountholderAccountholderidtypeAccountholderidActiveOK with default headers values
func NewGetAccountholderAccountholderidtypeAccountholderidActiveOK() *GetAccountholderAccountholderidtypeAccountholderidActiveOK {
	return &GetAccountholderAccountholderidtypeAccountholderidActiveOK{}
}

/*GetAccountholderAccountholderidtypeAccountholderidActiveOK handles this case with default header values.

Ok. True if account holder is registered and active, false if the account holder is not active or not found
*/
type GetAccountholderAccountholderidtypeAccountholderidActiveOK struct {
}

func (o *GetAccountholderAccountholderidtypeAccountholderidActiveOK) Error() string {
	return fmt.Sprintf("[GET /v1_0/accountholder/{accountHolderIdType}/{accountHolderId}/active][%d] getAccountholderAccountholderidtypeAccountholderidActiveOK ", 200)
}

func (o *GetAccountholderAccountholderidtypeAccountholderidActiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccountholderAccountholderidtypeAccountholderidActiveBadRequest creates a GetAccountholderAccountholderidtypeAccountholderidActiveBadRequest with default headers values
func NewGetAccountholderAccountholderidtypeAccountholderidActiveBadRequest() *GetAccountholderAccountholderidtypeAccountholderidActiveBadRequest {
	return &GetAccountholderAccountholderidtypeAccountholderidActiveBadRequest{}
}

/*GetAccountholderAccountholderidtypeAccountholderidActiveBadRequest handles this case with default header values.

Bad request, e.g. invalid data was sent in the request.
*/
type GetAccountholderAccountholderidtypeAccountholderidActiveBadRequest struct {
}

func (o *GetAccountholderAccountholderidtypeAccountholderidActiveBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1_0/accountholder/{accountHolderIdType}/{accountHolderId}/active][%d] getAccountholderAccountholderidtypeAccountholderidActiveBadRequest ", 400)
}

func (o *GetAccountholderAccountholderidtypeAccountholderidActiveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccountholderAccountholderidtypeAccountholderidActiveInternalServerError creates a GetAccountholderAccountholderidtypeAccountholderidActiveInternalServerError with default headers values
func NewGetAccountholderAccountholderidtypeAccountholderidActiveInternalServerError() *GetAccountholderAccountholderidtypeAccountholderidActiveInternalServerError {
	return &GetAccountholderAccountholderidtypeAccountholderidActiveInternalServerError{}
}

/*GetAccountholderAccountholderidtypeAccountholderidActiveInternalServerError handles this case with default header values.

Internal error. The returned response contains details.
*/
type GetAccountholderAccountholderidtypeAccountholderidActiveInternalServerError struct {
	Payload *models.ErrorReason
}

func (o *GetAccountholderAccountholderidtypeAccountholderidActiveInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1_0/accountholder/{accountHolderIdType}/{accountHolderId}/active][%d] getAccountholderAccountholderidtypeAccountholderidActiveInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAccountholderAccountholderidtypeAccountholderidActiveInternalServerError) GetPayload() *models.ErrorReason {
	return o.Payload
}

func (o *GetAccountholderAccountholderidtypeAccountholderidActiveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorReason)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
