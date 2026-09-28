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
	"reflect"
	"strings"
)

// checks if the UserFactorSignedNonce type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFactorSignedNonce{}

// UserFactorSignedNonce `signed_nonce` is the factor type for [Okta FastPass](https://help.okta.com/oie/en-us/content/topics/identity-engine/devices/fp/fp-main.htm). You can't use the Factors API to enroll or activate Okta FastPass (`signed_nonce`) for a user. Use the [Okta Verify](https://help.okta.com/en-us/content/topics/mobile/okta-verify-overview.htm) authenticator enrollment flow instead.  You can use the Factors API to list and delete `signed_nonce` factors.
type UserFactorSignedNonce struct {
	UserFactor
	FactorType           interface{}                   `json:"factorType,omitempty"`
	Profile              *UserFactorSignedNonceProfile `json:"profile,omitempty"`
	Provider             *string                       `json:"provider,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorSignedNonce UserFactorSignedNonce

// NewUserFactorSignedNonce instantiates a new UserFactorSignedNonce object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorSignedNonce() *UserFactorSignedNonce {
	this := UserFactorSignedNonce{}
	return &this
}

// NewUserFactorSignedNonceWithDefaults instantiates a new UserFactorSignedNonce object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorSignedNonceWithDefaults() *UserFactorSignedNonce {
	this := UserFactorSignedNonce{}
	return &this
}

// GetFactorType returns the FactorType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserFactorSignedNonce) GetFactorType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FactorType
}

// GetFactorTypeOk returns a tuple with the FactorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserFactorSignedNonce) GetFactorTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FactorType) {
		return nil, false
	}
	return &o.FactorType, true
}

// HasFactorType returns a boolean if a field has been set.
func (o *UserFactorSignedNonce) HasFactorType() bool {
	if o != nil && !IsNil(o.FactorType) {
		return true
	}

	return false
}

// SetFactorType gets a reference to the given interface{} and assigns it to the FactorType field.
func (o *UserFactorSignedNonce) SetFactorType(v interface{}) {
	o.FactorType = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactorSignedNonce) GetProfile() UserFactorSignedNonceProfile {
	if o == nil || IsNil(o.Profile) {
		var ret UserFactorSignedNonceProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonce) GetProfileOk() (*UserFactorSignedNonceProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactorSignedNonce) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserFactorSignedNonceProfile and assigns it to the Profile field.
func (o *UserFactorSignedNonce) SetProfile(v UserFactorSignedNonceProfile) {
	o.Profile = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *UserFactorSignedNonce) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonce) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *UserFactorSignedNonce) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *UserFactorSignedNonce) SetProvider(v string) {
	o.Provider = &v
}

func (o UserFactorSignedNonce) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFactorSignedNonce) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedUserFactor, errUserFactor := json.Marshal(o.UserFactor)
	if errUserFactor != nil {
		return map[string]interface{}{}, errUserFactor
	}
	errUserFactor = json.Unmarshal([]byte(serializedUserFactor), &toSerialize)
	if errUserFactor != nil {
		return map[string]interface{}{}, errUserFactor
	}
	if o.FactorType != nil {
		toSerialize["factorType"] = o.FactorType
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserFactorSignedNonce) UnmarshalJSON(data []byte) (err error) {
	type UserFactorSignedNonceWithoutEmbeddedStruct struct {
		FactorType interface{}                   `json:"factorType,omitempty"`
		Profile    *UserFactorSignedNonceProfile `json:"profile,omitempty"`
		Provider   *string                       `json:"provider,omitempty"`
	}

	varUserFactorSignedNonceWithoutEmbeddedStruct := UserFactorSignedNonceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varUserFactorSignedNonceWithoutEmbeddedStruct)
	if err == nil {
		varUserFactorSignedNonce := _UserFactorSignedNonce{}
		varUserFactorSignedNonce.FactorType = varUserFactorSignedNonceWithoutEmbeddedStruct.FactorType
		varUserFactorSignedNonce.Profile = varUserFactorSignedNonceWithoutEmbeddedStruct.Profile
		varUserFactorSignedNonce.Provider = varUserFactorSignedNonceWithoutEmbeddedStruct.Provider
		*o = UserFactorSignedNonce(varUserFactorSignedNonce)
	} else {
		return err
	}

	varUserFactorSignedNonce := _UserFactorSignedNonce{}

	err = json.Unmarshal(data, &varUserFactorSignedNonce)
	if err == nil {
		o.UserFactor = varUserFactorSignedNonce.UserFactor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "factorType")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "provider")

		// remove fields from embedded structs
		reflectUserFactor := reflect.ValueOf(o.UserFactor)
		for i := 0; i < reflectUserFactor.Type().NumField(); i++ {
			t := reflectUserFactor.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserFactorSignedNonce struct {
	value *UserFactorSignedNonce
	isSet bool
}

func (v NullableUserFactorSignedNonce) Get() *UserFactorSignedNonce {
	return v.value
}

func (v *NullableUserFactorSignedNonce) Set(val *UserFactorSignedNonce) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorSignedNonce) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorSignedNonce) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorSignedNonce(val *UserFactorSignedNonce) *NullableUserFactorSignedNonce {
	return &NullableUserFactorSignedNonce{value: val, isSet: true}
}

func (v NullableUserFactorSignedNonce) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorSignedNonce) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
