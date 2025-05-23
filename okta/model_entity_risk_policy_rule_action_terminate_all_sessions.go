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

// EntityRiskPolicyRuleActionTerminateAllSessions struct for EntityRiskPolicyRuleActionTerminateAllSessions
type EntityRiskPolicyRuleActionTerminateAllSessions struct {
	// This action revokes or terminates all of the user's active sessions.
	Action *string `json:"action,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntityRiskPolicyRuleActionTerminateAllSessions EntityRiskPolicyRuleActionTerminateAllSessions

// NewEntityRiskPolicyRuleActionTerminateAllSessions instantiates a new EntityRiskPolicyRuleActionTerminateAllSessions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityRiskPolicyRuleActionTerminateAllSessions() *EntityRiskPolicyRuleActionTerminateAllSessions {
	this := EntityRiskPolicyRuleActionTerminateAllSessions{}
	return &this
}

// NewEntityRiskPolicyRuleActionTerminateAllSessionsWithDefaults instantiates a new EntityRiskPolicyRuleActionTerminateAllSessions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityRiskPolicyRuleActionTerminateAllSessionsWithDefaults() *EntityRiskPolicyRuleActionTerminateAllSessions {
	this := EntityRiskPolicyRuleActionTerminateAllSessions{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *EntityRiskPolicyRuleActionTerminateAllSessions) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRiskPolicyRuleActionTerminateAllSessions) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *EntityRiskPolicyRuleActionTerminateAllSessions) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *EntityRiskPolicyRuleActionTerminateAllSessions) SetAction(v string) {
	o.Action = &v
}

func (o EntityRiskPolicyRuleActionTerminateAllSessions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntityRiskPolicyRuleActionTerminateAllSessions) UnmarshalJSON(bytes []byte) (err error) {
	varEntityRiskPolicyRuleActionTerminateAllSessions := _EntityRiskPolicyRuleActionTerminateAllSessions{}

	err = json.Unmarshal(bytes, &varEntityRiskPolicyRuleActionTerminateAllSessions)
	if err == nil {
		*o = EntityRiskPolicyRuleActionTerminateAllSessions(varEntityRiskPolicyRuleActionTerminateAllSessions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "action")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntityRiskPolicyRuleActionTerminateAllSessions struct {
	value *EntityRiskPolicyRuleActionTerminateAllSessions
	isSet bool
}

func (v NullableEntityRiskPolicyRuleActionTerminateAllSessions) Get() *EntityRiskPolicyRuleActionTerminateAllSessions {
	return v.value
}

func (v *NullableEntityRiskPolicyRuleActionTerminateAllSessions) Set(val *EntityRiskPolicyRuleActionTerminateAllSessions) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityRiskPolicyRuleActionTerminateAllSessions) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityRiskPolicyRuleActionTerminateAllSessions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityRiskPolicyRuleActionTerminateAllSessions(val *EntityRiskPolicyRuleActionTerminateAllSessions) *NullableEntityRiskPolicyRuleActionTerminateAllSessions {
	return &NullableEntityRiskPolicyRuleActionTerminateAllSessions{value: val, isSet: true}
}

func (v NullableEntityRiskPolicyRuleActionTerminateAllSessions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityRiskPolicyRuleActionTerminateAllSessions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

