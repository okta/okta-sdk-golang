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

// IdentityProviderCredentialsClient struct for IdentityProviderCredentialsClient
type IdentityProviderCredentialsClient struct {
	ClientId *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`
	// Require Proof Key for Code Exchange (PKCE) for additional verification
	PkceRequired *bool `json:"pkce_required,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProviderCredentialsClient IdentityProviderCredentialsClient

// NewIdentityProviderCredentialsClient instantiates a new IdentityProviderCredentialsClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderCredentialsClient() *IdentityProviderCredentialsClient {
	this := IdentityProviderCredentialsClient{}
	return &this
}

// NewIdentityProviderCredentialsClientWithDefaults instantiates a new IdentityProviderCredentialsClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderCredentialsClientWithDefaults() *IdentityProviderCredentialsClient {
	this := IdentityProviderCredentialsClient{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IdentityProviderCredentialsClient) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCredentialsClient) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IdentityProviderCredentialsClient) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IdentityProviderCredentialsClient) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *IdentityProviderCredentialsClient) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCredentialsClient) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *IdentityProviderCredentialsClient) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *IdentityProviderCredentialsClient) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetPkceRequired returns the PkceRequired field value if set, zero value otherwise.
func (o *IdentityProviderCredentialsClient) GetPkceRequired() bool {
	if o == nil || o.PkceRequired == nil {
		var ret bool
		return ret
	}
	return *o.PkceRequired
}

// GetPkceRequiredOk returns a tuple with the PkceRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCredentialsClient) GetPkceRequiredOk() (*bool, bool) {
	if o == nil || o.PkceRequired == nil {
		return nil, false
	}
	return o.PkceRequired, true
}

// HasPkceRequired returns a boolean if a field has been set.
func (o *IdentityProviderCredentialsClient) HasPkceRequired() bool {
	if o != nil && o.PkceRequired != nil {
		return true
	}

	return false
}

// SetPkceRequired gets a reference to the given bool and assigns it to the PkceRequired field.
func (o *IdentityProviderCredentialsClient) SetPkceRequired(v bool) {
	o.PkceRequired = &v
}

func (o IdentityProviderCredentialsClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.PkceRequired != nil {
		toSerialize["pkce_required"] = o.PkceRequired
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityProviderCredentialsClient) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityProviderCredentialsClient := _IdentityProviderCredentialsClient{}

	err = json.Unmarshal(bytes, &varIdentityProviderCredentialsClient)
	if err == nil {
		*o = IdentityProviderCredentialsClient(varIdentityProviderCredentialsClient)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "client_secret")
		delete(additionalProperties, "pkce_required")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdentityProviderCredentialsClient struct {
	value *IdentityProviderCredentialsClient
	isSet bool
}

func (v NullableIdentityProviderCredentialsClient) Get() *IdentityProviderCredentialsClient {
	return v.value
}

func (v *NullableIdentityProviderCredentialsClient) Set(val *IdentityProviderCredentialsClient) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderCredentialsClient) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderCredentialsClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderCredentialsClient(val *IdentityProviderCredentialsClient) *NullableIdentityProviderCredentialsClient {
	return &NullableIdentityProviderCredentialsClient{value: val, isSet: true}
}

func (v NullableIdentityProviderCredentialsClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderCredentialsClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

