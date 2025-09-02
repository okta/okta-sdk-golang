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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// BundleEntitlementsResponseLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification
type BundleEntitlementsResponseLinks struct {
	// Link to the next resource
	Next *HrefObject `json:"next,omitempty"`
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	// Link to the bundle resource
	Bundle *HrefObject `json:"bundle,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BundleEntitlementsResponseLinks BundleEntitlementsResponseLinks

// NewBundleEntitlementsResponseLinks instantiates a new BundleEntitlementsResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleEntitlementsResponseLinks() *BundleEntitlementsResponseLinks {
	this := BundleEntitlementsResponseLinks{}
	return &this
}

// NewBundleEntitlementsResponseLinksWithDefaults instantiates a new BundleEntitlementsResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleEntitlementsResponseLinksWithDefaults() *BundleEntitlementsResponseLinks {
	this := BundleEntitlementsResponseLinks{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *BundleEntitlementsResponseLinks) GetNext() HrefObject {
	if o == nil || o.Next == nil {
		var ret HrefObject
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleEntitlementsResponseLinks) GetNextOk() (*HrefObject, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *BundleEntitlementsResponseLinks) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given HrefObject and assigns it to the Next field.
func (o *BundleEntitlementsResponseLinks) SetNext(v HrefObject) {
	o.Next = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *BundleEntitlementsResponseLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleEntitlementsResponseLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *BundleEntitlementsResponseLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *BundleEntitlementsResponseLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetBundle returns the Bundle field value if set, zero value otherwise.
func (o *BundleEntitlementsResponseLinks) GetBundle() HrefObject {
	if o == nil || o.Bundle == nil {
		var ret HrefObject
		return ret
	}
	return *o.Bundle
}

// GetBundleOk returns a tuple with the Bundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleEntitlementsResponseLinks) GetBundleOk() (*HrefObject, bool) {
	if o == nil || o.Bundle == nil {
		return nil, false
	}
	return o.Bundle, true
}

// HasBundle returns a boolean if a field has been set.
func (o *BundleEntitlementsResponseLinks) HasBundle() bool {
	if o != nil && o.Bundle != nil {
		return true
	}

	return false
}

// SetBundle gets a reference to the given HrefObject and assigns it to the Bundle field.
func (o *BundleEntitlementsResponseLinks) SetBundle(v HrefObject) {
	o.Bundle = &v
}

func (o BundleEntitlementsResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Bundle != nil {
		toSerialize["bundle"] = o.Bundle
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BundleEntitlementsResponseLinks) UnmarshalJSON(bytes []byte) (err error) {
	varBundleEntitlementsResponseLinks := _BundleEntitlementsResponseLinks{}

	err = json.Unmarshal(bytes, &varBundleEntitlementsResponseLinks)
	if err == nil {
		*o = BundleEntitlementsResponseLinks(varBundleEntitlementsResponseLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "next")
		delete(additionalProperties, "self")
		delete(additionalProperties, "bundle")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBundleEntitlementsResponseLinks struct {
	value *BundleEntitlementsResponseLinks
	isSet bool
}

func (v NullableBundleEntitlementsResponseLinks) Get() *BundleEntitlementsResponseLinks {
	return v.value
}

func (v *NullableBundleEntitlementsResponseLinks) Set(val *BundleEntitlementsResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleEntitlementsResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleEntitlementsResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleEntitlementsResponseLinks(val *BundleEntitlementsResponseLinks) *NullableBundleEntitlementsResponseLinks {
	return &NullableBundleEntitlementsResponseLinks{value: val, isSet: true}
}

func (v NullableBundleEntitlementsResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleEntitlementsResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

