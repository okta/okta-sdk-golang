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

// ResourceSetBindingsLinks struct for ResourceSetBindingsLinks
type ResourceSetBindingsLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	// Link to the next list of bindings for the specified resource set
	Next *HrefObject `json:"next,omitempty"`
	ResourceSet *HrefObjectResourceSetLink `json:"resource-set,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetBindingsLinks ResourceSetBindingsLinks

// NewResourceSetBindingsLinks instantiates a new ResourceSetBindingsLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetBindingsLinks() *ResourceSetBindingsLinks {
	this := ResourceSetBindingsLinks{}
	return &this
}

// NewResourceSetBindingsLinksWithDefaults instantiates a new ResourceSetBindingsLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetBindingsLinksWithDefaults() *ResourceSetBindingsLinks {
	this := ResourceSetBindingsLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ResourceSetBindingsLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingsLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ResourceSetBindingsLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *ResourceSetBindingsLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ResourceSetBindingsLinks) GetNext() HrefObject {
	if o == nil || o.Next == nil {
		var ret HrefObject
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingsLinks) GetNextOk() (*HrefObject, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ResourceSetBindingsLinks) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given HrefObject and assigns it to the Next field.
func (o *ResourceSetBindingsLinks) SetNext(v HrefObject) {
	o.Next = &v
}

// GetResourceSet returns the ResourceSet field value if set, zero value otherwise.
func (o *ResourceSetBindingsLinks) GetResourceSet() HrefObjectResourceSetLink {
	if o == nil || o.ResourceSet == nil {
		var ret HrefObjectResourceSetLink
		return ret
	}
	return *o.ResourceSet
}

// GetResourceSetOk returns a tuple with the ResourceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingsLinks) GetResourceSetOk() (*HrefObjectResourceSetLink, bool) {
	if o == nil || o.ResourceSet == nil {
		return nil, false
	}
	return o.ResourceSet, true
}

// HasResourceSet returns a boolean if a field has been set.
func (o *ResourceSetBindingsLinks) HasResourceSet() bool {
	if o != nil && o.ResourceSet != nil {
		return true
	}

	return false
}

// SetResourceSet gets a reference to the given HrefObjectResourceSetLink and assigns it to the ResourceSet field.
func (o *ResourceSetBindingsLinks) SetResourceSet(v HrefObjectResourceSetLink) {
	o.ResourceSet = &v
}

func (o ResourceSetBindingsLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.ResourceSet != nil {
		toSerialize["resource-set"] = o.ResourceSet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSetBindingsLinks) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSetBindingsLinks := _ResourceSetBindingsLinks{}

	err = json.Unmarshal(bytes, &varResourceSetBindingsLinks)
	if err == nil {
		*o = ResourceSetBindingsLinks(varResourceSetBindingsLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "next")
		delete(additionalProperties, "resource-set")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSetBindingsLinks struct {
	value *ResourceSetBindingsLinks
	isSet bool
}

func (v NullableResourceSetBindingsLinks) Get() *ResourceSetBindingsLinks {
	return v.value
}

func (v *NullableResourceSetBindingsLinks) Set(val *ResourceSetBindingsLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetBindingsLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetBindingsLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetBindingsLinks(val *ResourceSetBindingsLinks) *NullableResourceSetBindingsLinks {
	return &NullableResourceSetBindingsLinks{value: val, isSet: true}
}

func (v NullableResourceSetBindingsLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetBindingsLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

