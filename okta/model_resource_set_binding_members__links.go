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

// ResourceSetBindingMembersLinks struct for ResourceSetBindingMembersLinks
type ResourceSetBindingMembersLinks struct {
	Next *HrefObject `json:"next,omitempty"`
	Binding *HrefObject `json:"binding,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetBindingMembersLinks ResourceSetBindingMembersLinks

// NewResourceSetBindingMembersLinks instantiates a new ResourceSetBindingMembersLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetBindingMembersLinks() *ResourceSetBindingMembersLinks {
	this := ResourceSetBindingMembersLinks{}
	return &this
}

// NewResourceSetBindingMembersLinksWithDefaults instantiates a new ResourceSetBindingMembersLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetBindingMembersLinksWithDefaults() *ResourceSetBindingMembersLinks {
	this := ResourceSetBindingMembersLinks{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ResourceSetBindingMembersLinks) GetNext() HrefObject {
	if o == nil || o.Next == nil {
		var ret HrefObject
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingMembersLinks) GetNextOk() (*HrefObject, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ResourceSetBindingMembersLinks) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given HrefObject and assigns it to the Next field.
func (o *ResourceSetBindingMembersLinks) SetNext(v HrefObject) {
	o.Next = &v
}

// GetBinding returns the Binding field value if set, zero value otherwise.
func (o *ResourceSetBindingMembersLinks) GetBinding() HrefObject {
	if o == nil || o.Binding == nil {
		var ret HrefObject
		return ret
	}
	return *o.Binding
}

// GetBindingOk returns a tuple with the Binding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingMembersLinks) GetBindingOk() (*HrefObject, bool) {
	if o == nil || o.Binding == nil {
		return nil, false
	}
	return o.Binding, true
}

// HasBinding returns a boolean if a field has been set.
func (o *ResourceSetBindingMembersLinks) HasBinding() bool {
	if o != nil && o.Binding != nil {
		return true
	}

	return false
}

// SetBinding gets a reference to the given HrefObject and assigns it to the Binding field.
func (o *ResourceSetBindingMembersLinks) SetBinding(v HrefObject) {
	o.Binding = &v
}

func (o ResourceSetBindingMembersLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Binding != nil {
		toSerialize["binding"] = o.Binding
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSetBindingMembersLinks) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSetBindingMembersLinks := _ResourceSetBindingMembersLinks{}

	err = json.Unmarshal(bytes, &varResourceSetBindingMembersLinks)
	if err == nil {
		*o = ResourceSetBindingMembersLinks(varResourceSetBindingMembersLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "next")
		delete(additionalProperties, "binding")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSetBindingMembersLinks struct {
	value *ResourceSetBindingMembersLinks
	isSet bool
}

func (v NullableResourceSetBindingMembersLinks) Get() *ResourceSetBindingMembersLinks {
	return v.value
}

func (v *NullableResourceSetBindingMembersLinks) Set(val *ResourceSetBindingMembersLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetBindingMembersLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetBindingMembersLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetBindingMembersLinks(val *ResourceSetBindingMembersLinks) *NullableResourceSetBindingMembersLinks {
	return &NullableResourceSetBindingMembersLinks{value: val, isSet: true}
}

func (v NullableResourceSetBindingMembersLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetBindingMembersLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

