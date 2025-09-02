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

// OrgAerialConsent struct for OrgAerialConsent
type OrgAerialConsent struct {
	// The unique ID of the Aerial account
	AccountId string `json:"accountId"`
	AdditionalProperties map[string]interface{}
}

type _OrgAerialConsent OrgAerialConsent

// NewOrgAerialConsent instantiates a new OrgAerialConsent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgAerialConsent(accountId string) *OrgAerialConsent {
	this := OrgAerialConsent{}
	this.AccountId = accountId
	return &this
}

// NewOrgAerialConsentWithDefaults instantiates a new OrgAerialConsent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgAerialConsentWithDefaults() *OrgAerialConsent {
	this := OrgAerialConsent{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *OrgAerialConsent) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *OrgAerialConsent) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *OrgAerialConsent) SetAccountId(v string) {
	o.AccountId = v
}

func (o OrgAerialConsent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountId"] = o.AccountId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OrgAerialConsent) UnmarshalJSON(bytes []byte) (err error) {
	varOrgAerialConsent := _OrgAerialConsent{}

	err = json.Unmarshal(bytes, &varOrgAerialConsent)
	if err == nil {
		*o = OrgAerialConsent(varOrgAerialConsent)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "accountId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOrgAerialConsent struct {
	value *OrgAerialConsent
	isSet bool
}

func (v NullableOrgAerialConsent) Get() *OrgAerialConsent {
	return v.value
}

func (v *NullableOrgAerialConsent) Set(val *OrgAerialConsent) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgAerialConsent) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgAerialConsent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgAerialConsent(val *OrgAerialConsent) *NullableOrgAerialConsent {
	return &NullableOrgAerialConsent{value: val, isSet: true}
}

func (v NullableOrgAerialConsent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgAerialConsent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

