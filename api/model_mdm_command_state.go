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

// MdmCommandState the model 'MdmCommandState'
type MdmCommandState string

// List of MdmCommandState
const (
	MDMCOMMANDSTATE_PENDING MdmCommandState = "PENDING"
	MDMCOMMANDSTATE_ACKNOWLEDGED MdmCommandState = "ACKNOWLEDGED"
	MDMCOMMANDSTATE_NOT_NOW MdmCommandState = "NOT_NOW"
	MDMCOMMANDSTATE_ERROR MdmCommandState = "ERROR"
)

// All allowed values of MdmCommandState enum
var AllowedMdmCommandStateEnumValues = []MdmCommandState{
	"PENDING",
	"ACKNOWLEDGED",
	"NOT_NOW",
	"ERROR",
}

func (v *MdmCommandState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MdmCommandState(value)
	for _, existing := range AllowedMdmCommandStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MdmCommandState", value)
}

// NewMdmCommandStateFromValue returns a pointer to a valid MdmCommandState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMdmCommandStateFromValue(v string) (*MdmCommandState, error) {
	ev := MdmCommandState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MdmCommandState: valid values are %v", v, AllowedMdmCommandStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MdmCommandState) IsValid() bool {
	for _, existing := range AllowedMdmCommandStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MdmCommandState value
func (v MdmCommandState) Ptr() *MdmCommandState {
	return &v
}

type NullableMdmCommandState struct {
	value *MdmCommandState
	isSet bool
}

func (v NullableMdmCommandState) Get() *MdmCommandState {
	return v.value
}

func (v *NullableMdmCommandState) Set(val *MdmCommandState) {
	v.value = val
	v.isSet = true
}

func (v NullableMdmCommandState) IsSet() bool {
	return v.isSet
}

func (v *NullableMdmCommandState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMdmCommandState(val *MdmCommandState) *NullableMdmCommandState {
	return &NullableMdmCommandState{value: val, isSet: true}
}

func (v NullableMdmCommandState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMdmCommandState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
