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

// SearchPatchPolicyLogParams struct for SearchPatchPolicyLogParams
type SearchPatchPolicyLogParams struct {
	PageNumber *int32 `json:"pageNumber,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	OrderBy []OrderBy `json:"orderBy,omitempty"`
	Filter []Filter `json:"filter,omitempty"`
	IsLoadToEnd *bool `json:"isLoadToEnd,omitempty"`
	PatchPolicyId *int32 `json:"patchPolicyId,omitempty"`
	IsLatest *bool `json:"isLatest,omitempty"`
}

// NewSearchPatchPolicyLogParams instantiates a new SearchPatchPolicyLogParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchPatchPolicyLogParams() *SearchPatchPolicyLogParams {
	this := SearchPatchPolicyLogParams{}
	var isLatest bool = false
	this.IsLatest = &isLatest
	return &this
}

// NewSearchPatchPolicyLogParamsWithDefaults instantiates a new SearchPatchPolicyLogParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchPatchPolicyLogParamsWithDefaults() *SearchPatchPolicyLogParams {
	this := SearchPatchPolicyLogParams{}
	var isLatest bool = false
	this.IsLatest = &isLatest
	return &this
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *SearchPatchPolicyLogParams) GetPageNumber() int32 {
	if o == nil || o.PageNumber == nil {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPatchPolicyLogParams) GetPageNumberOk() (*int32, bool) {
	if o == nil || o.PageNumber == nil {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *SearchPatchPolicyLogParams) HasPageNumber() bool {
	if o != nil && o.PageNumber != nil {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *SearchPatchPolicyLogParams) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *SearchPatchPolicyLogParams) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPatchPolicyLogParams) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *SearchPatchPolicyLogParams) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *SearchPatchPolicyLogParams) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *SearchPatchPolicyLogParams) GetOrderBy() []OrderBy {
	if o == nil || o.OrderBy == nil {
		var ret []OrderBy
		return ret
	}
	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPatchPolicyLogParams) GetOrderByOk() ([]OrderBy, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *SearchPatchPolicyLogParams) HasOrderBy() bool {
	if o != nil && o.OrderBy != nil {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given []OrderBy and assigns it to the OrderBy field.
func (o *SearchPatchPolicyLogParams) SetOrderBy(v []OrderBy) {
	o.OrderBy = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *SearchPatchPolicyLogParams) GetFilter() []Filter {
	if o == nil || o.Filter == nil {
		var ret []Filter
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPatchPolicyLogParams) GetFilterOk() ([]Filter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *SearchPatchPolicyLogParams) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []Filter and assigns it to the Filter field.
func (o *SearchPatchPolicyLogParams) SetFilter(v []Filter) {
	o.Filter = v
}

// GetIsLoadToEnd returns the IsLoadToEnd field value if set, zero value otherwise.
func (o *SearchPatchPolicyLogParams) GetIsLoadToEnd() bool {
	if o == nil || o.IsLoadToEnd == nil {
		var ret bool
		return ret
	}
	return *o.IsLoadToEnd
}

// GetIsLoadToEndOk returns a tuple with the IsLoadToEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPatchPolicyLogParams) GetIsLoadToEndOk() (*bool, bool) {
	if o == nil || o.IsLoadToEnd == nil {
		return nil, false
	}
	return o.IsLoadToEnd, true
}

// HasIsLoadToEnd returns a boolean if a field has been set.
func (o *SearchPatchPolicyLogParams) HasIsLoadToEnd() bool {
	if o != nil && o.IsLoadToEnd != nil {
		return true
	}

	return false
}

// SetIsLoadToEnd gets a reference to the given bool and assigns it to the IsLoadToEnd field.
func (o *SearchPatchPolicyLogParams) SetIsLoadToEnd(v bool) {
	o.IsLoadToEnd = &v
}

// GetPatchPolicyId returns the PatchPolicyId field value if set, zero value otherwise.
func (o *SearchPatchPolicyLogParams) GetPatchPolicyId() int32 {
	if o == nil || o.PatchPolicyId == nil {
		var ret int32
		return ret
	}
	return *o.PatchPolicyId
}

// GetPatchPolicyIdOk returns a tuple with the PatchPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPatchPolicyLogParams) GetPatchPolicyIdOk() (*int32, bool) {
	if o == nil || o.PatchPolicyId == nil {
		return nil, false
	}
	return o.PatchPolicyId, true
}

// HasPatchPolicyId returns a boolean if a field has been set.
func (o *SearchPatchPolicyLogParams) HasPatchPolicyId() bool {
	if o != nil && o.PatchPolicyId != nil {
		return true
	}

	return false
}

// SetPatchPolicyId gets a reference to the given int32 and assigns it to the PatchPolicyId field.
func (o *SearchPatchPolicyLogParams) SetPatchPolicyId(v int32) {
	o.PatchPolicyId = &v
}

// GetIsLatest returns the IsLatest field value if set, zero value otherwise.
func (o *SearchPatchPolicyLogParams) GetIsLatest() bool {
	if o == nil || o.IsLatest == nil {
		var ret bool
		return ret
	}
	return *o.IsLatest
}

// GetIsLatestOk returns a tuple with the IsLatest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPatchPolicyLogParams) GetIsLatestOk() (*bool, bool) {
	if o == nil || o.IsLatest == nil {
		return nil, false
	}
	return o.IsLatest, true
}

// HasIsLatest returns a boolean if a field has been set.
func (o *SearchPatchPolicyLogParams) HasIsLatest() bool {
	if o != nil && o.IsLatest != nil {
		return true
	}

	return false
}

// SetIsLatest gets a reference to the given bool and assigns it to the IsLatest field.
func (o *SearchPatchPolicyLogParams) SetIsLatest(v bool) {
	o.IsLatest = &v
}

func (o SearchPatchPolicyLogParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PageNumber != nil {
		toSerialize["pageNumber"] = o.PageNumber
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.OrderBy != nil {
		toSerialize["orderBy"] = o.OrderBy
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.IsLoadToEnd != nil {
		toSerialize["isLoadToEnd"] = o.IsLoadToEnd
	}
	if o.PatchPolicyId != nil {
		toSerialize["patchPolicyId"] = o.PatchPolicyId
	}
	if o.IsLatest != nil {
		toSerialize["isLatest"] = o.IsLatest
	}
	return json.Marshal(toSerialize)
}

type NullableSearchPatchPolicyLogParams struct {
	value *SearchPatchPolicyLogParams
	isSet bool
}

func (v NullableSearchPatchPolicyLogParams) Get() *SearchPatchPolicyLogParams {
	return v.value
}

func (v *NullableSearchPatchPolicyLogParams) Set(val *SearchPatchPolicyLogParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchPatchPolicyLogParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchPatchPolicyLogParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchPatchPolicyLogParams(val *SearchPatchPolicyLogParams) *NullableSearchPatchPolicyLogParams {
	return &NullableSearchPatchPolicyLogParams{value: val, isSet: true}
}

func (v NullableSearchPatchPolicyLogParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchPatchPolicyLogParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

