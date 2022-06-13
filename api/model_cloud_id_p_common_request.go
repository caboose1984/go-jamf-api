/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CloudIdPCommonRequest A Cloud Identity Provider information for request
type CloudIdPCommonRequest struct {
	DisplayName string `json:"displayName"`
	ProviderName string `json:"providerName"`
}

// NewCloudIdPCommonRequest instantiates a new CloudIdPCommonRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudIdPCommonRequest(displayName string, providerName string) *CloudIdPCommonRequest {
	this := CloudIdPCommonRequest{}
	this.DisplayName = displayName
	this.ProviderName = providerName
	return &this
}

// NewCloudIdPCommonRequestWithDefaults instantiates a new CloudIdPCommonRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudIdPCommonRequestWithDefaults() *CloudIdPCommonRequest {
	this := CloudIdPCommonRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *CloudIdPCommonRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CloudIdPCommonRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CloudIdPCommonRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetProviderName returns the ProviderName field value
func (o *CloudIdPCommonRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *CloudIdPCommonRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *CloudIdPCommonRequest) SetProviderName(v string) {
	o.ProviderName = v
}

func (o CloudIdPCommonRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["providerName"] = o.ProviderName
	}
	return json.Marshal(toSerialize)
}

type NullableCloudIdPCommonRequest struct {
	value *CloudIdPCommonRequest
	isSet bool
}

func (v NullableCloudIdPCommonRequest) Get() *CloudIdPCommonRequest {
	return v.value
}

func (v *NullableCloudIdPCommonRequest) Set(val *CloudIdPCommonRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudIdPCommonRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudIdPCommonRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudIdPCommonRequest(val *CloudIdPCommonRequest) *NullableCloudIdPCommonRequest {
	return &NullableCloudIdPCommonRequest{value: val, isSet: true}
}

func (v NullableCloudIdPCommonRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudIdPCommonRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


