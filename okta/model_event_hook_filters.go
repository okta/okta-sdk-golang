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

// checks if the EventHookFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventHookFilters{}

// EventHookFilters The optional filter defined on a specific event type  > **Note:** Event hook filters is a [self-service Early Access (EA)](/openapi/okta-management/guides/release-lifecycle/#early-access-ea) to enable. If you want to disable this feature, it's recommended to first remove all event filters.
type EventHookFilters struct {
	// The object that maps the filter to the event type
	EventFilterMap []EventHookFilterMapObject `json:"eventFilterMap,omitempty"`
	// The type of filter. Currently only supports `EXPRESSION_LANGUAGE`
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventHookFilters EventHookFilters

// NewEventHookFilters instantiates a new EventHookFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventHookFilters() *EventHookFilters {
	this := EventHookFilters{}
	return &this
}

// NewEventHookFiltersWithDefaults instantiates a new EventHookFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventHookFiltersWithDefaults() *EventHookFilters {
	this := EventHookFilters{}
	return &this
}

// GetEventFilterMap returns the EventFilterMap field value if set, zero value otherwise.
func (o *EventHookFilters) GetEventFilterMap() []EventHookFilterMapObject {
	if o == nil || IsNil(o.EventFilterMap) {
		var ret []EventHookFilterMapObject
		return ret
	}
	return o.EventFilterMap
}

// GetEventFilterMapOk returns a tuple with the EventFilterMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHookFilters) GetEventFilterMapOk() ([]EventHookFilterMapObject, bool) {
	if o == nil || IsNil(o.EventFilterMap) {
		return nil, false
	}
	return o.EventFilterMap, true
}

// HasEventFilterMap returns a boolean if a field has been set.
func (o *EventHookFilters) HasEventFilterMap() bool {
	if o != nil && !IsNil(o.EventFilterMap) {
		return true
	}

	return false
}

// SetEventFilterMap gets a reference to the given []EventHookFilterMapObject and assigns it to the EventFilterMap field.
func (o *EventHookFilters) SetEventFilterMap(v []EventHookFilterMapObject) {
	o.EventFilterMap = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EventHookFilters) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHookFilters) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EventHookFilters) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EventHookFilters) SetType(v string) {
	o.Type = &v
}

func (o EventHookFilters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventHookFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventFilterMap) {
		toSerialize["eventFilterMap"] = o.EventFilterMap
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventHookFilters) UnmarshalJSON(data []byte) (err error) {
	varEventHookFilters := _EventHookFilters{}

	err = json.Unmarshal(data, &varEventHookFilters)

	if err != nil {
		return err
	}

	*o = EventHookFilters(varEventHookFilters)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "eventFilterMap")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventHookFilters struct {
	value *EventHookFilters
	isSet bool
}

func (v NullableEventHookFilters) Get() *EventHookFilters {
	return v.value
}

func (v *NullableEventHookFilters) Set(val *EventHookFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableEventHookFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableEventHookFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventHookFilters(val *EventHookFilters) *NullableEventHookFilters {
	return &NullableEventHookFilters{value: val, isSet: true}
}

func (v NullableEventHookFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventHookFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
