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

// checks if the GracePeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GracePeriod{}

// GracePeriod <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>Represents the Grace Period configuration for the device assurance policy
type GracePeriod struct {
	Expiry *GracePeriodExpiry `json:"expiry,omitempty"`
	// Represents the type of Grace Period configured for the device assurance policy
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GracePeriod GracePeriod

// NewGracePeriod instantiates a new GracePeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGracePeriod() *GracePeriod {
	this := GracePeriod{}
	return &this
}

// NewGracePeriodWithDefaults instantiates a new GracePeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGracePeriodWithDefaults() *GracePeriod {
	this := GracePeriod{}
	return &this
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *GracePeriod) GetExpiry() GracePeriodExpiry {
	if o == nil || IsNil(o.Expiry) {
		var ret GracePeriodExpiry
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GracePeriod) GetExpiryOk() (*GracePeriodExpiry, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *GracePeriod) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given GracePeriodExpiry and assigns it to the Expiry field.
func (o *GracePeriod) SetExpiry(v GracePeriodExpiry) {
	o.Expiry = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GracePeriod) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GracePeriod) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GracePeriod) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GracePeriod) SetType(v string) {
	o.Type = &v
}

func (o GracePeriod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GracePeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GracePeriod) UnmarshalJSON(data []byte) (err error) {
	varGracePeriod := _GracePeriod{}

	err = json.Unmarshal(data, &varGracePeriod)

	if err != nil {
		return err
	}

	*o = GracePeriod(varGracePeriod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expiry")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGracePeriod struct {
	value *GracePeriod
	isSet bool
}

func (v NullableGracePeriod) Get() *GracePeriod {
	return v.value
}

func (v *NullableGracePeriod) Set(val *GracePeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableGracePeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableGracePeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGracePeriod(val *GracePeriod) *NullableGracePeriod {
	return &NullableGracePeriod{value: val, isSet: true}
}

func (v NullableGracePeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGracePeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
