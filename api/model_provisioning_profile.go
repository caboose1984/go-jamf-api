/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// ProvisioningProfile struct for ProvisioningProfile
type ProvisioningProfile struct {
	DisplayName *string `json:"displayName,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
}

// NewProvisioningProfile instantiates a new ProvisioningProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningProfile() *ProvisioningProfile {
	this := ProvisioningProfile{}
	return &this
}

// NewProvisioningProfileWithDefaults instantiates a new ProvisioningProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningProfileWithDefaults() *ProvisioningProfile {
	this := ProvisioningProfile{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ProvisioningProfile) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningProfile) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ProvisioningProfile) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ProvisioningProfile) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ProvisioningProfile) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningProfile) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ProvisioningProfile) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ProvisioningProfile) SetUuid(v string) {
	o.Uuid = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *ProvisioningProfile) GetExpirationDate() time.Time {
	if o == nil || o.ExpirationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningProfile) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ProvisioningProfile) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *ProvisioningProfile) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

func (o ProvisioningProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.ExpirationDate != nil {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	return json.Marshal(toSerialize)
}

type NullableProvisioningProfile struct {
	value *ProvisioningProfile
	isSet bool
}

func (v NullableProvisioningProfile) Get() *ProvisioningProfile {
	return v.value
}

func (v *NullableProvisioningProfile) Set(val *ProvisioningProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningProfile(val *ProvisioningProfile) *NullableProvisioningProfile {
	return &NullableProvisioningProfile{value: val, isSet: true}
}

func (v NullableProvisioningProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


