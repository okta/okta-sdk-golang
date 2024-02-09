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

API version: 5.1.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
)

// ImportScheduleObjectIncrementalImport struct for ImportScheduleObjectIncrementalImport
type ImportScheduleObjectIncrementalImport struct {
	// The import schedule in UNIX cron format
	Expression string `json:"expression"`
	// The import schedule time zone in Internet Assigned Numbers Authority (IANA) time zone name format
	Timezone             *string `json:"timezone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportScheduleObjectIncrementalImport ImportScheduleObjectIncrementalImport

// NewImportScheduleObjectIncrementalImport instantiates a new ImportScheduleObjectIncrementalImport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportScheduleObjectIncrementalImport(expression string) *ImportScheduleObjectIncrementalImport {
	this := ImportScheduleObjectIncrementalImport{}
	this.Expression = expression
	return &this
}

// NewImportScheduleObjectIncrementalImportWithDefaults instantiates a new ImportScheduleObjectIncrementalImport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportScheduleObjectIncrementalImportWithDefaults() *ImportScheduleObjectIncrementalImport {
	this := ImportScheduleObjectIncrementalImport{}
	return &this
}

// GetExpression returns the Expression field value
func (o *ImportScheduleObjectIncrementalImport) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *ImportScheduleObjectIncrementalImport) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *ImportScheduleObjectIncrementalImport) SetExpression(v string) {
	o.Expression = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ImportScheduleObjectIncrementalImport) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportScheduleObjectIncrementalImport) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ImportScheduleObjectIncrementalImport) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ImportScheduleObjectIncrementalImport) SetTimezone(v string) {
	o.Timezone = &v
}

func (o ImportScheduleObjectIncrementalImport) MarshalJSON() ([]byte, error) {
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

func (o *ImportScheduleObjectIncrementalImport) UnmarshalJSON(bytes []byte) (err error) {
	varImportScheduleObjectIncrementalImport := _ImportScheduleObjectIncrementalImport{}

	err = json.Unmarshal(bytes, &varImportScheduleObjectIncrementalImport)
	if err == nil {
		*o = ImportScheduleObjectIncrementalImport(varImportScheduleObjectIncrementalImport)
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

type NullableImportScheduleObjectIncrementalImport struct {
	value *ImportScheduleObjectIncrementalImport
	isSet bool
}

func (v NullableImportScheduleObjectIncrementalImport) Get() *ImportScheduleObjectIncrementalImport {
	return v.value
}

func (v *NullableImportScheduleObjectIncrementalImport) Set(val *ImportScheduleObjectIncrementalImport) {
	v.value = val
	v.isSet = true
}

func (v NullableImportScheduleObjectIncrementalImport) IsSet() bool {
	return v.isSet
}

func (v *NullableImportScheduleObjectIncrementalImport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportScheduleObjectIncrementalImport(val *ImportScheduleObjectIncrementalImport) *NullableImportScheduleObjectIncrementalImport {
	return &NullableImportScheduleObjectIncrementalImport{value: val, isSet: true}
}

func (v NullableImportScheduleObjectIncrementalImport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportScheduleObjectIncrementalImport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
