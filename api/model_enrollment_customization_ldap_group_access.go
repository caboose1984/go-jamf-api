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

// EnrollmentCustomizationLdapGroupAccess struct for EnrollmentCustomizationLdapGroupAccess
type EnrollmentCustomizationLdapGroupAccess struct {
	LdapServerId *int32 `json:"ldapServerId,omitempty"`
	GroupName *string `json:"groupName,omitempty"`
}

// NewEnrollmentCustomizationLdapGroupAccess instantiates a new EnrollmentCustomizationLdapGroupAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentCustomizationLdapGroupAccess() *EnrollmentCustomizationLdapGroupAccess {
	this := EnrollmentCustomizationLdapGroupAccess{}
	return &this
}

// NewEnrollmentCustomizationLdapGroupAccessWithDefaults instantiates a new EnrollmentCustomizationLdapGroupAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentCustomizationLdapGroupAccessWithDefaults() *EnrollmentCustomizationLdapGroupAccess {
	this := EnrollmentCustomizationLdapGroupAccess{}
	return &this
}

// GetLdapServerId returns the LdapServerId field value if set, zero value otherwise.
func (o *EnrollmentCustomizationLdapGroupAccess) GetLdapServerId() int32 {
	if o == nil || o.LdapServerId == nil {
		var ret int32
		return ret
	}
	return *o.LdapServerId
}

// GetLdapServerIdOk returns a tuple with the LdapServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationLdapGroupAccess) GetLdapServerIdOk() (*int32, bool) {
	if o == nil || o.LdapServerId == nil {
		return nil, false
	}
	return o.LdapServerId, true
}

// HasLdapServerId returns a boolean if a field has been set.
func (o *EnrollmentCustomizationLdapGroupAccess) HasLdapServerId() bool {
	if o != nil && o.LdapServerId != nil {
		return true
	}

	return false
}

// SetLdapServerId gets a reference to the given int32 and assigns it to the LdapServerId field.
func (o *EnrollmentCustomizationLdapGroupAccess) SetLdapServerId(v int32) {
	o.LdapServerId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *EnrollmentCustomizationLdapGroupAccess) GetGroupName() string {
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationLdapGroupAccess) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *EnrollmentCustomizationLdapGroupAccess) HasGroupName() bool {
	if o != nil && o.GroupName != nil {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *EnrollmentCustomizationLdapGroupAccess) SetGroupName(v string) {
	o.GroupName = &v
}

func (o EnrollmentCustomizationLdapGroupAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LdapServerId != nil {
		toSerialize["ldapServerId"] = o.LdapServerId
	}
	if o.GroupName != nil {
		toSerialize["groupName"] = o.GroupName
	}
	return json.Marshal(toSerialize)
}

type NullableEnrollmentCustomizationLdapGroupAccess struct {
	value *EnrollmentCustomizationLdapGroupAccess
	isSet bool
}

func (v NullableEnrollmentCustomizationLdapGroupAccess) Get() *EnrollmentCustomizationLdapGroupAccess {
	return v.value
}

func (v *NullableEnrollmentCustomizationLdapGroupAccess) Set(val *EnrollmentCustomizationLdapGroupAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentCustomizationLdapGroupAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentCustomizationLdapGroupAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentCustomizationLdapGroupAccess(val *EnrollmentCustomizationLdapGroupAccess) *NullableEnrollmentCustomizationLdapGroupAccess {
	return &NullableEnrollmentCustomizationLdapGroupAccess{value: val, isSet: true}
}

func (v NullableEnrollmentCustomizationLdapGroupAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentCustomizationLdapGroupAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


