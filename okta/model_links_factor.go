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

// LinksFactor struct for LinksFactor
type LinksFactor struct {
	Factor *LinksFactorFactor `json:"factor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksFactor LinksFactor

// NewLinksFactor instantiates a new LinksFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksFactor() *LinksFactor {
	this := LinksFactor{}
	return &this
}

// NewLinksFactorWithDefaults instantiates a new LinksFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksFactorWithDefaults() *LinksFactor {
	this := LinksFactor{}
	return &this
}

// GetFactor returns the Factor field value if set, zero value otherwise.
func (o *LinksFactor) GetFactor() LinksFactorFactor {
	if o == nil || o.Factor == nil {
		var ret LinksFactorFactor
		return ret
	}
	return *o.Factor
}

// GetFactorOk returns a tuple with the Factor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksFactor) GetFactorOk() (*LinksFactorFactor, bool) {
	if o == nil || o.Factor == nil {
		return nil, false
	}
	return o.Factor, true
}

// HasFactor returns a boolean if a field has been set.
func (o *LinksFactor) HasFactor() bool {
	if o != nil && o.Factor != nil {
		return true
	}

	return false
}

// SetFactor gets a reference to the given LinksFactorFactor and assigns it to the Factor field.
func (o *LinksFactor) SetFactor(v LinksFactorFactor) {
	o.Factor = &v
}

func (o LinksFactor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Factor != nil {
		toSerialize["factor"] = o.Factor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksFactor) UnmarshalJSON(bytes []byte) (err error) {
	varLinksFactor := _LinksFactor{}

	err = json.Unmarshal(bytes, &varLinksFactor)
	if err == nil {
		*o = LinksFactor(varLinksFactor)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "factor")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinksFactor struct {
	value *LinksFactor
	isSet bool
}

func (v NullableLinksFactor) Get() *LinksFactor {
	return v.value
}

func (v *NullableLinksFactor) Set(val *LinksFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksFactor(val *LinksFactor) *NullableLinksFactor {
	return &NullableLinksFactor{value: val, isSet: true}
}

func (v NullableLinksFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

