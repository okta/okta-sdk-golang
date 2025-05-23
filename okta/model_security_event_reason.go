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

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// SecurityEventReason struct for SecurityEventReason
type SecurityEventReason struct {
	// The event reason in English
	En string `json:"en"`
	AdditionalProperties map[string]interface{}
}

type _SecurityEventReason SecurityEventReason

// NewSecurityEventReason instantiates a new SecurityEventReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEventReason(en string) *SecurityEventReason {
	this := SecurityEventReason{}
	this.En = en
	return &this
}

// NewSecurityEventReasonWithDefaults instantiates a new SecurityEventReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEventReasonWithDefaults() *SecurityEventReason {
	this := SecurityEventReason{}
	return &this
}

// GetEn returns the En field value
func (o *SecurityEventReason) GetEn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.En
}

// GetEnOk returns a tuple with the En field value
// and a boolean to check if the value has been set.
func (o *SecurityEventReason) GetEnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.En, true
}

// SetEn sets field value
func (o *SecurityEventReason) SetEn(v string) {
	o.En = v
}

func (o SecurityEventReason) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["en"] = o.En
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityEventReason) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityEventReason := _SecurityEventReason{}

	err = json.Unmarshal(bytes, &varSecurityEventReason)
	if err == nil {
		*o = SecurityEventReason(varSecurityEventReason)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "en")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityEventReason struct {
	value *SecurityEventReason
	isSet bool
}

func (v NullableSecurityEventReason) Get() *SecurityEventReason {
	return v.value
}

func (v *NullableSecurityEventReason) Set(val *SecurityEventReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEventReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEventReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEventReason(val *SecurityEventReason) *NullableSecurityEventReason {
	return &NullableSecurityEventReason{value: val, isSet: true}
}

func (v NullableSecurityEventReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEventReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

