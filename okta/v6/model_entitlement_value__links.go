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

// EntitlementValueLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification
type EntitlementValueLinks struct {
	Group *HrefObjectGroupLink `json:"group,omitempty"`
	App *HrefObjectAppLink `json:"app,omitempty"`
	ResourceSet *HrefObjectResourceSetLink `json:"resource-set,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementValueLinks EntitlementValueLinks

// NewEntitlementValueLinks instantiates a new EntitlementValueLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementValueLinks() *EntitlementValueLinks {
	this := EntitlementValueLinks{}
	return &this
}

// NewEntitlementValueLinksWithDefaults instantiates a new EntitlementValueLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementValueLinksWithDefaults() *EntitlementValueLinks {
	this := EntitlementValueLinks{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *EntitlementValueLinks) GetGroup() HrefObjectGroupLink {
	if o == nil || o.Group == nil {
		var ret HrefObjectGroupLink
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueLinks) GetGroupOk() (*HrefObjectGroupLink, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *EntitlementValueLinks) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given HrefObjectGroupLink and assigns it to the Group field.
func (o *EntitlementValueLinks) SetGroup(v HrefObjectGroupLink) {
	o.Group = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *EntitlementValueLinks) GetApp() HrefObjectAppLink {
	if o == nil || o.App == nil {
		var ret HrefObjectAppLink
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueLinks) GetAppOk() (*HrefObjectAppLink, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *EntitlementValueLinks) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given HrefObjectAppLink and assigns it to the App field.
func (o *EntitlementValueLinks) SetApp(v HrefObjectAppLink) {
	o.App = &v
}

// GetResourceSet returns the ResourceSet field value if set, zero value otherwise.
func (o *EntitlementValueLinks) GetResourceSet() HrefObjectResourceSetLink {
	if o == nil || o.ResourceSet == nil {
		var ret HrefObjectResourceSetLink
		return ret
	}
	return *o.ResourceSet
}

// GetResourceSetOk returns a tuple with the ResourceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueLinks) GetResourceSetOk() (*HrefObjectResourceSetLink, bool) {
	if o == nil || o.ResourceSet == nil {
		return nil, false
	}
	return o.ResourceSet, true
}

// HasResourceSet returns a boolean if a field has been set.
func (o *EntitlementValueLinks) HasResourceSet() bool {
	if o != nil && o.ResourceSet != nil {
		return true
	}

	return false
}

// SetResourceSet gets a reference to the given HrefObjectResourceSetLink and assigns it to the ResourceSet field.
func (o *EntitlementValueLinks) SetResourceSet(v HrefObjectResourceSetLink) {
	o.ResourceSet = &v
}

func (o EntitlementValueLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	if o.ResourceSet != nil {
		toSerialize["resource-set"] = o.ResourceSet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementValueLinks) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementValueLinks := _EntitlementValueLinks{}

	err = json.Unmarshal(bytes, &varEntitlementValueLinks)
	if err == nil {
		*o = EntitlementValueLinks(varEntitlementValueLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "group")
		delete(additionalProperties, "app")
		delete(additionalProperties, "resource-set")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementValueLinks struct {
	value *EntitlementValueLinks
	isSet bool
}

func (v NullableEntitlementValueLinks) Get() *EntitlementValueLinks {
	return v.value
}

func (v *NullableEntitlementValueLinks) Set(val *EntitlementValueLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementValueLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementValueLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementValueLinks(val *EntitlementValueLinks) *NullableEntitlementValueLinks {
	return &NullableEntitlementValueLinks{value: val, isSet: true}
}

func (v NullableEntitlementValueLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementValueLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

