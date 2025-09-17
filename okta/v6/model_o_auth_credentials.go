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
)

// checks if the OAuthCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuthCredentials{}

// OAuthCredentials Client authentication credentials for an [OAuth 2.0 Authorization Server](https://tools.ietf.org/html/rfc6749#section-2.3)
type OAuthCredentials struct {
	Client               *OAuthCredentialsClient `json:"client,omitempty"`
	Signing              *AppleClientSigning     `json:"signing,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuthCredentials OAuthCredentials

// NewOAuthCredentials instantiates a new OAuthCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthCredentials() *OAuthCredentials {
	this := OAuthCredentials{}
	return &this
}

// NewOAuthCredentialsWithDefaults instantiates a new OAuthCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthCredentialsWithDefaults() *OAuthCredentials {
	this := OAuthCredentials{}
	return &this
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *OAuthCredentials) GetClient() OAuthCredentialsClient {
	if o == nil || IsNil(o.Client) {
		var ret OAuthCredentialsClient
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthCredentials) GetClientOk() (*OAuthCredentialsClient, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *OAuthCredentials) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given OAuthCredentialsClient and assigns it to the Client field.
func (o *OAuthCredentials) SetClient(v OAuthCredentialsClient) {
	o.Client = &v
}

// GetSigning returns the Signing field value if set, zero value otherwise.
func (o *OAuthCredentials) GetSigning() AppleClientSigning {
	if o == nil || IsNil(o.Signing) {
		var ret AppleClientSigning
		return ret
	}
	return *o.Signing
}

// GetSigningOk returns a tuple with the Signing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthCredentials) GetSigningOk() (*AppleClientSigning, bool) {
	if o == nil || IsNil(o.Signing) {
		return nil, false
	}
	return o.Signing, true
}

// HasSigning returns a boolean if a field has been set.
func (o *OAuthCredentials) HasSigning() bool {
	if o != nil && !IsNil(o.Signing) {
		return true
	}

	return false
}

// SetSigning gets a reference to the given AppleClientSigning and assigns it to the Signing field.
func (o *OAuthCredentials) SetSigning(v AppleClientSigning) {
	o.Signing = &v
}

func (o OAuthCredentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuthCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !IsNil(o.Signing) {
		toSerialize["signing"] = o.Signing
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuthCredentials) UnmarshalJSON(data []byte) (err error) {
	varOAuthCredentials := _OAuthCredentials{}

	err = json.Unmarshal(data, &varOAuthCredentials)

	if err != nil {
		return err
	}

	*o = OAuthCredentials(varOAuthCredentials)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "client")
		delete(additionalProperties, "signing")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuthCredentials struct {
	value *OAuthCredentials
	isSet bool
}

func (v NullableOAuthCredentials) Get() *OAuthCredentials {
	return v.value
}

func (v *NullableOAuthCredentials) Set(val *OAuthCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthCredentials(val *OAuthCredentials) *NullableOAuthCredentials {
	return &NullableOAuthCredentials{value: val, isSet: true}
}

func (v NullableOAuthCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
