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

// checks if the SessionIdentityProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionIdentityProvider{}

// SessionIdentityProvider struct for SessionIdentityProvider
type SessionIdentityProvider struct {
	// Identity Provider ID. If the `type` is `OKTA`, then the `id` is the org ID.
	Id                   *string `json:"id,omitempty"`
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionIdentityProvider SessionIdentityProvider

// NewSessionIdentityProvider instantiates a new SessionIdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionIdentityProvider() *SessionIdentityProvider {
	this := SessionIdentityProvider{}
	return &this
}

// NewSessionIdentityProviderWithDefaults instantiates a new SessionIdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionIdentityProviderWithDefaults() *SessionIdentityProvider {
	this := SessionIdentityProvider{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SessionIdentityProvider) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionIdentityProvider) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SessionIdentityProvider) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SessionIdentityProvider) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SessionIdentityProvider) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionIdentityProvider) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SessionIdentityProvider) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SessionIdentityProvider) SetType(v string) {
	o.Type = &v
}

func (o SessionIdentityProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionIdentityProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SessionIdentityProvider) UnmarshalJSON(data []byte) (err error) {
	varSessionIdentityProvider := _SessionIdentityProvider{}

	err = json.Unmarshal(data, &varSessionIdentityProvider)

	if err != nil {
		return err
	}

	*o = SessionIdentityProvider(varSessionIdentityProvider)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSessionIdentityProvider struct {
	value *SessionIdentityProvider
	isSet bool
}

func (v NullableSessionIdentityProvider) Get() *SessionIdentityProvider {
	return v.value
}

func (v *NullableSessionIdentityProvider) Set(val *SessionIdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionIdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionIdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionIdentityProvider(val *SessionIdentityProvider) *NullableSessionIdentityProvider {
	return &NullableSessionIdentityProvider{value: val, isSet: true}
}

func (v NullableSessionIdentityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionIdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
