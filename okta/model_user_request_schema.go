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

// checks if the UserRequestSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRequestSchema{}

// UserRequestSchema struct for UserRequestSchema
type UserRequestSchema struct {
	// The external ID of the user in the identity source
	ExternalId           *string                             `json:"externalId,omitempty"`
	Profile              *IdentitySourceUserProfileForUpsert `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserRequestSchema UserRequestSchema

// NewUserRequestSchema instantiates a new UserRequestSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRequestSchema() *UserRequestSchema {
	this := UserRequestSchema{}
	return &this
}

// NewUserRequestSchemaWithDefaults instantiates a new UserRequestSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRequestSchemaWithDefaults() *UserRequestSchema {
	this := UserRequestSchema{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *UserRequestSchema) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestSchema) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *UserRequestSchema) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *UserRequestSchema) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserRequestSchema) GetProfile() IdentitySourceUserProfileForUpsert {
	if o == nil || IsNil(o.Profile) {
		var ret IdentitySourceUserProfileForUpsert
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestSchema) GetProfileOk() (*IdentitySourceUserProfileForUpsert, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserRequestSchema) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given IdentitySourceUserProfileForUpsert and assigns it to the Profile field.
func (o *UserRequestSchema) SetProfile(v IdentitySourceUserProfileForUpsert) {
	o.Profile = &v
}

func (o UserRequestSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRequestSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserRequestSchema) UnmarshalJSON(data []byte) (err error) {
	varUserRequestSchema := _UserRequestSchema{}

	err = json.Unmarshal(data, &varUserRequestSchema)

	if err != nil {
		return err
	}

	*o = UserRequestSchema(varUserRequestSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserRequestSchema struct {
	value *UserRequestSchema
	isSet bool
}

func (v NullableUserRequestSchema) Get() *UserRequestSchema {
	return v.value
}

func (v *NullableUserRequestSchema) Set(val *UserRequestSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRequestSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRequestSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRequestSchema(val *UserRequestSchema) *NullableUserRequestSchema {
	return &NullableUserRequestSchema{value: val, isSet: true}
}

func (v NullableUserRequestSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRequestSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
