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

// AccessPolicyRuleActions struct for AccessPolicyRuleActions
type AccessPolicyRuleActions struct {
	AppSignOn *AccessPolicyRuleApplicationSignOn `json:"appSignOn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyRuleActions AccessPolicyRuleActions

// NewAccessPolicyRuleActions instantiates a new AccessPolicyRuleActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyRuleActions() *AccessPolicyRuleActions {
	this := AccessPolicyRuleActions{}
	return &this
}

// NewAccessPolicyRuleActionsWithDefaults instantiates a new AccessPolicyRuleActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyRuleActionsWithDefaults() *AccessPolicyRuleActions {
	this := AccessPolicyRuleActions{}
	return &this
}

// GetAppSignOn returns the AppSignOn field value if set, zero value otherwise.
func (o *AccessPolicyRuleActions) GetAppSignOn() AccessPolicyRuleApplicationSignOn {
	if o == nil || o.AppSignOn == nil {
		var ret AccessPolicyRuleApplicationSignOn
		return ret
	}
	return *o.AppSignOn
}

// GetAppSignOnOk returns a tuple with the AppSignOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleActions) GetAppSignOnOk() (*AccessPolicyRuleApplicationSignOn, bool) {
	if o == nil || o.AppSignOn == nil {
		return nil, false
	}
	return o.AppSignOn, true
}

// HasAppSignOn returns a boolean if a field has been set.
func (o *AccessPolicyRuleActions) HasAppSignOn() bool {
	if o != nil && o.AppSignOn != nil {
		return true
	}

	return false
}

// SetAppSignOn gets a reference to the given AccessPolicyRuleApplicationSignOn and assigns it to the AppSignOn field.
func (o *AccessPolicyRuleActions) SetAppSignOn(v AccessPolicyRuleApplicationSignOn) {
	o.AppSignOn = &v
}

func (o AccessPolicyRuleActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppSignOn != nil {
		toSerialize["appSignOn"] = o.AppSignOn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessPolicyRuleActions) UnmarshalJSON(bytes []byte) (err error) {
	varAccessPolicyRuleActions := _AccessPolicyRuleActions{}

	err = json.Unmarshal(bytes, &varAccessPolicyRuleActions)
	if err == nil {
		*o = AccessPolicyRuleActions(varAccessPolicyRuleActions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "appSignOn")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAccessPolicyRuleActions struct {
	value *AccessPolicyRuleActions
	isSet bool
}

func (v NullableAccessPolicyRuleActions) Get() *AccessPolicyRuleActions {
	return v.value
}

func (v *NullableAccessPolicyRuleActions) Set(val *AccessPolicyRuleActions) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyRuleActions) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyRuleActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyRuleActions(val *AccessPolicyRuleActions) *NullableAccessPolicyRuleActions {
	return &NullableAccessPolicyRuleActions{value: val, isSet: true}
}

func (v NullableAccessPolicyRuleActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyRuleActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

