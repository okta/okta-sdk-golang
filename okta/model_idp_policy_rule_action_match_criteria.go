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

// IdpPolicyRuleActionMatchCriteria struct for IdpPolicyRuleActionMatchCriteria
type IdpPolicyRuleActionMatchCriteria struct {
	// The IdP property that the evaluated string should match to
	PropertyName *string `json:"propertyName,omitempty"`
	// You can provide an Okta Expression Language expression with the Login Context that's evaluated with the IdP. For example, the value `login.identifier` refers to the user's username. If the user is signing in with the username `john.doe@mycompany.com`, the expression `login.identifier.substringAfter(@))` is evaluated to the domain name of the user, for example: `mycompany.com`. 
	ProviderExpression *string `json:"providerExpression,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdpPolicyRuleActionMatchCriteria IdpPolicyRuleActionMatchCriteria

// NewIdpPolicyRuleActionMatchCriteria instantiates a new IdpPolicyRuleActionMatchCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpPolicyRuleActionMatchCriteria() *IdpPolicyRuleActionMatchCriteria {
	this := IdpPolicyRuleActionMatchCriteria{}
	return &this
}

// NewIdpPolicyRuleActionMatchCriteriaWithDefaults instantiates a new IdpPolicyRuleActionMatchCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpPolicyRuleActionMatchCriteriaWithDefaults() *IdpPolicyRuleActionMatchCriteria {
	this := IdpPolicyRuleActionMatchCriteria{}
	return &this
}

// GetPropertyName returns the PropertyName field value if set, zero value otherwise.
func (o *IdpPolicyRuleActionMatchCriteria) GetPropertyName() string {
	if o == nil || o.PropertyName == nil {
		var ret string
		return ret
	}
	return *o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPolicyRuleActionMatchCriteria) GetPropertyNameOk() (*string, bool) {
	if o == nil || o.PropertyName == nil {
		return nil, false
	}
	return o.PropertyName, true
}

// HasPropertyName returns a boolean if a field has been set.
func (o *IdpPolicyRuleActionMatchCriteria) HasPropertyName() bool {
	if o != nil && o.PropertyName != nil {
		return true
	}

	return false
}

// SetPropertyName gets a reference to the given string and assigns it to the PropertyName field.
func (o *IdpPolicyRuleActionMatchCriteria) SetPropertyName(v string) {
	o.PropertyName = &v
}

// GetProviderExpression returns the ProviderExpression field value if set, zero value otherwise.
func (o *IdpPolicyRuleActionMatchCriteria) GetProviderExpression() string {
	if o == nil || o.ProviderExpression == nil {
		var ret string
		return ret
	}
	return *o.ProviderExpression
}

// GetProviderExpressionOk returns a tuple with the ProviderExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPolicyRuleActionMatchCriteria) GetProviderExpressionOk() (*string, bool) {
	if o == nil || o.ProviderExpression == nil {
		return nil, false
	}
	return o.ProviderExpression, true
}

// HasProviderExpression returns a boolean if a field has been set.
func (o *IdpPolicyRuleActionMatchCriteria) HasProviderExpression() bool {
	if o != nil && o.ProviderExpression != nil {
		return true
	}

	return false
}

// SetProviderExpression gets a reference to the given string and assigns it to the ProviderExpression field.
func (o *IdpPolicyRuleActionMatchCriteria) SetProviderExpression(v string) {
	o.ProviderExpression = &v
}

func (o IdpPolicyRuleActionMatchCriteria) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PropertyName != nil {
		toSerialize["propertyName"] = o.PropertyName
	}
	if o.ProviderExpression != nil {
		toSerialize["providerExpression"] = o.ProviderExpression
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdpPolicyRuleActionMatchCriteria) UnmarshalJSON(bytes []byte) (err error) {
	varIdpPolicyRuleActionMatchCriteria := _IdpPolicyRuleActionMatchCriteria{}

	err = json.Unmarshal(bytes, &varIdpPolicyRuleActionMatchCriteria)
	if err == nil {
		*o = IdpPolicyRuleActionMatchCriteria(varIdpPolicyRuleActionMatchCriteria)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "propertyName")
		delete(additionalProperties, "providerExpression")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdpPolicyRuleActionMatchCriteria struct {
	value *IdpPolicyRuleActionMatchCriteria
	isSet bool
}

func (v NullableIdpPolicyRuleActionMatchCriteria) Get() *IdpPolicyRuleActionMatchCriteria {
	return v.value
}

func (v *NullableIdpPolicyRuleActionMatchCriteria) Set(val *IdpPolicyRuleActionMatchCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpPolicyRuleActionMatchCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpPolicyRuleActionMatchCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpPolicyRuleActionMatchCriteria(val *IdpPolicyRuleActionMatchCriteria) *NullableIdpPolicyRuleActionMatchCriteria {
	return &NullableIdpPolicyRuleActionMatchCriteria{value: val, isSet: true}
}

func (v NullableIdpPolicyRuleActionMatchCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpPolicyRuleActionMatchCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

