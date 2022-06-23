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

// AuthAccount struct for AuthAccount
type AuthAccount struct {
	Id *int32 `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
	RealName *string `json:"realName,omitempty"`
	Email *string `json:"email,omitempty"`
	Preferences *AccountPreferences `json:"preferences,omitempty"`
	IsMultiSiteAdmin *bool `json:"isMultiSiteAdmin,omitempty"`
	AccessLevel *string `json:"accessLevel,omitempty"`
	PrivilegeSet *string `json:"privilegeSet,omitempty"`
	PrivilegesBySite *map[string][]string `json:"privilegesBySite,omitempty"`
	GroupIds []int32 `json:"groupIds,omitempty"`
	CurrentSiteId *int32 `json:"currentSiteId,omitempty"`
}

// NewAuthAccount instantiates a new AuthAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAccount() *AuthAccount {
	this := AuthAccount{}
	return &this
}

// NewAuthAccountWithDefaults instantiates a new AuthAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAccountWithDefaults() *AuthAccount {
	this := AuthAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthAccount) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccount) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthAccount) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AuthAccount) SetId(v int32) {
	o.Id = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *AuthAccount) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccount) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *AuthAccount) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *AuthAccount) SetUsername(v string) {
	o.Username = &v
}

// GetRealName returns the RealName field value if set, zero value otherwise.
func (o *AuthAccount) GetRealName() string {
	if o == nil || o.RealName == nil {
		var ret string
		return ret
	}
	return *o.RealName
}

// GetRealNameOk returns a tuple with the RealName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccount) GetRealNameOk() (*string, bool) {
	if o == nil || o.RealName == nil {
		return nil, false
	}
	return o.RealName, true
}

// HasRealName returns a boolean if a field has been set.
func (o *AuthAccount) HasRealName() bool {
	if o != nil && o.RealName != nil {
		return true
	}

	return false
}

// SetRealName gets a reference to the given string and assigns it to the RealName field.
func (o *AuthAccount) SetRealName(v string) {
	o.RealName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AuthAccount) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccount) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AuthAccount) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AuthAccount) SetEmail(v string) {
	o.Email = &v
}

// GetPreferences returns the Preferences field value if set, zero value otherwise.
func (o *AuthAccount) GetPreferences() AccountPreferences {
	if o == nil || o.Preferences == nil {
		var ret AccountPreferences
		return ret
	}
	return *o.Preferences
}

// GetPreferencesOk returns a tuple with the Preferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccount) GetPreferencesOk() (*AccountPreferences, bool) {
	if o == nil || o.Preferences == nil {
		return nil, false
	}
	return o.Preferences, true
}

// HasPreferences returns a boolean if a field has been set.
func (o *AuthAccount) HasPreferences() bool {
	if o != nil && o.Preferences != nil {
		return true
	}

	return false
}

// SetPreferences gets a reference to the given AccountPreferences and assigns it to the Preferences field.
func (o *AuthAccount) SetPreferences(v AccountPreferences) {
	o.Preferences = &v
}

// GetIsMultiSiteAdmin returns the IsMultiSiteAdmin field value if set, zero value otherwise.
func (o *AuthAccount) GetIsMultiSiteAdmin() bool {
	if o == nil || o.IsMultiSiteAdmin == nil {
		var ret bool
		return ret
	}
	return *o.IsMultiSiteAdmin
}

// GetIsMultiSiteAdminOk returns a tuple with the IsMultiSiteAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccount) GetIsMultiSiteAdminOk() (*bool, bool) {
	if o == nil || o.IsMultiSiteAdmin == nil {
		return nil, false
	}
	return o.IsMultiSiteAdmin, true
}

// HasIsMultiSiteAdmin returns a boolean if a field has been set.
func (o *AuthAccount) HasIsMultiSiteAdmin() bool {
	if o != nil && o.IsMultiSiteAdmin != nil {
		return true
	}

	return false
}

// SetIsMultiSiteAdmin gets a reference to the given bool and assigns it to the IsMultiSiteAdmin field.
func (o *AuthAccount) SetIsMultiSiteAdmin(v bool) {
	o.IsMultiSiteAdmin = &v
}

// GetAccessLevel returns the AccessLevel field value if set, zero value otherwise.
func (o *AuthAccount) GetAccessLevel() string {
	if o == nil || o.AccessLevel == nil {
		var ret string
		return ret
	}
	return *o.AccessLevel
}

// GetAccessLevelOk returns a tuple with the AccessLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccount) GetAccessLevelOk() (*string, bool) {
	if o == nil || o.AccessLevel == nil {
		return nil, false
	}
	return o.AccessLevel, true
}

// HasAccessLevel returns a boolean if a field has been set.
func (o *AuthAccount) HasAccessLevel() bool {
	if o != nil && o.AccessLevel != nil {
		return true
	}

	return false
}

// SetAccessLevel gets a reference to the given string and assigns it to the AccessLevel field.
func (o *AuthAccount) SetAccessLevel(v string) {
	o.AccessLevel = &v
}

// GetPrivilegeSet returns the PrivilegeSet field value if set, zero value otherwise.
func (o *AuthAccount) GetPrivilegeSet() string {
	if o == nil || o.PrivilegeSet == nil {
		var ret string
		return ret
	}
	return *o.PrivilegeSet
}

// GetPrivilegeSetOk returns a tuple with the PrivilegeSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccount) GetPrivilegeSetOk() (*string, bool) {
	if o == nil || o.PrivilegeSet == nil {
		return nil, false
	}
	return o.PrivilegeSet, true
}

// HasPrivilegeSet returns a boolean if a field has been set.
func (o *AuthAccount) HasPrivilegeSet() bool {
	if o != nil && o.PrivilegeSet != nil {
		return true
	}

	return false
}

// SetPrivilegeSet gets a reference to the given string and assigns it to the PrivilegeSet field.
func (o *AuthAccount) SetPrivilegeSet(v string) {
	o.PrivilegeSet = &v
}

// GetPrivilegesBySite returns the PrivilegesBySite field value if set, zero value otherwise.
func (o *AuthAccount) GetPrivilegesBySite() map[string][]string {
	if o == nil || o.PrivilegesBySite == nil {
		var ret map[string][]string
		return ret
	}
	return *o.PrivilegesBySite
}

// GetPrivilegesBySiteOk returns a tuple with the PrivilegesBySite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccount) GetPrivilegesBySiteOk() (*map[string][]string, bool) {
	if o == nil || o.PrivilegesBySite == nil {
		return nil, false
	}
	return o.PrivilegesBySite, true
}

// HasPrivilegesBySite returns a boolean if a field has been set.
func (o *AuthAccount) HasPrivilegesBySite() bool {
	if o != nil && o.PrivilegesBySite != nil {
		return true
	}

	return false
}

// SetPrivilegesBySite gets a reference to the given map[string][]string and assigns it to the PrivilegesBySite field.
func (o *AuthAccount) SetPrivilegesBySite(v map[string][]string) {
	o.PrivilegesBySite = &v
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *AuthAccount) GetGroupIds() []int32 {
	if o == nil || o.GroupIds == nil {
		var ret []int32
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccount) GetGroupIdsOk() ([]int32, bool) {
	if o == nil || o.GroupIds == nil {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *AuthAccount) HasGroupIds() bool {
	if o != nil && o.GroupIds != nil {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []int32 and assigns it to the GroupIds field.
func (o *AuthAccount) SetGroupIds(v []int32) {
	o.GroupIds = v
}

// GetCurrentSiteId returns the CurrentSiteId field value if set, zero value otherwise.
func (o *AuthAccount) GetCurrentSiteId() int32 {
	if o == nil || o.CurrentSiteId == nil {
		var ret int32
		return ret
	}
	return *o.CurrentSiteId
}

// GetCurrentSiteIdOk returns a tuple with the CurrentSiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccount) GetCurrentSiteIdOk() (*int32, bool) {
	if o == nil || o.CurrentSiteId == nil {
		return nil, false
	}
	return o.CurrentSiteId, true
}

// HasCurrentSiteId returns a boolean if a field has been set.
func (o *AuthAccount) HasCurrentSiteId() bool {
	if o != nil && o.CurrentSiteId != nil {
		return true
	}

	return false
}

// SetCurrentSiteId gets a reference to the given int32 and assigns it to the CurrentSiteId field.
func (o *AuthAccount) SetCurrentSiteId(v int32) {
	o.CurrentSiteId = &v
}

func (o AuthAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.RealName != nil {
		toSerialize["realName"] = o.RealName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Preferences != nil {
		toSerialize["preferences"] = o.Preferences
	}
	if o.IsMultiSiteAdmin != nil {
		toSerialize["isMultiSiteAdmin"] = o.IsMultiSiteAdmin
	}
	if o.AccessLevel != nil {
		toSerialize["accessLevel"] = o.AccessLevel
	}
	if o.PrivilegeSet != nil {
		toSerialize["privilegeSet"] = o.PrivilegeSet
	}
	if o.PrivilegesBySite != nil {
		toSerialize["privilegesBySite"] = o.PrivilegesBySite
	}
	if o.GroupIds != nil {
		toSerialize["groupIds"] = o.GroupIds
	}
	if o.CurrentSiteId != nil {
		toSerialize["currentSiteId"] = o.CurrentSiteId
	}
	return json.Marshal(toSerialize)
}

type NullableAuthAccount struct {
	value *AuthAccount
	isSet bool
}

func (v NullableAuthAccount) Get() *AuthAccount {
	return v.value
}

func (v *NullableAuthAccount) Set(val *AuthAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAccount(val *AuthAccount) *NullableAuthAccount {
	return &NullableAuthAccount{value: val, isSet: true}
}

func (v NullableAuthAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

