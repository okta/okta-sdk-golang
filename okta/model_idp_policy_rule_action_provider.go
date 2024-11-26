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

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
)

// IdpPolicyRuleActionProvider struct for IdpPolicyRuleActionProvider
type IdpPolicyRuleActionProvider struct {
	// IdP types of `OKTA`, `AgentlessDSSO`, and `IWA` don't require an ID.
	Id *string `json:"id,omitempty"`
	// Provider `name` in Okta. Optional. Supported in `IDENTITY ENGINE`.
	Name                 *string `json:"name,omitempty"`
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdpPolicyRuleActionProvider IdpPolicyRuleActionProvider

// NewIdpPolicyRuleActionProvider instantiates a new IdpPolicyRuleActionProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpPolicyRuleActionProvider() *IdpPolicyRuleActionProvider {
	this := IdpPolicyRuleActionProvider{}
	return &this
}

// NewIdpPolicyRuleActionProviderWithDefaults instantiates a new IdpPolicyRuleActionProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpPolicyRuleActionProviderWithDefaults() *IdpPolicyRuleActionProvider {
	this := IdpPolicyRuleActionProvider{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdpPolicyRuleActionProvider) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPolicyRuleActionProvider) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdpPolicyRuleActionProvider) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdpPolicyRuleActionProvider) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdpPolicyRuleActionProvider) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPolicyRuleActionProvider) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdpPolicyRuleActionProvider) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdpPolicyRuleActionProvider) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdpPolicyRuleActionProvider) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPolicyRuleActionProvider) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdpPolicyRuleActionProvider) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IdpPolicyRuleActionProvider) SetType(v string) {
	o.Type = &v
}

func (o IdpPolicyRuleActionProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdpPolicyRuleActionProvider) UnmarshalJSON(bytes []byte) (err error) {
	varIdpPolicyRuleActionProvider := _IdpPolicyRuleActionProvider{}

	err = json.Unmarshal(bytes, &varIdpPolicyRuleActionProvider)
	if err == nil {
		*o = IdpPolicyRuleActionProvider(varIdpPolicyRuleActionProvider)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdpPolicyRuleActionProvider struct {
	value *IdpPolicyRuleActionProvider
	isSet bool
}

func (v NullableIdpPolicyRuleActionProvider) Get() *IdpPolicyRuleActionProvider {
	return v.value
}

func (v *NullableIdpPolicyRuleActionProvider) Set(val *IdpPolicyRuleActionProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpPolicyRuleActionProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpPolicyRuleActionProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpPolicyRuleActionProvider(val *IdpPolicyRuleActionProvider) *NullableIdpPolicyRuleActionProvider {
	return &NullableIdpPolicyRuleActionProvider{value: val, isSet: true}
}

func (v NullableIdpPolicyRuleActionProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpPolicyRuleActionProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
