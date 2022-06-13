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

// VolumePurchasingSubscription struct for VolumePurchasingSubscription
type VolumePurchasingSubscription struct {
	Name string `json:"name"`
	Enabled *bool `json:"enabled,omitempty"`
	Triggers []string `json:"triggers,omitempty"`
	LocationIds []string `json:"locationIds,omitempty"`
	InternalRecipients []InternalRecipient `json:"internalRecipients,omitempty"`
	ExternalRecipients []ExternalRecipient `json:"externalRecipients,omitempty"`
	SiteId *string `json:"siteId,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewVolumePurchasingSubscription instantiates a new VolumePurchasingSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumePurchasingSubscription(name string) *VolumePurchasingSubscription {
	this := VolumePurchasingSubscription{}
	this.Name = name
	var enabled bool = true
	this.Enabled = &enabled
	var siteId string = "-1"
	this.SiteId = &siteId
	return &this
}

// NewVolumePurchasingSubscriptionWithDefaults instantiates a new VolumePurchasingSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumePurchasingSubscriptionWithDefaults() *VolumePurchasingSubscription {
	this := VolumePurchasingSubscription{}
	var enabled bool = true
	this.Enabled = &enabled
	var siteId string = "-1"
	this.SiteId = &siteId
	return &this
}

// GetName returns the Name field value
func (o *VolumePurchasingSubscription) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscription) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VolumePurchasingSubscription) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VolumePurchasingSubscription) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscription) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VolumePurchasingSubscription) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VolumePurchasingSubscription) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *VolumePurchasingSubscription) GetTriggers() []string {
	if o == nil || o.Triggers == nil {
		var ret []string
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscription) GetTriggersOk() ([]string, bool) {
	if o == nil || o.Triggers == nil {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *VolumePurchasingSubscription) HasTriggers() bool {
	if o != nil && o.Triggers != nil {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []string and assigns it to the Triggers field.
func (o *VolumePurchasingSubscription) SetTriggers(v []string) {
	o.Triggers = v
}

// GetLocationIds returns the LocationIds field value if set, zero value otherwise.
func (o *VolumePurchasingSubscription) GetLocationIds() []string {
	if o == nil || o.LocationIds == nil {
		var ret []string
		return ret
	}
	return o.LocationIds
}

// GetLocationIdsOk returns a tuple with the LocationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscription) GetLocationIdsOk() ([]string, bool) {
	if o == nil || o.LocationIds == nil {
		return nil, false
	}
	return o.LocationIds, true
}

// HasLocationIds returns a boolean if a field has been set.
func (o *VolumePurchasingSubscription) HasLocationIds() bool {
	if o != nil && o.LocationIds != nil {
		return true
	}

	return false
}

// SetLocationIds gets a reference to the given []string and assigns it to the LocationIds field.
func (o *VolumePurchasingSubscription) SetLocationIds(v []string) {
	o.LocationIds = v
}

// GetInternalRecipients returns the InternalRecipients field value if set, zero value otherwise.
func (o *VolumePurchasingSubscription) GetInternalRecipients() []InternalRecipient {
	if o == nil || o.InternalRecipients == nil {
		var ret []InternalRecipient
		return ret
	}
	return o.InternalRecipients
}

// GetInternalRecipientsOk returns a tuple with the InternalRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscription) GetInternalRecipientsOk() ([]InternalRecipient, bool) {
	if o == nil || o.InternalRecipients == nil {
		return nil, false
	}
	return o.InternalRecipients, true
}

// HasInternalRecipients returns a boolean if a field has been set.
func (o *VolumePurchasingSubscription) HasInternalRecipients() bool {
	if o != nil && o.InternalRecipients != nil {
		return true
	}

	return false
}

// SetInternalRecipients gets a reference to the given []InternalRecipient and assigns it to the InternalRecipients field.
func (o *VolumePurchasingSubscription) SetInternalRecipients(v []InternalRecipient) {
	o.InternalRecipients = v
}

// GetExternalRecipients returns the ExternalRecipients field value if set, zero value otherwise.
func (o *VolumePurchasingSubscription) GetExternalRecipients() []ExternalRecipient {
	if o == nil || o.ExternalRecipients == nil {
		var ret []ExternalRecipient
		return ret
	}
	return o.ExternalRecipients
}

// GetExternalRecipientsOk returns a tuple with the ExternalRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscription) GetExternalRecipientsOk() ([]ExternalRecipient, bool) {
	if o == nil || o.ExternalRecipients == nil {
		return nil, false
	}
	return o.ExternalRecipients, true
}

// HasExternalRecipients returns a boolean if a field has been set.
func (o *VolumePurchasingSubscription) HasExternalRecipients() bool {
	if o != nil && o.ExternalRecipients != nil {
		return true
	}

	return false
}

// SetExternalRecipients gets a reference to the given []ExternalRecipient and assigns it to the ExternalRecipients field.
func (o *VolumePurchasingSubscription) SetExternalRecipients(v []ExternalRecipient) {
	o.ExternalRecipients = v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *VolumePurchasingSubscription) GetSiteId() string {
	if o == nil || o.SiteId == nil {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscription) GetSiteIdOk() (*string, bool) {
	if o == nil || o.SiteId == nil {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *VolumePurchasingSubscription) HasSiteId() bool {
	if o != nil && o.SiteId != nil {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *VolumePurchasingSubscription) SetSiteId(v string) {
	o.SiteId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VolumePurchasingSubscription) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscription) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VolumePurchasingSubscription) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VolumePurchasingSubscription) SetId(v string) {
	o.Id = &v
}

func (o VolumePurchasingSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Triggers != nil {
		toSerialize["triggers"] = o.Triggers
	}
	if o.LocationIds != nil {
		toSerialize["locationIds"] = o.LocationIds
	}
	if o.InternalRecipients != nil {
		toSerialize["internalRecipients"] = o.InternalRecipients
	}
	if o.ExternalRecipients != nil {
		toSerialize["externalRecipients"] = o.ExternalRecipients
	}
	if o.SiteId != nil {
		toSerialize["siteId"] = o.SiteId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableVolumePurchasingSubscription struct {
	value *VolumePurchasingSubscription
	isSet bool
}

func (v NullableVolumePurchasingSubscription) Get() *VolumePurchasingSubscription {
	return v.value
}

func (v *NullableVolumePurchasingSubscription) Set(val *VolumePurchasingSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumePurchasingSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumePurchasingSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumePurchasingSubscription(val *VolumePurchasingSubscription) *NullableVolumePurchasingSubscription {
	return &NullableVolumePurchasingSubscription{value: val, isSet: true}
}

func (v NullableVolumePurchasingSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumePurchasingSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


