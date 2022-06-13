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

// SsoKeystoreDetails struct for SsoKeystoreDetails
type SsoKeystoreDetails struct {
	Keys []string `json:"keys,omitempty"`
	SerialNumber *int32 `json:"serialNumber,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Issuer *string `json:"issuer,omitempty"`
	Expiration *string `json:"expiration,omitempty"`
}

// NewSsoKeystoreDetails instantiates a new SsoKeystoreDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoKeystoreDetails() *SsoKeystoreDetails {
	this := SsoKeystoreDetails{}
	return &this
}

// NewSsoKeystoreDetailsWithDefaults instantiates a new SsoKeystoreDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoKeystoreDetailsWithDefaults() *SsoKeystoreDetails {
	this := SsoKeystoreDetails{}
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *SsoKeystoreDetails) GetKeys() []string {
	if o == nil || o.Keys == nil {
		var ret []string
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoKeystoreDetails) GetKeysOk() ([]string, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *SsoKeystoreDetails) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []string and assigns it to the Keys field.
func (o *SsoKeystoreDetails) SetKeys(v []string) {
	o.Keys = v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *SsoKeystoreDetails) GetSerialNumber() int32 {
	if o == nil || o.SerialNumber == nil {
		var ret int32
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoKeystoreDetails) GetSerialNumberOk() (*int32, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *SsoKeystoreDetails) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given int32 and assigns it to the SerialNumber field.
func (o *SsoKeystoreDetails) SetSerialNumber(v int32) {
	o.SerialNumber = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *SsoKeystoreDetails) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoKeystoreDetails) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *SsoKeystoreDetails) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *SsoKeystoreDetails) SetSubject(v string) {
	o.Subject = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *SsoKeystoreDetails) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoKeystoreDetails) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *SsoKeystoreDetails) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *SsoKeystoreDetails) SetIssuer(v string) {
	o.Issuer = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *SsoKeystoreDetails) GetExpiration() string {
	if o == nil || o.Expiration == nil {
		var ret string
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoKeystoreDetails) GetExpirationOk() (*string, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *SsoKeystoreDetails) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given string and assigns it to the Expiration field.
func (o *SsoKeystoreDetails) SetExpiration(v string) {
	o.Expiration = &v
}

func (o SsoKeystoreDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	if o.SerialNumber != nil {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.Expiration != nil {
		toSerialize["expiration"] = o.Expiration
	}
	return json.Marshal(toSerialize)
}

type NullableSsoKeystoreDetails struct {
	value *SsoKeystoreDetails
	isSet bool
}

func (v NullableSsoKeystoreDetails) Get() *SsoKeystoreDetails {
	return v.value
}

func (v *NullableSsoKeystoreDetails) Set(val *SsoKeystoreDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoKeystoreDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoKeystoreDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoKeystoreDetails(val *SsoKeystoreDetails) *NullableSsoKeystoreDetails {
	return &NullableSsoKeystoreDetails{value: val, isSet: true}
}

func (v NullableSsoKeystoreDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoKeystoreDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


