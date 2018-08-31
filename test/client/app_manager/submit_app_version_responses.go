// Code generated by go-swagger; DO NOT EDIT.

package app_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// SubmitAppVersionReader is a Reader for the SubmitAppVersion structure.
type SubmitAppVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitAppVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSubmitAppVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubmitAppVersionOK creates a SubmitAppVersionOK with default headers values
func NewSubmitAppVersionOK() *SubmitAppVersionOK {
	return &SubmitAppVersionOK{}
}

/*SubmitAppVersionOK handles this case with default header values.

SubmitAppVersionOK submit app version o k
*/
type SubmitAppVersionOK struct {
	Payload *models.OpenpitrixSubmitAppVersionResponse
}

func (o *SubmitAppVersionOK) Error() string {
	return fmt.Sprintf("[POST /v1/app_version/action/submit][%d] submitAppVersionOK  %+v", 200, o.Payload)
}

func (o *SubmitAppVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixSubmitAppVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}