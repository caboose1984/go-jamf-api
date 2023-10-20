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

// checks if the MobileDeviceIosInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MobileDeviceIosInventory{}

// MobileDeviceIosInventory struct for MobileDeviceIosInventory
type MobileDeviceIosInventory struct {
	MobileDeviceId *string `json:"mobileDeviceId,omitempty"`
	// Based on the value of this type either ios or appleTv objects will be populated.
	DeviceType string `json:"deviceType"`
	Hardware *MobileDeviceHardware `json:"hardware,omitempty"`
	UserAndLocation *MobileDeviceUserAndLocation `json:"userAndLocation,omitempty"`
	Purchasing *MobileDevicePurchasing `json:"purchasing,omitempty"`
	Applications []MobileDeviceApplicationInventoryDetail `json:"applications,omitempty"`
	Certificates []MobileDeviceCertificate `json:"certificates,omitempty"`
	Profiles []MobileDeviceProfile `json:"profiles,omitempty"`
	UserProfiles []MobileDeviceUserProfile `json:"userProfiles,omitempty"`
	ExtensionAttributes []MobileDeviceExtensionAttribute `json:"extensionAttributes,omitempty"`
	General *MobileDeviceIosGeneral `json:"general,omitempty"`
	Security *MobileDeviceSecurity `json:"security,omitempty"`
	Ebooks []MobileDeviceEbookInventoryDetail `json:"ebooks,omitempty"`
	Network *MobileDeviceNetwork `json:"network,omitempty"`
	ServiceSubscriptions []MobileDeviceServiceSubscriptions `json:"serviceSubscriptions,omitempty"`
	ProvisioningProfiles []MobileDeviceProvisioningProfiles `json:"provisioningProfiles,omitempty"`
	SharedUsers []MobileDeviceSharedUser `json:"sharedUsers,omitempty"`
}

// NewMobileDeviceIosInventory instantiates a new MobileDeviceIosInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileDeviceIosInventory(deviceType string) *MobileDeviceIosInventory {
	this := MobileDeviceIosInventory{}
	this.DeviceType = deviceType
	return &this
}

// NewMobileDeviceIosInventoryWithDefaults instantiates a new MobileDeviceIosInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileDeviceIosInventoryWithDefaults() *MobileDeviceIosInventory {
	this := MobileDeviceIosInventory{}
	return &this
}

// GetMobileDeviceId returns the MobileDeviceId field value if set, zero value otherwise.
func (o *MobileDeviceIosInventory) GetMobileDeviceId() string {
	if o == nil || IsNil(o.MobileDeviceId) {
		var ret string
		return ret
	}
	return *o.MobileDeviceId
}

// GetMobileDeviceIdOk returns a tuple with the MobileDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetMobileDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MobileDeviceId) {
		return nil, false
	}
	return o.MobileDeviceId, true
}

// HasMobileDeviceId returns a boolean if a field has been set.
func (o *MobileDeviceIosInventory) HasMobileDeviceId() bool {
	if o != nil && !IsNil(o.MobileDeviceId) {
		return true
	}

	return false
}

// SetMobileDeviceId gets a reference to the given string and assigns it to the MobileDeviceId field.
func (o *MobileDeviceIosInventory) SetMobileDeviceId(v string) {
	o.MobileDeviceId = &v
}

// GetDeviceType returns the DeviceType field value
func (o *MobileDeviceIosInventory) GetDeviceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetDeviceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *MobileDeviceIosInventory) SetDeviceType(v string) {
	o.DeviceType = v
}

// GetHardware returns the Hardware field value if set, zero value otherwise.
func (o *MobileDeviceIosInventory) GetHardware() MobileDeviceHardware {
	if o == nil || IsNil(o.Hardware) {
		var ret MobileDeviceHardware
		return ret
	}
	return *o.Hardware
}

// GetHardwareOk returns a tuple with the Hardware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetHardwareOk() (*MobileDeviceHardware, bool) {
	if o == nil || IsNil(o.Hardware) {
		return nil, false
	}
	return o.Hardware, true
}

// HasHardware returns a boolean if a field has been set.
func (o *MobileDeviceIosInventory) HasHardware() bool {
	if o != nil && !IsNil(o.Hardware) {
		return true
	}

	return false
}

// SetHardware gets a reference to the given MobileDeviceHardware and assigns it to the Hardware field.
func (o *MobileDeviceIosInventory) SetHardware(v MobileDeviceHardware) {
	o.Hardware = &v
}

// GetUserAndLocation returns the UserAndLocation field value if set, zero value otherwise.
func (o *MobileDeviceIosInventory) GetUserAndLocation() MobileDeviceUserAndLocation {
	if o == nil || IsNil(o.UserAndLocation) {
		var ret MobileDeviceUserAndLocation
		return ret
	}
	return *o.UserAndLocation
}

// GetUserAndLocationOk returns a tuple with the UserAndLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetUserAndLocationOk() (*MobileDeviceUserAndLocation, bool) {
	if o == nil || IsNil(o.UserAndLocation) {
		return nil, false
	}
	return o.UserAndLocation, true
}

// HasUserAndLocation returns a boolean if a field has been set.
func (o *MobileDeviceIosInventory) HasUserAndLocation() bool {
	if o != nil && !IsNil(o.UserAndLocation) {
		return true
	}

	return false
}

// SetUserAndLocation gets a reference to the given MobileDeviceUserAndLocation and assigns it to the UserAndLocation field.
func (o *MobileDeviceIosInventory) SetUserAndLocation(v MobileDeviceUserAndLocation) {
	o.UserAndLocation = &v
}

// GetPurchasing returns the Purchasing field value if set, zero value otherwise.
func (o *MobileDeviceIosInventory) GetPurchasing() MobileDevicePurchasing {
	if o == nil || IsNil(o.Purchasing) {
		var ret MobileDevicePurchasing
		return ret
	}
	return *o.Purchasing
}

// GetPurchasingOk returns a tuple with the Purchasing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetPurchasingOk() (*MobileDevicePurchasing, bool) {
	if o == nil || IsNil(o.Purchasing) {
		return nil, false
	}
	return o.Purchasing, true
}

// HasPurchasing returns a boolean if a field has been set.
func (o *MobileDeviceIosInventory) HasPurchasing() bool {
	if o != nil && !IsNil(o.Purchasing) {
		return true
	}

	return false
}

// SetPurchasing gets a reference to the given MobileDevicePurchasing and assigns it to the Purchasing field.
func (o *MobileDeviceIosInventory) SetPurchasing(v MobileDevicePurchasing) {
	o.Purchasing = &v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *MobileDeviceIosInventory) GetApplications() []MobileDeviceApplicationInventoryDetail {
	if o == nil || IsNil(o.Applications) {
		var ret []MobileDeviceApplicationInventoryDetail
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetApplicationsOk() ([]MobileDeviceApplicationInventoryDetail, bool) {
	if o == nil || IsNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *MobileDeviceIosInventory) HasApplications() bool {
	if o != nil && !IsNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []MobileDeviceApplicationInventoryDetail and assigns it to the Applications field.
func (o *MobileDeviceIosInventory) SetApplications(v []MobileDeviceApplicationInventoryDetail) {
	o.Applications = v
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *MobileDeviceIosInventory) GetCertificates() []MobileDeviceCertificate {
	if o == nil || IsNil(o.Certificates) {
		var ret []MobileDeviceCertificate
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetCertificatesOk() ([]MobileDeviceCertificate, bool) {
	if o == nil || IsNil(o.Certificates) {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *MobileDeviceIosInventory) HasCertificates() bool {
	if o != nil && !IsNil(o.Certificates) {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []MobileDeviceCertificate and assigns it to the Certificates field.
func (o *MobileDeviceIosInventory) SetCertificates(v []MobileDeviceCertificate) {
	o.Certificates = v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *MobileDeviceIosInventory) GetProfiles() []MobileDeviceProfile {
	if o == nil || IsNil(o.Profiles) {
		var ret []MobileDeviceProfile
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetProfilesOk() ([]MobileDeviceProfile, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *MobileDeviceIosInventory) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []MobileDeviceProfile and assigns it to the Profiles field.
func (o *MobileDeviceIosInventory) SetProfiles(v []MobileDeviceProfile) {
	o.Profiles = v
}

// GetUserProfiles returns the UserProfiles field value if set, zero value otherwise.
func (o *MobileDeviceIosInventory) GetUserProfiles() []MobileDeviceUserProfile {
	if o == nil || IsNil(o.UserProfiles) {
		var ret []MobileDeviceUserProfile
		return ret
	}
	return o.UserProfiles
}

// GetUserProfilesOk returns a tuple with the UserProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetUserProfilesOk() ([]MobileDeviceUserProfile, bool) {
	if o == nil || IsNil(o.UserProfiles) {
		return nil, false
	}
	return o.UserProfiles, true
}

// HasUserProfiles returns a boolean if a field has been set.
func (o *MobileDeviceIosInventory) HasUserProfiles() bool {
	if o != nil && !IsNil(o.UserProfiles) {
		return true
	}

	return false
}

// SetUserProfiles gets a reference to the given []MobileDeviceUserProfile and assigns it to the UserProfiles field.
func (o *MobileDeviceIosInventory) SetUserProfiles(v []MobileDeviceUserProfile) {
	o.UserProfiles = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *MobileDeviceIosInventory) GetExtensionAttributes() []MobileDeviceExtensionAttribute {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret []MobileDeviceExtensionAttribute
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetExtensionAttributesOk() ([]MobileDeviceExtensionAttribute, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *MobileDeviceIosInventory) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given []MobileDeviceExtensionAttribute and assigns it to the ExtensionAttributes field.
func (o *MobileDeviceIosInventory) SetExtensionAttributes(v []MobileDeviceExtensionAttribute) {
	o.ExtensionAttributes = v
}

// GetGeneral returns the General field value if set, zero value otherwise.
func (o *MobileDeviceIosInventory) GetGeneral() MobileDeviceIosGeneral {
	if o == nil || IsNil(o.General) {
		var ret MobileDeviceIosGeneral
		return ret
	}
	return *o.General
}

// GetGeneralOk returns a tuple with the General field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetGeneralOk() (*MobileDeviceIosGeneral, bool) {
	if o == nil || IsNil(o.General) {
		return nil, false
	}
	return o.General, true
}

// HasGeneral returns a boolean if a field has been set.
func (o *MobileDeviceIosInventory) HasGeneral() bool {
	if o != nil && !IsNil(o.General) {
		return true
	}

	return false
}

// SetGeneral gets a reference to the given MobileDeviceIosGeneral and assigns it to the General field.
func (o *MobileDeviceIosInventory) SetGeneral(v MobileDeviceIosGeneral) {
	o.General = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *MobileDeviceIosInventory) GetSecurity() MobileDeviceSecurity {
	if o == nil || IsNil(o.Security) {
		var ret MobileDeviceSecurity
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetSecurityOk() (*MobileDeviceSecurity, bool) {
	if o == nil || IsNil(o.Security) {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *MobileDeviceIosInventory) HasSecurity() bool {
	if o != nil && !IsNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given MobileDeviceSecurity and assigns it to the Security field.
func (o *MobileDeviceIosInventory) SetSecurity(v MobileDeviceSecurity) {
	o.Security = &v
}

// GetEbooks returns the Ebooks field value if set, zero value otherwise.
func (o *MobileDeviceIosInventory) GetEbooks() []MobileDeviceEbookInventoryDetail {
	if o == nil || IsNil(o.Ebooks) {
		var ret []MobileDeviceEbookInventoryDetail
		return ret
	}
	return o.Ebooks
}

// GetEbooksOk returns a tuple with the Ebooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetEbooksOk() ([]MobileDeviceEbookInventoryDetail, bool) {
	if o == nil || IsNil(o.Ebooks) {
		return nil, false
	}
	return o.Ebooks, true
}

// HasEbooks returns a boolean if a field has been set.
func (o *MobileDeviceIosInventory) HasEbooks() bool {
	if o != nil && !IsNil(o.Ebooks) {
		return true
	}

	return false
}

// SetEbooks gets a reference to the given []MobileDeviceEbookInventoryDetail and assigns it to the Ebooks field.
func (o *MobileDeviceIosInventory) SetEbooks(v []MobileDeviceEbookInventoryDetail) {
	o.Ebooks = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *MobileDeviceIosInventory) GetNetwork() MobileDeviceNetwork {
	if o == nil || IsNil(o.Network) {
		var ret MobileDeviceNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetNetworkOk() (*MobileDeviceNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *MobileDeviceIosInventory) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given MobileDeviceNetwork and assigns it to the Network field.
func (o *MobileDeviceIosInventory) SetNetwork(v MobileDeviceNetwork) {
	o.Network = &v
}

// GetServiceSubscriptions returns the ServiceSubscriptions field value if set, zero value otherwise.
func (o *MobileDeviceIosInventory) GetServiceSubscriptions() []MobileDeviceServiceSubscriptions {
	if o == nil || IsNil(o.ServiceSubscriptions) {
		var ret []MobileDeviceServiceSubscriptions
		return ret
	}
	return o.ServiceSubscriptions
}

// GetServiceSubscriptionsOk returns a tuple with the ServiceSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetServiceSubscriptionsOk() ([]MobileDeviceServiceSubscriptions, bool) {
	if o == nil || IsNil(o.ServiceSubscriptions) {
		return nil, false
	}
	return o.ServiceSubscriptions, true
}

// HasServiceSubscriptions returns a boolean if a field has been set.
func (o *MobileDeviceIosInventory) HasServiceSubscriptions() bool {
	if o != nil && !IsNil(o.ServiceSubscriptions) {
		return true
	}

	return false
}

// SetServiceSubscriptions gets a reference to the given []MobileDeviceServiceSubscriptions and assigns it to the ServiceSubscriptions field.
func (o *MobileDeviceIosInventory) SetServiceSubscriptions(v []MobileDeviceServiceSubscriptions) {
	o.ServiceSubscriptions = v
}

// GetProvisioningProfiles returns the ProvisioningProfiles field value if set, zero value otherwise.
func (o *MobileDeviceIosInventory) GetProvisioningProfiles() []MobileDeviceProvisioningProfiles {
	if o == nil || IsNil(o.ProvisioningProfiles) {
		var ret []MobileDeviceProvisioningProfiles
		return ret
	}
	return o.ProvisioningProfiles
}

// GetProvisioningProfilesOk returns a tuple with the ProvisioningProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetProvisioningProfilesOk() ([]MobileDeviceProvisioningProfiles, bool) {
	if o == nil || IsNil(o.ProvisioningProfiles) {
		return nil, false
	}
	return o.ProvisioningProfiles, true
}

// HasProvisioningProfiles returns a boolean if a field has been set.
func (o *MobileDeviceIosInventory) HasProvisioningProfiles() bool {
	if o != nil && !IsNil(o.ProvisioningProfiles) {
		return true
	}

	return false
}

// SetProvisioningProfiles gets a reference to the given []MobileDeviceProvisioningProfiles and assigns it to the ProvisioningProfiles field.
func (o *MobileDeviceIosInventory) SetProvisioningProfiles(v []MobileDeviceProvisioningProfiles) {
	o.ProvisioningProfiles = v
}

// GetSharedUsers returns the SharedUsers field value if set, zero value otherwise.
func (o *MobileDeviceIosInventory) GetSharedUsers() []MobileDeviceSharedUser {
	if o == nil || IsNil(o.SharedUsers) {
		var ret []MobileDeviceSharedUser
		return ret
	}
	return o.SharedUsers
}

// GetSharedUsersOk returns a tuple with the SharedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceIosInventory) GetSharedUsersOk() ([]MobileDeviceSharedUser, bool) {
	if o == nil || IsNil(o.SharedUsers) {
		return nil, false
	}
	return o.SharedUsers, true
}

// HasSharedUsers returns a boolean if a field has been set.
func (o *MobileDeviceIosInventory) HasSharedUsers() bool {
	if o != nil && !IsNil(o.SharedUsers) {
		return true
	}

	return false
}

// SetSharedUsers gets a reference to the given []MobileDeviceSharedUser and assigns it to the SharedUsers field.
func (o *MobileDeviceIosInventory) SetSharedUsers(v []MobileDeviceSharedUser) {
	o.SharedUsers = v
}

func (o MobileDeviceIosInventory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobileDeviceIosInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MobileDeviceId) {
		toSerialize["mobileDeviceId"] = o.MobileDeviceId
	}
	toSerialize["deviceType"] = o.DeviceType
	if !IsNil(o.Hardware) {
		toSerialize["hardware"] = o.Hardware
	}
	if !IsNil(o.UserAndLocation) {
		toSerialize["userAndLocation"] = o.UserAndLocation
	}
	if !IsNil(o.Purchasing) {
		toSerialize["purchasing"] = o.Purchasing
	}
	if !IsNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	if !IsNil(o.Certificates) {
		toSerialize["certificates"] = o.Certificates
	}
	if !IsNil(o.Profiles) {
		toSerialize["profiles"] = o.Profiles
	}
	if !IsNil(o.UserProfiles) {
		toSerialize["userProfiles"] = o.UserProfiles
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extensionAttributes"] = o.ExtensionAttributes
	}
	if !IsNil(o.General) {
		toSerialize["general"] = o.General
	}
	if !IsNil(o.Security) {
		toSerialize["security"] = o.Security
	}
	if !IsNil(o.Ebooks) {
		toSerialize["ebooks"] = o.Ebooks
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.ServiceSubscriptions) {
		toSerialize["serviceSubscriptions"] = o.ServiceSubscriptions
	}
	if !IsNil(o.ProvisioningProfiles) {
		toSerialize["provisioningProfiles"] = o.ProvisioningProfiles
	}
	if !IsNil(o.SharedUsers) {
		toSerialize["sharedUsers"] = o.SharedUsers
	}
	return toSerialize, nil
}

type NullableMobileDeviceIosInventory struct {
	value *MobileDeviceIosInventory
	isSet bool
}

func (v NullableMobileDeviceIosInventory) Get() *MobileDeviceIosInventory {
	return v.value
}

func (v *NullableMobileDeviceIosInventory) Set(val *MobileDeviceIosInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDeviceIosInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDeviceIosInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDeviceIosInventory(val *MobileDeviceIosInventory) *NullableMobileDeviceIosInventory {
	return &NullableMobileDeviceIosInventory{value: val, isSet: true}
}

func (v NullableMobileDeviceIosInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDeviceIosInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


