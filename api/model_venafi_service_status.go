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

// VenafiServiceStatus struct for VenafiServiceStatus
type VenafiServiceStatus struct {
	Status *string `json:"status,omitempty"`
}

// NewVenafiServiceStatus instantiates a new VenafiServiceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVenafiServiceStatus() *VenafiServiceStatus {
	this := VenafiServiceStatus{}
	return &this
}

// NewVenafiServiceStatusWithDefaults instantiates a new VenafiServiceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVenafiServiceStatusWithDefaults() *VenafiServiceStatus {
	this := VenafiServiceStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VenafiServiceStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VenafiServiceStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VenafiServiceStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VenafiServiceStatus) SetStatus(v string) {
	o.Status = &v
}

func (o VenafiServiceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableVenafiServiceStatus struct {
	value *VenafiServiceStatus
	isSet bool
}

func (v NullableVenafiServiceStatus) Get() *VenafiServiceStatus {
	return v.value
}

func (v *NullableVenafiServiceStatus) Set(val *VenafiServiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVenafiServiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVenafiServiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVenafiServiceStatus(val *VenafiServiceStatus) *NullableVenafiServiceStatus {
	return &NullableVenafiServiceStatus{value: val, isSet: true}
}

func (v NullableVenafiServiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVenafiServiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


