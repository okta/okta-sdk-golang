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

// checks if the UsersUpdateRequestSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsersUpdateRequestSchema{}

// UsersUpdateRequestSchema struct for UsersUpdateRequestSchema
type UsersUpdateRequestSchema struct {
	Profile              *IdentitySourceUserProfileForUpsert `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UsersUpdateRequestSchema UsersUpdateRequestSchema

// NewUsersUpdateRequestSchema instantiates a new UsersUpdateRequestSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersUpdateRequestSchema() *UsersUpdateRequestSchema {
	this := UsersUpdateRequestSchema{}
	return &this
}

// NewUsersUpdateRequestSchemaWithDefaults instantiates a new UsersUpdateRequestSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersUpdateRequestSchemaWithDefaults() *UsersUpdateRequestSchema {
	this := UsersUpdateRequestSchema{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UsersUpdateRequestSchema) GetProfile() IdentitySourceUserProfileForUpsert {
	if o == nil || IsNil(o.Profile) {
		var ret IdentitySourceUserProfileForUpsert
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersUpdateRequestSchema) GetProfileOk() (*IdentitySourceUserProfileForUpsert, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UsersUpdateRequestSchema) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given IdentitySourceUserProfileForUpsert and assigns it to the Profile field.
func (o *UsersUpdateRequestSchema) SetProfile(v IdentitySourceUserProfileForUpsert) {
	o.Profile = &v
}

func (o UsersUpdateRequestSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersUpdateRequestSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UsersUpdateRequestSchema) UnmarshalJSON(data []byte) (err error) {
	varUsersUpdateRequestSchema := _UsersUpdateRequestSchema{}

	err = json.Unmarshal(data, &varUsersUpdateRequestSchema)

	if err != nil {
		return err
	}

	*o = UsersUpdateRequestSchema(varUsersUpdateRequestSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsersUpdateRequestSchema struct {
	value *UsersUpdateRequestSchema
	isSet bool
}

func (v NullableUsersUpdateRequestSchema) Get() *UsersUpdateRequestSchema {
	return v.value
}

func (v *NullableUsersUpdateRequestSchema) Set(val *UsersUpdateRequestSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersUpdateRequestSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersUpdateRequestSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersUpdateRequestSchema(val *UsersUpdateRequestSchema) *NullableUsersUpdateRequestSchema {
	return &NullableUsersUpdateRequestSchema{value: val, isSet: true}
}

func (v NullableUsersUpdateRequestSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersUpdateRequestSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
