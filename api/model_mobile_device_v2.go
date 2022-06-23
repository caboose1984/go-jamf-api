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

// MobileDeviceV2 struct for MobileDeviceV2
type MobileDeviceV2 struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`
	WifiMacAddress *string `json:"wifiMacAddress,omitempty"`
	Udid *string `json:"udid,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Model *string `json:"model,omitempty"`
	ModelIdentifier *string `json:"modelIdentifier,omitempty"`
	Username *string `json:"username,omitempty"`
	Type *string `json:"type,omitempty"`
	ManagementId *string `json:"managementId,omitempty"`
	SoftwareUpdateDeviceId *string `json:"softwareUpdateDeviceId,omitempty"`
}

// NewMobileDeviceV2 instantiates a new MobileDeviceV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileDeviceV2() *MobileDeviceV2 {
	this := MobileDeviceV2{}
	return &this
}

// NewMobileDeviceV2WithDefaults instantiates a new MobileDeviceV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileDeviceV2WithDefaults() *MobileDeviceV2 {
	this := MobileDeviceV2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MobileDeviceV2) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceV2) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MobileDeviceV2) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MobileDeviceV2) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MobileDeviceV2) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceV2) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MobileDeviceV2) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MobileDeviceV2) SetName(v string) {
	o.Name = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *MobileDeviceV2) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceV2) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *MobileDeviceV2) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *MobileDeviceV2) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetWifiMacAddress returns the WifiMacAddress field value if set, zero value otherwise.
func (o *MobileDeviceV2) GetWifiMacAddress() string {
	if o == nil || o.WifiMacAddress == nil {
		var ret string
		return ret
	}
	return *o.WifiMacAddress
}

// GetWifiMacAddressOk returns a tuple with the WifiMacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceV2) GetWifiMacAddressOk() (*string, bool) {
	if o == nil || o.WifiMacAddress == nil {
		return nil, false
	}
	return o.WifiMacAddress, true
}

// HasWifiMacAddress returns a boolean if a field has been set.
func (o *MobileDeviceV2) HasWifiMacAddress() bool {
	if o != nil && o.WifiMacAddress != nil {
		return true
	}

	return false
}

// SetWifiMacAddress gets a reference to the given string and assigns it to the WifiMacAddress field.
func (o *MobileDeviceV2) SetWifiMacAddress(v string) {
	o.WifiMacAddress = &v
}

// GetUdid returns the Udid field value if set, zero value otherwise.
func (o *MobileDeviceV2) GetUdid() string {
	if o == nil || o.Udid == nil {
		var ret string
		return ret
	}
	return *o.Udid
}

// GetUdidOk returns a tuple with the Udid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceV2) GetUdidOk() (*string, bool) {
	if o == nil || o.Udid == nil {
		return nil, false
	}
	return o.Udid, true
}

// HasUdid returns a boolean if a field has been set.
func (o *MobileDeviceV2) HasUdid() bool {
	if o != nil && o.Udid != nil {
		return true
	}

	return false
}

// SetUdid gets a reference to the given string and assigns it to the Udid field.
func (o *MobileDeviceV2) SetUdid(v string) {
	o.Udid = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *MobileDeviceV2) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceV2) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *MobileDeviceV2) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *MobileDeviceV2) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *MobileDeviceV2) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceV2) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *MobileDeviceV2) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *MobileDeviceV2) SetModel(v string) {
	o.Model = &v
}

// GetModelIdentifier returns the ModelIdentifier field value if set, zero value otherwise.
func (o *MobileDeviceV2) GetModelIdentifier() string {
	if o == nil || o.ModelIdentifier == nil {
		var ret string
		return ret
	}
	return *o.ModelIdentifier
}

// GetModelIdentifierOk returns a tuple with the ModelIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceV2) GetModelIdentifierOk() (*string, bool) {
	if o == nil || o.ModelIdentifier == nil {
		return nil, false
	}
	return o.ModelIdentifier, true
}

// HasModelIdentifier returns a boolean if a field has been set.
func (o *MobileDeviceV2) HasModelIdentifier() bool {
	if o != nil && o.ModelIdentifier != nil {
		return true
	}

	return false
}

// SetModelIdentifier gets a reference to the given string and assigns it to the ModelIdentifier field.
func (o *MobileDeviceV2) SetModelIdentifier(v string) {
	o.ModelIdentifier = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *MobileDeviceV2) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceV2) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *MobileDeviceV2) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *MobileDeviceV2) SetUsername(v string) {
	o.Username = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MobileDeviceV2) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceV2) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MobileDeviceV2) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MobileDeviceV2) SetType(v string) {
	o.Type = &v
}

// GetManagementId returns the ManagementId field value if set, zero value otherwise.
func (o *MobileDeviceV2) GetManagementId() string {
	if o == nil || o.ManagementId == nil {
		var ret string
		return ret
	}
	return *o.ManagementId
}

// GetManagementIdOk returns a tuple with the ManagementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceV2) GetManagementIdOk() (*string, bool) {
	if o == nil || o.ManagementId == nil {
		return nil, false
	}
	return o.ManagementId, true
}

// HasManagementId returns a boolean if a field has been set.
func (o *MobileDeviceV2) HasManagementId() bool {
	if o != nil && o.ManagementId != nil {
		return true
	}

	return false
}

// SetManagementId gets a reference to the given string and assigns it to the ManagementId field.
func (o *MobileDeviceV2) SetManagementId(v string) {
	o.ManagementId = &v
}

// GetSoftwareUpdateDeviceId returns the SoftwareUpdateDeviceId field value if set, zero value otherwise.
func (o *MobileDeviceV2) GetSoftwareUpdateDeviceId() string {
	if o == nil || o.SoftwareUpdateDeviceId == nil {
		var ret string
		return ret
	}
	return *o.SoftwareUpdateDeviceId
}

// GetSoftwareUpdateDeviceIdOk returns a tuple with the SoftwareUpdateDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceV2) GetSoftwareUpdateDeviceIdOk() (*string, bool) {
	if o == nil || o.SoftwareUpdateDeviceId == nil {
		return nil, false
	}
	return o.SoftwareUpdateDeviceId, true
}

// HasSoftwareUpdateDeviceId returns a boolean if a field has been set.
func (o *MobileDeviceV2) HasSoftwareUpdateDeviceId() bool {
	if o != nil && o.SoftwareUpdateDeviceId != nil {
		return true
	}

	return false
}

// SetSoftwareUpdateDeviceId gets a reference to the given string and assigns it to the SoftwareUpdateDeviceId field.
func (o *MobileDeviceV2) SetSoftwareUpdateDeviceId(v string) {
	o.SoftwareUpdateDeviceId = &v
}

func (o MobileDeviceV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SerialNumber != nil {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if o.WifiMacAddress != nil {
		toSerialize["wifiMacAddress"] = o.WifiMacAddress
	}
	if o.Udid != nil {
		toSerialize["udid"] = o.Udid
	}
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.ModelIdentifier != nil {
		toSerialize["modelIdentifier"] = o.ModelIdentifier
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ManagementId != nil {
		toSerialize["managementId"] = o.ManagementId
	}
	if o.SoftwareUpdateDeviceId != nil {
		toSerialize["softwareUpdateDeviceId"] = o.SoftwareUpdateDeviceId
	}
	return json.Marshal(toSerialize)
}

type NullableMobileDeviceV2 struct {
	value *MobileDeviceV2
	isSet bool
}

func (v NullableMobileDeviceV2) Get() *MobileDeviceV2 {
	return v.value
}

func (v *NullableMobileDeviceV2) Set(val *MobileDeviceV2) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDeviceV2) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDeviceV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDeviceV2(val *MobileDeviceV2) *NullableMobileDeviceV2 {
	return &NullableMobileDeviceV2{value: val, isSet: true}
}

func (v NullableMobileDeviceV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDeviceV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

