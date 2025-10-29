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

// checks if the InlineHookLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineHookLinks{}

// InlineHookLinks struct for InlineHookLinks
type InlineHookLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	// URL to activate the inline hook
	Activate *HrefObject `json:"activate,omitempty"`
	// URL to deactivate the inline hook
	Deactivate *HrefObject `json:"deactivate,omitempty"`
	// URL to delete the inline hook
	Delete *HrefObject `json:"delete,omitempty"`
	// URL to test the inline hook
	Execute              *HrefObject `json:"execute,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookLinks InlineHookLinks

// NewInlineHookLinks instantiates a new InlineHookLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookLinks() *InlineHookLinks {
	this := InlineHookLinks{}
	return &this
}

// NewInlineHookLinksWithDefaults instantiates a new InlineHookLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookLinksWithDefaults() *InlineHookLinks {
	this := InlineHookLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *InlineHookLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *InlineHookLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *InlineHookLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *InlineHookLinks) GetActivate() HrefObject {
	if o == nil || IsNil(o.Activate) {
		var ret HrefObject
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookLinks) GetActivateOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Activate) {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *InlineHookLinks) HasActivate() bool {
	if o != nil && !IsNil(o.Activate) {
		return true
	}

	return false
}

// SetActivate gets a reference to the given HrefObject and assigns it to the Activate field.
func (o *InlineHookLinks) SetActivate(v HrefObject) {
	o.Activate = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *InlineHookLinks) GetDeactivate() HrefObject {
	if o == nil || IsNil(o.Deactivate) {
		var ret HrefObject
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookLinks) GetDeactivateOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Deactivate) {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *InlineHookLinks) HasDeactivate() bool {
	if o != nil && !IsNil(o.Deactivate) {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObject and assigns it to the Deactivate field.
func (o *InlineHookLinks) SetDeactivate(v HrefObject) {
	o.Deactivate = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *InlineHookLinks) GetDelete() HrefObject {
	if o == nil || IsNil(o.Delete) {
		var ret HrefObject
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookLinks) GetDeleteOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *InlineHookLinks) HasDelete() bool {
	if o != nil && !IsNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given HrefObject and assigns it to the Delete field.
func (o *InlineHookLinks) SetDelete(v HrefObject) {
	o.Delete = &v
}

// GetExecute returns the Execute field value if set, zero value otherwise.
func (o *InlineHookLinks) GetExecute() HrefObject {
	if o == nil || IsNil(o.Execute) {
		var ret HrefObject
		return ret
	}
	return *o.Execute
}

// GetExecuteOk returns a tuple with the Execute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookLinks) GetExecuteOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Execute) {
		return nil, false
	}
	return o.Execute, true
}

// HasExecute returns a boolean if a field has been set.
func (o *InlineHookLinks) HasExecute() bool {
	if o != nil && !IsNil(o.Execute) {
		return true
	}

	return false
}

// SetExecute gets a reference to the given HrefObject and assigns it to the Execute field.
func (o *InlineHookLinks) SetExecute(v HrefObject) {
	o.Execute = &v
}

func (o InlineHookLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineHookLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Activate) {
		toSerialize["activate"] = o.Activate
	}
	if !IsNil(o.Deactivate) {
		toSerialize["deactivate"] = o.Deactivate
	}
	if !IsNil(o.Delete) {
		toSerialize["delete"] = o.Delete
	}
	if !IsNil(o.Execute) {
		toSerialize["execute"] = o.Execute
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineHookLinks) UnmarshalJSON(data []byte) (err error) {
	varInlineHookLinks := _InlineHookLinks{}

	err = json.Unmarshal(data, &varInlineHookLinks)

	if err != nil {
		return err
	}

	*o = InlineHookLinks(varInlineHookLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "activate")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "delete")
		delete(additionalProperties, "execute")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineHookLinks struct {
	value *InlineHookLinks
	isSet bool
}

func (v NullableInlineHookLinks) Get() *InlineHookLinks {
	return v.value
}

func (v *NullableInlineHookLinks) Set(val *InlineHookLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookLinks(val *InlineHookLinks) *NullableInlineHookLinks {
	return &NullableInlineHookLinks{value: val, isSet: true}
}

func (v NullableInlineHookLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
