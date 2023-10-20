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

// checks if the AdvancedSearchSearchResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvancedSearchSearchResults{}

// AdvancedSearchSearchResults struct for AdvancedSearchSearchResults
type AdvancedSearchSearchResults struct {
	TotalCount *int32 `json:"totalCount,omitempty"`
	Results []AdvancedSearch `json:"results,omitempty"`
}

// NewAdvancedSearchSearchResults instantiates a new AdvancedSearchSearchResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvancedSearchSearchResults() *AdvancedSearchSearchResults {
	this := AdvancedSearchSearchResults{}
	return &this
}

// NewAdvancedSearchSearchResultsWithDefaults instantiates a new AdvancedSearchSearchResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvancedSearchSearchResultsWithDefaults() *AdvancedSearchSearchResults {
	this := AdvancedSearchSearchResults{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AdvancedSearchSearchResults) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedSearchSearchResults) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AdvancedSearchSearchResults) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *AdvancedSearchSearchResults) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *AdvancedSearchSearchResults) GetResults() []AdvancedSearch {
	if o == nil || IsNil(o.Results) {
		var ret []AdvancedSearch
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedSearchSearchResults) GetResultsOk() ([]AdvancedSearch, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *AdvancedSearchSearchResults) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []AdvancedSearch and assigns it to the Results field.
func (o *AdvancedSearchSearchResults) SetResults(v []AdvancedSearch) {
	o.Results = v
}

func (o AdvancedSearchSearchResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvancedSearchSearchResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableAdvancedSearchSearchResults struct {
	value *AdvancedSearchSearchResults
	isSet bool
}

func (v NullableAdvancedSearchSearchResults) Get() *AdvancedSearchSearchResults {
	return v.value
}

func (v *NullableAdvancedSearchSearchResults) Set(val *AdvancedSearchSearchResults) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvancedSearchSearchResults) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvancedSearchSearchResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvancedSearchSearchResults(val *AdvancedSearchSearchResults) *NullableAdvancedSearchSearchResults {
	return &NullableAdvancedSearchSearchResults{value: val, isSet: true}
}

func (v NullableAdvancedSearchSearchResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvancedSearchSearchResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


