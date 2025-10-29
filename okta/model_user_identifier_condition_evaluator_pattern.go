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
	"fmt"
)

// checks if the UserIdentifierConditionEvaluatorPattern type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserIdentifierConditionEvaluatorPattern{}

// UserIdentifierConditionEvaluatorPattern Specifies the details of the patterns to match against
type UserIdentifierConditionEvaluatorPattern struct {
	// The type of pattern. For regex, use `EXPRESSION`.
	MatchType string `json:"matchType"`
	// The regular expression or simple match string
	Value                string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _UserIdentifierConditionEvaluatorPattern UserIdentifierConditionEvaluatorPattern

// NewUserIdentifierConditionEvaluatorPattern instantiates a new UserIdentifierConditionEvaluatorPattern object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentifierConditionEvaluatorPattern(matchType string, value string) *UserIdentifierConditionEvaluatorPattern {
	this := UserIdentifierConditionEvaluatorPattern{}
	this.MatchType = matchType
	this.Value = value
	return &this
}

// NewUserIdentifierConditionEvaluatorPatternWithDefaults instantiates a new UserIdentifierConditionEvaluatorPattern object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentifierConditionEvaluatorPatternWithDefaults() *UserIdentifierConditionEvaluatorPattern {
	this := UserIdentifierConditionEvaluatorPattern{}
	return &this
}

// GetMatchType returns the MatchType field value
func (o *UserIdentifierConditionEvaluatorPattern) GetMatchType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *UserIdentifierConditionEvaluatorPattern) GetMatchTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *UserIdentifierConditionEvaluatorPattern) SetMatchType(v string) {
	o.MatchType = v
}

// GetValue returns the Value field value
func (o *UserIdentifierConditionEvaluatorPattern) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UserIdentifierConditionEvaluatorPattern) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UserIdentifierConditionEvaluatorPattern) SetValue(v string) {
	o.Value = v
}

func (o UserIdentifierConditionEvaluatorPattern) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserIdentifierConditionEvaluatorPattern) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["matchType"] = o.MatchType
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserIdentifierConditionEvaluatorPattern) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"matchType",
		"value",
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

	varUserIdentifierConditionEvaluatorPattern := _UserIdentifierConditionEvaluatorPattern{}

	err = json.Unmarshal(data, &varUserIdentifierConditionEvaluatorPattern)

	if err != nil {
		return err
	}

	*o = UserIdentifierConditionEvaluatorPattern(varUserIdentifierConditionEvaluatorPattern)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "matchType")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserIdentifierConditionEvaluatorPattern struct {
	value *UserIdentifierConditionEvaluatorPattern
	isSet bool
}

func (v NullableUserIdentifierConditionEvaluatorPattern) Get() *UserIdentifierConditionEvaluatorPattern {
	return v.value
}

func (v *NullableUserIdentifierConditionEvaluatorPattern) Set(val *UserIdentifierConditionEvaluatorPattern) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentifierConditionEvaluatorPattern) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentifierConditionEvaluatorPattern) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentifierConditionEvaluatorPattern(val *UserIdentifierConditionEvaluatorPattern) *NullableUserIdentifierConditionEvaluatorPattern {
	return &NullableUserIdentifierConditionEvaluatorPattern{value: val, isSet: true}
}

func (v NullableUserIdentifierConditionEvaluatorPattern) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentifierConditionEvaluatorPattern) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
