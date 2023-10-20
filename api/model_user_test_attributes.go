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

// checks if the UserTestAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserTestAttributes{}

// UserTestAttributes struct for UserTestAttributes
type UserTestAttributes struct {
	FullName *string `json:"fullName,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Position *string `json:"position,omitempty"`
	Room *string `json:"room,omitempty"`
	BuildingName *string `json:"buildingName,omitempty"`
	DepartmentName *string `json:"departmentName,omitempty"`
}

// NewUserTestAttributes instantiates a new UserTestAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTestAttributes() *UserTestAttributes {
	this := UserTestAttributes{}
	return &this
}

// NewUserTestAttributesWithDefaults instantiates a new UserTestAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTestAttributesWithDefaults() *UserTestAttributes {
	this := UserTestAttributes{}
	return &this
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *UserTestAttributes) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTestAttributes) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *UserTestAttributes) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *UserTestAttributes) SetFullName(v string) {
	o.FullName = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *UserTestAttributes) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTestAttributes) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *UserTestAttributes) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *UserTestAttributes) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *UserTestAttributes) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTestAttributes) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *UserTestAttributes) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *UserTestAttributes) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *UserTestAttributes) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTestAttributes) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *UserTestAttributes) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *UserTestAttributes) SetPosition(v string) {
	o.Position = &v
}

// GetRoom returns the Room field value if set, zero value otherwise.
func (o *UserTestAttributes) GetRoom() string {
	if o == nil || IsNil(o.Room) {
		var ret string
		return ret
	}
	return *o.Room
}

// GetRoomOk returns a tuple with the Room field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTestAttributes) GetRoomOk() (*string, bool) {
	if o == nil || IsNil(o.Room) {
		return nil, false
	}
	return o.Room, true
}

// HasRoom returns a boolean if a field has been set.
func (o *UserTestAttributes) HasRoom() bool {
	if o != nil && !IsNil(o.Room) {
		return true
	}

	return false
}

// SetRoom gets a reference to the given string and assigns it to the Room field.
func (o *UserTestAttributes) SetRoom(v string) {
	o.Room = &v
}

// GetBuildingName returns the BuildingName field value if set, zero value otherwise.
func (o *UserTestAttributes) GetBuildingName() string {
	if o == nil || IsNil(o.BuildingName) {
		var ret string
		return ret
	}
	return *o.BuildingName
}

// GetBuildingNameOk returns a tuple with the BuildingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTestAttributes) GetBuildingNameOk() (*string, bool) {
	if o == nil || IsNil(o.BuildingName) {
		return nil, false
	}
	return o.BuildingName, true
}

// HasBuildingName returns a boolean if a field has been set.
func (o *UserTestAttributes) HasBuildingName() bool {
	if o != nil && !IsNil(o.BuildingName) {
		return true
	}

	return false
}

// SetBuildingName gets a reference to the given string and assigns it to the BuildingName field.
func (o *UserTestAttributes) SetBuildingName(v string) {
	o.BuildingName = &v
}

// GetDepartmentName returns the DepartmentName field value if set, zero value otherwise.
func (o *UserTestAttributes) GetDepartmentName() string {
	if o == nil || IsNil(o.DepartmentName) {
		var ret string
		return ret
	}
	return *o.DepartmentName
}

// GetDepartmentNameOk returns a tuple with the DepartmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTestAttributes) GetDepartmentNameOk() (*string, bool) {
	if o == nil || IsNil(o.DepartmentName) {
		return nil, false
	}
	return o.DepartmentName, true
}

// HasDepartmentName returns a boolean if a field has been set.
func (o *UserTestAttributes) HasDepartmentName() bool {
	if o != nil && !IsNil(o.DepartmentName) {
		return true
	}

	return false
}

// SetDepartmentName gets a reference to the given string and assigns it to the DepartmentName field.
func (o *UserTestAttributes) SetDepartmentName(v string) {
	o.DepartmentName = &v
}

func (o UserTestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserTestAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Room) {
		toSerialize["room"] = o.Room
	}
	if !IsNil(o.BuildingName) {
		toSerialize["buildingName"] = o.BuildingName
	}
	if !IsNil(o.DepartmentName) {
		toSerialize["departmentName"] = o.DepartmentName
	}
	return toSerialize, nil
}

type NullableUserTestAttributes struct {
	value *UserTestAttributes
	isSet bool
}

func (v NullableUserTestAttributes) Get() *UserTestAttributes {
	return v.value
}

func (v *NullableUserTestAttributes) Set(val *UserTestAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTestAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTestAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTestAttributes(val *UserTestAttributes) *NullableUserTestAttributes {
	return &NullableUserTestAttributes{value: val, isSet: true}
}

func (v NullableUserTestAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTestAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


