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

// checks if the SelfServiceRegistrationPolicyRuleConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelfServiceRegistrationPolicyRuleConditions{}

// SelfServiceRegistrationPolicyRuleConditions <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Specifies conditions that must be met for this rule to apply. A maximum of 10 rules are supported per policy. A rule matches when all specified conditions are satisfied.
type SelfServiceRegistrationPolicyRuleConditions struct {
	EmailDomains         *SelfServiceRegistrationEmailDomainsCondition `json:"emailDomains,omitempty"`
	Network              *PolicyNetworkCondition                       `json:"network,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SelfServiceRegistrationPolicyRuleConditions SelfServiceRegistrationPolicyRuleConditions

// NewSelfServiceRegistrationPolicyRuleConditions instantiates a new SelfServiceRegistrationPolicyRuleConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceRegistrationPolicyRuleConditions() *SelfServiceRegistrationPolicyRuleConditions {
	this := SelfServiceRegistrationPolicyRuleConditions{}
	return &this
}

// NewSelfServiceRegistrationPolicyRuleConditionsWithDefaults instantiates a new SelfServiceRegistrationPolicyRuleConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceRegistrationPolicyRuleConditionsWithDefaults() *SelfServiceRegistrationPolicyRuleConditions {
	this := SelfServiceRegistrationPolicyRuleConditions{}
	return &this
}

// GetEmailDomains returns the EmailDomains field value if set, zero value otherwise.
func (o *SelfServiceRegistrationPolicyRuleConditions) GetEmailDomains() SelfServiceRegistrationEmailDomainsCondition {
	if o == nil || IsNil(o.EmailDomains) {
		var ret SelfServiceRegistrationEmailDomainsCondition
		return ret
	}
	return *o.EmailDomains
}

// GetEmailDomainsOk returns a tuple with the EmailDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceRegistrationPolicyRuleConditions) GetEmailDomainsOk() (*SelfServiceRegistrationEmailDomainsCondition, bool) {
	if o == nil || IsNil(o.EmailDomains) {
		return nil, false
	}
	return o.EmailDomains, true
}

// HasEmailDomains returns a boolean if a field has been set.
func (o *SelfServiceRegistrationPolicyRuleConditions) HasEmailDomains() bool {
	if o != nil && !IsNil(o.EmailDomains) {
		return true
	}

	return false
}

// SetEmailDomains gets a reference to the given SelfServiceRegistrationEmailDomainsCondition and assigns it to the EmailDomains field.
func (o *SelfServiceRegistrationPolicyRuleConditions) SetEmailDomains(v SelfServiceRegistrationEmailDomainsCondition) {
	o.EmailDomains = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *SelfServiceRegistrationPolicyRuleConditions) GetNetwork() PolicyNetworkCondition {
	if o == nil || IsNil(o.Network) {
		var ret PolicyNetworkCondition
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceRegistrationPolicyRuleConditions) GetNetworkOk() (*PolicyNetworkCondition, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *SelfServiceRegistrationPolicyRuleConditions) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given PolicyNetworkCondition and assigns it to the Network field.
func (o *SelfServiceRegistrationPolicyRuleConditions) SetNetwork(v PolicyNetworkCondition) {
	o.Network = &v
}

func (o SelfServiceRegistrationPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelfServiceRegistrationPolicyRuleConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailDomains) {
		toSerialize["emailDomains"] = o.EmailDomains
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SelfServiceRegistrationPolicyRuleConditions) UnmarshalJSON(data []byte) (err error) {
	varSelfServiceRegistrationPolicyRuleConditions := _SelfServiceRegistrationPolicyRuleConditions{}

	err = json.Unmarshal(data, &varSelfServiceRegistrationPolicyRuleConditions)

	if err != nil {
		return err
	}

	*o = SelfServiceRegistrationPolicyRuleConditions(varSelfServiceRegistrationPolicyRuleConditions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "emailDomains")
		delete(additionalProperties, "network")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelfServiceRegistrationPolicyRuleConditions struct {
	value *SelfServiceRegistrationPolicyRuleConditions
	isSet bool
}

func (v NullableSelfServiceRegistrationPolicyRuleConditions) Get() *SelfServiceRegistrationPolicyRuleConditions {
	return v.value
}

func (v *NullableSelfServiceRegistrationPolicyRuleConditions) Set(val *SelfServiceRegistrationPolicyRuleConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceRegistrationPolicyRuleConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceRegistrationPolicyRuleConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceRegistrationPolicyRuleConditions(val *SelfServiceRegistrationPolicyRuleConditions) *NullableSelfServiceRegistrationPolicyRuleConditions {
	return &NullableSelfServiceRegistrationPolicyRuleConditions{value: val, isSet: true}
}

func (v NullableSelfServiceRegistrationPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceRegistrationPolicyRuleConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
