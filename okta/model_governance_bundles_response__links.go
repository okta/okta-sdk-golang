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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the GovernanceBundlesResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GovernanceBundlesResponseLinks{}

// GovernanceBundlesResponseLinks struct for GovernanceBundlesResponseLinks
type GovernanceBundlesResponseLinks struct {
	Self                 *HrefObjectSelfLink `json:"self,omitempty"`
	Next                 *HrefObjectNextLink `json:"next,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GovernanceBundlesResponseLinks GovernanceBundlesResponseLinks

// NewGovernanceBundlesResponseLinks instantiates a new GovernanceBundlesResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGovernanceBundlesResponseLinks() *GovernanceBundlesResponseLinks {
	this := GovernanceBundlesResponseLinks{}
	return &this
}

// NewGovernanceBundlesResponseLinksWithDefaults instantiates a new GovernanceBundlesResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGovernanceBundlesResponseLinksWithDefaults() *GovernanceBundlesResponseLinks {
	this := GovernanceBundlesResponseLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *GovernanceBundlesResponseLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceBundlesResponseLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *GovernanceBundlesResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *GovernanceBundlesResponseLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *GovernanceBundlesResponseLinks) GetNext() HrefObjectNextLink {
	if o == nil || IsNil(o.Next) {
		var ret HrefObjectNextLink
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceBundlesResponseLinks) GetNextOk() (*HrefObjectNextLink, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *GovernanceBundlesResponseLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given HrefObjectNextLink and assigns it to the Next field.
func (o *GovernanceBundlesResponseLinks) SetNext(v HrefObjectNextLink) {
	o.Next = &v
}

func (o GovernanceBundlesResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GovernanceBundlesResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GovernanceBundlesResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varGovernanceBundlesResponseLinks := _GovernanceBundlesResponseLinks{}

	err = json.Unmarshal(data, &varGovernanceBundlesResponseLinks)

	if err != nil {
		return err
	}

	*o = GovernanceBundlesResponseLinks(varGovernanceBundlesResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGovernanceBundlesResponseLinks struct {
	value *GovernanceBundlesResponseLinks
	isSet bool
}

func (v NullableGovernanceBundlesResponseLinks) Get() *GovernanceBundlesResponseLinks {
	return v.value
}

func (v *NullableGovernanceBundlesResponseLinks) Set(val *GovernanceBundlesResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGovernanceBundlesResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGovernanceBundlesResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGovernanceBundlesResponseLinks(val *GovernanceBundlesResponseLinks) *NullableGovernanceBundlesResponseLinks {
	return &NullableGovernanceBundlesResponseLinks{value: val, isSet: true}
}

func (v NullableGovernanceBundlesResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGovernanceBundlesResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
