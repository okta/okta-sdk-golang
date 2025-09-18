/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the DeviceAssurance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAssurance{}

// DeviceAssurance struct for DeviceAssurance
type DeviceAssurance struct {
	CreatedBy           *string              `json:"createdBy,omitempty"`
	CreatedDate         *string              `json:"createdDate,omitempty"`
	DevicePostureChecks *DevicePostureChecks `json:"devicePostureChecks,omitempty"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>Represents the remediation mode of this device assurance policy when users are denied access due to device noncompliance
	DisplayRemediationMode *string      `json:"displayRemediationMode,omitempty"`
	GracePeriod            *GracePeriod `json:"gracePeriod,omitempty"`
	Id                     *string      `json:"id,omitempty"`
	LastUpdate             *string      `json:"lastUpdate,omitempty"`
	LastUpdatedBy          *string      `json:"lastUpdatedBy,omitempty"`
	// Display name of the device assurance policy
	Name                 *string    `json:"name,omitempty"`
	Platform             *string    `json:"platform,omitempty"`
	Links                *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAssurance DeviceAssurance

// NewDeviceAssurance instantiates a new DeviceAssurance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAssurance() *DeviceAssurance {
	this := DeviceAssurance{}
	return &this
}

// NewDeviceAssuranceWithDefaults instantiates a new DeviceAssurance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAssuranceWithDefaults() *DeviceAssurance {
	this := DeviceAssurance{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DeviceAssurance) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssurance) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DeviceAssurance) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *DeviceAssurance) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *DeviceAssurance) GetCreatedDate() string {
	if o == nil || IsNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssurance) GetCreatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *DeviceAssurance) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *DeviceAssurance) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetDevicePostureChecks returns the DevicePostureChecks field value if set, zero value otherwise.
func (o *DeviceAssurance) GetDevicePostureChecks() DevicePostureChecks {
	if o == nil || IsNil(o.DevicePostureChecks) {
		var ret DevicePostureChecks
		return ret
	}
	return *o.DevicePostureChecks
}

// GetDevicePostureChecksOk returns a tuple with the DevicePostureChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssurance) GetDevicePostureChecksOk() (*DevicePostureChecks, bool) {
	if o == nil || IsNil(o.DevicePostureChecks) {
		return nil, false
	}
	return o.DevicePostureChecks, true
}

// HasDevicePostureChecks returns a boolean if a field has been set.
func (o *DeviceAssurance) HasDevicePostureChecks() bool {
	if o != nil && !IsNil(o.DevicePostureChecks) {
		return true
	}

	return false
}

// SetDevicePostureChecks gets a reference to the given DevicePostureChecks and assigns it to the DevicePostureChecks field.
func (o *DeviceAssurance) SetDevicePostureChecks(v DevicePostureChecks) {
	o.DevicePostureChecks = &v
}

// GetDisplayRemediationMode returns the DisplayRemediationMode field value if set, zero value otherwise.
func (o *DeviceAssurance) GetDisplayRemediationMode() string {
	if o == nil || IsNil(o.DisplayRemediationMode) {
		var ret string
		return ret
	}
	return *o.DisplayRemediationMode
}

// GetDisplayRemediationModeOk returns a tuple with the DisplayRemediationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssurance) GetDisplayRemediationModeOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayRemediationMode) {
		return nil, false
	}
	return o.DisplayRemediationMode, true
}

// HasDisplayRemediationMode returns a boolean if a field has been set.
func (o *DeviceAssurance) HasDisplayRemediationMode() bool {
	if o != nil && !IsNil(o.DisplayRemediationMode) {
		return true
	}

	return false
}

// SetDisplayRemediationMode gets a reference to the given string and assigns it to the DisplayRemediationMode field.
func (o *DeviceAssurance) SetDisplayRemediationMode(v string) {
	o.DisplayRemediationMode = &v
}

// GetGracePeriod returns the GracePeriod field value if set, zero value otherwise.
func (o *DeviceAssurance) GetGracePeriod() GracePeriod {
	if o == nil || IsNil(o.GracePeriod) {
		var ret GracePeriod
		return ret
	}
	return *o.GracePeriod
}

// GetGracePeriodOk returns a tuple with the GracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssurance) GetGracePeriodOk() (*GracePeriod, bool) {
	if o == nil || IsNil(o.GracePeriod) {
		return nil, false
	}
	return o.GracePeriod, true
}

// HasGracePeriod returns a boolean if a field has been set.
func (o *DeviceAssurance) HasGracePeriod() bool {
	if o != nil && !IsNil(o.GracePeriod) {
		return true
	}

	return false
}

// SetGracePeriod gets a reference to the given GracePeriod and assigns it to the GracePeriod field.
func (o *DeviceAssurance) SetGracePeriod(v GracePeriod) {
	o.GracePeriod = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceAssurance) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssurance) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceAssurance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceAssurance) SetId(v string) {
	o.Id = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *DeviceAssurance) GetLastUpdate() string {
	if o == nil || IsNil(o.LastUpdate) {
		var ret string
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssurance) GetLastUpdateOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdate) {
		return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *DeviceAssurance) HasLastUpdate() bool {
	if o != nil && !IsNil(o.LastUpdate) {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given string and assigns it to the LastUpdate field.
func (o *DeviceAssurance) SetLastUpdate(v string) {
	o.LastUpdate = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *DeviceAssurance) GetLastUpdatedBy() string {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssurance) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *DeviceAssurance) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *DeviceAssurance) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeviceAssurance) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssurance) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeviceAssurance) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeviceAssurance) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *DeviceAssurance) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssurance) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *DeviceAssurance) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *DeviceAssurance) SetPlatform(v string) {
	o.Platform = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DeviceAssurance) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssurance) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DeviceAssurance) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *DeviceAssurance) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o DeviceAssurance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAssurance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.DevicePostureChecks) {
		toSerialize["devicePostureChecks"] = o.DevicePostureChecks
	}
	if !IsNil(o.DisplayRemediationMode) {
		toSerialize["displayRemediationMode"] = o.DisplayRemediationMode
	}
	if !IsNil(o.GracePeriod) {
		toSerialize["gracePeriod"] = o.GracePeriod
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdate) {
		toSerialize["lastUpdate"] = o.LastUpdate
	}
	if !IsNil(o.LastUpdatedBy) {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceAssurance) UnmarshalJSON(data []byte) (err error) {
	varDeviceAssurance := _DeviceAssurance{}

	err = json.Unmarshal(data, &varDeviceAssurance)

	if err != nil {
		return err
	}

	*o = DeviceAssurance(varDeviceAssurance)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "devicePostureChecks")
		delete(additionalProperties, "displayRemediationMode")
		delete(additionalProperties, "gracePeriod")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdate")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "name")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceAssurance struct {
	value *DeviceAssurance
	isSet bool
}

func (v NullableDeviceAssurance) Get() *DeviceAssurance {
	return v.value
}

func (v *NullableDeviceAssurance) Set(val *DeviceAssurance) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAssurance) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAssurance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAssurance(val *DeviceAssurance) *NullableDeviceAssurance {
	return &NullableDeviceAssurance{value: val, isSet: true}
}

func (v NullableDeviceAssurance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAssurance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
