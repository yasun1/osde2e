/*
Copyright (c) 2019 Red Hat, Inc.

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

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// AccessTokenServer represents the interface the manages the 'access_token' resource.
type AccessTokenServer interface {

	// Post handles a request for the 'post' method.
	//
	// Returns access token generated from registries in docker format.
	Post(ctx context.Context, request *AccessTokenPostServerRequest, response *AccessTokenPostServerResponse) error
}

// AccessTokenPostServerRequest is the request for the 'post' method.
type AccessTokenPostServerRequest struct {
}

// AccessTokenPostServerResponse is the response for the 'post' method.
type AccessTokenPostServerResponse struct {
	status int
	err    *errors.Error
	body   *AccessToken
}

// Body sets the value of the 'body' parameter.
//
//
func (r *AccessTokenPostServerResponse) Body(value *AccessToken) *AccessTokenPostServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *AccessTokenPostServerResponse) Status(value int) *AccessTokenPostServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'post' method.
func (r *AccessTokenPostServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// dispatchAccessToken navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchAccessToken(w http.ResponseWriter, r *http.Request, server AccessTokenServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "POST":
			adaptAccessTokenPostRequest(w, r, server)
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	} else {
		switch segments[0] {
		default:
			errors.SendNotFound(w, r)
			return
		}
	}
}

// readAccessTokenPostRequest reads the given HTTP requests and translates it
// into an object of type AccessTokenPostServerRequest.
func readAccessTokenPostRequest(r *http.Request) (*AccessTokenPostServerRequest, error) {
	var err error
	result := new(AccessTokenPostServerRequest)
	return result, err
}

// writeAccessTokenPostResponse translates the given request object into an
// HTTP response.
func writeAccessTokenPostResponse(w http.ResponseWriter, r *AccessTokenPostServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}

// adaptAccessTokenPostRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptAccessTokenPostRequest(w http.ResponseWriter, r *http.Request, server AccessTokenServer) {
	request, err := readAccessTokenPostRequest(r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := new(AccessTokenPostServerResponse)
	response.status = 201
	err = server.Post(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeAccessTokenPostResponse(w, response)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
