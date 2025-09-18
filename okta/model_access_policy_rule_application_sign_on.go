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

// checks if the AccessPolicyRuleApplicationSignOn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessPolicyRuleApplicationSignOn{}

// AccessPolicyRuleApplicationSignOn Specifies the results when a user attempts to sign in
type AccessPolicyRuleApplicationSignOn struct {
	Access               *string             `json:"access,omitempty"`
	KeepMeSignedIn       *KeepMeSignedIn     `json:"keepMeSignedIn,omitempty"`
	VerificationMethod   *VerificationMethod `json:"verificationMethod,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyRuleApplicationSignOn AccessPolicyRuleApplicationSignOn

// NewAccessPolicyRuleApplicationSignOn instantiates a new AccessPolicyRuleApplicationSignOn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyRuleApplicationSignOn() *AccessPolicyRuleApplicationSignOn {
	this := AccessPolicyRuleApplicationSignOn{}
	return &this
}

// NewAccessPolicyRuleApplicationSignOnWithDefaults instantiates a new AccessPolicyRuleApplicationSignOn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyRuleApplicationSignOnWithDefaults() *AccessPolicyRuleApplicationSignOn {
	this := AccessPolicyRuleApplicationSignOn{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *AccessPolicyRuleApplicationSignOn) GetAccess() string {
	if o == nil || IsNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleApplicationSignOn) GetAccessOk() (*string, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *AccessPolicyRuleApplicationSignOn) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *AccessPolicyRuleApplicationSignOn) SetAccess(v string) {
	o.Access = &v
}

// GetKeepMeSignedIn returns the KeepMeSignedIn field value if set, zero value otherwise.
func (o *AccessPolicyRuleApplicationSignOn) GetKeepMeSignedIn() KeepMeSignedIn {
	if o == nil || IsNil(o.KeepMeSignedIn) {
		var ret KeepMeSignedIn
		return ret
	}
	return *o.KeepMeSignedIn
}

// GetKeepMeSignedInOk returns a tuple with the KeepMeSignedIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleApplicationSignOn) GetKeepMeSignedInOk() (*KeepMeSignedIn, bool) {
	if o == nil || IsNil(o.KeepMeSignedIn) {
		return nil, false
	}
	return o.KeepMeSignedIn, true
}

// HasKeepMeSignedIn returns a boolean if a field has been set.
func (o *AccessPolicyRuleApplicationSignOn) HasKeepMeSignedIn() bool {
	if o != nil && !IsNil(o.KeepMeSignedIn) {
		return true
	}

	return false
}

// SetKeepMeSignedIn gets a reference to the given KeepMeSignedIn and assigns it to the KeepMeSignedIn field.
func (o *AccessPolicyRuleApplicationSignOn) SetKeepMeSignedIn(v KeepMeSignedIn) {
	o.KeepMeSignedIn = &v
}

// GetVerificationMethod returns the VerificationMethod field value if set, zero value otherwise.
func (o *AccessPolicyRuleApplicationSignOn) GetVerificationMethod() VerificationMethod {
	if o == nil || IsNil(o.VerificationMethod) {
		var ret VerificationMethod
		return ret
	}
	return *o.VerificationMethod
}

// GetVerificationMethodOk returns a tuple with the VerificationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleApplicationSignOn) GetVerificationMethodOk() (*VerificationMethod, bool) {
	if o == nil || IsNil(o.VerificationMethod) {
		return nil, false
	}
	return o.VerificationMethod, true
}

// HasVerificationMethod returns a boolean if a field has been set.
func (o *AccessPolicyRuleApplicationSignOn) HasVerificationMethod() bool {
	if o != nil && !IsNil(o.VerificationMethod) {
		return true
	}

	return false
}

// SetVerificationMethod gets a reference to the given VerificationMethod and assigns it to the VerificationMethod field.
func (o *AccessPolicyRuleApplicationSignOn) SetVerificationMethod(v VerificationMethod) {
	o.VerificationMethod = &v
}

func (o AccessPolicyRuleApplicationSignOn) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessPolicyRuleApplicationSignOn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !IsNil(o.KeepMeSignedIn) {
		toSerialize["keepMeSignedIn"] = o.KeepMeSignedIn
	}
	if !IsNil(o.VerificationMethod) {
		toSerialize["verificationMethod"] = o.VerificationMethod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessPolicyRuleApplicationSignOn) UnmarshalJSON(data []byte) (err error) {
	varAccessPolicyRuleApplicationSignOn := _AccessPolicyRuleApplicationSignOn{}

	err = json.Unmarshal(data, &varAccessPolicyRuleApplicationSignOn)

	if err != nil {
		return err
	}

	*o = AccessPolicyRuleApplicationSignOn(varAccessPolicyRuleApplicationSignOn)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access")
		delete(additionalProperties, "keepMeSignedIn")
		delete(additionalProperties, "verificationMethod")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessPolicyRuleApplicationSignOn struct {
	value *AccessPolicyRuleApplicationSignOn
	isSet bool
}

func (v NullableAccessPolicyRuleApplicationSignOn) Get() *AccessPolicyRuleApplicationSignOn {
	return v.value
}

func (v *NullableAccessPolicyRuleApplicationSignOn) Set(val *AccessPolicyRuleApplicationSignOn) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyRuleApplicationSignOn) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyRuleApplicationSignOn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyRuleApplicationSignOn(val *AccessPolicyRuleApplicationSignOn) *NullableAccessPolicyRuleApplicationSignOn {
	return &NullableAccessPolicyRuleApplicationSignOn{value: val, isSet: true}
}

func (v NullableAccessPolicyRuleApplicationSignOn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyRuleApplicationSignOn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
