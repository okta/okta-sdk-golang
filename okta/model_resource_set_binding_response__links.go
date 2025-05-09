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

// ResourceSetBindingResponseLinks struct for ResourceSetBindingResponseLinks
type ResourceSetBindingResponseLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Bindings *HrefObject `json:"bindings,omitempty"`
	ResourceSet *HrefObject `json:"resource-set,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetBindingResponseLinks ResourceSetBindingResponseLinks

// NewResourceSetBindingResponseLinks instantiates a new ResourceSetBindingResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetBindingResponseLinks() *ResourceSetBindingResponseLinks {
	this := ResourceSetBindingResponseLinks{}
	return &this
}

// NewResourceSetBindingResponseLinksWithDefaults instantiates a new ResourceSetBindingResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetBindingResponseLinksWithDefaults() *ResourceSetBindingResponseLinks {
	this := ResourceSetBindingResponseLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ResourceSetBindingResponseLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingResponseLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ResourceSetBindingResponseLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *ResourceSetBindingResponseLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetBindings returns the Bindings field value if set, zero value otherwise.
func (o *ResourceSetBindingResponseLinks) GetBindings() HrefObject {
	if o == nil || o.Bindings == nil {
		var ret HrefObject
		return ret
	}
	return *o.Bindings
}

// GetBindingsOk returns a tuple with the Bindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingResponseLinks) GetBindingsOk() (*HrefObject, bool) {
	if o == nil || o.Bindings == nil {
		return nil, false
	}
	return o.Bindings, true
}

// HasBindings returns a boolean if a field has been set.
func (o *ResourceSetBindingResponseLinks) HasBindings() bool {
	if o != nil && o.Bindings != nil {
		return true
	}

	return false
}

// SetBindings gets a reference to the given HrefObject and assigns it to the Bindings field.
func (o *ResourceSetBindingResponseLinks) SetBindings(v HrefObject) {
	o.Bindings = &v
}

// GetResourceSet returns the ResourceSet field value if set, zero value otherwise.
func (o *ResourceSetBindingResponseLinks) GetResourceSet() HrefObject {
	if o == nil || o.ResourceSet == nil {
		var ret HrefObject
		return ret
	}
	return *o.ResourceSet
}

// GetResourceSetOk returns a tuple with the ResourceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingResponseLinks) GetResourceSetOk() (*HrefObject, bool) {
	if o == nil || o.ResourceSet == nil {
		return nil, false
	}
	return o.ResourceSet, true
}

// HasResourceSet returns a boolean if a field has been set.
func (o *ResourceSetBindingResponseLinks) HasResourceSet() bool {
	if o != nil && o.ResourceSet != nil {
		return true
	}

	return false
}

// SetResourceSet gets a reference to the given HrefObject and assigns it to the ResourceSet field.
func (o *ResourceSetBindingResponseLinks) SetResourceSet(v HrefObject) {
	o.ResourceSet = &v
}

func (o ResourceSetBindingResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Bindings != nil {
		toSerialize["bindings"] = o.Bindings
	}
	if o.ResourceSet != nil {
		toSerialize["resource-set"] = o.ResourceSet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSetBindingResponseLinks) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSetBindingResponseLinks := _ResourceSetBindingResponseLinks{}

	err = json.Unmarshal(bytes, &varResourceSetBindingResponseLinks)
	if err == nil {
		*o = ResourceSetBindingResponseLinks(varResourceSetBindingResponseLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "bindings")
		delete(additionalProperties, "resource-set")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSetBindingResponseLinks struct {
	value *ResourceSetBindingResponseLinks
	isSet bool
}

func (v NullableResourceSetBindingResponseLinks) Get() *ResourceSetBindingResponseLinks {
	return v.value
}

func (v *NullableResourceSetBindingResponseLinks) Set(val *ResourceSetBindingResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetBindingResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetBindingResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetBindingResponseLinks(val *ResourceSetBindingResponseLinks) *NullableResourceSetBindingResponseLinks {
	return &NullableResourceSetBindingResponseLinks{value: val, isSet: true}
}

func (v NullableResourceSetBindingResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetBindingResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

