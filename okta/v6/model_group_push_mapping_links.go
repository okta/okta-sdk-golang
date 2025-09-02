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

// GroupPushMappingLinks Discoverable resources related to the group push mapping
type GroupPushMappingLinks struct {
	App *HrefObjectAppLink `json:"app,omitempty"`
	SourceGroup *HrefObjectGroupLink `json:"sourceGroup,omitempty"`
	TargetGroup *HrefObjectGroupLink `json:"targetGroup,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupPushMappingLinks GroupPushMappingLinks

// NewGroupPushMappingLinks instantiates a new GroupPushMappingLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupPushMappingLinks() *GroupPushMappingLinks {
	this := GroupPushMappingLinks{}
	return &this
}

// NewGroupPushMappingLinksWithDefaults instantiates a new GroupPushMappingLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupPushMappingLinksWithDefaults() *GroupPushMappingLinks {
	this := GroupPushMappingLinks{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *GroupPushMappingLinks) GetApp() HrefObjectAppLink {
	if o == nil || o.App == nil {
		var ret HrefObjectAppLink
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPushMappingLinks) GetAppOk() (*HrefObjectAppLink, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *GroupPushMappingLinks) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given HrefObjectAppLink and assigns it to the App field.
func (o *GroupPushMappingLinks) SetApp(v HrefObjectAppLink) {
	o.App = &v
}

// GetSourceGroup returns the SourceGroup field value if set, zero value otherwise.
func (o *GroupPushMappingLinks) GetSourceGroup() HrefObjectGroupLink {
	if o == nil || o.SourceGroup == nil {
		var ret HrefObjectGroupLink
		return ret
	}
	return *o.SourceGroup
}

// GetSourceGroupOk returns a tuple with the SourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPushMappingLinks) GetSourceGroupOk() (*HrefObjectGroupLink, bool) {
	if o == nil || o.SourceGroup == nil {
		return nil, false
	}
	return o.SourceGroup, true
}

// HasSourceGroup returns a boolean if a field has been set.
func (o *GroupPushMappingLinks) HasSourceGroup() bool {
	if o != nil && o.SourceGroup != nil {
		return true
	}

	return false
}

// SetSourceGroup gets a reference to the given HrefObjectGroupLink and assigns it to the SourceGroup field.
func (o *GroupPushMappingLinks) SetSourceGroup(v HrefObjectGroupLink) {
	o.SourceGroup = &v
}

// GetTargetGroup returns the TargetGroup field value if set, zero value otherwise.
func (o *GroupPushMappingLinks) GetTargetGroup() HrefObjectGroupLink {
	if o == nil || o.TargetGroup == nil {
		var ret HrefObjectGroupLink
		return ret
	}
	return *o.TargetGroup
}

// GetTargetGroupOk returns a tuple with the TargetGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPushMappingLinks) GetTargetGroupOk() (*HrefObjectGroupLink, bool) {
	if o == nil || o.TargetGroup == nil {
		return nil, false
	}
	return o.TargetGroup, true
}

// HasTargetGroup returns a boolean if a field has been set.
func (o *GroupPushMappingLinks) HasTargetGroup() bool {
	if o != nil && o.TargetGroup != nil {
		return true
	}

	return false
}

// SetTargetGroup gets a reference to the given HrefObjectGroupLink and assigns it to the TargetGroup field.
func (o *GroupPushMappingLinks) SetTargetGroup(v HrefObjectGroupLink) {
	o.TargetGroup = &v
}

func (o GroupPushMappingLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	if o.SourceGroup != nil {
		toSerialize["sourceGroup"] = o.SourceGroup
	}
	if o.TargetGroup != nil {
		toSerialize["targetGroup"] = o.TargetGroup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GroupPushMappingLinks) UnmarshalJSON(bytes []byte) (err error) {
	varGroupPushMappingLinks := _GroupPushMappingLinks{}

	err = json.Unmarshal(bytes, &varGroupPushMappingLinks)
	if err == nil {
		*o = GroupPushMappingLinks(varGroupPushMappingLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "app")
		delete(additionalProperties, "sourceGroup")
		delete(additionalProperties, "targetGroup")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableGroupPushMappingLinks struct {
	value *GroupPushMappingLinks
	isSet bool
}

func (v NullableGroupPushMappingLinks) Get() *GroupPushMappingLinks {
	return v.value
}

func (v *NullableGroupPushMappingLinks) Set(val *GroupPushMappingLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupPushMappingLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupPushMappingLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupPushMappingLinks(val *GroupPushMappingLinks) *NullableGroupPushMappingLinks {
	return &NullableGroupPushMappingLinks{value: val, isSet: true}
}

func (v NullableGroupPushMappingLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupPushMappingLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

