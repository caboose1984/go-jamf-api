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

// LinkedConnectProfileSearchResults struct for LinkedConnectProfileSearchResults
type LinkedConnectProfileSearchResults struct {
	TotalCount *int32 `json:"totalCount,omitempty"`
	Results []LinkedConnectProfile `json:"results,omitempty"`
}

// NewLinkedConnectProfileSearchResults instantiates a new LinkedConnectProfileSearchResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedConnectProfileSearchResults() *LinkedConnectProfileSearchResults {
	this := LinkedConnectProfileSearchResults{}
	return &this
}

// NewLinkedConnectProfileSearchResultsWithDefaults instantiates a new LinkedConnectProfileSearchResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedConnectProfileSearchResultsWithDefaults() *LinkedConnectProfileSearchResults {
	this := LinkedConnectProfileSearchResults{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *LinkedConnectProfileSearchResults) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedConnectProfileSearchResults) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *LinkedConnectProfileSearchResults) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *LinkedConnectProfileSearchResults) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *LinkedConnectProfileSearchResults) GetResults() []LinkedConnectProfile {
	if o == nil || o.Results == nil {
		var ret []LinkedConnectProfile
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedConnectProfileSearchResults) GetResultsOk() ([]LinkedConnectProfile, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *LinkedConnectProfileSearchResults) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []LinkedConnectProfile and assigns it to the Results field.
func (o *LinkedConnectProfileSearchResults) SetResults(v []LinkedConnectProfile) {
	o.Results = v
}

func (o LinkedConnectProfileSearchResults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableLinkedConnectProfileSearchResults struct {
	value *LinkedConnectProfileSearchResults
	isSet bool
}

func (v NullableLinkedConnectProfileSearchResults) Get() *LinkedConnectProfileSearchResults {
	return v.value
}

func (v *NullableLinkedConnectProfileSearchResults) Set(val *LinkedConnectProfileSearchResults) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedConnectProfileSearchResults) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedConnectProfileSearchResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedConnectProfileSearchResults(val *LinkedConnectProfileSearchResults) *NullableLinkedConnectProfileSearchResults {
	return &NullableLinkedConnectProfileSearchResults{value: val, isSet: true}
}

func (v NullableLinkedConnectProfileSearchResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedConnectProfileSearchResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

