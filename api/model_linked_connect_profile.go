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

// LinkedConnectProfile struct for LinkedConnectProfile
type LinkedConnectProfile struct {
	Uuid *string `json:"uuid,omitempty"`
	ProfileId *string `json:"profileId,omitempty"`
	ProfileName *string `json:"profileName,omitempty"`
	ProfileScopeDescription *string `json:"profileScopeDescription,omitempty"`
	// Must be a valid Jamf Connect version 2.3.0 or higher. Versions are listed here `https://www.jamf.com/resources/product-documentation/jamf-connect-administrators-guide/`
	Version *string `json:"version,omitempty"`
	// Determines how the server will behave regarding application updates and installs on the devices that have the configuration profile installed. * `PATCH_UPDATES` - Server handles initial installation of the application and any patch updates. * `MINOR_AND_PATCH_UPDATES` - Server handles initial installation of the application and any patch and minor updates. * `INITIAL_INSTALLATION_ONLY` - Server only handles initial installation of the application. Updates will have to be done manually. * `NONE` - Server does not handle any installations or updates for the application. Version is ignored for this type. 
	AutoDeploymentType *string `json:"autoDeploymentType,omitempty"`
}

// NewLinkedConnectProfile instantiates a new LinkedConnectProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedConnectProfile() *LinkedConnectProfile {
	this := LinkedConnectProfile{}
	var autoDeploymentType string = "NONE"
	this.AutoDeploymentType = &autoDeploymentType
	return &this
}

// NewLinkedConnectProfileWithDefaults instantiates a new LinkedConnectProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedConnectProfileWithDefaults() *LinkedConnectProfile {
	this := LinkedConnectProfile{}
	var autoDeploymentType string = "NONE"
	this.AutoDeploymentType = &autoDeploymentType
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *LinkedConnectProfile) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedConnectProfile) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *LinkedConnectProfile) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *LinkedConnectProfile) SetUuid(v string) {
	o.Uuid = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *LinkedConnectProfile) GetProfileId() string {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedConnectProfile) GetProfileIdOk() (*string, bool) {
	if o == nil || o.ProfileId == nil {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *LinkedConnectProfile) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *LinkedConnectProfile) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetProfileName returns the ProfileName field value if set, zero value otherwise.
func (o *LinkedConnectProfile) GetProfileName() string {
	if o == nil || o.ProfileName == nil {
		var ret string
		return ret
	}
	return *o.ProfileName
}

// GetProfileNameOk returns a tuple with the ProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedConnectProfile) GetProfileNameOk() (*string, bool) {
	if o == nil || o.ProfileName == nil {
		return nil, false
	}
	return o.ProfileName, true
}

// HasProfileName returns a boolean if a field has been set.
func (o *LinkedConnectProfile) HasProfileName() bool {
	if o != nil && o.ProfileName != nil {
		return true
	}

	return false
}

// SetProfileName gets a reference to the given string and assigns it to the ProfileName field.
func (o *LinkedConnectProfile) SetProfileName(v string) {
	o.ProfileName = &v
}

// GetProfileScopeDescription returns the ProfileScopeDescription field value if set, zero value otherwise.
func (o *LinkedConnectProfile) GetProfileScopeDescription() string {
	if o == nil || o.ProfileScopeDescription == nil {
		var ret string
		return ret
	}
	return *o.ProfileScopeDescription
}

// GetProfileScopeDescriptionOk returns a tuple with the ProfileScopeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedConnectProfile) GetProfileScopeDescriptionOk() (*string, bool) {
	if o == nil || o.ProfileScopeDescription == nil {
		return nil, false
	}
	return o.ProfileScopeDescription, true
}

// HasProfileScopeDescription returns a boolean if a field has been set.
func (o *LinkedConnectProfile) HasProfileScopeDescription() bool {
	if o != nil && o.ProfileScopeDescription != nil {
		return true
	}

	return false
}

// SetProfileScopeDescription gets a reference to the given string and assigns it to the ProfileScopeDescription field.
func (o *LinkedConnectProfile) SetProfileScopeDescription(v string) {
	o.ProfileScopeDescription = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LinkedConnectProfile) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedConnectProfile) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LinkedConnectProfile) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LinkedConnectProfile) SetVersion(v string) {
	o.Version = &v
}

// GetAutoDeploymentType returns the AutoDeploymentType field value if set, zero value otherwise.
func (o *LinkedConnectProfile) GetAutoDeploymentType() string {
	if o == nil || o.AutoDeploymentType == nil {
		var ret string
		return ret
	}
	return *o.AutoDeploymentType
}

// GetAutoDeploymentTypeOk returns a tuple with the AutoDeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedConnectProfile) GetAutoDeploymentTypeOk() (*string, bool) {
	if o == nil || o.AutoDeploymentType == nil {
		return nil, false
	}
	return o.AutoDeploymentType, true
}

// HasAutoDeploymentType returns a boolean if a field has been set.
func (o *LinkedConnectProfile) HasAutoDeploymentType() bool {
	if o != nil && o.AutoDeploymentType != nil {
		return true
	}

	return false
}

// SetAutoDeploymentType gets a reference to the given string and assigns it to the AutoDeploymentType field.
func (o *LinkedConnectProfile) SetAutoDeploymentType(v string) {
	o.AutoDeploymentType = &v
}

func (o LinkedConnectProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.ProfileId != nil {
		toSerialize["profileId"] = o.ProfileId
	}
	if o.ProfileName != nil {
		toSerialize["profileName"] = o.ProfileName
	}
	if o.ProfileScopeDescription != nil {
		toSerialize["profileScopeDescription"] = o.ProfileScopeDescription
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.AutoDeploymentType != nil {
		toSerialize["autoDeploymentType"] = o.AutoDeploymentType
	}
	return json.Marshal(toSerialize)
}

type NullableLinkedConnectProfile struct {
	value *LinkedConnectProfile
	isSet bool
}

func (v NullableLinkedConnectProfile) Get() *LinkedConnectProfile {
	return v.value
}

func (v *NullableLinkedConnectProfile) Set(val *LinkedConnectProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedConnectProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedConnectProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedConnectProfile(val *LinkedConnectProfile) *NullableLinkedConnectProfile {
	return &NullableLinkedConnectProfile{value: val, isSet: true}
}

func (v NullableLinkedConnectProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedConnectProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

