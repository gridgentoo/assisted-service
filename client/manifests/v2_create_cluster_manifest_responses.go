// Code generated by go-swagger; DO NOT EDIT.

package manifests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2CreateClusterManifestReader is a Reader for the V2CreateClusterManifest structure.
type V2CreateClusterManifestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2CreateClusterManifestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewV2CreateClusterManifestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV2CreateClusterManifestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV2CreateClusterManifestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2CreateClusterManifestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2CreateClusterManifestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2CreateClusterManifestMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV2CreateClusterManifestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2CreateClusterManifestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2CreateClusterManifestCreated creates a V2CreateClusterManifestCreated with default headers values
func NewV2CreateClusterManifestCreated() *V2CreateClusterManifestCreated {
	return &V2CreateClusterManifestCreated{}
}

/*
V2CreateClusterManifestCreated describes a response with status code 201, with default header values.

Success.
*/
type V2CreateClusterManifestCreated struct {
	Payload *models.Manifest
}

// IsSuccess returns true when this v2 create cluster manifest created response has a 2xx status code
func (o *V2CreateClusterManifestCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v2 create cluster manifest created response has a 3xx status code
func (o *V2CreateClusterManifestCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 create cluster manifest created response has a 4xx status code
func (o *V2CreateClusterManifestCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 create cluster manifest created response has a 5xx status code
func (o *V2CreateClusterManifestCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 create cluster manifest created response a status code equal to that given
func (o *V2CreateClusterManifestCreated) IsCode(code int) bool {
	return code == 201
}

func (o *V2CreateClusterManifestCreated) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestCreated  %+v", 201, o.Payload)
}

func (o *V2CreateClusterManifestCreated) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestCreated  %+v", 201, o.Payload)
}

func (o *V2CreateClusterManifestCreated) GetPayload() *models.Manifest {
	return o.Payload
}

func (o *V2CreateClusterManifestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manifest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CreateClusterManifestBadRequest creates a V2CreateClusterManifestBadRequest with default headers values
func NewV2CreateClusterManifestBadRequest() *V2CreateClusterManifestBadRequest {
	return &V2CreateClusterManifestBadRequest{}
}

/*
V2CreateClusterManifestBadRequest describes a response with status code 400, with default header values.

Error.
*/
type V2CreateClusterManifestBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 create cluster manifest bad request response has a 2xx status code
func (o *V2CreateClusterManifestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 create cluster manifest bad request response has a 3xx status code
func (o *V2CreateClusterManifestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 create cluster manifest bad request response has a 4xx status code
func (o *V2CreateClusterManifestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 create cluster manifest bad request response has a 5xx status code
func (o *V2CreateClusterManifestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 create cluster manifest bad request response a status code equal to that given
func (o *V2CreateClusterManifestBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *V2CreateClusterManifestBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestBadRequest  %+v", 400, o.Payload)
}

func (o *V2CreateClusterManifestBadRequest) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestBadRequest  %+v", 400, o.Payload)
}

func (o *V2CreateClusterManifestBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2CreateClusterManifestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CreateClusterManifestUnauthorized creates a V2CreateClusterManifestUnauthorized with default headers values
func NewV2CreateClusterManifestUnauthorized() *V2CreateClusterManifestUnauthorized {
	return &V2CreateClusterManifestUnauthorized{}
}

/*
V2CreateClusterManifestUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type V2CreateClusterManifestUnauthorized struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 create cluster manifest unauthorized response has a 2xx status code
func (o *V2CreateClusterManifestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 create cluster manifest unauthorized response has a 3xx status code
func (o *V2CreateClusterManifestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 create cluster manifest unauthorized response has a 4xx status code
func (o *V2CreateClusterManifestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 create cluster manifest unauthorized response has a 5xx status code
func (o *V2CreateClusterManifestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 create cluster manifest unauthorized response a status code equal to that given
func (o *V2CreateClusterManifestUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *V2CreateClusterManifestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestUnauthorized  %+v", 401, o.Payload)
}

func (o *V2CreateClusterManifestUnauthorized) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestUnauthorized  %+v", 401, o.Payload)
}

func (o *V2CreateClusterManifestUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2CreateClusterManifestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CreateClusterManifestForbidden creates a V2CreateClusterManifestForbidden with default headers values
func NewV2CreateClusterManifestForbidden() *V2CreateClusterManifestForbidden {
	return &V2CreateClusterManifestForbidden{}
}

/*
V2CreateClusterManifestForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type V2CreateClusterManifestForbidden struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 create cluster manifest forbidden response has a 2xx status code
func (o *V2CreateClusterManifestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 create cluster manifest forbidden response has a 3xx status code
func (o *V2CreateClusterManifestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 create cluster manifest forbidden response has a 4xx status code
func (o *V2CreateClusterManifestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 create cluster manifest forbidden response has a 5xx status code
func (o *V2CreateClusterManifestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 create cluster manifest forbidden response a status code equal to that given
func (o *V2CreateClusterManifestForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *V2CreateClusterManifestForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestForbidden  %+v", 403, o.Payload)
}

func (o *V2CreateClusterManifestForbidden) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestForbidden  %+v", 403, o.Payload)
}

func (o *V2CreateClusterManifestForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2CreateClusterManifestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CreateClusterManifestNotFound creates a V2CreateClusterManifestNotFound with default headers values
func NewV2CreateClusterManifestNotFound() *V2CreateClusterManifestNotFound {
	return &V2CreateClusterManifestNotFound{}
}

/*
V2CreateClusterManifestNotFound describes a response with status code 404, with default header values.

Error.
*/
type V2CreateClusterManifestNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 create cluster manifest not found response has a 2xx status code
func (o *V2CreateClusterManifestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 create cluster manifest not found response has a 3xx status code
func (o *V2CreateClusterManifestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 create cluster manifest not found response has a 4xx status code
func (o *V2CreateClusterManifestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 create cluster manifest not found response has a 5xx status code
func (o *V2CreateClusterManifestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 create cluster manifest not found response a status code equal to that given
func (o *V2CreateClusterManifestNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *V2CreateClusterManifestNotFound) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestNotFound  %+v", 404, o.Payload)
}

func (o *V2CreateClusterManifestNotFound) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestNotFound  %+v", 404, o.Payload)
}

func (o *V2CreateClusterManifestNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2CreateClusterManifestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CreateClusterManifestMethodNotAllowed creates a V2CreateClusterManifestMethodNotAllowed with default headers values
func NewV2CreateClusterManifestMethodNotAllowed() *V2CreateClusterManifestMethodNotAllowed {
	return &V2CreateClusterManifestMethodNotAllowed{}
}

/*
V2CreateClusterManifestMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type V2CreateClusterManifestMethodNotAllowed struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 create cluster manifest method not allowed response has a 2xx status code
func (o *V2CreateClusterManifestMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 create cluster manifest method not allowed response has a 3xx status code
func (o *V2CreateClusterManifestMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 create cluster manifest method not allowed response has a 4xx status code
func (o *V2CreateClusterManifestMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 create cluster manifest method not allowed response has a 5xx status code
func (o *V2CreateClusterManifestMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 create cluster manifest method not allowed response a status code equal to that given
func (o *V2CreateClusterManifestMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *V2CreateClusterManifestMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2CreateClusterManifestMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2CreateClusterManifestMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2CreateClusterManifestMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CreateClusterManifestConflict creates a V2CreateClusterManifestConflict with default headers values
func NewV2CreateClusterManifestConflict() *V2CreateClusterManifestConflict {
	return &V2CreateClusterManifestConflict{}
}

/*
V2CreateClusterManifestConflict describes a response with status code 409, with default header values.

Error.
*/
type V2CreateClusterManifestConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 create cluster manifest conflict response has a 2xx status code
func (o *V2CreateClusterManifestConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 create cluster manifest conflict response has a 3xx status code
func (o *V2CreateClusterManifestConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 create cluster manifest conflict response has a 4xx status code
func (o *V2CreateClusterManifestConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 create cluster manifest conflict response has a 5xx status code
func (o *V2CreateClusterManifestConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 create cluster manifest conflict response a status code equal to that given
func (o *V2CreateClusterManifestConflict) IsCode(code int) bool {
	return code == 409
}

func (o *V2CreateClusterManifestConflict) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestConflict  %+v", 409, o.Payload)
}

func (o *V2CreateClusterManifestConflict) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestConflict  %+v", 409, o.Payload)
}

func (o *V2CreateClusterManifestConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2CreateClusterManifestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CreateClusterManifestInternalServerError creates a V2CreateClusterManifestInternalServerError with default headers values
func NewV2CreateClusterManifestInternalServerError() *V2CreateClusterManifestInternalServerError {
	return &V2CreateClusterManifestInternalServerError{}
}

/*
V2CreateClusterManifestInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type V2CreateClusterManifestInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 create cluster manifest internal server error response has a 2xx status code
func (o *V2CreateClusterManifestInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 create cluster manifest internal server error response has a 3xx status code
func (o *V2CreateClusterManifestInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 create cluster manifest internal server error response has a 4xx status code
func (o *V2CreateClusterManifestInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 create cluster manifest internal server error response has a 5xx status code
func (o *V2CreateClusterManifestInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 create cluster manifest internal server error response a status code equal to that given
func (o *V2CreateClusterManifestInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *V2CreateClusterManifestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestInternalServerError  %+v", 500, o.Payload)
}

func (o *V2CreateClusterManifestInternalServerError) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestInternalServerError  %+v", 500, o.Payload)
}

func (o *V2CreateClusterManifestInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2CreateClusterManifestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
