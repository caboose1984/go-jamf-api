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

// GetEnrollmentCustomizationAllOf struct for GetEnrollmentCustomizationAllOf
type GetEnrollmentCustomizationAllOf struct {
	Id *int32 `json:"id,omitempty"`
}

// NewGetEnrollmentCustomizationAllOf instantiates a new GetEnrollmentCustomizationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEnrollmentCustomizationAllOf() *GetEnrollmentCustomizationAllOf {
	this := GetEnrollmentCustomizationAllOf{}
	return &this
}

// NewGetEnrollmentCustomizationAllOfWithDefaults instantiates a new GetEnrollmentCustomizationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEnrollmentCustomizationAllOfWithDefaults() *GetEnrollmentCustomizationAllOf {
	this := GetEnrollmentCustomizationAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetEnrollmentCustomizationAllOf) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationAllOf) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetEnrollmentCustomizationAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetEnrollmentCustomizationAllOf) SetId(v int32) {
	o.Id = &v
}

func (o GetEnrollmentCustomizationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGetEnrollmentCustomizationAllOf struct {
	value *GetEnrollmentCustomizationAllOf
	isSet bool
}

func (v NullableGetEnrollmentCustomizationAllOf) Get() *GetEnrollmentCustomizationAllOf {
	return v.value
}

func (v *NullableGetEnrollmentCustomizationAllOf) Set(val *GetEnrollmentCustomizationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEnrollmentCustomizationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEnrollmentCustomizationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEnrollmentCustomizationAllOf(val *GetEnrollmentCustomizationAllOf) *NullableGetEnrollmentCustomizationAllOf {
	return &NullableGetEnrollmentCustomizationAllOf{value: val, isSet: true}
}

func (v NullableGetEnrollmentCustomizationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEnrollmentCustomizationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


