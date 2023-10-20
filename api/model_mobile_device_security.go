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

// checks if the MobileDeviceSecurity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MobileDeviceSecurity{}

// MobileDeviceSecurity This section only avaiable for Ios type.
type MobileDeviceSecurity struct {
	DataProtected *bool `json:"dataProtected,omitempty"`
	BlockLevelEncryptionCapable *bool `json:"blockLevelEncryptionCapable,omitempty"`
	FileLevelEncryptionCapable *bool `json:"fileLevelEncryptionCapable,omitempty"`
	PasscodePresent *bool `json:"passcodePresent,omitempty"`
	PasscodeCompliant *bool `json:"passcodeCompliant,omitempty"`
	PasscodeCompliantWithProfile *bool `json:"passcodeCompliantWithProfile,omitempty"`
	HardwareEncryption *int32 `json:"hardwareEncryption,omitempty"`
	ActivationLockEnabled *bool `json:"activationLockEnabled,omitempty"`
	JailBreakDetected *bool `json:"jailBreakDetected,omitempty"`
	PasscodeLockGracePeriodEnforcedSeconds *int32 `json:"passcodeLockGracePeriodEnforcedSeconds,omitempty"`
	PersonalDeviceProfileCurrent *bool `json:"personalDeviceProfileCurrent,omitempty"`
	LostModeEnabled *bool `json:"lostModeEnabled,omitempty"`
	LostModePersistent *bool `json:"lostModePersistent,omitempty"`
	LostModeMessage *string `json:"lostModeMessage,omitempty"`
	LostModePhoneNumber *string `json:"lostModePhoneNumber,omitempty"`
	LostModeFootnote *string `json:"lostModeFootnote,omitempty"`
	LostModeLocation *MobileDeviceLostModeLocation `json:"lostModeLocation,omitempty"`
}

// NewMobileDeviceSecurity instantiates a new MobileDeviceSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileDeviceSecurity() *MobileDeviceSecurity {
	this := MobileDeviceSecurity{}
	return &this
}

// NewMobileDeviceSecurityWithDefaults instantiates a new MobileDeviceSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileDeviceSecurityWithDefaults() *MobileDeviceSecurity {
	this := MobileDeviceSecurity{}
	return &this
}

// GetDataProtected returns the DataProtected field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetDataProtected() bool {
	if o == nil || IsNil(o.DataProtected) {
		var ret bool
		return ret
	}
	return *o.DataProtected
}

// GetDataProtectedOk returns a tuple with the DataProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetDataProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.DataProtected) {
		return nil, false
	}
	return o.DataProtected, true
}

// HasDataProtected returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasDataProtected() bool {
	if o != nil && !IsNil(o.DataProtected) {
		return true
	}

	return false
}

// SetDataProtected gets a reference to the given bool and assigns it to the DataProtected field.
func (o *MobileDeviceSecurity) SetDataProtected(v bool) {
	o.DataProtected = &v
}

// GetBlockLevelEncryptionCapable returns the BlockLevelEncryptionCapable field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetBlockLevelEncryptionCapable() bool {
	if o == nil || IsNil(o.BlockLevelEncryptionCapable) {
		var ret bool
		return ret
	}
	return *o.BlockLevelEncryptionCapable
}

// GetBlockLevelEncryptionCapableOk returns a tuple with the BlockLevelEncryptionCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetBlockLevelEncryptionCapableOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockLevelEncryptionCapable) {
		return nil, false
	}
	return o.BlockLevelEncryptionCapable, true
}

// HasBlockLevelEncryptionCapable returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasBlockLevelEncryptionCapable() bool {
	if o != nil && !IsNil(o.BlockLevelEncryptionCapable) {
		return true
	}

	return false
}

// SetBlockLevelEncryptionCapable gets a reference to the given bool and assigns it to the BlockLevelEncryptionCapable field.
func (o *MobileDeviceSecurity) SetBlockLevelEncryptionCapable(v bool) {
	o.BlockLevelEncryptionCapable = &v
}

// GetFileLevelEncryptionCapable returns the FileLevelEncryptionCapable field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetFileLevelEncryptionCapable() bool {
	if o == nil || IsNil(o.FileLevelEncryptionCapable) {
		var ret bool
		return ret
	}
	return *o.FileLevelEncryptionCapable
}

// GetFileLevelEncryptionCapableOk returns a tuple with the FileLevelEncryptionCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetFileLevelEncryptionCapableOk() (*bool, bool) {
	if o == nil || IsNil(o.FileLevelEncryptionCapable) {
		return nil, false
	}
	return o.FileLevelEncryptionCapable, true
}

// HasFileLevelEncryptionCapable returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasFileLevelEncryptionCapable() bool {
	if o != nil && !IsNil(o.FileLevelEncryptionCapable) {
		return true
	}

	return false
}

// SetFileLevelEncryptionCapable gets a reference to the given bool and assigns it to the FileLevelEncryptionCapable field.
func (o *MobileDeviceSecurity) SetFileLevelEncryptionCapable(v bool) {
	o.FileLevelEncryptionCapable = &v
}

// GetPasscodePresent returns the PasscodePresent field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetPasscodePresent() bool {
	if o == nil || IsNil(o.PasscodePresent) {
		var ret bool
		return ret
	}
	return *o.PasscodePresent
}

// GetPasscodePresentOk returns a tuple with the PasscodePresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetPasscodePresentOk() (*bool, bool) {
	if o == nil || IsNil(o.PasscodePresent) {
		return nil, false
	}
	return o.PasscodePresent, true
}

// HasPasscodePresent returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasPasscodePresent() bool {
	if o != nil && !IsNil(o.PasscodePresent) {
		return true
	}

	return false
}

// SetPasscodePresent gets a reference to the given bool and assigns it to the PasscodePresent field.
func (o *MobileDeviceSecurity) SetPasscodePresent(v bool) {
	o.PasscodePresent = &v
}

// GetPasscodeCompliant returns the PasscodeCompliant field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetPasscodeCompliant() bool {
	if o == nil || IsNil(o.PasscodeCompliant) {
		var ret bool
		return ret
	}
	return *o.PasscodeCompliant
}

// GetPasscodeCompliantOk returns a tuple with the PasscodeCompliant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetPasscodeCompliantOk() (*bool, bool) {
	if o == nil || IsNil(o.PasscodeCompliant) {
		return nil, false
	}
	return o.PasscodeCompliant, true
}

// HasPasscodeCompliant returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasPasscodeCompliant() bool {
	if o != nil && !IsNil(o.PasscodeCompliant) {
		return true
	}

	return false
}

// SetPasscodeCompliant gets a reference to the given bool and assigns it to the PasscodeCompliant field.
func (o *MobileDeviceSecurity) SetPasscodeCompliant(v bool) {
	o.PasscodeCompliant = &v
}

// GetPasscodeCompliantWithProfile returns the PasscodeCompliantWithProfile field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetPasscodeCompliantWithProfile() bool {
	if o == nil || IsNil(o.PasscodeCompliantWithProfile) {
		var ret bool
		return ret
	}
	return *o.PasscodeCompliantWithProfile
}

// GetPasscodeCompliantWithProfileOk returns a tuple with the PasscodeCompliantWithProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetPasscodeCompliantWithProfileOk() (*bool, bool) {
	if o == nil || IsNil(o.PasscodeCompliantWithProfile) {
		return nil, false
	}
	return o.PasscodeCompliantWithProfile, true
}

// HasPasscodeCompliantWithProfile returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasPasscodeCompliantWithProfile() bool {
	if o != nil && !IsNil(o.PasscodeCompliantWithProfile) {
		return true
	}

	return false
}

// SetPasscodeCompliantWithProfile gets a reference to the given bool and assigns it to the PasscodeCompliantWithProfile field.
func (o *MobileDeviceSecurity) SetPasscodeCompliantWithProfile(v bool) {
	o.PasscodeCompliantWithProfile = &v
}

// GetHardwareEncryption returns the HardwareEncryption field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetHardwareEncryption() int32 {
	if o == nil || IsNil(o.HardwareEncryption) {
		var ret int32
		return ret
	}
	return *o.HardwareEncryption
}

// GetHardwareEncryptionOk returns a tuple with the HardwareEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetHardwareEncryptionOk() (*int32, bool) {
	if o == nil || IsNil(o.HardwareEncryption) {
		return nil, false
	}
	return o.HardwareEncryption, true
}

// HasHardwareEncryption returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasHardwareEncryption() bool {
	if o != nil && !IsNil(o.HardwareEncryption) {
		return true
	}

	return false
}

// SetHardwareEncryption gets a reference to the given int32 and assigns it to the HardwareEncryption field.
func (o *MobileDeviceSecurity) SetHardwareEncryption(v int32) {
	o.HardwareEncryption = &v
}

// GetActivationLockEnabled returns the ActivationLockEnabled field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetActivationLockEnabled() bool {
	if o == nil || IsNil(o.ActivationLockEnabled) {
		var ret bool
		return ret
	}
	return *o.ActivationLockEnabled
}

// GetActivationLockEnabledOk returns a tuple with the ActivationLockEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetActivationLockEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ActivationLockEnabled) {
		return nil, false
	}
	return o.ActivationLockEnabled, true
}

// HasActivationLockEnabled returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasActivationLockEnabled() bool {
	if o != nil && !IsNil(o.ActivationLockEnabled) {
		return true
	}

	return false
}

// SetActivationLockEnabled gets a reference to the given bool and assigns it to the ActivationLockEnabled field.
func (o *MobileDeviceSecurity) SetActivationLockEnabled(v bool) {
	o.ActivationLockEnabled = &v
}

// GetJailBreakDetected returns the JailBreakDetected field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetJailBreakDetected() bool {
	if o == nil || IsNil(o.JailBreakDetected) {
		var ret bool
		return ret
	}
	return *o.JailBreakDetected
}

// GetJailBreakDetectedOk returns a tuple with the JailBreakDetected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetJailBreakDetectedOk() (*bool, bool) {
	if o == nil || IsNil(o.JailBreakDetected) {
		return nil, false
	}
	return o.JailBreakDetected, true
}

// HasJailBreakDetected returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasJailBreakDetected() bool {
	if o != nil && !IsNil(o.JailBreakDetected) {
		return true
	}

	return false
}

// SetJailBreakDetected gets a reference to the given bool and assigns it to the JailBreakDetected field.
func (o *MobileDeviceSecurity) SetJailBreakDetected(v bool) {
	o.JailBreakDetected = &v
}

// GetPasscodeLockGracePeriodEnforcedSeconds returns the PasscodeLockGracePeriodEnforcedSeconds field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetPasscodeLockGracePeriodEnforcedSeconds() int32 {
	if o == nil || IsNil(o.PasscodeLockGracePeriodEnforcedSeconds) {
		var ret int32
		return ret
	}
	return *o.PasscodeLockGracePeriodEnforcedSeconds
}

// GetPasscodeLockGracePeriodEnforcedSecondsOk returns a tuple with the PasscodeLockGracePeriodEnforcedSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetPasscodeLockGracePeriodEnforcedSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.PasscodeLockGracePeriodEnforcedSeconds) {
		return nil, false
	}
	return o.PasscodeLockGracePeriodEnforcedSeconds, true
}

// HasPasscodeLockGracePeriodEnforcedSeconds returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasPasscodeLockGracePeriodEnforcedSeconds() bool {
	if o != nil && !IsNil(o.PasscodeLockGracePeriodEnforcedSeconds) {
		return true
	}

	return false
}

// SetPasscodeLockGracePeriodEnforcedSeconds gets a reference to the given int32 and assigns it to the PasscodeLockGracePeriodEnforcedSeconds field.
func (o *MobileDeviceSecurity) SetPasscodeLockGracePeriodEnforcedSeconds(v int32) {
	o.PasscodeLockGracePeriodEnforcedSeconds = &v
}

// GetPersonalDeviceProfileCurrent returns the PersonalDeviceProfileCurrent field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetPersonalDeviceProfileCurrent() bool {
	if o == nil || IsNil(o.PersonalDeviceProfileCurrent) {
		var ret bool
		return ret
	}
	return *o.PersonalDeviceProfileCurrent
}

// GetPersonalDeviceProfileCurrentOk returns a tuple with the PersonalDeviceProfileCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetPersonalDeviceProfileCurrentOk() (*bool, bool) {
	if o == nil || IsNil(o.PersonalDeviceProfileCurrent) {
		return nil, false
	}
	return o.PersonalDeviceProfileCurrent, true
}

// HasPersonalDeviceProfileCurrent returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasPersonalDeviceProfileCurrent() bool {
	if o != nil && !IsNil(o.PersonalDeviceProfileCurrent) {
		return true
	}

	return false
}

// SetPersonalDeviceProfileCurrent gets a reference to the given bool and assigns it to the PersonalDeviceProfileCurrent field.
func (o *MobileDeviceSecurity) SetPersonalDeviceProfileCurrent(v bool) {
	o.PersonalDeviceProfileCurrent = &v
}

// GetLostModeEnabled returns the LostModeEnabled field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetLostModeEnabled() bool {
	if o == nil || IsNil(o.LostModeEnabled) {
		var ret bool
		return ret
	}
	return *o.LostModeEnabled
}

// GetLostModeEnabledOk returns a tuple with the LostModeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetLostModeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LostModeEnabled) {
		return nil, false
	}
	return o.LostModeEnabled, true
}

// HasLostModeEnabled returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasLostModeEnabled() bool {
	if o != nil && !IsNil(o.LostModeEnabled) {
		return true
	}

	return false
}

// SetLostModeEnabled gets a reference to the given bool and assigns it to the LostModeEnabled field.
func (o *MobileDeviceSecurity) SetLostModeEnabled(v bool) {
	o.LostModeEnabled = &v
}

// GetLostModePersistent returns the LostModePersistent field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetLostModePersistent() bool {
	if o == nil || IsNil(o.LostModePersistent) {
		var ret bool
		return ret
	}
	return *o.LostModePersistent
}

// GetLostModePersistentOk returns a tuple with the LostModePersistent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetLostModePersistentOk() (*bool, bool) {
	if o == nil || IsNil(o.LostModePersistent) {
		return nil, false
	}
	return o.LostModePersistent, true
}

// HasLostModePersistent returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasLostModePersistent() bool {
	if o != nil && !IsNil(o.LostModePersistent) {
		return true
	}

	return false
}

// SetLostModePersistent gets a reference to the given bool and assigns it to the LostModePersistent field.
func (o *MobileDeviceSecurity) SetLostModePersistent(v bool) {
	o.LostModePersistent = &v
}

// GetLostModeMessage returns the LostModeMessage field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetLostModeMessage() string {
	if o == nil || IsNil(o.LostModeMessage) {
		var ret string
		return ret
	}
	return *o.LostModeMessage
}

// GetLostModeMessageOk returns a tuple with the LostModeMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetLostModeMessageOk() (*string, bool) {
	if o == nil || IsNil(o.LostModeMessage) {
		return nil, false
	}
	return o.LostModeMessage, true
}

// HasLostModeMessage returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasLostModeMessage() bool {
	if o != nil && !IsNil(o.LostModeMessage) {
		return true
	}

	return false
}

// SetLostModeMessage gets a reference to the given string and assigns it to the LostModeMessage field.
func (o *MobileDeviceSecurity) SetLostModeMessage(v string) {
	o.LostModeMessage = &v
}

// GetLostModePhoneNumber returns the LostModePhoneNumber field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetLostModePhoneNumber() string {
	if o == nil || IsNil(o.LostModePhoneNumber) {
		var ret string
		return ret
	}
	return *o.LostModePhoneNumber
}

// GetLostModePhoneNumberOk returns a tuple with the LostModePhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetLostModePhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.LostModePhoneNumber) {
		return nil, false
	}
	return o.LostModePhoneNumber, true
}

// HasLostModePhoneNumber returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasLostModePhoneNumber() bool {
	if o != nil && !IsNil(o.LostModePhoneNumber) {
		return true
	}

	return false
}

// SetLostModePhoneNumber gets a reference to the given string and assigns it to the LostModePhoneNumber field.
func (o *MobileDeviceSecurity) SetLostModePhoneNumber(v string) {
	o.LostModePhoneNumber = &v
}

// GetLostModeFootnote returns the LostModeFootnote field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetLostModeFootnote() string {
	if o == nil || IsNil(o.LostModeFootnote) {
		var ret string
		return ret
	}
	return *o.LostModeFootnote
}

// GetLostModeFootnoteOk returns a tuple with the LostModeFootnote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetLostModeFootnoteOk() (*string, bool) {
	if o == nil || IsNil(o.LostModeFootnote) {
		return nil, false
	}
	return o.LostModeFootnote, true
}

// HasLostModeFootnote returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasLostModeFootnote() bool {
	if o != nil && !IsNil(o.LostModeFootnote) {
		return true
	}

	return false
}

// SetLostModeFootnote gets a reference to the given string and assigns it to the LostModeFootnote field.
func (o *MobileDeviceSecurity) SetLostModeFootnote(v string) {
	o.LostModeFootnote = &v
}

// GetLostModeLocation returns the LostModeLocation field value if set, zero value otherwise.
func (o *MobileDeviceSecurity) GetLostModeLocation() MobileDeviceLostModeLocation {
	if o == nil || IsNil(o.LostModeLocation) {
		var ret MobileDeviceLostModeLocation
		return ret
	}
	return *o.LostModeLocation
}

// GetLostModeLocationOk returns a tuple with the LostModeLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSecurity) GetLostModeLocationOk() (*MobileDeviceLostModeLocation, bool) {
	if o == nil || IsNil(o.LostModeLocation) {
		return nil, false
	}
	return o.LostModeLocation, true
}

// HasLostModeLocation returns a boolean if a field has been set.
func (o *MobileDeviceSecurity) HasLostModeLocation() bool {
	if o != nil && !IsNil(o.LostModeLocation) {
		return true
	}

	return false
}

// SetLostModeLocation gets a reference to the given MobileDeviceLostModeLocation and assigns it to the LostModeLocation field.
func (o *MobileDeviceSecurity) SetLostModeLocation(v MobileDeviceLostModeLocation) {
	o.LostModeLocation = &v
}

func (o MobileDeviceSecurity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobileDeviceSecurity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataProtected) {
		toSerialize["dataProtected"] = o.DataProtected
	}
	if !IsNil(o.BlockLevelEncryptionCapable) {
		toSerialize["blockLevelEncryptionCapable"] = o.BlockLevelEncryptionCapable
	}
	if !IsNil(o.FileLevelEncryptionCapable) {
		toSerialize["fileLevelEncryptionCapable"] = o.FileLevelEncryptionCapable
	}
	if !IsNil(o.PasscodePresent) {
		toSerialize["passcodePresent"] = o.PasscodePresent
	}
	if !IsNil(o.PasscodeCompliant) {
		toSerialize["passcodeCompliant"] = o.PasscodeCompliant
	}
	if !IsNil(o.PasscodeCompliantWithProfile) {
		toSerialize["passcodeCompliantWithProfile"] = o.PasscodeCompliantWithProfile
	}
	if !IsNil(o.HardwareEncryption) {
		toSerialize["hardwareEncryption"] = o.HardwareEncryption
	}
	if !IsNil(o.ActivationLockEnabled) {
		toSerialize["activationLockEnabled"] = o.ActivationLockEnabled
	}
	if !IsNil(o.JailBreakDetected) {
		toSerialize["jailBreakDetected"] = o.JailBreakDetected
	}
	if !IsNil(o.PasscodeLockGracePeriodEnforcedSeconds) {
		toSerialize["passcodeLockGracePeriodEnforcedSeconds"] = o.PasscodeLockGracePeriodEnforcedSeconds
	}
	if !IsNil(o.PersonalDeviceProfileCurrent) {
		toSerialize["personalDeviceProfileCurrent"] = o.PersonalDeviceProfileCurrent
	}
	if !IsNil(o.LostModeEnabled) {
		toSerialize["lostModeEnabled"] = o.LostModeEnabled
	}
	if !IsNil(o.LostModePersistent) {
		toSerialize["lostModePersistent"] = o.LostModePersistent
	}
	if !IsNil(o.LostModeMessage) {
		toSerialize["lostModeMessage"] = o.LostModeMessage
	}
	if !IsNil(o.LostModePhoneNumber) {
		toSerialize["lostModePhoneNumber"] = o.LostModePhoneNumber
	}
	if !IsNil(o.LostModeFootnote) {
		toSerialize["lostModeFootnote"] = o.LostModeFootnote
	}
	if !IsNil(o.LostModeLocation) {
		toSerialize["lostModeLocation"] = o.LostModeLocation
	}
	return toSerialize, nil
}

type NullableMobileDeviceSecurity struct {
	value *MobileDeviceSecurity
	isSet bool
}

func (v NullableMobileDeviceSecurity) Get() *MobileDeviceSecurity {
	return v.value
}

func (v *NullableMobileDeviceSecurity) Set(val *MobileDeviceSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDeviceSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDeviceSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDeviceSecurity(val *MobileDeviceSecurity) *NullableMobileDeviceSecurity {
	return &NullableMobileDeviceSecurity{value: val, isSet: true}
}

func (v NullableMobileDeviceSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDeviceSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


