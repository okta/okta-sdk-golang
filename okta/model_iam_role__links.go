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

// IamRoleLinks struct for IamRoleLinks
type IamRoleLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Permissions *HrefObject `json:"permissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamRoleLinks IamRoleLinks

// NewIamRoleLinks instantiates a new IamRoleLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamRoleLinks() *IamRoleLinks {
	this := IamRoleLinks{}
	return &this
}

// NewIamRoleLinksWithDefaults instantiates a new IamRoleLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamRoleLinksWithDefaults() *IamRoleLinks {
	this := IamRoleLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *IamRoleLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRoleLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *IamRoleLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *IamRoleLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *IamRoleLinks) GetPermissions() HrefObject {
	if o == nil || o.Permissions == nil {
		var ret HrefObject
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRoleLinks) GetPermissionsOk() (*HrefObject, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *IamRoleLinks) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given HrefObject and assigns it to the Permissions field.
func (o *IamRoleLinks) SetPermissions(v HrefObject) {
	o.Permissions = &v
}

func (o IamRoleLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamRoleLinks) UnmarshalJSON(bytes []byte) (err error) {
	varIamRoleLinks := _IamRoleLinks{}

	err = json.Unmarshal(bytes, &varIamRoleLinks)
	if err == nil {
		*o = IamRoleLinks(varIamRoleLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIamRoleLinks struct {
	value *IamRoleLinks
	isSet bool
}

func (v NullableIamRoleLinks) Get() *IamRoleLinks {
	return v.value
}

func (v *NullableIamRoleLinks) Set(val *IamRoleLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableIamRoleLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableIamRoleLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamRoleLinks(val *IamRoleLinks) *NullableIamRoleLinks {
	return &NullableIamRoleLinks{value: val, isSet: true}
}

func (v NullableIamRoleLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamRoleLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

