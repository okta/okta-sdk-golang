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

// checks if the SignInPageRoot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignInPageRoot{}

// SignInPageRoot The sign-in page sub-resources. Extends `PageRoot` with the sign-in-page-only `widgetConfigurationSchema` expand value (requires the SIW_CONFIG_JSON_CUSTOMIZATION feature).
type SignInPageRoot struct {
	Embedded             map[string]interface{} `json:"_embedded,omitempty"`
	Links                *PageRootLinks         `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SignInPageRoot SignInPageRoot

// NewSignInPageRoot instantiates a new SignInPageRoot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignInPageRoot() *SignInPageRoot {
	this := SignInPageRoot{}
	return &this
}

// NewSignInPageRootWithDefaults instantiates a new SignInPageRoot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignInPageRootWithDefaults() *SignInPageRoot {
	this := SignInPageRoot{}
	return &this
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *SignInPageRoot) GetEmbedded() map[string]interface{} {
	if o == nil || IsNil(o.Embedded) {
		var ret map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageRoot) GetEmbeddedOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Embedded) {
		return map[string]interface{}{}, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *SignInPageRoot) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]interface{} and assigns it to the Embedded field.
func (o *SignInPageRoot) SetEmbedded(v map[string]interface{}) {
	o.Embedded = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SignInPageRoot) GetLinks() PageRootLinks {
	if o == nil || IsNil(o.Links) {
		var ret PageRootLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageRoot) GetLinksOk() (*PageRootLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SignInPageRoot) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PageRootLinks and assigns it to the Links field.
func (o *SignInPageRoot) SetLinks(v PageRootLinks) {
	o.Links = &v
}

func (o SignInPageRoot) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignInPageRoot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SignInPageRoot) UnmarshalJSON(data []byte) (err error) {
	varSignInPageRoot := _SignInPageRoot{}

	err = json.Unmarshal(data, &varSignInPageRoot)

	if err != nil {
		return err
	}

	*o = SignInPageRoot(varSignInPageRoot)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSignInPageRoot struct {
	value *SignInPageRoot
	isSet bool
}

func (v NullableSignInPageRoot) Get() *SignInPageRoot {
	return v.value
}

func (v *NullableSignInPageRoot) Set(val *SignInPageRoot) {
	v.value = val
	v.isSet = true
}

func (v NullableSignInPageRoot) IsSet() bool {
	return v.isSet
}

func (v *NullableSignInPageRoot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignInPageRoot(val *SignInPageRoot) *NullableSignInPageRoot {
	return &NullableSignInPageRoot{value: val, isSet: true}
}

func (v NullableSignInPageRoot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignInPageRoot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
