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

// checks if the OrgOktaCommunicationSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgOktaCommunicationSetting{}

// OrgOktaCommunicationSetting struct for OrgOktaCommunicationSetting
type OrgOktaCommunicationSetting struct {
	// Indicates whether org users receive Okta communication emails
	OptOutEmailUsers     *bool                             `json:"optOutEmailUsers,omitempty"`
	Links                *OrgOktaCommunicationSettingLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgOktaCommunicationSetting OrgOktaCommunicationSetting

// NewOrgOktaCommunicationSetting instantiates a new OrgOktaCommunicationSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgOktaCommunicationSetting() *OrgOktaCommunicationSetting {
	this := OrgOktaCommunicationSetting{}
	return &this
}

// NewOrgOktaCommunicationSettingWithDefaults instantiates a new OrgOktaCommunicationSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgOktaCommunicationSettingWithDefaults() *OrgOktaCommunicationSetting {
	this := OrgOktaCommunicationSetting{}
	return &this
}

// GetOptOutEmailUsers returns the OptOutEmailUsers field value if set, zero value otherwise.
func (o *OrgOktaCommunicationSetting) GetOptOutEmailUsers() bool {
	if o == nil || IsNil(o.OptOutEmailUsers) {
		var ret bool
		return ret
	}
	return *o.OptOutEmailUsers
}

// GetOptOutEmailUsersOk returns a tuple with the OptOutEmailUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOktaCommunicationSetting) GetOptOutEmailUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.OptOutEmailUsers) {
		return nil, false
	}
	return o.OptOutEmailUsers, true
}

// HasOptOutEmailUsers returns a boolean if a field has been set.
func (o *OrgOktaCommunicationSetting) HasOptOutEmailUsers() bool {
	if o != nil && !IsNil(o.OptOutEmailUsers) {
		return true
	}

	return false
}

// SetOptOutEmailUsers gets a reference to the given bool and assigns it to the OptOutEmailUsers field.
func (o *OrgOktaCommunicationSetting) SetOptOutEmailUsers(v bool) {
	o.OptOutEmailUsers = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrgOktaCommunicationSetting) GetLinks() OrgOktaCommunicationSettingLinks {
	if o == nil || IsNil(o.Links) {
		var ret OrgOktaCommunicationSettingLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOktaCommunicationSetting) GetLinksOk() (*OrgOktaCommunicationSettingLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrgOktaCommunicationSetting) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OrgOktaCommunicationSettingLinks and assigns it to the Links field.
func (o *OrgOktaCommunicationSetting) SetLinks(v OrgOktaCommunicationSettingLinks) {
	o.Links = &v
}

func (o OrgOktaCommunicationSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgOktaCommunicationSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptOutEmailUsers) {
		toSerialize["optOutEmailUsers"] = o.OptOutEmailUsers
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgOktaCommunicationSetting) UnmarshalJSON(data []byte) (err error) {
	varOrgOktaCommunicationSetting := _OrgOktaCommunicationSetting{}

	err = json.Unmarshal(data, &varOrgOktaCommunicationSetting)

	if err != nil {
		return err
	}

	*o = OrgOktaCommunicationSetting(varOrgOktaCommunicationSetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "optOutEmailUsers")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgOktaCommunicationSetting struct {
	value *OrgOktaCommunicationSetting
	isSet bool
}

func (v NullableOrgOktaCommunicationSetting) Get() *OrgOktaCommunicationSetting {
	return v.value
}

func (v *NullableOrgOktaCommunicationSetting) Set(val *OrgOktaCommunicationSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgOktaCommunicationSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgOktaCommunicationSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgOktaCommunicationSetting(val *OrgOktaCommunicationSetting) *NullableOrgOktaCommunicationSetting {
	return &NullableOrgOktaCommunicationSetting{value: val, isSet: true}
}

func (v NullableOrgOktaCommunicationSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgOktaCommunicationSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
