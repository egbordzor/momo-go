package disbursements_go

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wondenge/momo-go/pkg/models"
)

// TransferReferenceIDGETReader is a Reader for the TransferReferenceIDGET structure.
type TransferReferenceIDGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TransferReferenceIDGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTransferReferenceIDGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTransferReferenceIDGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTransferReferenceIDGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTransferReferenceIDGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTransferReferenceIDGETOK creates a TransferReferenceIDGETOK with default headers values
func NewTransferReferenceIDGETOK() *TransferReferenceIDGETOK {
	return &TransferReferenceIDGETOK{}
}

/*TransferReferenceIDGETOK handles this case with default header values.

OK. Note that a failed transfer will be returned with this status too. The 'status' of the TransferResult can be used to determine the outcome of the request. The 'reason' field can be used to retrieve a cause in case of failure.
*/
type TransferReferenceIDGETOK struct {
	Payload *models.TransferResult
}

func (o *TransferReferenceIDGETOK) Error() string {
	return fmt.Sprintf("[GET /v1_0/transfer/{referenceId}][%d] transferReferenceIdGETOK  %+v", 200, o.Payload)
}

func (o *TransferReferenceIDGETOK) GetPayload() *models.TransferResult {
	return o.Payload
}

func (o *TransferReferenceIDGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TransferResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransferReferenceIDGETBadRequest creates a TransferReferenceIDGETBadRequest with default headers values
func NewTransferReferenceIDGETBadRequest() *TransferReferenceIDGETBadRequest {
	return &TransferReferenceIDGETBadRequest{}
}

/*TransferReferenceIDGETBadRequest handles this case with default header values.

Bad request, e.g. an incorrectly formatted reference id was provided.
*/
type TransferReferenceIDGETBadRequest struct {
}

func (o *TransferReferenceIDGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1_0/transfer/{referenceId}][%d] transferReferenceIdGETBadRequest ", 400)
}

func (o *TransferReferenceIDGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransferReferenceIDGETNotFound creates a TransferReferenceIDGETNotFound with default headers values
func NewTransferReferenceIDGETNotFound() *TransferReferenceIDGETNotFound {
	return &TransferReferenceIDGETNotFound{}
}

/*TransferReferenceIDGETNotFound handles this case with default header values.

Resource not found.
*/
type TransferReferenceIDGETNotFound struct {
	Payload *models.ErrorReason
}

func (o *TransferReferenceIDGETNotFound) Error() string {
	return fmt.Sprintf("[GET /v1_0/transfer/{referenceId}][%d] transferReferenceIdGETNotFound  %+v", 404, o.Payload)
}

func (o *TransferReferenceIDGETNotFound) GetPayload() *models.ErrorReason {
	return o.Payload
}

func (o *TransferReferenceIDGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorReason)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransferReferenceIDGETInternalServerError creates a TransferReferenceIDGETInternalServerError with default headers values
func NewTransferReferenceIDGETInternalServerError() *TransferReferenceIDGETInternalServerError {
	return &TransferReferenceIDGETInternalServerError{}
}

/*TransferReferenceIDGETInternalServerError handles this case with default header values.

Internal Error. Note that if the retreieved transfer has failed, it will not cause this status to be returned. This status is only returned if the GET request itself fails.
*/
type TransferReferenceIDGETInternalServerError struct {
	Payload *models.ErrorReason
}

func (o *TransferReferenceIDGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1_0/transfer/{referenceId}][%d] transferReferenceIdGETInternalServerError  %+v", 500, o.Payload)
}

func (o *TransferReferenceIDGETInternalServerError) GetPayload() *models.ErrorReason {
	return o.Payload
}

func (o *TransferReferenceIDGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorReason)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
