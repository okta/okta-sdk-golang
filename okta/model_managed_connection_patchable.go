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

// checks if the ManagedConnectionPatchable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedConnectionPatchable{}

// ManagedConnectionPatchable Update an existing managed connection. This request only applies to identity assertion connections; attempting to patch other connection types results in an error.
type ManagedConnectionPatchable struct {
	// Determines how Okta evaluates requested scopes for the connection.
	ScopeCondition *string `json:"scopeCondition,omitempty"`
	// Optional array of scopes. Required when scopeCondition is INCLUDE_ONLY or EXCLUDE. Not used when scopeCondition is ALL_SCOPES.
	Scopes []string `json:"scopes,omitempty"`
}

// NewManagedConnectionPatchable instantiates a new ManagedConnectionPatchable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedConnectionPatchable() *ManagedConnectionPatchable {
	this := ManagedConnectionPatchable{}
	return &this
}

// NewManagedConnectionPatchableWithDefaults instantiates a new ManagedConnectionPatchable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedConnectionPatchableWithDefaults() *ManagedConnectionPatchable {
	this := ManagedConnectionPatchable{}
	return &this
}

// GetScopeCondition returns the ScopeCondition field value if set, zero value otherwise.
func (o *ManagedConnectionPatchable) GetScopeCondition() string {
	if o == nil || IsNil(o.ScopeCondition) {
		var ret string
		return ret
	}
	return *o.ScopeCondition
}

// GetScopeConditionOk returns a tuple with the ScopeCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedConnectionPatchable) GetScopeConditionOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeCondition) {
		return nil, false
	}
	return o.ScopeCondition, true
}

// HasScopeCondition returns a boolean if a field has been set.
func (o *ManagedConnectionPatchable) HasScopeCondition() bool {
	if o != nil && !IsNil(o.ScopeCondition) {
		return true
	}

	return false
}

// SetScopeCondition gets a reference to the given string and assigns it to the ScopeCondition field.
func (o *ManagedConnectionPatchable) SetScopeCondition(v string) {
	o.ScopeCondition = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ManagedConnectionPatchable) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedConnectionPatchable) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ManagedConnectionPatchable) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ManagedConnectionPatchable) SetScopes(v []string) {
	o.Scopes = v
}

func (o ManagedConnectionPatchable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedConnectionPatchable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScopeCondition) {
		toSerialize["scopeCondition"] = o.ScopeCondition
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullableManagedConnectionPatchable struct {
	value *ManagedConnectionPatchable
	isSet bool
}

func (v NullableManagedConnectionPatchable) Get() *ManagedConnectionPatchable {
	return v.value
}

func (v *NullableManagedConnectionPatchable) Set(val *ManagedConnectionPatchable) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedConnectionPatchable) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedConnectionPatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedConnectionPatchable(val *ManagedConnectionPatchable) *NullableManagedConnectionPatchable {
	return &NullableManagedConnectionPatchable{value: val, isSet: true}
}

func (v NullableManagedConnectionPatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedConnectionPatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
