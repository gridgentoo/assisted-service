/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// LimitedSupportReasonTemplatesServer represents the interface the manages the 'limited_support_reason_templates' resource.
type LimitedSupportReasonTemplatesServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves the list of templates.
	List(ctx context.Context, request *LimitedSupportReasonTemplatesListServerRequest, response *LimitedSupportReasonTemplatesListServerResponse) error

	// LimitedSupportReasonTemplate returns the target 'limited_support_reason_template' server for the given identifier.
	//
	// Reference to the service that manages an specific template.
	LimitedSupportReasonTemplate(id string) LimitedSupportReasonTemplateServer
}

// LimitedSupportReasonTemplatesListServerRequest is the request for the 'list' method.
type LimitedSupportReasonTemplatesListServerRequest struct {
	page *int
	size *int
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *LimitedSupportReasonTemplatesListServerRequest) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
// Index of the requested page, where one corresponds to the first page.
func (r *LimitedSupportReasonTemplatesListServerRequest) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// Size returns the value of the 'size' parameter.
//
// Number of items contained in the returned page.
func (r *LimitedSupportReasonTemplatesListServerRequest) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
// Number of items contained in the returned page.
func (r *LimitedSupportReasonTemplatesListServerRequest) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// LimitedSupportReasonTemplatesListServerResponse is the response for the 'list' method.
type LimitedSupportReasonTemplatesListServerResponse struct {
	status int
	err    *errors.Error
	items  *LimitedSupportReasonTemplateList
	page   *int
	size   *int
	total  *int
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of template.
func (r *LimitedSupportReasonTemplatesListServerResponse) Items(value *LimitedSupportReasonTemplateList) *LimitedSupportReasonTemplatesListServerResponse {
	r.items = value
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *LimitedSupportReasonTemplatesListServerResponse) Page(value int) *LimitedSupportReasonTemplatesListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Number of items contained in the returned page.
func (r *LimitedSupportReasonTemplatesListServerResponse) Size(value int) *LimitedSupportReasonTemplatesListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection.
func (r *LimitedSupportReasonTemplatesListServerResponse) Total(value int) *LimitedSupportReasonTemplatesListServerResponse {
	r.total = &value
	return r
}

// Status sets the status code.
func (r *LimitedSupportReasonTemplatesListServerResponse) Status(value int) *LimitedSupportReasonTemplatesListServerResponse {
	r.status = value
	return r
}

// dispatchLimitedSupportReasonTemplates navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchLimitedSupportReasonTemplates(w http.ResponseWriter, r *http.Request, server LimitedSupportReasonTemplatesServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "GET":
			adaptLimitedSupportReasonTemplatesListRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	default:
		target := server.LimitedSupportReasonTemplate(segments[0])
		if target == nil {
			errors.SendNotFound(w, r)
			return
		}
		dispatchLimitedSupportReasonTemplate(w, r, target, segments[1:])
	}
}

// adaptLimitedSupportReasonTemplatesListRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptLimitedSupportReasonTemplatesListRequest(w http.ResponseWriter, r *http.Request, server LimitedSupportReasonTemplatesServer) {
	request := &LimitedSupportReasonTemplatesListServerRequest{}
	err := readLimitedSupportReasonTemplatesListRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &LimitedSupportReasonTemplatesListServerResponse{}
	response.status = 200
	err = server.List(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeLimitedSupportReasonTemplatesListResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
