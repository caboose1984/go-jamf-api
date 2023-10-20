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

// checks if the ComputerHardware type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComputerHardware{}

// ComputerHardware struct for ComputerHardware
type ComputerHardware struct {
	Make *string `json:"make,omitempty"`
	Model *string `json:"model,omitempty"`
	ModelIdentifier *string `json:"modelIdentifier,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`
	// Processor Speed in MHz.
	ProcessorSpeedMhz *int64 `json:"processorSpeedMhz,omitempty"`
	ProcessorCount *int32 `json:"processorCount,omitempty"`
	CoreCount *int32 `json:"coreCount,omitempty"`
	ProcessorType *string `json:"processorType,omitempty"`
	ProcessorArchitecture *string `json:"processorArchitecture,omitempty"`
	BusSpeedMhz *int64 `json:"busSpeedMhz,omitempty"`
	// Cache Size in KB.
	CacheSizeKilobytes *int64 `json:"cacheSizeKilobytes,omitempty"`
	NetworkAdapterType *string `json:"networkAdapterType,omitempty"`
	MacAddress *string `json:"macAddress,omitempty"`
	AltNetworkAdapterType *string `json:"altNetworkAdapterType,omitempty"`
	AltMacAddress *string `json:"altMacAddress,omitempty"`
	// Total RAM Size in MB.
	TotalRamMegabytes *int64 `json:"totalRamMegabytes,omitempty"`
	// Available RAM slots.
	OpenRamSlots *int32 `json:"openRamSlots,omitempty"`
	// Remaining percentage of battery power.
	BatteryCapacityPercent *int32 `json:"batteryCapacityPercent,omitempty"`
	SmcVersion *string `json:"smcVersion,omitempty"`
	NicSpeed *string `json:"nicSpeed,omitempty"`
	OpticalDrive *string `json:"opticalDrive,omitempty"`
	BootRom *string `json:"bootRom,omitempty"`
	BleCapable *bool `json:"bleCapable,omitempty"`
	SupportsIosAppInstalls *bool `json:"supportsIosAppInstalls,omitempty"`
	AppleSilicon *bool `json:"appleSilicon,omitempty"`
	ExtensionAttributes []ComputerExtensionAttribute `json:"extensionAttributes,omitempty"`
}

// NewComputerHardware instantiates a new ComputerHardware object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerHardware() *ComputerHardware {
	this := ComputerHardware{}
	return &this
}

// NewComputerHardwareWithDefaults instantiates a new ComputerHardware object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerHardwareWithDefaults() *ComputerHardware {
	this := ComputerHardware{}
	return &this
}

// GetMake returns the Make field value if set, zero value otherwise.
func (o *ComputerHardware) GetMake() string {
	if o == nil || IsNil(o.Make) {
		var ret string
		return ret
	}
	return *o.Make
}

// GetMakeOk returns a tuple with the Make field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetMakeOk() (*string, bool) {
	if o == nil || IsNil(o.Make) {
		return nil, false
	}
	return o.Make, true
}

// HasMake returns a boolean if a field has been set.
func (o *ComputerHardware) HasMake() bool {
	if o != nil && !IsNil(o.Make) {
		return true
	}

	return false
}

// SetMake gets a reference to the given string and assigns it to the Make field.
func (o *ComputerHardware) SetMake(v string) {
	o.Make = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ComputerHardware) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ComputerHardware) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ComputerHardware) SetModel(v string) {
	o.Model = &v
}

// GetModelIdentifier returns the ModelIdentifier field value if set, zero value otherwise.
func (o *ComputerHardware) GetModelIdentifier() string {
	if o == nil || IsNil(o.ModelIdentifier) {
		var ret string
		return ret
	}
	return *o.ModelIdentifier
}

// GetModelIdentifierOk returns a tuple with the ModelIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetModelIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.ModelIdentifier) {
		return nil, false
	}
	return o.ModelIdentifier, true
}

// HasModelIdentifier returns a boolean if a field has been set.
func (o *ComputerHardware) HasModelIdentifier() bool {
	if o != nil && !IsNil(o.ModelIdentifier) {
		return true
	}

	return false
}

// SetModelIdentifier gets a reference to the given string and assigns it to the ModelIdentifier field.
func (o *ComputerHardware) SetModelIdentifier(v string) {
	o.ModelIdentifier = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *ComputerHardware) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *ComputerHardware) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *ComputerHardware) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetProcessorSpeedMhz returns the ProcessorSpeedMhz field value if set, zero value otherwise.
func (o *ComputerHardware) GetProcessorSpeedMhz() int64 {
	if o == nil || IsNil(o.ProcessorSpeedMhz) {
		var ret int64
		return ret
	}
	return *o.ProcessorSpeedMhz
}

// GetProcessorSpeedMhzOk returns a tuple with the ProcessorSpeedMhz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetProcessorSpeedMhzOk() (*int64, bool) {
	if o == nil || IsNil(o.ProcessorSpeedMhz) {
		return nil, false
	}
	return o.ProcessorSpeedMhz, true
}

// HasProcessorSpeedMhz returns a boolean if a field has been set.
func (o *ComputerHardware) HasProcessorSpeedMhz() bool {
	if o != nil && !IsNil(o.ProcessorSpeedMhz) {
		return true
	}

	return false
}

// SetProcessorSpeedMhz gets a reference to the given int64 and assigns it to the ProcessorSpeedMhz field.
func (o *ComputerHardware) SetProcessorSpeedMhz(v int64) {
	o.ProcessorSpeedMhz = &v
}

// GetProcessorCount returns the ProcessorCount field value if set, zero value otherwise.
func (o *ComputerHardware) GetProcessorCount() int32 {
	if o == nil || IsNil(o.ProcessorCount) {
		var ret int32
		return ret
	}
	return *o.ProcessorCount
}

// GetProcessorCountOk returns a tuple with the ProcessorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetProcessorCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ProcessorCount) {
		return nil, false
	}
	return o.ProcessorCount, true
}

// HasProcessorCount returns a boolean if a field has been set.
func (o *ComputerHardware) HasProcessorCount() bool {
	if o != nil && !IsNil(o.ProcessorCount) {
		return true
	}

	return false
}

// SetProcessorCount gets a reference to the given int32 and assigns it to the ProcessorCount field.
func (o *ComputerHardware) SetProcessorCount(v int32) {
	o.ProcessorCount = &v
}

// GetCoreCount returns the CoreCount field value if set, zero value otherwise.
func (o *ComputerHardware) GetCoreCount() int32 {
	if o == nil || IsNil(o.CoreCount) {
		var ret int32
		return ret
	}
	return *o.CoreCount
}

// GetCoreCountOk returns a tuple with the CoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetCoreCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CoreCount) {
		return nil, false
	}
	return o.CoreCount, true
}

// HasCoreCount returns a boolean if a field has been set.
func (o *ComputerHardware) HasCoreCount() bool {
	if o != nil && !IsNil(o.CoreCount) {
		return true
	}

	return false
}

// SetCoreCount gets a reference to the given int32 and assigns it to the CoreCount field.
func (o *ComputerHardware) SetCoreCount(v int32) {
	o.CoreCount = &v
}

// GetProcessorType returns the ProcessorType field value if set, zero value otherwise.
func (o *ComputerHardware) GetProcessorType() string {
	if o == nil || IsNil(o.ProcessorType) {
		var ret string
		return ret
	}
	return *o.ProcessorType
}

// GetProcessorTypeOk returns a tuple with the ProcessorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetProcessorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessorType) {
		return nil, false
	}
	return o.ProcessorType, true
}

// HasProcessorType returns a boolean if a field has been set.
func (o *ComputerHardware) HasProcessorType() bool {
	if o != nil && !IsNil(o.ProcessorType) {
		return true
	}

	return false
}

// SetProcessorType gets a reference to the given string and assigns it to the ProcessorType field.
func (o *ComputerHardware) SetProcessorType(v string) {
	o.ProcessorType = &v
}

// GetProcessorArchitecture returns the ProcessorArchitecture field value if set, zero value otherwise.
func (o *ComputerHardware) GetProcessorArchitecture() string {
	if o == nil || IsNil(o.ProcessorArchitecture) {
		var ret string
		return ret
	}
	return *o.ProcessorArchitecture
}

// GetProcessorArchitectureOk returns a tuple with the ProcessorArchitecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetProcessorArchitectureOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessorArchitecture) {
		return nil, false
	}
	return o.ProcessorArchitecture, true
}

// HasProcessorArchitecture returns a boolean if a field has been set.
func (o *ComputerHardware) HasProcessorArchitecture() bool {
	if o != nil && !IsNil(o.ProcessorArchitecture) {
		return true
	}

	return false
}

// SetProcessorArchitecture gets a reference to the given string and assigns it to the ProcessorArchitecture field.
func (o *ComputerHardware) SetProcessorArchitecture(v string) {
	o.ProcessorArchitecture = &v
}

// GetBusSpeedMhz returns the BusSpeedMhz field value if set, zero value otherwise.
func (o *ComputerHardware) GetBusSpeedMhz() int64 {
	if o == nil || IsNil(o.BusSpeedMhz) {
		var ret int64
		return ret
	}
	return *o.BusSpeedMhz
}

// GetBusSpeedMhzOk returns a tuple with the BusSpeedMhz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetBusSpeedMhzOk() (*int64, bool) {
	if o == nil || IsNil(o.BusSpeedMhz) {
		return nil, false
	}
	return o.BusSpeedMhz, true
}

// HasBusSpeedMhz returns a boolean if a field has been set.
func (o *ComputerHardware) HasBusSpeedMhz() bool {
	if o != nil && !IsNil(o.BusSpeedMhz) {
		return true
	}

	return false
}

// SetBusSpeedMhz gets a reference to the given int64 and assigns it to the BusSpeedMhz field.
func (o *ComputerHardware) SetBusSpeedMhz(v int64) {
	o.BusSpeedMhz = &v
}

// GetCacheSizeKilobytes returns the CacheSizeKilobytes field value if set, zero value otherwise.
func (o *ComputerHardware) GetCacheSizeKilobytes() int64 {
	if o == nil || IsNil(o.CacheSizeKilobytes) {
		var ret int64
		return ret
	}
	return *o.CacheSizeKilobytes
}

// GetCacheSizeKilobytesOk returns a tuple with the CacheSizeKilobytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetCacheSizeKilobytesOk() (*int64, bool) {
	if o == nil || IsNil(o.CacheSizeKilobytes) {
		return nil, false
	}
	return o.CacheSizeKilobytes, true
}

// HasCacheSizeKilobytes returns a boolean if a field has been set.
func (o *ComputerHardware) HasCacheSizeKilobytes() bool {
	if o != nil && !IsNil(o.CacheSizeKilobytes) {
		return true
	}

	return false
}

// SetCacheSizeKilobytes gets a reference to the given int64 and assigns it to the CacheSizeKilobytes field.
func (o *ComputerHardware) SetCacheSizeKilobytes(v int64) {
	o.CacheSizeKilobytes = &v
}

// GetNetworkAdapterType returns the NetworkAdapterType field value if set, zero value otherwise.
func (o *ComputerHardware) GetNetworkAdapterType() string {
	if o == nil || IsNil(o.NetworkAdapterType) {
		var ret string
		return ret
	}
	return *o.NetworkAdapterType
}

// GetNetworkAdapterTypeOk returns a tuple with the NetworkAdapterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetNetworkAdapterTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkAdapterType) {
		return nil, false
	}
	return o.NetworkAdapterType, true
}

// HasNetworkAdapterType returns a boolean if a field has been set.
func (o *ComputerHardware) HasNetworkAdapterType() bool {
	if o != nil && !IsNil(o.NetworkAdapterType) {
		return true
	}

	return false
}

// SetNetworkAdapterType gets a reference to the given string and assigns it to the NetworkAdapterType field.
func (o *ComputerHardware) SetNetworkAdapterType(v string) {
	o.NetworkAdapterType = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *ComputerHardware) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *ComputerHardware) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *ComputerHardware) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetAltNetworkAdapterType returns the AltNetworkAdapterType field value if set, zero value otherwise.
func (o *ComputerHardware) GetAltNetworkAdapterType() string {
	if o == nil || IsNil(o.AltNetworkAdapterType) {
		var ret string
		return ret
	}
	return *o.AltNetworkAdapterType
}

// GetAltNetworkAdapterTypeOk returns a tuple with the AltNetworkAdapterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetAltNetworkAdapterTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AltNetworkAdapterType) {
		return nil, false
	}
	return o.AltNetworkAdapterType, true
}

// HasAltNetworkAdapterType returns a boolean if a field has been set.
func (o *ComputerHardware) HasAltNetworkAdapterType() bool {
	if o != nil && !IsNil(o.AltNetworkAdapterType) {
		return true
	}

	return false
}

// SetAltNetworkAdapterType gets a reference to the given string and assigns it to the AltNetworkAdapterType field.
func (o *ComputerHardware) SetAltNetworkAdapterType(v string) {
	o.AltNetworkAdapterType = &v
}

// GetAltMacAddress returns the AltMacAddress field value if set, zero value otherwise.
func (o *ComputerHardware) GetAltMacAddress() string {
	if o == nil || IsNil(o.AltMacAddress) {
		var ret string
		return ret
	}
	return *o.AltMacAddress
}

// GetAltMacAddressOk returns a tuple with the AltMacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetAltMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.AltMacAddress) {
		return nil, false
	}
	return o.AltMacAddress, true
}

// HasAltMacAddress returns a boolean if a field has been set.
func (o *ComputerHardware) HasAltMacAddress() bool {
	if o != nil && !IsNil(o.AltMacAddress) {
		return true
	}

	return false
}

// SetAltMacAddress gets a reference to the given string and assigns it to the AltMacAddress field.
func (o *ComputerHardware) SetAltMacAddress(v string) {
	o.AltMacAddress = &v
}

// GetTotalRamMegabytes returns the TotalRamMegabytes field value if set, zero value otherwise.
func (o *ComputerHardware) GetTotalRamMegabytes() int64 {
	if o == nil || IsNil(o.TotalRamMegabytes) {
		var ret int64
		return ret
	}
	return *o.TotalRamMegabytes
}

// GetTotalRamMegabytesOk returns a tuple with the TotalRamMegabytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetTotalRamMegabytesOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalRamMegabytes) {
		return nil, false
	}
	return o.TotalRamMegabytes, true
}

// HasTotalRamMegabytes returns a boolean if a field has been set.
func (o *ComputerHardware) HasTotalRamMegabytes() bool {
	if o != nil && !IsNil(o.TotalRamMegabytes) {
		return true
	}

	return false
}

// SetTotalRamMegabytes gets a reference to the given int64 and assigns it to the TotalRamMegabytes field.
func (o *ComputerHardware) SetTotalRamMegabytes(v int64) {
	o.TotalRamMegabytes = &v
}

// GetOpenRamSlots returns the OpenRamSlots field value if set, zero value otherwise.
func (o *ComputerHardware) GetOpenRamSlots() int32 {
	if o == nil || IsNil(o.OpenRamSlots) {
		var ret int32
		return ret
	}
	return *o.OpenRamSlots
}

// GetOpenRamSlotsOk returns a tuple with the OpenRamSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetOpenRamSlotsOk() (*int32, bool) {
	if o == nil || IsNil(o.OpenRamSlots) {
		return nil, false
	}
	return o.OpenRamSlots, true
}

// HasOpenRamSlots returns a boolean if a field has been set.
func (o *ComputerHardware) HasOpenRamSlots() bool {
	if o != nil && !IsNil(o.OpenRamSlots) {
		return true
	}

	return false
}

// SetOpenRamSlots gets a reference to the given int32 and assigns it to the OpenRamSlots field.
func (o *ComputerHardware) SetOpenRamSlots(v int32) {
	o.OpenRamSlots = &v
}

// GetBatteryCapacityPercent returns the BatteryCapacityPercent field value if set, zero value otherwise.
func (o *ComputerHardware) GetBatteryCapacityPercent() int32 {
	if o == nil || IsNil(o.BatteryCapacityPercent) {
		var ret int32
		return ret
	}
	return *o.BatteryCapacityPercent
}

// GetBatteryCapacityPercentOk returns a tuple with the BatteryCapacityPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetBatteryCapacityPercentOk() (*int32, bool) {
	if o == nil || IsNil(o.BatteryCapacityPercent) {
		return nil, false
	}
	return o.BatteryCapacityPercent, true
}

// HasBatteryCapacityPercent returns a boolean if a field has been set.
func (o *ComputerHardware) HasBatteryCapacityPercent() bool {
	if o != nil && !IsNil(o.BatteryCapacityPercent) {
		return true
	}

	return false
}

// SetBatteryCapacityPercent gets a reference to the given int32 and assigns it to the BatteryCapacityPercent field.
func (o *ComputerHardware) SetBatteryCapacityPercent(v int32) {
	o.BatteryCapacityPercent = &v
}

// GetSmcVersion returns the SmcVersion field value if set, zero value otherwise.
func (o *ComputerHardware) GetSmcVersion() string {
	if o == nil || IsNil(o.SmcVersion) {
		var ret string
		return ret
	}
	return *o.SmcVersion
}

// GetSmcVersionOk returns a tuple with the SmcVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetSmcVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SmcVersion) {
		return nil, false
	}
	return o.SmcVersion, true
}

// HasSmcVersion returns a boolean if a field has been set.
func (o *ComputerHardware) HasSmcVersion() bool {
	if o != nil && !IsNil(o.SmcVersion) {
		return true
	}

	return false
}

// SetSmcVersion gets a reference to the given string and assigns it to the SmcVersion field.
func (o *ComputerHardware) SetSmcVersion(v string) {
	o.SmcVersion = &v
}

// GetNicSpeed returns the NicSpeed field value if set, zero value otherwise.
func (o *ComputerHardware) GetNicSpeed() string {
	if o == nil || IsNil(o.NicSpeed) {
		var ret string
		return ret
	}
	return *o.NicSpeed
}

// GetNicSpeedOk returns a tuple with the NicSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetNicSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.NicSpeed) {
		return nil, false
	}
	return o.NicSpeed, true
}

// HasNicSpeed returns a boolean if a field has been set.
func (o *ComputerHardware) HasNicSpeed() bool {
	if o != nil && !IsNil(o.NicSpeed) {
		return true
	}

	return false
}

// SetNicSpeed gets a reference to the given string and assigns it to the NicSpeed field.
func (o *ComputerHardware) SetNicSpeed(v string) {
	o.NicSpeed = &v
}

// GetOpticalDrive returns the OpticalDrive field value if set, zero value otherwise.
func (o *ComputerHardware) GetOpticalDrive() string {
	if o == nil || IsNil(o.OpticalDrive) {
		var ret string
		return ret
	}
	return *o.OpticalDrive
}

// GetOpticalDriveOk returns a tuple with the OpticalDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetOpticalDriveOk() (*string, bool) {
	if o == nil || IsNil(o.OpticalDrive) {
		return nil, false
	}
	return o.OpticalDrive, true
}

// HasOpticalDrive returns a boolean if a field has been set.
func (o *ComputerHardware) HasOpticalDrive() bool {
	if o != nil && !IsNil(o.OpticalDrive) {
		return true
	}

	return false
}

// SetOpticalDrive gets a reference to the given string and assigns it to the OpticalDrive field.
func (o *ComputerHardware) SetOpticalDrive(v string) {
	o.OpticalDrive = &v
}

// GetBootRom returns the BootRom field value if set, zero value otherwise.
func (o *ComputerHardware) GetBootRom() string {
	if o == nil || IsNil(o.BootRom) {
		var ret string
		return ret
	}
	return *o.BootRom
}

// GetBootRomOk returns a tuple with the BootRom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetBootRomOk() (*string, bool) {
	if o == nil || IsNil(o.BootRom) {
		return nil, false
	}
	return o.BootRom, true
}

// HasBootRom returns a boolean if a field has been set.
func (o *ComputerHardware) HasBootRom() bool {
	if o != nil && !IsNil(o.BootRom) {
		return true
	}

	return false
}

// SetBootRom gets a reference to the given string and assigns it to the BootRom field.
func (o *ComputerHardware) SetBootRom(v string) {
	o.BootRom = &v
}

// GetBleCapable returns the BleCapable field value if set, zero value otherwise.
func (o *ComputerHardware) GetBleCapable() bool {
	if o == nil || IsNil(o.BleCapable) {
		var ret bool
		return ret
	}
	return *o.BleCapable
}

// GetBleCapableOk returns a tuple with the BleCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetBleCapableOk() (*bool, bool) {
	if o == nil || IsNil(o.BleCapable) {
		return nil, false
	}
	return o.BleCapable, true
}

// HasBleCapable returns a boolean if a field has been set.
func (o *ComputerHardware) HasBleCapable() bool {
	if o != nil && !IsNil(o.BleCapable) {
		return true
	}

	return false
}

// SetBleCapable gets a reference to the given bool and assigns it to the BleCapable field.
func (o *ComputerHardware) SetBleCapable(v bool) {
	o.BleCapable = &v
}

// GetSupportsIosAppInstalls returns the SupportsIosAppInstalls field value if set, zero value otherwise.
func (o *ComputerHardware) GetSupportsIosAppInstalls() bool {
	if o == nil || IsNil(o.SupportsIosAppInstalls) {
		var ret bool
		return ret
	}
	return *o.SupportsIosAppInstalls
}

// GetSupportsIosAppInstallsOk returns a tuple with the SupportsIosAppInstalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetSupportsIosAppInstallsOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsIosAppInstalls) {
		return nil, false
	}
	return o.SupportsIosAppInstalls, true
}

// HasSupportsIosAppInstalls returns a boolean if a field has been set.
func (o *ComputerHardware) HasSupportsIosAppInstalls() bool {
	if o != nil && !IsNil(o.SupportsIosAppInstalls) {
		return true
	}

	return false
}

// SetSupportsIosAppInstalls gets a reference to the given bool and assigns it to the SupportsIosAppInstalls field.
func (o *ComputerHardware) SetSupportsIosAppInstalls(v bool) {
	o.SupportsIosAppInstalls = &v
}

// GetAppleSilicon returns the AppleSilicon field value if set, zero value otherwise.
func (o *ComputerHardware) GetAppleSilicon() bool {
	if o == nil || IsNil(o.AppleSilicon) {
		var ret bool
		return ret
	}
	return *o.AppleSilicon
}

// GetAppleSiliconOk returns a tuple with the AppleSilicon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetAppleSiliconOk() (*bool, bool) {
	if o == nil || IsNil(o.AppleSilicon) {
		return nil, false
	}
	return o.AppleSilicon, true
}

// HasAppleSilicon returns a boolean if a field has been set.
func (o *ComputerHardware) HasAppleSilicon() bool {
	if o != nil && !IsNil(o.AppleSilicon) {
		return true
	}

	return false
}

// SetAppleSilicon gets a reference to the given bool and assigns it to the AppleSilicon field.
func (o *ComputerHardware) SetAppleSilicon(v bool) {
	o.AppleSilicon = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *ComputerHardware) GetExtensionAttributes() []ComputerExtensionAttribute {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret []ComputerExtensionAttribute
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerHardware) GetExtensionAttributesOk() ([]ComputerExtensionAttribute, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *ComputerHardware) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given []ComputerExtensionAttribute and assigns it to the ExtensionAttributes field.
func (o *ComputerHardware) SetExtensionAttributes(v []ComputerExtensionAttribute) {
	o.ExtensionAttributes = v
}

func (o ComputerHardware) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComputerHardware) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Make) {
		toSerialize["make"] = o.Make
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.ModelIdentifier) {
		toSerialize["modelIdentifier"] = o.ModelIdentifier
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if !IsNil(o.ProcessorSpeedMhz) {
		toSerialize["processorSpeedMhz"] = o.ProcessorSpeedMhz
	}
	if !IsNil(o.ProcessorCount) {
		toSerialize["processorCount"] = o.ProcessorCount
	}
	if !IsNil(o.CoreCount) {
		toSerialize["coreCount"] = o.CoreCount
	}
	if !IsNil(o.ProcessorType) {
		toSerialize["processorType"] = o.ProcessorType
	}
	if !IsNil(o.ProcessorArchitecture) {
		toSerialize["processorArchitecture"] = o.ProcessorArchitecture
	}
	if !IsNil(o.BusSpeedMhz) {
		toSerialize["busSpeedMhz"] = o.BusSpeedMhz
	}
	if !IsNil(o.CacheSizeKilobytes) {
		toSerialize["cacheSizeKilobytes"] = o.CacheSizeKilobytes
	}
	if !IsNil(o.NetworkAdapterType) {
		toSerialize["networkAdapterType"] = o.NetworkAdapterType
	}
	if !IsNil(o.MacAddress) {
		toSerialize["macAddress"] = o.MacAddress
	}
	if !IsNil(o.AltNetworkAdapterType) {
		toSerialize["altNetworkAdapterType"] = o.AltNetworkAdapterType
	}
	if !IsNil(o.AltMacAddress) {
		toSerialize["altMacAddress"] = o.AltMacAddress
	}
	if !IsNil(o.TotalRamMegabytes) {
		toSerialize["totalRamMegabytes"] = o.TotalRamMegabytes
	}
	if !IsNil(o.OpenRamSlots) {
		toSerialize["openRamSlots"] = o.OpenRamSlots
	}
	if !IsNil(o.BatteryCapacityPercent) {
		toSerialize["batteryCapacityPercent"] = o.BatteryCapacityPercent
	}
	if !IsNil(o.SmcVersion) {
		toSerialize["smcVersion"] = o.SmcVersion
	}
	if !IsNil(o.NicSpeed) {
		toSerialize["nicSpeed"] = o.NicSpeed
	}
	if !IsNil(o.OpticalDrive) {
		toSerialize["opticalDrive"] = o.OpticalDrive
	}
	if !IsNil(o.BootRom) {
		toSerialize["bootRom"] = o.BootRom
	}
	if !IsNil(o.BleCapable) {
		toSerialize["bleCapable"] = o.BleCapable
	}
	if !IsNil(o.SupportsIosAppInstalls) {
		toSerialize["supportsIosAppInstalls"] = o.SupportsIosAppInstalls
	}
	if !IsNil(o.AppleSilicon) {
		toSerialize["appleSilicon"] = o.AppleSilicon
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extensionAttributes"] = o.ExtensionAttributes
	}
	return toSerialize, nil
}

type NullableComputerHardware struct {
	value *ComputerHardware
	isSet bool
}

func (v NullableComputerHardware) Get() *ComputerHardware {
	return v.value
}

func (v *NullableComputerHardware) Set(val *ComputerHardware) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerHardware) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerHardware) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerHardware(val *ComputerHardware) *NullableComputerHardware {
	return &NullableComputerHardware{value: val, isSet: true}
}

func (v NullableComputerHardware) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerHardware) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


