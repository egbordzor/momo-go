package remittance_go

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wondenge/momo-go/pkg/models"
)

// TransferPOSTReader is a Reader for the TransferPOST structure.
type TransferPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TransferPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewTransferPOSTAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTransferPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewTransferPOSTConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTransferPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTransferPOSTAccepted creates a TransferPOSTAccepted with default headers values
func NewTransferPOSTAccepted() *TransferPOSTAccepted {
	return &TransferPOSTAccepted{}
}

/*TransferPOSTAccepted handles this case with default header values.

Accepted
*/
type TransferPOSTAccepted struct {
}

func (o *TransferPOSTAccepted) Error() string {
	return fmt.Sprintf("[POST /v1_0/transfer][%d] transferPOSTAccepted ", 202)
}

func (o *TransferPOSTAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransferPOSTBadRequest creates a TransferPOSTBadRequest with default headers values
func NewTransferPOSTBadRequest() *TransferPOSTBadRequest {
	return &TransferPOSTBadRequest{}
}

/*TransferPOSTBadRequest handles this case with default header values.

Bad request, e.g. invalid data was sent in the request.
*/
type TransferPOSTBadRequest struct {
}

func (o *TransferPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1_0/transfer][%d] transferPOSTBadRequest ", 400)
}

func (o *TransferPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransferPOSTConflict creates a TransferPOSTConflict with default headers values
func NewTransferPOSTConflict() *TransferPOSTConflict {
	return &TransferPOSTConflict{}
}

/*TransferPOSTConflict handles this case with default header values.

Conflict, duplicated reference id
*/
type TransferPOSTConflict struct {
	Payload *models.ErrorReason
}

func (o *TransferPOSTConflict) Error() string {
	return fmt.Sprintf("[POST /v1_0/transfer][%d] transferPOSTConflict  %+v", 409, o.Payload)
}

func (o *TransferPOSTConflict) GetPayload() *models.ErrorReason {
	return o.Payload
}

func (o *TransferPOSTConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorReason)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransferPOSTInternalServerError creates a TransferPOSTInternalServerError with default headers values
func NewTransferPOSTInternalServerError() *TransferPOSTInternalServerError {
	return &TransferPOSTInternalServerError{}
}

/*TransferPOSTInternalServerError handles this case with default header values.

Internal Error.
*/
type TransferPOSTInternalServerError struct {
	Payload *models.ErrorReason
}

func (o *TransferPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1_0/transfer][%d] transferPOSTInternalServerError  %+v", 500, o.Payload)
}

func (o *TransferPOSTInternalServerError) GetPayload() *models.ErrorReason {
	return o.Payload
}

func (o *TransferPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorReason)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
