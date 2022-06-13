/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// VoiceRoamingSetting the model 'VoiceRoamingSetting'
type VoiceRoamingSetting string

// List of VoiceRoamingSetting
const (
	ENABLE_VOICE_ROAMING VoiceRoamingSetting = "ENABLE_VOICE_ROAMING"
	DISABLE_VOICE_ROAMING VoiceRoamingSetting = "DISABLE_VOICE_ROAMING"
)

// All allowed values of VoiceRoamingSetting enum
var AllowedVoiceRoamingSettingEnumValues = []VoiceRoamingSetting{
	"ENABLE_VOICE_ROAMING",
	"DISABLE_VOICE_ROAMING",
}

func (v *VoiceRoamingSetting) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VoiceRoamingSetting(value)
	for _, existing := range AllowedVoiceRoamingSettingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VoiceRoamingSetting", value)
}

// NewVoiceRoamingSettingFromValue returns a pointer to a valid VoiceRoamingSetting
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVoiceRoamingSettingFromValue(v string) (*VoiceRoamingSetting, error) {
	ev := VoiceRoamingSetting(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VoiceRoamingSetting: valid values are %v", v, AllowedVoiceRoamingSettingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VoiceRoamingSetting) IsValid() bool {
	for _, existing := range AllowedVoiceRoamingSettingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VoiceRoamingSetting value
func (v VoiceRoamingSetting) Ptr() *VoiceRoamingSetting {
	return &v
}

type NullableVoiceRoamingSetting struct {
	value *VoiceRoamingSetting
	isSet bool
}

func (v NullableVoiceRoamingSetting) Get() *VoiceRoamingSetting {
	return v.value
}

func (v *NullableVoiceRoamingSetting) Set(val *VoiceRoamingSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableVoiceRoamingSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableVoiceRoamingSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoiceRoamingSetting(val *VoiceRoamingSetting) *NullableVoiceRoamingSetting {
	return &NullableVoiceRoamingSetting{value: val, isSet: true}
}

func (v NullableVoiceRoamingSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoiceRoamingSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

