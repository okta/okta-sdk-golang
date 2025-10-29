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

// checks if the IdPCsrLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdPCsrLinks{}

// IdPCsrLinks struct for IdPCsrLinks
type IdPCsrLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	// Publish the CSR
	Publish              *HrefObject `json:"publish,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdPCsrLinks IdPCsrLinks

// NewIdPCsrLinks instantiates a new IdPCsrLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdPCsrLinks() *IdPCsrLinks {
	this := IdPCsrLinks{}
	return &this
}

// NewIdPCsrLinksWithDefaults instantiates a new IdPCsrLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdPCsrLinksWithDefaults() *IdPCsrLinks {
	this := IdPCsrLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *IdPCsrLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPCsrLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *IdPCsrLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *IdPCsrLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetPublish returns the Publish field value if set, zero value otherwise.
func (o *IdPCsrLinks) GetPublish() HrefObject {
	if o == nil || IsNil(o.Publish) {
		var ret HrefObject
		return ret
	}
	return *o.Publish
}

// GetPublishOk returns a tuple with the Publish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPCsrLinks) GetPublishOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Publish) {
		return nil, false
	}
	return o.Publish, true
}

// HasPublish returns a boolean if a field has been set.
func (o *IdPCsrLinks) HasPublish() bool {
	if o != nil && !IsNil(o.Publish) {
		return true
	}

	return false
}

// SetPublish gets a reference to the given HrefObject and assigns it to the Publish field.
func (o *IdPCsrLinks) SetPublish(v HrefObject) {
	o.Publish = &v
}

func (o IdPCsrLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdPCsrLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Publish) {
		toSerialize["publish"] = o.Publish
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdPCsrLinks) UnmarshalJSON(data []byte) (err error) {
	varIdPCsrLinks := _IdPCsrLinks{}

	err = json.Unmarshal(data, &varIdPCsrLinks)

	if err != nil {
		return err
	}

	*o = IdPCsrLinks(varIdPCsrLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "publish")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdPCsrLinks struct {
	value *IdPCsrLinks
	isSet bool
}

func (v NullableIdPCsrLinks) Get() *IdPCsrLinks {
	return v.value
}

func (v *NullableIdPCsrLinks) Set(val *IdPCsrLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableIdPCsrLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableIdPCsrLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdPCsrLinks(val *IdPCsrLinks) *NullableIdPCsrLinks {
	return &NullableIdPCsrLinks{value: val, isSet: true}
}

func (v NullableIdPCsrLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdPCsrLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
