package collection_go

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wondenge/momo-go/pkg/models"
)

// RequesttopayReferenceIDGETReader is a Reader for the RequesttopayReferenceIDGET structure.
type RequesttopayReferenceIDGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequesttopayReferenceIDGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRequesttopayReferenceIDGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRequesttopayReferenceIDGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRequesttopayReferenceIDGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRequesttopayReferenceIDGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRequesttopayReferenceIDGETOK creates a RequesttopayReferenceIDGETOK with default headers values
func NewRequesttopayReferenceIDGETOK() *RequesttopayReferenceIDGETOK {
	return &RequesttopayReferenceIDGETOK{}
}

/*RequesttopayReferenceIDGETOK handles this case with default header values.

OK. Note that a  failed request to pay will be returned with this status too. The 'status' of the RequestToPayResult can be used to determine the outcome of the request. The 'reason' field can be used to retrieve a cause in case of failure.
*/
type RequesttopayReferenceIDGETOK struct {
	Payload *models.RequestToPayResult
}

func (o *RequesttopayReferenceIDGETOK) Error() string {
	return fmt.Sprintf("[GET /v1_0/requesttopay/{referenceId}][%d] requesttopayReferenceIdGETOK  %+v", 200, o.Payload)
}

func (o *RequesttopayReferenceIDGETOK) GetPayload() *models.RequestToPayResult {
	return o.Payload
}

func (o *RequesttopayReferenceIDGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestToPayResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequesttopayReferenceIDGETBadRequest creates a RequesttopayReferenceIDGETBadRequest with default headers values
func NewRequesttopayReferenceIDGETBadRequest() *RequesttopayReferenceIDGETBadRequest {
	return &RequesttopayReferenceIDGETBadRequest{}
}

/*RequesttopayReferenceIDGETBadRequest handles this case with default header values.

Bad request, e.g. an incorrectly formatted reference id was provided.
*/
type RequesttopayReferenceIDGETBadRequest struct {
}

func (o *RequesttopayReferenceIDGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1_0/requesttopay/{referenceId}][%d] requesttopayReferenceIdGETBadRequest ", 400)
}

func (o *RequesttopayReferenceIDGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequesttopayReferenceIDGETNotFound creates a RequesttopayReferenceIDGETNotFound with default headers values
func NewRequesttopayReferenceIDGETNotFound() *RequesttopayReferenceIDGETNotFound {
	return &RequesttopayReferenceIDGETNotFound{}
}

/*RequesttopayReferenceIDGETNotFound handles this case with default header values.

Resource not found.
*/
type RequesttopayReferenceIDGETNotFound struct {
	Payload *models.ErrorReason
}

func (o *RequesttopayReferenceIDGETNotFound) Error() string {
	return fmt.Sprintf("[GET /v1_0/requesttopay/{referenceId}][%d] requesttopayReferenceIdGETNotFound  %+v", 404, o.Payload)
}

func (o *RequesttopayReferenceIDGETNotFound) GetPayload() *models.ErrorReason {
	return o.Payload
}

func (o *RequesttopayReferenceIDGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorReason)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequesttopayReferenceIDGETInternalServerError creates a RequesttopayReferenceIDGETInternalServerError with default headers values
func NewRequesttopayReferenceIDGETInternalServerError() *RequesttopayReferenceIDGETInternalServerError {
	return &RequesttopayReferenceIDGETInternalServerError{}
}

/*RequesttopayReferenceIDGETInternalServerError handles this case with default header values.

Internal Error. Note that if the retrieved request to pay has failed, it will not cause this status to be returned. This status is only returned if the GET request itself fails.
*/
type RequesttopayReferenceIDGETInternalServerError struct {
	Payload *models.ErrorReason
}

func (o *RequesttopayReferenceIDGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1_0/requesttopay/{referenceId}][%d] requesttopayReferenceIdGETInternalServerError  %+v", 500, o.Payload)
}

func (o *RequesttopayReferenceIDGETInternalServerError) GetPayload() *models.ErrorReason {
	return o.Payload
}

func (o *RequesttopayReferenceIDGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorReason)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
