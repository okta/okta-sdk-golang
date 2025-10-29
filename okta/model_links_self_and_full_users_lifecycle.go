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

// checks if the LinksSelfAndFullUsersLifecycle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksSelfAndFullUsersLifecycle{}

// LinksSelfAndFullUsersLifecycle struct for LinksSelfAndFullUsersLifecycle
type LinksSelfAndFullUsersLifecycle struct {
	Self       *HrefObjectSelfLink       `json:"self,omitempty"`
	Activate   *HrefObjectActivateLink   `json:"activate,omitempty"`
	Deactivate *HrefObjectDeactivateLink `json:"deactivate,omitempty"`
	Suspend    *HrefObjectSuspendLink    `json:"suspend,omitempty"`
	Unsuspend  *HrefObjectUnsuspendLink  `json:"unsuspend,omitempty"`
	// Link to device users
	Users                *HrefObject `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksSelfAndFullUsersLifecycle LinksSelfAndFullUsersLifecycle

// NewLinksSelfAndFullUsersLifecycle instantiates a new LinksSelfAndFullUsersLifecycle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksSelfAndFullUsersLifecycle() *LinksSelfAndFullUsersLifecycle {
	this := LinksSelfAndFullUsersLifecycle{}
	return &this
}

// NewLinksSelfAndFullUsersLifecycleWithDefaults instantiates a new LinksSelfAndFullUsersLifecycle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksSelfAndFullUsersLifecycleWithDefaults() *LinksSelfAndFullUsersLifecycle {
	this := LinksSelfAndFullUsersLifecycle{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *LinksSelfAndFullUsersLifecycle) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndFullUsersLifecycle) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *LinksSelfAndFullUsersLifecycle) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *LinksSelfAndFullUsersLifecycle) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *LinksSelfAndFullUsersLifecycle) GetActivate() HrefObjectActivateLink {
	if o == nil || IsNil(o.Activate) {
		var ret HrefObjectActivateLink
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndFullUsersLifecycle) GetActivateOk() (*HrefObjectActivateLink, bool) {
	if o == nil || IsNil(o.Activate) {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *LinksSelfAndFullUsersLifecycle) HasActivate() bool {
	if o != nil && !IsNil(o.Activate) {
		return true
	}

	return false
}

// SetActivate gets a reference to the given HrefObjectActivateLink and assigns it to the Activate field.
func (o *LinksSelfAndFullUsersLifecycle) SetActivate(v HrefObjectActivateLink) {
	o.Activate = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *LinksSelfAndFullUsersLifecycle) GetDeactivate() HrefObjectDeactivateLink {
	if o == nil || IsNil(o.Deactivate) {
		var ret HrefObjectDeactivateLink
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndFullUsersLifecycle) GetDeactivateOk() (*HrefObjectDeactivateLink, bool) {
	if o == nil || IsNil(o.Deactivate) {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *LinksSelfAndFullUsersLifecycle) HasDeactivate() bool {
	if o != nil && !IsNil(o.Deactivate) {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObjectDeactivateLink and assigns it to the Deactivate field.
func (o *LinksSelfAndFullUsersLifecycle) SetDeactivate(v HrefObjectDeactivateLink) {
	o.Deactivate = &v
}

// GetSuspend returns the Suspend field value if set, zero value otherwise.
func (o *LinksSelfAndFullUsersLifecycle) GetSuspend() HrefObjectSuspendLink {
	if o == nil || IsNil(o.Suspend) {
		var ret HrefObjectSuspendLink
		return ret
	}
	return *o.Suspend
}

// GetSuspendOk returns a tuple with the Suspend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndFullUsersLifecycle) GetSuspendOk() (*HrefObjectSuspendLink, bool) {
	if o == nil || IsNil(o.Suspend) {
		return nil, false
	}
	return o.Suspend, true
}

// HasSuspend returns a boolean if a field has been set.
func (o *LinksSelfAndFullUsersLifecycle) HasSuspend() bool {
	if o != nil && !IsNil(o.Suspend) {
		return true
	}

	return false
}

// SetSuspend gets a reference to the given HrefObjectSuspendLink and assigns it to the Suspend field.
func (o *LinksSelfAndFullUsersLifecycle) SetSuspend(v HrefObjectSuspendLink) {
	o.Suspend = &v
}

// GetUnsuspend returns the Unsuspend field value if set, zero value otherwise.
func (o *LinksSelfAndFullUsersLifecycle) GetUnsuspend() HrefObjectUnsuspendLink {
	if o == nil || IsNil(o.Unsuspend) {
		var ret HrefObjectUnsuspendLink
		return ret
	}
	return *o.Unsuspend
}

// GetUnsuspendOk returns a tuple with the Unsuspend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndFullUsersLifecycle) GetUnsuspendOk() (*HrefObjectUnsuspendLink, bool) {
	if o == nil || IsNil(o.Unsuspend) {
		return nil, false
	}
	return o.Unsuspend, true
}

// HasUnsuspend returns a boolean if a field has been set.
func (o *LinksSelfAndFullUsersLifecycle) HasUnsuspend() bool {
	if o != nil && !IsNil(o.Unsuspend) {
		return true
	}

	return false
}

// SetUnsuspend gets a reference to the given HrefObjectUnsuspendLink and assigns it to the Unsuspend field.
func (o *LinksSelfAndFullUsersLifecycle) SetUnsuspend(v HrefObjectUnsuspendLink) {
	o.Unsuspend = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *LinksSelfAndFullUsersLifecycle) GetUsers() HrefObject {
	if o == nil || IsNil(o.Users) {
		var ret HrefObject
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndFullUsersLifecycle) GetUsersOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *LinksSelfAndFullUsersLifecycle) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given HrefObject and assigns it to the Users field.
func (o *LinksSelfAndFullUsersLifecycle) SetUsers(v HrefObject) {
	o.Users = &v
}

func (o LinksSelfAndFullUsersLifecycle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksSelfAndFullUsersLifecycle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Activate) {
		toSerialize["activate"] = o.Activate
	}
	if !IsNil(o.Deactivate) {
		toSerialize["deactivate"] = o.Deactivate
	}
	if !IsNil(o.Suspend) {
		toSerialize["suspend"] = o.Suspend
	}
	if !IsNil(o.Unsuspend) {
		toSerialize["unsuspend"] = o.Unsuspend
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksSelfAndFullUsersLifecycle) UnmarshalJSON(data []byte) (err error) {
	varLinksSelfAndFullUsersLifecycle := _LinksSelfAndFullUsersLifecycle{}

	err = json.Unmarshal(data, &varLinksSelfAndFullUsersLifecycle)

	if err != nil {
		return err
	}

	*o = LinksSelfAndFullUsersLifecycle(varLinksSelfAndFullUsersLifecycle)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "activate")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "suspend")
		delete(additionalProperties, "unsuspend")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinksSelfAndFullUsersLifecycle struct {
	value *LinksSelfAndFullUsersLifecycle
	isSet bool
}

func (v NullableLinksSelfAndFullUsersLifecycle) Get() *LinksSelfAndFullUsersLifecycle {
	return v.value
}

func (v *NullableLinksSelfAndFullUsersLifecycle) Set(val *LinksSelfAndFullUsersLifecycle) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksSelfAndFullUsersLifecycle) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksSelfAndFullUsersLifecycle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksSelfAndFullUsersLifecycle(val *LinksSelfAndFullUsersLifecycle) *NullableLinksSelfAndFullUsersLifecycle {
	return &NullableLinksSelfAndFullUsersLifecycle{value: val, isSet: true}
}

func (v NullableLinksSelfAndFullUsersLifecycle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksSelfAndFullUsersLifecycle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
