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

// OktaSignOnPolicyRuleActions struct for OktaSignOnPolicyRuleActions
type OktaSignOnPolicyRuleActions struct {
	Signon *OktaSignOnPolicyRuleSignonActions `json:"signon,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaSignOnPolicyRuleActions OktaSignOnPolicyRuleActions

// NewOktaSignOnPolicyRuleActions instantiates a new OktaSignOnPolicyRuleActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaSignOnPolicyRuleActions() *OktaSignOnPolicyRuleActions {
	this := OktaSignOnPolicyRuleActions{}
	return &this
}

// NewOktaSignOnPolicyRuleActionsWithDefaults instantiates a new OktaSignOnPolicyRuleActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaSignOnPolicyRuleActionsWithDefaults() *OktaSignOnPolicyRuleActions {
	this := OktaSignOnPolicyRuleActions{}
	return &this
}

// GetSignon returns the Signon field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleActions) GetSignon() OktaSignOnPolicyRuleSignonActions {
	if o == nil || o.Signon == nil {
		var ret OktaSignOnPolicyRuleSignonActions
		return ret
	}
	return *o.Signon
}

// GetSignonOk returns a tuple with the Signon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleActions) GetSignonOk() (*OktaSignOnPolicyRuleSignonActions, bool) {
	if o == nil || o.Signon == nil {
		return nil, false
	}
	return o.Signon, true
}

// HasSignon returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleActions) HasSignon() bool {
	if o != nil && o.Signon != nil {
		return true
	}

	return false
}

// SetSignon gets a reference to the given OktaSignOnPolicyRuleSignonActions and assigns it to the Signon field.
func (o *OktaSignOnPolicyRuleActions) SetSignon(v OktaSignOnPolicyRuleSignonActions) {
	o.Signon = &v
}

func (o OktaSignOnPolicyRuleActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Signon != nil {
		toSerialize["signon"] = o.Signon
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OktaSignOnPolicyRuleActions) UnmarshalJSON(bytes []byte) (err error) {
	varOktaSignOnPolicyRuleActions := _OktaSignOnPolicyRuleActions{}

	err = json.Unmarshal(bytes, &varOktaSignOnPolicyRuleActions)
	if err == nil {
		*o = OktaSignOnPolicyRuleActions(varOktaSignOnPolicyRuleActions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "signon")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOktaSignOnPolicyRuleActions struct {
	value *OktaSignOnPolicyRuleActions
	isSet bool
}

func (v NullableOktaSignOnPolicyRuleActions) Get() *OktaSignOnPolicyRuleActions {
	return v.value
}

func (v *NullableOktaSignOnPolicyRuleActions) Set(val *OktaSignOnPolicyRuleActions) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaSignOnPolicyRuleActions) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaSignOnPolicyRuleActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaSignOnPolicyRuleActions(val *OktaSignOnPolicyRuleActions) *NullableOktaSignOnPolicyRuleActions {
	return &NullableOktaSignOnPolicyRuleActions{value: val, isSet: true}
}

func (v NullableOktaSignOnPolicyRuleActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaSignOnPolicyRuleActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

