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

// checks if the InlineHookLinksCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineHookLinksCreate{}

// InlineHookLinksCreate struct for InlineHookLinksCreate
type InlineHookLinksCreate struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	// URL to deactivate the inline hook
	Deactivate *HrefObject `json:"deactivate,omitempty"`
	// URL to test the inline hook
	Execute              *HrefObject `json:"execute,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookLinksCreate InlineHookLinksCreate

// NewInlineHookLinksCreate instantiates a new InlineHookLinksCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookLinksCreate() *InlineHookLinksCreate {
	this := InlineHookLinksCreate{}
	return &this
}

// NewInlineHookLinksCreateWithDefaults instantiates a new InlineHookLinksCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookLinksCreateWithDefaults() *InlineHookLinksCreate {
	this := InlineHookLinksCreate{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *InlineHookLinksCreate) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookLinksCreate) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *InlineHookLinksCreate) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *InlineHookLinksCreate) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *InlineHookLinksCreate) GetDeactivate() HrefObject {
	if o == nil || IsNil(o.Deactivate) {
		var ret HrefObject
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookLinksCreate) GetDeactivateOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Deactivate) {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *InlineHookLinksCreate) HasDeactivate() bool {
	if o != nil && !IsNil(o.Deactivate) {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObject and assigns it to the Deactivate field.
func (o *InlineHookLinksCreate) SetDeactivate(v HrefObject) {
	o.Deactivate = &v
}

// GetExecute returns the Execute field value if set, zero value otherwise.
func (o *InlineHookLinksCreate) GetExecute() HrefObject {
	if o == nil || IsNil(o.Execute) {
		var ret HrefObject
		return ret
	}
	return *o.Execute
}

// GetExecuteOk returns a tuple with the Execute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookLinksCreate) GetExecuteOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Execute) {
		return nil, false
	}
	return o.Execute, true
}

// HasExecute returns a boolean if a field has been set.
func (o *InlineHookLinksCreate) HasExecute() bool {
	if o != nil && !IsNil(o.Execute) {
		return true
	}

	return false
}

// SetExecute gets a reference to the given HrefObject and assigns it to the Execute field.
func (o *InlineHookLinksCreate) SetExecute(v HrefObject) {
	o.Execute = &v
}

func (o InlineHookLinksCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineHookLinksCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Deactivate) {
		toSerialize["deactivate"] = o.Deactivate
	}
	if !IsNil(o.Execute) {
		toSerialize["execute"] = o.Execute
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineHookLinksCreate) UnmarshalJSON(data []byte) (err error) {
	varInlineHookLinksCreate := _InlineHookLinksCreate{}

	err = json.Unmarshal(data, &varInlineHookLinksCreate)

	if err != nil {
		return err
	}

	*o = InlineHookLinksCreate(varInlineHookLinksCreate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "execute")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineHookLinksCreate struct {
	value *InlineHookLinksCreate
	isSet bool
}

func (v NullableInlineHookLinksCreate) Get() *InlineHookLinksCreate {
	return v.value
}

func (v *NullableInlineHookLinksCreate) Set(val *InlineHookLinksCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookLinksCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookLinksCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookLinksCreate(val *InlineHookLinksCreate) *NullableInlineHookLinksCreate {
	return &NullableInlineHookLinksCreate{value: val, isSet: true}
}

func (v NullableInlineHookLinksCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookLinksCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
