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

// OktaSignOnPolicyRuleSignonActions struct for OktaSignOnPolicyRuleSignonActions
type OktaSignOnPolicyRuleSignonActions struct {
	Access *string `json:"access,omitempty"`
	FactorLifetime *int32 `json:"factorLifetime,omitempty"`
	FactorPromptMode *string `json:"factorPromptMode,omitempty"`
	RememberDeviceByDefault *bool `json:"rememberDeviceByDefault,omitempty"`
	RequireFactor *bool `json:"requireFactor,omitempty"`
	Session *OktaSignOnPolicyRuleSignonSessionActions `json:"session,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaSignOnPolicyRuleSignonActions OktaSignOnPolicyRuleSignonActions

// NewOktaSignOnPolicyRuleSignonActions instantiates a new OktaSignOnPolicyRuleSignonActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaSignOnPolicyRuleSignonActions() *OktaSignOnPolicyRuleSignonActions {
	this := OktaSignOnPolicyRuleSignonActions{}
	var rememberDeviceByDefault bool = false
	this.RememberDeviceByDefault = &rememberDeviceByDefault
	var requireFactor bool = false
	this.RequireFactor = &requireFactor
	return &this
}

// NewOktaSignOnPolicyRuleSignonActionsWithDefaults instantiates a new OktaSignOnPolicyRuleSignonActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaSignOnPolicyRuleSignonActionsWithDefaults() *OktaSignOnPolicyRuleSignonActions {
	this := OktaSignOnPolicyRuleSignonActions{}
	var rememberDeviceByDefault bool = false
	this.RememberDeviceByDefault = &rememberDeviceByDefault
	var requireFactor bool = false
	this.RequireFactor = &requireFactor
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleSignonActions) GetAccess() string {
	if o == nil || o.Access == nil {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleSignonActions) GetAccessOk() (*string, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleSignonActions) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *OktaSignOnPolicyRuleSignonActions) SetAccess(v string) {
	o.Access = &v
}

// GetFactorLifetime returns the FactorLifetime field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleSignonActions) GetFactorLifetime() int32 {
	if o == nil || o.FactorLifetime == nil {
		var ret int32
		return ret
	}
	return *o.FactorLifetime
}

// GetFactorLifetimeOk returns a tuple with the FactorLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleSignonActions) GetFactorLifetimeOk() (*int32, bool) {
	if o == nil || o.FactorLifetime == nil {
		return nil, false
	}
	return o.FactorLifetime, true
}

// HasFactorLifetime returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleSignonActions) HasFactorLifetime() bool {
	if o != nil && o.FactorLifetime != nil {
		return true
	}

	return false
}

// SetFactorLifetime gets a reference to the given int32 and assigns it to the FactorLifetime field.
func (o *OktaSignOnPolicyRuleSignonActions) SetFactorLifetime(v int32) {
	o.FactorLifetime = &v
}

// GetFactorPromptMode returns the FactorPromptMode field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleSignonActions) GetFactorPromptMode() string {
	if o == nil || o.FactorPromptMode == nil {
		var ret string
		return ret
	}
	return *o.FactorPromptMode
}

// GetFactorPromptModeOk returns a tuple with the FactorPromptMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleSignonActions) GetFactorPromptModeOk() (*string, bool) {
	if o == nil || o.FactorPromptMode == nil {
		return nil, false
	}
	return o.FactorPromptMode, true
}

// HasFactorPromptMode returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleSignonActions) HasFactorPromptMode() bool {
	if o != nil && o.FactorPromptMode != nil {
		return true
	}

	return false
}

// SetFactorPromptMode gets a reference to the given string and assigns it to the FactorPromptMode field.
func (o *OktaSignOnPolicyRuleSignonActions) SetFactorPromptMode(v string) {
	o.FactorPromptMode = &v
}

// GetRememberDeviceByDefault returns the RememberDeviceByDefault field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleSignonActions) GetRememberDeviceByDefault() bool {
	if o == nil || o.RememberDeviceByDefault == nil {
		var ret bool
		return ret
	}
	return *o.RememberDeviceByDefault
}

// GetRememberDeviceByDefaultOk returns a tuple with the RememberDeviceByDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleSignonActions) GetRememberDeviceByDefaultOk() (*bool, bool) {
	if o == nil || o.RememberDeviceByDefault == nil {
		return nil, false
	}
	return o.RememberDeviceByDefault, true
}

// HasRememberDeviceByDefault returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleSignonActions) HasRememberDeviceByDefault() bool {
	if o != nil && o.RememberDeviceByDefault != nil {
		return true
	}

	return false
}

// SetRememberDeviceByDefault gets a reference to the given bool and assigns it to the RememberDeviceByDefault field.
func (o *OktaSignOnPolicyRuleSignonActions) SetRememberDeviceByDefault(v bool) {
	o.RememberDeviceByDefault = &v
}

// GetRequireFactor returns the RequireFactor field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleSignonActions) GetRequireFactor() bool {
	if o == nil || o.RequireFactor == nil {
		var ret bool
		return ret
	}
	return *o.RequireFactor
}

// GetRequireFactorOk returns a tuple with the RequireFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleSignonActions) GetRequireFactorOk() (*bool, bool) {
	if o == nil || o.RequireFactor == nil {
		return nil, false
	}
	return o.RequireFactor, true
}

// HasRequireFactor returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleSignonActions) HasRequireFactor() bool {
	if o != nil && o.RequireFactor != nil {
		return true
	}

	return false
}

// SetRequireFactor gets a reference to the given bool and assigns it to the RequireFactor field.
func (o *OktaSignOnPolicyRuleSignonActions) SetRequireFactor(v bool) {
	o.RequireFactor = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleSignonActions) GetSession() OktaSignOnPolicyRuleSignonSessionActions {
	if o == nil || o.Session == nil {
		var ret OktaSignOnPolicyRuleSignonSessionActions
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleSignonActions) GetSessionOk() (*OktaSignOnPolicyRuleSignonSessionActions, bool) {
	if o == nil || o.Session == nil {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleSignonActions) HasSession() bool {
	if o != nil && o.Session != nil {
		return true
	}

	return false
}

// SetSession gets a reference to the given OktaSignOnPolicyRuleSignonSessionActions and assigns it to the Session field.
func (o *OktaSignOnPolicyRuleSignonActions) SetSession(v OktaSignOnPolicyRuleSignonSessionActions) {
	o.Session = &v
}

func (o OktaSignOnPolicyRuleSignonActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.FactorLifetime != nil {
		toSerialize["factorLifetime"] = o.FactorLifetime
	}
	if o.FactorPromptMode != nil {
		toSerialize["factorPromptMode"] = o.FactorPromptMode
	}
	if o.RememberDeviceByDefault != nil {
		toSerialize["rememberDeviceByDefault"] = o.RememberDeviceByDefault
	}
	if o.RequireFactor != nil {
		toSerialize["requireFactor"] = o.RequireFactor
	}
	if o.Session != nil {
		toSerialize["session"] = o.Session
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OktaSignOnPolicyRuleSignonActions) UnmarshalJSON(bytes []byte) (err error) {
	varOktaSignOnPolicyRuleSignonActions := _OktaSignOnPolicyRuleSignonActions{}

	err = json.Unmarshal(bytes, &varOktaSignOnPolicyRuleSignonActions)
	if err == nil {
		*o = OktaSignOnPolicyRuleSignonActions(varOktaSignOnPolicyRuleSignonActions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "access")
		delete(additionalProperties, "factorLifetime")
		delete(additionalProperties, "factorPromptMode")
		delete(additionalProperties, "rememberDeviceByDefault")
		delete(additionalProperties, "requireFactor")
		delete(additionalProperties, "session")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOktaSignOnPolicyRuleSignonActions struct {
	value *OktaSignOnPolicyRuleSignonActions
	isSet bool
}

func (v NullableOktaSignOnPolicyRuleSignonActions) Get() *OktaSignOnPolicyRuleSignonActions {
	return v.value
}

func (v *NullableOktaSignOnPolicyRuleSignonActions) Set(val *OktaSignOnPolicyRuleSignonActions) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaSignOnPolicyRuleSignonActions) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaSignOnPolicyRuleSignonActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaSignOnPolicyRuleSignonActions(val *OktaSignOnPolicyRuleSignonActions) *NullableOktaSignOnPolicyRuleSignonActions {
	return &NullableOktaSignOnPolicyRuleSignonActions{value: val, isSet: true}
}

func (v NullableOktaSignOnPolicyRuleSignonActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaSignOnPolicyRuleSignonActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

