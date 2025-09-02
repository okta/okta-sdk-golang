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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// PolicyAccountLinkFilterUsers Filters on which users are available for account linking
type PolicyAccountLinkFilterUsers struct {
	// Specifies the blocklist of user identifiers to exclude from account linking
	Exclude []string `json:"exclude,omitempty"`
	// Specifies whether admin users should be excluded from account linking
	ExcludeAdmins *bool `json:"excludeAdmins,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAccountLinkFilterUsers PolicyAccountLinkFilterUsers

// NewPolicyAccountLinkFilterUsers instantiates a new PolicyAccountLinkFilterUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAccountLinkFilterUsers() *PolicyAccountLinkFilterUsers {
	this := PolicyAccountLinkFilterUsers{}
	var excludeAdmins bool = false
	this.ExcludeAdmins = &excludeAdmins
	return &this
}

// NewPolicyAccountLinkFilterUsersWithDefaults instantiates a new PolicyAccountLinkFilterUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAccountLinkFilterUsersWithDefaults() *PolicyAccountLinkFilterUsers {
	this := PolicyAccountLinkFilterUsers{}
	var excludeAdmins bool = false
	this.ExcludeAdmins = &excludeAdmins
	return &this
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *PolicyAccountLinkFilterUsers) GetExclude() []string {
	if o == nil || o.Exclude == nil {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAccountLinkFilterUsers) GetExcludeOk() ([]string, bool) {
	if o == nil || o.Exclude == nil {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *PolicyAccountLinkFilterUsers) HasExclude() bool {
	if o != nil && o.Exclude != nil {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *PolicyAccountLinkFilterUsers) SetExclude(v []string) {
	o.Exclude = v
}

// GetExcludeAdmins returns the ExcludeAdmins field value if set, zero value otherwise.
func (o *PolicyAccountLinkFilterUsers) GetExcludeAdmins() bool {
	if o == nil || o.ExcludeAdmins == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeAdmins
}

// GetExcludeAdminsOk returns a tuple with the ExcludeAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAccountLinkFilterUsers) GetExcludeAdminsOk() (*bool, bool) {
	if o == nil || o.ExcludeAdmins == nil {
		return nil, false
	}
	return o.ExcludeAdmins, true
}

// HasExcludeAdmins returns a boolean if a field has been set.
func (o *PolicyAccountLinkFilterUsers) HasExcludeAdmins() bool {
	if o != nil && o.ExcludeAdmins != nil {
		return true
	}

	return false
}

// SetExcludeAdmins gets a reference to the given bool and assigns it to the ExcludeAdmins field.
func (o *PolicyAccountLinkFilterUsers) SetExcludeAdmins(v bool) {
	o.ExcludeAdmins = &v
}

func (o PolicyAccountLinkFilterUsers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Exclude != nil {
		toSerialize["exclude"] = o.Exclude
	}
	if o.ExcludeAdmins != nil {
		toSerialize["excludeAdmins"] = o.ExcludeAdmins
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyAccountLinkFilterUsers) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyAccountLinkFilterUsers := _PolicyAccountLinkFilterUsers{}

	err = json.Unmarshal(bytes, &varPolicyAccountLinkFilterUsers)
	if err == nil {
		*o = PolicyAccountLinkFilterUsers(varPolicyAccountLinkFilterUsers)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "exclude")
		delete(additionalProperties, "excludeAdmins")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePolicyAccountLinkFilterUsers struct {
	value *PolicyAccountLinkFilterUsers
	isSet bool
}

func (v NullablePolicyAccountLinkFilterUsers) Get() *PolicyAccountLinkFilterUsers {
	return v.value
}

func (v *NullablePolicyAccountLinkFilterUsers) Set(val *PolicyAccountLinkFilterUsers) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAccountLinkFilterUsers) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAccountLinkFilterUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAccountLinkFilterUsers(val *PolicyAccountLinkFilterUsers) *NullablePolicyAccountLinkFilterUsers {
	return &NullablePolicyAccountLinkFilterUsers{value: val, isSet: true}
}

func (v NullablePolicyAccountLinkFilterUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAccountLinkFilterUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

