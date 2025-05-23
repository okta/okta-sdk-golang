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

// IdpPolicyRuleActionIdp struct for IdpPolicyRuleActionIdp
type IdpPolicyRuleActionIdp struct {
	// List of configured Identity Providers that a given Rule can route to. Ability to define multiple providers is a part of the Okta Identity Engine. This allows users to choose a Provider when they sign in. Contact support for information on the Identity Engine.
	Providers []IdpPolicyRuleActionProvider `json:"providers,omitempty"`
	IdpSelectionType *string `json:"idpSelectionType,omitempty"`
	// Required if `idpSelectionType` is set to `DYNAMIC`
	MatchCriteria []IdpPolicyRuleActionMatchCriteria `json:"matchCriteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdpPolicyRuleActionIdp IdpPolicyRuleActionIdp

// NewIdpPolicyRuleActionIdp instantiates a new IdpPolicyRuleActionIdp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpPolicyRuleActionIdp() *IdpPolicyRuleActionIdp {
	this := IdpPolicyRuleActionIdp{}
	return &this
}

// NewIdpPolicyRuleActionIdpWithDefaults instantiates a new IdpPolicyRuleActionIdp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpPolicyRuleActionIdpWithDefaults() *IdpPolicyRuleActionIdp {
	this := IdpPolicyRuleActionIdp{}
	return &this
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *IdpPolicyRuleActionIdp) GetProviders() []IdpPolicyRuleActionProvider {
	if o == nil || o.Providers == nil {
		var ret []IdpPolicyRuleActionProvider
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPolicyRuleActionIdp) GetProvidersOk() ([]IdpPolicyRuleActionProvider, bool) {
	if o == nil || o.Providers == nil {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *IdpPolicyRuleActionIdp) HasProviders() bool {
	if o != nil && o.Providers != nil {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []IdpPolicyRuleActionProvider and assigns it to the Providers field.
func (o *IdpPolicyRuleActionIdp) SetProviders(v []IdpPolicyRuleActionProvider) {
	o.Providers = v
}

// GetIdpSelectionType returns the IdpSelectionType field value if set, zero value otherwise.
func (o *IdpPolicyRuleActionIdp) GetIdpSelectionType() string {
	if o == nil || o.IdpSelectionType == nil {
		var ret string
		return ret
	}
	return *o.IdpSelectionType
}

// GetIdpSelectionTypeOk returns a tuple with the IdpSelectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPolicyRuleActionIdp) GetIdpSelectionTypeOk() (*string, bool) {
	if o == nil || o.IdpSelectionType == nil {
		return nil, false
	}
	return o.IdpSelectionType, true
}

// HasIdpSelectionType returns a boolean if a field has been set.
func (o *IdpPolicyRuleActionIdp) HasIdpSelectionType() bool {
	if o != nil && o.IdpSelectionType != nil {
		return true
	}

	return false
}

// SetIdpSelectionType gets a reference to the given string and assigns it to the IdpSelectionType field.
func (o *IdpPolicyRuleActionIdp) SetIdpSelectionType(v string) {
	o.IdpSelectionType = &v
}

// GetMatchCriteria returns the MatchCriteria field value if set, zero value otherwise.
func (o *IdpPolicyRuleActionIdp) GetMatchCriteria() []IdpPolicyRuleActionMatchCriteria {
	if o == nil || o.MatchCriteria == nil {
		var ret []IdpPolicyRuleActionMatchCriteria
		return ret
	}
	return o.MatchCriteria
}

// GetMatchCriteriaOk returns a tuple with the MatchCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPolicyRuleActionIdp) GetMatchCriteriaOk() ([]IdpPolicyRuleActionMatchCriteria, bool) {
	if o == nil || o.MatchCriteria == nil {
		return nil, false
	}
	return o.MatchCriteria, true
}

// HasMatchCriteria returns a boolean if a field has been set.
func (o *IdpPolicyRuleActionIdp) HasMatchCriteria() bool {
	if o != nil && o.MatchCriteria != nil {
		return true
	}

	return false
}

// SetMatchCriteria gets a reference to the given []IdpPolicyRuleActionMatchCriteria and assigns it to the MatchCriteria field.
func (o *IdpPolicyRuleActionIdp) SetMatchCriteria(v []IdpPolicyRuleActionMatchCriteria) {
	o.MatchCriteria = v
}

func (o IdpPolicyRuleActionIdp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Providers != nil {
		toSerialize["providers"] = o.Providers
	}
	if o.IdpSelectionType != nil {
		toSerialize["idpSelectionType"] = o.IdpSelectionType
	}
	if o.MatchCriteria != nil {
		toSerialize["matchCriteria"] = o.MatchCriteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdpPolicyRuleActionIdp) UnmarshalJSON(bytes []byte) (err error) {
	varIdpPolicyRuleActionIdp := _IdpPolicyRuleActionIdp{}

	err = json.Unmarshal(bytes, &varIdpPolicyRuleActionIdp)
	if err == nil {
		*o = IdpPolicyRuleActionIdp(varIdpPolicyRuleActionIdp)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "providers")
		delete(additionalProperties, "idpSelectionType")
		delete(additionalProperties, "matchCriteria")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdpPolicyRuleActionIdp struct {
	value *IdpPolicyRuleActionIdp
	isSet bool
}

func (v NullableIdpPolicyRuleActionIdp) Get() *IdpPolicyRuleActionIdp {
	return v.value
}

func (v *NullableIdpPolicyRuleActionIdp) Set(val *IdpPolicyRuleActionIdp) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpPolicyRuleActionIdp) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpPolicyRuleActionIdp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpPolicyRuleActionIdp(val *IdpPolicyRuleActionIdp) *NullableIdpPolicyRuleActionIdp {
	return &NullableIdpPolicyRuleActionIdp{value: val, isSet: true}
}

func (v NullableIdpPolicyRuleActionIdp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpPolicyRuleActionIdp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

