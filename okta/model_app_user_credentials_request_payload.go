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

// checks if the AppUserCredentialsRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserCredentialsRequestPayload{}

// AppUserCredentialsRequestPayload Updates the assigned user credentials
type AppUserCredentialsRequestPayload struct {
	Credentials          *AppUserCredentials `json:"credentials,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppUserCredentialsRequestPayload AppUserCredentialsRequestPayload

// NewAppUserCredentialsRequestPayload instantiates a new AppUserCredentialsRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUserCredentialsRequestPayload() *AppUserCredentialsRequestPayload {
	this := AppUserCredentialsRequestPayload{}
	return &this
}

// NewAppUserCredentialsRequestPayloadWithDefaults instantiates a new AppUserCredentialsRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserCredentialsRequestPayloadWithDefaults() *AppUserCredentialsRequestPayload {
	this := AppUserCredentialsRequestPayload{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *AppUserCredentialsRequestPayload) GetCredentials() AppUserCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret AppUserCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserCredentialsRequestPayload) GetCredentialsOk() (*AppUserCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *AppUserCredentialsRequestPayload) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given AppUserCredentials and assigns it to the Credentials field.
func (o *AppUserCredentialsRequestPayload) SetCredentials(v AppUserCredentials) {
	o.Credentials = &v
}

func (o AppUserCredentialsRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUserCredentialsRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppUserCredentialsRequestPayload) UnmarshalJSON(data []byte) (err error) {
	varAppUserCredentialsRequestPayload := _AppUserCredentialsRequestPayload{}

	err = json.Unmarshal(data, &varAppUserCredentialsRequestPayload)

	if err != nil {
		return err
	}

	*o = AppUserCredentialsRequestPayload(varAppUserCredentialsRequestPayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "credentials")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppUserCredentialsRequestPayload struct {
	value *AppUserCredentialsRequestPayload
	isSet bool
}

func (v NullableAppUserCredentialsRequestPayload) Get() *AppUserCredentialsRequestPayload {
	return v.value
}

func (v *NullableAppUserCredentialsRequestPayload) Set(val *AppUserCredentialsRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserCredentialsRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserCredentialsRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserCredentialsRequestPayload(val *AppUserCredentialsRequestPayload) *NullableAppUserCredentialsRequestPayload {
	return &NullableAppUserCredentialsRequestPayload{value: val, isSet: true}
}

func (v NullableAppUserCredentialsRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserCredentialsRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
