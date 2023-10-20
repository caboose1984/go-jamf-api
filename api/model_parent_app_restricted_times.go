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

// checks if the ParentAppRestrictedTimes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParentAppRestrictedTimes{}

// ParentAppRestrictedTimes struct for ParentAppRestrictedTimes
type ParentAppRestrictedTimes struct {
	Key *DayOfWeek `json:"key,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ParentAppRestrictedTimes ParentAppRestrictedTimes

// NewParentAppRestrictedTimes instantiates a new ParentAppRestrictedTimes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParentAppRestrictedTimes() *ParentAppRestrictedTimes {
	this := ParentAppRestrictedTimes{}
	return &this
}

// NewParentAppRestrictedTimesWithDefaults instantiates a new ParentAppRestrictedTimes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParentAppRestrictedTimesWithDefaults() *ParentAppRestrictedTimes {
	this := ParentAppRestrictedTimes{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ParentAppRestrictedTimes) GetKey() DayOfWeek {
	if o == nil || IsNil(o.Key) {
		var ret DayOfWeek
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentAppRestrictedTimes) GetKeyOk() (*DayOfWeek, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ParentAppRestrictedTimes) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given DayOfWeek and assigns it to the Key field.
func (o *ParentAppRestrictedTimes) SetKey(v DayOfWeek) {
	o.Key = &v
}

func (o ParentAppRestrictedTimes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParentAppRestrictedTimes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ParentAppRestrictedTimes) UnmarshalJSON(bytes []byte) (err error) {
	varParentAppRestrictedTimes := _ParentAppRestrictedTimes{}

	err = json.Unmarshal(bytes, &varParentAppRestrictedTimes)

	if err != nil {
		return err
	}

	*o = ParentAppRestrictedTimes(varParentAppRestrictedTimes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableParentAppRestrictedTimes struct {
	value *ParentAppRestrictedTimes
	isSet bool
}

func (v NullableParentAppRestrictedTimes) Get() *ParentAppRestrictedTimes {
	return v.value
}

func (v *NullableParentAppRestrictedTimes) Set(val *ParentAppRestrictedTimes) {
	v.value = val
	v.isSet = true
}

func (v NullableParentAppRestrictedTimes) IsSet() bool {
	return v.isSet
}

func (v *NullableParentAppRestrictedTimes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParentAppRestrictedTimes(val *ParentAppRestrictedTimes) *NullableParentAppRestrictedTimes {
	return &NullableParentAppRestrictedTimes{value: val, isSet: true}
}

func (v NullableParentAppRestrictedTimes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParentAppRestrictedTimes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


