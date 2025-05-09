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

// UserSchemaProperties struct for UserSchemaProperties
type UserSchemaProperties struct {
	Profile *UserSchemaPropertiesProfile `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSchemaProperties UserSchemaProperties

// NewUserSchemaProperties instantiates a new UserSchemaProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSchemaProperties() *UserSchemaProperties {
	this := UserSchemaProperties{}
	return &this
}

// NewUserSchemaPropertiesWithDefaults instantiates a new UserSchemaProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSchemaPropertiesWithDefaults() *UserSchemaProperties {
	this := UserSchemaProperties{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserSchemaProperties) GetProfile() UserSchemaPropertiesProfile {
	if o == nil || o.Profile == nil {
		var ret UserSchemaPropertiesProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaProperties) GetProfileOk() (*UserSchemaPropertiesProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserSchemaProperties) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserSchemaPropertiesProfile and assigns it to the Profile field.
func (o *UserSchemaProperties) SetProfile(v UserSchemaPropertiesProfile) {
	o.Profile = &v
}

func (o UserSchemaProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserSchemaProperties) UnmarshalJSON(bytes []byte) (err error) {
	varUserSchemaProperties := _UserSchemaProperties{}

	err = json.Unmarshal(bytes, &varUserSchemaProperties)
	if err == nil {
		*o = UserSchemaProperties(varUserSchemaProperties)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserSchemaProperties struct {
	value *UserSchemaProperties
	isSet bool
}

func (v NullableUserSchemaProperties) Get() *UserSchemaProperties {
	return v.value
}

func (v *NullableUserSchemaProperties) Set(val *UserSchemaProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSchemaProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSchemaProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSchemaProperties(val *UserSchemaProperties) *NullableUserSchemaProperties {
	return &NullableUserSchemaProperties{value: val, isSet: true}
}

func (v NullableUserSchemaProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSchemaProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

