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

// UserPolicyRuleCondition Specifies a set of Users to be included or excluded
type UserPolicyRuleCondition struct {
	// Users to be excluded
	Exclude []string `json:"exclude,omitempty"`
	Inactivity *InactivityPolicyRuleCondition `json:"inactivity,omitempty"`
	// Users to be included
	Include []string `json:"include,omitempty"`
	LifecycleExpiration *LifecycleExpirationPolicyRuleCondition `json:"lifecycleExpiration,omitempty"`
	PasswordExpiration *PasswordExpirationPolicyRuleCondition `json:"passwordExpiration,omitempty"`
	UserLifecycleAttribute *UserLifecycleAttributePolicyRuleCondition `json:"userLifecycleAttribute,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserPolicyRuleCondition UserPolicyRuleCondition

// NewUserPolicyRuleCondition instantiates a new UserPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPolicyRuleCondition() *UserPolicyRuleCondition {
	this := UserPolicyRuleCondition{}
	return &this
}

// NewUserPolicyRuleConditionWithDefaults instantiates a new UserPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPolicyRuleConditionWithDefaults() *UserPolicyRuleCondition {
	this := UserPolicyRuleCondition{}
	return &this
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *UserPolicyRuleCondition) GetExclude() []string {
	if o == nil || o.Exclude == nil {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPolicyRuleCondition) GetExcludeOk() ([]string, bool) {
	if o == nil || o.Exclude == nil {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *UserPolicyRuleCondition) HasExclude() bool {
	if o != nil && o.Exclude != nil {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *UserPolicyRuleCondition) SetExclude(v []string) {
	o.Exclude = v
}

// GetInactivity returns the Inactivity field value if set, zero value otherwise.
func (o *UserPolicyRuleCondition) GetInactivity() InactivityPolicyRuleCondition {
	if o == nil || o.Inactivity == nil {
		var ret InactivityPolicyRuleCondition
		return ret
	}
	return *o.Inactivity
}

// GetInactivityOk returns a tuple with the Inactivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPolicyRuleCondition) GetInactivityOk() (*InactivityPolicyRuleCondition, bool) {
	if o == nil || o.Inactivity == nil {
		return nil, false
	}
	return o.Inactivity, true
}

// HasInactivity returns a boolean if a field has been set.
func (o *UserPolicyRuleCondition) HasInactivity() bool {
	if o != nil && o.Inactivity != nil {
		return true
	}

	return false
}

// SetInactivity gets a reference to the given InactivityPolicyRuleCondition and assigns it to the Inactivity field.
func (o *UserPolicyRuleCondition) SetInactivity(v InactivityPolicyRuleCondition) {
	o.Inactivity = &v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *UserPolicyRuleCondition) GetInclude() []string {
	if o == nil || o.Include == nil {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPolicyRuleCondition) GetIncludeOk() ([]string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *UserPolicyRuleCondition) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *UserPolicyRuleCondition) SetInclude(v []string) {
	o.Include = v
}

// GetLifecycleExpiration returns the LifecycleExpiration field value if set, zero value otherwise.
func (o *UserPolicyRuleCondition) GetLifecycleExpiration() LifecycleExpirationPolicyRuleCondition {
	if o == nil || o.LifecycleExpiration == nil {
		var ret LifecycleExpirationPolicyRuleCondition
		return ret
	}
	return *o.LifecycleExpiration
}

// GetLifecycleExpirationOk returns a tuple with the LifecycleExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPolicyRuleCondition) GetLifecycleExpirationOk() (*LifecycleExpirationPolicyRuleCondition, bool) {
	if o == nil || o.LifecycleExpiration == nil {
		return nil, false
	}
	return o.LifecycleExpiration, true
}

// HasLifecycleExpiration returns a boolean if a field has been set.
func (o *UserPolicyRuleCondition) HasLifecycleExpiration() bool {
	if o != nil && o.LifecycleExpiration != nil {
		return true
	}

	return false
}

// SetLifecycleExpiration gets a reference to the given LifecycleExpirationPolicyRuleCondition and assigns it to the LifecycleExpiration field.
func (o *UserPolicyRuleCondition) SetLifecycleExpiration(v LifecycleExpirationPolicyRuleCondition) {
	o.LifecycleExpiration = &v
}

// GetPasswordExpiration returns the PasswordExpiration field value if set, zero value otherwise.
func (o *UserPolicyRuleCondition) GetPasswordExpiration() PasswordExpirationPolicyRuleCondition {
	if o == nil || o.PasswordExpiration == nil {
		var ret PasswordExpirationPolicyRuleCondition
		return ret
	}
	return *o.PasswordExpiration
}

// GetPasswordExpirationOk returns a tuple with the PasswordExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPolicyRuleCondition) GetPasswordExpirationOk() (*PasswordExpirationPolicyRuleCondition, bool) {
	if o == nil || o.PasswordExpiration == nil {
		return nil, false
	}
	return o.PasswordExpiration, true
}

// HasPasswordExpiration returns a boolean if a field has been set.
func (o *UserPolicyRuleCondition) HasPasswordExpiration() bool {
	if o != nil && o.PasswordExpiration != nil {
		return true
	}

	return false
}

// SetPasswordExpiration gets a reference to the given PasswordExpirationPolicyRuleCondition and assigns it to the PasswordExpiration field.
func (o *UserPolicyRuleCondition) SetPasswordExpiration(v PasswordExpirationPolicyRuleCondition) {
	o.PasswordExpiration = &v
}

// GetUserLifecycleAttribute returns the UserLifecycleAttribute field value if set, zero value otherwise.
func (o *UserPolicyRuleCondition) GetUserLifecycleAttribute() UserLifecycleAttributePolicyRuleCondition {
	if o == nil || o.UserLifecycleAttribute == nil {
		var ret UserLifecycleAttributePolicyRuleCondition
		return ret
	}
	return *o.UserLifecycleAttribute
}

// GetUserLifecycleAttributeOk returns a tuple with the UserLifecycleAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPolicyRuleCondition) GetUserLifecycleAttributeOk() (*UserLifecycleAttributePolicyRuleCondition, bool) {
	if o == nil || o.UserLifecycleAttribute == nil {
		return nil, false
	}
	return o.UserLifecycleAttribute, true
}

// HasUserLifecycleAttribute returns a boolean if a field has been set.
func (o *UserPolicyRuleCondition) HasUserLifecycleAttribute() bool {
	if o != nil && o.UserLifecycleAttribute != nil {
		return true
	}

	return false
}

// SetUserLifecycleAttribute gets a reference to the given UserLifecycleAttributePolicyRuleCondition and assigns it to the UserLifecycleAttribute field.
func (o *UserPolicyRuleCondition) SetUserLifecycleAttribute(v UserLifecycleAttributePolicyRuleCondition) {
	o.UserLifecycleAttribute = &v
}

func (o UserPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Exclude != nil {
		toSerialize["exclude"] = o.Exclude
	}
	if o.Inactivity != nil {
		toSerialize["inactivity"] = o.Inactivity
	}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}
	if o.LifecycleExpiration != nil {
		toSerialize["lifecycleExpiration"] = o.LifecycleExpiration
	}
	if o.PasswordExpiration != nil {
		toSerialize["passwordExpiration"] = o.PasswordExpiration
	}
	if o.UserLifecycleAttribute != nil {
		toSerialize["userLifecycleAttribute"] = o.UserLifecycleAttribute
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserPolicyRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varUserPolicyRuleCondition := _UserPolicyRuleCondition{}

	err = json.Unmarshal(bytes, &varUserPolicyRuleCondition)
	if err == nil {
		*o = UserPolicyRuleCondition(varUserPolicyRuleCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "exclude")
		delete(additionalProperties, "inactivity")
		delete(additionalProperties, "include")
		delete(additionalProperties, "lifecycleExpiration")
		delete(additionalProperties, "passwordExpiration")
		delete(additionalProperties, "userLifecycleAttribute")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserPolicyRuleCondition struct {
	value *UserPolicyRuleCondition
	isSet bool
}

func (v NullableUserPolicyRuleCondition) Get() *UserPolicyRuleCondition {
	return v.value
}

func (v *NullableUserPolicyRuleCondition) Set(val *UserPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPolicyRuleCondition(val *UserPolicyRuleCondition) *NullableUserPolicyRuleCondition {
	return &NullableUserPolicyRuleCondition{value: val, isSet: true}
}

func (v NullableUserPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

