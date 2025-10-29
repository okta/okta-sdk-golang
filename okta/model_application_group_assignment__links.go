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

// checks if the ApplicationGroupAssignmentLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationGroupAssignmentLinks{}

// ApplicationGroupAssignmentLinks struct for ApplicationGroupAssignmentLinks
type ApplicationGroupAssignmentLinks struct {
	Self                 *HrefObjectSelfLink  `json:"self,omitempty"`
	App                  *HrefObjectAppLink   `json:"app,omitempty"`
	Group                *HrefObjectGroupLink `json:"group,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationGroupAssignmentLinks ApplicationGroupAssignmentLinks

// NewApplicationGroupAssignmentLinks instantiates a new ApplicationGroupAssignmentLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationGroupAssignmentLinks() *ApplicationGroupAssignmentLinks {
	this := ApplicationGroupAssignmentLinks{}
	return &this
}

// NewApplicationGroupAssignmentLinksWithDefaults instantiates a new ApplicationGroupAssignmentLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationGroupAssignmentLinksWithDefaults() *ApplicationGroupAssignmentLinks {
	this := ApplicationGroupAssignmentLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ApplicationGroupAssignmentLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGroupAssignmentLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ApplicationGroupAssignmentLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *ApplicationGroupAssignmentLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *ApplicationGroupAssignmentLinks) GetApp() HrefObjectAppLink {
	if o == nil || IsNil(o.App) {
		var ret HrefObjectAppLink
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGroupAssignmentLinks) GetAppOk() (*HrefObjectAppLink, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *ApplicationGroupAssignmentLinks) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given HrefObjectAppLink and assigns it to the App field.
func (o *ApplicationGroupAssignmentLinks) SetApp(v HrefObjectAppLink) {
	o.App = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ApplicationGroupAssignmentLinks) GetGroup() HrefObjectGroupLink {
	if o == nil || IsNil(o.Group) {
		var ret HrefObjectGroupLink
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGroupAssignmentLinks) GetGroupOk() (*HrefObjectGroupLink, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ApplicationGroupAssignmentLinks) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given HrefObjectGroupLink and assigns it to the Group field.
func (o *ApplicationGroupAssignmentLinks) SetGroup(v HrefObjectGroupLink) {
	o.Group = &v
}

func (o ApplicationGroupAssignmentLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationGroupAssignmentLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationGroupAssignmentLinks) UnmarshalJSON(data []byte) (err error) {
	varApplicationGroupAssignmentLinks := _ApplicationGroupAssignmentLinks{}

	err = json.Unmarshal(data, &varApplicationGroupAssignmentLinks)

	if err != nil {
		return err
	}

	*o = ApplicationGroupAssignmentLinks(varApplicationGroupAssignmentLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "app")
		delete(additionalProperties, "group")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationGroupAssignmentLinks struct {
	value *ApplicationGroupAssignmentLinks
	isSet bool
}

func (v NullableApplicationGroupAssignmentLinks) Get() *ApplicationGroupAssignmentLinks {
	return v.value
}

func (v *NullableApplicationGroupAssignmentLinks) Set(val *ApplicationGroupAssignmentLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationGroupAssignmentLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationGroupAssignmentLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationGroupAssignmentLinks(val *ApplicationGroupAssignmentLinks) *NullableApplicationGroupAssignmentLinks {
	return &NullableApplicationGroupAssignmentLinks{value: val, isSet: true}
}

func (v NullableApplicationGroupAssignmentLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationGroupAssignmentLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
