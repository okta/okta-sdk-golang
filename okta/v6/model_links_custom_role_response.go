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

// LinksCustomRoleResponse Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification. This object is used for dynamic discovery of related resources.
type LinksCustomRoleResponse struct {
	Assignee *HrefObjectAssigneeLink `json:"assignee,omitempty"`
	Member *HrefObjectMemberLink `json:"member,omitempty"`
	Permissions *HrefObjectPermissionsLink `json:"permissions,omitempty"`
	ResourceSet *HrefObjectResourceSetLink `json:"resource-set,omitempty"`
	Role *HrefObjectRoleLink `json:"role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksCustomRoleResponse LinksCustomRoleResponse

// NewLinksCustomRoleResponse instantiates a new LinksCustomRoleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksCustomRoleResponse() *LinksCustomRoleResponse {
	this := LinksCustomRoleResponse{}
	return &this
}

// NewLinksCustomRoleResponseWithDefaults instantiates a new LinksCustomRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksCustomRoleResponseWithDefaults() *LinksCustomRoleResponse {
	this := LinksCustomRoleResponse{}
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *LinksCustomRoleResponse) GetAssignee() HrefObjectAssigneeLink {
	if o == nil || o.Assignee == nil {
		var ret HrefObjectAssigneeLink
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksCustomRoleResponse) GetAssigneeOk() (*HrefObjectAssigneeLink, bool) {
	if o == nil || o.Assignee == nil {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *LinksCustomRoleResponse) HasAssignee() bool {
	if o != nil && o.Assignee != nil {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given HrefObjectAssigneeLink and assigns it to the Assignee field.
func (o *LinksCustomRoleResponse) SetAssignee(v HrefObjectAssigneeLink) {
	o.Assignee = &v
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *LinksCustomRoleResponse) GetMember() HrefObjectMemberLink {
	if o == nil || o.Member == nil {
		var ret HrefObjectMemberLink
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksCustomRoleResponse) GetMemberOk() (*HrefObjectMemberLink, bool) {
	if o == nil || o.Member == nil {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *LinksCustomRoleResponse) HasMember() bool {
	if o != nil && o.Member != nil {
		return true
	}

	return false
}

// SetMember gets a reference to the given HrefObjectMemberLink and assigns it to the Member field.
func (o *LinksCustomRoleResponse) SetMember(v HrefObjectMemberLink) {
	o.Member = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *LinksCustomRoleResponse) GetPermissions() HrefObjectPermissionsLink {
	if o == nil || o.Permissions == nil {
		var ret HrefObjectPermissionsLink
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksCustomRoleResponse) GetPermissionsOk() (*HrefObjectPermissionsLink, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *LinksCustomRoleResponse) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given HrefObjectPermissionsLink and assigns it to the Permissions field.
func (o *LinksCustomRoleResponse) SetPermissions(v HrefObjectPermissionsLink) {
	o.Permissions = &v
}

// GetResourceSet returns the ResourceSet field value if set, zero value otherwise.
func (o *LinksCustomRoleResponse) GetResourceSet() HrefObjectResourceSetLink {
	if o == nil || o.ResourceSet == nil {
		var ret HrefObjectResourceSetLink
		return ret
	}
	return *o.ResourceSet
}

// GetResourceSetOk returns a tuple with the ResourceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksCustomRoleResponse) GetResourceSetOk() (*HrefObjectResourceSetLink, bool) {
	if o == nil || o.ResourceSet == nil {
		return nil, false
	}
	return o.ResourceSet, true
}

// HasResourceSet returns a boolean if a field has been set.
func (o *LinksCustomRoleResponse) HasResourceSet() bool {
	if o != nil && o.ResourceSet != nil {
		return true
	}

	return false
}

// SetResourceSet gets a reference to the given HrefObjectResourceSetLink and assigns it to the ResourceSet field.
func (o *LinksCustomRoleResponse) SetResourceSet(v HrefObjectResourceSetLink) {
	o.ResourceSet = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *LinksCustomRoleResponse) GetRole() HrefObjectRoleLink {
	if o == nil || o.Role == nil {
		var ret HrefObjectRoleLink
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksCustomRoleResponse) GetRoleOk() (*HrefObjectRoleLink, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *LinksCustomRoleResponse) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given HrefObjectRoleLink and assigns it to the Role field.
func (o *LinksCustomRoleResponse) SetRole(v HrefObjectRoleLink) {
	o.Role = &v
}

func (o LinksCustomRoleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assignee != nil {
		toSerialize["assignee"] = o.Assignee
	}
	if o.Member != nil {
		toSerialize["member"] = o.Member
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.ResourceSet != nil {
		toSerialize["resource-set"] = o.ResourceSet
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksCustomRoleResponse) UnmarshalJSON(bytes []byte) (err error) {
	varLinksCustomRoleResponse := _LinksCustomRoleResponse{}

	err = json.Unmarshal(bytes, &varLinksCustomRoleResponse)
	if err == nil {
		*o = LinksCustomRoleResponse(varLinksCustomRoleResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "assignee")
		delete(additionalProperties, "member")
		delete(additionalProperties, "permissions")
		delete(additionalProperties, "resource-set")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinksCustomRoleResponse struct {
	value *LinksCustomRoleResponse
	isSet bool
}

func (v NullableLinksCustomRoleResponse) Get() *LinksCustomRoleResponse {
	return v.value
}

func (v *NullableLinksCustomRoleResponse) Set(val *LinksCustomRoleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksCustomRoleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksCustomRoleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksCustomRoleResponse(val *LinksCustomRoleResponse) *NullableLinksCustomRoleResponse {
	return &NullableLinksCustomRoleResponse{value: val, isSet: true}
}

func (v NullableLinksCustomRoleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksCustomRoleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

