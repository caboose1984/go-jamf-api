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

// checks if the HistorySearchResultsV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistorySearchResultsV1{}

// HistorySearchResultsV1 struct for HistorySearchResultsV1
type HistorySearchResultsV1 struct {
	TotalCount *int32 `json:"totalCount,omitempty"`
	Results []ObjectHistoryV1 `json:"results,omitempty"`
}

// NewHistorySearchResultsV1 instantiates a new HistorySearchResultsV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistorySearchResultsV1() *HistorySearchResultsV1 {
	this := HistorySearchResultsV1{}
	return &this
}

// NewHistorySearchResultsV1WithDefaults instantiates a new HistorySearchResultsV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistorySearchResultsV1WithDefaults() *HistorySearchResultsV1 {
	this := HistorySearchResultsV1{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *HistorySearchResultsV1) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistorySearchResultsV1) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *HistorySearchResultsV1) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *HistorySearchResultsV1) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *HistorySearchResultsV1) GetResults() []ObjectHistoryV1 {
	if o == nil || IsNil(o.Results) {
		var ret []ObjectHistoryV1
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistorySearchResultsV1) GetResultsOk() ([]ObjectHistoryV1, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *HistorySearchResultsV1) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ObjectHistoryV1 and assigns it to the Results field.
func (o *HistorySearchResultsV1) SetResults(v []ObjectHistoryV1) {
	o.Results = v
}

func (o HistorySearchResultsV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistorySearchResultsV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableHistorySearchResultsV1 struct {
	value *HistorySearchResultsV1
	isSet bool
}

func (v NullableHistorySearchResultsV1) Get() *HistorySearchResultsV1 {
	return v.value
}

func (v *NullableHistorySearchResultsV1) Set(val *HistorySearchResultsV1) {
	v.value = val
	v.isSet = true
}

func (v NullableHistorySearchResultsV1) IsSet() bool {
	return v.isSet
}

func (v *NullableHistorySearchResultsV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistorySearchResultsV1(val *HistorySearchResultsV1) *NullableHistorySearchResultsV1 {
	return &NullableHistorySearchResultsV1{value: val, isSet: true}
}

func (v NullableHistorySearchResultsV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistorySearchResultsV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


