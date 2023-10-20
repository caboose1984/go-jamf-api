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

// checks if the ComputerContentCachingCacheDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComputerContentCachingCacheDetail{}

// ComputerContentCachingCacheDetail struct for ComputerContentCachingCacheDetail
type ComputerContentCachingCacheDetail struct {
	ComputerContentCachingCacheDetailsId *string `json:"computerContentCachingCacheDetailsId,omitempty"`
	CategoryName *string `json:"categoryName,omitempty"`
	DiskSpaceBytesUsed *int64 `json:"diskSpaceBytesUsed,omitempty"`
}

// NewComputerContentCachingCacheDetail instantiates a new ComputerContentCachingCacheDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerContentCachingCacheDetail() *ComputerContentCachingCacheDetail {
	this := ComputerContentCachingCacheDetail{}
	return &this
}

// NewComputerContentCachingCacheDetailWithDefaults instantiates a new ComputerContentCachingCacheDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerContentCachingCacheDetailWithDefaults() *ComputerContentCachingCacheDetail {
	this := ComputerContentCachingCacheDetail{}
	return &this
}

// GetComputerContentCachingCacheDetailsId returns the ComputerContentCachingCacheDetailsId field value if set, zero value otherwise.
func (o *ComputerContentCachingCacheDetail) GetComputerContentCachingCacheDetailsId() string {
	if o == nil || IsNil(o.ComputerContentCachingCacheDetailsId) {
		var ret string
		return ret
	}
	return *o.ComputerContentCachingCacheDetailsId
}

// GetComputerContentCachingCacheDetailsIdOk returns a tuple with the ComputerContentCachingCacheDetailsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerContentCachingCacheDetail) GetComputerContentCachingCacheDetailsIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComputerContentCachingCacheDetailsId) {
		return nil, false
	}
	return o.ComputerContentCachingCacheDetailsId, true
}

// HasComputerContentCachingCacheDetailsId returns a boolean if a field has been set.
func (o *ComputerContentCachingCacheDetail) HasComputerContentCachingCacheDetailsId() bool {
	if o != nil && !IsNil(o.ComputerContentCachingCacheDetailsId) {
		return true
	}

	return false
}

// SetComputerContentCachingCacheDetailsId gets a reference to the given string and assigns it to the ComputerContentCachingCacheDetailsId field.
func (o *ComputerContentCachingCacheDetail) SetComputerContentCachingCacheDetailsId(v string) {
	o.ComputerContentCachingCacheDetailsId = &v
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise.
func (o *ComputerContentCachingCacheDetail) GetCategoryName() string {
	if o == nil || IsNil(o.CategoryName) {
		var ret string
		return ret
	}
	return *o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerContentCachingCacheDetail) GetCategoryNameOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryName) {
		return nil, false
	}
	return o.CategoryName, true
}

// HasCategoryName returns a boolean if a field has been set.
func (o *ComputerContentCachingCacheDetail) HasCategoryName() bool {
	if o != nil && !IsNil(o.CategoryName) {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given string and assigns it to the CategoryName field.
func (o *ComputerContentCachingCacheDetail) SetCategoryName(v string) {
	o.CategoryName = &v
}

// GetDiskSpaceBytesUsed returns the DiskSpaceBytesUsed field value if set, zero value otherwise.
func (o *ComputerContentCachingCacheDetail) GetDiskSpaceBytesUsed() int64 {
	if o == nil || IsNil(o.DiskSpaceBytesUsed) {
		var ret int64
		return ret
	}
	return *o.DiskSpaceBytesUsed
}

// GetDiskSpaceBytesUsedOk returns a tuple with the DiskSpaceBytesUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerContentCachingCacheDetail) GetDiskSpaceBytesUsedOk() (*int64, bool) {
	if o == nil || IsNil(o.DiskSpaceBytesUsed) {
		return nil, false
	}
	return o.DiskSpaceBytesUsed, true
}

// HasDiskSpaceBytesUsed returns a boolean if a field has been set.
func (o *ComputerContentCachingCacheDetail) HasDiskSpaceBytesUsed() bool {
	if o != nil && !IsNil(o.DiskSpaceBytesUsed) {
		return true
	}

	return false
}

// SetDiskSpaceBytesUsed gets a reference to the given int64 and assigns it to the DiskSpaceBytesUsed field.
func (o *ComputerContentCachingCacheDetail) SetDiskSpaceBytesUsed(v int64) {
	o.DiskSpaceBytesUsed = &v
}

func (o ComputerContentCachingCacheDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComputerContentCachingCacheDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComputerContentCachingCacheDetailsId) {
		toSerialize["computerContentCachingCacheDetailsId"] = o.ComputerContentCachingCacheDetailsId
	}
	if !IsNil(o.CategoryName) {
		toSerialize["categoryName"] = o.CategoryName
	}
	if !IsNil(o.DiskSpaceBytesUsed) {
		toSerialize["diskSpaceBytesUsed"] = o.DiskSpaceBytesUsed
	}
	return toSerialize, nil
}

type NullableComputerContentCachingCacheDetail struct {
	value *ComputerContentCachingCacheDetail
	isSet bool
}

func (v NullableComputerContentCachingCacheDetail) Get() *ComputerContentCachingCacheDetail {
	return v.value
}

func (v *NullableComputerContentCachingCacheDetail) Set(val *ComputerContentCachingCacheDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerContentCachingCacheDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerContentCachingCacheDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerContentCachingCacheDetail(val *ComputerContentCachingCacheDetail) *NullableComputerContentCachingCacheDetail {
	return &NullableComputerContentCachingCacheDetail{value: val, isSet: true}
}

func (v NullableComputerContentCachingCacheDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerContentCachingCacheDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


