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

// checks if the PolicyAccountLinkFilterGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyAccountLinkFilterGroups{}

// PolicyAccountLinkFilterGroups Group memberships used to determine link candidates
type PolicyAccountLinkFilterGroups struct {
	// Specifies the allowlist of Group identifiers to match against. Group memberships are restricted to type `OKTA_GROUP`.
	Include              []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAccountLinkFilterGroups PolicyAccountLinkFilterGroups

// NewPolicyAccountLinkFilterGroups instantiates a new PolicyAccountLinkFilterGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAccountLinkFilterGroups() *PolicyAccountLinkFilterGroups {
	this := PolicyAccountLinkFilterGroups{}
	return &this
}

// NewPolicyAccountLinkFilterGroupsWithDefaults instantiates a new PolicyAccountLinkFilterGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAccountLinkFilterGroupsWithDefaults() *PolicyAccountLinkFilterGroups {
	this := PolicyAccountLinkFilterGroups{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *PolicyAccountLinkFilterGroups) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAccountLinkFilterGroups) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *PolicyAccountLinkFilterGroups) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *PolicyAccountLinkFilterGroups) SetInclude(v []string) {
	o.Include = v
}

func (o PolicyAccountLinkFilterGroups) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyAccountLinkFilterGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyAccountLinkFilterGroups) UnmarshalJSON(data []byte) (err error) {
	varPolicyAccountLinkFilterGroups := _PolicyAccountLinkFilterGroups{}

	err = json.Unmarshal(data, &varPolicyAccountLinkFilterGroups)

	if err != nil {
		return err
	}

	*o = PolicyAccountLinkFilterGroups(varPolicyAccountLinkFilterGroups)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyAccountLinkFilterGroups struct {
	value *PolicyAccountLinkFilterGroups
	isSet bool
}

func (v NullablePolicyAccountLinkFilterGroups) Get() *PolicyAccountLinkFilterGroups {
	return v.value
}

func (v *NullablePolicyAccountLinkFilterGroups) Set(val *PolicyAccountLinkFilterGroups) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAccountLinkFilterGroups) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAccountLinkFilterGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAccountLinkFilterGroups(val *PolicyAccountLinkFilterGroups) *NullablePolicyAccountLinkFilterGroups {
	return &NullablePolicyAccountLinkFilterGroups{value: val, isSet: true}
}

func (v NullablePolicyAccountLinkFilterGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAccountLinkFilterGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
