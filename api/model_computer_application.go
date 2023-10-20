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

// checks if the ComputerApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComputerApplication{}

// ComputerApplication struct for ComputerApplication
type ComputerApplication struct {
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
	Version *string `json:"version,omitempty"`
	MacAppStore *bool `json:"macAppStore,omitempty"`
	SizeMegabytes *int32 `json:"sizeMegabytes,omitempty"`
	BundleId *string `json:"bundleId,omitempty"`
	UpdateAvailable *bool `json:"updateAvailable,omitempty"`
	// The app's external version ID. It can be used in the iTunes Search API to decide if the app needs to be updated
	ExternalVersionId *string `json:"externalVersionId,omitempty"`
}

// NewComputerApplication instantiates a new ComputerApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerApplication() *ComputerApplication {
	this := ComputerApplication{}
	return &this
}

// NewComputerApplicationWithDefaults instantiates a new ComputerApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerApplicationWithDefaults() *ComputerApplication {
	this := ComputerApplication{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputerApplication) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerApplication) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputerApplication) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputerApplication) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ComputerApplication) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerApplication) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ComputerApplication) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ComputerApplication) SetPath(v string) {
	o.Path = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ComputerApplication) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerApplication) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ComputerApplication) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ComputerApplication) SetVersion(v string) {
	o.Version = &v
}

// GetMacAppStore returns the MacAppStore field value if set, zero value otherwise.
func (o *ComputerApplication) GetMacAppStore() bool {
	if o == nil || IsNil(o.MacAppStore) {
		var ret bool
		return ret
	}
	return *o.MacAppStore
}

// GetMacAppStoreOk returns a tuple with the MacAppStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerApplication) GetMacAppStoreOk() (*bool, bool) {
	if o == nil || IsNil(o.MacAppStore) {
		return nil, false
	}
	return o.MacAppStore, true
}

// HasMacAppStore returns a boolean if a field has been set.
func (o *ComputerApplication) HasMacAppStore() bool {
	if o != nil && !IsNil(o.MacAppStore) {
		return true
	}

	return false
}

// SetMacAppStore gets a reference to the given bool and assigns it to the MacAppStore field.
func (o *ComputerApplication) SetMacAppStore(v bool) {
	o.MacAppStore = &v
}

// GetSizeMegabytes returns the SizeMegabytes field value if set, zero value otherwise.
func (o *ComputerApplication) GetSizeMegabytes() int32 {
	if o == nil || IsNil(o.SizeMegabytes) {
		var ret int32
		return ret
	}
	return *o.SizeMegabytes
}

// GetSizeMegabytesOk returns a tuple with the SizeMegabytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerApplication) GetSizeMegabytesOk() (*int32, bool) {
	if o == nil || IsNil(o.SizeMegabytes) {
		return nil, false
	}
	return o.SizeMegabytes, true
}

// HasSizeMegabytes returns a boolean if a field has been set.
func (o *ComputerApplication) HasSizeMegabytes() bool {
	if o != nil && !IsNil(o.SizeMegabytes) {
		return true
	}

	return false
}

// SetSizeMegabytes gets a reference to the given int32 and assigns it to the SizeMegabytes field.
func (o *ComputerApplication) SetSizeMegabytes(v int32) {
	o.SizeMegabytes = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *ComputerApplication) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerApplication) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *ComputerApplication) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *ComputerApplication) SetBundleId(v string) {
	o.BundleId = &v
}

// GetUpdateAvailable returns the UpdateAvailable field value if set, zero value otherwise.
func (o *ComputerApplication) GetUpdateAvailable() bool {
	if o == nil || IsNil(o.UpdateAvailable) {
		var ret bool
		return ret
	}
	return *o.UpdateAvailable
}

// GetUpdateAvailableOk returns a tuple with the UpdateAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerApplication) GetUpdateAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateAvailable) {
		return nil, false
	}
	return o.UpdateAvailable, true
}

// HasUpdateAvailable returns a boolean if a field has been set.
func (o *ComputerApplication) HasUpdateAvailable() bool {
	if o != nil && !IsNil(o.UpdateAvailable) {
		return true
	}

	return false
}

// SetUpdateAvailable gets a reference to the given bool and assigns it to the UpdateAvailable field.
func (o *ComputerApplication) SetUpdateAvailable(v bool) {
	o.UpdateAvailable = &v
}

// GetExternalVersionId returns the ExternalVersionId field value if set, zero value otherwise.
func (o *ComputerApplication) GetExternalVersionId() string {
	if o == nil || IsNil(o.ExternalVersionId) {
		var ret string
		return ret
	}
	return *o.ExternalVersionId
}

// GetExternalVersionIdOk returns a tuple with the ExternalVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerApplication) GetExternalVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalVersionId) {
		return nil, false
	}
	return o.ExternalVersionId, true
}

// HasExternalVersionId returns a boolean if a field has been set.
func (o *ComputerApplication) HasExternalVersionId() bool {
	if o != nil && !IsNil(o.ExternalVersionId) {
		return true
	}

	return false
}

// SetExternalVersionId gets a reference to the given string and assigns it to the ExternalVersionId field.
func (o *ComputerApplication) SetExternalVersionId(v string) {
	o.ExternalVersionId = &v
}

func (o ComputerApplication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComputerApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.MacAppStore) {
		toSerialize["macAppStore"] = o.MacAppStore
	}
	if !IsNil(o.SizeMegabytes) {
		toSerialize["sizeMegabytes"] = o.SizeMegabytes
	}
	if !IsNil(o.BundleId) {
		toSerialize["bundleId"] = o.BundleId
	}
	if !IsNil(o.UpdateAvailable) {
		toSerialize["updateAvailable"] = o.UpdateAvailable
	}
	if !IsNil(o.ExternalVersionId) {
		toSerialize["externalVersionId"] = o.ExternalVersionId
	}
	return toSerialize, nil
}

type NullableComputerApplication struct {
	value *ComputerApplication
	isSet bool
}

func (v NullableComputerApplication) Get() *ComputerApplication {
	return v.value
}

func (v *NullableComputerApplication) Set(val *ComputerApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerApplication(val *ComputerApplication) *NullableComputerApplication {
	return &NullableComputerApplication{value: val, isSet: true}
}

func (v NullableComputerApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


