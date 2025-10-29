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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the ApplicationEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationEmbedded{}

// ApplicationEmbedded Embedded resources related to the app using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification. If the `expand=user/{userId}` query parameter is specified, then the assigned [Application User](/openapi/okta-management/management/tag/ApplicationUsers/) is embedded.
type ApplicationEmbedded struct {
	// The specified [Application User](/openapi/okta-management/management/tag/ApplicationUsers/) assigned to the app
	User                 map[string]map[string]interface{} `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationEmbedded ApplicationEmbedded

// NewApplicationEmbedded instantiates a new ApplicationEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationEmbedded() *ApplicationEmbedded {
	this := ApplicationEmbedded{}
	return &this
}

// NewApplicationEmbeddedWithDefaults instantiates a new ApplicationEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationEmbeddedWithDefaults() *ApplicationEmbedded {
	this := ApplicationEmbedded{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ApplicationEmbedded) GetUser() map[string]map[string]interface{} {
	if o == nil || IsNil(o.User) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEmbedded) GetUserOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.User) {
		return map[string]map[string]interface{}{}, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ApplicationEmbedded) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given map[string]map[string]interface{} and assigns it to the User field.
func (o *ApplicationEmbedded) SetUser(v map[string]map[string]interface{}) {
	o.User = v
}

func (o ApplicationEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationEmbedded) UnmarshalJSON(data []byte) (err error) {
	varApplicationEmbedded := _ApplicationEmbedded{}

	err = json.Unmarshal(data, &varApplicationEmbedded)

	if err != nil {
		return err
	}

	*o = ApplicationEmbedded(varApplicationEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationEmbedded struct {
	value *ApplicationEmbedded
	isSet bool
}

func (v NullableApplicationEmbedded) Get() *ApplicationEmbedded {
	return v.value
}

func (v *NullableApplicationEmbedded) Set(val *ApplicationEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationEmbedded(val *ApplicationEmbedded) *NullableApplicationEmbedded {
	return &NullableApplicationEmbedded{value: val, isSet: true}
}

func (v NullableApplicationEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
