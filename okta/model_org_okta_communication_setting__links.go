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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the OrgOktaCommunicationSettingLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgOktaCommunicationSettingLinks{}

// OrgOktaCommunicationSettingLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for this object using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification
type OrgOktaCommunicationSettingLinks struct {
	// Link to opt users in to communication emails
	OptIn *HrefObject `json:"optIn,omitempty"`
	// Link to opt users out of communication emails
	OptOut               *HrefObject `json:"optOut,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgOktaCommunicationSettingLinks OrgOktaCommunicationSettingLinks

// NewOrgOktaCommunicationSettingLinks instantiates a new OrgOktaCommunicationSettingLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgOktaCommunicationSettingLinks() *OrgOktaCommunicationSettingLinks {
	this := OrgOktaCommunicationSettingLinks{}
	return &this
}

// NewOrgOktaCommunicationSettingLinksWithDefaults instantiates a new OrgOktaCommunicationSettingLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgOktaCommunicationSettingLinksWithDefaults() *OrgOktaCommunicationSettingLinks {
	this := OrgOktaCommunicationSettingLinks{}
	return &this
}

// GetOptIn returns the OptIn field value if set, zero value otherwise.
func (o *OrgOktaCommunicationSettingLinks) GetOptIn() HrefObject {
	if o == nil || IsNil(o.OptIn) {
		var ret HrefObject
		return ret
	}
	return *o.OptIn
}

// GetOptInOk returns a tuple with the OptIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOktaCommunicationSettingLinks) GetOptInOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.OptIn) {
		return nil, false
	}
	return o.OptIn, true
}

// HasOptIn returns a boolean if a field has been set.
func (o *OrgOktaCommunicationSettingLinks) HasOptIn() bool {
	if o != nil && !IsNil(o.OptIn) {
		return true
	}

	return false
}

// SetOptIn gets a reference to the given HrefObject and assigns it to the OptIn field.
func (o *OrgOktaCommunicationSettingLinks) SetOptIn(v HrefObject) {
	o.OptIn = &v
}

// GetOptOut returns the OptOut field value if set, zero value otherwise.
func (o *OrgOktaCommunicationSettingLinks) GetOptOut() HrefObject {
	if o == nil || IsNil(o.OptOut) {
		var ret HrefObject
		return ret
	}
	return *o.OptOut
}

// GetOptOutOk returns a tuple with the OptOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOktaCommunicationSettingLinks) GetOptOutOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.OptOut) {
		return nil, false
	}
	return o.OptOut, true
}

// HasOptOut returns a boolean if a field has been set.
func (o *OrgOktaCommunicationSettingLinks) HasOptOut() bool {
	if o != nil && !IsNil(o.OptOut) {
		return true
	}

	return false
}

// SetOptOut gets a reference to the given HrefObject and assigns it to the OptOut field.
func (o *OrgOktaCommunicationSettingLinks) SetOptOut(v HrefObject) {
	o.OptOut = &v
}

func (o OrgOktaCommunicationSettingLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgOktaCommunicationSettingLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptIn) {
		toSerialize["optIn"] = o.OptIn
	}
	if !IsNil(o.OptOut) {
		toSerialize["optOut"] = o.OptOut
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgOktaCommunicationSettingLinks) UnmarshalJSON(data []byte) (err error) {
	varOrgOktaCommunicationSettingLinks := _OrgOktaCommunicationSettingLinks{}

	err = json.Unmarshal(data, &varOrgOktaCommunicationSettingLinks)

	if err != nil {
		return err
	}

	*o = OrgOktaCommunicationSettingLinks(varOrgOktaCommunicationSettingLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "optIn")
		delete(additionalProperties, "optOut")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgOktaCommunicationSettingLinks struct {
	value *OrgOktaCommunicationSettingLinks
	isSet bool
}

func (v NullableOrgOktaCommunicationSettingLinks) Get() *OrgOktaCommunicationSettingLinks {
	return v.value
}

func (v *NullableOrgOktaCommunicationSettingLinks) Set(val *OrgOktaCommunicationSettingLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgOktaCommunicationSettingLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgOktaCommunicationSettingLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgOktaCommunicationSettingLinks(val *OrgOktaCommunicationSettingLinks) *NullableOrgOktaCommunicationSettingLinks {
	return &NullableOrgOktaCommunicationSettingLinks{value: val, isSet: true}
}

func (v NullableOrgOktaCommunicationSettingLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgOktaCommunicationSettingLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
