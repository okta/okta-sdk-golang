/*
Okta Admin Management API

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

// checks if the CimdClientEntityBindingLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CimdClientEntityBindingLinks{}

// CimdClientEntityBindingLinks struct for CimdClientEntityBindingLinks
type CimdClientEntityBindingLinks struct {
	Self                 *CimdClientEntityLinksSelf `json:"self,omitempty"`
	CimdClientEntity     *CimdClientEntityLinksSelf `json:"cimdClientEntity,omitempty"`
	BoundEntity          *CimdClientEntityLinksSelf `json:"boundEntity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CimdClientEntityBindingLinks CimdClientEntityBindingLinks

// NewCimdClientEntityBindingLinks instantiates a new CimdClientEntityBindingLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCimdClientEntityBindingLinks() *CimdClientEntityBindingLinks {
	this := CimdClientEntityBindingLinks{}
	return &this
}

// NewCimdClientEntityBindingLinksWithDefaults instantiates a new CimdClientEntityBindingLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCimdClientEntityBindingLinksWithDefaults() *CimdClientEntityBindingLinks {
	this := CimdClientEntityBindingLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *CimdClientEntityBindingLinks) GetSelf() CimdClientEntityLinksSelf {
	if o == nil || IsNil(o.Self) {
		var ret CimdClientEntityLinksSelf
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntityBindingLinks) GetSelfOk() (*CimdClientEntityLinksSelf, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *CimdClientEntityBindingLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given CimdClientEntityLinksSelf and assigns it to the Self field.
func (o *CimdClientEntityBindingLinks) SetSelf(v CimdClientEntityLinksSelf) {
	o.Self = &v
}

// GetCimdClientEntity returns the CimdClientEntity field value if set, zero value otherwise.
func (o *CimdClientEntityBindingLinks) GetCimdClientEntity() CimdClientEntityLinksSelf {
	if o == nil || IsNil(o.CimdClientEntity) {
		var ret CimdClientEntityLinksSelf
		return ret
	}
	return *o.CimdClientEntity
}

// GetCimdClientEntityOk returns a tuple with the CimdClientEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntityBindingLinks) GetCimdClientEntityOk() (*CimdClientEntityLinksSelf, bool) {
	if o == nil || IsNil(o.CimdClientEntity) {
		return nil, false
	}
	return o.CimdClientEntity, true
}

// HasCimdClientEntity returns a boolean if a field has been set.
func (o *CimdClientEntityBindingLinks) HasCimdClientEntity() bool {
	if o != nil && !IsNil(o.CimdClientEntity) {
		return true
	}

	return false
}

// SetCimdClientEntity gets a reference to the given CimdClientEntityLinksSelf and assigns it to the CimdClientEntity field.
func (o *CimdClientEntityBindingLinks) SetCimdClientEntity(v CimdClientEntityLinksSelf) {
	o.CimdClientEntity = &v
}

// GetBoundEntity returns the BoundEntity field value if set, zero value otherwise.
func (o *CimdClientEntityBindingLinks) GetBoundEntity() CimdClientEntityLinksSelf {
	if o == nil || IsNil(o.BoundEntity) {
		var ret CimdClientEntityLinksSelf
		return ret
	}
	return *o.BoundEntity
}

// GetBoundEntityOk returns a tuple with the BoundEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntityBindingLinks) GetBoundEntityOk() (*CimdClientEntityLinksSelf, bool) {
	if o == nil || IsNil(o.BoundEntity) {
		return nil, false
	}
	return o.BoundEntity, true
}

// HasBoundEntity returns a boolean if a field has been set.
func (o *CimdClientEntityBindingLinks) HasBoundEntity() bool {
	if o != nil && !IsNil(o.BoundEntity) {
		return true
	}

	return false
}

// SetBoundEntity gets a reference to the given CimdClientEntityLinksSelf and assigns it to the BoundEntity field.
func (o *CimdClientEntityBindingLinks) SetBoundEntity(v CimdClientEntityLinksSelf) {
	o.BoundEntity = &v
}

func (o CimdClientEntityBindingLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CimdClientEntityBindingLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.CimdClientEntity) {
		toSerialize["cimdClientEntity"] = o.CimdClientEntity
	}
	if !IsNil(o.BoundEntity) {
		toSerialize["boundEntity"] = o.BoundEntity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CimdClientEntityBindingLinks) UnmarshalJSON(data []byte) (err error) {
	varCimdClientEntityBindingLinks := _CimdClientEntityBindingLinks{}

	err = json.Unmarshal(data, &varCimdClientEntityBindingLinks)

	if err != nil {
		return err
	}

	*o = CimdClientEntityBindingLinks(varCimdClientEntityBindingLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "cimdClientEntity")
		delete(additionalProperties, "boundEntity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCimdClientEntityBindingLinks struct {
	value *CimdClientEntityBindingLinks
	isSet bool
}

func (v NullableCimdClientEntityBindingLinks) Get() *CimdClientEntityBindingLinks {
	return v.value
}

func (v *NullableCimdClientEntityBindingLinks) Set(val *CimdClientEntityBindingLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableCimdClientEntityBindingLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableCimdClientEntityBindingLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCimdClientEntityBindingLinks(val *CimdClientEntityBindingLinks) *NullableCimdClientEntityBindingLinks {
	return &NullableCimdClientEntityBindingLinks{value: val, isSet: true}
}

func (v NullableCimdClientEntityBindingLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCimdClientEntityBindingLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
