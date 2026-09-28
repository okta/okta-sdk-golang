/*
Okta Admin Management API

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

// checks if the ProximityProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProximityProvider{}

// ProximityProvider struct for ProximityProvider
type ProximityProvider struct {
	AuthenticatorCharacteristics *ProximityProviderAuthenticatorCharacteristics `json:"authenticatorCharacteristics,omitempty"`
	// Unique key for the proximity provider
	Id *string `json:"id,omitempty"`
	// Display name of the proximity provider
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProximityProvider ProximityProvider

// NewProximityProvider instantiates a new ProximityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProximityProvider() *ProximityProvider {
	this := ProximityProvider{}
	return &this
}

// NewProximityProviderWithDefaults instantiates a new ProximityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProximityProviderWithDefaults() *ProximityProvider {
	this := ProximityProvider{}
	return &this
}

// GetAuthenticatorCharacteristics returns the AuthenticatorCharacteristics field value if set, zero value otherwise.
func (o *ProximityProvider) GetAuthenticatorCharacteristics() ProximityProviderAuthenticatorCharacteristics {
	if o == nil || IsNil(o.AuthenticatorCharacteristics) {
		var ret ProximityProviderAuthenticatorCharacteristics
		return ret
	}
	return *o.AuthenticatorCharacteristics
}

// GetAuthenticatorCharacteristicsOk returns a tuple with the AuthenticatorCharacteristics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProximityProvider) GetAuthenticatorCharacteristicsOk() (*ProximityProviderAuthenticatorCharacteristics, bool) {
	if o == nil || IsNil(o.AuthenticatorCharacteristics) {
		return nil, false
	}
	return o.AuthenticatorCharacteristics, true
}

// HasAuthenticatorCharacteristics returns a boolean if a field has been set.
func (o *ProximityProvider) HasAuthenticatorCharacteristics() bool {
	if o != nil && !IsNil(o.AuthenticatorCharacteristics) {
		return true
	}

	return false
}

// SetAuthenticatorCharacteristics gets a reference to the given ProximityProviderAuthenticatorCharacteristics and assigns it to the AuthenticatorCharacteristics field.
func (o *ProximityProvider) SetAuthenticatorCharacteristics(v ProximityProviderAuthenticatorCharacteristics) {
	o.AuthenticatorCharacteristics = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProximityProvider) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProximityProvider) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProximityProvider) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProximityProvider) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProximityProvider) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProximityProvider) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProximityProvider) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProximityProvider) SetName(v string) {
	o.Name = &v
}

func (o ProximityProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProximityProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenticatorCharacteristics) {
		toSerialize["authenticatorCharacteristics"] = o.AuthenticatorCharacteristics
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProximityProvider) UnmarshalJSON(data []byte) (err error) {
	varProximityProvider := _ProximityProvider{}

	err = json.Unmarshal(data, &varProximityProvider)

	if err != nil {
		return err
	}

	*o = ProximityProvider(varProximityProvider)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticatorCharacteristics")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProximityProvider struct {
	value *ProximityProvider
	isSet bool
}

func (v NullableProximityProvider) Get() *ProximityProvider {
	return v.value
}

func (v *NullableProximityProvider) Set(val *ProximityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableProximityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableProximityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProximityProvider(val *ProximityProvider) *NullableProximityProvider {
	return &NullableProximityProvider{value: val, isSet: true}
}

func (v NullableProximityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProximityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
