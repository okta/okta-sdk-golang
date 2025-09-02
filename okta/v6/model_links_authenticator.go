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

// LinksAuthenticator struct for LinksAuthenticator
type LinksAuthenticator struct {
	// Returns information about a specific authenticator. See [Retrieve an authenticator](/openapi/okta-management/management/tag/Authenticator/#tag/Authenticator/operation/getAuthenticator).
	Authenticator *HrefObject `json:"authenticator,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksAuthenticator LinksAuthenticator

// NewLinksAuthenticator instantiates a new LinksAuthenticator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksAuthenticator() *LinksAuthenticator {
	this := LinksAuthenticator{}
	return &this
}

// NewLinksAuthenticatorWithDefaults instantiates a new LinksAuthenticator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksAuthenticatorWithDefaults() *LinksAuthenticator {
	this := LinksAuthenticator{}
	return &this
}

// GetAuthenticator returns the Authenticator field value if set, zero value otherwise.
func (o *LinksAuthenticator) GetAuthenticator() HrefObject {
	if o == nil || o.Authenticator == nil {
		var ret HrefObject
		return ret
	}
	return *o.Authenticator
}

// GetAuthenticatorOk returns a tuple with the Authenticator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksAuthenticator) GetAuthenticatorOk() (*HrefObject, bool) {
	if o == nil || o.Authenticator == nil {
		return nil, false
	}
	return o.Authenticator, true
}

// HasAuthenticator returns a boolean if a field has been set.
func (o *LinksAuthenticator) HasAuthenticator() bool {
	if o != nil && o.Authenticator != nil {
		return true
	}

	return false
}

// SetAuthenticator gets a reference to the given HrefObject and assigns it to the Authenticator field.
func (o *LinksAuthenticator) SetAuthenticator(v HrefObject) {
	o.Authenticator = &v
}

func (o LinksAuthenticator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Authenticator != nil {
		toSerialize["authenticator"] = o.Authenticator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksAuthenticator) UnmarshalJSON(bytes []byte) (err error) {
	varLinksAuthenticator := _LinksAuthenticator{}

	err = json.Unmarshal(bytes, &varLinksAuthenticator)
	if err == nil {
		*o = LinksAuthenticator(varLinksAuthenticator)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authenticator")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinksAuthenticator struct {
	value *LinksAuthenticator
	isSet bool
}

func (v NullableLinksAuthenticator) Get() *LinksAuthenticator {
	return v.value
}

func (v *NullableLinksAuthenticator) Set(val *LinksAuthenticator) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksAuthenticator) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksAuthenticator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksAuthenticator(val *LinksAuthenticator) *NullableLinksAuthenticator {
	return &NullableLinksAuthenticator{value: val, isSet: true}
}

func (v NullableLinksAuthenticator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksAuthenticator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

