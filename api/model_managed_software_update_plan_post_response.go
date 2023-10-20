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

// checks if the ManagedSoftwareUpdatePlanPostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedSoftwareUpdatePlanPostResponse{}

// ManagedSoftwareUpdatePlanPostResponse struct for ManagedSoftwareUpdatePlanPostResponse
type ManagedSoftwareUpdatePlanPostResponse struct {
	Plans []PlanDeviceResponse `json:"plans,omitempty"`
}

// NewManagedSoftwareUpdatePlanPostResponse instantiates a new ManagedSoftwareUpdatePlanPostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedSoftwareUpdatePlanPostResponse() *ManagedSoftwareUpdatePlanPostResponse {
	this := ManagedSoftwareUpdatePlanPostResponse{}
	return &this
}

// NewManagedSoftwareUpdatePlanPostResponseWithDefaults instantiates a new ManagedSoftwareUpdatePlanPostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedSoftwareUpdatePlanPostResponseWithDefaults() *ManagedSoftwareUpdatePlanPostResponse {
	this := ManagedSoftwareUpdatePlanPostResponse{}
	return &this
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *ManagedSoftwareUpdatePlanPostResponse) GetPlans() []PlanDeviceResponse {
	if o == nil || IsNil(o.Plans) {
		var ret []PlanDeviceResponse
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedSoftwareUpdatePlanPostResponse) GetPlansOk() ([]PlanDeviceResponse, bool) {
	if o == nil || IsNil(o.Plans) {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *ManagedSoftwareUpdatePlanPostResponse) HasPlans() bool {
	if o != nil && !IsNil(o.Plans) {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []PlanDeviceResponse and assigns it to the Plans field.
func (o *ManagedSoftwareUpdatePlanPostResponse) SetPlans(v []PlanDeviceResponse) {
	o.Plans = v
}

func (o ManagedSoftwareUpdatePlanPostResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedSoftwareUpdatePlanPostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plans) {
		toSerialize["plans"] = o.Plans
	}
	return toSerialize, nil
}

type NullableManagedSoftwareUpdatePlanPostResponse struct {
	value *ManagedSoftwareUpdatePlanPostResponse
	isSet bool
}

func (v NullableManagedSoftwareUpdatePlanPostResponse) Get() *ManagedSoftwareUpdatePlanPostResponse {
	return v.value
}

func (v *NullableManagedSoftwareUpdatePlanPostResponse) Set(val *ManagedSoftwareUpdatePlanPostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedSoftwareUpdatePlanPostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedSoftwareUpdatePlanPostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedSoftwareUpdatePlanPostResponse(val *ManagedSoftwareUpdatePlanPostResponse) *NullableManagedSoftwareUpdatePlanPostResponse {
	return &NullableManagedSoftwareUpdatePlanPostResponse{value: val, isSet: true}
}

func (v NullableManagedSoftwareUpdatePlanPostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedSoftwareUpdatePlanPostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


