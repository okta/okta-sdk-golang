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

// checks if the GroupsResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsResponseSchema{}

// GroupsResponseSchema struct for GroupsResponseSchema
type GroupsResponseSchema struct {
	// The external ID of the identity source group
	ExternalId *string `json:"externalId,omitempty"`
	// The Okta group ID of the identity source group
	Id                   *string                      `json:"id,omitempty"`
	Profile              *GroupsResponseSchemaProfile `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupsResponseSchema GroupsResponseSchema

// NewGroupsResponseSchema instantiates a new GroupsResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsResponseSchema() *GroupsResponseSchema {
	this := GroupsResponseSchema{}
	return &this
}

// NewGroupsResponseSchemaWithDefaults instantiates a new GroupsResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsResponseSchemaWithDefaults() *GroupsResponseSchema {
	this := GroupsResponseSchema{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *GroupsResponseSchema) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsResponseSchema) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *GroupsResponseSchema) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *GroupsResponseSchema) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupsResponseSchema) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsResponseSchema) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupsResponseSchema) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupsResponseSchema) SetId(v string) {
	o.Id = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *GroupsResponseSchema) GetProfile() GroupsResponseSchemaProfile {
	if o == nil || IsNil(o.Profile) {
		var ret GroupsResponseSchemaProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsResponseSchema) GetProfileOk() (*GroupsResponseSchemaProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *GroupsResponseSchema) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given GroupsResponseSchemaProfile and assigns it to the Profile field.
func (o *GroupsResponseSchema) SetProfile(v GroupsResponseSchemaProfile) {
	o.Profile = &v
}

func (o GroupsResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupsResponseSchema) UnmarshalJSON(data []byte) (err error) {
	varGroupsResponseSchema := _GroupsResponseSchema{}

	err = json.Unmarshal(data, &varGroupsResponseSchema)

	if err != nil {
		return err
	}

	*o = GroupsResponseSchema(varGroupsResponseSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupsResponseSchema struct {
	value *GroupsResponseSchema
	isSet bool
}

func (v NullableGroupsResponseSchema) Get() *GroupsResponseSchema {
	return v.value
}

func (v *NullableGroupsResponseSchema) Set(val *GroupsResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsResponseSchema(val *GroupsResponseSchema) *NullableGroupsResponseSchema {
	return &NullableGroupsResponseSchema{value: val, isSet: true}
}

func (v NullableGroupsResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
