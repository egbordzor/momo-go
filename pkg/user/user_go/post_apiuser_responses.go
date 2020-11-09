package user_go

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wondenge/momo-go/pkg/models"
)

// PostApiuserReader is a Reader for the PostApiuser structure.
type PostApiuserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostApiuserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostApiuserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostApiuserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostApiuserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostApiuserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostApiuserCreated creates a PostApiuserCreated with default headers values
func NewPostApiuserCreated() *PostApiuserCreated {
	return &PostApiuserCreated{}
}

/*PostApiuserCreated handles this case with default header values.

Created
*/
type PostApiuserCreated struct {
}

func (o *PostApiuserCreated) Error() string {
	return fmt.Sprintf("[POST /v1_0/apiuser][%d] postApiuserCreated ", 201)
}

func (o *PostApiuserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostApiuserBadRequest creates a PostApiuserBadRequest with default headers values
func NewPostApiuserBadRequest() *PostApiuserBadRequest {
	return &PostApiuserBadRequest{}
}

/*PostApiuserBadRequest handles this case with default header values.

Bad request, e.g. invalid data was sent in the request.
*/
type PostApiuserBadRequest struct {
}

func (o *PostApiuserBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1_0/apiuser][%d] postApiuserBadRequest ", 400)
}

func (o *PostApiuserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostApiuserConflict creates a PostApiuserConflict with default headers values
func NewPostApiuserConflict() *PostApiuserConflict {
	return &PostApiuserConflict{}
}

/*PostApiuserConflict handles this case with default header values.

Conflict, duplicated reference id
*/
type PostApiuserConflict struct {
	Payload *models.ErrorReason
}

func (o *PostApiuserConflict) Error() string {
	return fmt.Sprintf("[POST /v1_0/apiuser][%d] postApiuserConflict  %+v", 409, o.Payload)
}

func (o *PostApiuserConflict) GetPayload() *models.ErrorReason {
	return o.Payload
}

func (o *PostApiuserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorReason)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApiuserInternalServerError creates a PostApiuserInternalServerError with default headers values
func NewPostApiuserInternalServerError() *PostApiuserInternalServerError {
	return &PostApiuserInternalServerError{}
}

/*PostApiuserInternalServerError handles this case with default header values.

Internal error. Check log for information.
*/
type PostApiuserInternalServerError struct {
}

func (o *PostApiuserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1_0/apiuser][%d] postApiuserInternalServerError ", 500)
}

func (o *PostApiuserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
