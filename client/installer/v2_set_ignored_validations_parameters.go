// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// NewV2SetIgnoredValidationsParams creates a new V2SetIgnoredValidationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV2SetIgnoredValidationsParams() *V2SetIgnoredValidationsParams {
	return &V2SetIgnoredValidationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV2SetIgnoredValidationsParamsWithTimeout creates a new V2SetIgnoredValidationsParams object
// with the ability to set a timeout on a request.
func NewV2SetIgnoredValidationsParamsWithTimeout(timeout time.Duration) *V2SetIgnoredValidationsParams {
	return &V2SetIgnoredValidationsParams{
		timeout: timeout,
	}
}

// NewV2SetIgnoredValidationsParamsWithContext creates a new V2SetIgnoredValidationsParams object
// with the ability to set a context for a request.
func NewV2SetIgnoredValidationsParamsWithContext(ctx context.Context) *V2SetIgnoredValidationsParams {
	return &V2SetIgnoredValidationsParams{
		Context: ctx,
	}
}

// NewV2SetIgnoredValidationsParamsWithHTTPClient creates a new V2SetIgnoredValidationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewV2SetIgnoredValidationsParamsWithHTTPClient(client *http.Client) *V2SetIgnoredValidationsParams {
	return &V2SetIgnoredValidationsParams{
		HTTPClient: client,
	}
}

/*
V2SetIgnoredValidationsParams contains all the parameters to send to the API endpoint

	for the v2 set ignored validations operation.

	Typically these are written to a http.Request.
*/
type V2SetIgnoredValidationsParams struct {

	/* ClusterID.

	   The cluster whose failing validations should be ignored according to this list.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	/* IgnoredValidations.

	   The validations to be ignored.
	*/
	IgnoredValidations *models.IgnoredValidations

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v2 set ignored validations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2SetIgnoredValidationsParams) WithDefaults() *V2SetIgnoredValidationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v2 set ignored validations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2SetIgnoredValidationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v2 set ignored validations params
func (o *V2SetIgnoredValidationsParams) WithTimeout(timeout time.Duration) *V2SetIgnoredValidationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 set ignored validations params
func (o *V2SetIgnoredValidationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 set ignored validations params
func (o *V2SetIgnoredValidationsParams) WithContext(ctx context.Context) *V2SetIgnoredValidationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 set ignored validations params
func (o *V2SetIgnoredValidationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 set ignored validations params
func (o *V2SetIgnoredValidationsParams) WithHTTPClient(client *http.Client) *V2SetIgnoredValidationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 set ignored validations params
func (o *V2SetIgnoredValidationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the v2 set ignored validations params
func (o *V2SetIgnoredValidationsParams) WithClusterID(clusterID strfmt.UUID) *V2SetIgnoredValidationsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the v2 set ignored validations params
func (o *V2SetIgnoredValidationsParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithIgnoredValidations adds the ignoredValidations to the v2 set ignored validations params
func (o *V2SetIgnoredValidationsParams) WithIgnoredValidations(ignoredValidations *models.IgnoredValidations) *V2SetIgnoredValidationsParams {
	o.SetIgnoredValidations(ignoredValidations)
	return o
}

// SetIgnoredValidations adds the ignoredValidations to the v2 set ignored validations params
func (o *V2SetIgnoredValidationsParams) SetIgnoredValidations(ignoredValidations *models.IgnoredValidations) {
	o.IgnoredValidations = ignoredValidations
}

// WriteToRequest writes these params to a swagger request
func (o *V2SetIgnoredValidationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}
	if o.IgnoredValidations != nil {
		if err := r.SetBodyParam(o.IgnoredValidations); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
