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
	"reflect"
	"strings"
)

// UserFactorCustomHOTP struct for UserFactorCustomHOTP
type UserFactorCustomHOTP struct {
	UserFactor
	// ID of an existing Custom TOTP Factor profile. To create this, see [Custom TOTP Factor](https://help.okta.com/okta_help.htm?id=ext-mfa-totp).
	FactorProfileId *string `json:"factorProfileId,omitempty"`
	FactorType interface{} `json:"factorType,omitempty"`
	Profile *UserFactorCustomHOTPProfile `json:"profile,omitempty"`
	Provider *string `json:"provider,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorCustomHOTP UserFactorCustomHOTP

// NewUserFactorCustomHOTP instantiates a new UserFactorCustomHOTP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorCustomHOTP() *UserFactorCustomHOTP {
	this := UserFactorCustomHOTP{}
	return &this
}

// NewUserFactorCustomHOTPWithDefaults instantiates a new UserFactorCustomHOTP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorCustomHOTPWithDefaults() *UserFactorCustomHOTP {
	this := UserFactorCustomHOTP{}
	return &this
}

// GetFactorProfileId returns the FactorProfileId field value if set, zero value otherwise.
func (o *UserFactorCustomHOTP) GetFactorProfileId() string {
	if o == nil || o.FactorProfileId == nil {
		var ret string
		return ret
	}
	return *o.FactorProfileId
}

// GetFactorProfileIdOk returns a tuple with the FactorProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorCustomHOTP) GetFactorProfileIdOk() (*string, bool) {
	if o == nil || o.FactorProfileId == nil {
		return nil, false
	}
	return o.FactorProfileId, true
}

// HasFactorProfileId returns a boolean if a field has been set.
func (o *UserFactorCustomHOTP) HasFactorProfileId() bool {
	if o != nil && o.FactorProfileId != nil {
		return true
	}

	return false
}

// SetFactorProfileId gets a reference to the given string and assigns it to the FactorProfileId field.
func (o *UserFactorCustomHOTP) SetFactorProfileId(v string) {
	o.FactorProfileId = &v
}

// GetFactorType returns the FactorType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserFactorCustomHOTP) GetFactorType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FactorType
}

// GetFactorTypeOk returns a tuple with the FactorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserFactorCustomHOTP) GetFactorTypeOk() (*interface{}, bool) {
	if o == nil || o.FactorType == nil {
		return nil, false
	}
	return &o.FactorType, true
}

// HasFactorType returns a boolean if a field has been set.
func (o *UserFactorCustomHOTP) HasFactorType() bool {
	if o != nil && o.FactorType != nil {
		return true
	}

	return false
}

// SetFactorType gets a reference to the given interface{} and assigns it to the FactorType field.
func (o *UserFactorCustomHOTP) SetFactorType(v interface{}) {
	o.FactorType = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactorCustomHOTP) GetProfile() UserFactorCustomHOTPProfile {
	if o == nil || o.Profile == nil {
		var ret UserFactorCustomHOTPProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorCustomHOTP) GetProfileOk() (*UserFactorCustomHOTPProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactorCustomHOTP) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserFactorCustomHOTPProfile and assigns it to the Profile field.
func (o *UserFactorCustomHOTP) SetProfile(v UserFactorCustomHOTPProfile) {
	o.Profile = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *UserFactorCustomHOTP) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorCustomHOTP) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *UserFactorCustomHOTP) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *UserFactorCustomHOTP) SetProvider(v string) {
	o.Provider = &v
}

func (o UserFactorCustomHOTP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedUserFactor, errUserFactor := json.Marshal(o.UserFactor)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	errUserFactor = json.Unmarshal([]byte(serializedUserFactor), &toSerialize)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	if o.FactorProfileId != nil {
		toSerialize["factorProfileId"] = o.FactorProfileId
	}
	if o.FactorType != nil {
		toSerialize["factorType"] = o.FactorType
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorCustomHOTP) UnmarshalJSON(bytes []byte) (err error) {
	type UserFactorCustomHOTPWithoutEmbeddedStruct struct {
		// ID of an existing Custom TOTP Factor profile. To create this, see [Custom TOTP Factor](https://help.okta.com/okta_help.htm?id=ext-mfa-totp).
		FactorProfileId *string `json:"factorProfileId,omitempty"`
		FactorType interface{} `json:"factorType,omitempty"`
		Profile *UserFactorCustomHOTPProfile `json:"profile,omitempty"`
		Provider *string `json:"provider,omitempty"`
	}

	varUserFactorCustomHOTPWithoutEmbeddedStruct := UserFactorCustomHOTPWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varUserFactorCustomHOTPWithoutEmbeddedStruct)
	if err == nil {
		varUserFactorCustomHOTP := _UserFactorCustomHOTP{}
		varUserFactorCustomHOTP.FactorProfileId = varUserFactorCustomHOTPWithoutEmbeddedStruct.FactorProfileId
		varUserFactorCustomHOTP.FactorType = varUserFactorCustomHOTPWithoutEmbeddedStruct.FactorType
		varUserFactorCustomHOTP.Profile = varUserFactorCustomHOTPWithoutEmbeddedStruct.Profile
		varUserFactorCustomHOTP.Provider = varUserFactorCustomHOTPWithoutEmbeddedStruct.Provider
		*o = UserFactorCustomHOTP(varUserFactorCustomHOTP)
	} else {
		return err
	}

	varUserFactorCustomHOTP := _UserFactorCustomHOTP{}

	err = json.Unmarshal(bytes, &varUserFactorCustomHOTP)
	if err == nil {
		o.UserFactor = varUserFactorCustomHOTP.UserFactor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
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
	} else {
		return err
	}

	return err
}

type NullableUserFactorCustomHOTP struct {
	value *UserFactorCustomHOTP
	isSet bool
}

func (v NullableUserFactorCustomHOTP) Get() *UserFactorCustomHOTP {
	return v.value
}

func (v *NullableUserFactorCustomHOTP) Set(val *UserFactorCustomHOTP) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorCustomHOTP) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorCustomHOTP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorCustomHOTP(val *UserFactorCustomHOTP) *NullableUserFactorCustomHOTP {
	return &NullableUserFactorCustomHOTP{value: val, isSet: true}
}

func (v NullableUserFactorCustomHOTP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorCustomHOTP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

