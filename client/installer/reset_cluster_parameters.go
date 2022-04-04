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
)

// NewResetClusterParams creates a new ResetClusterParams object
// with the default values initialized.
func NewResetClusterParams() *ResetClusterParams {
	var ()
	return &ResetClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResetClusterParamsWithTimeout creates a new ResetClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResetClusterParamsWithTimeout(timeout time.Duration) *ResetClusterParams {
	var ()
	return &ResetClusterParams{

		timeout: timeout,
	}
}

// NewResetClusterParamsWithContext creates a new ResetClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewResetClusterParamsWithContext(ctx context.Context) *ResetClusterParams {
	var ()
	return &ResetClusterParams{

		Context: ctx,
	}
}

// NewResetClusterParamsWithHTTPClient creates a new ResetClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResetClusterParamsWithHTTPClient(client *http.Client) *ResetClusterParams {
	var ()
	return &ResetClusterParams{
		HTTPClient: client,
	}
}

/*ResetClusterParams contains all the parameters to send to the API endpoint
for the reset cluster operation typically these are written to a http.Request
*/
type ResetClusterParams struct {

	/*ClusterID
	  The cluster whose installation is to be reset.

	*/
	ClusterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reset cluster params
func (o *ResetClusterParams) WithTimeout(timeout time.Duration) *ResetClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset cluster params
func (o *ResetClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset cluster params
func (o *ResetClusterParams) WithContext(ctx context.Context) *ResetClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset cluster params
func (o *ResetClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset cluster params
func (o *ResetClusterParams) WithHTTPClient(client *http.Client) *ResetClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset cluster params
func (o *ResetClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the reset cluster params
func (o *ResetClusterParams) WithClusterID(clusterID strfmt.UUID) *ResetClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the reset cluster params
func (o *ResetClusterParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *ResetClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}