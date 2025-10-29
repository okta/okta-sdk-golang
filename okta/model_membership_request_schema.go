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

// checks if the MembershipRequestSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembershipRequestSchema{}

// MembershipRequestSchema struct for MembershipRequestSchema
type MembershipRequestSchema struct {
	// The external ID of the user to be added as a member of the group in Okta
	MemberExternalId     *string `json:"memberExternalId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MembershipRequestSchema MembershipRequestSchema

// NewMembershipRequestSchema instantiates a new MembershipRequestSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipRequestSchema() *MembershipRequestSchema {
	this := MembershipRequestSchema{}
	return &this
}

// NewMembershipRequestSchemaWithDefaults instantiates a new MembershipRequestSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipRequestSchemaWithDefaults() *MembershipRequestSchema {
	this := MembershipRequestSchema{}
	return &this
}

// GetMemberExternalId returns the MemberExternalId field value if set, zero value otherwise.
func (o *MembershipRequestSchema) GetMemberExternalId() string {
	if o == nil || IsNil(o.MemberExternalId) {
		var ret string
		return ret
	}
	return *o.MemberExternalId
}

// GetMemberExternalIdOk returns a tuple with the MemberExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipRequestSchema) GetMemberExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.MemberExternalId) {
		return nil, false
	}
	return o.MemberExternalId, true
}

// HasMemberExternalId returns a boolean if a field has been set.
func (o *MembershipRequestSchema) HasMemberExternalId() bool {
	if o != nil && !IsNil(o.MemberExternalId) {
		return true
	}

	return false
}

// SetMemberExternalId gets a reference to the given string and assigns it to the MemberExternalId field.
func (o *MembershipRequestSchema) SetMemberExternalId(v string) {
	o.MemberExternalId = &v
}

func (o MembershipRequestSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembershipRequestSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MemberExternalId) {
		toSerialize["memberExternalId"] = o.MemberExternalId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MembershipRequestSchema) UnmarshalJSON(data []byte) (err error) {
	varMembershipRequestSchema := _MembershipRequestSchema{}

	err = json.Unmarshal(data, &varMembershipRequestSchema)

	if err != nil {
		return err
	}

	*o = MembershipRequestSchema(varMembershipRequestSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "memberExternalId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMembershipRequestSchema struct {
	value *MembershipRequestSchema
	isSet bool
}

func (v NullableMembershipRequestSchema) Get() *MembershipRequestSchema {
	return v.value
}

func (v *NullableMembershipRequestSchema) Set(val *MembershipRequestSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipRequestSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipRequestSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipRequestSchema(val *MembershipRequestSchema) *NullableMembershipRequestSchema {
	return &NullableMembershipRequestSchema{value: val, isSet: true}
}

func (v NullableMembershipRequestSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipRequestSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
