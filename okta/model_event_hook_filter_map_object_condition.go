/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// EventHookFilterMapObjectCondition struct for EventHookFilterMapObjectCondition
type EventHookFilterMapObjectCondition struct {
	// The Okta Expression language statement that filters the event type
	Expression *string `json:"expression,omitempty"`
	// Internal field
	Version NullableString `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventHookFilterMapObjectCondition EventHookFilterMapObjectCondition

// NewEventHookFilterMapObjectCondition instantiates a new EventHookFilterMapObjectCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventHookFilterMapObjectCondition() *EventHookFilterMapObjectCondition {
	this := EventHookFilterMapObjectCondition{}
	return &this
}

// NewEventHookFilterMapObjectConditionWithDefaults instantiates a new EventHookFilterMapObjectCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventHookFilterMapObjectConditionWithDefaults() *EventHookFilterMapObjectCondition {
	this := EventHookFilterMapObjectCondition{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *EventHookFilterMapObjectCondition) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHookFilterMapObjectCondition) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *EventHookFilterMapObjectCondition) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *EventHookFilterMapObjectCondition) SetExpression(v string) {
	o.Expression = &v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventHookFilterMapObjectCondition) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventHookFilterMapObjectCondition) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *EventHookFilterMapObjectCondition) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *EventHookFilterMapObjectCondition) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *EventHookFilterMapObjectCondition) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *EventHookFilterMapObjectCondition) UnsetVersion() {
	o.Version.Unset()
}

func (o EventHookFilterMapObjectCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EventHookFilterMapObjectCondition) UnmarshalJSON(bytes []byte) (err error) {
	varEventHookFilterMapObjectCondition := _EventHookFilterMapObjectCondition{}

	err = json.Unmarshal(bytes, &varEventHookFilterMapObjectCondition)
	if err == nil {
		*o = EventHookFilterMapObjectCondition(varEventHookFilterMapObjectCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "expression")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEventHookFilterMapObjectCondition struct {
	value *EventHookFilterMapObjectCondition
	isSet bool
}

func (v NullableEventHookFilterMapObjectCondition) Get() *EventHookFilterMapObjectCondition {
	return v.value
}

func (v *NullableEventHookFilterMapObjectCondition) Set(val *EventHookFilterMapObjectCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableEventHookFilterMapObjectCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableEventHookFilterMapObjectCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventHookFilterMapObjectCondition(val *EventHookFilterMapObjectCondition) *NullableEventHookFilterMapObjectCondition {
	return &NullableEventHookFilterMapObjectCondition{value: val, isSet: true}
}

func (v NullableEventHookFilterMapObjectCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventHookFilterMapObjectCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

