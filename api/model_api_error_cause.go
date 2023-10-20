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

// checks if the ApiErrorCause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiErrorCause{}

// ApiErrorCause struct for ApiErrorCause
type ApiErrorCause struct {
	// Error-specific code that can be used to identify localization string, etc.
	Code *string `json:"code,omitempty"`
	// Name of the field that caused the error.
	Field string `json:"field"`
	// A general description of error for troubleshooting/debugging. Generally this text should not be displayed to a user; instead refer to errorCode and it's localized text
	Description *string `json:"description,omitempty"`
	// id of object with error. Optional.
	Id NullableString `json:"id,omitempty"`
}

// NewApiErrorCause instantiates a new ApiErrorCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiErrorCause(field string) *ApiErrorCause {
	this := ApiErrorCause{}
	this.Field = field
	return &this
}

// NewApiErrorCauseWithDefaults instantiates a new ApiErrorCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorCauseWithDefaults() *ApiErrorCause {
	this := ApiErrorCause{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiErrorCause) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiErrorCause) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiErrorCause) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ApiErrorCause) SetCode(v string) {
	o.Code = &v
}

// GetField returns the Field field value
func (o *ApiErrorCause) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *ApiErrorCause) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *ApiErrorCause) SetField(v string) {
	o.Field = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiErrorCause) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiErrorCause) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiErrorCause) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiErrorCause) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiErrorCause) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiErrorCause) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ApiErrorCause) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ApiErrorCause) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ApiErrorCause) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ApiErrorCause) UnsetId() {
	o.Id.Unset()
}

func (o ApiErrorCause) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiErrorCause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	toSerialize["field"] = o.Field
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	return toSerialize, nil
}

type NullableApiErrorCause struct {
	value *ApiErrorCause
	isSet bool
}

func (v NullableApiErrorCause) Get() *ApiErrorCause {
	return v.value
}

func (v *NullableApiErrorCause) Set(val *ApiErrorCause) {
	v.value = val
	v.isSet = true
}

func (v NullableApiErrorCause) IsSet() bool {
	return v.isSet
}

func (v *NullableApiErrorCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiErrorCause(val *ApiErrorCause) *NullableApiErrorCause {
	return &NullableApiErrorCause{value: val, isSet: true}
}

func (v NullableApiErrorCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiErrorCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


