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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// RoleGovernanceResource The resource of a grant
type RoleGovernanceResource struct {
	// The resource name
	Label *string `json:"label,omitempty"`
	// The resources id
	Resource *string `json:"resource,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleGovernanceResource RoleGovernanceResource

// NewRoleGovernanceResource instantiates a new RoleGovernanceResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleGovernanceResource() *RoleGovernanceResource {
	this := RoleGovernanceResource{}
	return &this
}

// NewRoleGovernanceResourceWithDefaults instantiates a new RoleGovernanceResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleGovernanceResourceWithDefaults() *RoleGovernanceResource {
	this := RoleGovernanceResource{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *RoleGovernanceResource) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGovernanceResource) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *RoleGovernanceResource) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *RoleGovernanceResource) SetLabel(v string) {
	o.Label = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *RoleGovernanceResource) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGovernanceResource) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *RoleGovernanceResource) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *RoleGovernanceResource) SetResource(v string) {
	o.Resource = &v
}

func (o RoleGovernanceResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleGovernanceResource) UnmarshalJSON(bytes []byte) (err error) {
	varRoleGovernanceResource := _RoleGovernanceResource{}

	err = json.Unmarshal(bytes, &varRoleGovernanceResource)
	if err == nil {
		*o = RoleGovernanceResource(varRoleGovernanceResource)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "resource")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRoleGovernanceResource struct {
	value *RoleGovernanceResource
	isSet bool
}

func (v NullableRoleGovernanceResource) Get() *RoleGovernanceResource {
	return v.value
}

func (v *NullableRoleGovernanceResource) Set(val *RoleGovernanceResource) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleGovernanceResource) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleGovernanceResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleGovernanceResource(val *RoleGovernanceResource) *NullableRoleGovernanceResource {
	return &NullableRoleGovernanceResource{value: val, isSet: true}
}

func (v NullableRoleGovernanceResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleGovernanceResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

