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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the ResourceSetBindingMembersLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceSetBindingMembersLinks{}

// ResourceSetBindingMembersLinks struct for ResourceSetBindingMembersLinks
type ResourceSetBindingMembersLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	// Link to the next list of binding members for the specified role and resource set
	Next                 *HrefObject            `json:"next,omitempty"`
	Binding              *HrefObjectBindingLink `json:"binding,omitempty"`
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

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ResourceSetBindingMembersLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingMembersLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ResourceSetBindingMembersLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *ResourceSetBindingMembersLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ResourceSetBindingMembersLinks) GetNext() HrefObject {
	if o == nil || IsNil(o.Next) {
		var ret HrefObject
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingMembersLinks) GetNextOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ResourceSetBindingMembersLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given HrefObject and assigns it to the Next field.
func (o *ResourceSetBindingMembersLinks) SetNext(v HrefObject) {
	o.Next = &v
}

// GetBinding returns the Binding field value if set, zero value otherwise.
func (o *ResourceSetBindingMembersLinks) GetBinding() HrefObjectBindingLink {
	if o == nil || IsNil(o.Binding) {
		var ret HrefObjectBindingLink
		return ret
	}
	return *o.Binding
}

// GetBindingOk returns a tuple with the Binding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingMembersLinks) GetBindingOk() (*HrefObjectBindingLink, bool) {
	if o == nil || IsNil(o.Binding) {
		return nil, false
	}
	return o.Binding, true
}

// HasBinding returns a boolean if a field has been set.
func (o *ResourceSetBindingMembersLinks) HasBinding() bool {
	if o != nil && !IsNil(o.Binding) {
		return true
	}

	return false
}

// SetBinding gets a reference to the given HrefObjectBindingLink and assigns it to the Binding field.
func (o *ResourceSetBindingMembersLinks) SetBinding(v HrefObjectBindingLink) {
	o.Binding = &v
}

func (o ResourceSetBindingMembersLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceSetBindingMembersLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Binding) {
		toSerialize["binding"] = o.Binding
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceSetBindingMembersLinks) UnmarshalJSON(data []byte) (err error) {
	varResourceSetBindingMembersLinks := _ResourceSetBindingMembersLinks{}

	err = json.Unmarshal(data, &varResourceSetBindingMembersLinks)

	if err != nil {
		return err
	}

	*o = ResourceSetBindingMembersLinks(varResourceSetBindingMembersLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "next")
		delete(additionalProperties, "binding")
		o.AdditionalProperties = additionalProperties
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
