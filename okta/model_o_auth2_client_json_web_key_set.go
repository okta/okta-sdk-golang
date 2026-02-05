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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the OAuth2ClientJsonWebKeySet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ClientJsonWebKeySet{}

// OAuth2ClientJsonWebKeySet A JSON Web Key Set (JWKS) containing the OAuth 2.0 client's JSON Web Keys
type OAuth2ClientJsonWebKeySet struct {
	// The JSON Web Keys in this key set
	Keys                 []AddJwk201Response `json:"keys,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientJsonWebKeySet OAuth2ClientJsonWebKeySet

// NewOAuth2ClientJsonWebKeySet instantiates a new OAuth2ClientJsonWebKeySet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientJsonWebKeySet() *OAuth2ClientJsonWebKeySet {
	this := OAuth2ClientJsonWebKeySet{}
	return &this
}

// NewOAuth2ClientJsonWebKeySetWithDefaults instantiates a new OAuth2ClientJsonWebKeySet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientJsonWebKeySetWithDefaults() *OAuth2ClientJsonWebKeySet {
	this := OAuth2ClientJsonWebKeySet{}
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeySet) GetKeys() []AddJwk201Response {
	if o == nil || IsNil(o.Keys) {
		var ret []AddJwk201Response
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeySet) GetKeysOk() ([]AddJwk201Response, bool) {
	if o == nil || IsNil(o.Keys) {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeySet) HasKeys() bool {
	if o != nil && !IsNil(o.Keys) {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []AddJwk201Response and assigns it to the Keys field.
func (o *OAuth2ClientJsonWebKeySet) SetKeys(v []AddJwk201Response) {
	o.Keys = v
}

func (o OAuth2ClientJsonWebKeySet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ClientJsonWebKeySet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Keys) {
		toSerialize["keys"] = o.Keys
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2ClientJsonWebKeySet) UnmarshalJSON(data []byte) (err error) {
	varOAuth2ClientJsonWebKeySet := _OAuth2ClientJsonWebKeySet{}

	err = json.Unmarshal(data, &varOAuth2ClientJsonWebKeySet)

	if err != nil {
		return err
	}

	*o = OAuth2ClientJsonWebKeySet(varOAuth2ClientJsonWebKeySet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "keys")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2ClientJsonWebKeySet struct {
	value *OAuth2ClientJsonWebKeySet
	isSet bool
}

func (v NullableOAuth2ClientJsonWebKeySet) Get() *OAuth2ClientJsonWebKeySet {
	return v.value
}

func (v *NullableOAuth2ClientJsonWebKeySet) Set(val *OAuth2ClientJsonWebKeySet) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientJsonWebKeySet) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientJsonWebKeySet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientJsonWebKeySet(val *OAuth2ClientJsonWebKeySet) *NullableOAuth2ClientJsonWebKeySet {
	return &NullableOAuth2ClientJsonWebKeySet{value: val, isSet: true}
}

func (v NullableOAuth2ClientJsonWebKeySet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientJsonWebKeySet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
