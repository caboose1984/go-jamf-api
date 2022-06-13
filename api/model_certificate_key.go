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

// CertificateKey struct for CertificateKey
type CertificateKey struct {
	Id *string `json:"id,omitempty"`
	Valid *bool `json:"valid,omitempty"`
}

// NewCertificateKey instantiates a new CertificateKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateKey() *CertificateKey {
	this := CertificateKey{}
	return &this
}

// NewCertificateKeyWithDefaults instantiates a new CertificateKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateKeyWithDefaults() *CertificateKey {
	this := CertificateKey{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CertificateKey) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateKey) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CertificateKey) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CertificateKey) SetId(v string) {
	o.Id = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *CertificateKey) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateKey) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *CertificateKey) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *CertificateKey) SetValid(v bool) {
	o.Valid = &v
}

func (o CertificateKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateKey struct {
	value *CertificateKey
	isSet bool
}

func (v NullableCertificateKey) Get() *CertificateKey {
	return v.value
}

func (v *NullableCertificateKey) Set(val *CertificateKey) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateKey) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateKey(val *CertificateKey) *NullableCertificateKey {
	return &NullableCertificateKey{value: val, isSet: true}
}

func (v NullableCertificateKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


