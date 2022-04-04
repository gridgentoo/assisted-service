// Code generated by go-swagger; DO NOT EDIT.

package assisted_service_iso

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the assisted service iso client
type API interface {
	/*
	   CreateISOAndUploadToS3 Creates ISO for the user and uploads to S3.*/
	CreateISOAndUploadToS3(ctx context.Context, params *CreateISOAndUploadToS3Params) (*CreateISOAndUploadToS3Created, error)
	/*
	   DownloadISO Downloads the Assisted Service ISO.*/
	DownloadISO(ctx context.Context, params *DownloadISOParams, writer io.Writer) (*DownloadISOOK, error)
	/*
	   GetPresignedForAssistedServiceISO Retrieves a pre-signed S3 URL for downloading assisted-service ISO.*/
	GetPresignedForAssistedServiceISO(ctx context.Context, params *GetPresignedForAssistedServiceISOParams) (*GetPresignedForAssistedServiceISOOK, error)
}

// New creates a new assisted service iso API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for assisted service iso API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CreateISOAndUploadToS3 Creates ISO for the user and uploads to S3.
*/
func (a *Client) CreateISOAndUploadToS3(ctx context.Context, params *CreateISOAndUploadToS3Params) (*CreateISOAndUploadToS3Created, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateISOAndUploadToS3",
		Method:             "POST",
		PathPattern:        "/v1/assisted-service-iso",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateISOAndUploadToS3Reader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateISOAndUploadToS3Created), nil

}

/*
DownloadISO Downloads the Assisted Service ISO.
*/
func (a *Client) DownloadISO(ctx context.Context, params *DownloadISOParams, writer io.Writer) (*DownloadISOOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DownloadISO",
		Method:             "GET",
		PathPattern:        "/v1/assisted-service-iso/data",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DownloadISOReader{formats: a.formats, writer: writer},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DownloadISOOK), nil

}

/*
GetPresignedForAssistedServiceISO Retrieves a pre-signed S3 URL for downloading assisted-service ISO.
*/
func (a *Client) GetPresignedForAssistedServiceISO(ctx context.Context, params *GetPresignedForAssistedServiceISOParams) (*GetPresignedForAssistedServiceISOOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPresignedForAssistedServiceISO",
		Method:             "GET",
		PathPattern:        "/v1/assisted-service-iso/presigned",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPresignedForAssistedServiceISOReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPresignedForAssistedServiceISOOK), nil

}