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

// LinksSelfAndFullUsersLifecycle struct for LinksSelfAndFullUsersLifecycle
type LinksSelfAndFullUsersLifecycle struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Activate *HrefObjectActivateLink `json:"activate,omitempty"`
	Deactivate *HrefObjectDeactivateLink `json:"deactivate,omitempty"`
	Suspend *HrefObjectSuspendLink `json:"suspend,omitempty"`
	Unsuspend *HrefObjectUnsuspendLink `json:"unsuspend,omitempty"`
	Users *HrefObject `json:"users,omitempty"`
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
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndFullUsersLifecycle) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *LinksSelfAndFullUsersLifecycle) HasSelf() bool {
	if o != nil && o.Self != nil {
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
	if o == nil || o.Activate == nil {
		var ret HrefObjectActivateLink
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndFullUsersLifecycle) GetActivateOk() (*HrefObjectActivateLink, bool) {
	if o == nil || o.Activate == nil {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *LinksSelfAndFullUsersLifecycle) HasActivate() bool {
	if o != nil && o.Activate != nil {
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
	if o == nil || o.Deactivate == nil {
		var ret HrefObjectDeactivateLink
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndFullUsersLifecycle) GetDeactivateOk() (*HrefObjectDeactivateLink, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *LinksSelfAndFullUsersLifecycle) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
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
	if o == nil || o.Suspend == nil {
		var ret HrefObjectSuspendLink
		return ret
	}
	return *o.Suspend
}

// GetSuspendOk returns a tuple with the Suspend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndFullUsersLifecycle) GetSuspendOk() (*HrefObjectSuspendLink, bool) {
	if o == nil || o.Suspend == nil {
		return nil, false
	}
	return o.Suspend, true
}

// HasSuspend returns a boolean if a field has been set.
func (o *LinksSelfAndFullUsersLifecycle) HasSuspend() bool {
	if o != nil && o.Suspend != nil {
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
	if o == nil || o.Unsuspend == nil {
		var ret HrefObjectUnsuspendLink
		return ret
	}
	return *o.Unsuspend
}

// GetUnsuspendOk returns a tuple with the Unsuspend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndFullUsersLifecycle) GetUnsuspendOk() (*HrefObjectUnsuspendLink, bool) {
	if o == nil || o.Unsuspend == nil {
		return nil, false
	}
	return o.Unsuspend, true
}

// HasUnsuspend returns a boolean if a field has been set.
func (o *LinksSelfAndFullUsersLifecycle) HasUnsuspend() bool {
	if o != nil && o.Unsuspend != nil {
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
	if o == nil || o.Users == nil {
		var ret HrefObject
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndFullUsersLifecycle) GetUsersOk() (*HrefObject, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *LinksSelfAndFullUsersLifecycle) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given HrefObject and assigns it to the Users field.
func (o *LinksSelfAndFullUsersLifecycle) SetUsers(v HrefObject) {
	o.Users = &v
}

func (o LinksSelfAndFullUsersLifecycle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Activate != nil {
		toSerialize["activate"] = o.Activate
	}
	if o.Deactivate != nil {
		toSerialize["deactivate"] = o.Deactivate
	}
	if o.Suspend != nil {
		toSerialize["suspend"] = o.Suspend
	}
	if o.Unsuspend != nil {
		toSerialize["unsuspend"] = o.Unsuspend
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksSelfAndFullUsersLifecycle) UnmarshalJSON(bytes []byte) (err error) {
	varLinksSelfAndFullUsersLifecycle := _LinksSelfAndFullUsersLifecycle{}

	err = json.Unmarshal(bytes, &varLinksSelfAndFullUsersLifecycle)
	if err == nil {
		*o = LinksSelfAndFullUsersLifecycle(varLinksSelfAndFullUsersLifecycle)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "activate")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "suspend")
		delete(additionalProperties, "unsuspend")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

