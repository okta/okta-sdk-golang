/*
Okta Admin Management API

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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the ImportRoadblockAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportRoadblockAction{}

// ImportRoadblockAction An action to apply to the active import roadblock of an app instance
type ImportRoadblockAction struct {
	// The action to apply:  * `RESTART`: Restarts the import that triggered the roadblock, bypassing the roadblock safeguard for that import. * `CANCEL`: Cancels the import that triggered the roadblock and resumes the user imports paused by that roadblock. The resumed imports run with the roadblock thresholds still enforced: no safeguard is skipped.
	Action string `json:"action"`
	// ID of the app instance whose roadblock to act on
	AppInstanceId        string `json:"appInstanceId"`
	AdditionalProperties map[string]interface{}
}

type _ImportRoadblockAction ImportRoadblockAction

// NewImportRoadblockAction instantiates a new ImportRoadblockAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportRoadblockAction(action string, appInstanceId string) *ImportRoadblockAction {
	this := ImportRoadblockAction{}
	this.Action = action
	this.AppInstanceId = appInstanceId
	return &this
}

// NewImportRoadblockActionWithDefaults instantiates a new ImportRoadblockAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportRoadblockActionWithDefaults() *ImportRoadblockAction {
	this := ImportRoadblockAction{}
	return &this
}

// GetAction returns the Action field value
func (o *ImportRoadblockAction) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ImportRoadblockAction) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ImportRoadblockAction) SetAction(v string) {
	o.Action = v
}

// GetAppInstanceId returns the AppInstanceId field value
func (o *ImportRoadblockAction) GetAppInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppInstanceId
}

// GetAppInstanceIdOk returns a tuple with the AppInstanceId field value
// and a boolean to check if the value has been set.
func (o *ImportRoadblockAction) GetAppInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppInstanceId, true
}

// SetAppInstanceId sets field value
func (o *ImportRoadblockAction) SetAppInstanceId(v string) {
	o.AppInstanceId = v
}

func (o ImportRoadblockAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportRoadblockAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["appInstanceId"] = o.AppInstanceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportRoadblockAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"appInstanceId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varImportRoadblockAction := _ImportRoadblockAction{}

	err = json.Unmarshal(data, &varImportRoadblockAction)

	if err != nil {
		return err
	}

	*o = ImportRoadblockAction(varImportRoadblockAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "appInstanceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportRoadblockAction struct {
	value *ImportRoadblockAction
	isSet bool
}

func (v NullableImportRoadblockAction) Get() *ImportRoadblockAction {
	return v.value
}

func (v *NullableImportRoadblockAction) Set(val *ImportRoadblockAction) {
	v.value = val
	v.isSet = true
}

func (v NullableImportRoadblockAction) IsSet() bool {
	return v.isSet
}

func (v *NullableImportRoadblockAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportRoadblockAction(val *ImportRoadblockAction) *NullableImportRoadblockAction {
	return &NullableImportRoadblockAction{value: val, isSet: true}
}

func (v NullableImportRoadblockAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportRoadblockAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
