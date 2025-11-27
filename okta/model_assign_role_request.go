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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the AssignRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignRoleRequest{}

// AssignRoleRequest struct for AssignRoleRequest
type AssignRoleRequest struct {
	// | Role type                    | Description                                                 | |------------------------------|-------------------------------------------------------------| | ACCESS_CERTIFICATIONS_ADMIN  | Access Certifications Administrator IAM-based standard role | | ACCESS_REQUESTS_ADMIN        | Access Requests Administrator IAM-based standard role       | | API_ACCESS_MANAGEMENT_ADMIN  | Access Management Administrator standard role               | | APP_ADMIN                    | Application Administrator standard role                     | | CUSTOM                       | Custom admin role                                           | | GROUP_MEMBERSHIP_ADMIN       | Group Membership Administrator standard role                | | HELP_DESK_ADMIN              | Help Desk Administrator standard role                       | | ORG_ADMIN                    | Organizational Administrator standard role                  | | READ_ONLY_ADMIN              | Read-Only Administrator standard role                       | | REPORT_ADMIN                 | Report Administrator standard role                          | | SUPER_ADMIN                  | Super Administrator standard role                           | | USER_ADMIN                   | User Administrator standard role                            | | WORKFLOWS_ADMIN              | Workflows Administrator IAM-based standard role             |
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssignRoleRequest AssignRoleRequest

// NewAssignRoleRequest instantiates a new AssignRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignRoleRequest() *AssignRoleRequest {
	this := AssignRoleRequest{}
	return &this
}

// NewAssignRoleRequestWithDefaults instantiates a new AssignRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignRoleRequestWithDefaults() *AssignRoleRequest {
	this := AssignRoleRequest{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AssignRoleRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignRoleRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AssignRoleRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AssignRoleRequest) SetType(v string) {
	o.Type = &v
}

func (o AssignRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssignRoleRequest) UnmarshalJSON(data []byte) (err error) {
	varAssignRoleRequest := _AssignRoleRequest{}

	err = json.Unmarshal(data, &varAssignRoleRequest)

	if err != nil {
		return err
	}

	*o = AssignRoleRequest(varAssignRoleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssignRoleRequest struct {
	value *AssignRoleRequest
	isSet bool
}

func (v NullableAssignRoleRequest) Get() *AssignRoleRequest {
	return v.value
}

func (v *NullableAssignRoleRequest) Set(val *AssignRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignRoleRequest(val *AssignRoleRequest) *NullableAssignRoleRequest {
	return &NullableAssignRoleRequest{value: val, isSet: true}
}

func (v NullableAssignRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
