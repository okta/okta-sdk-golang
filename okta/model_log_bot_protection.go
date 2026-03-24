/*
Okta Admin Management API

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

// checks if the LogBotProtection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogBotProtection{}

// LogBotProtection <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>The result of the bot protection detection associated with the event
type LogBotProtection struct {
	// The bot detected level associated with the bot protection configuration target
	Level                NullableString `json:"level,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogBotProtection LogBotProtection

// NewLogBotProtection instantiates a new LogBotProtection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogBotProtection() *LogBotProtection {
	this := LogBotProtection{}
	return &this
}

// NewLogBotProtectionWithDefaults instantiates a new LogBotProtection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogBotProtectionWithDefaults() *LogBotProtection {
	this := LogBotProtection{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogBotProtection) GetLevel() string {
	if o == nil || IsNil(o.Level.Get()) {
		var ret string
		return ret
	}
	return *o.Level.Get()
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogBotProtection) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Level.Get(), o.Level.IsSet()
}

// HasLevel returns a boolean if a field has been set.
func (o *LogBotProtection) HasLevel() bool {
	if o != nil && o.Level.IsSet() {
		return true
	}

	return false
}

// SetLevel gets a reference to the given NullableString and assigns it to the Level field.
func (o *LogBotProtection) SetLevel(v string) {
	o.Level.Set(&v)
}

// SetLevelNil sets the value for Level to be an explicit nil
func (o *LogBotProtection) SetLevelNil() {
	o.Level.Set(nil)
}

// UnsetLevel ensures that no value is present for Level, not even an explicit nil
func (o *LogBotProtection) UnsetLevel() {
	o.Level.Unset()
}

func (o LogBotProtection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogBotProtection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Level.IsSet() {
		toSerialize["level"] = o.Level.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogBotProtection) UnmarshalJSON(data []byte) (err error) {
	varLogBotProtection := _LogBotProtection{}

	err = json.Unmarshal(data, &varLogBotProtection)

	if err != nil {
		return err
	}

	*o = LogBotProtection(varLogBotProtection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "level")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogBotProtection struct {
	value *LogBotProtection
	isSet bool
}

func (v NullableLogBotProtection) Get() *LogBotProtection {
	return v.value
}

func (v *NullableLogBotProtection) Set(val *LogBotProtection) {
	v.value = val
	v.isSet = true
}

func (v NullableLogBotProtection) IsSet() bool {
	return v.isSet
}

func (v *NullableLogBotProtection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogBotProtection(val *LogBotProtection) *NullableLogBotProtection {
	return &NullableLogBotProtection{value: val, isSet: true}
}

func (v NullableLogBotProtection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogBotProtection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
