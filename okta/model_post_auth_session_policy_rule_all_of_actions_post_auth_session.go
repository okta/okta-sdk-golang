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

// checks if the PostAuthSessionPolicyRuleAllOfActionsPostAuthSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAuthSessionPolicyRuleAllOfActionsPostAuthSession{}

// PostAuthSessionPolicyRuleAllOfActionsPostAuthSession This object contains a `failureActions` array that defines the specific action to take when the session protection policy detects a failure
type PostAuthSessionPolicyRuleAllOfActionsPostAuthSession struct {
	// An array of objects that define the action. It can be empty or contain two `action` value pairs.
	FailureActions       []PostAuthSessionFailureActionsObject `json:"failureActions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostAuthSessionPolicyRuleAllOfActionsPostAuthSession PostAuthSessionPolicyRuleAllOfActionsPostAuthSession

// NewPostAuthSessionPolicyRuleAllOfActionsPostAuthSession instantiates a new PostAuthSessionPolicyRuleAllOfActionsPostAuthSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAuthSessionPolicyRuleAllOfActionsPostAuthSession() *PostAuthSessionPolicyRuleAllOfActionsPostAuthSession {
	this := PostAuthSessionPolicyRuleAllOfActionsPostAuthSession{}
	return &this
}

// NewPostAuthSessionPolicyRuleAllOfActionsPostAuthSessionWithDefaults instantiates a new PostAuthSessionPolicyRuleAllOfActionsPostAuthSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAuthSessionPolicyRuleAllOfActionsPostAuthSessionWithDefaults() *PostAuthSessionPolicyRuleAllOfActionsPostAuthSession {
	this := PostAuthSessionPolicyRuleAllOfActionsPostAuthSession{}
	return &this
}

// GetFailureActions returns the FailureActions field value if set, zero value otherwise.
func (o *PostAuthSessionPolicyRuleAllOfActionsPostAuthSession) GetFailureActions() []PostAuthSessionFailureActionsObject {
	if o == nil || IsNil(o.FailureActions) {
		var ret []PostAuthSessionFailureActionsObject
		return ret
	}
	return o.FailureActions
}

// GetFailureActionsOk returns a tuple with the FailureActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthSessionPolicyRuleAllOfActionsPostAuthSession) GetFailureActionsOk() ([]PostAuthSessionFailureActionsObject, bool) {
	if o == nil || IsNil(o.FailureActions) {
		return nil, false
	}
	return o.FailureActions, true
}

// HasFailureActions returns a boolean if a field has been set.
func (o *PostAuthSessionPolicyRuleAllOfActionsPostAuthSession) HasFailureActions() bool {
	if o != nil && !IsNil(o.FailureActions) {
		return true
	}

	return false
}

// SetFailureActions gets a reference to the given []PostAuthSessionFailureActionsObject and assigns it to the FailureActions field.
func (o *PostAuthSessionPolicyRuleAllOfActionsPostAuthSession) SetFailureActions(v []PostAuthSessionFailureActionsObject) {
	o.FailureActions = v
}

func (o PostAuthSessionPolicyRuleAllOfActionsPostAuthSession) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAuthSessionPolicyRuleAllOfActionsPostAuthSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FailureActions) {
		toSerialize["failureActions"] = o.FailureActions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostAuthSessionPolicyRuleAllOfActionsPostAuthSession) UnmarshalJSON(data []byte) (err error) {
	varPostAuthSessionPolicyRuleAllOfActionsPostAuthSession := _PostAuthSessionPolicyRuleAllOfActionsPostAuthSession{}

	err = json.Unmarshal(data, &varPostAuthSessionPolicyRuleAllOfActionsPostAuthSession)

	if err != nil {
		return err
	}

	*o = PostAuthSessionPolicyRuleAllOfActionsPostAuthSession(varPostAuthSessionPolicyRuleAllOfActionsPostAuthSession)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "failureActions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostAuthSessionPolicyRuleAllOfActionsPostAuthSession struct {
	value *PostAuthSessionPolicyRuleAllOfActionsPostAuthSession
	isSet bool
}

func (v NullablePostAuthSessionPolicyRuleAllOfActionsPostAuthSession) Get() *PostAuthSessionPolicyRuleAllOfActionsPostAuthSession {
	return v.value
}

func (v *NullablePostAuthSessionPolicyRuleAllOfActionsPostAuthSession) Set(val *PostAuthSessionPolicyRuleAllOfActionsPostAuthSession) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAuthSessionPolicyRuleAllOfActionsPostAuthSession) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAuthSessionPolicyRuleAllOfActionsPostAuthSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAuthSessionPolicyRuleAllOfActionsPostAuthSession(val *PostAuthSessionPolicyRuleAllOfActionsPostAuthSession) *NullablePostAuthSessionPolicyRuleAllOfActionsPostAuthSession {
	return &NullablePostAuthSessionPolicyRuleAllOfActionsPostAuthSession{value: val, isSet: true}
}

func (v NullablePostAuthSessionPolicyRuleAllOfActionsPostAuthSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAuthSessionPolicyRuleAllOfActionsPostAuthSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
