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

// checks if the OrgCAPTCHASettingsLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgCAPTCHASettingsLinks{}

// OrgCAPTCHASettingsLinks Link relations for the CAPTCHA settings object
type OrgCAPTCHASettingsLinks struct {
	Self                 *HrefObject `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgCAPTCHASettingsLinks OrgCAPTCHASettingsLinks

// NewOrgCAPTCHASettingsLinks instantiates a new OrgCAPTCHASettingsLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgCAPTCHASettingsLinks() *OrgCAPTCHASettingsLinks {
	this := OrgCAPTCHASettingsLinks{}
	return &this
}

// NewOrgCAPTCHASettingsLinksWithDefaults instantiates a new OrgCAPTCHASettingsLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgCAPTCHASettingsLinksWithDefaults() *OrgCAPTCHASettingsLinks {
	this := OrgCAPTCHASettingsLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *OrgCAPTCHASettingsLinks) GetSelf() HrefObject {
	if o == nil || IsNil(o.Self) {
		var ret HrefObject
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCAPTCHASettingsLinks) GetSelfOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *OrgCAPTCHASettingsLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObject and assigns it to the Self field.
func (o *OrgCAPTCHASettingsLinks) SetSelf(v HrefObject) {
	o.Self = &v
}

func (o OrgCAPTCHASettingsLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgCAPTCHASettingsLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgCAPTCHASettingsLinks) UnmarshalJSON(data []byte) (err error) {
	varOrgCAPTCHASettingsLinks := _OrgCAPTCHASettingsLinks{}

	err = json.Unmarshal(data, &varOrgCAPTCHASettingsLinks)

	if err != nil {
		return err
	}

	*o = OrgCAPTCHASettingsLinks(varOrgCAPTCHASettingsLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgCAPTCHASettingsLinks struct {
	value *OrgCAPTCHASettingsLinks
	isSet bool
}

func (v NullableOrgCAPTCHASettingsLinks) Get() *OrgCAPTCHASettingsLinks {
	return v.value
}

func (v *NullableOrgCAPTCHASettingsLinks) Set(val *OrgCAPTCHASettingsLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgCAPTCHASettingsLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgCAPTCHASettingsLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgCAPTCHASettingsLinks(val *OrgCAPTCHASettingsLinks) *NullableOrgCAPTCHASettingsLinks {
	return &NullableOrgCAPTCHASettingsLinks{value: val, isSet: true}
}

func (v NullableOrgCAPTCHASettingsLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgCAPTCHASettingsLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
