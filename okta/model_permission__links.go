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

// PermissionLinks struct for PermissionLinks
type PermissionLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Role *HrefObject `json:"role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PermissionLinks PermissionLinks

// NewPermissionLinks instantiates a new PermissionLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionLinks() *PermissionLinks {
	this := PermissionLinks{}
	return &this
}

// NewPermissionLinksWithDefaults instantiates a new PermissionLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionLinksWithDefaults() *PermissionLinks {
	this := PermissionLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *PermissionLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *PermissionLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *PermissionLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *PermissionLinks) GetRole() HrefObject {
	if o == nil || o.Role == nil {
		var ret HrefObject
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionLinks) GetRoleOk() (*HrefObject, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *PermissionLinks) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given HrefObject and assigns it to the Role field.
func (o *PermissionLinks) SetRole(v HrefObject) {
	o.Role = &v
}

func (o PermissionLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PermissionLinks) UnmarshalJSON(bytes []byte) (err error) {
	varPermissionLinks := _PermissionLinks{}

	err = json.Unmarshal(bytes, &varPermissionLinks)
	if err == nil {
		*o = PermissionLinks(varPermissionLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePermissionLinks struct {
	value *PermissionLinks
	isSet bool
}

func (v NullablePermissionLinks) Get() *PermissionLinks {
	return v.value
}

func (v *NullablePermissionLinks) Set(val *PermissionLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionLinks(val *PermissionLinks) *NullablePermissionLinks {
	return &NullablePermissionLinks{value: val, isSet: true}
}

func (v NullablePermissionLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

