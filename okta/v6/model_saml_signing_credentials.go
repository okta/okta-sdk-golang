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

// SamlSigningCredentials Key used for signing requests to the IdP
type SamlSigningCredentials struct {
	// IdP key credential reference to the Okta X.509 signature certificate
	Kid *string `json:"kid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlSigningCredentials SamlSigningCredentials

// NewSamlSigningCredentials instantiates a new SamlSigningCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlSigningCredentials() *SamlSigningCredentials {
	this := SamlSigningCredentials{}
	return &this
}

// NewSamlSigningCredentialsWithDefaults instantiates a new SamlSigningCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlSigningCredentialsWithDefaults() *SamlSigningCredentials {
	this := SamlSigningCredentials{}
	return &this
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *SamlSigningCredentials) GetKid() string {
	if o == nil || o.Kid == nil {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSigningCredentials) GetKidOk() (*string, bool) {
	if o == nil || o.Kid == nil {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *SamlSigningCredentials) HasKid() bool {
	if o != nil && o.Kid != nil {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *SamlSigningCredentials) SetKid(v string) {
	o.Kid = &v
}

func (o SamlSigningCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kid != nil {
		toSerialize["kid"] = o.Kid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SamlSigningCredentials) UnmarshalJSON(bytes []byte) (err error) {
	varSamlSigningCredentials := _SamlSigningCredentials{}

	err = json.Unmarshal(bytes, &varSamlSigningCredentials)
	if err == nil {
		*o = SamlSigningCredentials(varSamlSigningCredentials)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "kid")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSamlSigningCredentials struct {
	value *SamlSigningCredentials
	isSet bool
}

func (v NullableSamlSigningCredentials) Get() *SamlSigningCredentials {
	return v.value
}

func (v *NullableSamlSigningCredentials) Set(val *SamlSigningCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlSigningCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlSigningCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlSigningCredentials(val *SamlSigningCredentials) *NullableSamlSigningCredentials {
	return &NullableSamlSigningCredentials{value: val, isSet: true}
}

func (v NullableSamlSigningCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlSigningCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

