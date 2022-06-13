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

// ComputerSection the model 'ComputerSection'
type ComputerSection string

// List of ComputerSection
const (
	GENERAL ComputerSection = "GENERAL"
	DISK_ENCRYPTION ComputerSection = "DISK_ENCRYPTION"
	PURCHASING ComputerSection = "PURCHASING"
	APPLICATIONS ComputerSection = "APPLICATIONS"
	STORAGE ComputerSection = "STORAGE"
	USER_AND_LOCATION ComputerSection = "USER_AND_LOCATION"
	CONFIGURATION_PROFILES ComputerSection = "CONFIGURATION_PROFILES"
	PRINTERS ComputerSection = "PRINTERS"
	SERVICES ComputerSection = "SERVICES"
	HARDWARE ComputerSection = "HARDWARE"
	LOCAL_USER_ACCOUNTS ComputerSection = "LOCAL_USER_ACCOUNTS"
	CERTIFICATES ComputerSection = "CERTIFICATES"
	ATTACHMENTS ComputerSection = "ATTACHMENTS"
	PLUGINS ComputerSection = "PLUGINS"
	PACKAGE_RECEIPTS ComputerSection = "PACKAGE_RECEIPTS"
	FONTS ComputerSection = "FONTS"
	SECURITY ComputerSection = "SECURITY"
	OPERATING_SYSTEM ComputerSection = "OPERATING_SYSTEM"
	LICENSED_SOFTWARE ComputerSection = "LICENSED_SOFTWARE"
	IBEACONS ComputerSection = "IBEACONS"
	SOFTWARE_UPDATES ComputerSection = "SOFTWARE_UPDATES"
	EXTENSION_ATTRIBUTES ComputerSection = "EXTENSION_ATTRIBUTES"
	CONTENT_CACHING ComputerSection = "CONTENT_CACHING"
	GROUP_MEMBERSHIPS ComputerSection = "GROUP_MEMBERSHIPS"
)

// All allowed values of ComputerSection enum
var AllowedComputerSectionEnumValues = []ComputerSection{
	"GENERAL",
	"DISK_ENCRYPTION",
	"PURCHASING",
	"APPLICATIONS",
	"STORAGE",
	"USER_AND_LOCATION",
	"CONFIGURATION_PROFILES",
	"PRINTERS",
	"SERVICES",
	"HARDWARE",
	"LOCAL_USER_ACCOUNTS",
	"CERTIFICATES",
	"ATTACHMENTS",
	"PLUGINS",
	"PACKAGE_RECEIPTS",
	"FONTS",
	"SECURITY",
	"OPERATING_SYSTEM",
	"LICENSED_SOFTWARE",
	"IBEACONS",
	"SOFTWARE_UPDATES",
	"EXTENSION_ATTRIBUTES",
	"CONTENT_CACHING",
	"GROUP_MEMBERSHIPS",
}

func (v *ComputerSection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ComputerSection(value)
	for _, existing := range AllowedComputerSectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ComputerSection", value)
}

// NewComputerSectionFromValue returns a pointer to a valid ComputerSection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComputerSectionFromValue(v string) (*ComputerSection, error) {
	ev := ComputerSection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ComputerSection: valid values are %v", v, AllowedComputerSectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComputerSection) IsValid() bool {
	for _, existing := range AllowedComputerSectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ComputerSection value
func (v ComputerSection) Ptr() *ComputerSection {
	return &v
}

type NullableComputerSection struct {
	value *ComputerSection
	isSet bool
}

func (v NullableComputerSection) Get() *ComputerSection {
	return v.value
}

func (v *NullableComputerSection) Set(val *ComputerSection) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerSection) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerSection(val *ComputerSection) *NullableComputerSection {
	return &NullableComputerSection{value: val, isSet: true}
}

func (v NullableComputerSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

