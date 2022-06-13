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

// Network struct for Network
type Network struct {
	CellularTechnology *string `json:"cellularTechnology,omitempty"`
	IsVoiceRoamingEnabled *bool `json:"isVoiceRoamingEnabled,omitempty"`
	Imei *string `json:"imei,omitempty"`
	Iccid *string `json:"iccid,omitempty"`
	Meid *string `json:"meid,omitempty"`
	CarrierSettingsVersion *string `json:"carrierSettingsVersion,omitempty"`
	CurrentCarrierNetwork *string `json:"currentCarrierNetwork,omitempty"`
	CurrentMobileCountryCode *string `json:"currentMobileCountryCode,omitempty"`
	CurrentMobileNetworkCode *string `json:"currentMobileNetworkCode,omitempty"`
	HomeCarrierNetwork *string `json:"homeCarrierNetwork,omitempty"`
	HomeMobileCountryCode *string `json:"homeMobileCountryCode,omitempty"`
	HomeMobileNetworkCode *string `json:"homeMobileNetworkCode,omitempty"`
	IsDataRoamingEnabled *bool `json:"isDataRoamingEnabled,omitempty"`
	IsRoaming *bool `json:"isRoaming,omitempty"`
	IsPersonalHotspotEnabled *bool `json:"isPersonalHotspotEnabled,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// NewNetwork instantiates a new Network object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetwork() *Network {
	this := Network{}
	return &this
}

// NewNetworkWithDefaults instantiates a new Network object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkWithDefaults() *Network {
	this := Network{}
	return &this
}

// GetCellularTechnology returns the CellularTechnology field value if set, zero value otherwise.
func (o *Network) GetCellularTechnology() string {
	if o == nil || o.CellularTechnology == nil {
		var ret string
		return ret
	}
	return *o.CellularTechnology
}

// GetCellularTechnologyOk returns a tuple with the CellularTechnology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetCellularTechnologyOk() (*string, bool) {
	if o == nil || o.CellularTechnology == nil {
		return nil, false
	}
	return o.CellularTechnology, true
}

// HasCellularTechnology returns a boolean if a field has been set.
func (o *Network) HasCellularTechnology() bool {
	if o != nil && o.CellularTechnology != nil {
		return true
	}

	return false
}

// SetCellularTechnology gets a reference to the given string and assigns it to the CellularTechnology field.
func (o *Network) SetCellularTechnology(v string) {
	o.CellularTechnology = &v
}

// GetIsVoiceRoamingEnabled returns the IsVoiceRoamingEnabled field value if set, zero value otherwise.
func (o *Network) GetIsVoiceRoamingEnabled() bool {
	if o == nil || o.IsVoiceRoamingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsVoiceRoamingEnabled
}

// GetIsVoiceRoamingEnabledOk returns a tuple with the IsVoiceRoamingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetIsVoiceRoamingEnabledOk() (*bool, bool) {
	if o == nil || o.IsVoiceRoamingEnabled == nil {
		return nil, false
	}
	return o.IsVoiceRoamingEnabled, true
}

// HasIsVoiceRoamingEnabled returns a boolean if a field has been set.
func (o *Network) HasIsVoiceRoamingEnabled() bool {
	if o != nil && o.IsVoiceRoamingEnabled != nil {
		return true
	}

	return false
}

// SetIsVoiceRoamingEnabled gets a reference to the given bool and assigns it to the IsVoiceRoamingEnabled field.
func (o *Network) SetIsVoiceRoamingEnabled(v bool) {
	o.IsVoiceRoamingEnabled = &v
}

// GetImei returns the Imei field value if set, zero value otherwise.
func (o *Network) GetImei() string {
	if o == nil || o.Imei == nil {
		var ret string
		return ret
	}
	return *o.Imei
}

// GetImeiOk returns a tuple with the Imei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetImeiOk() (*string, bool) {
	if o == nil || o.Imei == nil {
		return nil, false
	}
	return o.Imei, true
}

// HasImei returns a boolean if a field has been set.
func (o *Network) HasImei() bool {
	if o != nil && o.Imei != nil {
		return true
	}

	return false
}

// SetImei gets a reference to the given string and assigns it to the Imei field.
func (o *Network) SetImei(v string) {
	o.Imei = &v
}

// GetIccid returns the Iccid field value if set, zero value otherwise.
func (o *Network) GetIccid() string {
	if o == nil || o.Iccid == nil {
		var ret string
		return ret
	}
	return *o.Iccid
}

// GetIccidOk returns a tuple with the Iccid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetIccidOk() (*string, bool) {
	if o == nil || o.Iccid == nil {
		return nil, false
	}
	return o.Iccid, true
}

// HasIccid returns a boolean if a field has been set.
func (o *Network) HasIccid() bool {
	if o != nil && o.Iccid != nil {
		return true
	}

	return false
}

// SetIccid gets a reference to the given string and assigns it to the Iccid field.
func (o *Network) SetIccid(v string) {
	o.Iccid = &v
}

// GetMeid returns the Meid field value if set, zero value otherwise.
func (o *Network) GetMeid() string {
	if o == nil || o.Meid == nil {
		var ret string
		return ret
	}
	return *o.Meid
}

// GetMeidOk returns a tuple with the Meid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetMeidOk() (*string, bool) {
	if o == nil || o.Meid == nil {
		return nil, false
	}
	return o.Meid, true
}

// HasMeid returns a boolean if a field has been set.
func (o *Network) HasMeid() bool {
	if o != nil && o.Meid != nil {
		return true
	}

	return false
}

// SetMeid gets a reference to the given string and assigns it to the Meid field.
func (o *Network) SetMeid(v string) {
	o.Meid = &v
}

// GetCarrierSettingsVersion returns the CarrierSettingsVersion field value if set, zero value otherwise.
func (o *Network) GetCarrierSettingsVersion() string {
	if o == nil || o.CarrierSettingsVersion == nil {
		var ret string
		return ret
	}
	return *o.CarrierSettingsVersion
}

// GetCarrierSettingsVersionOk returns a tuple with the CarrierSettingsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetCarrierSettingsVersionOk() (*string, bool) {
	if o == nil || o.CarrierSettingsVersion == nil {
		return nil, false
	}
	return o.CarrierSettingsVersion, true
}

// HasCarrierSettingsVersion returns a boolean if a field has been set.
func (o *Network) HasCarrierSettingsVersion() bool {
	if o != nil && o.CarrierSettingsVersion != nil {
		return true
	}

	return false
}

// SetCarrierSettingsVersion gets a reference to the given string and assigns it to the CarrierSettingsVersion field.
func (o *Network) SetCarrierSettingsVersion(v string) {
	o.CarrierSettingsVersion = &v
}

// GetCurrentCarrierNetwork returns the CurrentCarrierNetwork field value if set, zero value otherwise.
func (o *Network) GetCurrentCarrierNetwork() string {
	if o == nil || o.CurrentCarrierNetwork == nil {
		var ret string
		return ret
	}
	return *o.CurrentCarrierNetwork
}

// GetCurrentCarrierNetworkOk returns a tuple with the CurrentCarrierNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetCurrentCarrierNetworkOk() (*string, bool) {
	if o == nil || o.CurrentCarrierNetwork == nil {
		return nil, false
	}
	return o.CurrentCarrierNetwork, true
}

// HasCurrentCarrierNetwork returns a boolean if a field has been set.
func (o *Network) HasCurrentCarrierNetwork() bool {
	if o != nil && o.CurrentCarrierNetwork != nil {
		return true
	}

	return false
}

// SetCurrentCarrierNetwork gets a reference to the given string and assigns it to the CurrentCarrierNetwork field.
func (o *Network) SetCurrentCarrierNetwork(v string) {
	o.CurrentCarrierNetwork = &v
}

// GetCurrentMobileCountryCode returns the CurrentMobileCountryCode field value if set, zero value otherwise.
func (o *Network) GetCurrentMobileCountryCode() string {
	if o == nil || o.CurrentMobileCountryCode == nil {
		var ret string
		return ret
	}
	return *o.CurrentMobileCountryCode
}

// GetCurrentMobileCountryCodeOk returns a tuple with the CurrentMobileCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetCurrentMobileCountryCodeOk() (*string, bool) {
	if o == nil || o.CurrentMobileCountryCode == nil {
		return nil, false
	}
	return o.CurrentMobileCountryCode, true
}

// HasCurrentMobileCountryCode returns a boolean if a field has been set.
func (o *Network) HasCurrentMobileCountryCode() bool {
	if o != nil && o.CurrentMobileCountryCode != nil {
		return true
	}

	return false
}

// SetCurrentMobileCountryCode gets a reference to the given string and assigns it to the CurrentMobileCountryCode field.
func (o *Network) SetCurrentMobileCountryCode(v string) {
	o.CurrentMobileCountryCode = &v
}

// GetCurrentMobileNetworkCode returns the CurrentMobileNetworkCode field value if set, zero value otherwise.
func (o *Network) GetCurrentMobileNetworkCode() string {
	if o == nil || o.CurrentMobileNetworkCode == nil {
		var ret string
		return ret
	}
	return *o.CurrentMobileNetworkCode
}

// GetCurrentMobileNetworkCodeOk returns a tuple with the CurrentMobileNetworkCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetCurrentMobileNetworkCodeOk() (*string, bool) {
	if o == nil || o.CurrentMobileNetworkCode == nil {
		return nil, false
	}
	return o.CurrentMobileNetworkCode, true
}

// HasCurrentMobileNetworkCode returns a boolean if a field has been set.
func (o *Network) HasCurrentMobileNetworkCode() bool {
	if o != nil && o.CurrentMobileNetworkCode != nil {
		return true
	}

	return false
}

// SetCurrentMobileNetworkCode gets a reference to the given string and assigns it to the CurrentMobileNetworkCode field.
func (o *Network) SetCurrentMobileNetworkCode(v string) {
	o.CurrentMobileNetworkCode = &v
}

// GetHomeCarrierNetwork returns the HomeCarrierNetwork field value if set, zero value otherwise.
func (o *Network) GetHomeCarrierNetwork() string {
	if o == nil || o.HomeCarrierNetwork == nil {
		var ret string
		return ret
	}
	return *o.HomeCarrierNetwork
}

// GetHomeCarrierNetworkOk returns a tuple with the HomeCarrierNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetHomeCarrierNetworkOk() (*string, bool) {
	if o == nil || o.HomeCarrierNetwork == nil {
		return nil, false
	}
	return o.HomeCarrierNetwork, true
}

// HasHomeCarrierNetwork returns a boolean if a field has been set.
func (o *Network) HasHomeCarrierNetwork() bool {
	if o != nil && o.HomeCarrierNetwork != nil {
		return true
	}

	return false
}

// SetHomeCarrierNetwork gets a reference to the given string and assigns it to the HomeCarrierNetwork field.
func (o *Network) SetHomeCarrierNetwork(v string) {
	o.HomeCarrierNetwork = &v
}

// GetHomeMobileCountryCode returns the HomeMobileCountryCode field value if set, zero value otherwise.
func (o *Network) GetHomeMobileCountryCode() string {
	if o == nil || o.HomeMobileCountryCode == nil {
		var ret string
		return ret
	}
	return *o.HomeMobileCountryCode
}

// GetHomeMobileCountryCodeOk returns a tuple with the HomeMobileCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetHomeMobileCountryCodeOk() (*string, bool) {
	if o == nil || o.HomeMobileCountryCode == nil {
		return nil, false
	}
	return o.HomeMobileCountryCode, true
}

// HasHomeMobileCountryCode returns a boolean if a field has been set.
func (o *Network) HasHomeMobileCountryCode() bool {
	if o != nil && o.HomeMobileCountryCode != nil {
		return true
	}

	return false
}

// SetHomeMobileCountryCode gets a reference to the given string and assigns it to the HomeMobileCountryCode field.
func (o *Network) SetHomeMobileCountryCode(v string) {
	o.HomeMobileCountryCode = &v
}

// GetHomeMobileNetworkCode returns the HomeMobileNetworkCode field value if set, zero value otherwise.
func (o *Network) GetHomeMobileNetworkCode() string {
	if o == nil || o.HomeMobileNetworkCode == nil {
		var ret string
		return ret
	}
	return *o.HomeMobileNetworkCode
}

// GetHomeMobileNetworkCodeOk returns a tuple with the HomeMobileNetworkCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetHomeMobileNetworkCodeOk() (*string, bool) {
	if o == nil || o.HomeMobileNetworkCode == nil {
		return nil, false
	}
	return o.HomeMobileNetworkCode, true
}

// HasHomeMobileNetworkCode returns a boolean if a field has been set.
func (o *Network) HasHomeMobileNetworkCode() bool {
	if o != nil && o.HomeMobileNetworkCode != nil {
		return true
	}

	return false
}

// SetHomeMobileNetworkCode gets a reference to the given string and assigns it to the HomeMobileNetworkCode field.
func (o *Network) SetHomeMobileNetworkCode(v string) {
	o.HomeMobileNetworkCode = &v
}

// GetIsDataRoamingEnabled returns the IsDataRoamingEnabled field value if set, zero value otherwise.
func (o *Network) GetIsDataRoamingEnabled() bool {
	if o == nil || o.IsDataRoamingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsDataRoamingEnabled
}

// GetIsDataRoamingEnabledOk returns a tuple with the IsDataRoamingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetIsDataRoamingEnabledOk() (*bool, bool) {
	if o == nil || o.IsDataRoamingEnabled == nil {
		return nil, false
	}
	return o.IsDataRoamingEnabled, true
}

// HasIsDataRoamingEnabled returns a boolean if a field has been set.
func (o *Network) HasIsDataRoamingEnabled() bool {
	if o != nil && o.IsDataRoamingEnabled != nil {
		return true
	}

	return false
}

// SetIsDataRoamingEnabled gets a reference to the given bool and assigns it to the IsDataRoamingEnabled field.
func (o *Network) SetIsDataRoamingEnabled(v bool) {
	o.IsDataRoamingEnabled = &v
}

// GetIsRoaming returns the IsRoaming field value if set, zero value otherwise.
func (o *Network) GetIsRoaming() bool {
	if o == nil || o.IsRoaming == nil {
		var ret bool
		return ret
	}
	return *o.IsRoaming
}

// GetIsRoamingOk returns a tuple with the IsRoaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetIsRoamingOk() (*bool, bool) {
	if o == nil || o.IsRoaming == nil {
		return nil, false
	}
	return o.IsRoaming, true
}

// HasIsRoaming returns a boolean if a field has been set.
func (o *Network) HasIsRoaming() bool {
	if o != nil && o.IsRoaming != nil {
		return true
	}

	return false
}

// SetIsRoaming gets a reference to the given bool and assigns it to the IsRoaming field.
func (o *Network) SetIsRoaming(v bool) {
	o.IsRoaming = &v
}

// GetIsPersonalHotspotEnabled returns the IsPersonalHotspotEnabled field value if set, zero value otherwise.
func (o *Network) GetIsPersonalHotspotEnabled() bool {
	if o == nil || o.IsPersonalHotspotEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsPersonalHotspotEnabled
}

// GetIsPersonalHotspotEnabledOk returns a tuple with the IsPersonalHotspotEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetIsPersonalHotspotEnabledOk() (*bool, bool) {
	if o == nil || o.IsPersonalHotspotEnabled == nil {
		return nil, false
	}
	return o.IsPersonalHotspotEnabled, true
}

// HasIsPersonalHotspotEnabled returns a boolean if a field has been set.
func (o *Network) HasIsPersonalHotspotEnabled() bool {
	if o != nil && o.IsPersonalHotspotEnabled != nil {
		return true
	}

	return false
}

// SetIsPersonalHotspotEnabled gets a reference to the given bool and assigns it to the IsPersonalHotspotEnabled field.
func (o *Network) SetIsPersonalHotspotEnabled(v bool) {
	o.IsPersonalHotspotEnabled = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *Network) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *Network) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *Network) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o Network) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CellularTechnology != nil {
		toSerialize["cellularTechnology"] = o.CellularTechnology
	}
	if o.IsVoiceRoamingEnabled != nil {
		toSerialize["isVoiceRoamingEnabled"] = o.IsVoiceRoamingEnabled
	}
	if o.Imei != nil {
		toSerialize["imei"] = o.Imei
	}
	if o.Iccid != nil {
		toSerialize["iccid"] = o.Iccid
	}
	if o.Meid != nil {
		toSerialize["meid"] = o.Meid
	}
	if o.CarrierSettingsVersion != nil {
		toSerialize["carrierSettingsVersion"] = o.CarrierSettingsVersion
	}
	if o.CurrentCarrierNetwork != nil {
		toSerialize["currentCarrierNetwork"] = o.CurrentCarrierNetwork
	}
	if o.CurrentMobileCountryCode != nil {
		toSerialize["currentMobileCountryCode"] = o.CurrentMobileCountryCode
	}
	if o.CurrentMobileNetworkCode != nil {
		toSerialize["currentMobileNetworkCode"] = o.CurrentMobileNetworkCode
	}
	if o.HomeCarrierNetwork != nil {
		toSerialize["homeCarrierNetwork"] = o.HomeCarrierNetwork
	}
	if o.HomeMobileCountryCode != nil {
		toSerialize["homeMobileCountryCode"] = o.HomeMobileCountryCode
	}
	if o.HomeMobileNetworkCode != nil {
		toSerialize["homeMobileNetworkCode"] = o.HomeMobileNetworkCode
	}
	if o.IsDataRoamingEnabled != nil {
		toSerialize["isDataRoamingEnabled"] = o.IsDataRoamingEnabled
	}
	if o.IsRoaming != nil {
		toSerialize["isRoaming"] = o.IsRoaming
	}
	if o.IsPersonalHotspotEnabled != nil {
		toSerialize["isPersonalHotspotEnabled"] = o.IsPersonalHotspotEnabled
	}
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	return json.Marshal(toSerialize)
}

type NullableNetwork struct {
	value *Network
	isSet bool
}

func (v NullableNetwork) Get() *Network {
	return v.value
}

func (v *NullableNetwork) Set(val *Network) {
	v.value = val
	v.isSet = true
}

func (v NullableNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetwork(val *Network) *NullableNetwork {
	return &NullableNetwork{value: val, isSet: true}
}

func (v NullableNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


