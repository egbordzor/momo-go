package collection_go

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wondenge/momo-go/pkg/models"
)

// RequesttopayPOSTReader is a Reader for the RequesttopayPOST structure.
type RequesttopayPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequesttopayPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewRequesttopayPOSTAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRequesttopayPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRequesttopayPOSTConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRequesttopayPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRequesttopayPOSTAccepted creates a RequesttopayPOSTAccepted with default headers values
func NewRequesttopayPOSTAccepted() *RequesttopayPOSTAccepted {
	return &RequesttopayPOSTAccepted{}
}

/*RequesttopayPOSTAccepted handles this case with default header values.

Accepted
*/
type RequesttopayPOSTAccepted struct {
}

func (o *RequesttopayPOSTAccepted) Error() string {
	return fmt.Sprintf("[POST /v1_0/requesttopay][%d] requesttopayPOSTAccepted ", 202)
}

func (o *RequesttopayPOSTAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequesttopayPOSTBadRequest creates a RequesttopayPOSTBadRequest with default headers values
func NewRequesttopayPOSTBadRequest() *RequesttopayPOSTBadRequest {
	return &RequesttopayPOSTBadRequest{}
}

/*RequesttopayPOSTBadRequest handles this case with default header values.

Bad request, e.g. invalid data was sent in the request.
*/
type RequesttopayPOSTBadRequest struct {
}

func (o *RequesttopayPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1_0/requesttopay][%d] requesttopayPOSTBadRequest ", 400)
}

func (o *RequesttopayPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequesttopayPOSTConflict creates a RequesttopayPOSTConflict with default headers values
func NewRequesttopayPOSTConflict() *RequesttopayPOSTConflict {
	return &RequesttopayPOSTConflict{}
}

/*RequesttopayPOSTConflict handles this case with default header values.

Conflict, duplicated reference id
*/
type RequesttopayPOSTConflict struct {
	Payload *models.ErrorReason
}

func (o *RequesttopayPOSTConflict) Error() string {
	return fmt.Sprintf("[POST /v1_0/requesttopay][%d] requesttopayPOSTConflict  %+v", 409, o.Payload)
}

func (o *RequesttopayPOSTConflict) GetPayload() *models.ErrorReason {
	return o.Payload
}

func (o *RequesttopayPOSTConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorReason)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequesttopayPOSTInternalServerError creates a RequesttopayPOSTInternalServerError with default headers values
func NewRequesttopayPOSTInternalServerError() *RequesttopayPOSTInternalServerError {
	return &RequesttopayPOSTInternalServerError{}
}

/*RequesttopayPOSTInternalServerError handles this case with default header values.

Internal Error.
*/
type RequesttopayPOSTInternalServerError struct {
	Payload *models.ErrorReason
}

func (o *RequesttopayPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1_0/requesttopay][%d] requesttopayPOSTInternalServerError  %+v", 500, o.Payload)
}

func (o *RequesttopayPOSTInternalServerError) GetPayload() *models.ErrorReason {
	return o.Payload
}

func (o *RequesttopayPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorReason)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
