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

// ProvisioningConnectionResponseProfile struct for ProvisioningConnectionResponseProfile
type ProvisioningConnectionResponseProfile struct {
	// Defines the method of authentication
	AuthScheme string `json:"authScheme"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningConnectionResponseProfile ProvisioningConnectionResponseProfile

// NewProvisioningConnectionResponseProfile instantiates a new ProvisioningConnectionResponseProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningConnectionResponseProfile(authScheme string) *ProvisioningConnectionResponseProfile {
	this := ProvisioningConnectionResponseProfile{}
	this.AuthScheme = authScheme
	return &this
}

// NewProvisioningConnectionResponseProfileWithDefaults instantiates a new ProvisioningConnectionResponseProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningConnectionResponseProfileWithDefaults() *ProvisioningConnectionResponseProfile {
	this := ProvisioningConnectionResponseProfile{}
	return &this
}

// GetAuthScheme returns the AuthScheme field value
func (o *ProvisioningConnectionResponseProfile) GetAuthScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthScheme
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionResponseProfile) GetAuthSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthScheme, true
}

// SetAuthScheme sets field value
func (o *ProvisioningConnectionResponseProfile) SetAuthScheme(v string) {
	o.AuthScheme = v
}

func (o ProvisioningConnectionResponseProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authScheme"] = o.AuthScheme
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProvisioningConnectionResponseProfile) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningConnectionResponseProfile := _ProvisioningConnectionResponseProfile{}

	err = json.Unmarshal(bytes, &varProvisioningConnectionResponseProfile)
	if err == nil {
		*o = ProvisioningConnectionResponseProfile(varProvisioningConnectionResponseProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authScheme")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProvisioningConnectionResponseProfile struct {
	value *ProvisioningConnectionResponseProfile
	isSet bool
}

func (v NullableProvisioningConnectionResponseProfile) Get() *ProvisioningConnectionResponseProfile {
	return v.value
}

func (v *NullableProvisioningConnectionResponseProfile) Set(val *ProvisioningConnectionResponseProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConnectionResponseProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConnectionResponseProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConnectionResponseProfile(val *ProvisioningConnectionResponseProfile) *NullableProvisioningConnectionResponseProfile {
	return &NullableProvisioningConnectionResponseProfile{value: val, isSet: true}
}

func (v NullableProvisioningConnectionResponseProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConnectionResponseProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

