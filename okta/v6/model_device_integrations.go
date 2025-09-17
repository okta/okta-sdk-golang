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

// checks if the DeviceIntegrations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceIntegrations{}

// DeviceIntegrations struct for DeviceIntegrations
type DeviceIntegrations struct {
	// The display name of the device integration
	DisplayName *string `json:"displayName,omitempty"`
	// The ID of the device integration
	Id       *string                     `json:"id,omitempty"`
	Metadata *DeviceIntegrationsMetadata `json:"metadata,omitempty"`
	// The namespace of the device integration
	Name     *string `json:"name,omitempty"`
	Platform *string `json:"platform,omitempty"`
	// The status of the device integration
	Status               *string                `json:"status,omitempty"`
	Links                *LinksSelfAndLifecycle `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceIntegrations DeviceIntegrations

// NewDeviceIntegrations instantiates a new DeviceIntegrations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceIntegrations() *DeviceIntegrations {
	this := DeviceIntegrations{}
	return &this
}

// NewDeviceIntegrationsWithDefaults instantiates a new DeviceIntegrations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceIntegrationsWithDefaults() *DeviceIntegrations {
	this := DeviceIntegrations{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *DeviceIntegrations) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceIntegrations) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DeviceIntegrations) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *DeviceIntegrations) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceIntegrations) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceIntegrations) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceIntegrations) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceIntegrations) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DeviceIntegrations) GetMetadata() DeviceIntegrationsMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret DeviceIntegrationsMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceIntegrations) GetMetadataOk() (*DeviceIntegrationsMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DeviceIntegrations) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given DeviceIntegrationsMetadata and assigns it to the Metadata field.
func (o *DeviceIntegrations) SetMetadata(v DeviceIntegrationsMetadata) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeviceIntegrations) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceIntegrations) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeviceIntegrations) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeviceIntegrations) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *DeviceIntegrations) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceIntegrations) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *DeviceIntegrations) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *DeviceIntegrations) SetPlatform(v string) {
	o.Platform = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeviceIntegrations) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceIntegrations) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeviceIntegrations) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeviceIntegrations) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DeviceIntegrations) GetLinks() LinksSelfAndLifecycle {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelfAndLifecycle
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceIntegrations) GetLinksOk() (*LinksSelfAndLifecycle, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DeviceIntegrations) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelfAndLifecycle and assigns it to the Links field.
func (o *DeviceIntegrations) SetLinks(v LinksSelfAndLifecycle) {
	o.Links = &v
}

func (o DeviceIntegrations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceIntegrations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceIntegrations) UnmarshalJSON(data []byte) (err error) {
	varDeviceIntegrations := _DeviceIntegrations{}

	err = json.Unmarshal(data, &varDeviceIntegrations)

	if err != nil {
		return err
	}

	*o = DeviceIntegrations(varDeviceIntegrations)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "id")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "name")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceIntegrations struct {
	value *DeviceIntegrations
	isSet bool
}

func (v NullableDeviceIntegrations) Get() *DeviceIntegrations {
	return v.value
}

func (v *NullableDeviceIntegrations) Set(val *DeviceIntegrations) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceIntegrations) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceIntegrations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceIntegrations(val *DeviceIntegrations) *NullableDeviceIntegrations {
	return &NullableDeviceIntegrations{value: val, isSet: true}
}

func (v NullableDeviceIntegrations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceIntegrations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
