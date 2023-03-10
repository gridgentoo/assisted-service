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

// V2CompleteInstallationReader is a Reader for the V2CompleteInstallation structure.
type V2CompleteInstallationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2CompleteInstallationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewV2CompleteInstallationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV2CompleteInstallationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2CompleteInstallationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2CompleteInstallationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2CompleteInstallationMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV2CompleteInstallationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2CompleteInstallationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewV2CompleteInstallationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2CompleteInstallationAccepted creates a V2CompleteInstallationAccepted with default headers values
func NewV2CompleteInstallationAccepted() *V2CompleteInstallationAccepted {
	return &V2CompleteInstallationAccepted{}
}

/*
V2CompleteInstallationAccepted describes a response with status code 202, with default header values.

Success.
*/
type V2CompleteInstallationAccepted struct {
	Payload *models.Cluster
}

// IsSuccess returns true when this v2 complete installation accepted response has a 2xx status code
func (o *V2CompleteInstallationAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v2 complete installation accepted response has a 3xx status code
func (o *V2CompleteInstallationAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 complete installation accepted response has a 4xx status code
func (o *V2CompleteInstallationAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 complete installation accepted response has a 5xx status code
func (o *V2CompleteInstallationAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 complete installation accepted response a status code equal to that given
func (o *V2CompleteInstallationAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *V2CompleteInstallationAccepted) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/complete-installation][%d] v2CompleteInstallationAccepted  %+v", 202, o.Payload)
}

func (o *V2CompleteInstallationAccepted) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/complete-installation][%d] v2CompleteInstallationAccepted  %+v", 202, o.Payload)
}

func (o *V2CompleteInstallationAccepted) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *V2CompleteInstallationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CompleteInstallationUnauthorized creates a V2CompleteInstallationUnauthorized with default headers values
func NewV2CompleteInstallationUnauthorized() *V2CompleteInstallationUnauthorized {
	return &V2CompleteInstallationUnauthorized{}
}

/*
V2CompleteInstallationUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type V2CompleteInstallationUnauthorized struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 complete installation unauthorized response has a 2xx status code
func (o *V2CompleteInstallationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 complete installation unauthorized response has a 3xx status code
func (o *V2CompleteInstallationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 complete installation unauthorized response has a 4xx status code
func (o *V2CompleteInstallationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 complete installation unauthorized response has a 5xx status code
func (o *V2CompleteInstallationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 complete installation unauthorized response a status code equal to that given
func (o *V2CompleteInstallationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *V2CompleteInstallationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/complete-installation][%d] v2CompleteInstallationUnauthorized  %+v", 401, o.Payload)
}

func (o *V2CompleteInstallationUnauthorized) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/complete-installation][%d] v2CompleteInstallationUnauthorized  %+v", 401, o.Payload)
}

func (o *V2CompleteInstallationUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2CompleteInstallationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CompleteInstallationForbidden creates a V2CompleteInstallationForbidden with default headers values
func NewV2CompleteInstallationForbidden() *V2CompleteInstallationForbidden {
	return &V2CompleteInstallationForbidden{}
}

/*
V2CompleteInstallationForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type V2CompleteInstallationForbidden struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 complete installation forbidden response has a 2xx status code
func (o *V2CompleteInstallationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 complete installation forbidden response has a 3xx status code
func (o *V2CompleteInstallationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 complete installation forbidden response has a 4xx status code
func (o *V2CompleteInstallationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 complete installation forbidden response has a 5xx status code
func (o *V2CompleteInstallationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 complete installation forbidden response a status code equal to that given
func (o *V2CompleteInstallationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *V2CompleteInstallationForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/complete-installation][%d] v2CompleteInstallationForbidden  %+v", 403, o.Payload)
}

func (o *V2CompleteInstallationForbidden) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/complete-installation][%d] v2CompleteInstallationForbidden  %+v", 403, o.Payload)
}

func (o *V2CompleteInstallationForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2CompleteInstallationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CompleteInstallationNotFound creates a V2CompleteInstallationNotFound with default headers values
func NewV2CompleteInstallationNotFound() *V2CompleteInstallationNotFound {
	return &V2CompleteInstallationNotFound{}
}

/*
V2CompleteInstallationNotFound describes a response with status code 404, with default header values.

Error.
*/
type V2CompleteInstallationNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 complete installation not found response has a 2xx status code
func (o *V2CompleteInstallationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 complete installation not found response has a 3xx status code
func (o *V2CompleteInstallationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 complete installation not found response has a 4xx status code
func (o *V2CompleteInstallationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 complete installation not found response has a 5xx status code
func (o *V2CompleteInstallationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 complete installation not found response a status code equal to that given
func (o *V2CompleteInstallationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *V2CompleteInstallationNotFound) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/complete-installation][%d] v2CompleteInstallationNotFound  %+v", 404, o.Payload)
}

func (o *V2CompleteInstallationNotFound) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/complete-installation][%d] v2CompleteInstallationNotFound  %+v", 404, o.Payload)
}

func (o *V2CompleteInstallationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2CompleteInstallationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CompleteInstallationMethodNotAllowed creates a V2CompleteInstallationMethodNotAllowed with default headers values
func NewV2CompleteInstallationMethodNotAllowed() *V2CompleteInstallationMethodNotAllowed {
	return &V2CompleteInstallationMethodNotAllowed{}
}

/*
V2CompleteInstallationMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type V2CompleteInstallationMethodNotAllowed struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 complete installation method not allowed response has a 2xx status code
func (o *V2CompleteInstallationMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 complete installation method not allowed response has a 3xx status code
func (o *V2CompleteInstallationMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 complete installation method not allowed response has a 4xx status code
func (o *V2CompleteInstallationMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 complete installation method not allowed response has a 5xx status code
func (o *V2CompleteInstallationMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 complete installation method not allowed response a status code equal to that given
func (o *V2CompleteInstallationMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *V2CompleteInstallationMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/complete-installation][%d] v2CompleteInstallationMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2CompleteInstallationMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/complete-installation][%d] v2CompleteInstallationMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2CompleteInstallationMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2CompleteInstallationMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CompleteInstallationConflict creates a V2CompleteInstallationConflict with default headers values
func NewV2CompleteInstallationConflict() *V2CompleteInstallationConflict {
	return &V2CompleteInstallationConflict{}
}

/*
V2CompleteInstallationConflict describes a response with status code 409, with default header values.

Error.
*/
type V2CompleteInstallationConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 complete installation conflict response has a 2xx status code
func (o *V2CompleteInstallationConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 complete installation conflict response has a 3xx status code
func (o *V2CompleteInstallationConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 complete installation conflict response has a 4xx status code
func (o *V2CompleteInstallationConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 complete installation conflict response has a 5xx status code
func (o *V2CompleteInstallationConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 complete installation conflict response a status code equal to that given
func (o *V2CompleteInstallationConflict) IsCode(code int) bool {
	return code == 409
}

func (o *V2CompleteInstallationConflict) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/complete-installation][%d] v2CompleteInstallationConflict  %+v", 409, o.Payload)
}

func (o *V2CompleteInstallationConflict) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/complete-installation][%d] v2CompleteInstallationConflict  %+v", 409, o.Payload)
}

func (o *V2CompleteInstallationConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2CompleteInstallationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CompleteInstallationInternalServerError creates a V2CompleteInstallationInternalServerError with default headers values
func NewV2CompleteInstallationInternalServerError() *V2CompleteInstallationInternalServerError {
	return &V2CompleteInstallationInternalServerError{}
}

/*
V2CompleteInstallationInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type V2CompleteInstallationInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 complete installation internal server error response has a 2xx status code
func (o *V2CompleteInstallationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 complete installation internal server error response has a 3xx status code
func (o *V2CompleteInstallationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 complete installation internal server error response has a 4xx status code
func (o *V2CompleteInstallationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 complete installation internal server error response has a 5xx status code
func (o *V2CompleteInstallationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 complete installation internal server error response a status code equal to that given
func (o *V2CompleteInstallationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *V2CompleteInstallationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/complete-installation][%d] v2CompleteInstallationInternalServerError  %+v", 500, o.Payload)
}

func (o *V2CompleteInstallationInternalServerError) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/complete-installation][%d] v2CompleteInstallationInternalServerError  %+v", 500, o.Payload)
}

func (o *V2CompleteInstallationInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2CompleteInstallationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CompleteInstallationServiceUnavailable creates a V2CompleteInstallationServiceUnavailable with default headers values
func NewV2CompleteInstallationServiceUnavailable() *V2CompleteInstallationServiceUnavailable {
	return &V2CompleteInstallationServiceUnavailable{}
}

/*
V2CompleteInstallationServiceUnavailable describes a response with status code 503, with default header values.

Unavailable.
*/
type V2CompleteInstallationServiceUnavailable struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 complete installation service unavailable response has a 2xx status code
func (o *V2CompleteInstallationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 complete installation service unavailable response has a 3xx status code
func (o *V2CompleteInstallationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 complete installation service unavailable response has a 4xx status code
func (o *V2CompleteInstallationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 complete installation service unavailable response has a 5xx status code
func (o *V2CompleteInstallationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 complete installation service unavailable response a status code equal to that given
func (o *V2CompleteInstallationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *V2CompleteInstallationServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/complete-installation][%d] v2CompleteInstallationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *V2CompleteInstallationServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/complete-installation][%d] v2CompleteInstallationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *V2CompleteInstallationServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2CompleteInstallationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
