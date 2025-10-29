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
	"fmt"
)

// checks if the LogStreamLinksSelfAndLifecycle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogStreamLinksSelfAndLifecycle{}

// LogStreamLinksSelfAndLifecycle Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the current status of an application using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification. This object is used for dynamic discovery of related resources and lifecycle operations.
type LogStreamLinksSelfAndLifecycle struct {
	Activate             *LogStreamActivateLink   `json:"activate,omitempty"`
	Deactivate           *LogStreamDeactivateLink `json:"deactivate,omitempty"`
	Self                 LogStreamSelfLink        `json:"self"`
	AdditionalProperties map[string]interface{}
}

type _LogStreamLinksSelfAndLifecycle LogStreamLinksSelfAndLifecycle

// NewLogStreamLinksSelfAndLifecycle instantiates a new LogStreamLinksSelfAndLifecycle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStreamLinksSelfAndLifecycle(self LogStreamSelfLink) *LogStreamLinksSelfAndLifecycle {
	this := LogStreamLinksSelfAndLifecycle{}
	this.Self = self
	return &this
}

// NewLogStreamLinksSelfAndLifecycleWithDefaults instantiates a new LogStreamLinksSelfAndLifecycle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamLinksSelfAndLifecycleWithDefaults() *LogStreamLinksSelfAndLifecycle {
	this := LogStreamLinksSelfAndLifecycle{}
	return &this
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *LogStreamLinksSelfAndLifecycle) GetActivate() LogStreamActivateLink {
	if o == nil || IsNil(o.Activate) {
		var ret LogStreamActivateLink
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamLinksSelfAndLifecycle) GetActivateOk() (*LogStreamActivateLink, bool) {
	if o == nil || IsNil(o.Activate) {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *LogStreamLinksSelfAndLifecycle) HasActivate() bool {
	if o != nil && !IsNil(o.Activate) {
		return true
	}

	return false
}

// SetActivate gets a reference to the given LogStreamActivateLink and assigns it to the Activate field.
func (o *LogStreamLinksSelfAndLifecycle) SetActivate(v LogStreamActivateLink) {
	o.Activate = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *LogStreamLinksSelfAndLifecycle) GetDeactivate() LogStreamDeactivateLink {
	if o == nil || IsNil(o.Deactivate) {
		var ret LogStreamDeactivateLink
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamLinksSelfAndLifecycle) GetDeactivateOk() (*LogStreamDeactivateLink, bool) {
	if o == nil || IsNil(o.Deactivate) {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *LogStreamLinksSelfAndLifecycle) HasDeactivate() bool {
	if o != nil && !IsNil(o.Deactivate) {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given LogStreamDeactivateLink and assigns it to the Deactivate field.
func (o *LogStreamLinksSelfAndLifecycle) SetDeactivate(v LogStreamDeactivateLink) {
	o.Deactivate = &v
}

// GetSelf returns the Self field value
func (o *LogStreamLinksSelfAndLifecycle) GetSelf() LogStreamSelfLink {
	if o == nil {
		var ret LogStreamSelfLink
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *LogStreamLinksSelfAndLifecycle) GetSelfOk() (*LogStreamSelfLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *LogStreamLinksSelfAndLifecycle) SetSelf(v LogStreamSelfLink) {
	o.Self = v
}

func (o LogStreamLinksSelfAndLifecycle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogStreamLinksSelfAndLifecycle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Activate) {
		toSerialize["activate"] = o.Activate
	}
	if !IsNil(o.Deactivate) {
		toSerialize["deactivate"] = o.Deactivate
	}
	toSerialize["self"] = o.Self

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogStreamLinksSelfAndLifecycle) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
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

	varLogStreamLinksSelfAndLifecycle := _LogStreamLinksSelfAndLifecycle{}

	err = json.Unmarshal(data, &varLogStreamLinksSelfAndLifecycle)

	if err != nil {
		return err
	}

	*o = LogStreamLinksSelfAndLifecycle(varLogStreamLinksSelfAndLifecycle)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activate")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogStreamLinksSelfAndLifecycle struct {
	value *LogStreamLinksSelfAndLifecycle
	isSet bool
}

func (v NullableLogStreamLinksSelfAndLifecycle) Get() *LogStreamLinksSelfAndLifecycle {
	return v.value
}

func (v *NullableLogStreamLinksSelfAndLifecycle) Set(val *LogStreamLinksSelfAndLifecycle) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStreamLinksSelfAndLifecycle) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStreamLinksSelfAndLifecycle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStreamLinksSelfAndLifecycle(val *LogStreamLinksSelfAndLifecycle) *NullableLogStreamLinksSelfAndLifecycle {
	return &NullableLogStreamLinksSelfAndLifecycle{value: val, isSet: true}
}

func (v NullableLogStreamLinksSelfAndLifecycle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStreamLinksSelfAndLifecycle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
