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

// checks if the CimdClientEntityLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CimdClientEntityLinks{}

// CimdClientEntityLinks struct for CimdClientEntityLinks
type CimdClientEntityLinks struct {
	Self                 *CimdClientEntityLinksSelf                 `json:"self,omitempty"`
	MetadataRequirements *CimdClientEntityLinksMetadataRequirements `json:"metadataRequirements,omitempty"`
	// Links to the entities that are bound to the CIMD client entity
	Bindings             []CimdClientEntityLinksSelf `json:"bindings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CimdClientEntityLinks CimdClientEntityLinks

// NewCimdClientEntityLinks instantiates a new CimdClientEntityLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCimdClientEntityLinks() *CimdClientEntityLinks {
	this := CimdClientEntityLinks{}
	return &this
}

// NewCimdClientEntityLinksWithDefaults instantiates a new CimdClientEntityLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCimdClientEntityLinksWithDefaults() *CimdClientEntityLinks {
	this := CimdClientEntityLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *CimdClientEntityLinks) GetSelf() CimdClientEntityLinksSelf {
	if o == nil || IsNil(o.Self) {
		var ret CimdClientEntityLinksSelf
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntityLinks) GetSelfOk() (*CimdClientEntityLinksSelf, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *CimdClientEntityLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given CimdClientEntityLinksSelf and assigns it to the Self field.
func (o *CimdClientEntityLinks) SetSelf(v CimdClientEntityLinksSelf) {
	o.Self = &v
}

// GetMetadataRequirements returns the MetadataRequirements field value if set, zero value otherwise.
func (o *CimdClientEntityLinks) GetMetadataRequirements() CimdClientEntityLinksMetadataRequirements {
	if o == nil || IsNil(o.MetadataRequirements) {
		var ret CimdClientEntityLinksMetadataRequirements
		return ret
	}
	return *o.MetadataRequirements
}

// GetMetadataRequirementsOk returns a tuple with the MetadataRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntityLinks) GetMetadataRequirementsOk() (*CimdClientEntityLinksMetadataRequirements, bool) {
	if o == nil || IsNil(o.MetadataRequirements) {
		return nil, false
	}
	return o.MetadataRequirements, true
}

// HasMetadataRequirements returns a boolean if a field has been set.
func (o *CimdClientEntityLinks) HasMetadataRequirements() bool {
	if o != nil && !IsNil(o.MetadataRequirements) {
		return true
	}

	return false
}

// SetMetadataRequirements gets a reference to the given CimdClientEntityLinksMetadataRequirements and assigns it to the MetadataRequirements field.
func (o *CimdClientEntityLinks) SetMetadataRequirements(v CimdClientEntityLinksMetadataRequirements) {
	o.MetadataRequirements = &v
}

// GetBindings returns the Bindings field value if set, zero value otherwise.
func (o *CimdClientEntityLinks) GetBindings() []CimdClientEntityLinksSelf {
	if o == nil || IsNil(o.Bindings) {
		var ret []CimdClientEntityLinksSelf
		return ret
	}
	return o.Bindings
}

// GetBindingsOk returns a tuple with the Bindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntityLinks) GetBindingsOk() ([]CimdClientEntityLinksSelf, bool) {
	if o == nil || IsNil(o.Bindings) {
		return nil, false
	}
	return o.Bindings, true
}

// HasBindings returns a boolean if a field has been set.
func (o *CimdClientEntityLinks) HasBindings() bool {
	if o != nil && !IsNil(o.Bindings) {
		return true
	}

	return false
}

// SetBindings gets a reference to the given []CimdClientEntityLinksSelf and assigns it to the Bindings field.
func (o *CimdClientEntityLinks) SetBindings(v []CimdClientEntityLinksSelf) {
	o.Bindings = v
}

func (o CimdClientEntityLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CimdClientEntityLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.MetadataRequirements) {
		toSerialize["metadataRequirements"] = o.MetadataRequirements
	}
	if !IsNil(o.Bindings) {
		toSerialize["bindings"] = o.Bindings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CimdClientEntityLinks) UnmarshalJSON(data []byte) (err error) {
	varCimdClientEntityLinks := _CimdClientEntityLinks{}

	err = json.Unmarshal(data, &varCimdClientEntityLinks)

	if err != nil {
		return err
	}

	*o = CimdClientEntityLinks(varCimdClientEntityLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "metadataRequirements")
		delete(additionalProperties, "bindings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCimdClientEntityLinks struct {
	value *CimdClientEntityLinks
	isSet bool
}

func (v NullableCimdClientEntityLinks) Get() *CimdClientEntityLinks {
	return v.value
}

func (v *NullableCimdClientEntityLinks) Set(val *CimdClientEntityLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableCimdClientEntityLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableCimdClientEntityLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCimdClientEntityLinks(val *CimdClientEntityLinks) *NullableCimdClientEntityLinks {
	return &NullableCimdClientEntityLinks{value: val, isSet: true}
}

func (v NullableCimdClientEntityLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCimdClientEntityLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
