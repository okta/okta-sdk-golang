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

// LogTargetChangeDetails Details on the target's changes. Not all event types support the `changeDetails` property, and not all target objects contain the `changeDetails` property.You must include a property within the object. When querying on this property, you can't search on the `to` or `from` objects alone. You must include a property within the object.
type LogTargetChangeDetails struct {
	// The original properties of the target
	From map[string]interface{} `json:"from,omitempty"`
	// The updated properties of the target
	To map[string]interface{} `json:"to,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogTargetChangeDetails LogTargetChangeDetails

// NewLogTargetChangeDetails instantiates a new LogTargetChangeDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogTargetChangeDetails() *LogTargetChangeDetails {
	this := LogTargetChangeDetails{}
	return &this
}

// NewLogTargetChangeDetailsWithDefaults instantiates a new LogTargetChangeDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogTargetChangeDetailsWithDefaults() *LogTargetChangeDetails {
	this := LogTargetChangeDetails{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *LogTargetChangeDetails) GetFrom() map[string]interface{} {
	if o == nil || o.From == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogTargetChangeDetails) GetFromOk() (map[string]interface{}, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *LogTargetChangeDetails) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given map[string]interface{} and assigns it to the From field.
func (o *LogTargetChangeDetails) SetFrom(v map[string]interface{}) {
	o.From = v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *LogTargetChangeDetails) GetTo() map[string]interface{} {
	if o == nil || o.To == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogTargetChangeDetails) GetToOk() (map[string]interface{}, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *LogTargetChangeDetails) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given map[string]interface{} and assigns it to the To field.
func (o *LogTargetChangeDetails) SetTo(v map[string]interface{}) {
	o.To = v
}

func (o LogTargetChangeDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogTargetChangeDetails) UnmarshalJSON(bytes []byte) (err error) {
	varLogTargetChangeDetails := _LogTargetChangeDetails{}

	err = json.Unmarshal(bytes, &varLogTargetChangeDetails)
	if err == nil {
		*o = LogTargetChangeDetails(varLogTargetChangeDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogTargetChangeDetails struct {
	value *LogTargetChangeDetails
	isSet bool
}

func (v NullableLogTargetChangeDetails) Get() *LogTargetChangeDetails {
	return v.value
}

func (v *NullableLogTargetChangeDetails) Set(val *LogTargetChangeDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableLogTargetChangeDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableLogTargetChangeDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogTargetChangeDetails(val *LogTargetChangeDetails) *NullableLogTargetChangeDetails {
	return &NullableLogTargetChangeDetails{value: val, isSet: true}
}

func (v NullableLogTargetChangeDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogTargetChangeDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

