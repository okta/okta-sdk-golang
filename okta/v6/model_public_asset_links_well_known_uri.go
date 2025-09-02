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

// PublicAssetLinksWellKnownURI The well-known URI content in a JSON array of objects format
type PublicAssetLinksWellKnownURI struct {
	[]string
}

// NewPublicAssetLinksWellKnownURI instantiates a new PublicAssetLinksWellKnownURI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicAssetLinksWellKnownURI() *PublicAssetLinksWellKnownURI {
	this := PublicAssetLinksWellKnownURI{}
	return &this
}

// NewPublicAssetLinksWellKnownURIWithDefaults instantiates a new PublicAssetLinksWellKnownURI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicAssetLinksWellKnownURIWithDefaults() *PublicAssetLinksWellKnownURI {
	this := PublicAssetLinksWellKnownURI{}
	return &this
}

func (o PublicAssetLinksWellKnownURI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serialized[]string, err[]string := json.Marshal(o.[]string)
	if err[]string != nil {
		return []byte{}, err[]string
	}
	err[]string = json.Unmarshal([]byte(serialized[]string), &toSerialize)
	if err[]string != nil {
		return []byte{}, err[]string
	}
	return json.Marshal(toSerialize)
}

type NullablePublicAssetLinksWellKnownURI struct {
	value *PublicAssetLinksWellKnownURI
	isSet bool
}

func (v NullablePublicAssetLinksWellKnownURI) Get() *PublicAssetLinksWellKnownURI {
	return v.value
}

func (v *NullablePublicAssetLinksWellKnownURI) Set(val *PublicAssetLinksWellKnownURI) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicAssetLinksWellKnownURI) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicAssetLinksWellKnownURI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicAssetLinksWellKnownURI(val *PublicAssetLinksWellKnownURI) *NullablePublicAssetLinksWellKnownURI {
	return &NullablePublicAssetLinksWellKnownURI{value: val, isSet: true}
}

func (v NullablePublicAssetLinksWellKnownURI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicAssetLinksWellKnownURI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

