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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the ProvisioningConnectionTokenRequestProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningConnectionTokenRequestProfile{}

// ProvisioningConnectionTokenRequestProfile struct for ProvisioningConnectionTokenRequestProfile
type ProvisioningConnectionTokenRequestProfile struct {
	// A token is used to authenticate with the app. This property is only returned for the `TOKEN` authentication scheme.
	AuthScheme string `json:"authScheme"`
	// Token used to authenticate with the app
	Token                *string `json:"token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningConnectionTokenRequestProfile ProvisioningConnectionTokenRequestProfile

// NewProvisioningConnectionTokenRequestProfile instantiates a new ProvisioningConnectionTokenRequestProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningConnectionTokenRequestProfile(authScheme string) *ProvisioningConnectionTokenRequestProfile {
	this := ProvisioningConnectionTokenRequestProfile{}
	this.AuthScheme = authScheme
	return &this
}

// NewProvisioningConnectionTokenRequestProfileWithDefaults instantiates a new ProvisioningConnectionTokenRequestProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningConnectionTokenRequestProfileWithDefaults() *ProvisioningConnectionTokenRequestProfile {
	this := ProvisioningConnectionTokenRequestProfile{}
	return &this
}

// GetAuthScheme returns the AuthScheme field value
func (o *ProvisioningConnectionTokenRequestProfile) GetAuthScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthScheme
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionTokenRequestProfile) GetAuthSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthScheme, true
}

// SetAuthScheme sets field value
func (o *ProvisioningConnectionTokenRequestProfile) SetAuthScheme(v string) {
	o.AuthScheme = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ProvisioningConnectionTokenRequestProfile) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionTokenRequestProfile) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ProvisioningConnectionTokenRequestProfile) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ProvisioningConnectionTokenRequestProfile) SetToken(v string) {
	o.Token = &v
}

func (o ProvisioningConnectionTokenRequestProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningConnectionTokenRequestProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authScheme"] = o.AuthScheme
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningConnectionTokenRequestProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authScheme",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProvisioningConnectionTokenRequestProfile := _ProvisioningConnectionTokenRequestProfile{}

	err = json.Unmarshal(data, &varProvisioningConnectionTokenRequestProfile)

	if err != nil {
		return err
	}

	*o = ProvisioningConnectionTokenRequestProfile(varProvisioningConnectionTokenRequestProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authScheme")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningConnectionTokenRequestProfile struct {
	value *ProvisioningConnectionTokenRequestProfile
	isSet bool
}

func (v NullableProvisioningConnectionTokenRequestProfile) Get() *ProvisioningConnectionTokenRequestProfile {
	return v.value
}

func (v *NullableProvisioningConnectionTokenRequestProfile) Set(val *ProvisioningConnectionTokenRequestProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConnectionTokenRequestProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConnectionTokenRequestProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConnectionTokenRequestProfile(val *ProvisioningConnectionTokenRequestProfile) *NullableProvisioningConnectionTokenRequestProfile {
	return &NullableProvisioningConnectionTokenRequestProfile{value: val, isSet: true}
}

func (v NullableProvisioningConnectionTokenRequestProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConnectionTokenRequestProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
