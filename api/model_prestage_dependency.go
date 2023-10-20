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

// checks if the PrestageDependency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrestageDependency{}

// PrestageDependency struct for PrestageDependency
type PrestageDependency struct {
	Name *string `json:"name,omitempty"`
	HumanReadableName *string `json:"humanReadableName,omitempty"`
	Hyperlink *string `json:"hyperlink,omitempty"`
}

// NewPrestageDependency instantiates a new PrestageDependency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrestageDependency() *PrestageDependency {
	this := PrestageDependency{}
	return &this
}

// NewPrestageDependencyWithDefaults instantiates a new PrestageDependency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrestageDependencyWithDefaults() *PrestageDependency {
	this := PrestageDependency{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PrestageDependency) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageDependency) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PrestageDependency) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PrestageDependency) SetName(v string) {
	o.Name = &v
}

// GetHumanReadableName returns the HumanReadableName field value if set, zero value otherwise.
func (o *PrestageDependency) GetHumanReadableName() string {
	if o == nil || IsNil(o.HumanReadableName) {
		var ret string
		return ret
	}
	return *o.HumanReadableName
}

// GetHumanReadableNameOk returns a tuple with the HumanReadableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageDependency) GetHumanReadableNameOk() (*string, bool) {
	if o == nil || IsNil(o.HumanReadableName) {
		return nil, false
	}
	return o.HumanReadableName, true
}

// HasHumanReadableName returns a boolean if a field has been set.
func (o *PrestageDependency) HasHumanReadableName() bool {
	if o != nil && !IsNil(o.HumanReadableName) {
		return true
	}

	return false
}

// SetHumanReadableName gets a reference to the given string and assigns it to the HumanReadableName field.
func (o *PrestageDependency) SetHumanReadableName(v string) {
	o.HumanReadableName = &v
}

// GetHyperlink returns the Hyperlink field value if set, zero value otherwise.
func (o *PrestageDependency) GetHyperlink() string {
	if o == nil || IsNil(o.Hyperlink) {
		var ret string
		return ret
	}
	return *o.Hyperlink
}

// GetHyperlinkOk returns a tuple with the Hyperlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageDependency) GetHyperlinkOk() (*string, bool) {
	if o == nil || IsNil(o.Hyperlink) {
		return nil, false
	}
	return o.Hyperlink, true
}

// HasHyperlink returns a boolean if a field has been set.
func (o *PrestageDependency) HasHyperlink() bool {
	if o != nil && !IsNil(o.Hyperlink) {
		return true
	}

	return false
}

// SetHyperlink gets a reference to the given string and assigns it to the Hyperlink field.
func (o *PrestageDependency) SetHyperlink(v string) {
	o.Hyperlink = &v
}

func (o PrestageDependency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrestageDependency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.HumanReadableName) {
		toSerialize["humanReadableName"] = o.HumanReadableName
	}
	if !IsNil(o.Hyperlink) {
		toSerialize["hyperlink"] = o.Hyperlink
	}
	return toSerialize, nil
}

type NullablePrestageDependency struct {
	value *PrestageDependency
	isSet bool
}

func (v NullablePrestageDependency) Get() *PrestageDependency {
	return v.value
}

func (v *NullablePrestageDependency) Set(val *PrestageDependency) {
	v.value = val
	v.isSet = true
}

func (v NullablePrestageDependency) IsSet() bool {
	return v.isSet
}

func (v *NullablePrestageDependency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrestageDependency(val *PrestageDependency) *NullablePrestageDependency {
	return &NullablePrestageDependency{value: val, isSet: true}
}

func (v NullablePrestageDependency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrestageDependency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


