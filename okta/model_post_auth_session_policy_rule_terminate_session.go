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

// checks if the PostAuthSessionPolicyRuleTerminateSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAuthSessionPolicyRuleTerminateSession{}

// PostAuthSessionPolicyRuleTerminateSession struct for PostAuthSessionPolicyRuleTerminateSession
type PostAuthSessionPolicyRuleTerminateSession struct {
	// The action to take when the session protection policy detects a failure.
	Action               *string `json:"action,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostAuthSessionPolicyRuleTerminateSession PostAuthSessionPolicyRuleTerminateSession

// NewPostAuthSessionPolicyRuleTerminateSession instantiates a new PostAuthSessionPolicyRuleTerminateSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAuthSessionPolicyRuleTerminateSession() *PostAuthSessionPolicyRuleTerminateSession {
	this := PostAuthSessionPolicyRuleTerminateSession{}
	return &this
}

// NewPostAuthSessionPolicyRuleTerminateSessionWithDefaults instantiates a new PostAuthSessionPolicyRuleTerminateSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAuthSessionPolicyRuleTerminateSessionWithDefaults() *PostAuthSessionPolicyRuleTerminateSession {
	this := PostAuthSessionPolicyRuleTerminateSession{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PostAuthSessionPolicyRuleTerminateSession) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthSessionPolicyRuleTerminateSession) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PostAuthSessionPolicyRuleTerminateSession) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PostAuthSessionPolicyRuleTerminateSession) SetAction(v string) {
	o.Action = &v
}

func (o PostAuthSessionPolicyRuleTerminateSession) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAuthSessionPolicyRuleTerminateSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostAuthSessionPolicyRuleTerminateSession) UnmarshalJSON(data []byte) (err error) {
	varPostAuthSessionPolicyRuleTerminateSession := _PostAuthSessionPolicyRuleTerminateSession{}

	err = json.Unmarshal(data, &varPostAuthSessionPolicyRuleTerminateSession)

	if err != nil {
		return err
	}

	*o = PostAuthSessionPolicyRuleTerminateSession(varPostAuthSessionPolicyRuleTerminateSession)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostAuthSessionPolicyRuleTerminateSession struct {
	value *PostAuthSessionPolicyRuleTerminateSession
	isSet bool
}

func (v NullablePostAuthSessionPolicyRuleTerminateSession) Get() *PostAuthSessionPolicyRuleTerminateSession {
	return v.value
}

func (v *NullablePostAuthSessionPolicyRuleTerminateSession) Set(val *PostAuthSessionPolicyRuleTerminateSession) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAuthSessionPolicyRuleTerminateSession) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAuthSessionPolicyRuleTerminateSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAuthSessionPolicyRuleTerminateSession(val *PostAuthSessionPolicyRuleTerminateSession) *NullablePostAuthSessionPolicyRuleTerminateSession {
	return &NullablePostAuthSessionPolicyRuleTerminateSession{value: val, isSet: true}
}

func (v NullablePostAuthSessionPolicyRuleTerminateSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAuthSessionPolicyRuleTerminateSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
