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

// DeprecatedServerUpdate An old Cloud Identity Provider LDAP server configuration for updates
type DeprecatedServerUpdate struct {
	Id string `json:"id"`
	Enabled bool `json:"enabled"`
	ProviderName string `json:"providerName"`
	DisplayName string `json:"displayName"`
	ServerUrl string `json:"serverUrl"`
	DomainName string `json:"domainName"`
	Port int32 `json:"port"`
	Keystore *CloudLdapKeystoreFile `json:"keystore,omitempty"`
	ConnectionTimeout int32 `json:"connectionTimeout"`
	SearchTimeout int32 `json:"searchTimeout"`
	UseWildcards bool `json:"useWildcards"`
	ConnectionType string `json:"connectionType"`
}

// NewDeprecatedServerUpdate instantiates a new DeprecatedServerUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedServerUpdate(id string, enabled bool, providerName string, displayName string, serverUrl string, domainName string, port int32, connectionTimeout int32, searchTimeout int32, useWildcards bool, connectionType string) *DeprecatedServerUpdate {
	this := DeprecatedServerUpdate{}
	this.Id = id
	this.Enabled = enabled
	this.ProviderName = providerName
	this.DisplayName = displayName
	this.ServerUrl = serverUrl
	this.DomainName = domainName
	this.Port = port
	this.ConnectionTimeout = connectionTimeout
	this.SearchTimeout = searchTimeout
	this.UseWildcards = useWildcards
	this.ConnectionType = connectionType
	return &this
}

// NewDeprecatedServerUpdateWithDefaults instantiates a new DeprecatedServerUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedServerUpdateWithDefaults() *DeprecatedServerUpdate {
	this := DeprecatedServerUpdate{}
	return &this
}

// GetId returns the Id field value
func (o *DeprecatedServerUpdate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerUpdate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeprecatedServerUpdate) SetId(v string) {
	o.Id = v
}

// GetEnabled returns the Enabled field value
func (o *DeprecatedServerUpdate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DeprecatedServerUpdate) SetEnabled(v bool) {
	o.Enabled = v
}

// GetProviderName returns the ProviderName field value
func (o *DeprecatedServerUpdate) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerUpdate) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *DeprecatedServerUpdate) SetProviderName(v string) {
	o.ProviderName = v
}

// GetDisplayName returns the DisplayName field value
func (o *DeprecatedServerUpdate) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerUpdate) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *DeprecatedServerUpdate) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetServerUrl returns the ServerUrl field value
func (o *DeprecatedServerUpdate) GetServerUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerUrl
}

// GetServerUrlOk returns a tuple with the ServerUrl field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerUpdate) GetServerUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerUrl, true
}

// SetServerUrl sets field value
func (o *DeprecatedServerUpdate) SetServerUrl(v string) {
	o.ServerUrl = v
}

// GetDomainName returns the DomainName field value
func (o *DeprecatedServerUpdate) GetDomainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerUpdate) GetDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainName, true
}

// SetDomainName sets field value
func (o *DeprecatedServerUpdate) SetDomainName(v string) {
	o.DomainName = v
}

// GetPort returns the Port field value
func (o *DeprecatedServerUpdate) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerUpdate) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *DeprecatedServerUpdate) SetPort(v int32) {
	o.Port = v
}

// GetKeystore returns the Keystore field value if set, zero value otherwise.
func (o *DeprecatedServerUpdate) GetKeystore() CloudLdapKeystoreFile {
	if o == nil || o.Keystore == nil {
		var ret CloudLdapKeystoreFile
		return ret
	}
	return *o.Keystore
}

// GetKeystoreOk returns a tuple with the Keystore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedServerUpdate) GetKeystoreOk() (*CloudLdapKeystoreFile, bool) {
	if o == nil || o.Keystore == nil {
		return nil, false
	}
	return o.Keystore, true
}

// HasKeystore returns a boolean if a field has been set.
func (o *DeprecatedServerUpdate) HasKeystore() bool {
	if o != nil && o.Keystore != nil {
		return true
	}

	return false
}

// SetKeystore gets a reference to the given CloudLdapKeystoreFile and assigns it to the Keystore field.
func (o *DeprecatedServerUpdate) SetKeystore(v CloudLdapKeystoreFile) {
	o.Keystore = &v
}

// GetConnectionTimeout returns the ConnectionTimeout field value
func (o *DeprecatedServerUpdate) GetConnectionTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConnectionTimeout
}

// GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerUpdate) GetConnectionTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionTimeout, true
}

// SetConnectionTimeout sets field value
func (o *DeprecatedServerUpdate) SetConnectionTimeout(v int32) {
	o.ConnectionTimeout = v
}

// GetSearchTimeout returns the SearchTimeout field value
func (o *DeprecatedServerUpdate) GetSearchTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SearchTimeout
}

// GetSearchTimeoutOk returns a tuple with the SearchTimeout field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerUpdate) GetSearchTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchTimeout, true
}

// SetSearchTimeout sets field value
func (o *DeprecatedServerUpdate) SetSearchTimeout(v int32) {
	o.SearchTimeout = v
}

// GetUseWildcards returns the UseWildcards field value
func (o *DeprecatedServerUpdate) GetUseWildcards() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseWildcards
}

// GetUseWildcardsOk returns a tuple with the UseWildcards field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerUpdate) GetUseWildcardsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseWildcards, true
}

// SetUseWildcards sets field value
func (o *DeprecatedServerUpdate) SetUseWildcards(v bool) {
	o.UseWildcards = v
}

// GetConnectionType returns the ConnectionType field value
func (o *DeprecatedServerUpdate) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerUpdate) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *DeprecatedServerUpdate) SetConnectionType(v string) {
	o.ConnectionType = v
}

func (o DeprecatedServerUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["providerName"] = o.ProviderName
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["serverUrl"] = o.ServerUrl
	}
	if true {
		toSerialize["domainName"] = o.DomainName
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if o.Keystore != nil {
		toSerialize["keystore"] = o.Keystore
	}
	if true {
		toSerialize["connectionTimeout"] = o.ConnectionTimeout
	}
	if true {
		toSerialize["searchTimeout"] = o.SearchTimeout
	}
	if true {
		toSerialize["useWildcards"] = o.UseWildcards
	}
	if true {
		toSerialize["connectionType"] = o.ConnectionType
	}
	return json.Marshal(toSerialize)
}

type NullableDeprecatedServerUpdate struct {
	value *DeprecatedServerUpdate
	isSet bool
}

func (v NullableDeprecatedServerUpdate) Get() *DeprecatedServerUpdate {
	return v.value
}

func (v *NullableDeprecatedServerUpdate) Set(val *DeprecatedServerUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedServerUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedServerUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedServerUpdate(val *DeprecatedServerUpdate) *NullableDeprecatedServerUpdate {
	return &NullableDeprecatedServerUpdate{value: val, isSet: true}
}

func (v NullableDeprecatedServerUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedServerUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


