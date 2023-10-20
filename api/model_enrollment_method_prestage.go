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

// checks if the EnrollmentMethodPrestage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrollmentMethodPrestage{}

// EnrollmentMethodPrestage struct for EnrollmentMethodPrestage
type EnrollmentMethodPrestage struct {
	MobileDevicePrestageId *string `json:"mobileDevicePrestageId,omitempty"`
	ProfileName *string `json:"profileName,omitempty"`
}

// NewEnrollmentMethodPrestage instantiates a new EnrollmentMethodPrestage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentMethodPrestage() *EnrollmentMethodPrestage {
	this := EnrollmentMethodPrestage{}
	return &this
}

// NewEnrollmentMethodPrestageWithDefaults instantiates a new EnrollmentMethodPrestage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentMethodPrestageWithDefaults() *EnrollmentMethodPrestage {
	this := EnrollmentMethodPrestage{}
	return &this
}

// GetMobileDevicePrestageId returns the MobileDevicePrestageId field value if set, zero value otherwise.
func (o *EnrollmentMethodPrestage) GetMobileDevicePrestageId() string {
	if o == nil || IsNil(o.MobileDevicePrestageId) {
		var ret string
		return ret
	}
	return *o.MobileDevicePrestageId
}

// GetMobileDevicePrestageIdOk returns a tuple with the MobileDevicePrestageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentMethodPrestage) GetMobileDevicePrestageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MobileDevicePrestageId) {
		return nil, false
	}
	return o.MobileDevicePrestageId, true
}

// HasMobileDevicePrestageId returns a boolean if a field has been set.
func (o *EnrollmentMethodPrestage) HasMobileDevicePrestageId() bool {
	if o != nil && !IsNil(o.MobileDevicePrestageId) {
		return true
	}

	return false
}

// SetMobileDevicePrestageId gets a reference to the given string and assigns it to the MobileDevicePrestageId field.
func (o *EnrollmentMethodPrestage) SetMobileDevicePrestageId(v string) {
	o.MobileDevicePrestageId = &v
}

// GetProfileName returns the ProfileName field value if set, zero value otherwise.
func (o *EnrollmentMethodPrestage) GetProfileName() string {
	if o == nil || IsNil(o.ProfileName) {
		var ret string
		return ret
	}
	return *o.ProfileName
}

// GetProfileNameOk returns a tuple with the ProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentMethodPrestage) GetProfileNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileName) {
		return nil, false
	}
	return o.ProfileName, true
}

// HasProfileName returns a boolean if a field has been set.
func (o *EnrollmentMethodPrestage) HasProfileName() bool {
	if o != nil && !IsNil(o.ProfileName) {
		return true
	}

	return false
}

// SetProfileName gets a reference to the given string and assigns it to the ProfileName field.
func (o *EnrollmentMethodPrestage) SetProfileName(v string) {
	o.ProfileName = &v
}

func (o EnrollmentMethodPrestage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrollmentMethodPrestage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MobileDevicePrestageId) {
		toSerialize["mobileDevicePrestageId"] = o.MobileDevicePrestageId
	}
	if !IsNil(o.ProfileName) {
		toSerialize["profileName"] = o.ProfileName
	}
	return toSerialize, nil
}

type NullableEnrollmentMethodPrestage struct {
	value *EnrollmentMethodPrestage
	isSet bool
}

func (v NullableEnrollmentMethodPrestage) Get() *EnrollmentMethodPrestage {
	return v.value
}

func (v *NullableEnrollmentMethodPrestage) Set(val *EnrollmentMethodPrestage) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentMethodPrestage) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentMethodPrestage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentMethodPrestage(val *EnrollmentMethodPrestage) *NullableEnrollmentMethodPrestage {
	return &NullableEnrollmentMethodPrestage{value: val, isSet: true}
}

func (v NullableEnrollmentMethodPrestage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentMethodPrestage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


