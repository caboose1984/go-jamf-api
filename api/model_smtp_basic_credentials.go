/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SmtpBasicCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmtpBasicCredentials{}

// SmtpBasicCredentials struct for SmtpBasicCredentials
type SmtpBasicCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type _SmtpBasicCredentials SmtpBasicCredentials

// NewSmtpBasicCredentials instantiates a new SmtpBasicCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmtpBasicCredentials(username string, password string) *SmtpBasicCredentials {
	this := SmtpBasicCredentials{}
	this.Username = username
	this.Password = password
	return &this
}

// NewSmtpBasicCredentialsWithDefaults instantiates a new SmtpBasicCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmtpBasicCredentialsWithDefaults() *SmtpBasicCredentials {
	this := SmtpBasicCredentials{}
	var username string = ""
	this.Username = username
	var password string = ""
	this.Password = password
	return &this
}

// GetUsername returns the Username field value
func (o *SmtpBasicCredentials) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *SmtpBasicCredentials) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *SmtpBasicCredentials) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *SmtpBasicCredentials) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SmtpBasicCredentials) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SmtpBasicCredentials) SetPassword(v string) {
	o.Password = v
}

func (o SmtpBasicCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmtpBasicCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *SmtpBasicCredentials) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
		"password",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSmtpBasicCredentials := _SmtpBasicCredentials{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmtpBasicCredentials)

	if err != nil {
		return err
	}

	*o = SmtpBasicCredentials(varSmtpBasicCredentials)

	return err
}

type NullableSmtpBasicCredentials struct {
	value *SmtpBasicCredentials
	isSet bool
}

func (v NullableSmtpBasicCredentials) Get() *SmtpBasicCredentials {
	return v.value
}

func (v *NullableSmtpBasicCredentials) Set(val *SmtpBasicCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableSmtpBasicCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableSmtpBasicCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmtpBasicCredentials(val *SmtpBasicCredentials) *NullableSmtpBasicCredentials {
	return &NullableSmtpBasicCredentials{value: val, isSet: true}
}

func (v NullableSmtpBasicCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmtpBasicCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

