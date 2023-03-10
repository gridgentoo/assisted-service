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

// GetInfraEnvDownloadURLReader is a Reader for the GetInfraEnvDownloadURL structure.
type GetInfraEnvDownloadURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInfraEnvDownloadURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInfraEnvDownloadURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInfraEnvDownloadURLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInfraEnvDownloadURLUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInfraEnvDownloadURLForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInfraEnvDownloadURLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetInfraEnvDownloadURLMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInfraEnvDownloadURLInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewGetInfraEnvDownloadURLNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetInfraEnvDownloadURLServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInfraEnvDownloadURLOK creates a GetInfraEnvDownloadURLOK with default headers values
func NewGetInfraEnvDownloadURLOK() *GetInfraEnvDownloadURLOK {
	return &GetInfraEnvDownloadURLOK{}
}

/*
GetInfraEnvDownloadURLOK describes a response with status code 200, with default header values.

Success.
*/
type GetInfraEnvDownloadURLOK struct {
	Payload *models.PresignedURL
}

// IsSuccess returns true when this get infra env download Url o k response has a 2xx status code
func (o *GetInfraEnvDownloadURLOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get infra env download Url o k response has a 3xx status code
func (o *GetInfraEnvDownloadURLOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get infra env download Url o k response has a 4xx status code
func (o *GetInfraEnvDownloadURLOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get infra env download Url o k response has a 5xx status code
func (o *GetInfraEnvDownloadURLOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get infra env download Url o k response a status code equal to that given
func (o *GetInfraEnvDownloadURLOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInfraEnvDownloadURLOK) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlOK  %+v", 200, o.Payload)
}

func (o *GetInfraEnvDownloadURLOK) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlOK  %+v", 200, o.Payload)
}

func (o *GetInfraEnvDownloadURLOK) GetPayload() *models.PresignedURL {
	return o.Payload
}

func (o *GetInfraEnvDownloadURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PresignedURL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInfraEnvDownloadURLBadRequest creates a GetInfraEnvDownloadURLBadRequest with default headers values
func NewGetInfraEnvDownloadURLBadRequest() *GetInfraEnvDownloadURLBadRequest {
	return &GetInfraEnvDownloadURLBadRequest{}
}

/*
GetInfraEnvDownloadURLBadRequest describes a response with status code 400, with default header values.

Bad Request.
*/
type GetInfraEnvDownloadURLBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get infra env download Url bad request response has a 2xx status code
func (o *GetInfraEnvDownloadURLBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get infra env download Url bad request response has a 3xx status code
func (o *GetInfraEnvDownloadURLBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get infra env download Url bad request response has a 4xx status code
func (o *GetInfraEnvDownloadURLBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get infra env download Url bad request response has a 5xx status code
func (o *GetInfraEnvDownloadURLBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get infra env download Url bad request response a status code equal to that given
func (o *GetInfraEnvDownloadURLBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetInfraEnvDownloadURLBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlBadRequest  %+v", 400, o.Payload)
}

func (o *GetInfraEnvDownloadURLBadRequest) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlBadRequest  %+v", 400, o.Payload)
}

func (o *GetInfraEnvDownloadURLBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetInfraEnvDownloadURLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInfraEnvDownloadURLUnauthorized creates a GetInfraEnvDownloadURLUnauthorized with default headers values
func NewGetInfraEnvDownloadURLUnauthorized() *GetInfraEnvDownloadURLUnauthorized {
	return &GetInfraEnvDownloadURLUnauthorized{}
}

/*
GetInfraEnvDownloadURLUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type GetInfraEnvDownloadURLUnauthorized struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this get infra env download Url unauthorized response has a 2xx status code
func (o *GetInfraEnvDownloadURLUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get infra env download Url unauthorized response has a 3xx status code
func (o *GetInfraEnvDownloadURLUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get infra env download Url unauthorized response has a 4xx status code
func (o *GetInfraEnvDownloadURLUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get infra env download Url unauthorized response has a 5xx status code
func (o *GetInfraEnvDownloadURLUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get infra env download Url unauthorized response a status code equal to that given
func (o *GetInfraEnvDownloadURLUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetInfraEnvDownloadURLUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInfraEnvDownloadURLUnauthorized) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInfraEnvDownloadURLUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *GetInfraEnvDownloadURLUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInfraEnvDownloadURLForbidden creates a GetInfraEnvDownloadURLForbidden with default headers values
func NewGetInfraEnvDownloadURLForbidden() *GetInfraEnvDownloadURLForbidden {
	return &GetInfraEnvDownloadURLForbidden{}
}

/*
GetInfraEnvDownloadURLForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type GetInfraEnvDownloadURLForbidden struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this get infra env download Url forbidden response has a 2xx status code
func (o *GetInfraEnvDownloadURLForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get infra env download Url forbidden response has a 3xx status code
func (o *GetInfraEnvDownloadURLForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get infra env download Url forbidden response has a 4xx status code
func (o *GetInfraEnvDownloadURLForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get infra env download Url forbidden response has a 5xx status code
func (o *GetInfraEnvDownloadURLForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get infra env download Url forbidden response a status code equal to that given
func (o *GetInfraEnvDownloadURLForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetInfraEnvDownloadURLForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlForbidden  %+v", 403, o.Payload)
}

func (o *GetInfraEnvDownloadURLForbidden) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlForbidden  %+v", 403, o.Payload)
}

func (o *GetInfraEnvDownloadURLForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *GetInfraEnvDownloadURLForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInfraEnvDownloadURLNotFound creates a GetInfraEnvDownloadURLNotFound with default headers values
func NewGetInfraEnvDownloadURLNotFound() *GetInfraEnvDownloadURLNotFound {
	return &GetInfraEnvDownloadURLNotFound{}
}

/*
GetInfraEnvDownloadURLNotFound describes a response with status code 404, with default header values.

Error.
*/
type GetInfraEnvDownloadURLNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get infra env download Url not found response has a 2xx status code
func (o *GetInfraEnvDownloadURLNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get infra env download Url not found response has a 3xx status code
func (o *GetInfraEnvDownloadURLNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get infra env download Url not found response has a 4xx status code
func (o *GetInfraEnvDownloadURLNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get infra env download Url not found response has a 5xx status code
func (o *GetInfraEnvDownloadURLNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get infra env download Url not found response a status code equal to that given
func (o *GetInfraEnvDownloadURLNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetInfraEnvDownloadURLNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlNotFound  %+v", 404, o.Payload)
}

func (o *GetInfraEnvDownloadURLNotFound) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlNotFound  %+v", 404, o.Payload)
}

func (o *GetInfraEnvDownloadURLNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetInfraEnvDownloadURLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInfraEnvDownloadURLMethodNotAllowed creates a GetInfraEnvDownloadURLMethodNotAllowed with default headers values
func NewGetInfraEnvDownloadURLMethodNotAllowed() *GetInfraEnvDownloadURLMethodNotAllowed {
	return &GetInfraEnvDownloadURLMethodNotAllowed{}
}

/*
GetInfraEnvDownloadURLMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type GetInfraEnvDownloadURLMethodNotAllowed struct {
	Payload *models.Error
}

// IsSuccess returns true when this get infra env download Url method not allowed response has a 2xx status code
func (o *GetInfraEnvDownloadURLMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get infra env download Url method not allowed response has a 3xx status code
func (o *GetInfraEnvDownloadURLMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get infra env download Url method not allowed response has a 4xx status code
func (o *GetInfraEnvDownloadURLMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this get infra env download Url method not allowed response has a 5xx status code
func (o *GetInfraEnvDownloadURLMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this get infra env download Url method not allowed response a status code equal to that given
func (o *GetInfraEnvDownloadURLMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *GetInfraEnvDownloadURLMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetInfraEnvDownloadURLMethodNotAllowed) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetInfraEnvDownloadURLMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetInfraEnvDownloadURLMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInfraEnvDownloadURLInternalServerError creates a GetInfraEnvDownloadURLInternalServerError with default headers values
func NewGetInfraEnvDownloadURLInternalServerError() *GetInfraEnvDownloadURLInternalServerError {
	return &GetInfraEnvDownloadURLInternalServerError{}
}

/*
GetInfraEnvDownloadURLInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type GetInfraEnvDownloadURLInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get infra env download Url internal server error response has a 2xx status code
func (o *GetInfraEnvDownloadURLInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get infra env download Url internal server error response has a 3xx status code
func (o *GetInfraEnvDownloadURLInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get infra env download Url internal server error response has a 4xx status code
func (o *GetInfraEnvDownloadURLInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get infra env download Url internal server error response has a 5xx status code
func (o *GetInfraEnvDownloadURLInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get infra env download Url internal server error response a status code equal to that given
func (o *GetInfraEnvDownloadURLInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetInfraEnvDownloadURLInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInfraEnvDownloadURLInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInfraEnvDownloadURLInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetInfraEnvDownloadURLInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInfraEnvDownloadURLNotImplemented creates a GetInfraEnvDownloadURLNotImplemented with default headers values
func NewGetInfraEnvDownloadURLNotImplemented() *GetInfraEnvDownloadURLNotImplemented {
	return &GetInfraEnvDownloadURLNotImplemented{}
}

/*
GetInfraEnvDownloadURLNotImplemented describes a response with status code 501, with default header values.

Not implemented.
*/
type GetInfraEnvDownloadURLNotImplemented struct {
	Payload *models.Error
}

// IsSuccess returns true when this get infra env download Url not implemented response has a 2xx status code
func (o *GetInfraEnvDownloadURLNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get infra env download Url not implemented response has a 3xx status code
func (o *GetInfraEnvDownloadURLNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get infra env download Url not implemented response has a 4xx status code
func (o *GetInfraEnvDownloadURLNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this get infra env download Url not implemented response has a 5xx status code
func (o *GetInfraEnvDownloadURLNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this get infra env download Url not implemented response a status code equal to that given
func (o *GetInfraEnvDownloadURLNotImplemented) IsCode(code int) bool {
	return code == 501
}

func (o *GetInfraEnvDownloadURLNotImplemented) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlNotImplemented  %+v", 501, o.Payload)
}

func (o *GetInfraEnvDownloadURLNotImplemented) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlNotImplemented  %+v", 501, o.Payload)
}

func (o *GetInfraEnvDownloadURLNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetInfraEnvDownloadURLNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInfraEnvDownloadURLServiceUnavailable creates a GetInfraEnvDownloadURLServiceUnavailable with default headers values
func NewGetInfraEnvDownloadURLServiceUnavailable() *GetInfraEnvDownloadURLServiceUnavailable {
	return &GetInfraEnvDownloadURLServiceUnavailable{}
}

/*
GetInfraEnvDownloadURLServiceUnavailable describes a response with status code 503, with default header values.

Unavailable.
*/
type GetInfraEnvDownloadURLServiceUnavailable struct {
	Payload *models.Error
}

// IsSuccess returns true when this get infra env download Url service unavailable response has a 2xx status code
func (o *GetInfraEnvDownloadURLServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get infra env download Url service unavailable response has a 3xx status code
func (o *GetInfraEnvDownloadURLServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get infra env download Url service unavailable response has a 4xx status code
func (o *GetInfraEnvDownloadURLServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get infra env download Url service unavailable response has a 5xx status code
func (o *GetInfraEnvDownloadURLServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get infra env download Url service unavailable response a status code equal to that given
func (o *GetInfraEnvDownloadURLServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetInfraEnvDownloadURLServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetInfraEnvDownloadURLServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/downloads/image-url][%d] getInfraEnvDownloadUrlServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetInfraEnvDownloadURLServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetInfraEnvDownloadURLServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
