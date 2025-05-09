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

// LinkedHrefObject struct for LinkedHrefObject
type LinkedHrefObject struct {
	// Link URI
	Href string `json:"href"`
	AdditionalProperties map[string]interface{}
}

type _LinkedHrefObject LinkedHrefObject

// NewLinkedHrefObject instantiates a new LinkedHrefObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedHrefObject(href string) *LinkedHrefObject {
	this := LinkedHrefObject{}
	this.Href = href
	return &this
}

// NewLinkedHrefObjectWithDefaults instantiates a new LinkedHrefObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedHrefObjectWithDefaults() *LinkedHrefObject {
	this := LinkedHrefObject{}
	return &this
}

// GetHref returns the Href field value
func (o *LinkedHrefObject) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *LinkedHrefObject) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *LinkedHrefObject) SetHref(v string) {
	o.Href = v
}

func (o LinkedHrefObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["href"] = o.Href
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkedHrefObject) UnmarshalJSON(bytes []byte) (err error) {
	varLinkedHrefObject := _LinkedHrefObject{}

	err = json.Unmarshal(bytes, &varLinkedHrefObject)
	if err == nil {
		*o = LinkedHrefObject(varLinkedHrefObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinkedHrefObject struct {
	value *LinkedHrefObject
	isSet bool
}

func (v NullableLinkedHrefObject) Get() *LinkedHrefObject {
	return v.value
}

func (v *NullableLinkedHrefObject) Set(val *LinkedHrefObject) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedHrefObject) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedHrefObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedHrefObject(val *LinkedHrefObject) *NullableLinkedHrefObject {
	return &NullableLinkedHrefObject{value: val, isSet: true}
}

func (v NullableLinkedHrefObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedHrefObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

