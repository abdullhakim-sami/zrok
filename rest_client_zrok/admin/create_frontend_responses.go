// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/zrok/rest_model_zrok"
)

// CreateFrontendReader is a Reader for the CreateFrontend structure.
type CreateFrontendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFrontendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateFrontendCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateFrontendBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateFrontendUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateFrontendNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateFrontendInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateFrontendCreated creates a CreateFrontendCreated with default headers values
func NewCreateFrontendCreated() *CreateFrontendCreated {
	return &CreateFrontendCreated{}
}

/*
CreateFrontendCreated describes a response with status code 201, with default header values.

frontend created
*/
type CreateFrontendCreated struct {
	Payload *rest_model_zrok.CreateFrontendResponse
}

// IsSuccess returns true when this create frontend created response has a 2xx status code
func (o *CreateFrontendCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create frontend created response has a 3xx status code
func (o *CreateFrontendCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create frontend created response has a 4xx status code
func (o *CreateFrontendCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create frontend created response has a 5xx status code
func (o *CreateFrontendCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create frontend created response a status code equal to that given
func (o *CreateFrontendCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateFrontendCreated) Error() string {
	return fmt.Sprintf("[POST /frontend][%d] createFrontendCreated  %+v", 201, o.Payload)
}

func (o *CreateFrontendCreated) String() string {
	return fmt.Sprintf("[POST /frontend][%d] createFrontendCreated  %+v", 201, o.Payload)
}

func (o *CreateFrontendCreated) GetPayload() *rest_model_zrok.CreateFrontendResponse {
	return o.Payload
}

func (o *CreateFrontendCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model_zrok.CreateFrontendResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFrontendBadRequest creates a CreateFrontendBadRequest with default headers values
func NewCreateFrontendBadRequest() *CreateFrontendBadRequest {
	return &CreateFrontendBadRequest{}
}

/*
CreateFrontendBadRequest describes a response with status code 400, with default header values.

bad request
*/
type CreateFrontendBadRequest struct {
}

// IsSuccess returns true when this create frontend bad request response has a 2xx status code
func (o *CreateFrontendBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create frontend bad request response has a 3xx status code
func (o *CreateFrontendBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create frontend bad request response has a 4xx status code
func (o *CreateFrontendBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create frontend bad request response has a 5xx status code
func (o *CreateFrontendBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create frontend bad request response a status code equal to that given
func (o *CreateFrontendBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateFrontendBadRequest) Error() string {
	return fmt.Sprintf("[POST /frontend][%d] createFrontendBadRequest ", 400)
}

func (o *CreateFrontendBadRequest) String() string {
	return fmt.Sprintf("[POST /frontend][%d] createFrontendBadRequest ", 400)
}

func (o *CreateFrontendBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFrontendUnauthorized creates a CreateFrontendUnauthorized with default headers values
func NewCreateFrontendUnauthorized() *CreateFrontendUnauthorized {
	return &CreateFrontendUnauthorized{}
}

/*
CreateFrontendUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type CreateFrontendUnauthorized struct {
}

// IsSuccess returns true when this create frontend unauthorized response has a 2xx status code
func (o *CreateFrontendUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create frontend unauthorized response has a 3xx status code
func (o *CreateFrontendUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create frontend unauthorized response has a 4xx status code
func (o *CreateFrontendUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create frontend unauthorized response has a 5xx status code
func (o *CreateFrontendUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create frontend unauthorized response a status code equal to that given
func (o *CreateFrontendUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateFrontendUnauthorized) Error() string {
	return fmt.Sprintf("[POST /frontend][%d] createFrontendUnauthorized ", 401)
}

func (o *CreateFrontendUnauthorized) String() string {
	return fmt.Sprintf("[POST /frontend][%d] createFrontendUnauthorized ", 401)
}

func (o *CreateFrontendUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFrontendNotFound creates a CreateFrontendNotFound with default headers values
func NewCreateFrontendNotFound() *CreateFrontendNotFound {
	return &CreateFrontendNotFound{}
}

/*
CreateFrontendNotFound describes a response with status code 404, with default header values.

not found
*/
type CreateFrontendNotFound struct {
}

// IsSuccess returns true when this create frontend not found response has a 2xx status code
func (o *CreateFrontendNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create frontend not found response has a 3xx status code
func (o *CreateFrontendNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create frontend not found response has a 4xx status code
func (o *CreateFrontendNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create frontend not found response has a 5xx status code
func (o *CreateFrontendNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create frontend not found response a status code equal to that given
func (o *CreateFrontendNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateFrontendNotFound) Error() string {
	return fmt.Sprintf("[POST /frontend][%d] createFrontendNotFound ", 404)
}

func (o *CreateFrontendNotFound) String() string {
	return fmt.Sprintf("[POST /frontend][%d] createFrontendNotFound ", 404)
}

func (o *CreateFrontendNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFrontendInternalServerError creates a CreateFrontendInternalServerError with default headers values
func NewCreateFrontendInternalServerError() *CreateFrontendInternalServerError {
	return &CreateFrontendInternalServerError{}
}

/*
CreateFrontendInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type CreateFrontendInternalServerError struct {
}

// IsSuccess returns true when this create frontend internal server error response has a 2xx status code
func (o *CreateFrontendInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create frontend internal server error response has a 3xx status code
func (o *CreateFrontendInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create frontend internal server error response has a 4xx status code
func (o *CreateFrontendInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create frontend internal server error response has a 5xx status code
func (o *CreateFrontendInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create frontend internal server error response a status code equal to that given
func (o *CreateFrontendInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateFrontendInternalServerError) Error() string {
	return fmt.Sprintf("[POST /frontend][%d] createFrontendInternalServerError ", 500)
}

func (o *CreateFrontendInternalServerError) String() string {
	return fmt.Sprintf("[POST /frontend][%d] createFrontendInternalServerError ", 500)
}

func (o *CreateFrontendInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
