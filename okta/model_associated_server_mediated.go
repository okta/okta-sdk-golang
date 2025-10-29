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

// checks if the AssociatedServerMediated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssociatedServerMediated{}

// AssociatedServerMediated struct for AssociatedServerMediated
type AssociatedServerMediated struct {
	// A list of the authorization server IDs
	Trusted              []string `json:"trusted,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssociatedServerMediated AssociatedServerMediated

// NewAssociatedServerMediated instantiates a new AssociatedServerMediated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociatedServerMediated() *AssociatedServerMediated {
	this := AssociatedServerMediated{}
	return &this
}

// NewAssociatedServerMediatedWithDefaults instantiates a new AssociatedServerMediated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociatedServerMediatedWithDefaults() *AssociatedServerMediated {
	this := AssociatedServerMediated{}
	return &this
}

// GetTrusted returns the Trusted field value if set, zero value otherwise.
func (o *AssociatedServerMediated) GetTrusted() []string {
	if o == nil || IsNil(o.Trusted) {
		var ret []string
		return ret
	}
	return o.Trusted
}

// GetTrustedOk returns a tuple with the Trusted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedServerMediated) GetTrustedOk() ([]string, bool) {
	if o == nil || IsNil(o.Trusted) {
		return nil, false
	}
	return o.Trusted, true
}

// HasTrusted returns a boolean if a field has been set.
func (o *AssociatedServerMediated) HasTrusted() bool {
	if o != nil && !IsNil(o.Trusted) {
		return true
	}

	return false
}

// SetTrusted gets a reference to the given []string and assigns it to the Trusted field.
func (o *AssociatedServerMediated) SetTrusted(v []string) {
	o.Trusted = v
}

func (o AssociatedServerMediated) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssociatedServerMediated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Trusted) {
		toSerialize["trusted"] = o.Trusted
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssociatedServerMediated) UnmarshalJSON(data []byte) (err error) {
	varAssociatedServerMediated := _AssociatedServerMediated{}

	err = json.Unmarshal(data, &varAssociatedServerMediated)

	if err != nil {
		return err
	}

	*o = AssociatedServerMediated(varAssociatedServerMediated)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "trusted")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssociatedServerMediated struct {
	value *AssociatedServerMediated
	isSet bool
}

func (v NullableAssociatedServerMediated) Get() *AssociatedServerMediated {
	return v.value
}

func (v *NullableAssociatedServerMediated) Set(val *AssociatedServerMediated) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociatedServerMediated) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociatedServerMediated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociatedServerMediated(val *AssociatedServerMediated) *NullableAssociatedServerMediated {
	return &NullableAssociatedServerMediated{value: val, isSet: true}
}

func (v NullableAssociatedServerMediated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociatedServerMediated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
