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

// checks if the SessionCandidateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionCandidateRequest{}

// SessionCandidateRequest Request to crate new remote session. Ultimately this allows connection between an admin and an end-user
type SessionCandidateRequest struct {
	// Device identifier
	DeviceId string `json:"deviceId"`
	// Device type
	DeviceType string `json:"deviceType"`
	// Session description. To be used for additional context on the reason of the session
	Description string `json:"description"`
}

// NewSessionCandidateRequest instantiates a new SessionCandidateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionCandidateRequest(deviceId string, deviceType string, description string) *SessionCandidateRequest {
	this := SessionCandidateRequest{}
	this.DeviceId = deviceId
	this.DeviceType = deviceType
	this.Description = description
	return &this
}

// NewSessionCandidateRequestWithDefaults instantiates a new SessionCandidateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionCandidateRequestWithDefaults() *SessionCandidateRequest {
	this := SessionCandidateRequest{}
	return &this
}

// GetDeviceId returns the DeviceId field value
func (o *SessionCandidateRequest) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *SessionCandidateRequest) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *SessionCandidateRequest) SetDeviceId(v string) {
	o.DeviceId = v
}

// GetDeviceType returns the DeviceType field value
func (o *SessionCandidateRequest) GetDeviceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *SessionCandidateRequest) GetDeviceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *SessionCandidateRequest) SetDeviceType(v string) {
	o.DeviceType = v
}

// GetDescription returns the Description field value
func (o *SessionCandidateRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SessionCandidateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SessionCandidateRequest) SetDescription(v string) {
	o.Description = v
}

func (o SessionCandidateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionCandidateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deviceId"] = o.DeviceId
	toSerialize["deviceType"] = o.DeviceType
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

type NullableSessionCandidateRequest struct {
	value *SessionCandidateRequest
	isSet bool
}

func (v NullableSessionCandidateRequest) Get() *SessionCandidateRequest {
	return v.value
}

func (v *NullableSessionCandidateRequest) Set(val *SessionCandidateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionCandidateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionCandidateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionCandidateRequest(val *SessionCandidateRequest) *NullableSessionCandidateRequest {
	return &NullableSessionCandidateRequest{value: val, isSet: true}
}

func (v NullableSessionCandidateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionCandidateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


