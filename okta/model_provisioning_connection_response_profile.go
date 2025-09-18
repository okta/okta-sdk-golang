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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the ProvisioningConnectionResponseProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningConnectionResponseProfile{}

// ProvisioningConnectionResponseProfile struct for ProvisioningConnectionResponseProfile
type ProvisioningConnectionResponseProfile struct {
	// Defines the method of authentication
	AuthScheme           string                                   `json:"authScheme"`
	Signing              *Org2OrgProvisioningOAuthSigningSettings `json:"signing,omitempty"`
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

// GetSigning returns the Signing field value if set, zero value otherwise.
func (o *ProvisioningConnectionResponseProfile) GetSigning() Org2OrgProvisioningOAuthSigningSettings {
	if o == nil || IsNil(o.Signing) {
		var ret Org2OrgProvisioningOAuthSigningSettings
		return ret
	}
	return *o.Signing
}

// GetSigningOk returns a tuple with the Signing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionResponseProfile) GetSigningOk() (*Org2OrgProvisioningOAuthSigningSettings, bool) {
	if o == nil || IsNil(o.Signing) {
		return nil, false
	}
	return o.Signing, true
}

// HasSigning returns a boolean if a field has been set.
func (o *ProvisioningConnectionResponseProfile) HasSigning() bool {
	if o != nil && !IsNil(o.Signing) {
		return true
	}

	return false
}

// SetSigning gets a reference to the given Org2OrgProvisioningOAuthSigningSettings and assigns it to the Signing field.
func (o *ProvisioningConnectionResponseProfile) SetSigning(v Org2OrgProvisioningOAuthSigningSettings) {
	o.Signing = &v
}

func (o ProvisioningConnectionResponseProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningConnectionResponseProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authScheme"] = o.AuthScheme
	if !IsNil(o.Signing) {
		toSerialize["signing"] = o.Signing
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningConnectionResponseProfile) UnmarshalJSON(data []byte) (err error) {
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

	varProvisioningConnectionResponseProfile := _ProvisioningConnectionResponseProfile{}

	err = json.Unmarshal(data, &varProvisioningConnectionResponseProfile)

	if err != nil {
		return err
	}

	*o = ProvisioningConnectionResponseProfile(varProvisioningConnectionResponseProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authScheme")
		delete(additionalProperties, "signing")
		o.AdditionalProperties = additionalProperties
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
