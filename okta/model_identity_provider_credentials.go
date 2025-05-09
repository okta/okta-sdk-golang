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

// IdentityProviderCredentials struct for IdentityProviderCredentials
type IdentityProviderCredentials struct {
	Client *IdentityProviderCredentialsClient `json:"client,omitempty"`
	Signing *IdentityProviderCredentialsSigning `json:"signing,omitempty"`
	Trust *IdentityProviderCredentialsTrust `json:"trust,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProviderCredentials IdentityProviderCredentials

// NewIdentityProviderCredentials instantiates a new IdentityProviderCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderCredentials() *IdentityProviderCredentials {
	this := IdentityProviderCredentials{}
	return &this
}

// NewIdentityProviderCredentialsWithDefaults instantiates a new IdentityProviderCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderCredentialsWithDefaults() *IdentityProviderCredentials {
	this := IdentityProviderCredentials{}
	return &this
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *IdentityProviderCredentials) GetClient() IdentityProviderCredentialsClient {
	if o == nil || o.Client == nil {
		var ret IdentityProviderCredentialsClient
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCredentials) GetClientOk() (*IdentityProviderCredentialsClient, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *IdentityProviderCredentials) HasClient() bool {
	if o != nil && o.Client != nil {
		return true
	}

	return false
}

// SetClient gets a reference to the given IdentityProviderCredentialsClient and assigns it to the Client field.
func (o *IdentityProviderCredentials) SetClient(v IdentityProviderCredentialsClient) {
	o.Client = &v
}

// GetSigning returns the Signing field value if set, zero value otherwise.
func (o *IdentityProviderCredentials) GetSigning() IdentityProviderCredentialsSigning {
	if o == nil || o.Signing == nil {
		var ret IdentityProviderCredentialsSigning
		return ret
	}
	return *o.Signing
}

// GetSigningOk returns a tuple with the Signing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCredentials) GetSigningOk() (*IdentityProviderCredentialsSigning, bool) {
	if o == nil || o.Signing == nil {
		return nil, false
	}
	return o.Signing, true
}

// HasSigning returns a boolean if a field has been set.
func (o *IdentityProviderCredentials) HasSigning() bool {
	if o != nil && o.Signing != nil {
		return true
	}

	return false
}

// SetSigning gets a reference to the given IdentityProviderCredentialsSigning and assigns it to the Signing field.
func (o *IdentityProviderCredentials) SetSigning(v IdentityProviderCredentialsSigning) {
	o.Signing = &v
}

// GetTrust returns the Trust field value if set, zero value otherwise.
func (o *IdentityProviderCredentials) GetTrust() IdentityProviderCredentialsTrust {
	if o == nil || o.Trust == nil {
		var ret IdentityProviderCredentialsTrust
		return ret
	}
	return *o.Trust
}

// GetTrustOk returns a tuple with the Trust field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCredentials) GetTrustOk() (*IdentityProviderCredentialsTrust, bool) {
	if o == nil || o.Trust == nil {
		return nil, false
	}
	return o.Trust, true
}

// HasTrust returns a boolean if a field has been set.
func (o *IdentityProviderCredentials) HasTrust() bool {
	if o != nil && o.Trust != nil {
		return true
	}

	return false
}

// SetTrust gets a reference to the given IdentityProviderCredentialsTrust and assigns it to the Trust field.
func (o *IdentityProviderCredentials) SetTrust(v IdentityProviderCredentialsTrust) {
	o.Trust = &v
}

func (o IdentityProviderCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Client != nil {
		toSerialize["client"] = o.Client
	}
	if o.Signing != nil {
		toSerialize["signing"] = o.Signing
	}
	if o.Trust != nil {
		toSerialize["trust"] = o.Trust
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityProviderCredentials) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityProviderCredentials := _IdentityProviderCredentials{}

	err = json.Unmarshal(bytes, &varIdentityProviderCredentials)
	if err == nil {
		*o = IdentityProviderCredentials(varIdentityProviderCredentials)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "client")
		delete(additionalProperties, "signing")
		delete(additionalProperties, "trust")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdentityProviderCredentials struct {
	value *IdentityProviderCredentials
	isSet bool
}

func (v NullableIdentityProviderCredentials) Get() *IdentityProviderCredentials {
	return v.value
}

func (v *NullableIdentityProviderCredentials) Set(val *IdentityProviderCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderCredentials(val *IdentityProviderCredentials) *NullableIdentityProviderCredentials {
	return &NullableIdentityProviderCredentials{value: val, isSet: true}
}

func (v NullableIdentityProviderCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

