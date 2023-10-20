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

// checks if the LapsSettingsResponseV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LapsSettingsResponseV2{}

// LapsSettingsResponseV2 struct for LapsSettingsResponseV2
type LapsSettingsResponseV2 struct {
	// When enabled, all appropriate computers will have the SetAutoAdminPassword command sent to them automatically.
	AutoDeployEnabled *bool `json:"autoDeployEnabled,omitempty"`
	// The amount of time in seconds that the local admin password will be rotated after viewing.
	PasswordRotationTime *int32 `json:"passwordRotationTime,omitempty"`
	// When enabled, all appropriate computers will automatically have their password expired and rotated after the configured autoRotateExpirationTime
	AutoRotateEnabled *bool `json:"autoRotateEnabled,omitempty"`
	// The amount of time in seconds that the local admin password will be rotated automatically if it is never viewed.
	AutoRotateExpirationTime *int32 `json:"autoRotateExpirationTime,omitempty"`
}

// NewLapsSettingsResponseV2 instantiates a new LapsSettingsResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLapsSettingsResponseV2() *LapsSettingsResponseV2 {
	this := LapsSettingsResponseV2{}
	return &this
}

// NewLapsSettingsResponseV2WithDefaults instantiates a new LapsSettingsResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLapsSettingsResponseV2WithDefaults() *LapsSettingsResponseV2 {
	this := LapsSettingsResponseV2{}
	return &this
}

// GetAutoDeployEnabled returns the AutoDeployEnabled field value if set, zero value otherwise.
func (o *LapsSettingsResponseV2) GetAutoDeployEnabled() bool {
	if o == nil || IsNil(o.AutoDeployEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoDeployEnabled
}

// GetAutoDeployEnabledOk returns a tuple with the AutoDeployEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LapsSettingsResponseV2) GetAutoDeployEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoDeployEnabled) {
		return nil, false
	}
	return o.AutoDeployEnabled, true
}

// HasAutoDeployEnabled returns a boolean if a field has been set.
func (o *LapsSettingsResponseV2) HasAutoDeployEnabled() bool {
	if o != nil && !IsNil(o.AutoDeployEnabled) {
		return true
	}

	return false
}

// SetAutoDeployEnabled gets a reference to the given bool and assigns it to the AutoDeployEnabled field.
func (o *LapsSettingsResponseV2) SetAutoDeployEnabled(v bool) {
	o.AutoDeployEnabled = &v
}

// GetPasswordRotationTime returns the PasswordRotationTime field value if set, zero value otherwise.
func (o *LapsSettingsResponseV2) GetPasswordRotationTime() int32 {
	if o == nil || IsNil(o.PasswordRotationTime) {
		var ret int32
		return ret
	}
	return *o.PasswordRotationTime
}

// GetPasswordRotationTimeOk returns a tuple with the PasswordRotationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LapsSettingsResponseV2) GetPasswordRotationTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.PasswordRotationTime) {
		return nil, false
	}
	return o.PasswordRotationTime, true
}

// HasPasswordRotationTime returns a boolean if a field has been set.
func (o *LapsSettingsResponseV2) HasPasswordRotationTime() bool {
	if o != nil && !IsNil(o.PasswordRotationTime) {
		return true
	}

	return false
}

// SetPasswordRotationTime gets a reference to the given int32 and assigns it to the PasswordRotationTime field.
func (o *LapsSettingsResponseV2) SetPasswordRotationTime(v int32) {
	o.PasswordRotationTime = &v
}

// GetAutoRotateEnabled returns the AutoRotateEnabled field value if set, zero value otherwise.
func (o *LapsSettingsResponseV2) GetAutoRotateEnabled() bool {
	if o == nil || IsNil(o.AutoRotateEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoRotateEnabled
}

// GetAutoRotateEnabledOk returns a tuple with the AutoRotateEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LapsSettingsResponseV2) GetAutoRotateEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoRotateEnabled) {
		return nil, false
	}
	return o.AutoRotateEnabled, true
}

// HasAutoRotateEnabled returns a boolean if a field has been set.
func (o *LapsSettingsResponseV2) HasAutoRotateEnabled() bool {
	if o != nil && !IsNil(o.AutoRotateEnabled) {
		return true
	}

	return false
}

// SetAutoRotateEnabled gets a reference to the given bool and assigns it to the AutoRotateEnabled field.
func (o *LapsSettingsResponseV2) SetAutoRotateEnabled(v bool) {
	o.AutoRotateEnabled = &v
}

// GetAutoRotateExpirationTime returns the AutoRotateExpirationTime field value if set, zero value otherwise.
func (o *LapsSettingsResponseV2) GetAutoRotateExpirationTime() int32 {
	if o == nil || IsNil(o.AutoRotateExpirationTime) {
		var ret int32
		return ret
	}
	return *o.AutoRotateExpirationTime
}

// GetAutoRotateExpirationTimeOk returns a tuple with the AutoRotateExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LapsSettingsResponseV2) GetAutoRotateExpirationTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.AutoRotateExpirationTime) {
		return nil, false
	}
	return o.AutoRotateExpirationTime, true
}

// HasAutoRotateExpirationTime returns a boolean if a field has been set.
func (o *LapsSettingsResponseV2) HasAutoRotateExpirationTime() bool {
	if o != nil && !IsNil(o.AutoRotateExpirationTime) {
		return true
	}

	return false
}

// SetAutoRotateExpirationTime gets a reference to the given int32 and assigns it to the AutoRotateExpirationTime field.
func (o *LapsSettingsResponseV2) SetAutoRotateExpirationTime(v int32) {
	o.AutoRotateExpirationTime = &v
}

func (o LapsSettingsResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LapsSettingsResponseV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoDeployEnabled) {
		toSerialize["autoDeployEnabled"] = o.AutoDeployEnabled
	}
	if !IsNil(o.PasswordRotationTime) {
		toSerialize["passwordRotationTime"] = o.PasswordRotationTime
	}
	if !IsNil(o.AutoRotateEnabled) {
		toSerialize["autoRotateEnabled"] = o.AutoRotateEnabled
	}
	if !IsNil(o.AutoRotateExpirationTime) {
		toSerialize["autoRotateExpirationTime"] = o.AutoRotateExpirationTime
	}
	return toSerialize, nil
}

type NullableLapsSettingsResponseV2 struct {
	value *LapsSettingsResponseV2
	isSet bool
}

func (v NullableLapsSettingsResponseV2) Get() *LapsSettingsResponseV2 {
	return v.value
}

func (v *NullableLapsSettingsResponseV2) Set(val *LapsSettingsResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableLapsSettingsResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableLapsSettingsResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLapsSettingsResponseV2(val *LapsSettingsResponseV2) *NullableLapsSettingsResponseV2 {
	return &NullableLapsSettingsResponseV2{value: val, isSet: true}
}

func (v NullableLapsSettingsResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLapsSettingsResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

