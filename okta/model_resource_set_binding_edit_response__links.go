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

// checks if the ResourceSetBindingEditResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceSetBindingEditResponseLinks{}

// ResourceSetBindingEditResponseLinks struct for ResourceSetBindingEditResponseLinks
type ResourceSetBindingEditResponseLinks struct {
	Self                 *HrefObjectSelfLink        `json:"self,omitempty"`
	ResourceSet          *HrefObjectResourceSetLink `json:"resource-set,omitempty"`
	Bindings             *HrefObjectBindingsLink    `json:"bindings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetBindingEditResponseLinks ResourceSetBindingEditResponseLinks

// NewResourceSetBindingEditResponseLinks instantiates a new ResourceSetBindingEditResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetBindingEditResponseLinks() *ResourceSetBindingEditResponseLinks {
	this := ResourceSetBindingEditResponseLinks{}
	return &this
}

// NewResourceSetBindingEditResponseLinksWithDefaults instantiates a new ResourceSetBindingEditResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetBindingEditResponseLinksWithDefaults() *ResourceSetBindingEditResponseLinks {
	this := ResourceSetBindingEditResponseLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ResourceSetBindingEditResponseLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingEditResponseLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ResourceSetBindingEditResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *ResourceSetBindingEditResponseLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetResourceSet returns the ResourceSet field value if set, zero value otherwise.
func (o *ResourceSetBindingEditResponseLinks) GetResourceSet() HrefObjectResourceSetLink {
	if o == nil || IsNil(o.ResourceSet) {
		var ret HrefObjectResourceSetLink
		return ret
	}
	return *o.ResourceSet
}

// GetResourceSetOk returns a tuple with the ResourceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingEditResponseLinks) GetResourceSetOk() (*HrefObjectResourceSetLink, bool) {
	if o == nil || IsNil(o.ResourceSet) {
		return nil, false
	}
	return o.ResourceSet, true
}

// HasResourceSet returns a boolean if a field has been set.
func (o *ResourceSetBindingEditResponseLinks) HasResourceSet() bool {
	if o != nil && !IsNil(o.ResourceSet) {
		return true
	}

	return false
}

// SetResourceSet gets a reference to the given HrefObjectResourceSetLink and assigns it to the ResourceSet field.
func (o *ResourceSetBindingEditResponseLinks) SetResourceSet(v HrefObjectResourceSetLink) {
	o.ResourceSet = &v
}

// GetBindings returns the Bindings field value if set, zero value otherwise.
func (o *ResourceSetBindingEditResponseLinks) GetBindings() HrefObjectBindingsLink {
	if o == nil || IsNil(o.Bindings) {
		var ret HrefObjectBindingsLink
		return ret
	}
	return *o.Bindings
}

// GetBindingsOk returns a tuple with the Bindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingEditResponseLinks) GetBindingsOk() (*HrefObjectBindingsLink, bool) {
	if o == nil || IsNil(o.Bindings) {
		return nil, false
	}
	return o.Bindings, true
}

// HasBindings returns a boolean if a field has been set.
func (o *ResourceSetBindingEditResponseLinks) HasBindings() bool {
	if o != nil && !IsNil(o.Bindings) {
		return true
	}

	return false
}

// SetBindings gets a reference to the given HrefObjectBindingsLink and assigns it to the Bindings field.
func (o *ResourceSetBindingEditResponseLinks) SetBindings(v HrefObjectBindingsLink) {
	o.Bindings = &v
}

func (o ResourceSetBindingEditResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceSetBindingEditResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.ResourceSet) {
		toSerialize["resource-set"] = o.ResourceSet
	}
	if !IsNil(o.Bindings) {
		toSerialize["bindings"] = o.Bindings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceSetBindingEditResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varResourceSetBindingEditResponseLinks := _ResourceSetBindingEditResponseLinks{}

	err = json.Unmarshal(data, &varResourceSetBindingEditResponseLinks)

	if err != nil {
		return err
	}

	*o = ResourceSetBindingEditResponseLinks(varResourceSetBindingEditResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "resource-set")
		delete(additionalProperties, "bindings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceSetBindingEditResponseLinks struct {
	value *ResourceSetBindingEditResponseLinks
	isSet bool
}

func (v NullableResourceSetBindingEditResponseLinks) Get() *ResourceSetBindingEditResponseLinks {
	return v.value
}

func (v *NullableResourceSetBindingEditResponseLinks) Set(val *ResourceSetBindingEditResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetBindingEditResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetBindingEditResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetBindingEditResponseLinks(val *ResourceSetBindingEditResponseLinks) *NullableResourceSetBindingEditResponseLinks {
	return &NullableResourceSetBindingEditResponseLinks{value: val, isSet: true}
}

func (v NullableResourceSetBindingEditResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetBindingEditResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
