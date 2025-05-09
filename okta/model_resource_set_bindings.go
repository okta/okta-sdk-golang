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

// ResourceSetBindings struct for ResourceSetBindings
type ResourceSetBindings struct {
	Roles []ResourceSetBindingRole `json:"roles,omitempty"`
	Links *ResourceSetBindingResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetBindings ResourceSetBindings

// NewResourceSetBindings instantiates a new ResourceSetBindings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetBindings() *ResourceSetBindings {
	this := ResourceSetBindings{}
	return &this
}

// NewResourceSetBindingsWithDefaults instantiates a new ResourceSetBindings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetBindingsWithDefaults() *ResourceSetBindings {
	this := ResourceSetBindings{}
	return &this
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ResourceSetBindings) GetRoles() []ResourceSetBindingRole {
	if o == nil || o.Roles == nil {
		var ret []ResourceSetBindingRole
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindings) GetRolesOk() ([]ResourceSetBindingRole, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ResourceSetBindings) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []ResourceSetBindingRole and assigns it to the Roles field.
func (o *ResourceSetBindings) SetRoles(v []ResourceSetBindingRole) {
	o.Roles = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResourceSetBindings) GetLinks() ResourceSetBindingResponseLinks {
	if o == nil || o.Links == nil {
		var ret ResourceSetBindingResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindings) GetLinksOk() (*ResourceSetBindingResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceSetBindings) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceSetBindingResponseLinks and assigns it to the Links field.
func (o *ResourceSetBindings) SetLinks(v ResourceSetBindingResponseLinks) {
	o.Links = &v
}

func (o ResourceSetBindings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSetBindings) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSetBindings := _ResourceSetBindings{}

	err = json.Unmarshal(bytes, &varResourceSetBindings)
	if err == nil {
		*o = ResourceSetBindings(varResourceSetBindings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "roles")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSetBindings struct {
	value *ResourceSetBindings
	isSet bool
}

func (v NullableResourceSetBindings) Get() *ResourceSetBindings {
	return v.value
}

func (v *NullableResourceSetBindings) Set(val *ResourceSetBindings) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetBindings) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetBindings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetBindings(val *ResourceSetBindings) *NullableResourceSetBindings {
	return &NullableResourceSetBindings{value: val, isSet: true}
}

func (v NullableResourceSetBindings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetBindings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

