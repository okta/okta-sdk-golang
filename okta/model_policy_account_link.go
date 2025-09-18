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

// checks if the PolicyAccountLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyAccountLink{}

// PolicyAccountLink Specifies the behavior for linking an IdP user to an existing Okta user
type PolicyAccountLink struct {
	// Specifies the account linking action for an IdP user
	Action               *string                  `json:"action,omitempty"`
	Filter               *PolicyAccountLinkFilter `json:"filter,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAccountLink PolicyAccountLink

// NewPolicyAccountLink instantiates a new PolicyAccountLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAccountLink() *PolicyAccountLink {
	this := PolicyAccountLink{}
	return &this
}

// NewPolicyAccountLinkWithDefaults instantiates a new PolicyAccountLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAccountLinkWithDefaults() *PolicyAccountLink {
	this := PolicyAccountLink{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PolicyAccountLink) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAccountLink) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PolicyAccountLink) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PolicyAccountLink) SetAction(v string) {
	o.Action = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *PolicyAccountLink) GetFilter() PolicyAccountLinkFilter {
	if o == nil || IsNil(o.Filter) {
		var ret PolicyAccountLinkFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAccountLink) GetFilterOk() (*PolicyAccountLinkFilter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *PolicyAccountLink) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given PolicyAccountLinkFilter and assigns it to the Filter field.
func (o *PolicyAccountLink) SetFilter(v PolicyAccountLinkFilter) {
	o.Filter = &v
}

func (o PolicyAccountLink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyAccountLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyAccountLink) UnmarshalJSON(data []byte) (err error) {
	varPolicyAccountLink := _PolicyAccountLink{}

	err = json.Unmarshal(data, &varPolicyAccountLink)

	if err != nil {
		return err
	}

	*o = PolicyAccountLink(varPolicyAccountLink)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "filter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyAccountLink struct {
	value *PolicyAccountLink
	isSet bool
}

func (v NullablePolicyAccountLink) Get() *PolicyAccountLink {
	return v.value
}

func (v *NullablePolicyAccountLink) Set(val *PolicyAccountLink) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAccountLink) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAccountLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAccountLink(val *PolicyAccountLink) *NullablePolicyAccountLink {
	return &NullablePolicyAccountLink{value: val, isSet: true}
}

func (v NullablePolicyAccountLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAccountLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
