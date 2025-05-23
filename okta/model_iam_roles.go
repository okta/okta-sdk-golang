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

// IamRoles struct for IamRoles
type IamRoles struct {
	Roles []IamRole `json:"roles,omitempty"`
	Links *LinksNext `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamRoles IamRoles

// NewIamRoles instantiates a new IamRoles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamRoles() *IamRoles {
	this := IamRoles{}
	return &this
}

// NewIamRolesWithDefaults instantiates a new IamRoles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamRolesWithDefaults() *IamRoles {
	this := IamRoles{}
	return &this
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *IamRoles) GetRoles() []IamRole {
	if o == nil || o.Roles == nil {
		var ret []IamRole
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRoles) GetRolesOk() ([]IamRole, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *IamRoles) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []IamRole and assigns it to the Roles field.
func (o *IamRoles) SetRoles(v []IamRole) {
	o.Roles = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *IamRoles) GetLinks() LinksNext {
	if o == nil || o.Links == nil {
		var ret LinksNext
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRoles) GetLinksOk() (*LinksNext, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *IamRoles) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksNext and assigns it to the Links field.
func (o *IamRoles) SetLinks(v LinksNext) {
	o.Links = &v
}

func (o IamRoles) MarshalJSON() ([]byte, error) {
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

func (o *IamRoles) UnmarshalJSON(bytes []byte) (err error) {
	varIamRoles := _IamRoles{}

	err = json.Unmarshal(bytes, &varIamRoles)
	if err == nil {
		*o = IamRoles(varIamRoles)
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

type NullableIamRoles struct {
	value *IamRoles
	isSet bool
}

func (v NullableIamRoles) Get() *IamRoles {
	return v.value
}

func (v *NullableIamRoles) Set(val *IamRoles) {
	v.value = val
	v.isSet = true
}

func (v NullableIamRoles) IsSet() bool {
	return v.isSet
}

func (v *NullableIamRoles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamRoles(val *IamRoles) *NullableIamRoles {
	return &NullableIamRoles{value: val, isSet: true}
}

func (v NullableIamRoles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamRoles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

