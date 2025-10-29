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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the DevicePostureCheck type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicePostureCheck{}

// DevicePostureCheck struct for DevicePostureCheck
type DevicePostureCheck struct {
	// User who created the device posture check
	CreatedBy *string `json:"createdBy,omitempty"`
	// Time the device posture check was created
	CreatedDate *string `json:"createdDate,omitempty"`
	// Description of the device posture check
	Description *string `json:"description,omitempty"`
	// The ID of the device posture check
	Id *string `json:"id,omitempty"`
	// Time the device posture check was updated
	LastUpdate *string `json:"lastUpdate,omitempty"`
	// User who updated the device posture check
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	// Represents how the device posture check is rendered in device assurance policies
	MappingType *string `json:"mappingType,omitempty"`
	// Display name of the device posture check
	Name     *string `json:"name,omitempty"`
	Platform *string `json:"platform,omitempty"`
	// OSQuery for the device posture check
	Query               *string                                 `json:"query,omitempty"`
	RemediationSettings *DevicePostureChecksRemediationSettings `json:"remediationSettings,omitempty"`
	Type                *string                                 `json:"type,omitempty"`
	// Unique name of the device posture check
	VariableName         *string    `json:"variableName,omitempty"`
	Links                *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DevicePostureCheck DevicePostureCheck

// NewDevicePostureCheck instantiates a new DevicePostureCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicePostureCheck() *DevicePostureCheck {
	this := DevicePostureCheck{}
	return &this
}

// NewDevicePostureCheckWithDefaults instantiates a new DevicePostureCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicePostureCheckWithDefaults() *DevicePostureCheck {
	this := DevicePostureCheck{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DevicePostureCheck) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureCheck) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DevicePostureCheck) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *DevicePostureCheck) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *DevicePostureCheck) GetCreatedDate() string {
	if o == nil || IsNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureCheck) GetCreatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *DevicePostureCheck) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *DevicePostureCheck) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DevicePostureCheck) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureCheck) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DevicePostureCheck) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DevicePostureCheck) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DevicePostureCheck) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureCheck) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DevicePostureCheck) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DevicePostureCheck) SetId(v string) {
	o.Id = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *DevicePostureCheck) GetLastUpdate() string {
	if o == nil || IsNil(o.LastUpdate) {
		var ret string
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureCheck) GetLastUpdateOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdate) {
		return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *DevicePostureCheck) HasLastUpdate() bool {
	if o != nil && !IsNil(o.LastUpdate) {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given string and assigns it to the LastUpdate field.
func (o *DevicePostureCheck) SetLastUpdate(v string) {
	o.LastUpdate = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *DevicePostureCheck) GetLastUpdatedBy() string {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureCheck) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *DevicePostureCheck) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *DevicePostureCheck) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetMappingType returns the MappingType field value if set, zero value otherwise.
func (o *DevicePostureCheck) GetMappingType() string {
	if o == nil || IsNil(o.MappingType) {
		var ret string
		return ret
	}
	return *o.MappingType
}

// GetMappingTypeOk returns a tuple with the MappingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureCheck) GetMappingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MappingType) {
		return nil, false
	}
	return o.MappingType, true
}

// HasMappingType returns a boolean if a field has been set.
func (o *DevicePostureCheck) HasMappingType() bool {
	if o != nil && !IsNil(o.MappingType) {
		return true
	}

	return false
}

// SetMappingType gets a reference to the given string and assigns it to the MappingType field.
func (o *DevicePostureCheck) SetMappingType(v string) {
	o.MappingType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DevicePostureCheck) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureCheck) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DevicePostureCheck) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DevicePostureCheck) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *DevicePostureCheck) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureCheck) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *DevicePostureCheck) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *DevicePostureCheck) SetPlatform(v string) {
	o.Platform = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *DevicePostureCheck) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureCheck) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *DevicePostureCheck) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *DevicePostureCheck) SetQuery(v string) {
	o.Query = &v
}

// GetRemediationSettings returns the RemediationSettings field value if set, zero value otherwise.
func (o *DevicePostureCheck) GetRemediationSettings() DevicePostureChecksRemediationSettings {
	if o == nil || IsNil(o.RemediationSettings) {
		var ret DevicePostureChecksRemediationSettings
		return ret
	}
	return *o.RemediationSettings
}

// GetRemediationSettingsOk returns a tuple with the RemediationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureCheck) GetRemediationSettingsOk() (*DevicePostureChecksRemediationSettings, bool) {
	if o == nil || IsNil(o.RemediationSettings) {
		return nil, false
	}
	return o.RemediationSettings, true
}

// HasRemediationSettings returns a boolean if a field has been set.
func (o *DevicePostureCheck) HasRemediationSettings() bool {
	if o != nil && !IsNil(o.RemediationSettings) {
		return true
	}

	return false
}

// SetRemediationSettings gets a reference to the given DevicePostureChecksRemediationSettings and assigns it to the RemediationSettings field.
func (o *DevicePostureCheck) SetRemediationSettings(v DevicePostureChecksRemediationSettings) {
	o.RemediationSettings = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DevicePostureCheck) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureCheck) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DevicePostureCheck) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DevicePostureCheck) SetType(v string) {
	o.Type = &v
}

// GetVariableName returns the VariableName field value if set, zero value otherwise.
func (o *DevicePostureCheck) GetVariableName() string {
	if o == nil || IsNil(o.VariableName) {
		var ret string
		return ret
	}
	return *o.VariableName
}

// GetVariableNameOk returns a tuple with the VariableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureCheck) GetVariableNameOk() (*string, bool) {
	if o == nil || IsNil(o.VariableName) {
		return nil, false
	}
	return o.VariableName, true
}

// HasVariableName returns a boolean if a field has been set.
func (o *DevicePostureCheck) HasVariableName() bool {
	if o != nil && !IsNil(o.VariableName) {
		return true
	}

	return false
}

// SetVariableName gets a reference to the given string and assigns it to the VariableName field.
func (o *DevicePostureCheck) SetVariableName(v string) {
	o.VariableName = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DevicePostureCheck) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureCheck) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DevicePostureCheck) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *DevicePostureCheck) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o DevicePostureCheck) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicePostureCheck) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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
	if !IsNil(o.MappingType) {
		toSerialize["mappingType"] = o.MappingType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.RemediationSettings) {
		toSerialize["remediationSettings"] = o.RemediationSettings
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.VariableName) {
		toSerialize["variableName"] = o.VariableName
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DevicePostureCheck) UnmarshalJSON(data []byte) (err error) {
	varDevicePostureCheck := _DevicePostureCheck{}

	err = json.Unmarshal(data, &varDevicePostureCheck)

	if err != nil {
		return err
	}

	*o = DevicePostureCheck(varDevicePostureCheck)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdate")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "mappingType")
		delete(additionalProperties, "name")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "query")
		delete(additionalProperties, "remediationSettings")
		delete(additionalProperties, "type")
		delete(additionalProperties, "variableName")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDevicePostureCheck struct {
	value *DevicePostureCheck
	isSet bool
}

func (v NullableDevicePostureCheck) Get() *DevicePostureCheck {
	return v.value
}

func (v *NullableDevicePostureCheck) Set(val *DevicePostureCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicePostureCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicePostureCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicePostureCheck(val *DevicePostureCheck) *NullableDevicePostureCheck {
	return &NullableDevicePostureCheck{value: val, isSet: true}
}

func (v NullableDevicePostureCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicePostureCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
