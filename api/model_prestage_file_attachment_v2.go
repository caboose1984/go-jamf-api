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

// PrestageFileAttachmentV2 struct for PrestageFileAttachmentV2
type PrestageFileAttachmentV2 struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	FileType *string `json:"fileType,omitempty"`
}

// NewPrestageFileAttachmentV2 instantiates a new PrestageFileAttachmentV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrestageFileAttachmentV2() *PrestageFileAttachmentV2 {
	this := PrestageFileAttachmentV2{}
	return &this
}

// NewPrestageFileAttachmentV2WithDefaults instantiates a new PrestageFileAttachmentV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrestageFileAttachmentV2WithDefaults() *PrestageFileAttachmentV2 {
	this := PrestageFileAttachmentV2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PrestageFileAttachmentV2) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageFileAttachmentV2) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PrestageFileAttachmentV2) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PrestageFileAttachmentV2) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PrestageFileAttachmentV2) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageFileAttachmentV2) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PrestageFileAttachmentV2) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PrestageFileAttachmentV2) SetName(v string) {
	o.Name = &v
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *PrestageFileAttachmentV2) GetFileType() string {
	if o == nil || o.FileType == nil {
		var ret string
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageFileAttachmentV2) GetFileTypeOk() (*string, bool) {
	if o == nil || o.FileType == nil {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *PrestageFileAttachmentV2) HasFileType() bool {
	if o != nil && o.FileType != nil {
		return true
	}

	return false
}

// SetFileType gets a reference to the given string and assigns it to the FileType field.
func (o *PrestageFileAttachmentV2) SetFileType(v string) {
	o.FileType = &v
}

func (o PrestageFileAttachmentV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FileType != nil {
		toSerialize["fileType"] = o.FileType
	}
	return json.Marshal(toSerialize)
}

type NullablePrestageFileAttachmentV2 struct {
	value *PrestageFileAttachmentV2
	isSet bool
}

func (v NullablePrestageFileAttachmentV2) Get() *PrestageFileAttachmentV2 {
	return v.value
}

func (v *NullablePrestageFileAttachmentV2) Set(val *PrestageFileAttachmentV2) {
	v.value = val
	v.isSet = true
}

func (v NullablePrestageFileAttachmentV2) IsSet() bool {
	return v.isSet
}

func (v *NullablePrestageFileAttachmentV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrestageFileAttachmentV2(val *PrestageFileAttachmentV2) *NullablePrestageFileAttachmentV2 {
	return &NullablePrestageFileAttachmentV2{value: val, isSet: true}
}

func (v NullablePrestageFileAttachmentV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrestageFileAttachmentV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


