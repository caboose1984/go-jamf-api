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

// checks if the AccountPreferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountPreferences{}

// AccountPreferences struct for AccountPreferences
type AccountPreferences struct {
	Language *string `json:"language,omitempty"`
	DateFormat *string `json:"dateFormat,omitempty"`
	Region *string `json:"region,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	IsDisableRelativeDates *bool `json:"isDisableRelativeDates,omitempty"`
}

// NewAccountPreferences instantiates a new AccountPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountPreferences() *AccountPreferences {
	this := AccountPreferences{}
	return &this
}

// NewAccountPreferencesWithDefaults instantiates a new AccountPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountPreferencesWithDefaults() *AccountPreferences {
	this := AccountPreferences{}
	return &this
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AccountPreferences) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPreferences) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AccountPreferences) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *AccountPreferences) SetLanguage(v string) {
	o.Language = &v
}

// GetDateFormat returns the DateFormat field value if set, zero value otherwise.
func (o *AccountPreferences) GetDateFormat() string {
	if o == nil || IsNil(o.DateFormat) {
		var ret string
		return ret
	}
	return *o.DateFormat
}

// GetDateFormatOk returns a tuple with the DateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPreferences) GetDateFormatOk() (*string, bool) {
	if o == nil || IsNil(o.DateFormat) {
		return nil, false
	}
	return o.DateFormat, true
}

// HasDateFormat returns a boolean if a field has been set.
func (o *AccountPreferences) HasDateFormat() bool {
	if o != nil && !IsNil(o.DateFormat) {
		return true
	}

	return false
}

// SetDateFormat gets a reference to the given string and assigns it to the DateFormat field.
func (o *AccountPreferences) SetDateFormat(v string) {
	o.DateFormat = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AccountPreferences) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPreferences) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AccountPreferences) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AccountPreferences) SetRegion(v string) {
	o.Region = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *AccountPreferences) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPreferences) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *AccountPreferences) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *AccountPreferences) SetTimezone(v string) {
	o.Timezone = &v
}

// GetIsDisableRelativeDates returns the IsDisableRelativeDates field value if set, zero value otherwise.
func (o *AccountPreferences) GetIsDisableRelativeDates() bool {
	if o == nil || IsNil(o.IsDisableRelativeDates) {
		var ret bool
		return ret
	}
	return *o.IsDisableRelativeDates
}

// GetIsDisableRelativeDatesOk returns a tuple with the IsDisableRelativeDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPreferences) GetIsDisableRelativeDatesOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDisableRelativeDates) {
		return nil, false
	}
	return o.IsDisableRelativeDates, true
}

// HasIsDisableRelativeDates returns a boolean if a field has been set.
func (o *AccountPreferences) HasIsDisableRelativeDates() bool {
	if o != nil && !IsNil(o.IsDisableRelativeDates) {
		return true
	}

	return false
}

// SetIsDisableRelativeDates gets a reference to the given bool and assigns it to the IsDisableRelativeDates field.
func (o *AccountPreferences) SetIsDisableRelativeDates(v bool) {
	o.IsDisableRelativeDates = &v
}

func (o AccountPreferences) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountPreferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.DateFormat) {
		toSerialize["dateFormat"] = o.DateFormat
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.IsDisableRelativeDates) {
		toSerialize["isDisableRelativeDates"] = o.IsDisableRelativeDates
	}
	return toSerialize, nil
}

type NullableAccountPreferences struct {
	value *AccountPreferences
	isSet bool
}

func (v NullableAccountPreferences) Get() *AccountPreferences {
	return v.value
}

func (v *NullableAccountPreferences) Set(val *AccountPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountPreferences(val *AccountPreferences) *NullableAccountPreferences {
	return &NullableAccountPreferences{value: val, isSet: true}
}

func (v NullableAccountPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


