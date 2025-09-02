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

// RevokeRefreshTokenHrefObject struct for RevokeRefreshTokenHrefObject
type RevokeRefreshTokenHrefObject struct {
	// Link URI
	Href *string `json:"href,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RevokeRefreshTokenHrefObject RevokeRefreshTokenHrefObject

// NewRevokeRefreshTokenHrefObject instantiates a new RevokeRefreshTokenHrefObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevokeRefreshTokenHrefObject() *RevokeRefreshTokenHrefObject {
	this := RevokeRefreshTokenHrefObject{}
	return &this
}

// NewRevokeRefreshTokenHrefObjectWithDefaults instantiates a new RevokeRefreshTokenHrefObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokeRefreshTokenHrefObjectWithDefaults() *RevokeRefreshTokenHrefObject {
	this := RevokeRefreshTokenHrefObject{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RevokeRefreshTokenHrefObject) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevokeRefreshTokenHrefObject) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RevokeRefreshTokenHrefObject) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RevokeRefreshTokenHrefObject) SetHref(v string) {
	o.Href = &v
}

func (o RevokeRefreshTokenHrefObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RevokeRefreshTokenHrefObject) UnmarshalJSON(bytes []byte) (err error) {
	varRevokeRefreshTokenHrefObject := _RevokeRefreshTokenHrefObject{}

	err = json.Unmarshal(bytes, &varRevokeRefreshTokenHrefObject)
	if err == nil {
		*o = RevokeRefreshTokenHrefObject(varRevokeRefreshTokenHrefObject)
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

type NullableRevokeRefreshTokenHrefObject struct {
	value *RevokeRefreshTokenHrefObject
	isSet bool
}

func (v NullableRevokeRefreshTokenHrefObject) Get() *RevokeRefreshTokenHrefObject {
	return v.value
}

func (v *NullableRevokeRefreshTokenHrefObject) Set(val *RevokeRefreshTokenHrefObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokeRefreshTokenHrefObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokeRefreshTokenHrefObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokeRefreshTokenHrefObject(val *RevokeRefreshTokenHrefObject) *NullableRevokeRefreshTokenHrefObject {
	return &NullableRevokeRefreshTokenHrefObject{value: val, isSet: true}
}

func (v NullableRevokeRefreshTokenHrefObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokeRefreshTokenHrefObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

