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

// checks if the EntitlementValuesResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementValuesResponseLinks{}

// EntitlementValuesResponseLinks struct for EntitlementValuesResponseLinks
type EntitlementValuesResponseLinks struct {
	Self                 *HrefObjectSelfLink `json:"self,omitempty"`
	Bundle               *BundleLink         `json:"bundle,omitempty"`
	Entitlements         *EntitlementsLink   `json:"entitlements,omitempty"`
	Next                 *HrefObjectNextLink `json:"next,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementValuesResponseLinks EntitlementValuesResponseLinks

// NewEntitlementValuesResponseLinks instantiates a new EntitlementValuesResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementValuesResponseLinks() *EntitlementValuesResponseLinks {
	this := EntitlementValuesResponseLinks{}
	return &this
}

// NewEntitlementValuesResponseLinksWithDefaults instantiates a new EntitlementValuesResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementValuesResponseLinksWithDefaults() *EntitlementValuesResponseLinks {
	this := EntitlementValuesResponseLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *EntitlementValuesResponseLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValuesResponseLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *EntitlementValuesResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *EntitlementValuesResponseLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetBundle returns the Bundle field value if set, zero value otherwise.
func (o *EntitlementValuesResponseLinks) GetBundle() BundleLink {
	if o == nil || IsNil(o.Bundle) {
		var ret BundleLink
		return ret
	}
	return *o.Bundle
}

// GetBundleOk returns a tuple with the Bundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValuesResponseLinks) GetBundleOk() (*BundleLink, bool) {
	if o == nil || IsNil(o.Bundle) {
		return nil, false
	}
	return o.Bundle, true
}

// HasBundle returns a boolean if a field has been set.
func (o *EntitlementValuesResponseLinks) HasBundle() bool {
	if o != nil && !IsNil(o.Bundle) {
		return true
	}

	return false
}

// SetBundle gets a reference to the given BundleLink and assigns it to the Bundle field.
func (o *EntitlementValuesResponseLinks) SetBundle(v BundleLink) {
	o.Bundle = &v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *EntitlementValuesResponseLinks) GetEntitlements() EntitlementsLink {
	if o == nil || IsNil(o.Entitlements) {
		var ret EntitlementsLink
		return ret
	}
	return *o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValuesResponseLinks) GetEntitlementsOk() (*EntitlementsLink, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *EntitlementValuesResponseLinks) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given EntitlementsLink and assigns it to the Entitlements field.
func (o *EntitlementValuesResponseLinks) SetEntitlements(v EntitlementsLink) {
	o.Entitlements = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *EntitlementValuesResponseLinks) GetNext() HrefObjectNextLink {
	if o == nil || IsNil(o.Next) {
		var ret HrefObjectNextLink
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValuesResponseLinks) GetNextOk() (*HrefObjectNextLink, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *EntitlementValuesResponseLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given HrefObjectNextLink and assigns it to the Next field.
func (o *EntitlementValuesResponseLinks) SetNext(v HrefObjectNextLink) {
	o.Next = &v
}

func (o EntitlementValuesResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementValuesResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Bundle) {
		toSerialize["bundle"] = o.Bundle
	}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementValuesResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varEntitlementValuesResponseLinks := _EntitlementValuesResponseLinks{}

	err = json.Unmarshal(data, &varEntitlementValuesResponseLinks)

	if err != nil {
		return err
	}

	*o = EntitlementValuesResponseLinks(varEntitlementValuesResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "bundle")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementValuesResponseLinks struct {
	value *EntitlementValuesResponseLinks
	isSet bool
}

func (v NullableEntitlementValuesResponseLinks) Get() *EntitlementValuesResponseLinks {
	return v.value
}

func (v *NullableEntitlementValuesResponseLinks) Set(val *EntitlementValuesResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementValuesResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementValuesResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementValuesResponseLinks(val *EntitlementValuesResponseLinks) *NullableEntitlementValuesResponseLinks {
	return &NullableEntitlementValuesResponseLinks{value: val, isSet: true}
}

func (v NullableEntitlementValuesResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementValuesResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
