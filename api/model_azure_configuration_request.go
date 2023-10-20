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

// checks if the AzureConfigurationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureConfigurationRequest{}

// AzureConfigurationRequest A Cloud Identity Provider Azure configuration for responses
type AzureConfigurationRequest struct {
	CloudIdPCommon CloudIdPCommonRequest `json:"cloudIdPCommon"`
	Server AzureServerConfigurationRequest `json:"server"`
}

// NewAzureConfigurationRequest instantiates a new AzureConfigurationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureConfigurationRequest(cloudIdPCommon CloudIdPCommonRequest, server AzureServerConfigurationRequest) *AzureConfigurationRequest {
	this := AzureConfigurationRequest{}
	this.CloudIdPCommon = cloudIdPCommon
	this.Server = server
	return &this
}

// NewAzureConfigurationRequestWithDefaults instantiates a new AzureConfigurationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureConfigurationRequestWithDefaults() *AzureConfigurationRequest {
	this := AzureConfigurationRequest{}
	return &this
}

// GetCloudIdPCommon returns the CloudIdPCommon field value
func (o *AzureConfigurationRequest) GetCloudIdPCommon() CloudIdPCommonRequest {
	if o == nil {
		var ret CloudIdPCommonRequest
		return ret
	}

	return o.CloudIdPCommon
}

// GetCloudIdPCommonOk returns a tuple with the CloudIdPCommon field value
// and a boolean to check if the value has been set.
func (o *AzureConfigurationRequest) GetCloudIdPCommonOk() (*CloudIdPCommonRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudIdPCommon, true
}

// SetCloudIdPCommon sets field value
func (o *AzureConfigurationRequest) SetCloudIdPCommon(v CloudIdPCommonRequest) {
	o.CloudIdPCommon = v
}

// GetServer returns the Server field value
func (o *AzureConfigurationRequest) GetServer() AzureServerConfigurationRequest {
	if o == nil {
		var ret AzureServerConfigurationRequest
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *AzureConfigurationRequest) GetServerOk() (*AzureServerConfigurationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *AzureConfigurationRequest) SetServer(v AzureServerConfigurationRequest) {
	o.Server = v
}

func (o AzureConfigurationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureConfigurationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudIdPCommon"] = o.CloudIdPCommon
	toSerialize["server"] = o.Server
	return toSerialize, nil
}

type NullableAzureConfigurationRequest struct {
	value *AzureConfigurationRequest
	isSet bool
}

func (v NullableAzureConfigurationRequest) Get() *AzureConfigurationRequest {
	return v.value
}

func (v *NullableAzureConfigurationRequest) Set(val *AzureConfigurationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureConfigurationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureConfigurationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureConfigurationRequest(val *AzureConfigurationRequest) *NullableAzureConfigurationRequest {
	return &NullableAzureConfigurationRequest{value: val, isSet: true}
}

func (v NullableAzureConfigurationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureConfigurationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


