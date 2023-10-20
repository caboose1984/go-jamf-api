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

// checks if the ComputerPackageReceipts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComputerPackageReceipts{}

// ComputerPackageReceipts All package receipts are listed by their package name
type ComputerPackageReceipts struct {
	InstalledByJamfPro []string `json:"installedByJamfPro,omitempty"`
	InstalledByInstallerSwu []string `json:"installedByInstallerSwu,omitempty"`
	Cached []string `json:"cached,omitempty"`
}

// NewComputerPackageReceipts instantiates a new ComputerPackageReceipts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerPackageReceipts() *ComputerPackageReceipts {
	this := ComputerPackageReceipts{}
	return &this
}

// NewComputerPackageReceiptsWithDefaults instantiates a new ComputerPackageReceipts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerPackageReceiptsWithDefaults() *ComputerPackageReceipts {
	this := ComputerPackageReceipts{}
	return &this
}

// GetInstalledByJamfPro returns the InstalledByJamfPro field value if set, zero value otherwise.
func (o *ComputerPackageReceipts) GetInstalledByJamfPro() []string {
	if o == nil || IsNil(o.InstalledByJamfPro) {
		var ret []string
		return ret
	}
	return o.InstalledByJamfPro
}

// GetInstalledByJamfProOk returns a tuple with the InstalledByJamfPro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPackageReceipts) GetInstalledByJamfProOk() ([]string, bool) {
	if o == nil || IsNil(o.InstalledByJamfPro) {
		return nil, false
	}
	return o.InstalledByJamfPro, true
}

// HasInstalledByJamfPro returns a boolean if a field has been set.
func (o *ComputerPackageReceipts) HasInstalledByJamfPro() bool {
	if o != nil && !IsNil(o.InstalledByJamfPro) {
		return true
	}

	return false
}

// SetInstalledByJamfPro gets a reference to the given []string and assigns it to the InstalledByJamfPro field.
func (o *ComputerPackageReceipts) SetInstalledByJamfPro(v []string) {
	o.InstalledByJamfPro = v
}

// GetInstalledByInstallerSwu returns the InstalledByInstallerSwu field value if set, zero value otherwise.
func (o *ComputerPackageReceipts) GetInstalledByInstallerSwu() []string {
	if o == nil || IsNil(o.InstalledByInstallerSwu) {
		var ret []string
		return ret
	}
	return o.InstalledByInstallerSwu
}

// GetInstalledByInstallerSwuOk returns a tuple with the InstalledByInstallerSwu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPackageReceipts) GetInstalledByInstallerSwuOk() ([]string, bool) {
	if o == nil || IsNil(o.InstalledByInstallerSwu) {
		return nil, false
	}
	return o.InstalledByInstallerSwu, true
}

// HasInstalledByInstallerSwu returns a boolean if a field has been set.
func (o *ComputerPackageReceipts) HasInstalledByInstallerSwu() bool {
	if o != nil && !IsNil(o.InstalledByInstallerSwu) {
		return true
	}

	return false
}

// SetInstalledByInstallerSwu gets a reference to the given []string and assigns it to the InstalledByInstallerSwu field.
func (o *ComputerPackageReceipts) SetInstalledByInstallerSwu(v []string) {
	o.InstalledByInstallerSwu = v
}

// GetCached returns the Cached field value if set, zero value otherwise.
func (o *ComputerPackageReceipts) GetCached() []string {
	if o == nil || IsNil(o.Cached) {
		var ret []string
		return ret
	}
	return o.Cached
}

// GetCachedOk returns a tuple with the Cached field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPackageReceipts) GetCachedOk() ([]string, bool) {
	if o == nil || IsNil(o.Cached) {
		return nil, false
	}
	return o.Cached, true
}

// HasCached returns a boolean if a field has been set.
func (o *ComputerPackageReceipts) HasCached() bool {
	if o != nil && !IsNil(o.Cached) {
		return true
	}

	return false
}

// SetCached gets a reference to the given []string and assigns it to the Cached field.
func (o *ComputerPackageReceipts) SetCached(v []string) {
	o.Cached = v
}

func (o ComputerPackageReceipts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComputerPackageReceipts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstalledByJamfPro) {
		toSerialize["installedByJamfPro"] = o.InstalledByJamfPro
	}
	if !IsNil(o.InstalledByInstallerSwu) {
		toSerialize["installedByInstallerSwu"] = o.InstalledByInstallerSwu
	}
	if !IsNil(o.Cached) {
		toSerialize["cached"] = o.Cached
	}
	return toSerialize, nil
}

type NullableComputerPackageReceipts struct {
	value *ComputerPackageReceipts
	isSet bool
}

func (v NullableComputerPackageReceipts) Get() *ComputerPackageReceipts {
	return v.value
}

func (v *NullableComputerPackageReceipts) Set(val *ComputerPackageReceipts) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerPackageReceipts) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerPackageReceipts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerPackageReceipts(val *ComputerPackageReceipts) *NullableComputerPackageReceipts {
	return &NullableComputerPackageReceipts{value: val, isSet: true}
}

func (v NullableComputerPackageReceipts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerPackageReceipts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


