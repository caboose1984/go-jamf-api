/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// MacosManagedSoftwareUpdatesApiService MacosManagedSoftwareUpdatesApi service
type MacosManagedSoftwareUpdatesApiService service

type ApiV1MacosManagedSoftwareUpdatesAvailableUpdatesGetRequest struct {
	ctx context.Context
	ApiService *MacosManagedSoftwareUpdatesApiService
}

func (r ApiV1MacosManagedSoftwareUpdatesAvailableUpdatesGetRequest) Execute() (*AvailableUpdates, *http.Response, error) {
	return r.ApiService.V1MacosManagedSoftwareUpdatesAvailableUpdatesGetExecute(r)
}

/*
V1MacosManagedSoftwareUpdatesAvailableUpdatesGet Retrieve available MacOs Managed Software Updates

Retrieves available MacOs Managed Software Updates

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1MacosManagedSoftwareUpdatesAvailableUpdatesGetRequest
*/
func (a *MacosManagedSoftwareUpdatesApiService) V1MacosManagedSoftwareUpdatesAvailableUpdatesGet(ctx context.Context) ApiV1MacosManagedSoftwareUpdatesAvailableUpdatesGetRequest {
	return ApiV1MacosManagedSoftwareUpdatesAvailableUpdatesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AvailableUpdates
func (a *MacosManagedSoftwareUpdatesApiService) V1MacosManagedSoftwareUpdatesAvailableUpdatesGetExecute(r ApiV1MacosManagedSoftwareUpdatesAvailableUpdatesGetRequest) (*AvailableUpdates, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AvailableUpdates
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MacosManagedSoftwareUpdatesApiService.V1MacosManagedSoftwareUpdatesAvailableUpdatesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/macos-managed-software-updates/available-updates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV1MacosManagedSoftwareUpdatesSendUpdatesPostRequest struct {
	ctx context.Context
	ApiService *MacosManagedSoftwareUpdatesApiService
	macOsManagedSoftwareUpdate *MacOsManagedSoftwareUpdate
}

// MacOs Managed Software Update to send
func (r ApiV1MacosManagedSoftwareUpdatesSendUpdatesPostRequest) MacOsManagedSoftwareUpdate(macOsManagedSoftwareUpdate MacOsManagedSoftwareUpdate) ApiV1MacosManagedSoftwareUpdatesSendUpdatesPostRequest {
	r.macOsManagedSoftwareUpdate = &macOsManagedSoftwareUpdate
	return r
}

func (r ApiV1MacosManagedSoftwareUpdatesSendUpdatesPostRequest) Execute() (*MacOsManagedSoftwareUpdateResponse, *http.Response, error) {
	return r.ApiService.V1MacosManagedSoftwareUpdatesSendUpdatesPostExecute(r)
}

/*
V1MacosManagedSoftwareUpdatesSendUpdatesPost Send MacOs Managed Software Updates

Sends MacOs Managed Software Updates

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1MacosManagedSoftwareUpdatesSendUpdatesPostRequest
*/
func (a *MacosManagedSoftwareUpdatesApiService) V1MacosManagedSoftwareUpdatesSendUpdatesPost(ctx context.Context) ApiV1MacosManagedSoftwareUpdatesSendUpdatesPostRequest {
	return ApiV1MacosManagedSoftwareUpdatesSendUpdatesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MacOsManagedSoftwareUpdateResponse
func (a *MacosManagedSoftwareUpdatesApiService) V1MacosManagedSoftwareUpdatesSendUpdatesPostExecute(r ApiV1MacosManagedSoftwareUpdatesSendUpdatesPostRequest) (*MacOsManagedSoftwareUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MacOsManagedSoftwareUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MacosManagedSoftwareUpdatesApiService.V1MacosManagedSoftwareUpdatesSendUpdatesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/macos-managed-software-updates/send-updates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.macOsManagedSoftwareUpdate == nil {
		return localVarReturnValue, nil, reportError("macOsManagedSoftwareUpdate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.macOsManagedSoftwareUpdate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
