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

// checks if the BrandingImageUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandingImageUrl{}

// BrandingImageUrl struct for BrandingImageUrl
type BrandingImageUrl struct {
	Url *string `json:"url,omitempty"`
}

// NewBrandingImageUrl instantiates a new BrandingImageUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingImageUrl() *BrandingImageUrl {
	this := BrandingImageUrl{}
	return &this
}

// NewBrandingImageUrlWithDefaults instantiates a new BrandingImageUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingImageUrlWithDefaults() *BrandingImageUrl {
	this := BrandingImageUrl{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *BrandingImageUrl) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingImageUrl) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *BrandingImageUrl) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *BrandingImageUrl) SetUrl(v string) {
	o.Url = &v
}

func (o BrandingImageUrl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandingImageUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableBrandingImageUrl struct {
	value *BrandingImageUrl
	isSet bool
}

func (v NullableBrandingImageUrl) Get() *BrandingImageUrl {
	return v.value
}

func (v *NullableBrandingImageUrl) Set(val *BrandingImageUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingImageUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingImageUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingImageUrl(val *BrandingImageUrl) *NullableBrandingImageUrl {
	return &NullableBrandingImageUrl{value: val, isSet: true}
}

func (v NullableBrandingImageUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingImageUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


