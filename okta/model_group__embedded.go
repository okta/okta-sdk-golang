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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the GroupEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupEmbedded{}

// GroupEmbedded Embedded resources related to the group
type GroupEmbedded struct {
	Stats                *GroupEmbeddedStats `json:"stats,omitempty"`
	App                  *GroupEmbeddedApp   `json:"app,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupEmbedded GroupEmbedded

// NewGroupEmbedded instantiates a new GroupEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupEmbedded() *GroupEmbedded {
	this := GroupEmbedded{}
	return &this
}

// NewGroupEmbeddedWithDefaults instantiates a new GroupEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupEmbeddedWithDefaults() *GroupEmbedded {
	this := GroupEmbedded{}
	return &this
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *GroupEmbedded) GetStats() GroupEmbeddedStats {
	if o == nil || IsNil(o.Stats) {
		var ret GroupEmbeddedStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupEmbedded) GetStatsOk() (*GroupEmbeddedStats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *GroupEmbedded) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given GroupEmbeddedStats and assigns it to the Stats field.
func (o *GroupEmbedded) SetStats(v GroupEmbeddedStats) {
	o.Stats = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *GroupEmbedded) GetApp() GroupEmbeddedApp {
	if o == nil || IsNil(o.App) {
		var ret GroupEmbeddedApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupEmbedded) GetAppOk() (*GroupEmbeddedApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *GroupEmbedded) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given GroupEmbeddedApp and assigns it to the App field.
func (o *GroupEmbedded) SetApp(v GroupEmbeddedApp) {
	o.App = &v
}

func (o GroupEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupEmbedded) UnmarshalJSON(data []byte) (err error) {
	varGroupEmbedded := _GroupEmbedded{}

	err = json.Unmarshal(data, &varGroupEmbedded)

	if err != nil {
		return err
	}

	*o = GroupEmbedded(varGroupEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "stats")
		delete(additionalProperties, "app")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupEmbedded struct {
	value *GroupEmbedded
	isSet bool
}

func (v NullableGroupEmbedded) Get() *GroupEmbedded {
	return v.value
}

func (v *NullableGroupEmbedded) Set(val *GroupEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupEmbedded(val *GroupEmbedded) *NullableGroupEmbedded {
	return &NullableGroupEmbedded{value: val, isSet: true}
}

func (v NullableGroupEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
