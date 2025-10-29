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

// checks if the LinkedObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkedObject{}

// LinkedObject struct for LinkedObject
type LinkedObject struct {
	Associated           *LinkedObjectDetails   `json:"associated,omitempty"`
	Primary              *LinkedObjectDetails   `json:"primary,omitempty"`
	Links                *LinkedObjectLinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkedObject LinkedObject

// NewLinkedObject instantiates a new LinkedObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedObject() *LinkedObject {
	this := LinkedObject{}
	return &this
}

// NewLinkedObjectWithDefaults instantiates a new LinkedObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedObjectWithDefaults() *LinkedObject {
	this := LinkedObject{}
	return &this
}

// GetAssociated returns the Associated field value if set, zero value otherwise.
func (o *LinkedObject) GetAssociated() LinkedObjectDetails {
	if o == nil || IsNil(o.Associated) {
		var ret LinkedObjectDetails
		return ret
	}
	return *o.Associated
}

// GetAssociatedOk returns a tuple with the Associated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedObject) GetAssociatedOk() (*LinkedObjectDetails, bool) {
	if o == nil || IsNil(o.Associated) {
		return nil, false
	}
	return o.Associated, true
}

// HasAssociated returns a boolean if a field has been set.
func (o *LinkedObject) HasAssociated() bool {
	if o != nil && !IsNil(o.Associated) {
		return true
	}

	return false
}

// SetAssociated gets a reference to the given LinkedObjectDetails and assigns it to the Associated field.
func (o *LinkedObject) SetAssociated(v LinkedObjectDetails) {
	o.Associated = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *LinkedObject) GetPrimary() LinkedObjectDetails {
	if o == nil || IsNil(o.Primary) {
		var ret LinkedObjectDetails
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedObject) GetPrimaryOk() (*LinkedObjectDetails, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *LinkedObject) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given LinkedObjectDetails and assigns it to the Primary field.
func (o *LinkedObject) SetPrimary(v LinkedObjectDetails) {
	o.Primary = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *LinkedObject) GetLinks() LinkedObjectLinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinkedObjectLinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedObject) GetLinksOk() (*LinkedObjectLinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LinkedObject) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinkedObjectLinksSelf and assigns it to the Links field.
func (o *LinkedObject) SetLinks(v LinkedObjectLinksSelf) {
	o.Links = &v
}

func (o LinkedObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkedObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Associated) {
		toSerialize["associated"] = o.Associated
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinkedObject) UnmarshalJSON(data []byte) (err error) {
	varLinkedObject := _LinkedObject{}

	err = json.Unmarshal(data, &varLinkedObject)

	if err != nil {
		return err
	}

	*o = LinkedObject(varLinkedObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "associated")
		delete(additionalProperties, "primary")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkedObject struct {
	value *LinkedObject
	isSet bool
}

func (v NullableLinkedObject) Get() *LinkedObject {
	return v.value
}

func (v *NullableLinkedObject) Set(val *LinkedObject) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedObject) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedObject(val *LinkedObject) *NullableLinkedObject {
	return &NullableLinkedObject{value: val, isSet: true}
}

func (v NullableLinkedObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
