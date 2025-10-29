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

// checks if the Permissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Permissions{}

// Permissions Permissions assigned to the role
type Permissions struct {
	// Array of permissions assigned to the role. See [Permissions](/openapi/okta-management/guides/permissions).
	Permissions          []Permission `json:"permissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Permissions Permissions

// NewPermissions instantiates a new Permissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissions() *Permissions {
	this := Permissions{}
	return &this
}

// NewPermissionsWithDefaults instantiates a new Permissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsWithDefaults() *Permissions {
	this := Permissions{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *Permissions) GetPermissions() []Permission {
	if o == nil || IsNil(o.Permissions) {
		var ret []Permission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetPermissionsOk() ([]Permission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *Permissions) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []Permission and assigns it to the Permissions field.
func (o *Permissions) SetPermissions(v []Permission) {
	o.Permissions = v
}

func (o Permissions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Permissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Permissions) UnmarshalJSON(data []byte) (err error) {
	varPermissions := _Permissions{}

	err = json.Unmarshal(data, &varPermissions)

	if err != nil {
		return err
	}

	*o = Permissions(varPermissions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePermissions struct {
	value *Permissions
	isSet bool
}

func (v NullablePermissions) Get() *Permissions {
	return v.value
}

func (v *NullablePermissions) Set(val *Permissions) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissions(val *Permissions) *NullablePermissions {
	return &NullablePermissions{value: val, isSet: true}
}

func (v NullablePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
