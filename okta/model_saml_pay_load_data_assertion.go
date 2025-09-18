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

// checks if the SAMLPayLoadDataAssertion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLPayLoadDataAssertion{}

// SAMLPayLoadDataAssertion Details of the SAML assertion that was generated
type SAMLPayLoadDataAssertion struct {
	Subject        *SAMLPayLoadDataAssertionSubject        `json:"subject,omitempty"`
	Authentication *SAMLPayLoadDataAssertionAuthentication `json:"authentication,omitempty"`
	Conditions     *SAMLPayLoadDataAssertionConditions     `json:"conditions,omitempty"`
	// Provides a JSON representation of the `<saml:AttributeStatement>` element contained in the generated SAML assertion. Contains any optional SAML attribute statements that you have defined for the app using the Admin Console's **SAML Settings**.
	Claims               *map[string]SAMLPayLoadDataAssertionClaimsValue `json:"claims,omitempty"`
	Lifetime             *SAMLPayLoadDataAssertionLifetime               `json:"lifetime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLPayLoadDataAssertion SAMLPayLoadDataAssertion

// NewSAMLPayLoadDataAssertion instantiates a new SAMLPayLoadDataAssertion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPayLoadDataAssertion() *SAMLPayLoadDataAssertion {
	this := SAMLPayLoadDataAssertion{}
	return &this
}

// NewSAMLPayLoadDataAssertionWithDefaults instantiates a new SAMLPayLoadDataAssertion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPayLoadDataAssertionWithDefaults() *SAMLPayLoadDataAssertion {
	this := SAMLPayLoadDataAssertion{}
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertion) GetSubject() SAMLPayLoadDataAssertionSubject {
	if o == nil || IsNil(o.Subject) {
		var ret SAMLPayLoadDataAssertionSubject
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertion) GetSubjectOk() (*SAMLPayLoadDataAssertionSubject, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertion) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given SAMLPayLoadDataAssertionSubject and assigns it to the Subject field.
func (o *SAMLPayLoadDataAssertion) SetSubject(v SAMLPayLoadDataAssertionSubject) {
	o.Subject = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertion) GetAuthentication() SAMLPayLoadDataAssertionAuthentication {
	if o == nil || IsNil(o.Authentication) {
		var ret SAMLPayLoadDataAssertionAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertion) GetAuthenticationOk() (*SAMLPayLoadDataAssertionAuthentication, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertion) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given SAMLPayLoadDataAssertionAuthentication and assigns it to the Authentication field.
func (o *SAMLPayLoadDataAssertion) SetAuthentication(v SAMLPayLoadDataAssertionAuthentication) {
	o.Authentication = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertion) GetConditions() SAMLPayLoadDataAssertionConditions {
	if o == nil || IsNil(o.Conditions) {
		var ret SAMLPayLoadDataAssertionConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertion) GetConditionsOk() (*SAMLPayLoadDataAssertionConditions, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertion) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given SAMLPayLoadDataAssertionConditions and assigns it to the Conditions field.
func (o *SAMLPayLoadDataAssertion) SetConditions(v SAMLPayLoadDataAssertionConditions) {
	o.Conditions = &v
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertion) GetClaims() map[string]SAMLPayLoadDataAssertionClaimsValue {
	if o == nil || IsNil(o.Claims) {
		var ret map[string]SAMLPayLoadDataAssertionClaimsValue
		return ret
	}
	return *o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertion) GetClaimsOk() (*map[string]SAMLPayLoadDataAssertionClaimsValue, bool) {
	if o == nil || IsNil(o.Claims) {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertion) HasClaims() bool {
	if o != nil && !IsNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given map[string]SAMLPayLoadDataAssertionClaimsValue and assigns it to the Claims field.
func (o *SAMLPayLoadDataAssertion) SetClaims(v map[string]SAMLPayLoadDataAssertionClaimsValue) {
	o.Claims = &v
}

// GetLifetime returns the Lifetime field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertion) GetLifetime() SAMLPayLoadDataAssertionLifetime {
	if o == nil || IsNil(o.Lifetime) {
		var ret SAMLPayLoadDataAssertionLifetime
		return ret
	}
	return *o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertion) GetLifetimeOk() (*SAMLPayLoadDataAssertionLifetime, bool) {
	if o == nil || IsNil(o.Lifetime) {
		return nil, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertion) HasLifetime() bool {
	if o != nil && !IsNil(o.Lifetime) {
		return true
	}

	return false
}

// SetLifetime gets a reference to the given SAMLPayLoadDataAssertionLifetime and assigns it to the Lifetime field.
func (o *SAMLPayLoadDataAssertion) SetLifetime(v SAMLPayLoadDataAssertionLifetime) {
	o.Lifetime = &v
}

func (o SAMLPayLoadDataAssertion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLPayLoadDataAssertion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Claims) {
		toSerialize["claims"] = o.Claims
	}
	if !IsNil(o.Lifetime) {
		toSerialize["lifetime"] = o.Lifetime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SAMLPayLoadDataAssertion) UnmarshalJSON(data []byte) (err error) {
	varSAMLPayLoadDataAssertion := _SAMLPayLoadDataAssertion{}

	err = json.Unmarshal(data, &varSAMLPayLoadDataAssertion)

	if err != nil {
		return err
	}

	*o = SAMLPayLoadDataAssertion(varSAMLPayLoadDataAssertion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "subject")
		delete(additionalProperties, "authentication")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "claims")
		delete(additionalProperties, "lifetime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSAMLPayLoadDataAssertion struct {
	value *SAMLPayLoadDataAssertion
	isSet bool
}

func (v NullableSAMLPayLoadDataAssertion) Get() *SAMLPayLoadDataAssertion {
	return v.value
}

func (v *NullableSAMLPayLoadDataAssertion) Set(val *SAMLPayLoadDataAssertion) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPayLoadDataAssertion) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPayLoadDataAssertion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPayLoadDataAssertion(val *SAMLPayLoadDataAssertion) *NullableSAMLPayLoadDataAssertion {
	return &NullableSAMLPayLoadDataAssertion{value: val, isSet: true}
}

func (v NullableSAMLPayLoadDataAssertion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPayLoadDataAssertion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
