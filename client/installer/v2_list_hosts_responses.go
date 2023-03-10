// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2ListHostsReader is a Reader for the V2ListHosts structure.
type V2ListHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2ListHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2ListHostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV2ListHostsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2ListHostsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2ListHostsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2ListHostsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewV2ListHostsNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewV2ListHostsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2ListHostsOK creates a V2ListHostsOK with default headers values
func NewV2ListHostsOK() *V2ListHostsOK {
	return &V2ListHostsOK{}
}

/*
V2ListHostsOK describes a response with status code 200, with default header values.

Success.
*/
type V2ListHostsOK struct {
	Payload models.HostList
}

// IsSuccess returns true when this v2 list hosts o k response has a 2xx status code
func (o *V2ListHostsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v2 list hosts o k response has a 3xx status code
func (o *V2ListHostsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list hosts o k response has a 4xx status code
func (o *V2ListHostsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 list hosts o k response has a 5xx status code
func (o *V2ListHostsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 list hosts o k response a status code equal to that given
func (o *V2ListHostsOK) IsCode(code int) bool {
	return code == 200
}

func (o *V2ListHostsOK) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts][%d] v2ListHostsOK  %+v", 200, o.Payload)
}

func (o *V2ListHostsOK) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts][%d] v2ListHostsOK  %+v", 200, o.Payload)
}

func (o *V2ListHostsOK) GetPayload() models.HostList {
	return o.Payload
}

func (o *V2ListHostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListHostsUnauthorized creates a V2ListHostsUnauthorized with default headers values
func NewV2ListHostsUnauthorized() *V2ListHostsUnauthorized {
	return &V2ListHostsUnauthorized{}
}

/*
V2ListHostsUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type V2ListHostsUnauthorized struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 list hosts unauthorized response has a 2xx status code
func (o *V2ListHostsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 list hosts unauthorized response has a 3xx status code
func (o *V2ListHostsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list hosts unauthorized response has a 4xx status code
func (o *V2ListHostsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 list hosts unauthorized response has a 5xx status code
func (o *V2ListHostsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 list hosts unauthorized response a status code equal to that given
func (o *V2ListHostsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *V2ListHostsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts][%d] v2ListHostsUnauthorized  %+v", 401, o.Payload)
}

func (o *V2ListHostsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts][%d] v2ListHostsUnauthorized  %+v", 401, o.Payload)
}

func (o *V2ListHostsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ListHostsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListHostsForbidden creates a V2ListHostsForbidden with default headers values
func NewV2ListHostsForbidden() *V2ListHostsForbidden {
	return &V2ListHostsForbidden{}
}

/*
V2ListHostsForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type V2ListHostsForbidden struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 list hosts forbidden response has a 2xx status code
func (o *V2ListHostsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 list hosts forbidden response has a 3xx status code
func (o *V2ListHostsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list hosts forbidden response has a 4xx status code
func (o *V2ListHostsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 list hosts forbidden response has a 5xx status code
func (o *V2ListHostsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 list hosts forbidden response a status code equal to that given
func (o *V2ListHostsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *V2ListHostsForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts][%d] v2ListHostsForbidden  %+v", 403, o.Payload)
}

func (o *V2ListHostsForbidden) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts][%d] v2ListHostsForbidden  %+v", 403, o.Payload)
}

func (o *V2ListHostsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ListHostsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListHostsMethodNotAllowed creates a V2ListHostsMethodNotAllowed with default headers values
func NewV2ListHostsMethodNotAllowed() *V2ListHostsMethodNotAllowed {
	return &V2ListHostsMethodNotAllowed{}
}

/*
V2ListHostsMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type V2ListHostsMethodNotAllowed struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 list hosts method not allowed response has a 2xx status code
func (o *V2ListHostsMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 list hosts method not allowed response has a 3xx status code
func (o *V2ListHostsMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list hosts method not allowed response has a 4xx status code
func (o *V2ListHostsMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 list hosts method not allowed response has a 5xx status code
func (o *V2ListHostsMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 list hosts method not allowed response a status code equal to that given
func (o *V2ListHostsMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *V2ListHostsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts][%d] v2ListHostsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2ListHostsMethodNotAllowed) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts][%d] v2ListHostsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2ListHostsMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListHostsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListHostsInternalServerError creates a V2ListHostsInternalServerError with default headers values
func NewV2ListHostsInternalServerError() *V2ListHostsInternalServerError {
	return &V2ListHostsInternalServerError{}
}

/*
V2ListHostsInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type V2ListHostsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 list hosts internal server error response has a 2xx status code
func (o *V2ListHostsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 list hosts internal server error response has a 3xx status code
func (o *V2ListHostsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list hosts internal server error response has a 4xx status code
func (o *V2ListHostsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 list hosts internal server error response has a 5xx status code
func (o *V2ListHostsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 list hosts internal server error response a status code equal to that given
func (o *V2ListHostsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *V2ListHostsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts][%d] v2ListHostsInternalServerError  %+v", 500, o.Payload)
}

func (o *V2ListHostsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts][%d] v2ListHostsInternalServerError  %+v", 500, o.Payload)
}

func (o *V2ListHostsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListHostsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListHostsNotImplemented creates a V2ListHostsNotImplemented with default headers values
func NewV2ListHostsNotImplemented() *V2ListHostsNotImplemented {
	return &V2ListHostsNotImplemented{}
}

/*
V2ListHostsNotImplemented describes a response with status code 501, with default header values.

Not implemented.
*/
type V2ListHostsNotImplemented struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 list hosts not implemented response has a 2xx status code
func (o *V2ListHostsNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 list hosts not implemented response has a 3xx status code
func (o *V2ListHostsNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list hosts not implemented response has a 4xx status code
func (o *V2ListHostsNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 list hosts not implemented response has a 5xx status code
func (o *V2ListHostsNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 list hosts not implemented response a status code equal to that given
func (o *V2ListHostsNotImplemented) IsCode(code int) bool {
	return code == 501
}

func (o *V2ListHostsNotImplemented) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts][%d] v2ListHostsNotImplemented  %+v", 501, o.Payload)
}

func (o *V2ListHostsNotImplemented) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts][%d] v2ListHostsNotImplemented  %+v", 501, o.Payload)
}

func (o *V2ListHostsNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListHostsNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListHostsServiceUnavailable creates a V2ListHostsServiceUnavailable with default headers values
func NewV2ListHostsServiceUnavailable() *V2ListHostsServiceUnavailable {
	return &V2ListHostsServiceUnavailable{}
}

/*
V2ListHostsServiceUnavailable describes a response with status code 503, with default header values.

Unavailable.
*/
type V2ListHostsServiceUnavailable struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 list hosts service unavailable response has a 2xx status code
func (o *V2ListHostsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 list hosts service unavailable response has a 3xx status code
func (o *V2ListHostsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list hosts service unavailable response has a 4xx status code
func (o *V2ListHostsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 list hosts service unavailable response has a 5xx status code
func (o *V2ListHostsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 list hosts service unavailable response a status code equal to that given
func (o *V2ListHostsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *V2ListHostsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts][%d] v2ListHostsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *V2ListHostsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts][%d] v2ListHostsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *V2ListHostsServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListHostsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
