/*
Okta Admin Management API

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

// checks if the SelfServiceRegistrationEmailDomainsCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelfServiceRegistrationEmailDomainsCondition{}

// SelfServiceRegistrationEmailDomainsCondition <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Specifies email domain conditions for the rule
type SelfServiceRegistrationEmailDomainsCondition struct {
	// List of email domain patterns to match. Only `ALL_SUSPICIOUS_EMAIL_DOMAINS` is supported.
	Include              []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SelfServiceRegistrationEmailDomainsCondition SelfServiceRegistrationEmailDomainsCondition

// NewSelfServiceRegistrationEmailDomainsCondition instantiates a new SelfServiceRegistrationEmailDomainsCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceRegistrationEmailDomainsCondition() *SelfServiceRegistrationEmailDomainsCondition {
	this := SelfServiceRegistrationEmailDomainsCondition{}
	return &this
}

// NewSelfServiceRegistrationEmailDomainsConditionWithDefaults instantiates a new SelfServiceRegistrationEmailDomainsCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceRegistrationEmailDomainsConditionWithDefaults() *SelfServiceRegistrationEmailDomainsCondition {
	this := SelfServiceRegistrationEmailDomainsCondition{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *SelfServiceRegistrationEmailDomainsCondition) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceRegistrationEmailDomainsCondition) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *SelfServiceRegistrationEmailDomainsCondition) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *SelfServiceRegistrationEmailDomainsCondition) SetInclude(v []string) {
	o.Include = v
}

func (o SelfServiceRegistrationEmailDomainsCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelfServiceRegistrationEmailDomainsCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SelfServiceRegistrationEmailDomainsCondition) UnmarshalJSON(data []byte) (err error) {
	varSelfServiceRegistrationEmailDomainsCondition := _SelfServiceRegistrationEmailDomainsCondition{}

	err = json.Unmarshal(data, &varSelfServiceRegistrationEmailDomainsCondition)

	if err != nil {
		return err
	}

	*o = SelfServiceRegistrationEmailDomainsCondition(varSelfServiceRegistrationEmailDomainsCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelfServiceRegistrationEmailDomainsCondition struct {
	value *SelfServiceRegistrationEmailDomainsCondition
	isSet bool
}

func (v NullableSelfServiceRegistrationEmailDomainsCondition) Get() *SelfServiceRegistrationEmailDomainsCondition {
	return v.value
}

func (v *NullableSelfServiceRegistrationEmailDomainsCondition) Set(val *SelfServiceRegistrationEmailDomainsCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceRegistrationEmailDomainsCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceRegistrationEmailDomainsCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceRegistrationEmailDomainsCondition(val *SelfServiceRegistrationEmailDomainsCondition) *NullableSelfServiceRegistrationEmailDomainsCondition {
	return &NullableSelfServiceRegistrationEmailDomainsCondition{value: val, isSet: true}
}

func (v NullableSelfServiceRegistrationEmailDomainsCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceRegistrationEmailDomainsCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
