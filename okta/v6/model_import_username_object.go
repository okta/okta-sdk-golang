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

// checks if the ImportUsernameObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportUsernameObject{}

// ImportUsernameObject Determines the Okta username for the imported user
type ImportUsernameObject struct {
	// For `usernameFormat=CUSTOM`, specifies the Okta Expression Language statement for a username format that imported users use to sign in to Okta
	UserNameExpression *string `json:"userNameExpression,omitempty"`
	// Determines the username format when users sign in to Okta
	UsernameFormat       string `json:"usernameFormat"`
	AdditionalProperties map[string]interface{}
}

type _ImportUsernameObject ImportUsernameObject

// NewImportUsernameObject instantiates a new ImportUsernameObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportUsernameObject(usernameFormat string) *ImportUsernameObject {
	this := ImportUsernameObject{}
	this.UsernameFormat = usernameFormat
	return &this
}

// NewImportUsernameObjectWithDefaults instantiates a new ImportUsernameObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportUsernameObjectWithDefaults() *ImportUsernameObject {
	this := ImportUsernameObject{}
	var usernameFormat string = "EMAIL"
	this.UsernameFormat = usernameFormat
	return &this
}

// GetUserNameExpression returns the UserNameExpression field value if set, zero value otherwise.
func (o *ImportUsernameObject) GetUserNameExpression() string {
	if o == nil || IsNil(o.UserNameExpression) {
		var ret string
		return ret
	}
	return *o.UserNameExpression
}

// GetUserNameExpressionOk returns a tuple with the UserNameExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUsernameObject) GetUserNameExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.UserNameExpression) {
		return nil, false
	}
	return o.UserNameExpression, true
}

// HasUserNameExpression returns a boolean if a field has been set.
func (o *ImportUsernameObject) HasUserNameExpression() bool {
	if o != nil && !IsNil(o.UserNameExpression) {
		return true
	}

	return false
}

// SetUserNameExpression gets a reference to the given string and assigns it to the UserNameExpression field.
func (o *ImportUsernameObject) SetUserNameExpression(v string) {
	o.UserNameExpression = &v
}

// GetUsernameFormat returns the UsernameFormat field value
func (o *ImportUsernameObject) GetUsernameFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UsernameFormat
}

// GetUsernameFormatOk returns a tuple with the UsernameFormat field value
// and a boolean to check if the value has been set.
func (o *ImportUsernameObject) GetUsernameFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsernameFormat, true
}

// SetUsernameFormat sets field value
func (o *ImportUsernameObject) SetUsernameFormat(v string) {
	o.UsernameFormat = v
}

func (o ImportUsernameObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportUsernameObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserNameExpression) {
		toSerialize["userNameExpression"] = o.UserNameExpression
	}
	toSerialize["usernameFormat"] = o.UsernameFormat

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportUsernameObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"usernameFormat",
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

	varImportUsernameObject := _ImportUsernameObject{}

	err = json.Unmarshal(data, &varImportUsernameObject)

	if err != nil {
		return err
	}

	*o = ImportUsernameObject(varImportUsernameObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "userNameExpression")
		delete(additionalProperties, "usernameFormat")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportUsernameObject struct {
	value *ImportUsernameObject
	isSet bool
}

func (v NullableImportUsernameObject) Get() *ImportUsernameObject {
	return v.value
}

func (v *NullableImportUsernameObject) Set(val *ImportUsernameObject) {
	v.value = val
	v.isSet = true
}

func (v NullableImportUsernameObject) IsSet() bool {
	return v.isSet
}

func (v *NullableImportUsernameObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportUsernameObject(val *ImportUsernameObject) *NullableImportUsernameObject {
	return &NullableImportUsernameObject{value: val, isSet: true}
}

func (v NullableImportUsernameObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportUsernameObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
