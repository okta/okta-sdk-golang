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
	"fmt"
)

// checks if the ImportScheduleSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportScheduleSettings{}

// ImportScheduleSettings struct for ImportScheduleSettings
type ImportScheduleSettings struct {
	// The import schedule in UNIX cron format
	Expression string `json:"expression"`
	// The import schedule time zone in Internet Assigned Numbers Authority (IANA) time zone name format
	Timezone             *string `json:"timezone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportScheduleSettings ImportScheduleSettings

// NewImportScheduleSettings instantiates a new ImportScheduleSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportScheduleSettings(expression string) *ImportScheduleSettings {
	this := ImportScheduleSettings{}
	this.Expression = expression
	return &this
}

// NewImportScheduleSettingsWithDefaults instantiates a new ImportScheduleSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportScheduleSettingsWithDefaults() *ImportScheduleSettings {
	this := ImportScheduleSettings{}
	return &this
}

// GetExpression returns the Expression field value
func (o *ImportScheduleSettings) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *ImportScheduleSettings) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *ImportScheduleSettings) SetExpression(v string) {
	o.Expression = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ImportScheduleSettings) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportScheduleSettings) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ImportScheduleSettings) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ImportScheduleSettings) SetTimezone(v string) {
	o.Timezone = &v
}

func (o ImportScheduleSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportScheduleSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportScheduleSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"expression",
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

	varImportScheduleSettings := _ImportScheduleSettings{}

	err = json.Unmarshal(data, &varImportScheduleSettings)

	if err != nil {
		return err
	}

	*o = ImportScheduleSettings(varImportScheduleSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expression")
		delete(additionalProperties, "timezone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportScheduleSettings struct {
	value *ImportScheduleSettings
	isSet bool
}

func (v NullableImportScheduleSettings) Get() *ImportScheduleSettings {
	return v.value
}

func (v *NullableImportScheduleSettings) Set(val *ImportScheduleSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableImportScheduleSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableImportScheduleSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportScheduleSettings(val *ImportScheduleSettings) *NullableImportScheduleSettings {
	return &NullableImportScheduleSettings{value: val, isSet: true}
}

func (v NullableImportScheduleSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportScheduleSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
