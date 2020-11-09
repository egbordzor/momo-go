package disbursements_go

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wondenge/momo-go/pkg/models"
)

// TokenPOSTReader is a Reader for the TokenPOST structure.
type TokenPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokenPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTokenPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewTokenPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTokenPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTokenPOSTOK creates a TokenPOSTOK with default headers values
func NewTokenPOSTOK() *TokenPOSTOK {
	return &TokenPOSTOK{}
}

/*TokenPOSTOK handles this case with default header values.

OK
*/
type TokenPOSTOK struct {
	Payload *models.TokenPost200ApplicationJSONResponse
}

func (o *TokenPOSTOK) Error() string {
	return fmt.Sprintf("[POST /token/][%d] tokenPOSTOK  %+v", 200, o.Payload)
}

func (o *TokenPOSTOK) GetPayload() *models.TokenPost200ApplicationJSONResponse {
	return o.Payload
}

func (o *TokenPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenPost200ApplicationJSONResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenPOSTUnauthorized creates a TokenPOSTUnauthorized with default headers values
func NewTokenPOSTUnauthorized() *TokenPOSTUnauthorized {
	return &TokenPOSTUnauthorized{}
}

/*TokenPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type TokenPOSTUnauthorized struct {
	Payload *models.TokenPost401ApplicationJSONResponse
}

func (o *TokenPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /token/][%d] tokenPOSTUnauthorized  %+v", 401, o.Payload)
}

func (o *TokenPOSTUnauthorized) GetPayload() *models.TokenPost401ApplicationJSONResponse {
	return o.Payload
}

func (o *TokenPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenPost401ApplicationJSONResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenPOSTInternalServerError creates a TokenPOSTInternalServerError with default headers values
func NewTokenPOSTInternalServerError() *TokenPOSTInternalServerError {
	return &TokenPOSTInternalServerError{}
}

/*TokenPOSTInternalServerError handles this case with default header values.

Error
*/
type TokenPOSTInternalServerError struct {
}

func (o *TokenPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /token/][%d] tokenPOSTInternalServerError ", 500)
}

func (o *TokenPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
