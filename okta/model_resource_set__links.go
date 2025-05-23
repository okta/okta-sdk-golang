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

// ResourceSetLinks struct for ResourceSetLinks
type ResourceSetLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Resources *HrefObject `json:"resources,omitempty"`
	Bindings *HrefObject `json:"bindings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetLinks ResourceSetLinks

// NewResourceSetLinks instantiates a new ResourceSetLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetLinks() *ResourceSetLinks {
	this := ResourceSetLinks{}
	return &this
}

// NewResourceSetLinksWithDefaults instantiates a new ResourceSetLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetLinksWithDefaults() *ResourceSetLinks {
	this := ResourceSetLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ResourceSetLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ResourceSetLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *ResourceSetLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ResourceSetLinks) GetResources() HrefObject {
	if o == nil || o.Resources == nil {
		var ret HrefObject
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetLinks) GetResourcesOk() (*HrefObject, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ResourceSetLinks) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given HrefObject and assigns it to the Resources field.
func (o *ResourceSetLinks) SetResources(v HrefObject) {
	o.Resources = &v
}

// GetBindings returns the Bindings field value if set, zero value otherwise.
func (o *ResourceSetLinks) GetBindings() HrefObject {
	if o == nil || o.Bindings == nil {
		var ret HrefObject
		return ret
	}
	return *o.Bindings
}

// GetBindingsOk returns a tuple with the Bindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetLinks) GetBindingsOk() (*HrefObject, bool) {
	if o == nil || o.Bindings == nil {
		return nil, false
	}
	return o.Bindings, true
}

// HasBindings returns a boolean if a field has been set.
func (o *ResourceSetLinks) HasBindings() bool {
	if o != nil && o.Bindings != nil {
		return true
	}

	return false
}

// SetBindings gets a reference to the given HrefObject and assigns it to the Bindings field.
func (o *ResourceSetLinks) SetBindings(v HrefObject) {
	o.Bindings = &v
}

func (o ResourceSetLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.Bindings != nil {
		toSerialize["bindings"] = o.Bindings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSetLinks) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSetLinks := _ResourceSetLinks{}

	err = json.Unmarshal(bytes, &varResourceSetLinks)
	if err == nil {
		*o = ResourceSetLinks(varResourceSetLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "resources")
		delete(additionalProperties, "bindings")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSetLinks struct {
	value *ResourceSetLinks
	isSet bool
}

func (v NullableResourceSetLinks) Get() *ResourceSetLinks {
	return v.value
}

func (v *NullableResourceSetLinks) Set(val *ResourceSetLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetLinks(val *ResourceSetLinks) *NullableResourceSetLinks {
	return &NullableResourceSetLinks{value: val, isSet: true}
}

func (v NullableResourceSetLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

