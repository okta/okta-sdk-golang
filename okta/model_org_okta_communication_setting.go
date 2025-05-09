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

// OrgOktaCommunicationSetting struct for OrgOktaCommunicationSetting
type OrgOktaCommunicationSetting struct {
	OptOutEmailUsers *bool `json:"optOutEmailUsers,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
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
	if o == nil || o.OptOutEmailUsers == nil {
		var ret bool
		return ret
	}
	return *o.OptOutEmailUsers
}

// GetOptOutEmailUsersOk returns a tuple with the OptOutEmailUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOktaCommunicationSetting) GetOptOutEmailUsersOk() (*bool, bool) {
	if o == nil || o.OptOutEmailUsers == nil {
		return nil, false
	}
	return o.OptOutEmailUsers, true
}

// HasOptOutEmailUsers returns a boolean if a field has been set.
func (o *OrgOktaCommunicationSetting) HasOptOutEmailUsers() bool {
	if o != nil && o.OptOutEmailUsers != nil {
		return true
	}

	return false
}

// SetOptOutEmailUsers gets a reference to the given bool and assigns it to the OptOutEmailUsers field.
func (o *OrgOktaCommunicationSetting) SetOptOutEmailUsers(v bool) {
	o.OptOutEmailUsers = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrgOktaCommunicationSetting) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOktaCommunicationSetting) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrgOktaCommunicationSetting) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *OrgOktaCommunicationSetting) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o OrgOktaCommunicationSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OptOutEmailUsers != nil {
		toSerialize["optOutEmailUsers"] = o.OptOutEmailUsers
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OrgOktaCommunicationSetting) UnmarshalJSON(bytes []byte) (err error) {
	varOrgOktaCommunicationSetting := _OrgOktaCommunicationSetting{}

	err = json.Unmarshal(bytes, &varOrgOktaCommunicationSetting)
	if err == nil {
		*o = OrgOktaCommunicationSetting(varOrgOktaCommunicationSetting)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "optOutEmailUsers")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

