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

// EventHookLinks struct for EventHookLinks
type EventHookLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Deactivate *HrefObject `json:"deactivate,omitempty"`
	Verify *HrefObject `json:"verify,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventHookLinks EventHookLinks

// NewEventHookLinks instantiates a new EventHookLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventHookLinks() *EventHookLinks {
	this := EventHookLinks{}
	return &this
}

// NewEventHookLinksWithDefaults instantiates a new EventHookLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventHookLinksWithDefaults() *EventHookLinks {
	this := EventHookLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *EventHookLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHookLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *EventHookLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *EventHookLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *EventHookLinks) GetDeactivate() HrefObject {
	if o == nil || o.Deactivate == nil {
		var ret HrefObject
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHookLinks) GetDeactivateOk() (*HrefObject, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *EventHookLinks) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObject and assigns it to the Deactivate field.
func (o *EventHookLinks) SetDeactivate(v HrefObject) {
	o.Deactivate = &v
}

// GetVerify returns the Verify field value if set, zero value otherwise.
func (o *EventHookLinks) GetVerify() HrefObject {
	if o == nil || o.Verify == nil {
		var ret HrefObject
		return ret
	}
	return *o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHookLinks) GetVerifyOk() (*HrefObject, bool) {
	if o == nil || o.Verify == nil {
		return nil, false
	}
	return o.Verify, true
}

// HasVerify returns a boolean if a field has been set.
func (o *EventHookLinks) HasVerify() bool {
	if o != nil && o.Verify != nil {
		return true
	}

	return false
}

// SetVerify gets a reference to the given HrefObject and assigns it to the Verify field.
func (o *EventHookLinks) SetVerify(v HrefObject) {
	o.Verify = &v
}

func (o EventHookLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Deactivate != nil {
		toSerialize["deactivate"] = o.Deactivate
	}
	if o.Verify != nil {
		toSerialize["verify"] = o.Verify
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EventHookLinks) UnmarshalJSON(bytes []byte) (err error) {
	varEventHookLinks := _EventHookLinks{}

	err = json.Unmarshal(bytes, &varEventHookLinks)
	if err == nil {
		*o = EventHookLinks(varEventHookLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "verify")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEventHookLinks struct {
	value *EventHookLinks
	isSet bool
}

func (v NullableEventHookLinks) Get() *EventHookLinks {
	return v.value
}

func (v *NullableEventHookLinks) Set(val *EventHookLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEventHookLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEventHookLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventHookLinks(val *EventHookLinks) *NullableEventHookLinks {
	return &NullableEventHookLinks{value: val, isSet: true}
}

func (v NullableEventHookLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventHookLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

