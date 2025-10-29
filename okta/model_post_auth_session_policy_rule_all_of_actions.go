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

// checks if the PostAuthSessionPolicyRuleAllOfActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAuthSessionPolicyRuleAllOfActions{}

// PostAuthSessionPolicyRuleAllOfActions The action to take in response to a failure of the reevaluated global session policy or authentication polices
type PostAuthSessionPolicyRuleAllOfActions struct {
	PostAuthSession      *PostAuthSessionPolicyRuleAllOfActionsPostAuthSession `json:"postAuthSession,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostAuthSessionPolicyRuleAllOfActions PostAuthSessionPolicyRuleAllOfActions

// NewPostAuthSessionPolicyRuleAllOfActions instantiates a new PostAuthSessionPolicyRuleAllOfActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAuthSessionPolicyRuleAllOfActions() *PostAuthSessionPolicyRuleAllOfActions {
	this := PostAuthSessionPolicyRuleAllOfActions{}
	return &this
}

// NewPostAuthSessionPolicyRuleAllOfActionsWithDefaults instantiates a new PostAuthSessionPolicyRuleAllOfActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAuthSessionPolicyRuleAllOfActionsWithDefaults() *PostAuthSessionPolicyRuleAllOfActions {
	this := PostAuthSessionPolicyRuleAllOfActions{}
	return &this
}

// GetPostAuthSession returns the PostAuthSession field value if set, zero value otherwise.
func (o *PostAuthSessionPolicyRuleAllOfActions) GetPostAuthSession() PostAuthSessionPolicyRuleAllOfActionsPostAuthSession {
	if o == nil || IsNil(o.PostAuthSession) {
		var ret PostAuthSessionPolicyRuleAllOfActionsPostAuthSession
		return ret
	}
	return *o.PostAuthSession
}

// GetPostAuthSessionOk returns a tuple with the PostAuthSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthSessionPolicyRuleAllOfActions) GetPostAuthSessionOk() (*PostAuthSessionPolicyRuleAllOfActionsPostAuthSession, bool) {
	if o == nil || IsNil(o.PostAuthSession) {
		return nil, false
	}
	return o.PostAuthSession, true
}

// HasPostAuthSession returns a boolean if a field has been set.
func (o *PostAuthSessionPolicyRuleAllOfActions) HasPostAuthSession() bool {
	if o != nil && !IsNil(o.PostAuthSession) {
		return true
	}

	return false
}

// SetPostAuthSession gets a reference to the given PostAuthSessionPolicyRuleAllOfActionsPostAuthSession and assigns it to the PostAuthSession field.
func (o *PostAuthSessionPolicyRuleAllOfActions) SetPostAuthSession(v PostAuthSessionPolicyRuleAllOfActionsPostAuthSession) {
	o.PostAuthSession = &v
}

func (o PostAuthSessionPolicyRuleAllOfActions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAuthSessionPolicyRuleAllOfActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PostAuthSession) {
		toSerialize["postAuthSession"] = o.PostAuthSession
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostAuthSessionPolicyRuleAllOfActions) UnmarshalJSON(data []byte) (err error) {
	varPostAuthSessionPolicyRuleAllOfActions := _PostAuthSessionPolicyRuleAllOfActions{}

	err = json.Unmarshal(data, &varPostAuthSessionPolicyRuleAllOfActions)

	if err != nil {
		return err
	}

	*o = PostAuthSessionPolicyRuleAllOfActions(varPostAuthSessionPolicyRuleAllOfActions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "postAuthSession")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostAuthSessionPolicyRuleAllOfActions struct {
	value *PostAuthSessionPolicyRuleAllOfActions
	isSet bool
}

func (v NullablePostAuthSessionPolicyRuleAllOfActions) Get() *PostAuthSessionPolicyRuleAllOfActions {
	return v.value
}

func (v *NullablePostAuthSessionPolicyRuleAllOfActions) Set(val *PostAuthSessionPolicyRuleAllOfActions) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAuthSessionPolicyRuleAllOfActions) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAuthSessionPolicyRuleAllOfActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAuthSessionPolicyRuleAllOfActions(val *PostAuthSessionPolicyRuleAllOfActions) *NullablePostAuthSessionPolicyRuleAllOfActions {
	return &NullablePostAuthSessionPolicyRuleAllOfActions{value: val, isSet: true}
}

func (v NullablePostAuthSessionPolicyRuleAllOfActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAuthSessionPolicyRuleAllOfActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
