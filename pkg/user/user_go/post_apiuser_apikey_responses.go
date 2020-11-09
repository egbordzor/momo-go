package user_go

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wondenge/momo-go/pkg/models"
)

// PostApiuserApikeyReader is a Reader for the PostApiuserApikey structure.
type PostApiuserApikeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostApiuserApikeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostApiuserApikeyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostApiuserApikeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostApiuserApikeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostApiuserApikeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostApiuserApikeyCreated creates a PostApiuserApikeyCreated with default headers values
func NewPostApiuserApikeyCreated() *PostApiuserApikeyCreated {
	return &PostApiuserApikeyCreated{}
}

/*PostApiuserApikeyCreated handles this case with default header values.

Created
*/
type PostApiuserApikeyCreated struct {
	Payload *models.APIUserKeyResult
}

func (o *PostApiuserApikeyCreated) Error() string {
	return fmt.Sprintf("[POST /v1_0/apiuser/{X-Reference-Id}/apikey][%d] postApiuserApikeyCreated  %+v", 201, o.Payload)
}

func (o *PostApiuserApikeyCreated) GetPayload() *models.APIUserKeyResult {
	return o.Payload
}

func (o *PostApiuserApikeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIUserKeyResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApiuserApikeyBadRequest creates a PostApiuserApikeyBadRequest with default headers values
func NewPostApiuserApikeyBadRequest() *PostApiuserApikeyBadRequest {
	return &PostApiuserApikeyBadRequest{}
}

/*PostApiuserApikeyBadRequest handles this case with default header values.

Bad request, e.g. invalid data was sent in the request.
*/
type PostApiuserApikeyBadRequest struct {
}

func (o *PostApiuserApikeyBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1_0/apiuser/{X-Reference-Id}/apikey][%d] postApiuserApikeyBadRequest ", 400)
}

func (o *PostApiuserApikeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostApiuserApikeyNotFound creates a PostApiuserApikeyNotFound with default headers values
func NewPostApiuserApikeyNotFound() *PostApiuserApikeyNotFound {
	return &PostApiuserApikeyNotFound{}
}

/*PostApiuserApikeyNotFound handles this case with default header values.

Not found, reference id not found or closed in sandbox
*/
type PostApiuserApikeyNotFound struct {
	Payload *models.ErrorReason
}

func (o *PostApiuserApikeyNotFound) Error() string {
	return fmt.Sprintf("[POST /v1_0/apiuser/{X-Reference-Id}/apikey][%d] postApiuserApikeyNotFound  %+v", 404, o.Payload)
}

func (o *PostApiuserApikeyNotFound) GetPayload() *models.ErrorReason {
	return o.Payload
}

func (o *PostApiuserApikeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorReason)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApiuserApikeyInternalServerError creates a PostApiuserApikeyInternalServerError with default headers values
func NewPostApiuserApikeyInternalServerError() *PostApiuserApikeyInternalServerError {
	return &PostApiuserApikeyInternalServerError{}
}

/*PostApiuserApikeyInternalServerError handles this case with default header values.

Internal error. Check log for information.
*/
type PostApiuserApikeyInternalServerError struct {
}

func (o *PostApiuserApikeyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1_0/apiuser/{X-Reference-Id}/apikey][%d] postApiuserApikeyInternalServerError ", 500)
}

func (o *PostApiuserApikeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
