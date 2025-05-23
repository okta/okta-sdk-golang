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

// LinkedObject struct for LinkedObject
type LinkedObject struct {
	Associated *LinkedObjectDetails `json:"associated,omitempty"`
	Primary *LinkedObjectDetails `json:"primary,omitempty"`
	Links *LinkedObjectLinksSelf `json:"_links,omitempty"`
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
	if o == nil || o.Associated == nil {
		var ret LinkedObjectDetails
		return ret
	}
	return *o.Associated
}

// GetAssociatedOk returns a tuple with the Associated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedObject) GetAssociatedOk() (*LinkedObjectDetails, bool) {
	if o == nil || o.Associated == nil {
		return nil, false
	}
	return o.Associated, true
}

// HasAssociated returns a boolean if a field has been set.
func (o *LinkedObject) HasAssociated() bool {
	if o != nil && o.Associated != nil {
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
	if o == nil || o.Primary == nil {
		var ret LinkedObjectDetails
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedObject) GetPrimaryOk() (*LinkedObjectDetails, bool) {
	if o == nil || o.Primary == nil {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *LinkedObject) HasPrimary() bool {
	if o != nil && o.Primary != nil {
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
	if o == nil || o.Links == nil {
		var ret LinkedObjectLinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedObject) GetLinksOk() (*LinkedObjectLinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LinkedObject) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinkedObjectLinksSelf and assigns it to the Links field.
func (o *LinkedObject) SetLinks(v LinkedObjectLinksSelf) {
	o.Links = &v
}

func (o LinkedObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Associated != nil {
		toSerialize["associated"] = o.Associated
	}
	if o.Primary != nil {
		toSerialize["primary"] = o.Primary
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkedObject) UnmarshalJSON(bytes []byte) (err error) {
	varLinkedObject := _LinkedObject{}

	err = json.Unmarshal(bytes, &varLinkedObject)
	if err == nil {
		*o = LinkedObject(varLinkedObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "associated")
		delete(additionalProperties, "primary")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

