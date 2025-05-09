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

// NetworkZoneAddress Specifies the value of an IP address expressed using either `range` or `CIDR` form.
type NetworkZoneAddress struct {
	// Format of the IP addresses
	Type *string `json:"type,omitempty"`
	// Value in CIDR/range form, depending on the `type` specified
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkZoneAddress NetworkZoneAddress

// NewNetworkZoneAddress instantiates a new NetworkZoneAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkZoneAddress() *NetworkZoneAddress {
	this := NetworkZoneAddress{}
	return &this
}

// NewNetworkZoneAddressWithDefaults instantiates a new NetworkZoneAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkZoneAddressWithDefaults() *NetworkZoneAddress {
	this := NetworkZoneAddress{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkZoneAddress) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkZoneAddress) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkZoneAddress) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworkZoneAddress) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *NetworkZoneAddress) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkZoneAddress) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *NetworkZoneAddress) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *NetworkZoneAddress) SetValue(v string) {
	o.Value = &v
}

func (o NetworkZoneAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkZoneAddress) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkZoneAddress := _NetworkZoneAddress{}

	err = json.Unmarshal(bytes, &varNetworkZoneAddress)
	if err == nil {
		*o = NetworkZoneAddress(varNetworkZoneAddress)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableNetworkZoneAddress struct {
	value *NetworkZoneAddress
	isSet bool
}

func (v NullableNetworkZoneAddress) Get() *NetworkZoneAddress {
	return v.value
}

func (v *NullableNetworkZoneAddress) Set(val *NetworkZoneAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkZoneAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkZoneAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkZoneAddress(val *NetworkZoneAddress) *NullableNetworkZoneAddress {
	return &NullableNetworkZoneAddress{value: val, isSet: true}
}

func (v NullableNetworkZoneAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkZoneAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

