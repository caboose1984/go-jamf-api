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

// VppAdminAccount struct for VppAdminAccount
type VppAdminAccount struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	LicenseCount *int32 `json:"licenseCount,omitempty"`
	UsedLicenseCount *int32 `json:"usedLicenseCount,omitempty"`
	Location *string `json:"location,omitempty"`
	ExpirationDate *string `json:"expirationDate,omitempty"`
	Site *Site `json:"site,omitempty"`
}

// NewVppAdminAccount instantiates a new VppAdminAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVppAdminAccount() *VppAdminAccount {
	this := VppAdminAccount{}
	return &this
}

// NewVppAdminAccountWithDefaults instantiates a new VppAdminAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVppAdminAccountWithDefaults() *VppAdminAccount {
	this := VppAdminAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VppAdminAccount) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VppAdminAccount) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VppAdminAccount) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VppAdminAccount) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VppAdminAccount) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VppAdminAccount) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VppAdminAccount) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VppAdminAccount) SetName(v string) {
	o.Name = &v
}

// GetLicenseCount returns the LicenseCount field value if set, zero value otherwise.
func (o *VppAdminAccount) GetLicenseCount() int32 {
	if o == nil || o.LicenseCount == nil {
		var ret int32
		return ret
	}
	return *o.LicenseCount
}

// GetLicenseCountOk returns a tuple with the LicenseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VppAdminAccount) GetLicenseCountOk() (*int32, bool) {
	if o == nil || o.LicenseCount == nil {
		return nil, false
	}
	return o.LicenseCount, true
}

// HasLicenseCount returns a boolean if a field has been set.
func (o *VppAdminAccount) HasLicenseCount() bool {
	if o != nil && o.LicenseCount != nil {
		return true
	}

	return false
}

// SetLicenseCount gets a reference to the given int32 and assigns it to the LicenseCount field.
func (o *VppAdminAccount) SetLicenseCount(v int32) {
	o.LicenseCount = &v
}

// GetUsedLicenseCount returns the UsedLicenseCount field value if set, zero value otherwise.
func (o *VppAdminAccount) GetUsedLicenseCount() int32 {
	if o == nil || o.UsedLicenseCount == nil {
		var ret int32
		return ret
	}
	return *o.UsedLicenseCount
}

// GetUsedLicenseCountOk returns a tuple with the UsedLicenseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VppAdminAccount) GetUsedLicenseCountOk() (*int32, bool) {
	if o == nil || o.UsedLicenseCount == nil {
		return nil, false
	}
	return o.UsedLicenseCount, true
}

// HasUsedLicenseCount returns a boolean if a field has been set.
func (o *VppAdminAccount) HasUsedLicenseCount() bool {
	if o != nil && o.UsedLicenseCount != nil {
		return true
	}

	return false
}

// SetUsedLicenseCount gets a reference to the given int32 and assigns it to the UsedLicenseCount field.
func (o *VppAdminAccount) SetUsedLicenseCount(v int32) {
	o.UsedLicenseCount = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *VppAdminAccount) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VppAdminAccount) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *VppAdminAccount) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *VppAdminAccount) SetLocation(v string) {
	o.Location = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *VppAdminAccount) GetExpirationDate() string {
	if o == nil || o.ExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VppAdminAccount) GetExpirationDateOk() (*string, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *VppAdminAccount) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *VppAdminAccount) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *VppAdminAccount) GetSite() Site {
	if o == nil || o.Site == nil {
		var ret Site
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VppAdminAccount) GetSiteOk() (*Site, bool) {
	if o == nil || o.Site == nil {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *VppAdminAccount) HasSite() bool {
	if o != nil && o.Site != nil {
		return true
	}

	return false
}

// SetSite gets a reference to the given Site and assigns it to the Site field.
func (o *VppAdminAccount) SetSite(v Site) {
	o.Site = &v
}

func (o VppAdminAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.LicenseCount != nil {
		toSerialize["licenseCount"] = o.LicenseCount
	}
	if o.UsedLicenseCount != nil {
		toSerialize["usedLicenseCount"] = o.UsedLicenseCount
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.ExpirationDate != nil {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if o.Site != nil {
		toSerialize["site"] = o.Site
	}
	return json.Marshal(toSerialize)
}

type NullableVppAdminAccount struct {
	value *VppAdminAccount
	isSet bool
}

func (v NullableVppAdminAccount) Get() *VppAdminAccount {
	return v.value
}

func (v *NullableVppAdminAccount) Set(val *VppAdminAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableVppAdminAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableVppAdminAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVppAdminAccount(val *VppAdminAccount) *NullableVppAdminAccount {
	return &NullableVppAdminAccount{value: val, isSet: true}
}

func (v NullableVppAdminAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVppAdminAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

