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
	"time"
)

// checks if the RoleTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleTarget{}

// RoleTarget struct for RoleTarget
type RoleTarget struct {
	// The assignment type of how the user receives this target
	AssignmentType *string `json:"assignmentType,omitempty"`
	// The expiry time stamp of the associated target. It's only included in the response if the associated target will expire.
	Expiration *time.Time `json:"expiration,omitempty"`
	// The [Okta Resource Name (ORN)](https://support.okta.com/help/s/article/understanding-okta-resource-name-orn) of the app target or group target
	Orn                  *string    `json:"orn,omitempty"`
	Links                *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleTarget RoleTarget

// NewRoleTarget instantiates a new RoleTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleTarget() *RoleTarget {
	this := RoleTarget{}
	return &this
}

// NewRoleTargetWithDefaults instantiates a new RoleTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleTargetWithDefaults() *RoleTarget {
	this := RoleTarget{}
	return &this
}

// GetAssignmentType returns the AssignmentType field value if set, zero value otherwise.
func (o *RoleTarget) GetAssignmentType() string {
	if o == nil || IsNil(o.AssignmentType) {
		var ret string
		return ret
	}
	return *o.AssignmentType
}

// GetAssignmentTypeOk returns a tuple with the AssignmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleTarget) GetAssignmentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssignmentType) {
		return nil, false
	}
	return o.AssignmentType, true
}

// HasAssignmentType returns a boolean if a field has been set.
func (o *RoleTarget) HasAssignmentType() bool {
	if o != nil && !IsNil(o.AssignmentType) {
		return true
	}

	return false
}

// SetAssignmentType gets a reference to the given string and assigns it to the AssignmentType field.
func (o *RoleTarget) SetAssignmentType(v string) {
	o.AssignmentType = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *RoleTarget) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration) {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleTarget) GetExpirationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *RoleTarget) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *RoleTarget) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *RoleTarget) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleTarget) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *RoleTarget) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *RoleTarget) SetOrn(v string) {
	o.Orn = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RoleTarget) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleTarget) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RoleTarget) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *RoleTarget) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o RoleTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssignmentType) {
		toSerialize["assignmentType"] = o.AssignmentType
	}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleTarget) UnmarshalJSON(data []byte) (err error) {
	varRoleTarget := _RoleTarget{}

	err = json.Unmarshal(data, &varRoleTarget)

	if err != nil {
		return err
	}

	*o = RoleTarget(varRoleTarget)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assignmentType")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleTarget struct {
	value *RoleTarget
	isSet bool
}

func (v NullableRoleTarget) Get() *RoleTarget {
	return v.value
}

func (v *NullableRoleTarget) Set(val *RoleTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleTarget(val *RoleTarget) *NullableRoleTarget {
	return &NullableRoleTarget{value: val, isSet: true}
}

func (v NullableRoleTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
