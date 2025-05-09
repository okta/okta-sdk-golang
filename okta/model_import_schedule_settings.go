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

// ImportScheduleSettings struct for ImportScheduleSettings
type ImportScheduleSettings struct {
	// The import schedule in UNIX cron format
	Expression string `json:"expression"`
	// The import schedule time zone in Internet Assigned Numbers Authority (IANA) time zone name format
	Timezone *string `json:"timezone,omitempty"`
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
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportScheduleSettings) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ImportScheduleSettings) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ImportScheduleSettings) SetTimezone(v string) {
	o.Timezone = &v
}

func (o ImportScheduleSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["expression"] = o.Expression
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ImportScheduleSettings) UnmarshalJSON(bytes []byte) (err error) {
	varImportScheduleSettings := _ImportScheduleSettings{}

	err = json.Unmarshal(bytes, &varImportScheduleSettings)
	if err == nil {
		*o = ImportScheduleSettings(varImportScheduleSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "expression")
		delete(additionalProperties, "timezone")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

