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
	"io"
	"net/http"
	"net/url"
)


type JamfProInitializationAPI interface {

	/*
	V1SystemInitializeDatabaseConnectionPost Provide Database Password during startup 

	Provide database password during startup. Endpoint is accessible when database password was not configured and Jamf Pro server has not been initialized yet.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return JamfProInitializationAPIV1SystemInitializeDatabaseConnectionPostRequest
	*/
	V1SystemInitializeDatabaseConnectionPost(ctx context.Context) JamfProInitializationAPIV1SystemInitializeDatabaseConnectionPostRequest

	// V1SystemInitializeDatabaseConnectionPostExecute executes the request
	V1SystemInitializeDatabaseConnectionPostExecute(r JamfProInitializationAPIV1SystemInitializeDatabaseConnectionPostRequest) (*http.Response, error)

	/*
	V1SystemInitializePost Set up fresh installed Jamf Pro Server 

	Set up fresh installed Jamf Pro Server


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return JamfProInitializationAPIV1SystemInitializePostRequest
	*/
	V1SystemInitializePost(ctx context.Context) JamfProInitializationAPIV1SystemInitializePostRequest

	// V1SystemInitializePostExecute executes the request
	V1SystemInitializePostExecute(r JamfProInitializationAPIV1SystemInitializePostRequest) (*http.Response, error)
}

// JamfProInitializationAPIService JamfProInitializationAPI service
type JamfProInitializationAPIService service

type JamfProInitializationAPIV1SystemInitializeDatabaseConnectionPostRequest struct {
	ctx context.Context
	ApiService JamfProInitializationAPI
	databasePassword *DatabasePassword
}

func (r JamfProInitializationAPIV1SystemInitializeDatabaseConnectionPostRequest) DatabasePassword(databasePassword DatabasePassword) JamfProInitializationAPIV1SystemInitializeDatabaseConnectionPostRequest {
	r.databasePassword = &databasePassword
	return r
}

func (r JamfProInitializationAPIV1SystemInitializeDatabaseConnectionPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1SystemInitializeDatabaseConnectionPostExecute(r)
}

/*
V1SystemInitializeDatabaseConnectionPost Provide Database Password during startup 

Provide database password during startup. Endpoint is accessible when database password was not configured and Jamf Pro server has not been initialized yet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return JamfProInitializationAPIV1SystemInitializeDatabaseConnectionPostRequest
*/
func (a *JamfProInitializationAPIService) V1SystemInitializeDatabaseConnectionPost(ctx context.Context) JamfProInitializationAPIV1SystemInitializeDatabaseConnectionPostRequest {
	return JamfProInitializationAPIV1SystemInitializeDatabaseConnectionPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *JamfProInitializationAPIService) V1SystemInitializeDatabaseConnectionPostExecute(r JamfProInitializationAPIV1SystemInitializeDatabaseConnectionPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfProInitializationAPIService.V1SystemInitializeDatabaseConnectionPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/system/initialize-database-connection"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.databasePassword == nil {
		return nil, reportError("databasePassword is required and must be specified")
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
	localVarPostBody = r.databasePassword
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type JamfProInitializationAPIV1SystemInitializePostRequest struct {
	ctx context.Context
	ApiService JamfProInitializationAPI
	initializeV1 *InitializeV1
}

func (r JamfProInitializationAPIV1SystemInitializePostRequest) InitializeV1(initializeV1 InitializeV1) JamfProInitializationAPIV1SystemInitializePostRequest {
	r.initializeV1 = &initializeV1
	return r
}

func (r JamfProInitializationAPIV1SystemInitializePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1SystemInitializePostExecute(r)
}

/*
V1SystemInitializePost Set up fresh installed Jamf Pro Server 

Set up fresh installed Jamf Pro Server


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return JamfProInitializationAPIV1SystemInitializePostRequest
*/
func (a *JamfProInitializationAPIService) V1SystemInitializePost(ctx context.Context) JamfProInitializationAPIV1SystemInitializePostRequest {
	return JamfProInitializationAPIV1SystemInitializePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *JamfProInitializationAPIService) V1SystemInitializePostExecute(r JamfProInitializationAPIV1SystemInitializePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfProInitializationAPIService.V1SystemInitializePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/system/initialize"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.initializeV1 == nil {
		return nil, reportError("initializeV1 is required and must be specified")
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
	localVarPostBody = r.initializeV1
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
