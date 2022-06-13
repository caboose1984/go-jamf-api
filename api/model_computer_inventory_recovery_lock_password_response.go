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

// ComputerInventoryRecoveryLockPasswordResponse struct for ComputerInventoryRecoveryLockPasswordResponse
type ComputerInventoryRecoveryLockPasswordResponse struct {
	RecoveryLockPassword *string `json:"recoveryLockPassword,omitempty"`
}

// NewComputerInventoryRecoveryLockPasswordResponse instantiates a new ComputerInventoryRecoveryLockPasswordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerInventoryRecoveryLockPasswordResponse() *ComputerInventoryRecoveryLockPasswordResponse {
	this := ComputerInventoryRecoveryLockPasswordResponse{}
	return &this
}

// NewComputerInventoryRecoveryLockPasswordResponseWithDefaults instantiates a new ComputerInventoryRecoveryLockPasswordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerInventoryRecoveryLockPasswordResponseWithDefaults() *ComputerInventoryRecoveryLockPasswordResponse {
	this := ComputerInventoryRecoveryLockPasswordResponse{}
	return &this
}

// GetRecoveryLockPassword returns the RecoveryLockPassword field value if set, zero value otherwise.
func (o *ComputerInventoryRecoveryLockPasswordResponse) GetRecoveryLockPassword() string {
	if o == nil || o.RecoveryLockPassword == nil {
		var ret string
		return ret
	}
	return *o.RecoveryLockPassword
}

// GetRecoveryLockPasswordOk returns a tuple with the RecoveryLockPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerInventoryRecoveryLockPasswordResponse) GetRecoveryLockPasswordOk() (*string, bool) {
	if o == nil || o.RecoveryLockPassword == nil {
		return nil, false
	}
	return o.RecoveryLockPassword, true
}

// HasRecoveryLockPassword returns a boolean if a field has been set.
func (o *ComputerInventoryRecoveryLockPasswordResponse) HasRecoveryLockPassword() bool {
	if o != nil && o.RecoveryLockPassword != nil {
		return true
	}

	return false
}

// SetRecoveryLockPassword gets a reference to the given string and assigns it to the RecoveryLockPassword field.
func (o *ComputerInventoryRecoveryLockPasswordResponse) SetRecoveryLockPassword(v string) {
	o.RecoveryLockPassword = &v
}

func (o ComputerInventoryRecoveryLockPasswordResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecoveryLockPassword != nil {
		toSerialize["recoveryLockPassword"] = o.RecoveryLockPassword
	}
	return json.Marshal(toSerialize)
}

type NullableComputerInventoryRecoveryLockPasswordResponse struct {
	value *ComputerInventoryRecoveryLockPasswordResponse
	isSet bool
}

func (v NullableComputerInventoryRecoveryLockPasswordResponse) Get() *ComputerInventoryRecoveryLockPasswordResponse {
	return v.value
}

func (v *NullableComputerInventoryRecoveryLockPasswordResponse) Set(val *ComputerInventoryRecoveryLockPasswordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerInventoryRecoveryLockPasswordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerInventoryRecoveryLockPasswordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerInventoryRecoveryLockPasswordResponse(val *ComputerInventoryRecoveryLockPasswordResponse) *NullableComputerInventoryRecoveryLockPasswordResponse {
	return &NullableComputerInventoryRecoveryLockPasswordResponse{value: val, isSet: true}
}

func (v NullableComputerInventoryRecoveryLockPasswordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerInventoryRecoveryLockPasswordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


