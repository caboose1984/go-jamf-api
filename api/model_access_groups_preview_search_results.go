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

// checks if the AccessGroupsPreviewSearchResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessGroupsPreviewSearchResults{}

// AccessGroupsPreviewSearchResults struct for AccessGroupsPreviewSearchResults
type AccessGroupsPreviewSearchResults struct {
	TotalCount *int32 `json:"totalCount,omitempty"`
	Results []EnrollmentAccessGroupPreview `json:"results,omitempty"`
}

// NewAccessGroupsPreviewSearchResults instantiates a new AccessGroupsPreviewSearchResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessGroupsPreviewSearchResults() *AccessGroupsPreviewSearchResults {
	this := AccessGroupsPreviewSearchResults{}
	return &this
}

// NewAccessGroupsPreviewSearchResultsWithDefaults instantiates a new AccessGroupsPreviewSearchResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessGroupsPreviewSearchResultsWithDefaults() *AccessGroupsPreviewSearchResults {
	this := AccessGroupsPreviewSearchResults{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AccessGroupsPreviewSearchResults) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessGroupsPreviewSearchResults) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AccessGroupsPreviewSearchResults) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *AccessGroupsPreviewSearchResults) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *AccessGroupsPreviewSearchResults) GetResults() []EnrollmentAccessGroupPreview {
	if o == nil || IsNil(o.Results) {
		var ret []EnrollmentAccessGroupPreview
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessGroupsPreviewSearchResults) GetResultsOk() ([]EnrollmentAccessGroupPreview, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *AccessGroupsPreviewSearchResults) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []EnrollmentAccessGroupPreview and assigns it to the Results field.
func (o *AccessGroupsPreviewSearchResults) SetResults(v []EnrollmentAccessGroupPreview) {
	o.Results = v
}

func (o AccessGroupsPreviewSearchResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessGroupsPreviewSearchResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableAccessGroupsPreviewSearchResults struct {
	value *AccessGroupsPreviewSearchResults
	isSet bool
}

func (v NullableAccessGroupsPreviewSearchResults) Get() *AccessGroupsPreviewSearchResults {
	return v.value
}

func (v *NullableAccessGroupsPreviewSearchResults) Set(val *AccessGroupsPreviewSearchResults) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessGroupsPreviewSearchResults) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessGroupsPreviewSearchResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessGroupsPreviewSearchResults(val *AccessGroupsPreviewSearchResults) *NullableAccessGroupsPreviewSearchResults {
	return &NullableAccessGroupsPreviewSearchResults{value: val, isSet: true}
}

func (v NullableAccessGroupsPreviewSearchResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessGroupsPreviewSearchResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


