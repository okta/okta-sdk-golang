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

// checks if the AIAgentOperationListResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIAgentOperationListResponseLinks{}

// AIAgentOperationListResponseLinks struct for AIAgentOperationListResponseLinks
type AIAgentOperationListResponseLinks struct {
	Self                 *HrefObjectSelfLink `json:"self,omitempty"`
	Next                 *HrefObjectNextLink `json:"next,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AIAgentOperationListResponseLinks AIAgentOperationListResponseLinks

// NewAIAgentOperationListResponseLinks instantiates a new AIAgentOperationListResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIAgentOperationListResponseLinks() *AIAgentOperationListResponseLinks {
	this := AIAgentOperationListResponseLinks{}
	return &this
}

// NewAIAgentOperationListResponseLinksWithDefaults instantiates a new AIAgentOperationListResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIAgentOperationListResponseLinksWithDefaults() *AIAgentOperationListResponseLinks {
	this := AIAgentOperationListResponseLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *AIAgentOperationListResponseLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIAgentOperationListResponseLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *AIAgentOperationListResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *AIAgentOperationListResponseLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *AIAgentOperationListResponseLinks) GetNext() HrefObjectNextLink {
	if o == nil || IsNil(o.Next) {
		var ret HrefObjectNextLink
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIAgentOperationListResponseLinks) GetNextOk() (*HrefObjectNextLink, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *AIAgentOperationListResponseLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given HrefObjectNextLink and assigns it to the Next field.
func (o *AIAgentOperationListResponseLinks) SetNext(v HrefObjectNextLink) {
	o.Next = &v
}

func (o AIAgentOperationListResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AIAgentOperationListResponseLinks) ToMap() (map[string]interface{}, error) {
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

func (o *AIAgentOperationListResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varAIAgentOperationListResponseLinks := _AIAgentOperationListResponseLinks{}

	err = json.Unmarshal(data, &varAIAgentOperationListResponseLinks)

	if err != nil {
		return err
	}

	*o = AIAgentOperationListResponseLinks(varAIAgentOperationListResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAIAgentOperationListResponseLinks struct {
	value *AIAgentOperationListResponseLinks
	isSet bool
}

func (v NullableAIAgentOperationListResponseLinks) Get() *AIAgentOperationListResponseLinks {
	return v.value
}

func (v *NullableAIAgentOperationListResponseLinks) Set(val *AIAgentOperationListResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableAIAgentOperationListResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableAIAgentOperationListResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIAgentOperationListResponseLinks(val *AIAgentOperationListResponseLinks) *NullableAIAgentOperationListResponseLinks {
	return &NullableAIAgentOperationListResponseLinks{value: val, isSet: true}
}

func (v NullableAIAgentOperationListResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIAgentOperationListResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
