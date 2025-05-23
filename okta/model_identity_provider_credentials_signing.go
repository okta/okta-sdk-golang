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

// IdentityProviderCredentialsSigning struct for IdentityProviderCredentialsSigning
type IdentityProviderCredentialsSigning struct {
	Kid *string `json:"kid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProviderCredentialsSigning IdentityProviderCredentialsSigning

// NewIdentityProviderCredentialsSigning instantiates a new IdentityProviderCredentialsSigning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderCredentialsSigning() *IdentityProviderCredentialsSigning {
	this := IdentityProviderCredentialsSigning{}
	return &this
}

// NewIdentityProviderCredentialsSigningWithDefaults instantiates a new IdentityProviderCredentialsSigning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderCredentialsSigningWithDefaults() *IdentityProviderCredentialsSigning {
	this := IdentityProviderCredentialsSigning{}
	return &this
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *IdentityProviderCredentialsSigning) GetKid() string {
	if o == nil || o.Kid == nil {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCredentialsSigning) GetKidOk() (*string, bool) {
	if o == nil || o.Kid == nil {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *IdentityProviderCredentialsSigning) HasKid() bool {
	if o != nil && o.Kid != nil {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *IdentityProviderCredentialsSigning) SetKid(v string) {
	o.Kid = &v
}

func (o IdentityProviderCredentialsSigning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kid != nil {
		toSerialize["kid"] = o.Kid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityProviderCredentialsSigning) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityProviderCredentialsSigning := _IdentityProviderCredentialsSigning{}

	err = json.Unmarshal(bytes, &varIdentityProviderCredentialsSigning)
	if err == nil {
		*o = IdentityProviderCredentialsSigning(varIdentityProviderCredentialsSigning)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "kid")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdentityProviderCredentialsSigning struct {
	value *IdentityProviderCredentialsSigning
	isSet bool
}

func (v NullableIdentityProviderCredentialsSigning) Get() *IdentityProviderCredentialsSigning {
	return v.value
}

func (v *NullableIdentityProviderCredentialsSigning) Set(val *IdentityProviderCredentialsSigning) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderCredentialsSigning) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderCredentialsSigning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderCredentialsSigning(val *IdentityProviderCredentialsSigning) *NullableIdentityProviderCredentialsSigning {
	return &NullableIdentityProviderCredentialsSigning{value: val, isSet: true}
}

func (v NullableIdentityProviderCredentialsSigning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderCredentialsSigning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

