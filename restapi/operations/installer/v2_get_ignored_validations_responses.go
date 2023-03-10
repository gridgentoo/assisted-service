// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2GetIgnoredValidationsOKCode is the HTTP code returned for type V2GetIgnoredValidationsOK
const V2GetIgnoredValidationsOKCode int = 200

/*
V2GetIgnoredValidationsOK Success.

swagger:response v2GetIgnoredValidationsOK
*/
type V2GetIgnoredValidationsOK struct {

	/*
	  In: Body
	*/
	Payload *models.IgnoredValidations `json:"body,omitempty"`
}

// NewV2GetIgnoredValidationsOK creates V2GetIgnoredValidationsOK with default headers values
func NewV2GetIgnoredValidationsOK() *V2GetIgnoredValidationsOK {

	return &V2GetIgnoredValidationsOK{}
}

// WithPayload adds the payload to the v2 get ignored validations o k response
func (o *V2GetIgnoredValidationsOK) WithPayload(payload *models.IgnoredValidations) *V2GetIgnoredValidationsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get ignored validations o k response
func (o *V2GetIgnoredValidationsOK) SetPayload(payload *models.IgnoredValidations) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetIgnoredValidationsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetIgnoredValidationsBadRequestCode is the HTTP code returned for type V2GetIgnoredValidationsBadRequest
const V2GetIgnoredValidationsBadRequestCode int = 400

/*
V2GetIgnoredValidationsBadRequest Error.

swagger:response v2GetIgnoredValidationsBadRequest
*/
type V2GetIgnoredValidationsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetIgnoredValidationsBadRequest creates V2GetIgnoredValidationsBadRequest with default headers values
func NewV2GetIgnoredValidationsBadRequest() *V2GetIgnoredValidationsBadRequest {

	return &V2GetIgnoredValidationsBadRequest{}
}

// WithPayload adds the payload to the v2 get ignored validations bad request response
func (o *V2GetIgnoredValidationsBadRequest) WithPayload(payload *models.Error) *V2GetIgnoredValidationsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get ignored validations bad request response
func (o *V2GetIgnoredValidationsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetIgnoredValidationsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetIgnoredValidationsNotFoundCode is the HTTP code returned for type V2GetIgnoredValidationsNotFound
const V2GetIgnoredValidationsNotFoundCode int = 404

/*
V2GetIgnoredValidationsNotFound Error.

swagger:response v2GetIgnoredValidationsNotFound
*/
type V2GetIgnoredValidationsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetIgnoredValidationsNotFound creates V2GetIgnoredValidationsNotFound with default headers values
func NewV2GetIgnoredValidationsNotFound() *V2GetIgnoredValidationsNotFound {

	return &V2GetIgnoredValidationsNotFound{}
}

// WithPayload adds the payload to the v2 get ignored validations not found response
func (o *V2GetIgnoredValidationsNotFound) WithPayload(payload *models.Error) *V2GetIgnoredValidationsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get ignored validations not found response
func (o *V2GetIgnoredValidationsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetIgnoredValidationsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
