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
	"fmt"
	"time"
)

// checks if the NetworkZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkZone{}

// NetworkZone struct for NetworkZone
type NetworkZone struct {
	// Timestamp when the object was created
	Created *time.Time `json:"created,omitempty"`
	// Unique identifier for the Network Zone
	Id *string `json:"id,omitempty"`
	// Timestamp when the object was last modified
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Unique name for this Network Zone
	Name string `json:"name"`
	// Network Zone status
	Status *string `json:"status,omitempty"`
	// Indicates a system Network Zone: * `true` for system Network Zones * `false` for custom Network Zones  The Okta org provides the following default system Network Zones: * `LegacyIpZone` * `BlockedIpZone` * `DefaultEnhancedDynamicZone` * `DefaultExemptIpZone`  Admins can modify the name of the default system Network Zone and add up to 5000 gateway or proxy IP entries.
	System *bool `json:"system,omitempty"`
	// The type of Network Zone
	Type string `json:"type"`
	// The usage of the Network Zone
	Usage                *string                `json:"usage,omitempty"`
	Links                *LinksSelfAndLifecycle `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkZone NetworkZone

// NewNetworkZone instantiates a new NetworkZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkZone(name string, type_ string) *NetworkZone {
	this := NetworkZone{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewNetworkZoneWithDefaults instantiates a new NetworkZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkZoneWithDefaults() *NetworkZone {
	this := NetworkZone{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *NetworkZone) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkZone) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *NetworkZone) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *NetworkZone) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkZone) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkZone) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkZone) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworkZone) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *NetworkZone) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkZone) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *NetworkZone) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *NetworkZone) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value
func (o *NetworkZone) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkZone) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkZone) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NetworkZone) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkZone) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NetworkZone) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NetworkZone) SetStatus(v string) {
	o.Status = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *NetworkZone) GetSystem() bool {
	if o == nil || IsNil(o.System) {
		var ret bool
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkZone) GetSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *NetworkZone) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given bool and assigns it to the System field.
func (o *NetworkZone) SetSystem(v bool) {
	o.System = &v
}

// GetType returns the Type field value
func (o *NetworkZone) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NetworkZone) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NetworkZone) SetType(v string) {
	o.Type = v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *NetworkZone) GetUsage() string {
	if o == nil || IsNil(o.Usage) {
		var ret string
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkZone) GetUsageOk() (*string, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *NetworkZone) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given string and assigns it to the Usage field.
func (o *NetworkZone) SetUsage(v string) {
	o.Usage = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *NetworkZone) GetLinks() LinksSelfAndLifecycle {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelfAndLifecycle
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkZone) GetLinksOk() (*LinksSelfAndLifecycle, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *NetworkZone) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelfAndLifecycle and assigns it to the Links field.
func (o *NetworkZone) SetLinks(v LinksSelfAndLifecycle) {
	o.Links = &v
}

func (o NetworkZone) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkZone) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNetworkZone := _NetworkZone{}

	err = json.Unmarshal(data, &varNetworkZone)

	if err != nil {
		return err
	}

	*o = NetworkZone(varNetworkZone)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "system")
		delete(additionalProperties, "type")
		delete(additionalProperties, "usage")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkZone struct {
	value *NetworkZone
	isSet bool
}

func (v NullableNetworkZone) Get() *NetworkZone {
	return v.value
}

func (v *NullableNetworkZone) Set(val *NetworkZone) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkZone) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkZone(val *NetworkZone) *NullableNetworkZone {
	return &NullableNetworkZone{value: val, isSet: true}
}

func (v NullableNetworkZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
