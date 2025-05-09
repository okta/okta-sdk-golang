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

// ProvisioningConnectionTokenRequestProfile struct for ProvisioningConnectionTokenRequestProfile
type ProvisioningConnectionTokenRequestProfile struct {
	// A token is used to authenticate with the app. This property is only returned for the `TOKEN` authentication scheme.
	AuthScheme string `json:"authScheme"`
	// Token used to authenticate with the app
	Token *string `json:"token,omitempty"`
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
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionTokenRequestProfile) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ProvisioningConnectionTokenRequestProfile) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ProvisioningConnectionTokenRequestProfile) SetToken(v string) {
	o.Token = &v
}

func (o ProvisioningConnectionTokenRequestProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authScheme"] = o.AuthScheme
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProvisioningConnectionTokenRequestProfile) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningConnectionTokenRequestProfile := _ProvisioningConnectionTokenRequestProfile{}

	err = json.Unmarshal(bytes, &varProvisioningConnectionTokenRequestProfile)
	if err == nil {
		*o = ProvisioningConnectionTokenRequestProfile(varProvisioningConnectionTokenRequestProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authScheme")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

