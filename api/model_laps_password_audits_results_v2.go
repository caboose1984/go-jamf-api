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

// checks if the LapsPasswordAuditsResultsV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LapsPasswordAuditsResultsV2{}

// LapsPasswordAuditsResultsV2 struct for LapsPasswordAuditsResultsV2
type LapsPasswordAuditsResultsV2 struct {
	TotalCount *int32 `json:"totalCount,omitempty"`
	Results []LapsPasswordAndAuditsV2 `json:"results,omitempty"`
}

// NewLapsPasswordAuditsResultsV2 instantiates a new LapsPasswordAuditsResultsV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLapsPasswordAuditsResultsV2() *LapsPasswordAuditsResultsV2 {
	this := LapsPasswordAuditsResultsV2{}
	return &this
}

// NewLapsPasswordAuditsResultsV2WithDefaults instantiates a new LapsPasswordAuditsResultsV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLapsPasswordAuditsResultsV2WithDefaults() *LapsPasswordAuditsResultsV2 {
	this := LapsPasswordAuditsResultsV2{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *LapsPasswordAuditsResultsV2) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LapsPasswordAuditsResultsV2) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *LapsPasswordAuditsResultsV2) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *LapsPasswordAuditsResultsV2) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *LapsPasswordAuditsResultsV2) GetResults() []LapsPasswordAndAuditsV2 {
	if o == nil || IsNil(o.Results) {
		var ret []LapsPasswordAndAuditsV2
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LapsPasswordAuditsResultsV2) GetResultsOk() ([]LapsPasswordAndAuditsV2, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *LapsPasswordAuditsResultsV2) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []LapsPasswordAndAuditsV2 and assigns it to the Results field.
func (o *LapsPasswordAuditsResultsV2) SetResults(v []LapsPasswordAndAuditsV2) {
	o.Results = v
}

func (o LapsPasswordAuditsResultsV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LapsPasswordAuditsResultsV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableLapsPasswordAuditsResultsV2 struct {
	value *LapsPasswordAuditsResultsV2
	isSet bool
}

func (v NullableLapsPasswordAuditsResultsV2) Get() *LapsPasswordAuditsResultsV2 {
	return v.value
}

func (v *NullableLapsPasswordAuditsResultsV2) Set(val *LapsPasswordAuditsResultsV2) {
	v.value = val
	v.isSet = true
}

func (v NullableLapsPasswordAuditsResultsV2) IsSet() bool {
	return v.isSet
}

func (v *NullableLapsPasswordAuditsResultsV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLapsPasswordAuditsResultsV2(val *LapsPasswordAuditsResultsV2) *NullableLapsPasswordAuditsResultsV2 {
	return &NullableLapsPasswordAuditsResultsV2{value: val, isSet: true}
}

func (v NullableLapsPasswordAuditsResultsV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLapsPasswordAuditsResultsV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

