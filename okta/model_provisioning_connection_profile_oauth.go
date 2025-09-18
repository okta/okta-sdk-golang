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

// checks if the ProvisioningConnectionProfileOauth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningConnectionProfileOauth{}

// ProvisioningConnectionProfileOauth The app provisioning connection profile used to configure the method of authentication and the credentials. Currently, token-based and OAuth 2.0-based authentication are supported.
type ProvisioningConnectionProfileOauth struct {
	// OAuth 2.0 is used to authenticate with the app.
	AuthScheme           string  `json:"authScheme"`
	ClientId             *string `json:"clientId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningConnectionProfileOauth ProvisioningConnectionProfileOauth

// NewProvisioningConnectionProfileOauth instantiates a new ProvisioningConnectionProfileOauth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningConnectionProfileOauth(authScheme string) *ProvisioningConnectionProfileOauth {
	this := ProvisioningConnectionProfileOauth{}
	this.AuthScheme = authScheme
	return &this
}

// NewProvisioningConnectionProfileOauthWithDefaults instantiates a new ProvisioningConnectionProfileOauth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningConnectionProfileOauthWithDefaults() *ProvisioningConnectionProfileOauth {
	this := ProvisioningConnectionProfileOauth{}
	return &this
}

// GetAuthScheme returns the AuthScheme field value
func (o *ProvisioningConnectionProfileOauth) GetAuthScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthScheme
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionProfileOauth) GetAuthSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthScheme, true
}

// SetAuthScheme sets field value
func (o *ProvisioningConnectionProfileOauth) SetAuthScheme(v string) {
	o.AuthScheme = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ProvisioningConnectionProfileOauth) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionProfileOauth) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProvisioningConnectionProfileOauth) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ProvisioningConnectionProfileOauth) SetClientId(v string) {
	o.ClientId = &v
}

func (o ProvisioningConnectionProfileOauth) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningConnectionProfileOauth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authScheme"] = o.AuthScheme
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningConnectionProfileOauth) UnmarshalJSON(data []byte) (err error) {
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

	varProvisioningConnectionProfileOauth := _ProvisioningConnectionProfileOauth{}

	err = json.Unmarshal(data, &varProvisioningConnectionProfileOauth)

	if err != nil {
		return err
	}

	*o = ProvisioningConnectionProfileOauth(varProvisioningConnectionProfileOauth)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authScheme")
		delete(additionalProperties, "clientId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningConnectionProfileOauth struct {
	value *ProvisioningConnectionProfileOauth
	isSet bool
}

func (v NullableProvisioningConnectionProfileOauth) Get() *ProvisioningConnectionProfileOauth {
	return v.value
}

func (v *NullableProvisioningConnectionProfileOauth) Set(val *ProvisioningConnectionProfileOauth) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConnectionProfileOauth) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConnectionProfileOauth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConnectionProfileOauth(val *ProvisioningConnectionProfileOauth) *NullableProvisioningConnectionProfileOauth {
	return &NullableProvisioningConnectionProfileOauth{value: val, isSet: true}
}

func (v NullableProvisioningConnectionProfileOauth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConnectionProfileOauth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
