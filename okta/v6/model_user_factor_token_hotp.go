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
	"reflect"
	"strings"
)

// checks if the UserFactorTokenHOTP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFactorTokenHOTP{}

// UserFactorTokenHOTP struct for UserFactorTokenHOTP
type UserFactorTokenHOTP struct {
	UserFactor
	// ID of an existing Custom TOTP factor profile. To create this, see [Custom TOTP factor](https://help.okta.com/okta_help.htm?id=ext-mfa-totp).
	FactorProfileId      *string                     `json:"factorProfileId,omitempty"`
	FactorType           interface{}                 `json:"factorType,omitempty"`
	Profile              *UserFactorTokenHOTPProfile `json:"profile,omitempty"`
	Provider             *string                     `json:"provider,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorTokenHOTP UserFactorTokenHOTP

// NewUserFactorTokenHOTP instantiates a new UserFactorTokenHOTP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorTokenHOTP() *UserFactorTokenHOTP {
	this := UserFactorTokenHOTP{}
	return &this
}

// NewUserFactorTokenHOTPWithDefaults instantiates a new UserFactorTokenHOTP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorTokenHOTPWithDefaults() *UserFactorTokenHOTP {
	this := UserFactorTokenHOTP{}
	return &this
}

// GetFactorProfileId returns the FactorProfileId field value if set, zero value otherwise.
func (o *UserFactorTokenHOTP) GetFactorProfileId() string {
	if o == nil || IsNil(o.FactorProfileId) {
		var ret string
		return ret
	}
	return *o.FactorProfileId
}

// GetFactorProfileIdOk returns a tuple with the FactorProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenHOTP) GetFactorProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.FactorProfileId) {
		return nil, false
	}
	return o.FactorProfileId, true
}

// HasFactorProfileId returns a boolean if a field has been set.
func (o *UserFactorTokenHOTP) HasFactorProfileId() bool {
	if o != nil && !IsNil(o.FactorProfileId) {
		return true
	}

	return false
}

// SetFactorProfileId gets a reference to the given string and assigns it to the FactorProfileId field.
func (o *UserFactorTokenHOTP) SetFactorProfileId(v string) {
	o.FactorProfileId = &v
}

// GetFactorType returns the FactorType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserFactorTokenHOTP) GetFactorType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FactorType
}

// GetFactorTypeOk returns a tuple with the FactorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserFactorTokenHOTP) GetFactorTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FactorType) {
		return nil, false
	}
	return &o.FactorType, true
}

// HasFactorType returns a boolean if a field has been set.
func (o *UserFactorTokenHOTP) HasFactorType() bool {
	if o != nil && !IsNil(o.FactorType) {
		return true
	}

	return false
}

// SetFactorType gets a reference to the given interface{} and assigns it to the FactorType field.
func (o *UserFactorTokenHOTP) SetFactorType(v interface{}) {
	o.FactorType = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactorTokenHOTP) GetProfile() UserFactorTokenHOTPProfile {
	if o == nil || IsNil(o.Profile) {
		var ret UserFactorTokenHOTPProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenHOTP) GetProfileOk() (*UserFactorTokenHOTPProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactorTokenHOTP) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserFactorTokenHOTPProfile and assigns it to the Profile field.
func (o *UserFactorTokenHOTP) SetProfile(v UserFactorTokenHOTPProfile) {
	o.Profile = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *UserFactorTokenHOTP) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenHOTP) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *UserFactorTokenHOTP) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *UserFactorTokenHOTP) SetProvider(v string) {
	o.Provider = &v
}

func (o UserFactorTokenHOTP) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFactorTokenHOTP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedUserFactor, errUserFactor := json.Marshal(o.UserFactor)
	if errUserFactor != nil {
		return map[string]interface{}{}, errUserFactor
	}
	errUserFactor = json.Unmarshal([]byte(serializedUserFactor), &toSerialize)
	if errUserFactor != nil {
		return map[string]interface{}{}, errUserFactor
	}
	if !IsNil(o.FactorProfileId) {
		toSerialize["factorProfileId"] = o.FactorProfileId
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

func (o *UserFactorTokenHOTP) UnmarshalJSON(data []byte) (err error) {
	type UserFactorTokenHOTPWithoutEmbeddedStruct struct {
		// ID of an existing Custom TOTP factor profile. To create this, see [Custom TOTP factor](https://help.okta.com/okta_help.htm?id=ext-mfa-totp).
		FactorProfileId *string                     `json:"factorProfileId,omitempty"`
		FactorType      interface{}                 `json:"factorType,omitempty"`
		Profile         *UserFactorTokenHOTPProfile `json:"profile,omitempty"`
		Provider        *string                     `json:"provider,omitempty"`
	}

	varUserFactorTokenHOTPWithoutEmbeddedStruct := UserFactorTokenHOTPWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varUserFactorTokenHOTPWithoutEmbeddedStruct)
	if err == nil {
		varUserFactorTokenHOTP := _UserFactorTokenHOTP{}
		varUserFactorTokenHOTP.FactorProfileId = varUserFactorTokenHOTPWithoutEmbeddedStruct.FactorProfileId
		varUserFactorTokenHOTP.FactorType = varUserFactorTokenHOTPWithoutEmbeddedStruct.FactorType
		varUserFactorTokenHOTP.Profile = varUserFactorTokenHOTPWithoutEmbeddedStruct.Profile
		varUserFactorTokenHOTP.Provider = varUserFactorTokenHOTPWithoutEmbeddedStruct.Provider
		*o = UserFactorTokenHOTP(varUserFactorTokenHOTP)
	} else {
		return err
	}

	varUserFactorTokenHOTP := _UserFactorTokenHOTP{}

	err = json.Unmarshal(data, &varUserFactorTokenHOTP)
	if err == nil {
		o.UserFactor = varUserFactorTokenHOTP.UserFactor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "factorProfileId")
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

type NullableUserFactorTokenHOTP struct {
	value *UserFactorTokenHOTP
	isSet bool
}

func (v NullableUserFactorTokenHOTP) Get() *UserFactorTokenHOTP {
	return v.value
}

func (v *NullableUserFactorTokenHOTP) Set(val *UserFactorTokenHOTP) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorTokenHOTP) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorTokenHOTP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorTokenHOTP(val *UserFactorTokenHOTP) *NullableUserFactorTokenHOTP {
	return &NullableUserFactorTokenHOTP{value: val, isSet: true}
}

func (v NullableUserFactorTokenHOTP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorTokenHOTP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
