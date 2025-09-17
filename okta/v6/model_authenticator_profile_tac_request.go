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

// checks if the AuthenticatorProfileTacRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorProfileTacRequest{}

// AuthenticatorProfileTacRequest Defines the authenticator specific parameters
type AuthenticatorProfileTacRequest struct {
	// Determines whether the enrollment can be used more than once. To enable multi-use, the org-level authenticator’s configuration must allow multi-use.
	MultiUse *bool `json:"multiUse,omitempty"`
	// Time-to-live (TTL) in minutes.  Specifies how long the TAC enrollment is valid after it's created and activated. The configured value must be between 10 minutes (`10`) and 10 days (`14400`), inclusive. The actual allowed range depends on the org-level authenticator configuration.
	Ttl                  *string `json:"ttl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorProfileTacRequest AuthenticatorProfileTacRequest

// NewAuthenticatorProfileTacRequest instantiates a new AuthenticatorProfileTacRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorProfileTacRequest() *AuthenticatorProfileTacRequest {
	this := AuthenticatorProfileTacRequest{}
	return &this
}

// NewAuthenticatorProfileTacRequestWithDefaults instantiates a new AuthenticatorProfileTacRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorProfileTacRequestWithDefaults() *AuthenticatorProfileTacRequest {
	this := AuthenticatorProfileTacRequest{}
	return &this
}

// GetMultiUse returns the MultiUse field value if set, zero value otherwise.
func (o *AuthenticatorProfileTacRequest) GetMultiUse() bool {
	if o == nil || IsNil(o.MultiUse) {
		var ret bool
		return ret
	}
	return *o.MultiUse
}

// GetMultiUseOk returns a tuple with the MultiUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorProfileTacRequest) GetMultiUseOk() (*bool, bool) {
	if o == nil || IsNil(o.MultiUse) {
		return nil, false
	}
	return o.MultiUse, true
}

// HasMultiUse returns a boolean if a field has been set.
func (o *AuthenticatorProfileTacRequest) HasMultiUse() bool {
	if o != nil && !IsNil(o.MultiUse) {
		return true
	}

	return false
}

// SetMultiUse gets a reference to the given bool and assigns it to the MultiUse field.
func (o *AuthenticatorProfileTacRequest) SetMultiUse(v bool) {
	o.MultiUse = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *AuthenticatorProfileTacRequest) GetTtl() string {
	if o == nil || IsNil(o.Ttl) {
		var ret string
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorProfileTacRequest) GetTtlOk() (*string, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *AuthenticatorProfileTacRequest) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given string and assigns it to the Ttl field.
func (o *AuthenticatorProfileTacRequest) SetTtl(v string) {
	o.Ttl = &v
}

func (o AuthenticatorProfileTacRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorProfileTacRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MultiUse) {
		toSerialize["multiUse"] = o.MultiUse
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorProfileTacRequest) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorProfileTacRequest := _AuthenticatorProfileTacRequest{}

	err = json.Unmarshal(data, &varAuthenticatorProfileTacRequest)

	if err != nil {
		return err
	}

	*o = AuthenticatorProfileTacRequest(varAuthenticatorProfileTacRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "multiUse")
		delete(additionalProperties, "ttl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorProfileTacRequest struct {
	value *AuthenticatorProfileTacRequest
	isSet bool
}

func (v NullableAuthenticatorProfileTacRequest) Get() *AuthenticatorProfileTacRequest {
	return v.value
}

func (v *NullableAuthenticatorProfileTacRequest) Set(val *AuthenticatorProfileTacRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorProfileTacRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorProfileTacRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorProfileTacRequest(val *AuthenticatorProfileTacRequest) *NullableAuthenticatorProfileTacRequest {
	return &NullableAuthenticatorProfileTacRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorProfileTacRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorProfileTacRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
