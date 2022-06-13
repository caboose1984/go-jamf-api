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

// SessionStatus Session status
type SessionStatus struct {
	// Session state
	SessionState *string `json:"sessionState,omitempty"`
	// Defines if the end user is online
	Online *bool `json:"online,omitempty"`
}

// NewSessionStatus instantiates a new SessionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionStatus() *SessionStatus {
	this := SessionStatus{}
	return &this
}

// NewSessionStatusWithDefaults instantiates a new SessionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionStatusWithDefaults() *SessionStatus {
	this := SessionStatus{}
	return &this
}

// GetSessionState returns the SessionState field value if set, zero value otherwise.
func (o *SessionStatus) GetSessionState() string {
	if o == nil || o.SessionState == nil {
		var ret string
		return ret
	}
	return *o.SessionState
}

// GetSessionStateOk returns a tuple with the SessionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStatus) GetSessionStateOk() (*string, bool) {
	if o == nil || o.SessionState == nil {
		return nil, false
	}
	return o.SessionState, true
}

// HasSessionState returns a boolean if a field has been set.
func (o *SessionStatus) HasSessionState() bool {
	if o != nil && o.SessionState != nil {
		return true
	}

	return false
}

// SetSessionState gets a reference to the given string and assigns it to the SessionState field.
func (o *SessionStatus) SetSessionState(v string) {
	o.SessionState = &v
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *SessionStatus) GetOnline() bool {
	if o == nil || o.Online == nil {
		var ret bool
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStatus) GetOnlineOk() (*bool, bool) {
	if o == nil || o.Online == nil {
		return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *SessionStatus) HasOnline() bool {
	if o != nil && o.Online != nil {
		return true
	}

	return false
}

// SetOnline gets a reference to the given bool and assigns it to the Online field.
func (o *SessionStatus) SetOnline(v bool) {
	o.Online = &v
}

func (o SessionStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SessionState != nil {
		toSerialize["sessionState"] = o.SessionState
	}
	if o.Online != nil {
		toSerialize["online"] = o.Online
	}
	return json.Marshal(toSerialize)
}

type NullableSessionStatus struct {
	value *SessionStatus
	isSet bool
}

func (v NullableSessionStatus) Get() *SessionStatus {
	return v.value
}

func (v *NullableSessionStatus) Set(val *SessionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionStatus(val *SessionStatus) *NullableSessionStatus {
	return &NullableSessionStatus{value: val, isSet: true}
}

func (v NullableSessionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


