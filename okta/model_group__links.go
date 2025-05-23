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

// GroupLinks struct for GroupLinks
type GroupLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Apps *HrefObject `json:"apps,omitempty"`
	Logo []HrefObject `json:"logo,omitempty"`
	Source *HrefObject `json:"source,omitempty"`
	Users *HrefObject `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupLinks GroupLinks

// NewGroupLinks instantiates a new GroupLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupLinks() *GroupLinks {
	this := GroupLinks{}
	return &this
}

// NewGroupLinksWithDefaults instantiates a new GroupLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupLinksWithDefaults() *GroupLinks {
	this := GroupLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *GroupLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *GroupLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *GroupLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *GroupLinks) GetApps() HrefObject {
	if o == nil || o.Apps == nil {
		var ret HrefObject
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLinks) GetAppsOk() (*HrefObject, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *GroupLinks) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given HrefObject and assigns it to the Apps field.
func (o *GroupLinks) SetApps(v HrefObject) {
	o.Apps = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *GroupLinks) GetLogo() []HrefObject {
	if o == nil || o.Logo == nil {
		var ret []HrefObject
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLinks) GetLogoOk() ([]HrefObject, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *GroupLinks) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given []HrefObject and assigns it to the Logo field.
func (o *GroupLinks) SetLogo(v []HrefObject) {
	o.Logo = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *GroupLinks) GetSource() HrefObject {
	if o == nil || o.Source == nil {
		var ret HrefObject
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLinks) GetSourceOk() (*HrefObject, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *GroupLinks) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given HrefObject and assigns it to the Source field.
func (o *GroupLinks) SetSource(v HrefObject) {
	o.Source = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *GroupLinks) GetUsers() HrefObject {
	if o == nil || o.Users == nil {
		var ret HrefObject
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLinks) GetUsersOk() (*HrefObject, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *GroupLinks) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given HrefObject and assigns it to the Users field.
func (o *GroupLinks) SetUsers(v HrefObject) {
	o.Users = &v
}

func (o GroupLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GroupLinks) UnmarshalJSON(bytes []byte) (err error) {
	varGroupLinks := _GroupLinks{}

	err = json.Unmarshal(bytes, &varGroupLinks)
	if err == nil {
		*o = GroupLinks(varGroupLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "apps")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "source")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableGroupLinks struct {
	value *GroupLinks
	isSet bool
}

func (v NullableGroupLinks) Get() *GroupLinks {
	return v.value
}

func (v *NullableGroupLinks) Set(val *GroupLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupLinks(val *GroupLinks) *NullableGroupLinks {
	return &NullableGroupLinks{value: val, isSet: true}
}

func (v NullableGroupLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

