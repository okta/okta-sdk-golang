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
	"fmt"
)

// AuthenticationMethod struct for AuthenticationMethod
type AuthenticationMethod struct {
	// Indicates if any secrets or private keys used during authentication must be hardware protected and not exportable. This property is only set for `POSSESSION` constraints.
	HardwareProtection *string `json:"hardwareProtection,omitempty"`
	// An ID that identifies the authenticator
	Id *string `json:"id,omitempty"`
	// A label that identifies the authenticator
	Key string `json:"key"`
	// Specifies the method used for the authenticator
	Method string `json:"method"`
	// Indicates if phishing-resistant Factors are required. This property is only set for `POSSESSION` constraints
	PhishingResistant *string `json:"phishingResistant,omitempty"`
	// Indicates if a user is required to be verified with a verification method.
	UserVerification *string `json:"userVerification,omitempty"`
	// Indicates which methods can be used for user verification. `userVerificationMethods` can only be used when `userVerification` is `REQUIRED`. `BIOMETRICS` is currently the only supported method.
	UserVerificationMethods []string `json:"userVerificationMethods,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticationMethod AuthenticationMethod

// NewAuthenticationMethod instantiates a new AuthenticationMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationMethod(key string, method string) *AuthenticationMethod {
	this := AuthenticationMethod{}
	var hardwareProtection string = "OPTIONAL"
	this.HardwareProtection = &hardwareProtection
	this.Key = key
	this.Method = method
	var phishingResistant string = "OPTIONAL"
	this.PhishingResistant = &phishingResistant
	var userVerification string = "OPTIONAL"
	this.UserVerification = &userVerification
	return &this
}

// NewAuthenticationMethodWithDefaults instantiates a new AuthenticationMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationMethodWithDefaults() *AuthenticationMethod {
	this := AuthenticationMethod{}
	var hardwareProtection string = "OPTIONAL"
	this.HardwareProtection = &hardwareProtection
	var phishingResistant string = "OPTIONAL"
	this.PhishingResistant = &phishingResistant
	var userVerification string = "OPTIONAL"
	this.UserVerification = &userVerification
	return &this
}

// GetHardwareProtection returns the HardwareProtection field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetHardwareProtection() string {
	if o == nil || o.HardwareProtection == nil {
		var ret string
		return ret
	}
	return *o.HardwareProtection
}

// GetHardwareProtectionOk returns a tuple with the HardwareProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetHardwareProtectionOk() (*string, bool) {
	if o == nil || o.HardwareProtection == nil {
		return nil, false
	}
	return o.HardwareProtection, true
}

// HasHardwareProtection returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasHardwareProtection() bool {
	if o != nil && o.HardwareProtection != nil {
		return true
	}

	return false
}

// SetHardwareProtection gets a reference to the given string and assigns it to the HardwareProtection field.
func (o *AuthenticationMethod) SetHardwareProtection(v string) {
	o.HardwareProtection = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthenticationMethod) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value
func (o *AuthenticationMethod) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *AuthenticationMethod) SetKey(v string) {
	o.Key = v
}

// GetMethod returns the Method field value
func (o *AuthenticationMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *AuthenticationMethod) SetMethod(v string) {
	o.Method = v
}

// GetPhishingResistant returns the PhishingResistant field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetPhishingResistant() string {
	if o == nil || o.PhishingResistant == nil {
		var ret string
		return ret
	}
	return *o.PhishingResistant
}

// GetPhishingResistantOk returns a tuple with the PhishingResistant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetPhishingResistantOk() (*string, bool) {
	if o == nil || o.PhishingResistant == nil {
		return nil, false
	}
	return o.PhishingResistant, true
}

// HasPhishingResistant returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasPhishingResistant() bool {
	if o != nil && o.PhishingResistant != nil {
		return true
	}

	return false
}

// SetPhishingResistant gets a reference to the given string and assigns it to the PhishingResistant field.
func (o *AuthenticationMethod) SetPhishingResistant(v string) {
	o.PhishingResistant = &v
}

// GetUserVerification returns the UserVerification field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetUserVerification() string {
	if o == nil || o.UserVerification == nil {
		var ret string
		return ret
	}
	return *o.UserVerification
}

// GetUserVerificationOk returns a tuple with the UserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetUserVerificationOk() (*string, bool) {
	if o == nil || o.UserVerification == nil {
		return nil, false
	}
	return o.UserVerification, true
}

// HasUserVerification returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasUserVerification() bool {
	if o != nil && o.UserVerification != nil {
		return true
	}

	return false
}

// SetUserVerification gets a reference to the given string and assigns it to the UserVerification field.
func (o *AuthenticationMethod) SetUserVerification(v string) {
	o.UserVerification = &v
}

// GetUserVerificationMethods returns the UserVerificationMethods field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetUserVerificationMethods() []string {
	if o == nil || o.UserVerificationMethods == nil {
		var ret []string
		return ret
	}
	return o.UserVerificationMethods
}

// GetUserVerificationMethodsOk returns a tuple with the UserVerificationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetUserVerificationMethodsOk() ([]string, bool) {
	if o == nil || o.UserVerificationMethods == nil {
		return nil, false
	}
	return o.UserVerificationMethods, true
}

// HasUserVerificationMethods returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasUserVerificationMethods() bool {
	if o != nil && o.UserVerificationMethods != nil {
		return true
	}

	return false
}

// SetUserVerificationMethods gets a reference to the given []string and assigns it to the UserVerificationMethods field.
func (o *AuthenticationMethod) SetUserVerificationMethods(v []string) {
	o.UserVerificationMethods = v
}

func (o AuthenticationMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HardwareProtection != nil {
		toSerialize["hardwareProtection"] = o.HardwareProtection
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if o.PhishingResistant != nil {
		toSerialize["phishingResistant"] = o.PhishingResistant
	}
	if o.UserVerification != nil {
		toSerialize["userVerification"] = o.UserVerification
	}
	if o.UserVerificationMethods != nil {
		toSerialize["userVerificationMethods"] = o.UserVerificationMethods
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticationMethod) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticationMethod := _AuthenticationMethod{}

	err = json.Unmarshal(bytes, &varAuthenticationMethod)
	if err == nil {
		*o = AuthenticationMethod(varAuthenticationMethod)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "hardwareProtection")
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "method")
		delete(additionalProperties, "phishingResistant")
		delete(additionalProperties, "userVerification")
		delete(additionalProperties, "userVerificationMethods")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticationMethod struct {
	value *AuthenticationMethod
	isSet bool
}

func (v NullableAuthenticationMethod) Get() *AuthenticationMethod {
	return v.value
}

func (v *NullableAuthenticationMethod) Set(val *AuthenticationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationMethod(val *AuthenticationMethod) *NullableAuthenticationMethod {
	return &NullableAuthenticationMethod{value: val, isSet: true}
}

func (v NullableAuthenticationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

