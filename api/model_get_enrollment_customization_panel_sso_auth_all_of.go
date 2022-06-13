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

// GetEnrollmentCustomizationPanelSsoAuthAllOf struct for GetEnrollmentCustomizationPanelSsoAuthAllOf
type GetEnrollmentCustomizationPanelSsoAuthAllOf struct {
	Id *int32 `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewGetEnrollmentCustomizationPanelSsoAuthAllOf instantiates a new GetEnrollmentCustomizationPanelSsoAuthAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEnrollmentCustomizationPanelSsoAuthAllOf() *GetEnrollmentCustomizationPanelSsoAuthAllOf {
	this := GetEnrollmentCustomizationPanelSsoAuthAllOf{}
	return &this
}

// NewGetEnrollmentCustomizationPanelSsoAuthAllOfWithDefaults instantiates a new GetEnrollmentCustomizationPanelSsoAuthAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEnrollmentCustomizationPanelSsoAuthAllOfWithDefaults() *GetEnrollmentCustomizationPanelSsoAuthAllOf {
	this := GetEnrollmentCustomizationPanelSsoAuthAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetEnrollmentCustomizationPanelSsoAuthAllOf) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationPanelSsoAuthAllOf) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetEnrollmentCustomizationPanelSsoAuthAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetEnrollmentCustomizationPanelSsoAuthAllOf) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetEnrollmentCustomizationPanelSsoAuthAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationPanelSsoAuthAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetEnrollmentCustomizationPanelSsoAuthAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetEnrollmentCustomizationPanelSsoAuthAllOf) SetType(v string) {
	o.Type = &v
}

func (o GetEnrollmentCustomizationPanelSsoAuthAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableGetEnrollmentCustomizationPanelSsoAuthAllOf struct {
	value *GetEnrollmentCustomizationPanelSsoAuthAllOf
	isSet bool
}

func (v NullableGetEnrollmentCustomizationPanelSsoAuthAllOf) Get() *GetEnrollmentCustomizationPanelSsoAuthAllOf {
	return v.value
}

func (v *NullableGetEnrollmentCustomizationPanelSsoAuthAllOf) Set(val *GetEnrollmentCustomizationPanelSsoAuthAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEnrollmentCustomizationPanelSsoAuthAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEnrollmentCustomizationPanelSsoAuthAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEnrollmentCustomizationPanelSsoAuthAllOf(val *GetEnrollmentCustomizationPanelSsoAuthAllOf) *NullableGetEnrollmentCustomizationPanelSsoAuthAllOf {
	return &NullableGetEnrollmentCustomizationPanelSsoAuthAllOf{value: val, isSet: true}
}

func (v NullableGetEnrollmentCustomizationPanelSsoAuthAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEnrollmentCustomizationPanelSsoAuthAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


