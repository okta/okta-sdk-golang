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

// EventHookFilterMapObject struct for EventHookFilterMapObject
type EventHookFilterMapObject struct {
	Condition *EventHookFilterMapObjectCondition `json:"condition,omitempty"`
	// The filtered event type
	Event *string `json:"event,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventHookFilterMapObject EventHookFilterMapObject

// NewEventHookFilterMapObject instantiates a new EventHookFilterMapObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventHookFilterMapObject() *EventHookFilterMapObject {
	this := EventHookFilterMapObject{}
	return &this
}

// NewEventHookFilterMapObjectWithDefaults instantiates a new EventHookFilterMapObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventHookFilterMapObjectWithDefaults() *EventHookFilterMapObject {
	this := EventHookFilterMapObject{}
	return &this
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *EventHookFilterMapObject) GetCondition() EventHookFilterMapObjectCondition {
	if o == nil || o.Condition == nil {
		var ret EventHookFilterMapObjectCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHookFilterMapObject) GetConditionOk() (*EventHookFilterMapObjectCondition, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *EventHookFilterMapObject) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given EventHookFilterMapObjectCondition and assigns it to the Condition field.
func (o *EventHookFilterMapObject) SetCondition(v EventHookFilterMapObjectCondition) {
	o.Condition = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *EventHookFilterMapObject) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHookFilterMapObject) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *EventHookFilterMapObject) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *EventHookFilterMapObject) SetEvent(v string) {
	o.Event = &v
}

func (o EventHookFilterMapObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EventHookFilterMapObject) UnmarshalJSON(bytes []byte) (err error) {
	varEventHookFilterMapObject := _EventHookFilterMapObject{}

	err = json.Unmarshal(bytes, &varEventHookFilterMapObject)
	if err == nil {
		*o = EventHookFilterMapObject(varEventHookFilterMapObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "condition")
		delete(additionalProperties, "event")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEventHookFilterMapObject struct {
	value *EventHookFilterMapObject
	isSet bool
}

func (v NullableEventHookFilterMapObject) Get() *EventHookFilterMapObject {
	return v.value
}

func (v *NullableEventHookFilterMapObject) Set(val *EventHookFilterMapObject) {
	v.value = val
	v.isSet = true
}

func (v NullableEventHookFilterMapObject) IsSet() bool {
	return v.isSet
}

func (v *NullableEventHookFilterMapObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventHookFilterMapObject(val *EventHookFilterMapObject) *NullableEventHookFilterMapObject {
	return &NullableEventHookFilterMapObject{value: val, isSet: true}
}

func (v NullableEventHookFilterMapObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventHookFilterMapObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

