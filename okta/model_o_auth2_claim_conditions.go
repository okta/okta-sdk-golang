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

// OAuth2ClaimConditions Specifies the scopes for the Claim
type OAuth2ClaimConditions struct {
	Scopes []string `json:"scopes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClaimConditions OAuth2ClaimConditions

// NewOAuth2ClaimConditions instantiates a new OAuth2ClaimConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClaimConditions() *OAuth2ClaimConditions {
	this := OAuth2ClaimConditions{}
	return &this
}

// NewOAuth2ClaimConditionsWithDefaults instantiates a new OAuth2ClaimConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClaimConditionsWithDefaults() *OAuth2ClaimConditions {
	this := OAuth2ClaimConditions{}
	return &this
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *OAuth2ClaimConditions) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClaimConditions) GetScopesOk() ([]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *OAuth2ClaimConditions) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *OAuth2ClaimConditions) SetScopes(v []string) {
	o.Scopes = v
}

func (o OAuth2ClaimConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2ClaimConditions) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2ClaimConditions := _OAuth2ClaimConditions{}

	err = json.Unmarshal(bytes, &varOAuth2ClaimConditions)
	if err == nil {
		*o = OAuth2ClaimConditions(varOAuth2ClaimConditions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "scopes")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOAuth2ClaimConditions struct {
	value *OAuth2ClaimConditions
	isSet bool
}

func (v NullableOAuth2ClaimConditions) Get() *OAuth2ClaimConditions {
	return v.value
}

func (v *NullableOAuth2ClaimConditions) Set(val *OAuth2ClaimConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClaimConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClaimConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClaimConditions(val *OAuth2ClaimConditions) *NullableOAuth2ClaimConditions {
	return &NullableOAuth2ClaimConditions{value: val, isSet: true}
}

func (v NullableOAuth2ClaimConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClaimConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

