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

// ContinuousAccessPolicyRuleTerminateSessionSlo struct for ContinuousAccessPolicyRuleTerminateSessionSlo
type ContinuousAccessPolicyRuleTerminateSessionSlo struct {
	// This property defines the session to terminate - everyone, no one, or a specific app instance.
	AppSelectionMode *string `json:"appSelectionMode,omitempty"`
	// This property defines the app instance access to terminate. Only include this property when `appSelectionMode` is set to `SPECIFIC`.
	AppInstanceIds []string `json:"appInstanceIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContinuousAccessPolicyRuleTerminateSessionSlo ContinuousAccessPolicyRuleTerminateSessionSlo

// NewContinuousAccessPolicyRuleTerminateSessionSlo instantiates a new ContinuousAccessPolicyRuleTerminateSessionSlo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinuousAccessPolicyRuleTerminateSessionSlo() *ContinuousAccessPolicyRuleTerminateSessionSlo {
	this := ContinuousAccessPolicyRuleTerminateSessionSlo{}
	return &this
}

// NewContinuousAccessPolicyRuleTerminateSessionSloWithDefaults instantiates a new ContinuousAccessPolicyRuleTerminateSessionSlo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinuousAccessPolicyRuleTerminateSessionSloWithDefaults() *ContinuousAccessPolicyRuleTerminateSessionSlo {
	this := ContinuousAccessPolicyRuleTerminateSessionSlo{}
	return &this
}

// GetAppSelectionMode returns the AppSelectionMode field value if set, zero value otherwise.
func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) GetAppSelectionMode() string {
	if o == nil || o.AppSelectionMode == nil {
		var ret string
		return ret
	}
	return *o.AppSelectionMode
}

// GetAppSelectionModeOk returns a tuple with the AppSelectionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) GetAppSelectionModeOk() (*string, bool) {
	if o == nil || o.AppSelectionMode == nil {
		return nil, false
	}
	return o.AppSelectionMode, true
}

// HasAppSelectionMode returns a boolean if a field has been set.
func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) HasAppSelectionMode() bool {
	if o != nil && o.AppSelectionMode != nil {
		return true
	}

	return false
}

// SetAppSelectionMode gets a reference to the given string and assigns it to the AppSelectionMode field.
func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) SetAppSelectionMode(v string) {
	o.AppSelectionMode = &v
}

// GetAppInstanceIds returns the AppInstanceIds field value if set, zero value otherwise.
func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) GetAppInstanceIds() []string {
	if o == nil || o.AppInstanceIds == nil {
		var ret []string
		return ret
	}
	return o.AppInstanceIds
}

// GetAppInstanceIdsOk returns a tuple with the AppInstanceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) GetAppInstanceIdsOk() ([]string, bool) {
	if o == nil || o.AppInstanceIds == nil {
		return nil, false
	}
	return o.AppInstanceIds, true
}

// HasAppInstanceIds returns a boolean if a field has been set.
func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) HasAppInstanceIds() bool {
	if o != nil && o.AppInstanceIds != nil {
		return true
	}

	return false
}

// SetAppInstanceIds gets a reference to the given []string and assigns it to the AppInstanceIds field.
func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) SetAppInstanceIds(v []string) {
	o.AppInstanceIds = v
}

func (o ContinuousAccessPolicyRuleTerminateSessionSlo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppSelectionMode != nil {
		toSerialize["appSelectionMode"] = o.AppSelectionMode
	}
	if o.AppInstanceIds != nil {
		toSerialize["appInstanceIds"] = o.AppInstanceIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) UnmarshalJSON(bytes []byte) (err error) {
	varContinuousAccessPolicyRuleTerminateSessionSlo := _ContinuousAccessPolicyRuleTerminateSessionSlo{}

	err = json.Unmarshal(bytes, &varContinuousAccessPolicyRuleTerminateSessionSlo)
	if err == nil {
		*o = ContinuousAccessPolicyRuleTerminateSessionSlo(varContinuousAccessPolicyRuleTerminateSessionSlo)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "appSelectionMode")
		delete(additionalProperties, "appInstanceIds")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableContinuousAccessPolicyRuleTerminateSessionSlo struct {
	value *ContinuousAccessPolicyRuleTerminateSessionSlo
	isSet bool
}

func (v NullableContinuousAccessPolicyRuleTerminateSessionSlo) Get() *ContinuousAccessPolicyRuleTerminateSessionSlo {
	return v.value
}

func (v *NullableContinuousAccessPolicyRuleTerminateSessionSlo) Set(val *ContinuousAccessPolicyRuleTerminateSessionSlo) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousAccessPolicyRuleTerminateSessionSlo) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousAccessPolicyRuleTerminateSessionSlo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousAccessPolicyRuleTerminateSessionSlo(val *ContinuousAccessPolicyRuleTerminateSessionSlo) *NullableContinuousAccessPolicyRuleTerminateSessionSlo {
	return &NullableContinuousAccessPolicyRuleTerminateSessionSlo{value: val, isSet: true}
}

func (v NullableContinuousAccessPolicyRuleTerminateSessionSlo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousAccessPolicyRuleTerminateSessionSlo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

