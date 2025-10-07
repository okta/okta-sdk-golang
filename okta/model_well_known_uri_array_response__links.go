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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the WellKnownURIArrayResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WellKnownURIArrayResponseLinks{}

// WellKnownURIArrayResponseLinks struct for WellKnownURIArrayResponseLinks
type WellKnownURIArrayResponseLinks struct {
	Self                 *HrefObject `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WellKnownURIArrayResponseLinks WellKnownURIArrayResponseLinks

// NewWellKnownURIArrayResponseLinks instantiates a new WellKnownURIArrayResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWellKnownURIArrayResponseLinks() *WellKnownURIArrayResponseLinks {
	this := WellKnownURIArrayResponseLinks{}
	return &this
}

// NewWellKnownURIArrayResponseLinksWithDefaults instantiates a new WellKnownURIArrayResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellKnownURIArrayResponseLinksWithDefaults() *WellKnownURIArrayResponseLinks {
	this := WellKnownURIArrayResponseLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *WellKnownURIArrayResponseLinks) GetSelf() HrefObject {
	if o == nil || IsNil(o.Self) {
		var ret HrefObject
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownURIArrayResponseLinks) GetSelfOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *WellKnownURIArrayResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObject and assigns it to the Self field.
func (o *WellKnownURIArrayResponseLinks) SetSelf(v HrefObject) {
	o.Self = &v
}

func (o WellKnownURIArrayResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WellKnownURIArrayResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WellKnownURIArrayResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varWellKnownURIArrayResponseLinks := _WellKnownURIArrayResponseLinks{}

	err = json.Unmarshal(data, &varWellKnownURIArrayResponseLinks)

	if err != nil {
		return err
	}

	*o = WellKnownURIArrayResponseLinks(varWellKnownURIArrayResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWellKnownURIArrayResponseLinks struct {
	value *WellKnownURIArrayResponseLinks
	isSet bool
}

func (v NullableWellKnownURIArrayResponseLinks) Get() *WellKnownURIArrayResponseLinks {
	return v.value
}

func (v *NullableWellKnownURIArrayResponseLinks) Set(val *WellKnownURIArrayResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableWellKnownURIArrayResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableWellKnownURIArrayResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWellKnownURIArrayResponseLinks(val *WellKnownURIArrayResponseLinks) *NullableWellKnownURIArrayResponseLinks {
	return &NullableWellKnownURIArrayResponseLinks{value: val, isSet: true}
}

func (v NullableWellKnownURIArrayResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWellKnownURIArrayResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
