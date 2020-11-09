package user_go

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wondenge/momo-go/pkg/models"
)

// GetApiuserReader is a Reader for the GetApiuser structure.
type GetApiuserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApiuserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApiuserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApiuserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetApiuserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetApiuserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetApiuserOK creates a GetApiuserOK with default headers values
func NewGetApiuserOK() *GetApiuserOK {
	return &GetApiuserOK{}
}

/*GetApiuserOK handles this case with default header values.

Ok
*/
type GetApiuserOK struct {
}

func (o *GetApiuserOK) Error() string {
	return fmt.Sprintf("[GET /v1_0/apiuser/{X-Reference-Id}][%d] getApiuserOK ", 200)
}

func (o *GetApiuserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApiuserBadRequest creates a GetApiuserBadRequest with default headers values
func NewGetApiuserBadRequest() *GetApiuserBadRequest {
	return &GetApiuserBadRequest{}
}

/*GetApiuserBadRequest handles this case with default header values.

Bad request, e.g. invalid data was sent in the request.
*/
type GetApiuserBadRequest struct {
}

func (o *GetApiuserBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1_0/apiuser/{X-Reference-Id}][%d] getApiuserBadRequest ", 400)
}

func (o *GetApiuserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApiuserNotFound creates a GetApiuserNotFound with default headers values
func NewGetApiuserNotFound() *GetApiuserNotFound {
	return &GetApiuserNotFound{}
}

/*GetApiuserNotFound handles this case with default header values.

Not found, reference id not found or closed in sandbox
*/
type GetApiuserNotFound struct {
	Payload *models.ErrorReason
}

func (o *GetApiuserNotFound) Error() string {
	return fmt.Sprintf("[GET /v1_0/apiuser/{X-Reference-Id}][%d] getApiuserNotFound  %+v", 404, o.Payload)
}

func (o *GetApiuserNotFound) GetPayload() *models.ErrorReason {
	return o.Payload
}

func (o *GetApiuserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorReason)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApiuserInternalServerError creates a GetApiuserInternalServerError with default headers values
func NewGetApiuserInternalServerError() *GetApiuserInternalServerError {
	return &GetApiuserInternalServerError{}
}

/*GetApiuserInternalServerError handles this case with default header values.

Internal error. Check log for information.
*/
type GetApiuserInternalServerError struct {
}

func (o *GetApiuserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1_0/apiuser/{X-Reference-Id}][%d] getApiuserInternalServerError ", 500)
}

func (o *GetApiuserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
