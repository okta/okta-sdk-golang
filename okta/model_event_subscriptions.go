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

// EventSubscriptions struct for EventSubscriptions
type EventSubscriptions struct {
	Filter NullableEventHookFilters `json:"filter,omitempty"`
	// The subscribed event types that trigger the event hook. When you register an event hook you need to specify which events you want to subscribe to. To see the list of event types currently eligible for use in event hooks, use the [Event Types catalog](/docs/reference/api/event-types/#catalog) and search with the parameter `event-hook-eligible`.
	Items []string `json:"items"`
	// The events object type. Currently supports `EVENT_TYPE`.
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _EventSubscriptions EventSubscriptions

// NewEventSubscriptions instantiates a new EventSubscriptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSubscriptions(items []string, type_ string) *EventSubscriptions {
	this := EventSubscriptions{}
	this.Items = items
	this.Type = type_
	return &this
}

// NewEventSubscriptionsWithDefaults instantiates a new EventSubscriptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSubscriptionsWithDefaults() *EventSubscriptions {
	this := EventSubscriptions{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSubscriptions) GetFilter() EventHookFilters {
	if o == nil || o.Filter.Get() == nil {
		var ret EventHookFilters
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSubscriptions) GetFilterOk() (*EventHookFilters, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a field has been set.
func (o *EventSubscriptions) HasFilter() bool {
	if o != nil && o.Filter.IsSet() {
		return true
	}

	return false
}

// SetFilter gets a reference to the given NullableEventHookFilters and assigns it to the Filter field.
func (o *EventSubscriptions) SetFilter(v EventHookFilters) {
	o.Filter.Set(&v)
}
// SetFilterNil sets the value for Filter to be an explicit nil
func (o *EventSubscriptions) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnsetFilter ensures that no value is present for Filter, not even an explicit nil
func (o *EventSubscriptions) UnsetFilter() {
	o.Filter.Unset()
}

// GetItems returns the Items field value
func (o *EventSubscriptions) GetItems() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *EventSubscriptions) GetItemsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *EventSubscriptions) SetItems(v []string) {
	o.Items = v
}

// GetType returns the Type field value
func (o *EventSubscriptions) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventSubscriptions) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EventSubscriptions) SetType(v string) {
	o.Type = v
}

func (o EventSubscriptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter.IsSet() {
		toSerialize["filter"] = o.Filter.Get()
	}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EventSubscriptions) UnmarshalJSON(bytes []byte) (err error) {
	varEventSubscriptions := _EventSubscriptions{}

	err = json.Unmarshal(bytes, &varEventSubscriptions)
	if err == nil {
		*o = EventSubscriptions(varEventSubscriptions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "filter")
		delete(additionalProperties, "items")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEventSubscriptions struct {
	value *EventSubscriptions
	isSet bool
}

func (v NullableEventSubscriptions) Get() *EventSubscriptions {
	return v.value
}

func (v *NullableEventSubscriptions) Set(val *EventSubscriptions) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSubscriptions) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSubscriptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSubscriptions(val *EventSubscriptions) *NullableEventSubscriptions {
	return &NullableEventSubscriptions{value: val, isSet: true}
}

func (v NullableEventSubscriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSubscriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

