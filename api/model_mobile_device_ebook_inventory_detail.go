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

// checks if the MobileDeviceEbookInventoryDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MobileDeviceEbookInventoryDetail{}

// MobileDeviceEbookInventoryDetail struct for MobileDeviceEbookInventoryDetail
type MobileDeviceEbookInventoryDetail struct {
	Author *string `json:"author,omitempty"`
	Title *string `json:"title,omitempty"`
	Version *string `json:"version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	ManagementState *string `json:"managementState,omitempty"`
}

// NewMobileDeviceEbookInventoryDetail instantiates a new MobileDeviceEbookInventoryDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileDeviceEbookInventoryDetail() *MobileDeviceEbookInventoryDetail {
	this := MobileDeviceEbookInventoryDetail{}
	return &this
}

// NewMobileDeviceEbookInventoryDetailWithDefaults instantiates a new MobileDeviceEbookInventoryDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileDeviceEbookInventoryDetailWithDefaults() *MobileDeviceEbookInventoryDetail {
	this := MobileDeviceEbookInventoryDetail{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *MobileDeviceEbookInventoryDetail) GetAuthor() string {
	if o == nil || IsNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceEbookInventoryDetail) GetAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *MobileDeviceEbookInventoryDetail) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *MobileDeviceEbookInventoryDetail) SetAuthor(v string) {
	o.Author = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *MobileDeviceEbookInventoryDetail) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceEbookInventoryDetail) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *MobileDeviceEbookInventoryDetail) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *MobileDeviceEbookInventoryDetail) SetTitle(v string) {
	o.Title = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MobileDeviceEbookInventoryDetail) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceEbookInventoryDetail) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MobileDeviceEbookInventoryDetail) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *MobileDeviceEbookInventoryDetail) SetVersion(v string) {
	o.Version = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *MobileDeviceEbookInventoryDetail) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceEbookInventoryDetail) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *MobileDeviceEbookInventoryDetail) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *MobileDeviceEbookInventoryDetail) SetKind(v string) {
	o.Kind = &v
}

// GetManagementState returns the ManagementState field value if set, zero value otherwise.
func (o *MobileDeviceEbookInventoryDetail) GetManagementState() string {
	if o == nil || IsNil(o.ManagementState) {
		var ret string
		return ret
	}
	return *o.ManagementState
}

// GetManagementStateOk returns a tuple with the ManagementState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceEbookInventoryDetail) GetManagementStateOk() (*string, bool) {
	if o == nil || IsNil(o.ManagementState) {
		return nil, false
	}
	return o.ManagementState, true
}

// HasManagementState returns a boolean if a field has been set.
func (o *MobileDeviceEbookInventoryDetail) HasManagementState() bool {
	if o != nil && !IsNil(o.ManagementState) {
		return true
	}

	return false
}

// SetManagementState gets a reference to the given string and assigns it to the ManagementState field.
func (o *MobileDeviceEbookInventoryDetail) SetManagementState(v string) {
	o.ManagementState = &v
}

func (o MobileDeviceEbookInventoryDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobileDeviceEbookInventoryDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.ManagementState) {
		toSerialize["managementState"] = o.ManagementState
	}
	return toSerialize, nil
}

type NullableMobileDeviceEbookInventoryDetail struct {
	value *MobileDeviceEbookInventoryDetail
	isSet bool
}

func (v NullableMobileDeviceEbookInventoryDetail) Get() *MobileDeviceEbookInventoryDetail {
	return v.value
}

func (v *NullableMobileDeviceEbookInventoryDetail) Set(val *MobileDeviceEbookInventoryDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDeviceEbookInventoryDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDeviceEbookInventoryDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDeviceEbookInventoryDetail(val *MobileDeviceEbookInventoryDetail) *NullableMobileDeviceEbookInventoryDetail {
	return &NullableMobileDeviceEbookInventoryDetail{value: val, isSet: true}
}

func (v NullableMobileDeviceEbookInventoryDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDeviceEbookInventoryDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


