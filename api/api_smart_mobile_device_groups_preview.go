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
	"strings"
)


type SmartMobileDeviceGroupsPreviewAPI interface {

	/*
	V1MobileDevicesIdRecalculateSmartGroupsPost Recalculate all smart groups for the given device id and then return count of smart groups that device fall into 

	Recalculates all smart groups for the given device id and then
returns the count of smart groups the device falls into


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id id of mobile device
	@return SmartMobileDeviceGroupsPreviewAPIV1MobileDevicesIdRecalculateSmartGroupsPostRequest
	*/
	V1MobileDevicesIdRecalculateSmartGroupsPost(ctx context.Context, id int32) SmartMobileDeviceGroupsPreviewAPIV1MobileDevicesIdRecalculateSmartGroupsPostRequest

	// V1MobileDevicesIdRecalculateSmartGroupsPostExecute executes the request
	//  @return RecalculationResults
	V1MobileDevicesIdRecalculateSmartGroupsPostExecute(r SmartMobileDeviceGroupsPreviewAPIV1MobileDevicesIdRecalculateSmartGroupsPostRequest) (*RecalculationResults, *http.Response, error)

	/*
	V1SmartMobileDeviceGroupsIdRecalculatePost Recalculate a smart group for the given id then return the ids for the devices in the smart group 

	recalculates a smart group for the given id and then
returns the ids for the devices in the smart group


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of smart group
	@return SmartMobileDeviceGroupsPreviewAPIV1SmartMobileDeviceGroupsIdRecalculatePostRequest
	*/
	V1SmartMobileDeviceGroupsIdRecalculatePost(ctx context.Context, id int32) SmartMobileDeviceGroupsPreviewAPIV1SmartMobileDeviceGroupsIdRecalculatePostRequest

	// V1SmartMobileDeviceGroupsIdRecalculatePostExecute executes the request
	//  @return RecalculationResults
	V1SmartMobileDeviceGroupsIdRecalculatePostExecute(r SmartMobileDeviceGroupsPreviewAPIV1SmartMobileDeviceGroupsIdRecalculatePostRequest) (*RecalculationResults, *http.Response, error)
}

// SmartMobileDeviceGroupsPreviewAPIService SmartMobileDeviceGroupsPreviewAPI service
type SmartMobileDeviceGroupsPreviewAPIService service

type SmartMobileDeviceGroupsPreviewAPIV1MobileDevicesIdRecalculateSmartGroupsPostRequest struct {
	ctx context.Context
	ApiService SmartMobileDeviceGroupsPreviewAPI
	id int32
}

func (r SmartMobileDeviceGroupsPreviewAPIV1MobileDevicesIdRecalculateSmartGroupsPostRequest) Execute() (*RecalculationResults, *http.Response, error) {
	return r.ApiService.V1MobileDevicesIdRecalculateSmartGroupsPostExecute(r)
}

/*
V1MobileDevicesIdRecalculateSmartGroupsPost Recalculate all smart groups for the given device id and then return count of smart groups that device fall into 

Recalculates all smart groups for the given device id and then
returns the count of smart groups the device falls into


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id id of mobile device
 @return SmartMobileDeviceGroupsPreviewAPIV1MobileDevicesIdRecalculateSmartGroupsPostRequest
*/
func (a *SmartMobileDeviceGroupsPreviewAPIService) V1MobileDevicesIdRecalculateSmartGroupsPost(ctx context.Context, id int32) SmartMobileDeviceGroupsPreviewAPIV1MobileDevicesIdRecalculateSmartGroupsPostRequest {
	return SmartMobileDeviceGroupsPreviewAPIV1MobileDevicesIdRecalculateSmartGroupsPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return RecalculationResults
func (a *SmartMobileDeviceGroupsPreviewAPIService) V1MobileDevicesIdRecalculateSmartGroupsPostExecute(r SmartMobileDeviceGroupsPreviewAPIV1MobileDevicesIdRecalculateSmartGroupsPostRequest) (*RecalculationResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RecalculationResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartMobileDeviceGroupsPreviewAPIService.V1MobileDevicesIdRecalculateSmartGroupsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/mobile-devices/{id}/recalculate-smart-groups"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

type SmartMobileDeviceGroupsPreviewAPIV1SmartMobileDeviceGroupsIdRecalculatePostRequest struct {
	ctx context.Context
	ApiService SmartMobileDeviceGroupsPreviewAPI
	id int32
}

func (r SmartMobileDeviceGroupsPreviewAPIV1SmartMobileDeviceGroupsIdRecalculatePostRequest) Execute() (*RecalculationResults, *http.Response, error) {
	return r.ApiService.V1SmartMobileDeviceGroupsIdRecalculatePostExecute(r)
}

/*
V1SmartMobileDeviceGroupsIdRecalculatePost Recalculate a smart group for the given id then return the ids for the devices in the smart group 

recalculates a smart group for the given id and then
returns the ids for the devices in the smart group


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of smart group
 @return SmartMobileDeviceGroupsPreviewAPIV1SmartMobileDeviceGroupsIdRecalculatePostRequest
*/
func (a *SmartMobileDeviceGroupsPreviewAPIService) V1SmartMobileDeviceGroupsIdRecalculatePost(ctx context.Context, id int32) SmartMobileDeviceGroupsPreviewAPIV1SmartMobileDeviceGroupsIdRecalculatePostRequest {
	return SmartMobileDeviceGroupsPreviewAPIV1SmartMobileDeviceGroupsIdRecalculatePostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return RecalculationResults
func (a *SmartMobileDeviceGroupsPreviewAPIService) V1SmartMobileDeviceGroupsIdRecalculatePostExecute(r SmartMobileDeviceGroupsPreviewAPIV1SmartMobileDeviceGroupsIdRecalculatePostRequest) (*RecalculationResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RecalculationResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartMobileDeviceGroupsPreviewAPIService.V1SmartMobileDeviceGroupsIdRecalculatePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/smart-mobile-device-groups/{id}/recalculate"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
