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

// VolumePurchasingLocationPost struct for VolumePurchasingLocationPost
type VolumePurchasingLocationPost struct {
	// If no value is provided when creating a VolumePurchasingLocation object, the 'name' will default to the 'locationName' value
	Name *string `json:"name,omitempty"`
	AutomaticallyPopulatePurchasedContent *bool `json:"automaticallyPopulatePurchasedContent,omitempty"`
	SendNotificationWhenNoLongerAssigned *bool `json:"sendNotificationWhenNoLongerAssigned,omitempty"`
	AutoRegisterManagedUsers *bool `json:"autoRegisterManagedUsers,omitempty"`
	SiteId *string `json:"siteId,omitempty"`
	ServiceToken string `json:"serviceToken"`
}

// NewVolumePurchasingLocationPost instantiates a new VolumePurchasingLocationPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumePurchasingLocationPost(serviceToken string) *VolumePurchasingLocationPost {
	this := VolumePurchasingLocationPost{}
	var automaticallyPopulatePurchasedContent bool = false
	this.AutomaticallyPopulatePurchasedContent = &automaticallyPopulatePurchasedContent
	var sendNotificationWhenNoLongerAssigned bool = false
	this.SendNotificationWhenNoLongerAssigned = &sendNotificationWhenNoLongerAssigned
	var autoRegisterManagedUsers bool = false
	this.AutoRegisterManagedUsers = &autoRegisterManagedUsers
	var siteId string = "-1"
	this.SiteId = &siteId
	this.ServiceToken = serviceToken
	return &this
}

// NewVolumePurchasingLocationPostWithDefaults instantiates a new VolumePurchasingLocationPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumePurchasingLocationPostWithDefaults() *VolumePurchasingLocationPost {
	this := VolumePurchasingLocationPost{}
	var automaticallyPopulatePurchasedContent bool = false
	this.AutomaticallyPopulatePurchasedContent = &automaticallyPopulatePurchasedContent
	var sendNotificationWhenNoLongerAssigned bool = false
	this.SendNotificationWhenNoLongerAssigned = &sendNotificationWhenNoLongerAssigned
	var autoRegisterManagedUsers bool = false
	this.AutoRegisterManagedUsers = &autoRegisterManagedUsers
	var siteId string = "-1"
	this.SiteId = &siteId
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VolumePurchasingLocationPost) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationPost) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VolumePurchasingLocationPost) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VolumePurchasingLocationPost) SetName(v string) {
	o.Name = &v
}

// GetAutomaticallyPopulatePurchasedContent returns the AutomaticallyPopulatePurchasedContent field value if set, zero value otherwise.
func (o *VolumePurchasingLocationPost) GetAutomaticallyPopulatePurchasedContent() bool {
	if o == nil || o.AutomaticallyPopulatePurchasedContent == nil {
		var ret bool
		return ret
	}
	return *o.AutomaticallyPopulatePurchasedContent
}

// GetAutomaticallyPopulatePurchasedContentOk returns a tuple with the AutomaticallyPopulatePurchasedContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationPost) GetAutomaticallyPopulatePurchasedContentOk() (*bool, bool) {
	if o == nil || o.AutomaticallyPopulatePurchasedContent == nil {
		return nil, false
	}
	return o.AutomaticallyPopulatePurchasedContent, true
}

// HasAutomaticallyPopulatePurchasedContent returns a boolean if a field has been set.
func (o *VolumePurchasingLocationPost) HasAutomaticallyPopulatePurchasedContent() bool {
	if o != nil && o.AutomaticallyPopulatePurchasedContent != nil {
		return true
	}

	return false
}

// SetAutomaticallyPopulatePurchasedContent gets a reference to the given bool and assigns it to the AutomaticallyPopulatePurchasedContent field.
func (o *VolumePurchasingLocationPost) SetAutomaticallyPopulatePurchasedContent(v bool) {
	o.AutomaticallyPopulatePurchasedContent = &v
}

// GetSendNotificationWhenNoLongerAssigned returns the SendNotificationWhenNoLongerAssigned field value if set, zero value otherwise.
func (o *VolumePurchasingLocationPost) GetSendNotificationWhenNoLongerAssigned() bool {
	if o == nil || o.SendNotificationWhenNoLongerAssigned == nil {
		var ret bool
		return ret
	}
	return *o.SendNotificationWhenNoLongerAssigned
}

// GetSendNotificationWhenNoLongerAssignedOk returns a tuple with the SendNotificationWhenNoLongerAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationPost) GetSendNotificationWhenNoLongerAssignedOk() (*bool, bool) {
	if o == nil || o.SendNotificationWhenNoLongerAssigned == nil {
		return nil, false
	}
	return o.SendNotificationWhenNoLongerAssigned, true
}

// HasSendNotificationWhenNoLongerAssigned returns a boolean if a field has been set.
func (o *VolumePurchasingLocationPost) HasSendNotificationWhenNoLongerAssigned() bool {
	if o != nil && o.SendNotificationWhenNoLongerAssigned != nil {
		return true
	}

	return false
}

// SetSendNotificationWhenNoLongerAssigned gets a reference to the given bool and assigns it to the SendNotificationWhenNoLongerAssigned field.
func (o *VolumePurchasingLocationPost) SetSendNotificationWhenNoLongerAssigned(v bool) {
	o.SendNotificationWhenNoLongerAssigned = &v
}

// GetAutoRegisterManagedUsers returns the AutoRegisterManagedUsers field value if set, zero value otherwise.
func (o *VolumePurchasingLocationPost) GetAutoRegisterManagedUsers() bool {
	if o == nil || o.AutoRegisterManagedUsers == nil {
		var ret bool
		return ret
	}
	return *o.AutoRegisterManagedUsers
}

// GetAutoRegisterManagedUsersOk returns a tuple with the AutoRegisterManagedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationPost) GetAutoRegisterManagedUsersOk() (*bool, bool) {
	if o == nil || o.AutoRegisterManagedUsers == nil {
		return nil, false
	}
	return o.AutoRegisterManagedUsers, true
}

// HasAutoRegisterManagedUsers returns a boolean if a field has been set.
func (o *VolumePurchasingLocationPost) HasAutoRegisterManagedUsers() bool {
	if o != nil && o.AutoRegisterManagedUsers != nil {
		return true
	}

	return false
}

// SetAutoRegisterManagedUsers gets a reference to the given bool and assigns it to the AutoRegisterManagedUsers field.
func (o *VolumePurchasingLocationPost) SetAutoRegisterManagedUsers(v bool) {
	o.AutoRegisterManagedUsers = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *VolumePurchasingLocationPost) GetSiteId() string {
	if o == nil || o.SiteId == nil {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationPost) GetSiteIdOk() (*string, bool) {
	if o == nil || o.SiteId == nil {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *VolumePurchasingLocationPost) HasSiteId() bool {
	if o != nil && o.SiteId != nil {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *VolumePurchasingLocationPost) SetSiteId(v string) {
	o.SiteId = &v
}

// GetServiceToken returns the ServiceToken field value
func (o *VolumePurchasingLocationPost) GetServiceToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceToken
}

// GetServiceTokenOk returns a tuple with the ServiceToken field value
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationPost) GetServiceTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceToken, true
}

// SetServiceToken sets field value
func (o *VolumePurchasingLocationPost) SetServiceToken(v string) {
	o.ServiceToken = v
}

func (o VolumePurchasingLocationPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.AutomaticallyPopulatePurchasedContent != nil {
		toSerialize["automaticallyPopulatePurchasedContent"] = o.AutomaticallyPopulatePurchasedContent
	}
	if o.SendNotificationWhenNoLongerAssigned != nil {
		toSerialize["sendNotificationWhenNoLongerAssigned"] = o.SendNotificationWhenNoLongerAssigned
	}
	if o.AutoRegisterManagedUsers != nil {
		toSerialize["autoRegisterManagedUsers"] = o.AutoRegisterManagedUsers
	}
	if o.SiteId != nil {
		toSerialize["siteId"] = o.SiteId
	}
	if true {
		toSerialize["serviceToken"] = o.ServiceToken
	}
	return json.Marshal(toSerialize)
}

type NullableVolumePurchasingLocationPost struct {
	value *VolumePurchasingLocationPost
	isSet bool
}

func (v NullableVolumePurchasingLocationPost) Get() *VolumePurchasingLocationPost {
	return v.value
}

func (v *NullableVolumePurchasingLocationPost) Set(val *VolumePurchasingLocationPost) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumePurchasingLocationPost) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumePurchasingLocationPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumePurchasingLocationPost(val *VolumePurchasingLocationPost) *NullableVolumePurchasingLocationPost {
	return &NullableVolumePurchasingLocationPost{value: val, isSet: true}
}

func (v NullableVolumePurchasingLocationPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumePurchasingLocationPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


