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


type ApiAuthenticationAPI interface {

	/*
	AuthCurrentPost Get the authorization details associated with the current API token 

	Get the authorization details associated with the current API token for the users current site

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAuthenticationAPIAuthCurrentPostRequest

	Deprecated
	*/
	AuthCurrentPost(ctx context.Context) ApiAuthenticationAPIAuthCurrentPostRequest

	// AuthCurrentPostExecute executes the request
	//  @return CurrentAuthorization
	// Deprecated
	AuthCurrentPostExecute(r ApiAuthenticationAPIAuthCurrentPostRequest) (*CurrentAuthorization, *http.Response, error)

	/*
	AuthGet Get all the Authorization details associated with the current api 

	Get all the authorization details associated with the current api token

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAuthenticationAPIAuthGetRequest

	Deprecated
	*/
	AuthGet(ctx context.Context) ApiAuthenticationAPIAuthGetRequest

	// AuthGetExecute executes the request
	//  @return Authorization
	// Deprecated
	AuthGetExecute(r ApiAuthenticationAPIAuthGetRequest) (*Authorization, *http.Response, error)

	/*
	AuthInvalidateTokenPost Invalidate current token 

	Invalidates current token

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAuthenticationAPIAuthInvalidateTokenPostRequest

	Deprecated
	*/
	AuthInvalidateTokenPost(ctx context.Context) ApiAuthenticationAPIAuthInvalidateTokenPostRequest

	// AuthInvalidateTokenPostExecute executes the request
	// Deprecated
	AuthInvalidateTokenPostExecute(r ApiAuthenticationAPIAuthInvalidateTokenPostRequest) (*http.Response, error)

	/*
	AuthKeepAlivePost Invalidate existing token and generates new token 

	Invalidates existing token and generates new token with extended expiration based on existing token credentials.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAuthenticationAPIAuthKeepAlivePostRequest

	Deprecated
	*/
	AuthKeepAlivePost(ctx context.Context) ApiAuthenticationAPIAuthKeepAlivePostRequest

	// AuthKeepAlivePostExecute executes the request
	//  @return AuthToken
	// Deprecated
	AuthKeepAlivePostExecute(r ApiAuthenticationAPIAuthKeepAlivePostRequest) (*AuthToken, *http.Response, error)

	/*
	AuthTokensPost Create a token based on other authentication details (basic, etc.) 

	Create a token based on other authentication details (basic, etc.)


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAuthenticationAPIAuthTokensPostRequest

	Deprecated
	*/
	AuthTokensPost(ctx context.Context) ApiAuthenticationAPIAuthTokensPostRequest

	// AuthTokensPostExecute executes the request
	//  @return AuthToken
	// Deprecated
	AuthTokensPostExecute(r ApiAuthenticationAPIAuthTokensPostRequest) (*AuthToken, *http.Response, error)

	/*
	V1AuthGet Get all the Authorization details associated with the current api 

	Get all the authorization details associated with the current api token

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAuthenticationAPIV1AuthGetRequest
	*/
	V1AuthGet(ctx context.Context) ApiAuthenticationAPIV1AuthGetRequest

	// V1AuthGetExecute executes the request
	//  @return AuthorizationV1
	V1AuthGetExecute(r ApiAuthenticationAPIV1AuthGetRequest) (*AuthorizationV1, *http.Response, error)

	/*
	V1AuthInvalidateTokenPost Invalidate current token 

	Invalidates current token

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAuthenticationAPIV1AuthInvalidateTokenPostRequest
	*/
	V1AuthInvalidateTokenPost(ctx context.Context) ApiAuthenticationAPIV1AuthInvalidateTokenPostRequest

	// V1AuthInvalidateTokenPostExecute executes the request
	V1AuthInvalidateTokenPostExecute(r ApiAuthenticationAPIV1AuthInvalidateTokenPostRequest) (*http.Response, error)

	/*
	V1AuthKeepAlivePost Invalidate existing token and generates new token 

	Invalidates existing token and generates new token with extended expiration based on existing token credentials.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAuthenticationAPIV1AuthKeepAlivePostRequest
	*/
	V1AuthKeepAlivePost(ctx context.Context) ApiAuthenticationAPIV1AuthKeepAlivePostRequest

	// V1AuthKeepAlivePostExecute executes the request
	//  @return AuthTokenV1
	V1AuthKeepAlivePostExecute(r ApiAuthenticationAPIV1AuthKeepAlivePostRequest) (*AuthTokenV1, *http.Response, error)

	/*
	V1AuthTokenPost Create a token based on other authentication details (basic, etc.) 

	Create a token based on other authentication details (basic, etc.)


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAuthenticationAPIV1AuthTokenPostRequest
	*/
	V1AuthTokenPost(ctx context.Context) ApiAuthenticationAPIV1AuthTokenPostRequest

	// V1AuthTokenPostExecute executes the request
	//  @return AuthTokenV1
	V1AuthTokenPostExecute(r ApiAuthenticationAPIV1AuthTokenPostRequest) (*AuthTokenV1, *http.Response, error)
}

// ApiAuthenticationAPIService ApiAuthenticationAPI service
type ApiAuthenticationAPIService service

type ApiAuthenticationAPIAuthCurrentPostRequest struct {
	ctx context.Context
	ApiService ApiAuthenticationAPI
}

func (r ApiAuthenticationAPIAuthCurrentPostRequest) Execute() (*CurrentAuthorization, *http.Response, error) {
	return r.ApiService.AuthCurrentPostExecute(r)
}

/*
AuthCurrentPost Get the authorization details associated with the current API token 

Get the authorization details associated with the current API token for the users current site

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthenticationAPIAuthCurrentPostRequest

Deprecated
*/
func (a *ApiAuthenticationAPIService) AuthCurrentPost(ctx context.Context) ApiAuthenticationAPIAuthCurrentPostRequest {
	return ApiAuthenticationAPIAuthCurrentPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CurrentAuthorization
// Deprecated
func (a *ApiAuthenticationAPIService) AuthCurrentPostExecute(r ApiAuthenticationAPIAuthCurrentPostRequest) (*CurrentAuthorization, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CurrentAuthorization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiAuthenticationAPIService.AuthCurrentPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auth/current"

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiAuthenticationAPIAuthGetRequest struct {
	ctx context.Context
	ApiService ApiAuthenticationAPI
}

func (r ApiAuthenticationAPIAuthGetRequest) Execute() (*Authorization, *http.Response, error) {
	return r.ApiService.AuthGetExecute(r)
}

/*
AuthGet Get all the Authorization details associated with the current api 

Get all the authorization details associated with the current api token

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthenticationAPIAuthGetRequest

Deprecated
*/
func (a *ApiAuthenticationAPIService) AuthGet(ctx context.Context) ApiAuthenticationAPIAuthGetRequest {
	return ApiAuthenticationAPIAuthGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Authorization
// Deprecated
func (a *ApiAuthenticationAPIService) AuthGetExecute(r ApiAuthenticationAPIAuthGetRequest) (*Authorization, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Authorization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiAuthenticationAPIService.AuthGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auth"

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiAuthenticationAPIAuthInvalidateTokenPostRequest struct {
	ctx context.Context
	ApiService ApiAuthenticationAPI
}

func (r ApiAuthenticationAPIAuthInvalidateTokenPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.AuthInvalidateTokenPostExecute(r)
}

/*
AuthInvalidateTokenPost Invalidate current token 

Invalidates current token

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthenticationAPIAuthInvalidateTokenPostRequest

Deprecated
*/
func (a *ApiAuthenticationAPIService) AuthInvalidateTokenPost(ctx context.Context) ApiAuthenticationAPIAuthInvalidateTokenPostRequest {
	return ApiAuthenticationAPIAuthInvalidateTokenPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
// Deprecated
func (a *ApiAuthenticationAPIService) AuthInvalidateTokenPostExecute(r ApiAuthenticationAPIAuthInvalidateTokenPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiAuthenticationAPIService.AuthInvalidateTokenPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auth/invalidateToken"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAuthenticationAPIAuthKeepAlivePostRequest struct {
	ctx context.Context
	ApiService ApiAuthenticationAPI
}

func (r ApiAuthenticationAPIAuthKeepAlivePostRequest) Execute() (*AuthToken, *http.Response, error) {
	return r.ApiService.AuthKeepAlivePostExecute(r)
}

/*
AuthKeepAlivePost Invalidate existing token and generates new token 

Invalidates existing token and generates new token with extended expiration based on existing token credentials.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthenticationAPIAuthKeepAlivePostRequest

Deprecated
*/
func (a *ApiAuthenticationAPIService) AuthKeepAlivePost(ctx context.Context) ApiAuthenticationAPIAuthKeepAlivePostRequest {
	return ApiAuthenticationAPIAuthKeepAlivePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuthToken
// Deprecated
func (a *ApiAuthenticationAPIService) AuthKeepAlivePostExecute(r ApiAuthenticationAPIAuthKeepAlivePostRequest) (*AuthToken, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiAuthenticationAPIService.AuthKeepAlivePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auth/keepAlive"

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiAuthenticationAPIAuthTokensPostRequest struct {
	ctx context.Context
	ApiService ApiAuthenticationAPI
}

func (r ApiAuthenticationAPIAuthTokensPostRequest) Execute() (*AuthToken, *http.Response, error) {
	return r.ApiService.AuthTokensPostExecute(r)
}

/*
AuthTokensPost Create a token based on other authentication details (basic, etc.) 

Create a token based on other authentication details (basic, etc.)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthenticationAPIAuthTokensPostRequest

Deprecated
*/
func (a *ApiAuthenticationAPIService) AuthTokensPost(ctx context.Context) ApiAuthenticationAPIAuthTokensPostRequest {
	return ApiAuthenticationAPIAuthTokensPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuthToken
// Deprecated
func (a *ApiAuthenticationAPIService) AuthTokensPostExecute(r ApiAuthenticationAPIAuthTokensPostRequest) (*AuthToken, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiAuthenticationAPIService.AuthTokensPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auth/tokens"

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiAuthenticationAPIV1AuthGetRequest struct {
	ctx context.Context
	ApiService ApiAuthenticationAPI
}

func (r ApiAuthenticationAPIV1AuthGetRequest) Execute() (*AuthorizationV1, *http.Response, error) {
	return r.ApiService.V1AuthGetExecute(r)
}

/*
V1AuthGet Get all the Authorization details associated with the current api 

Get all the authorization details associated with the current api token

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthenticationAPIV1AuthGetRequest
*/
func (a *ApiAuthenticationAPIService) V1AuthGet(ctx context.Context) ApiAuthenticationAPIV1AuthGetRequest {
	return ApiAuthenticationAPIV1AuthGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuthorizationV1
func (a *ApiAuthenticationAPIService) V1AuthGetExecute(r ApiAuthenticationAPIV1AuthGetRequest) (*AuthorizationV1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthorizationV1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiAuthenticationAPIService.V1AuthGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/auth"

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiAuthenticationAPIV1AuthInvalidateTokenPostRequest struct {
	ctx context.Context
	ApiService ApiAuthenticationAPI
}

func (r ApiAuthenticationAPIV1AuthInvalidateTokenPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1AuthInvalidateTokenPostExecute(r)
}

/*
V1AuthInvalidateTokenPost Invalidate current token 

Invalidates current token

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthenticationAPIV1AuthInvalidateTokenPostRequest
*/
func (a *ApiAuthenticationAPIService) V1AuthInvalidateTokenPost(ctx context.Context) ApiAuthenticationAPIV1AuthInvalidateTokenPostRequest {
	return ApiAuthenticationAPIV1AuthInvalidateTokenPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ApiAuthenticationAPIService) V1AuthInvalidateTokenPostExecute(r ApiAuthenticationAPIV1AuthInvalidateTokenPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiAuthenticationAPIService.V1AuthInvalidateTokenPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/auth/invalidate-token"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAuthenticationAPIV1AuthKeepAlivePostRequest struct {
	ctx context.Context
	ApiService ApiAuthenticationAPI
}

func (r ApiAuthenticationAPIV1AuthKeepAlivePostRequest) Execute() (*AuthTokenV1, *http.Response, error) {
	return r.ApiService.V1AuthKeepAlivePostExecute(r)
}

/*
V1AuthKeepAlivePost Invalidate existing token and generates new token 

Invalidates existing token and generates new token with extended expiration based on existing token credentials.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthenticationAPIV1AuthKeepAlivePostRequest
*/
func (a *ApiAuthenticationAPIService) V1AuthKeepAlivePost(ctx context.Context) ApiAuthenticationAPIV1AuthKeepAlivePostRequest {
	return ApiAuthenticationAPIV1AuthKeepAlivePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuthTokenV1
func (a *ApiAuthenticationAPIService) V1AuthKeepAlivePostExecute(r ApiAuthenticationAPIV1AuthKeepAlivePostRequest) (*AuthTokenV1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthTokenV1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiAuthenticationAPIService.V1AuthKeepAlivePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/auth/keep-alive"

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiAuthenticationAPIV1AuthTokenPostRequest struct {
	ctx context.Context
	ApiService ApiAuthenticationAPI
}

func (r ApiAuthenticationAPIV1AuthTokenPostRequest) Execute() (*AuthTokenV1, *http.Response, error) {
	return r.ApiService.V1AuthTokenPostExecute(r)
}

/*
V1AuthTokenPost Create a token based on other authentication details (basic, etc.) 

Create a token based on other authentication details (basic, etc.)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthenticationAPIV1AuthTokenPostRequest
*/
func (a *ApiAuthenticationAPIService) V1AuthTokenPost(ctx context.Context) ApiAuthenticationAPIV1AuthTokenPostRequest {
	return ApiAuthenticationAPIV1AuthTokenPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuthTokenV1
func (a *ApiAuthenticationAPIService) V1AuthTokenPostExecute(r ApiAuthenticationAPIV1AuthTokenPostRequest) (*AuthTokenV1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthTokenV1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiAuthenticationAPIService.V1AuthTokenPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/auth/token"

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
