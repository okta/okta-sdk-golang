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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the SAMLPayLoadDataAssertionAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLPayLoadDataAssertionAuthentication{}

// SAMLPayLoadDataAssertionAuthentication Provides a JSON representation of the `<saml:AuthnStatement>` element of the SAML assertion
type SAMLPayLoadDataAssertionAuthentication struct {
	// The unique identifier describing the assertion statement
	SessionIndex         *string                                             `json:"sessionIndex,omitempty"`
	AuthnContext         *SAMLPayLoadDataAssertionAuthenticationAuthnContext `json:"authnContext,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLPayLoadDataAssertionAuthentication SAMLPayLoadDataAssertionAuthentication

// NewSAMLPayLoadDataAssertionAuthentication instantiates a new SAMLPayLoadDataAssertionAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPayLoadDataAssertionAuthentication() *SAMLPayLoadDataAssertionAuthentication {
	this := SAMLPayLoadDataAssertionAuthentication{}
	return &this
}

// NewSAMLPayLoadDataAssertionAuthenticationWithDefaults instantiates a new SAMLPayLoadDataAssertionAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPayLoadDataAssertionAuthenticationWithDefaults() *SAMLPayLoadDataAssertionAuthentication {
	this := SAMLPayLoadDataAssertionAuthentication{}
	return &this
}

// GetSessionIndex returns the SessionIndex field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertionAuthentication) GetSessionIndex() string {
	if o == nil || IsNil(o.SessionIndex) {
		var ret string
		return ret
	}
	return *o.SessionIndex
}

// GetSessionIndexOk returns a tuple with the SessionIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertionAuthentication) GetSessionIndexOk() (*string, bool) {
	if o == nil || IsNil(o.SessionIndex) {
		return nil, false
	}
	return o.SessionIndex, true
}

// HasSessionIndex returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertionAuthentication) HasSessionIndex() bool {
	if o != nil && !IsNil(o.SessionIndex) {
		return true
	}

	return false
}

// SetSessionIndex gets a reference to the given string and assigns it to the SessionIndex field.
func (o *SAMLPayLoadDataAssertionAuthentication) SetSessionIndex(v string) {
	o.SessionIndex = &v
}

// GetAuthnContext returns the AuthnContext field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertionAuthentication) GetAuthnContext() SAMLPayLoadDataAssertionAuthenticationAuthnContext {
	if o == nil || IsNil(o.AuthnContext) {
		var ret SAMLPayLoadDataAssertionAuthenticationAuthnContext
		return ret
	}
	return *o.AuthnContext
}

// GetAuthnContextOk returns a tuple with the AuthnContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertionAuthentication) GetAuthnContextOk() (*SAMLPayLoadDataAssertionAuthenticationAuthnContext, bool) {
	if o == nil || IsNil(o.AuthnContext) {
		return nil, false
	}
	return o.AuthnContext, true
}

// HasAuthnContext returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertionAuthentication) HasAuthnContext() bool {
	if o != nil && !IsNil(o.AuthnContext) {
		return true
	}

	return false
}

// SetAuthnContext gets a reference to the given SAMLPayLoadDataAssertionAuthenticationAuthnContext and assigns it to the AuthnContext field.
func (o *SAMLPayLoadDataAssertionAuthentication) SetAuthnContext(v SAMLPayLoadDataAssertionAuthenticationAuthnContext) {
	o.AuthnContext = &v
}

func (o SAMLPayLoadDataAssertionAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLPayLoadDataAssertionAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SessionIndex) {
		toSerialize["sessionIndex"] = o.SessionIndex
	}
	if !IsNil(o.AuthnContext) {
		toSerialize["authnContext"] = o.AuthnContext
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SAMLPayLoadDataAssertionAuthentication) UnmarshalJSON(data []byte) (err error) {
	varSAMLPayLoadDataAssertionAuthentication := _SAMLPayLoadDataAssertionAuthentication{}

	err = json.Unmarshal(data, &varSAMLPayLoadDataAssertionAuthentication)

	if err != nil {
		return err
	}

	*o = SAMLPayLoadDataAssertionAuthentication(varSAMLPayLoadDataAssertionAuthentication)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sessionIndex")
		delete(additionalProperties, "authnContext")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSAMLPayLoadDataAssertionAuthentication struct {
	value *SAMLPayLoadDataAssertionAuthentication
	isSet bool
}

func (v NullableSAMLPayLoadDataAssertionAuthentication) Get() *SAMLPayLoadDataAssertionAuthentication {
	return v.value
}

func (v *NullableSAMLPayLoadDataAssertionAuthentication) Set(val *SAMLPayLoadDataAssertionAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPayLoadDataAssertionAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPayLoadDataAssertionAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPayLoadDataAssertionAuthentication(val *SAMLPayLoadDataAssertionAuthentication) *NullableSAMLPayLoadDataAssertionAuthentication {
	return &NullableSAMLPayLoadDataAssertionAuthentication{value: val, isSet: true}
}

func (v NullableSAMLPayLoadDataAssertionAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPayLoadDataAssertionAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
