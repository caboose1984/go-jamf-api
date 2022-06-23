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


type UserSessionPreviewApi interface {

	/*
	UserGet Return all Jamf Pro user acounts 

	Return all Jamf Pro user acounts.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUserGetRequest
	*/
	UserGet(ctx context.Context) ApiUserGetRequest

	// UserGetExecute executes the request
	//  @return []Account
	UserGetExecute(r ApiUserGetRequest) ([]Account, *http.Response, error)

	/*
	UserUpdateSessionPost Update values in the User's current session 

	Updates values in the user's current session.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUserUpdateSessionPostRequest
	*/
	UserUpdateSessionPost(ctx context.Context) ApiUserUpdateSessionPostRequest

	// UserUpdateSessionPostExecute executes the request
	//  @return Session
	UserUpdateSessionPostExecute(r ApiUserUpdateSessionPostRequest) (*Session, *http.Response, error)
}

// UserSessionPreviewApiService UserSessionPreviewApi service
type UserSessionPreviewApiService service

type ApiUserGetRequest struct {
	ctx context.Context
	ApiService UserSessionPreviewApi
}

func (r ApiUserGetRequest) Execute() ([]Account, *http.Response, error) {
	return r.ApiService.UserGetExecute(r)
}

/*
UserGet Return all Jamf Pro user acounts 

Return all Jamf Pro user acounts.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserGetRequest
*/
func (a *UserSessionPreviewApiService) UserGet(ctx context.Context) ApiUserGetRequest {
	return ApiUserGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Account
func (a *UserSessionPreviewApiService) UserGetExecute(r ApiUserGetRequest) ([]Account, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Account
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserSessionPreviewApiService.UserGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user"

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

type ApiUserUpdateSessionPostRequest struct {
	ctx context.Context
	ApiService UserSessionPreviewApi
	session *Session
}

// Values to update in user&#39;s current session.
func (r ApiUserUpdateSessionPostRequest) Session(session Session) ApiUserUpdateSessionPostRequest {
	r.session = &session
	return r
}

func (r ApiUserUpdateSessionPostRequest) Execute() (*Session, *http.Response, error) {
	return r.ApiService.UserUpdateSessionPostExecute(r)
}

/*
UserUpdateSessionPost Update values in the User's current session 

Updates values in the user's current session.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserUpdateSessionPostRequest
*/
func (a *UserSessionPreviewApiService) UserUpdateSessionPost(ctx context.Context) ApiUserUpdateSessionPostRequest {
	return ApiUserUpdateSessionPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Session
func (a *UserSessionPreviewApiService) UserUpdateSessionPostExecute(r ApiUserUpdateSessionPostRequest) (*Session, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Session
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserSessionPreviewApiService.UserUpdateSessionPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/updateSession"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.session
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