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

// OktaSignOnPolicyRuleConditions struct for OktaSignOnPolicyRuleConditions
type OktaSignOnPolicyRuleConditions struct {
	AuthContext *PolicyRuleAuthContextCondition `json:"authContext,omitempty"`
	IdentityProvider *IdentityProviderPolicyRuleCondition `json:"identityProvider,omitempty"`
	Network *PolicyNetworkCondition `json:"network,omitempty"`
	People *PolicyPeopleCondition `json:"people,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaSignOnPolicyRuleConditions OktaSignOnPolicyRuleConditions

// NewOktaSignOnPolicyRuleConditions instantiates a new OktaSignOnPolicyRuleConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaSignOnPolicyRuleConditions() *OktaSignOnPolicyRuleConditions {
	this := OktaSignOnPolicyRuleConditions{}
	return &this
}

// NewOktaSignOnPolicyRuleConditionsWithDefaults instantiates a new OktaSignOnPolicyRuleConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaSignOnPolicyRuleConditionsWithDefaults() *OktaSignOnPolicyRuleConditions {
	this := OktaSignOnPolicyRuleConditions{}
	return &this
}

// GetAuthContext returns the AuthContext field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleConditions) GetAuthContext() PolicyRuleAuthContextCondition {
	if o == nil || o.AuthContext == nil {
		var ret PolicyRuleAuthContextCondition
		return ret
	}
	return *o.AuthContext
}

// GetAuthContextOk returns a tuple with the AuthContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleConditions) GetAuthContextOk() (*PolicyRuleAuthContextCondition, bool) {
	if o == nil || o.AuthContext == nil {
		return nil, false
	}
	return o.AuthContext, true
}

// HasAuthContext returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleConditions) HasAuthContext() bool {
	if o != nil && o.AuthContext != nil {
		return true
	}

	return false
}

// SetAuthContext gets a reference to the given PolicyRuleAuthContextCondition and assigns it to the AuthContext field.
func (o *OktaSignOnPolicyRuleConditions) SetAuthContext(v PolicyRuleAuthContextCondition) {
	o.AuthContext = &v
}

// GetIdentityProvider returns the IdentityProvider field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleConditions) GetIdentityProvider() IdentityProviderPolicyRuleCondition {
	if o == nil || o.IdentityProvider == nil {
		var ret IdentityProviderPolicyRuleCondition
		return ret
	}
	return *o.IdentityProvider
}

// GetIdentityProviderOk returns a tuple with the IdentityProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleConditions) GetIdentityProviderOk() (*IdentityProviderPolicyRuleCondition, bool) {
	if o == nil || o.IdentityProvider == nil {
		return nil, false
	}
	return o.IdentityProvider, true
}

// HasIdentityProvider returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleConditions) HasIdentityProvider() bool {
	if o != nil && o.IdentityProvider != nil {
		return true
	}

	return false
}

// SetIdentityProvider gets a reference to the given IdentityProviderPolicyRuleCondition and assigns it to the IdentityProvider field.
func (o *OktaSignOnPolicyRuleConditions) SetIdentityProvider(v IdentityProviderPolicyRuleCondition) {
	o.IdentityProvider = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleConditions) GetNetwork() PolicyNetworkCondition {
	if o == nil || o.Network == nil {
		var ret PolicyNetworkCondition
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleConditions) GetNetworkOk() (*PolicyNetworkCondition, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleConditions) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given PolicyNetworkCondition and assigns it to the Network field.
func (o *OktaSignOnPolicyRuleConditions) SetNetwork(v PolicyNetworkCondition) {
	o.Network = &v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleConditions) GetPeople() PolicyPeopleCondition {
	if o == nil || o.People == nil {
		var ret PolicyPeopleCondition
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleConditions) GetPeopleOk() (*PolicyPeopleCondition, bool) {
	if o == nil || o.People == nil {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleConditions) HasPeople() bool {
	if o != nil && o.People != nil {
		return true
	}

	return false
}

// SetPeople gets a reference to the given PolicyPeopleCondition and assigns it to the People field.
func (o *OktaSignOnPolicyRuleConditions) SetPeople(v PolicyPeopleCondition) {
	o.People = &v
}

func (o OktaSignOnPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthContext != nil {
		toSerialize["authContext"] = o.AuthContext
	}
	if o.IdentityProvider != nil {
		toSerialize["identityProvider"] = o.IdentityProvider
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.People != nil {
		toSerialize["people"] = o.People
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OktaSignOnPolicyRuleConditions) UnmarshalJSON(bytes []byte) (err error) {
	varOktaSignOnPolicyRuleConditions := _OktaSignOnPolicyRuleConditions{}

	err = json.Unmarshal(bytes, &varOktaSignOnPolicyRuleConditions)
	if err == nil {
		*o = OktaSignOnPolicyRuleConditions(varOktaSignOnPolicyRuleConditions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authContext")
		delete(additionalProperties, "identityProvider")
		delete(additionalProperties, "network")
		delete(additionalProperties, "people")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOktaSignOnPolicyRuleConditions struct {
	value *OktaSignOnPolicyRuleConditions
	isSet bool
}

func (v NullableOktaSignOnPolicyRuleConditions) Get() *OktaSignOnPolicyRuleConditions {
	return v.value
}

func (v *NullableOktaSignOnPolicyRuleConditions) Set(val *OktaSignOnPolicyRuleConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaSignOnPolicyRuleConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaSignOnPolicyRuleConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaSignOnPolicyRuleConditions(val *OktaSignOnPolicyRuleConditions) *NullableOktaSignOnPolicyRuleConditions {
	return &NullableOktaSignOnPolicyRuleConditions{value: val, isSet: true}
}

func (v NullableOktaSignOnPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaSignOnPolicyRuleConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

