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

// AzureServerConfigurationUpdate Azure Cloud Identity Provider configuration update
type AzureServerConfigurationUpdate struct {
	Id string `json:"id"`
	Enabled bool `json:"enabled"`
	Mappings AzureMappings `json:"mappings"`
	SearchTimeout int32 `json:"searchTimeout"`
	// Use this field to enable transitive membership lookup with Single Sign On
	TransitiveMembershipEnabled bool `json:"transitiveMembershipEnabled"`
	// Use this field to set user field mapping for transitive membership lookup with Single Sign On
	TransitiveMembershipUserField string `json:"transitiveMembershipUserField"`
	// Use this field to enable transitive membership lookup. This setting would not apply to Single Sign On
	TransitiveDirectoryMembershipEnabled bool `json:"transitiveDirectoryMembershipEnabled"`
}

// NewAzureServerConfigurationUpdate instantiates a new AzureServerConfigurationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureServerConfigurationUpdate(id string, enabled bool, mappings AzureMappings, searchTimeout int32, transitiveMembershipEnabled bool, transitiveMembershipUserField string, transitiveDirectoryMembershipEnabled bool) *AzureServerConfigurationUpdate {
	this := AzureServerConfigurationUpdate{}
	this.Id = id
	this.Enabled = enabled
	this.Mappings = mappings
	this.SearchTimeout = searchTimeout
	this.TransitiveMembershipEnabled = transitiveMembershipEnabled
	this.TransitiveMembershipUserField = transitiveMembershipUserField
	this.TransitiveDirectoryMembershipEnabled = transitiveDirectoryMembershipEnabled
	return &this
}

// NewAzureServerConfigurationUpdateWithDefaults instantiates a new AzureServerConfigurationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureServerConfigurationUpdateWithDefaults() *AzureServerConfigurationUpdate {
	this := AzureServerConfigurationUpdate{}
	return &this
}

// GetId returns the Id field value
func (o *AzureServerConfigurationUpdate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AzureServerConfigurationUpdate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AzureServerConfigurationUpdate) SetId(v string) {
	o.Id = v
}

// GetEnabled returns the Enabled field value
func (o *AzureServerConfigurationUpdate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AzureServerConfigurationUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AzureServerConfigurationUpdate) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMappings returns the Mappings field value
func (o *AzureServerConfigurationUpdate) GetMappings() AzureMappings {
	if o == nil {
		var ret AzureMappings
		return ret
	}

	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value
// and a boolean to check if the value has been set.
func (o *AzureServerConfigurationUpdate) GetMappingsOk() (*AzureMappings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mappings, true
}

// SetMappings sets field value
func (o *AzureServerConfigurationUpdate) SetMappings(v AzureMappings) {
	o.Mappings = v
}

// GetSearchTimeout returns the SearchTimeout field value
func (o *AzureServerConfigurationUpdate) GetSearchTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SearchTimeout
}

// GetSearchTimeoutOk returns a tuple with the SearchTimeout field value
// and a boolean to check if the value has been set.
func (o *AzureServerConfigurationUpdate) GetSearchTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchTimeout, true
}

// SetSearchTimeout sets field value
func (o *AzureServerConfigurationUpdate) SetSearchTimeout(v int32) {
	o.SearchTimeout = v
}

// GetTransitiveMembershipEnabled returns the TransitiveMembershipEnabled field value
func (o *AzureServerConfigurationUpdate) GetTransitiveMembershipEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TransitiveMembershipEnabled
}

// GetTransitiveMembershipEnabledOk returns a tuple with the TransitiveMembershipEnabled field value
// and a boolean to check if the value has been set.
func (o *AzureServerConfigurationUpdate) GetTransitiveMembershipEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransitiveMembershipEnabled, true
}

// SetTransitiveMembershipEnabled sets field value
func (o *AzureServerConfigurationUpdate) SetTransitiveMembershipEnabled(v bool) {
	o.TransitiveMembershipEnabled = v
}

// GetTransitiveMembershipUserField returns the TransitiveMembershipUserField field value
func (o *AzureServerConfigurationUpdate) GetTransitiveMembershipUserField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransitiveMembershipUserField
}

// GetTransitiveMembershipUserFieldOk returns a tuple with the TransitiveMembershipUserField field value
// and a boolean to check if the value has been set.
func (o *AzureServerConfigurationUpdate) GetTransitiveMembershipUserFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransitiveMembershipUserField, true
}

// SetTransitiveMembershipUserField sets field value
func (o *AzureServerConfigurationUpdate) SetTransitiveMembershipUserField(v string) {
	o.TransitiveMembershipUserField = v
}

// GetTransitiveDirectoryMembershipEnabled returns the TransitiveDirectoryMembershipEnabled field value
func (o *AzureServerConfigurationUpdate) GetTransitiveDirectoryMembershipEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TransitiveDirectoryMembershipEnabled
}

// GetTransitiveDirectoryMembershipEnabledOk returns a tuple with the TransitiveDirectoryMembershipEnabled field value
// and a boolean to check if the value has been set.
func (o *AzureServerConfigurationUpdate) GetTransitiveDirectoryMembershipEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransitiveDirectoryMembershipEnabled, true
}

// SetTransitiveDirectoryMembershipEnabled sets field value
func (o *AzureServerConfigurationUpdate) SetTransitiveDirectoryMembershipEnabled(v bool) {
	o.TransitiveDirectoryMembershipEnabled = v
}

func (o AzureServerConfigurationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["mappings"] = o.Mappings
	}
	if true {
		toSerialize["searchTimeout"] = o.SearchTimeout
	}
	if true {
		toSerialize["transitiveMembershipEnabled"] = o.TransitiveMembershipEnabled
	}
	if true {
		toSerialize["transitiveMembershipUserField"] = o.TransitiveMembershipUserField
	}
	if true {
		toSerialize["transitiveDirectoryMembershipEnabled"] = o.TransitiveDirectoryMembershipEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableAzureServerConfigurationUpdate struct {
	value *AzureServerConfigurationUpdate
	isSet bool
}

func (v NullableAzureServerConfigurationUpdate) Get() *AzureServerConfigurationUpdate {
	return v.value
}

func (v *NullableAzureServerConfigurationUpdate) Set(val *AzureServerConfigurationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureServerConfigurationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureServerConfigurationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureServerConfigurationUpdate(val *AzureServerConfigurationUpdate) *NullableAzureServerConfigurationUpdate {
	return &NullableAzureServerConfigurationUpdate{value: val, isSet: true}
}

func (v NullableAzureServerConfigurationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureServerConfigurationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


