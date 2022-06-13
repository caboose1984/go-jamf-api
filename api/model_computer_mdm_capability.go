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

// ComputerMdmCapability struct for ComputerMdmCapability
type ComputerMdmCapability struct {
	Capable *bool `json:"capable,omitempty"`
	CapableUsers []string `json:"capableUsers,omitempty"`
}

// NewComputerMdmCapability instantiates a new ComputerMdmCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerMdmCapability() *ComputerMdmCapability {
	this := ComputerMdmCapability{}
	return &this
}

// NewComputerMdmCapabilityWithDefaults instantiates a new ComputerMdmCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerMdmCapabilityWithDefaults() *ComputerMdmCapability {
	this := ComputerMdmCapability{}
	return &this
}

// GetCapable returns the Capable field value if set, zero value otherwise.
func (o *ComputerMdmCapability) GetCapable() bool {
	if o == nil || o.Capable == nil {
		var ret bool
		return ret
	}
	return *o.Capable
}

// GetCapableOk returns a tuple with the Capable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerMdmCapability) GetCapableOk() (*bool, bool) {
	if o == nil || o.Capable == nil {
		return nil, false
	}
	return o.Capable, true
}

// HasCapable returns a boolean if a field has been set.
func (o *ComputerMdmCapability) HasCapable() bool {
	if o != nil && o.Capable != nil {
		return true
	}

	return false
}

// SetCapable gets a reference to the given bool and assigns it to the Capable field.
func (o *ComputerMdmCapability) SetCapable(v bool) {
	o.Capable = &v
}

// GetCapableUsers returns the CapableUsers field value if set, zero value otherwise.
func (o *ComputerMdmCapability) GetCapableUsers() []string {
	if o == nil || o.CapableUsers == nil {
		var ret []string
		return ret
	}
	return o.CapableUsers
}

// GetCapableUsersOk returns a tuple with the CapableUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerMdmCapability) GetCapableUsersOk() ([]string, bool) {
	if o == nil || o.CapableUsers == nil {
		return nil, false
	}
	return o.CapableUsers, true
}

// HasCapableUsers returns a boolean if a field has been set.
func (o *ComputerMdmCapability) HasCapableUsers() bool {
	if o != nil && o.CapableUsers != nil {
		return true
	}

	return false
}

// SetCapableUsers gets a reference to the given []string and assigns it to the CapableUsers field.
func (o *ComputerMdmCapability) SetCapableUsers(v []string) {
	o.CapableUsers = v
}

func (o ComputerMdmCapability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Capable != nil {
		toSerialize["capable"] = o.Capable
	}
	if o.CapableUsers != nil {
		toSerialize["capableUsers"] = o.CapableUsers
	}
	return json.Marshal(toSerialize)
}

type NullableComputerMdmCapability struct {
	value *ComputerMdmCapability
	isSet bool
}

func (v NullableComputerMdmCapability) Get() *ComputerMdmCapability {
	return v.value
}

func (v *NullableComputerMdmCapability) Set(val *ComputerMdmCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerMdmCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerMdmCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerMdmCapability(val *ComputerMdmCapability) *NullableComputerMdmCapability {
	return &NullableComputerMdmCapability{value: val, isSet: true}
}

func (v NullableComputerMdmCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerMdmCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


