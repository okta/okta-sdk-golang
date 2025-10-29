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

// checks if the SAMLPayLoadDataAssertionAuthenticationAuthnContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLPayLoadDataAssertionAuthenticationAuthnContext{}

// SAMLPayLoadDataAssertionAuthenticationAuthnContext Details of the authentication methods used for the SAML assertion
type SAMLPayLoadDataAssertionAuthenticationAuthnContext struct {
	// Describes the identity provider's supported authentication context classes
	AuthnContextClassRef *string `json:"authnContextClassRef,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLPayLoadDataAssertionAuthenticationAuthnContext SAMLPayLoadDataAssertionAuthenticationAuthnContext

// NewSAMLPayLoadDataAssertionAuthenticationAuthnContext instantiates a new SAMLPayLoadDataAssertionAuthenticationAuthnContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPayLoadDataAssertionAuthenticationAuthnContext() *SAMLPayLoadDataAssertionAuthenticationAuthnContext {
	this := SAMLPayLoadDataAssertionAuthenticationAuthnContext{}
	return &this
}

// NewSAMLPayLoadDataAssertionAuthenticationAuthnContextWithDefaults instantiates a new SAMLPayLoadDataAssertionAuthenticationAuthnContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPayLoadDataAssertionAuthenticationAuthnContextWithDefaults() *SAMLPayLoadDataAssertionAuthenticationAuthnContext {
	this := SAMLPayLoadDataAssertionAuthenticationAuthnContext{}
	return &this
}

// GetAuthnContextClassRef returns the AuthnContextClassRef field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertionAuthenticationAuthnContext) GetAuthnContextClassRef() string {
	if o == nil || IsNil(o.AuthnContextClassRef) {
		var ret string
		return ret
	}
	return *o.AuthnContextClassRef
}

// GetAuthnContextClassRefOk returns a tuple with the AuthnContextClassRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertionAuthenticationAuthnContext) GetAuthnContextClassRefOk() (*string, bool) {
	if o == nil || IsNil(o.AuthnContextClassRef) {
		return nil, false
	}
	return o.AuthnContextClassRef, true
}

// HasAuthnContextClassRef returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertionAuthenticationAuthnContext) HasAuthnContextClassRef() bool {
	if o != nil && !IsNil(o.AuthnContextClassRef) {
		return true
	}

	return false
}

// SetAuthnContextClassRef gets a reference to the given string and assigns it to the AuthnContextClassRef field.
func (o *SAMLPayLoadDataAssertionAuthenticationAuthnContext) SetAuthnContextClassRef(v string) {
	o.AuthnContextClassRef = &v
}

func (o SAMLPayLoadDataAssertionAuthenticationAuthnContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLPayLoadDataAssertionAuthenticationAuthnContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthnContextClassRef) {
		toSerialize["authnContextClassRef"] = o.AuthnContextClassRef
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SAMLPayLoadDataAssertionAuthenticationAuthnContext) UnmarshalJSON(data []byte) (err error) {
	varSAMLPayLoadDataAssertionAuthenticationAuthnContext := _SAMLPayLoadDataAssertionAuthenticationAuthnContext{}

	err = json.Unmarshal(data, &varSAMLPayLoadDataAssertionAuthenticationAuthnContext)

	if err != nil {
		return err
	}

	*o = SAMLPayLoadDataAssertionAuthenticationAuthnContext(varSAMLPayLoadDataAssertionAuthenticationAuthnContext)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authnContextClassRef")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSAMLPayLoadDataAssertionAuthenticationAuthnContext struct {
	value *SAMLPayLoadDataAssertionAuthenticationAuthnContext
	isSet bool
}

func (v NullableSAMLPayLoadDataAssertionAuthenticationAuthnContext) Get() *SAMLPayLoadDataAssertionAuthenticationAuthnContext {
	return v.value
}

func (v *NullableSAMLPayLoadDataAssertionAuthenticationAuthnContext) Set(val *SAMLPayLoadDataAssertionAuthenticationAuthnContext) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPayLoadDataAssertionAuthenticationAuthnContext) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPayLoadDataAssertionAuthenticationAuthnContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPayLoadDataAssertionAuthenticationAuthnContext(val *SAMLPayLoadDataAssertionAuthenticationAuthnContext) *NullableSAMLPayLoadDataAssertionAuthenticationAuthnContext {
	return &NullableSAMLPayLoadDataAssertionAuthenticationAuthnContext{value: val, isSet: true}
}

func (v NullableSAMLPayLoadDataAssertionAuthenticationAuthnContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPayLoadDataAssertionAuthenticationAuthnContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
