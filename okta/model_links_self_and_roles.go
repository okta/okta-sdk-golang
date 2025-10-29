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

// checks if the LinksSelfAndRoles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksSelfAndRoles{}

// LinksSelfAndRoles struct for LinksSelfAndRoles
type LinksSelfAndRoles struct {
	Self                 *HrefObjectSelfLink `json:"self,omitempty"`
	Roles                *HrefObjectRoleLink `json:"roles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksSelfAndRoles LinksSelfAndRoles

// NewLinksSelfAndRoles instantiates a new LinksSelfAndRoles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksSelfAndRoles() *LinksSelfAndRoles {
	this := LinksSelfAndRoles{}
	return &this
}

// NewLinksSelfAndRolesWithDefaults instantiates a new LinksSelfAndRoles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksSelfAndRolesWithDefaults() *LinksSelfAndRoles {
	this := LinksSelfAndRoles{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *LinksSelfAndRoles) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndRoles) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *LinksSelfAndRoles) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *LinksSelfAndRoles) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *LinksSelfAndRoles) GetRoles() HrefObjectRoleLink {
	if o == nil || IsNil(o.Roles) {
		var ret HrefObjectRoleLink
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndRoles) GetRolesOk() (*HrefObjectRoleLink, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *LinksSelfAndRoles) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given HrefObjectRoleLink and assigns it to the Roles field.
func (o *LinksSelfAndRoles) SetRoles(v HrefObjectRoleLink) {
	o.Roles = &v
}

func (o LinksSelfAndRoles) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksSelfAndRoles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksSelfAndRoles) UnmarshalJSON(data []byte) (err error) {
	varLinksSelfAndRoles := _LinksSelfAndRoles{}

	err = json.Unmarshal(data, &varLinksSelfAndRoles)

	if err != nil {
		return err
	}

	*o = LinksSelfAndRoles(varLinksSelfAndRoles)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "roles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinksSelfAndRoles struct {
	value *LinksSelfAndRoles
	isSet bool
}

func (v NullableLinksSelfAndRoles) Get() *LinksSelfAndRoles {
	return v.value
}

func (v *NullableLinksSelfAndRoles) Set(val *LinksSelfAndRoles) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksSelfAndRoles) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksSelfAndRoles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksSelfAndRoles(val *LinksSelfAndRoles) *NullableLinksSelfAndRoles {
	return &NullableLinksSelfAndRoles{value: val, isSet: true}
}

func (v NullableLinksSelfAndRoles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksSelfAndRoles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
