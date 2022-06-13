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

// InventoryPreloadCsvValidationSuccess struct for InventoryPreloadCsvValidationSuccess
type InventoryPreloadCsvValidationSuccess struct {
	RecordCount *int32 `json:"recordCount,omitempty"`
}

// NewInventoryPreloadCsvValidationSuccess instantiates a new InventoryPreloadCsvValidationSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryPreloadCsvValidationSuccess() *InventoryPreloadCsvValidationSuccess {
	this := InventoryPreloadCsvValidationSuccess{}
	return &this
}

// NewInventoryPreloadCsvValidationSuccessWithDefaults instantiates a new InventoryPreloadCsvValidationSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryPreloadCsvValidationSuccessWithDefaults() *InventoryPreloadCsvValidationSuccess {
	this := InventoryPreloadCsvValidationSuccess{}
	return &this
}

// GetRecordCount returns the RecordCount field value if set, zero value otherwise.
func (o *InventoryPreloadCsvValidationSuccess) GetRecordCount() int32 {
	if o == nil || o.RecordCount == nil {
		var ret int32
		return ret
	}
	return *o.RecordCount
}

// GetRecordCountOk returns a tuple with the RecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadCsvValidationSuccess) GetRecordCountOk() (*int32, bool) {
	if o == nil || o.RecordCount == nil {
		return nil, false
	}
	return o.RecordCount, true
}

// HasRecordCount returns a boolean if a field has been set.
func (o *InventoryPreloadCsvValidationSuccess) HasRecordCount() bool {
	if o != nil && o.RecordCount != nil {
		return true
	}

	return false
}

// SetRecordCount gets a reference to the given int32 and assigns it to the RecordCount field.
func (o *InventoryPreloadCsvValidationSuccess) SetRecordCount(v int32) {
	o.RecordCount = &v
}

func (o InventoryPreloadCsvValidationSuccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecordCount != nil {
		toSerialize["recordCount"] = o.RecordCount
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryPreloadCsvValidationSuccess struct {
	value *InventoryPreloadCsvValidationSuccess
	isSet bool
}

func (v NullableInventoryPreloadCsvValidationSuccess) Get() *InventoryPreloadCsvValidationSuccess {
	return v.value
}

func (v *NullableInventoryPreloadCsvValidationSuccess) Set(val *InventoryPreloadCsvValidationSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryPreloadCsvValidationSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryPreloadCsvValidationSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryPreloadCsvValidationSuccess(val *InventoryPreloadCsvValidationSuccess) *NullableInventoryPreloadCsvValidationSuccess {
	return &NullableInventoryPreloadCsvValidationSuccess{value: val, isSet: true}
}

func (v NullableInventoryPreloadCsvValidationSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryPreloadCsvValidationSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


