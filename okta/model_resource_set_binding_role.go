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

// ResourceSetBindingRole struct for ResourceSetBindingRole
type ResourceSetBindingRole struct {
	Id *string `json:"id,omitempty"`
	Links *ResourceSetBindingRoleLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetBindingRole ResourceSetBindingRole

// NewResourceSetBindingRole instantiates a new ResourceSetBindingRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetBindingRole() *ResourceSetBindingRole {
	this := ResourceSetBindingRole{}
	return &this
}

// NewResourceSetBindingRoleWithDefaults instantiates a new ResourceSetBindingRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetBindingRoleWithDefaults() *ResourceSetBindingRole {
	this := ResourceSetBindingRole{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceSetBindingRole) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingRole) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceSetBindingRole) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceSetBindingRole) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResourceSetBindingRole) GetLinks() ResourceSetBindingRoleLinks {
	if o == nil || o.Links == nil {
		var ret ResourceSetBindingRoleLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingRole) GetLinksOk() (*ResourceSetBindingRoleLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceSetBindingRole) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceSetBindingRoleLinks and assigns it to the Links field.
func (o *ResourceSetBindingRole) SetLinks(v ResourceSetBindingRoleLinks) {
	o.Links = &v
}

func (o ResourceSetBindingRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSetBindingRole) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSetBindingRole := _ResourceSetBindingRole{}

	err = json.Unmarshal(bytes, &varResourceSetBindingRole)
	if err == nil {
		*o = ResourceSetBindingRole(varResourceSetBindingRole)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSetBindingRole struct {
	value *ResourceSetBindingRole
	isSet bool
}

func (v NullableResourceSetBindingRole) Get() *ResourceSetBindingRole {
	return v.value
}

func (v *NullableResourceSetBindingRole) Set(val *ResourceSetBindingRole) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetBindingRole) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetBindingRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetBindingRole(val *ResourceSetBindingRole) *NullableResourceSetBindingRole {
	return &NullableResourceSetBindingRole{value: val, isSet: true}
}

func (v NullableResourceSetBindingRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetBindingRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

