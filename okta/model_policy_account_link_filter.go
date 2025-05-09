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

// PolicyAccountLinkFilter struct for PolicyAccountLinkFilter
type PolicyAccountLinkFilter struct {
	Groups *PolicyAccountLinkFilterGroups `json:"groups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAccountLinkFilter PolicyAccountLinkFilter

// NewPolicyAccountLinkFilter instantiates a new PolicyAccountLinkFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAccountLinkFilter() *PolicyAccountLinkFilter {
	this := PolicyAccountLinkFilter{}
	return &this
}

// NewPolicyAccountLinkFilterWithDefaults instantiates a new PolicyAccountLinkFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAccountLinkFilterWithDefaults() *PolicyAccountLinkFilter {
	this := PolicyAccountLinkFilter{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *PolicyAccountLinkFilter) GetGroups() PolicyAccountLinkFilterGroups {
	if o == nil || o.Groups == nil {
		var ret PolicyAccountLinkFilterGroups
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAccountLinkFilter) GetGroupsOk() (*PolicyAccountLinkFilterGroups, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *PolicyAccountLinkFilter) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given PolicyAccountLinkFilterGroups and assigns it to the Groups field.
func (o *PolicyAccountLinkFilter) SetGroups(v PolicyAccountLinkFilterGroups) {
	o.Groups = &v
}

func (o PolicyAccountLinkFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyAccountLinkFilter) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyAccountLinkFilter := _PolicyAccountLinkFilter{}

	err = json.Unmarshal(bytes, &varPolicyAccountLinkFilter)
	if err == nil {
		*o = PolicyAccountLinkFilter(varPolicyAccountLinkFilter)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "groups")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePolicyAccountLinkFilter struct {
	value *PolicyAccountLinkFilter
	isSet bool
}

func (v NullablePolicyAccountLinkFilter) Get() *PolicyAccountLinkFilter {
	return v.value
}

func (v *NullablePolicyAccountLinkFilter) Set(val *PolicyAccountLinkFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAccountLinkFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAccountLinkFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAccountLinkFilter(val *PolicyAccountLinkFilter) *NullablePolicyAccountLinkFilter {
	return &NullablePolicyAccountLinkFilter{value: val, isSet: true}
}

func (v NullablePolicyAccountLinkFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAccountLinkFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

