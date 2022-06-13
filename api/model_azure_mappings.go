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

// AzureMappings Azure Cloud Identity Provider mappings
type AzureMappings struct {
	UserId string `json:"userId"`
	UserName string `json:"userName"`
	RealName string `json:"realName"`
	Email string `json:"email"`
	Department string `json:"department"`
	Building string `json:"building"`
	Room string `json:"room"`
	Phone string `json:"phone"`
	Position string `json:"position"`
	GroupId string `json:"groupId"`
	GroupName string `json:"groupName"`
}

// NewAzureMappings instantiates a new AzureMappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMappings(userId string, userName string, realName string, email string, department string, building string, room string, phone string, position string, groupId string, groupName string) *AzureMappings {
	this := AzureMappings{}
	this.UserId = userId
	this.UserName = userName
	this.RealName = realName
	this.Email = email
	this.Department = department
	this.Building = building
	this.Room = room
	this.Phone = phone
	this.Position = position
	this.GroupId = groupId
	this.GroupName = groupName
	return &this
}

// NewAzureMappingsWithDefaults instantiates a new AzureMappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMappingsWithDefaults() *AzureMappings {
	this := AzureMappings{}
	return &this
}

// GetUserId returns the UserId field value
func (o *AzureMappings) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *AzureMappings) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *AzureMappings) SetUserId(v string) {
	o.UserId = v
}

// GetUserName returns the UserName field value
func (o *AzureMappings) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *AzureMappings) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *AzureMappings) SetUserName(v string) {
	o.UserName = v
}

// GetRealName returns the RealName field value
func (o *AzureMappings) GetRealName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RealName
}

// GetRealNameOk returns a tuple with the RealName field value
// and a boolean to check if the value has been set.
func (o *AzureMappings) GetRealNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RealName, true
}

// SetRealName sets field value
func (o *AzureMappings) SetRealName(v string) {
	o.RealName = v
}

// GetEmail returns the Email field value
func (o *AzureMappings) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AzureMappings) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AzureMappings) SetEmail(v string) {
	o.Email = v
}

// GetDepartment returns the Department field value
func (o *AzureMappings) GetDepartment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Department
}

// GetDepartmentOk returns a tuple with the Department field value
// and a boolean to check if the value has been set.
func (o *AzureMappings) GetDepartmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Department, true
}

// SetDepartment sets field value
func (o *AzureMappings) SetDepartment(v string) {
	o.Department = v
}

// GetBuilding returns the Building field value
func (o *AzureMappings) GetBuilding() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Building
}

// GetBuildingOk returns a tuple with the Building field value
// and a boolean to check if the value has been set.
func (o *AzureMappings) GetBuildingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Building, true
}

// SetBuilding sets field value
func (o *AzureMappings) SetBuilding(v string) {
	o.Building = v
}

// GetRoom returns the Room field value
func (o *AzureMappings) GetRoom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Room
}

// GetRoomOk returns a tuple with the Room field value
// and a boolean to check if the value has been set.
func (o *AzureMappings) GetRoomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Room, true
}

// SetRoom sets field value
func (o *AzureMappings) SetRoom(v string) {
	o.Room = v
}

// GetPhone returns the Phone field value
func (o *AzureMappings) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *AzureMappings) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *AzureMappings) SetPhone(v string) {
	o.Phone = v
}

// GetPosition returns the Position field value
func (o *AzureMappings) GetPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *AzureMappings) GetPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *AzureMappings) SetPosition(v string) {
	o.Position = v
}

// GetGroupId returns the GroupId field value
func (o *AzureMappings) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *AzureMappings) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *AzureMappings) SetGroupId(v string) {
	o.GroupId = v
}

// GetGroupName returns the GroupName field value
func (o *AzureMappings) GetGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value
// and a boolean to check if the value has been set.
func (o *AzureMappings) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupName, true
}

// SetGroupName sets field value
func (o *AzureMappings) SetGroupName(v string) {
	o.GroupName = v
}

func (o AzureMappings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["userId"] = o.UserId
	}
	if true {
		toSerialize["userName"] = o.UserName
	}
	if true {
		toSerialize["realName"] = o.RealName
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["department"] = o.Department
	}
	if true {
		toSerialize["building"] = o.Building
	}
	if true {
		toSerialize["room"] = o.Room
	}
	if true {
		toSerialize["phone"] = o.Phone
	}
	if true {
		toSerialize["position"] = o.Position
	}
	if true {
		toSerialize["groupId"] = o.GroupId
	}
	if true {
		toSerialize["groupName"] = o.GroupName
	}
	return json.Marshal(toSerialize)
}

type NullableAzureMappings struct {
	value *AzureMappings
	isSet bool
}

func (v NullableAzureMappings) Get() *AzureMappings {
	return v.value
}

func (v *NullableAzureMappings) Set(val *AzureMappings) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMappings) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMappings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMappings(val *AzureMappings) *NullableAzureMappings {
	return &NullableAzureMappings{value: val, isSet: true}
}

func (v NullableAzureMappings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMappings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


