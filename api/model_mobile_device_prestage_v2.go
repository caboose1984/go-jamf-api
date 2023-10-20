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

// checks if the MobileDevicePrestageV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MobileDevicePrestageV2{}

// MobileDevicePrestageV2 struct for MobileDevicePrestageV2
type MobileDevicePrestageV2 struct {
	DisplayName string `json:"displayName"`
	Mandatory bool `json:"mandatory"`
	MdmRemovable bool `json:"mdmRemovable"`
	SupportPhoneNumber string `json:"supportPhoneNumber"`
	SupportEmailAddress string `json:"supportEmailAddress"`
	Department string `json:"department"`
	DefaultPrestage bool `json:"defaultPrestage"`
	EnrollmentSiteId string `json:"enrollmentSiteId"`
	KeepExistingSiteMembership bool `json:"keepExistingSiteMembership"`
	KeepExistingLocationInformation bool `json:"keepExistingLocationInformation"`
	RequireAuthentication bool `json:"requireAuthentication"`
	AuthenticationPrompt string `json:"authenticationPrompt"`
	PreventActivationLock bool `json:"preventActivationLock"`
	EnableDeviceBasedActivationLock bool `json:"enableDeviceBasedActivationLock"`
	DeviceEnrollmentProgramInstanceId string `json:"deviceEnrollmentProgramInstanceId"`
	SkipSetupItems *map[string]bool `json:"skipSetupItems,omitempty"`
	LocationInformation LocationInformationV2 `json:"locationInformation"`
	PurchasingInformation PrestagePurchasingInformationV2 `json:"purchasingInformation"`
	// The Base64 encoded PEM Certificate
	AnchorCertificates []string `json:"anchorCertificates,omitempty"`
	EnrollmentCustomizationId *string `json:"enrollmentCustomizationId,omitempty"`
	Language *string `json:"language,omitempty"`
	Region *string `json:"region,omitempty"`
	AutoAdvanceSetup bool `json:"autoAdvanceSetup"`
	AllowPairing bool `json:"allowPairing"`
	MultiUser bool `json:"multiUser"`
	Supervised bool `json:"supervised"`
	MaximumSharedAccounts int32 `json:"maximumSharedAccounts"`
	ConfigureDeviceBeforeSetupAssistant bool `json:"configureDeviceBeforeSetupAssistant"`
	Names *MobileDevicePrestageNamesV2 `json:"names,omitempty"`
	SendTimezone bool `json:"sendTimezone"`
	Timezone string `json:"timezone"`
	StorageQuotaSizeMegabytes int32 `json:"storageQuotaSizeMegabytes"`
	UseStorageQuotaSize bool `json:"useStorageQuotaSize"`
	TemporarySessionOnly *bool `json:"temporarySessionOnly,omitempty"`
	EnforceTemporarySessionTimeout *bool `json:"enforceTemporarySessionTimeout,omitempty"`
	TemporarySessionTimeout *int32 `json:"temporarySessionTimeout,omitempty"`
	EnforceUserSessionTimeout *bool `json:"enforceUserSessionTimeout,omitempty"`
	UserSessionTimeout *int32 `json:"userSessionTimeout,omitempty"`
}

// NewMobileDevicePrestageV2 instantiates a new MobileDevicePrestageV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileDevicePrestageV2(displayName string, mandatory bool, mdmRemovable bool, supportPhoneNumber string, supportEmailAddress string, department string, defaultPrestage bool, enrollmentSiteId string, keepExistingSiteMembership bool, keepExistingLocationInformation bool, requireAuthentication bool, authenticationPrompt string, preventActivationLock bool, enableDeviceBasedActivationLock bool, deviceEnrollmentProgramInstanceId string, locationInformation LocationInformationV2, purchasingInformation PrestagePurchasingInformationV2, autoAdvanceSetup bool, allowPairing bool, multiUser bool, supervised bool, maximumSharedAccounts int32, configureDeviceBeforeSetupAssistant bool, sendTimezone bool, timezone string, storageQuotaSizeMegabytes int32, useStorageQuotaSize bool) *MobileDevicePrestageV2 {
	this := MobileDevicePrestageV2{}
	this.DisplayName = displayName
	this.Mandatory = mandatory
	this.MdmRemovable = mdmRemovable
	this.SupportPhoneNumber = supportPhoneNumber
	this.SupportEmailAddress = supportEmailAddress
	this.Department = department
	this.DefaultPrestage = defaultPrestage
	this.EnrollmentSiteId = enrollmentSiteId
	this.KeepExistingSiteMembership = keepExistingSiteMembership
	this.KeepExistingLocationInformation = keepExistingLocationInformation
	this.RequireAuthentication = requireAuthentication
	this.AuthenticationPrompt = authenticationPrompt
	this.PreventActivationLock = preventActivationLock
	this.EnableDeviceBasedActivationLock = enableDeviceBasedActivationLock
	this.DeviceEnrollmentProgramInstanceId = deviceEnrollmentProgramInstanceId
	this.LocationInformation = locationInformation
	this.PurchasingInformation = purchasingInformation
	this.AutoAdvanceSetup = autoAdvanceSetup
	this.AllowPairing = allowPairing
	this.MultiUser = multiUser
	this.Supervised = supervised
	this.MaximumSharedAccounts = maximumSharedAccounts
	this.ConfigureDeviceBeforeSetupAssistant = configureDeviceBeforeSetupAssistant
	this.SendTimezone = sendTimezone
	this.Timezone = timezone
	this.StorageQuotaSizeMegabytes = storageQuotaSizeMegabytes
	this.UseStorageQuotaSize = useStorageQuotaSize
	return &this
}

// NewMobileDevicePrestageV2WithDefaults instantiates a new MobileDevicePrestageV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileDevicePrestageV2WithDefaults() *MobileDevicePrestageV2 {
	this := MobileDevicePrestageV2{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *MobileDevicePrestageV2) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *MobileDevicePrestageV2) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetMandatory returns the Mandatory field value
func (o *MobileDevicePrestageV2) GetMandatory() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Mandatory
}

// GetMandatoryOk returns a tuple with the Mandatory field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetMandatoryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mandatory, true
}

// SetMandatory sets field value
func (o *MobileDevicePrestageV2) SetMandatory(v bool) {
	o.Mandatory = v
}

// GetMdmRemovable returns the MdmRemovable field value
func (o *MobileDevicePrestageV2) GetMdmRemovable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MdmRemovable
}

// GetMdmRemovableOk returns a tuple with the MdmRemovable field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetMdmRemovableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MdmRemovable, true
}

// SetMdmRemovable sets field value
func (o *MobileDevicePrestageV2) SetMdmRemovable(v bool) {
	o.MdmRemovable = v
}

// GetSupportPhoneNumber returns the SupportPhoneNumber field value
func (o *MobileDevicePrestageV2) GetSupportPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportPhoneNumber
}

// GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetSupportPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportPhoneNumber, true
}

// SetSupportPhoneNumber sets field value
func (o *MobileDevicePrestageV2) SetSupportPhoneNumber(v string) {
	o.SupportPhoneNumber = v
}

// GetSupportEmailAddress returns the SupportEmailAddress field value
func (o *MobileDevicePrestageV2) GetSupportEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportEmailAddress
}

// GetSupportEmailAddressOk returns a tuple with the SupportEmailAddress field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetSupportEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportEmailAddress, true
}

// SetSupportEmailAddress sets field value
func (o *MobileDevicePrestageV2) SetSupportEmailAddress(v string) {
	o.SupportEmailAddress = v
}

// GetDepartment returns the Department field value
func (o *MobileDevicePrestageV2) GetDepartment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Department
}

// GetDepartmentOk returns a tuple with the Department field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetDepartmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Department, true
}

// SetDepartment sets field value
func (o *MobileDevicePrestageV2) SetDepartment(v string) {
	o.Department = v
}

// GetDefaultPrestage returns the DefaultPrestage field value
func (o *MobileDevicePrestageV2) GetDefaultPrestage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultPrestage
}

// GetDefaultPrestageOk returns a tuple with the DefaultPrestage field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetDefaultPrestageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultPrestage, true
}

// SetDefaultPrestage sets field value
func (o *MobileDevicePrestageV2) SetDefaultPrestage(v bool) {
	o.DefaultPrestage = v
}

// GetEnrollmentSiteId returns the EnrollmentSiteId field value
func (o *MobileDevicePrestageV2) GetEnrollmentSiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnrollmentSiteId
}

// GetEnrollmentSiteIdOk returns a tuple with the EnrollmentSiteId field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetEnrollmentSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnrollmentSiteId, true
}

// SetEnrollmentSiteId sets field value
func (o *MobileDevicePrestageV2) SetEnrollmentSiteId(v string) {
	o.EnrollmentSiteId = v
}

// GetKeepExistingSiteMembership returns the KeepExistingSiteMembership field value
func (o *MobileDevicePrestageV2) GetKeepExistingSiteMembership() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.KeepExistingSiteMembership
}

// GetKeepExistingSiteMembershipOk returns a tuple with the KeepExistingSiteMembership field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetKeepExistingSiteMembershipOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeepExistingSiteMembership, true
}

// SetKeepExistingSiteMembership sets field value
func (o *MobileDevicePrestageV2) SetKeepExistingSiteMembership(v bool) {
	o.KeepExistingSiteMembership = v
}

// GetKeepExistingLocationInformation returns the KeepExistingLocationInformation field value
func (o *MobileDevicePrestageV2) GetKeepExistingLocationInformation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.KeepExistingLocationInformation
}

// GetKeepExistingLocationInformationOk returns a tuple with the KeepExistingLocationInformation field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetKeepExistingLocationInformationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeepExistingLocationInformation, true
}

// SetKeepExistingLocationInformation sets field value
func (o *MobileDevicePrestageV2) SetKeepExistingLocationInformation(v bool) {
	o.KeepExistingLocationInformation = v
}

// GetRequireAuthentication returns the RequireAuthentication field value
func (o *MobileDevicePrestageV2) GetRequireAuthentication() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireAuthentication
}

// GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetRequireAuthenticationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireAuthentication, true
}

// SetRequireAuthentication sets field value
func (o *MobileDevicePrestageV2) SetRequireAuthentication(v bool) {
	o.RequireAuthentication = v
}

// GetAuthenticationPrompt returns the AuthenticationPrompt field value
func (o *MobileDevicePrestageV2) GetAuthenticationPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationPrompt
}

// GetAuthenticationPromptOk returns a tuple with the AuthenticationPrompt field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetAuthenticationPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationPrompt, true
}

// SetAuthenticationPrompt sets field value
func (o *MobileDevicePrestageV2) SetAuthenticationPrompt(v string) {
	o.AuthenticationPrompt = v
}

// GetPreventActivationLock returns the PreventActivationLock field value
func (o *MobileDevicePrestageV2) GetPreventActivationLock() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PreventActivationLock
}

// GetPreventActivationLockOk returns a tuple with the PreventActivationLock field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetPreventActivationLockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreventActivationLock, true
}

// SetPreventActivationLock sets field value
func (o *MobileDevicePrestageV2) SetPreventActivationLock(v bool) {
	o.PreventActivationLock = v
}

// GetEnableDeviceBasedActivationLock returns the EnableDeviceBasedActivationLock field value
func (o *MobileDevicePrestageV2) GetEnableDeviceBasedActivationLock() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableDeviceBasedActivationLock
}

// GetEnableDeviceBasedActivationLockOk returns a tuple with the EnableDeviceBasedActivationLock field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetEnableDeviceBasedActivationLockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableDeviceBasedActivationLock, true
}

// SetEnableDeviceBasedActivationLock sets field value
func (o *MobileDevicePrestageV2) SetEnableDeviceBasedActivationLock(v bool) {
	o.EnableDeviceBasedActivationLock = v
}

// GetDeviceEnrollmentProgramInstanceId returns the DeviceEnrollmentProgramInstanceId field value
func (o *MobileDevicePrestageV2) GetDeviceEnrollmentProgramInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceEnrollmentProgramInstanceId
}

// GetDeviceEnrollmentProgramInstanceIdOk returns a tuple with the DeviceEnrollmentProgramInstanceId field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetDeviceEnrollmentProgramInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceEnrollmentProgramInstanceId, true
}

// SetDeviceEnrollmentProgramInstanceId sets field value
func (o *MobileDevicePrestageV2) SetDeviceEnrollmentProgramInstanceId(v string) {
	o.DeviceEnrollmentProgramInstanceId = v
}

// GetSkipSetupItems returns the SkipSetupItems field value if set, zero value otherwise.
func (o *MobileDevicePrestageV2) GetSkipSetupItems() map[string]bool {
	if o == nil || IsNil(o.SkipSetupItems) {
		var ret map[string]bool
		return ret
	}
	return *o.SkipSetupItems
}

// GetSkipSetupItemsOk returns a tuple with the SkipSetupItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetSkipSetupItemsOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.SkipSetupItems) {
		return nil, false
	}
	return o.SkipSetupItems, true
}

// HasSkipSetupItems returns a boolean if a field has been set.
func (o *MobileDevicePrestageV2) HasSkipSetupItems() bool {
	if o != nil && !IsNil(o.SkipSetupItems) {
		return true
	}

	return false
}

// SetSkipSetupItems gets a reference to the given map[string]bool and assigns it to the SkipSetupItems field.
func (o *MobileDevicePrestageV2) SetSkipSetupItems(v map[string]bool) {
	o.SkipSetupItems = &v
}

// GetLocationInformation returns the LocationInformation field value
func (o *MobileDevicePrestageV2) GetLocationInformation() LocationInformationV2 {
	if o == nil {
		var ret LocationInformationV2
		return ret
	}

	return o.LocationInformation
}

// GetLocationInformationOk returns a tuple with the LocationInformation field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetLocationInformationOk() (*LocationInformationV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationInformation, true
}

// SetLocationInformation sets field value
func (o *MobileDevicePrestageV2) SetLocationInformation(v LocationInformationV2) {
	o.LocationInformation = v
}

// GetPurchasingInformation returns the PurchasingInformation field value
func (o *MobileDevicePrestageV2) GetPurchasingInformation() PrestagePurchasingInformationV2 {
	if o == nil {
		var ret PrestagePurchasingInformationV2
		return ret
	}

	return o.PurchasingInformation
}

// GetPurchasingInformationOk returns a tuple with the PurchasingInformation field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetPurchasingInformationOk() (*PrestagePurchasingInformationV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchasingInformation, true
}

// SetPurchasingInformation sets field value
func (o *MobileDevicePrestageV2) SetPurchasingInformation(v PrestagePurchasingInformationV2) {
	o.PurchasingInformation = v
}

// GetAnchorCertificates returns the AnchorCertificates field value if set, zero value otherwise.
func (o *MobileDevicePrestageV2) GetAnchorCertificates() []string {
	if o == nil || IsNil(o.AnchorCertificates) {
		var ret []string
		return ret
	}
	return o.AnchorCertificates
}

// GetAnchorCertificatesOk returns a tuple with the AnchorCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetAnchorCertificatesOk() ([]string, bool) {
	if o == nil || IsNil(o.AnchorCertificates) {
		return nil, false
	}
	return o.AnchorCertificates, true
}

// HasAnchorCertificates returns a boolean if a field has been set.
func (o *MobileDevicePrestageV2) HasAnchorCertificates() bool {
	if o != nil && !IsNil(o.AnchorCertificates) {
		return true
	}

	return false
}

// SetAnchorCertificates gets a reference to the given []string and assigns it to the AnchorCertificates field.
func (o *MobileDevicePrestageV2) SetAnchorCertificates(v []string) {
	o.AnchorCertificates = v
}

// GetEnrollmentCustomizationId returns the EnrollmentCustomizationId field value if set, zero value otherwise.
func (o *MobileDevicePrestageV2) GetEnrollmentCustomizationId() string {
	if o == nil || IsNil(o.EnrollmentCustomizationId) {
		var ret string
		return ret
	}
	return *o.EnrollmentCustomizationId
}

// GetEnrollmentCustomizationIdOk returns a tuple with the EnrollmentCustomizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetEnrollmentCustomizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnrollmentCustomizationId) {
		return nil, false
	}
	return o.EnrollmentCustomizationId, true
}

// HasEnrollmentCustomizationId returns a boolean if a field has been set.
func (o *MobileDevicePrestageV2) HasEnrollmentCustomizationId() bool {
	if o != nil && !IsNil(o.EnrollmentCustomizationId) {
		return true
	}

	return false
}

// SetEnrollmentCustomizationId gets a reference to the given string and assigns it to the EnrollmentCustomizationId field.
func (o *MobileDevicePrestageV2) SetEnrollmentCustomizationId(v string) {
	o.EnrollmentCustomizationId = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *MobileDevicePrestageV2) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *MobileDevicePrestageV2) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *MobileDevicePrestageV2) SetLanguage(v string) {
	o.Language = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *MobileDevicePrestageV2) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *MobileDevicePrestageV2) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *MobileDevicePrestageV2) SetRegion(v string) {
	o.Region = &v
}

// GetAutoAdvanceSetup returns the AutoAdvanceSetup field value
func (o *MobileDevicePrestageV2) GetAutoAdvanceSetup() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoAdvanceSetup
}

// GetAutoAdvanceSetupOk returns a tuple with the AutoAdvanceSetup field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetAutoAdvanceSetupOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoAdvanceSetup, true
}

// SetAutoAdvanceSetup sets field value
func (o *MobileDevicePrestageV2) SetAutoAdvanceSetup(v bool) {
	o.AutoAdvanceSetup = v
}

// GetAllowPairing returns the AllowPairing field value
func (o *MobileDevicePrestageV2) GetAllowPairing() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowPairing
}

// GetAllowPairingOk returns a tuple with the AllowPairing field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetAllowPairingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowPairing, true
}

// SetAllowPairing sets field value
func (o *MobileDevicePrestageV2) SetAllowPairing(v bool) {
	o.AllowPairing = v
}

// GetMultiUser returns the MultiUser field value
func (o *MobileDevicePrestageV2) GetMultiUser() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MultiUser
}

// GetMultiUserOk returns a tuple with the MultiUser field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetMultiUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MultiUser, true
}

// SetMultiUser sets field value
func (o *MobileDevicePrestageV2) SetMultiUser(v bool) {
	o.MultiUser = v
}

// GetSupervised returns the Supervised field value
func (o *MobileDevicePrestageV2) GetSupervised() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Supervised
}

// GetSupervisedOk returns a tuple with the Supervised field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetSupervisedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supervised, true
}

// SetSupervised sets field value
func (o *MobileDevicePrestageV2) SetSupervised(v bool) {
	o.Supervised = v
}

// GetMaximumSharedAccounts returns the MaximumSharedAccounts field value
func (o *MobileDevicePrestageV2) GetMaximumSharedAccounts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaximumSharedAccounts
}

// GetMaximumSharedAccountsOk returns a tuple with the MaximumSharedAccounts field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetMaximumSharedAccountsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaximumSharedAccounts, true
}

// SetMaximumSharedAccounts sets field value
func (o *MobileDevicePrestageV2) SetMaximumSharedAccounts(v int32) {
	o.MaximumSharedAccounts = v
}

// GetConfigureDeviceBeforeSetupAssistant returns the ConfigureDeviceBeforeSetupAssistant field value
func (o *MobileDevicePrestageV2) GetConfigureDeviceBeforeSetupAssistant() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ConfigureDeviceBeforeSetupAssistant
}

// GetConfigureDeviceBeforeSetupAssistantOk returns a tuple with the ConfigureDeviceBeforeSetupAssistant field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetConfigureDeviceBeforeSetupAssistantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigureDeviceBeforeSetupAssistant, true
}

// SetConfigureDeviceBeforeSetupAssistant sets field value
func (o *MobileDevicePrestageV2) SetConfigureDeviceBeforeSetupAssistant(v bool) {
	o.ConfigureDeviceBeforeSetupAssistant = v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *MobileDevicePrestageV2) GetNames() MobileDevicePrestageNamesV2 {
	if o == nil || IsNil(o.Names) {
		var ret MobileDevicePrestageNamesV2
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetNamesOk() (*MobileDevicePrestageNamesV2, bool) {
	if o == nil || IsNil(o.Names) {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *MobileDevicePrestageV2) HasNames() bool {
	if o != nil && !IsNil(o.Names) {
		return true
	}

	return false
}

// SetNames gets a reference to the given MobileDevicePrestageNamesV2 and assigns it to the Names field.
func (o *MobileDevicePrestageV2) SetNames(v MobileDevicePrestageNamesV2) {
	o.Names = &v
}

// GetSendTimezone returns the SendTimezone field value
func (o *MobileDevicePrestageV2) GetSendTimezone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SendTimezone
}

// GetSendTimezoneOk returns a tuple with the SendTimezone field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetSendTimezoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SendTimezone, true
}

// SetSendTimezone sets field value
func (o *MobileDevicePrestageV2) SetSendTimezone(v bool) {
	o.SendTimezone = v
}

// GetTimezone returns the Timezone field value
func (o *MobileDevicePrestageV2) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *MobileDevicePrestageV2) SetTimezone(v string) {
	o.Timezone = v
}

// GetStorageQuotaSizeMegabytes returns the StorageQuotaSizeMegabytes field value
func (o *MobileDevicePrestageV2) GetStorageQuotaSizeMegabytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StorageQuotaSizeMegabytes
}

// GetStorageQuotaSizeMegabytesOk returns a tuple with the StorageQuotaSizeMegabytes field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetStorageQuotaSizeMegabytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageQuotaSizeMegabytes, true
}

// SetStorageQuotaSizeMegabytes sets field value
func (o *MobileDevicePrestageV2) SetStorageQuotaSizeMegabytes(v int32) {
	o.StorageQuotaSizeMegabytes = v
}

// GetUseStorageQuotaSize returns the UseStorageQuotaSize field value
func (o *MobileDevicePrestageV2) GetUseStorageQuotaSize() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseStorageQuotaSize
}

// GetUseStorageQuotaSizeOk returns a tuple with the UseStorageQuotaSize field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetUseStorageQuotaSizeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseStorageQuotaSize, true
}

// SetUseStorageQuotaSize sets field value
func (o *MobileDevicePrestageV2) SetUseStorageQuotaSize(v bool) {
	o.UseStorageQuotaSize = v
}

// GetTemporarySessionOnly returns the TemporarySessionOnly field value if set, zero value otherwise.
func (o *MobileDevicePrestageV2) GetTemporarySessionOnly() bool {
	if o == nil || IsNil(o.TemporarySessionOnly) {
		var ret bool
		return ret
	}
	return *o.TemporarySessionOnly
}

// GetTemporarySessionOnlyOk returns a tuple with the TemporarySessionOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetTemporarySessionOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.TemporarySessionOnly) {
		return nil, false
	}
	return o.TemporarySessionOnly, true
}

// HasTemporarySessionOnly returns a boolean if a field has been set.
func (o *MobileDevicePrestageV2) HasTemporarySessionOnly() bool {
	if o != nil && !IsNil(o.TemporarySessionOnly) {
		return true
	}

	return false
}

// SetTemporarySessionOnly gets a reference to the given bool and assigns it to the TemporarySessionOnly field.
func (o *MobileDevicePrestageV2) SetTemporarySessionOnly(v bool) {
	o.TemporarySessionOnly = &v
}

// GetEnforceTemporarySessionTimeout returns the EnforceTemporarySessionTimeout field value if set, zero value otherwise.
func (o *MobileDevicePrestageV2) GetEnforceTemporarySessionTimeout() bool {
	if o == nil || IsNil(o.EnforceTemporarySessionTimeout) {
		var ret bool
		return ret
	}
	return *o.EnforceTemporarySessionTimeout
}

// GetEnforceTemporarySessionTimeoutOk returns a tuple with the EnforceTemporarySessionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetEnforceTemporarySessionTimeoutOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceTemporarySessionTimeout) {
		return nil, false
	}
	return o.EnforceTemporarySessionTimeout, true
}

// HasEnforceTemporarySessionTimeout returns a boolean if a field has been set.
func (o *MobileDevicePrestageV2) HasEnforceTemporarySessionTimeout() bool {
	if o != nil && !IsNil(o.EnforceTemporarySessionTimeout) {
		return true
	}

	return false
}

// SetEnforceTemporarySessionTimeout gets a reference to the given bool and assigns it to the EnforceTemporarySessionTimeout field.
func (o *MobileDevicePrestageV2) SetEnforceTemporarySessionTimeout(v bool) {
	o.EnforceTemporarySessionTimeout = &v
}

// GetTemporarySessionTimeout returns the TemporarySessionTimeout field value if set, zero value otherwise.
func (o *MobileDevicePrestageV2) GetTemporarySessionTimeout() int32 {
	if o == nil || IsNil(o.TemporarySessionTimeout) {
		var ret int32
		return ret
	}
	return *o.TemporarySessionTimeout
}

// GetTemporarySessionTimeoutOk returns a tuple with the TemporarySessionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetTemporarySessionTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.TemporarySessionTimeout) {
		return nil, false
	}
	return o.TemporarySessionTimeout, true
}

// HasTemporarySessionTimeout returns a boolean if a field has been set.
func (o *MobileDevicePrestageV2) HasTemporarySessionTimeout() bool {
	if o != nil && !IsNil(o.TemporarySessionTimeout) {
		return true
	}

	return false
}

// SetTemporarySessionTimeout gets a reference to the given int32 and assigns it to the TemporarySessionTimeout field.
func (o *MobileDevicePrestageV2) SetTemporarySessionTimeout(v int32) {
	o.TemporarySessionTimeout = &v
}

// GetEnforceUserSessionTimeout returns the EnforceUserSessionTimeout field value if set, zero value otherwise.
func (o *MobileDevicePrestageV2) GetEnforceUserSessionTimeout() bool {
	if o == nil || IsNil(o.EnforceUserSessionTimeout) {
		var ret bool
		return ret
	}
	return *o.EnforceUserSessionTimeout
}

// GetEnforceUserSessionTimeoutOk returns a tuple with the EnforceUserSessionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetEnforceUserSessionTimeoutOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceUserSessionTimeout) {
		return nil, false
	}
	return o.EnforceUserSessionTimeout, true
}

// HasEnforceUserSessionTimeout returns a boolean if a field has been set.
func (o *MobileDevicePrestageV2) HasEnforceUserSessionTimeout() bool {
	if o != nil && !IsNil(o.EnforceUserSessionTimeout) {
		return true
	}

	return false
}

// SetEnforceUserSessionTimeout gets a reference to the given bool and assigns it to the EnforceUserSessionTimeout field.
func (o *MobileDevicePrestageV2) SetEnforceUserSessionTimeout(v bool) {
	o.EnforceUserSessionTimeout = &v
}

// GetUserSessionTimeout returns the UserSessionTimeout field value if set, zero value otherwise.
func (o *MobileDevicePrestageV2) GetUserSessionTimeout() int32 {
	if o == nil || IsNil(o.UserSessionTimeout) {
		var ret int32
		return ret
	}
	return *o.UserSessionTimeout
}

// GetUserSessionTimeoutOk returns a tuple with the UserSessionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageV2) GetUserSessionTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.UserSessionTimeout) {
		return nil, false
	}
	return o.UserSessionTimeout, true
}

// HasUserSessionTimeout returns a boolean if a field has been set.
func (o *MobileDevicePrestageV2) HasUserSessionTimeout() bool {
	if o != nil && !IsNil(o.UserSessionTimeout) {
		return true
	}

	return false
}

// SetUserSessionTimeout gets a reference to the given int32 and assigns it to the UserSessionTimeout field.
func (o *MobileDevicePrestageV2) SetUserSessionTimeout(v int32) {
	o.UserSessionTimeout = &v
}

func (o MobileDevicePrestageV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobileDevicePrestageV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["mandatory"] = o.Mandatory
	toSerialize["mdmRemovable"] = o.MdmRemovable
	toSerialize["supportPhoneNumber"] = o.SupportPhoneNumber
	toSerialize["supportEmailAddress"] = o.SupportEmailAddress
	toSerialize["department"] = o.Department
	toSerialize["defaultPrestage"] = o.DefaultPrestage
	toSerialize["enrollmentSiteId"] = o.EnrollmentSiteId
	toSerialize["keepExistingSiteMembership"] = o.KeepExistingSiteMembership
	toSerialize["keepExistingLocationInformation"] = o.KeepExistingLocationInformation
	toSerialize["requireAuthentication"] = o.RequireAuthentication
	toSerialize["authenticationPrompt"] = o.AuthenticationPrompt
	toSerialize["preventActivationLock"] = o.PreventActivationLock
	toSerialize["enableDeviceBasedActivationLock"] = o.EnableDeviceBasedActivationLock
	toSerialize["deviceEnrollmentProgramInstanceId"] = o.DeviceEnrollmentProgramInstanceId
	if !IsNil(o.SkipSetupItems) {
		toSerialize["skipSetupItems"] = o.SkipSetupItems
	}
	toSerialize["locationInformation"] = o.LocationInformation
	toSerialize["purchasingInformation"] = o.PurchasingInformation
	if !IsNil(o.AnchorCertificates) {
		toSerialize["anchorCertificates"] = o.AnchorCertificates
	}
	if !IsNil(o.EnrollmentCustomizationId) {
		toSerialize["enrollmentCustomizationId"] = o.EnrollmentCustomizationId
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	toSerialize["autoAdvanceSetup"] = o.AutoAdvanceSetup
	toSerialize["allowPairing"] = o.AllowPairing
	toSerialize["multiUser"] = o.MultiUser
	toSerialize["supervised"] = o.Supervised
	toSerialize["maximumSharedAccounts"] = o.MaximumSharedAccounts
	toSerialize["configureDeviceBeforeSetupAssistant"] = o.ConfigureDeviceBeforeSetupAssistant
	if !IsNil(o.Names) {
		toSerialize["names"] = o.Names
	}
	toSerialize["sendTimezone"] = o.SendTimezone
	toSerialize["timezone"] = o.Timezone
	toSerialize["storageQuotaSizeMegabytes"] = o.StorageQuotaSizeMegabytes
	toSerialize["useStorageQuotaSize"] = o.UseStorageQuotaSize
	if !IsNil(o.TemporarySessionOnly) {
		toSerialize["temporarySessionOnly"] = o.TemporarySessionOnly
	}
	if !IsNil(o.EnforceTemporarySessionTimeout) {
		toSerialize["enforceTemporarySessionTimeout"] = o.EnforceTemporarySessionTimeout
	}
	if !IsNil(o.TemporarySessionTimeout) {
		toSerialize["temporarySessionTimeout"] = o.TemporarySessionTimeout
	}
	if !IsNil(o.EnforceUserSessionTimeout) {
		toSerialize["enforceUserSessionTimeout"] = o.EnforceUserSessionTimeout
	}
	if !IsNil(o.UserSessionTimeout) {
		toSerialize["userSessionTimeout"] = o.UserSessionTimeout
	}
	return toSerialize, nil
}

type NullableMobileDevicePrestageV2 struct {
	value *MobileDevicePrestageV2
	isSet bool
}

func (v NullableMobileDevicePrestageV2) Get() *MobileDevicePrestageV2 {
	return v.value
}

func (v *NullableMobileDevicePrestageV2) Set(val *MobileDevicePrestageV2) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDevicePrestageV2) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDevicePrestageV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDevicePrestageV2(val *MobileDevicePrestageV2) *NullableMobileDevicePrestageV2 {
	return &NullableMobileDevicePrestageV2{value: val, isSet: true}
}

func (v NullableMobileDevicePrestageV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDevicePrestageV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


