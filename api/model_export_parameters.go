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

// ExportParameters struct for ExportParameters
type ExportParameters struct {
	Page NullableInt32 `json:"page,omitempty"`
	PageSize NullableInt32 `json:"pageSize,omitempty"`
	// Sorting criteria in the format: [<property>[:asc/desc]. Default direction when not stated is ascending.
	Sort []string `json:"sort,omitempty"`
	Filter NullableString `json:"filter,omitempty"`
	// Used to change default order or ignore some of the fields. When null or empty array, all fields will be exported.
	Fields []ExportField `json:"fields,omitempty"`
}

// NewExportParameters instantiates a new ExportParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportParameters() *ExportParameters {
	this := ExportParameters{}
	var page int32 = 0
	this.Page = *NewNullableInt32(&page)
	var pageSize int32 = 100
	this.PageSize = *NewNullableInt32(&pageSize)
	return &this
}

// NewExportParametersWithDefaults instantiates a new ExportParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportParametersWithDefaults() *ExportParameters {
	this := ExportParameters{}
	var page int32 = 0
	this.Page = *NewNullableInt32(&page)
	var pageSize int32 = 100
	this.PageSize = *NewNullableInt32(&pageSize)
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportParameters) GetPage() int32 {
	if o == nil || o.Page.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Page.Get()
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportParameters) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Page.Get(), o.Page.IsSet()
}

// HasPage returns a boolean if a field has been set.
func (o *ExportParameters) HasPage() bool {
	if o != nil && o.Page.IsSet() {
		return true
	}

	return false
}

// SetPage gets a reference to the given NullableInt32 and assigns it to the Page field.
func (o *ExportParameters) SetPage(v int32) {
	o.Page.Set(&v)
}
// SetPageNil sets the value for Page to be an explicit nil
func (o *ExportParameters) SetPageNil() {
	o.Page.Set(nil)
}

// UnsetPage ensures that no value is present for Page, not even an explicit nil
func (o *ExportParameters) UnsetPage() {
	o.Page.Unset()
}

// GetPageSize returns the PageSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportParameters) GetPageSize() int32 {
	if o == nil || o.PageSize.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PageSize.Get()
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportParameters) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageSize.Get(), o.PageSize.IsSet()
}

// HasPageSize returns a boolean if a field has been set.
func (o *ExportParameters) HasPageSize() bool {
	if o != nil && o.PageSize.IsSet() {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given NullableInt32 and assigns it to the PageSize field.
func (o *ExportParameters) SetPageSize(v int32) {
	o.PageSize.Set(&v)
}
// SetPageSizeNil sets the value for PageSize to be an explicit nil
func (o *ExportParameters) SetPageSizeNil() {
	o.PageSize.Set(nil)
}

// UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
func (o *ExportParameters) UnsetPageSize() {
	o.PageSize.Unset()
}

// GetSort returns the Sort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportParameters) GetSort() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportParameters) GetSortOk() ([]string, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ExportParameters) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given []string and assigns it to the Sort field.
func (o *ExportParameters) SetSort(v []string) {
	o.Sort = v
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportParameters) GetFilter() string {
	if o == nil || o.Filter.Get() == nil {
		var ret string
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportParameters) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a field has been set.
func (o *ExportParameters) HasFilter() bool {
	if o != nil && o.Filter.IsSet() {
		return true
	}

	return false
}

// SetFilter gets a reference to the given NullableString and assigns it to the Filter field.
func (o *ExportParameters) SetFilter(v string) {
	o.Filter.Set(&v)
}
// SetFilterNil sets the value for Filter to be an explicit nil
func (o *ExportParameters) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnsetFilter ensures that no value is present for Filter, not even an explicit nil
func (o *ExportParameters) UnsetFilter() {
	o.Filter.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportParameters) GetFields() []ExportField {
	if o == nil {
		var ret []ExportField
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportParameters) GetFieldsOk() ([]ExportField, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ExportParameters) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []ExportField and assigns it to the Fields field.
func (o *ExportParameters) SetFields(v []ExportField) {
	o.Fields = v
}

func (o ExportParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Page.IsSet() {
		toSerialize["page"] = o.Page.Get()
	}
	if o.PageSize.IsSet() {
		toSerialize["pageSize"] = o.PageSize.Get()
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Filter.IsSet() {
		toSerialize["filter"] = o.Filter.Get()
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableExportParameters struct {
	value *ExportParameters
	isSet bool
}

func (v NullableExportParameters) Get() *ExportParameters {
	return v.value
}

func (v *NullableExportParameters) Set(val *ExportParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableExportParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableExportParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportParameters(val *ExportParameters) *NullableExportParameters {
	return &NullableExportParameters{value: val, isSet: true}
}

func (v NullableExportParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


