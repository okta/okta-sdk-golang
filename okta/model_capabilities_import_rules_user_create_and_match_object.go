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

// CapabilitiesImportRulesUserCreateAndMatchObject Rules for matching and creating users
type CapabilitiesImportRulesUserCreateAndMatchObject struct {
	// Allows user import upon partial matching. Partial matching occurs when the first and last names of an imported user match those of an existing Okta user, even if the username or email attributes don't match.
	AllowPartialMatch *bool `json:"allowPartialMatch,omitempty"`
	// If set to `true`, imported new users are automatically activated.
	AutoActivateNewUsers *bool `json:"autoActivateNewUsers,omitempty"`
	// If set to `true`, exact-matched users are automatically confirmed on activation. If set to `false`, exact-matched users need to be confirmed manually.
	AutoConfirmExactMatch *bool `json:"autoConfirmExactMatch,omitempty"`
	// If set to `true`, imported new users are automatically confirmed on activation. This doesn't apply to imported users that already exist in Okta.
	AutoConfirmNewUsers *bool `json:"autoConfirmNewUsers,omitempty"`
	// If set to `true`, partially matched users are automatically confirmed on activation. If set to `false`, partially matched users need to be confirmed manually.
	AutoConfirmPartialMatch *bool `json:"autoConfirmPartialMatch,omitempty"`
	// Determines the attribute to match users
	ExactMatchCriteria *string `json:"exactMatchCriteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitiesImportRulesUserCreateAndMatchObject CapabilitiesImportRulesUserCreateAndMatchObject

// NewCapabilitiesImportRulesUserCreateAndMatchObject instantiates a new CapabilitiesImportRulesUserCreateAndMatchObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitiesImportRulesUserCreateAndMatchObject() *CapabilitiesImportRulesUserCreateAndMatchObject {
	this := CapabilitiesImportRulesUserCreateAndMatchObject{}
	return &this
}

// NewCapabilitiesImportRulesUserCreateAndMatchObjectWithDefaults instantiates a new CapabilitiesImportRulesUserCreateAndMatchObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitiesImportRulesUserCreateAndMatchObjectWithDefaults() *CapabilitiesImportRulesUserCreateAndMatchObject {
	this := CapabilitiesImportRulesUserCreateAndMatchObject{}
	return &this
}

// GetAllowPartialMatch returns the AllowPartialMatch field value if set, zero value otherwise.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAllowPartialMatch() bool {
	if o == nil || o.AllowPartialMatch == nil {
		var ret bool
		return ret
	}
	return *o.AllowPartialMatch
}

// GetAllowPartialMatchOk returns a tuple with the AllowPartialMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAllowPartialMatchOk() (*bool, bool) {
	if o == nil || o.AllowPartialMatch == nil {
		return nil, false
	}
	return o.AllowPartialMatch, true
}

// HasAllowPartialMatch returns a boolean if a field has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasAllowPartialMatch() bool {
	if o != nil && o.AllowPartialMatch != nil {
		return true
	}

	return false
}

// SetAllowPartialMatch gets a reference to the given bool and assigns it to the AllowPartialMatch field.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) SetAllowPartialMatch(v bool) {
	o.AllowPartialMatch = &v
}

// GetAutoActivateNewUsers returns the AutoActivateNewUsers field value if set, zero value otherwise.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoActivateNewUsers() bool {
	if o == nil || o.AutoActivateNewUsers == nil {
		var ret bool
		return ret
	}
	return *o.AutoActivateNewUsers
}

// GetAutoActivateNewUsersOk returns a tuple with the AutoActivateNewUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoActivateNewUsersOk() (*bool, bool) {
	if o == nil || o.AutoActivateNewUsers == nil {
		return nil, false
	}
	return o.AutoActivateNewUsers, true
}

// HasAutoActivateNewUsers returns a boolean if a field has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasAutoActivateNewUsers() bool {
	if o != nil && o.AutoActivateNewUsers != nil {
		return true
	}

	return false
}

// SetAutoActivateNewUsers gets a reference to the given bool and assigns it to the AutoActivateNewUsers field.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) SetAutoActivateNewUsers(v bool) {
	o.AutoActivateNewUsers = &v
}

// GetAutoConfirmExactMatch returns the AutoConfirmExactMatch field value if set, zero value otherwise.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoConfirmExactMatch() bool {
	if o == nil || o.AutoConfirmExactMatch == nil {
		var ret bool
		return ret
	}
	return *o.AutoConfirmExactMatch
}

// GetAutoConfirmExactMatchOk returns a tuple with the AutoConfirmExactMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoConfirmExactMatchOk() (*bool, bool) {
	if o == nil || o.AutoConfirmExactMatch == nil {
		return nil, false
	}
	return o.AutoConfirmExactMatch, true
}

// HasAutoConfirmExactMatch returns a boolean if a field has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasAutoConfirmExactMatch() bool {
	if o != nil && o.AutoConfirmExactMatch != nil {
		return true
	}

	return false
}

// SetAutoConfirmExactMatch gets a reference to the given bool and assigns it to the AutoConfirmExactMatch field.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) SetAutoConfirmExactMatch(v bool) {
	o.AutoConfirmExactMatch = &v
}

// GetAutoConfirmNewUsers returns the AutoConfirmNewUsers field value if set, zero value otherwise.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoConfirmNewUsers() bool {
	if o == nil || o.AutoConfirmNewUsers == nil {
		var ret bool
		return ret
	}
	return *o.AutoConfirmNewUsers
}

// GetAutoConfirmNewUsersOk returns a tuple with the AutoConfirmNewUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoConfirmNewUsersOk() (*bool, bool) {
	if o == nil || o.AutoConfirmNewUsers == nil {
		return nil, false
	}
	return o.AutoConfirmNewUsers, true
}

// HasAutoConfirmNewUsers returns a boolean if a field has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasAutoConfirmNewUsers() bool {
	if o != nil && o.AutoConfirmNewUsers != nil {
		return true
	}

	return false
}

// SetAutoConfirmNewUsers gets a reference to the given bool and assigns it to the AutoConfirmNewUsers field.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) SetAutoConfirmNewUsers(v bool) {
	o.AutoConfirmNewUsers = &v
}

// GetAutoConfirmPartialMatch returns the AutoConfirmPartialMatch field value if set, zero value otherwise.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoConfirmPartialMatch() bool {
	if o == nil || o.AutoConfirmPartialMatch == nil {
		var ret bool
		return ret
	}
	return *o.AutoConfirmPartialMatch
}

// GetAutoConfirmPartialMatchOk returns a tuple with the AutoConfirmPartialMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoConfirmPartialMatchOk() (*bool, bool) {
	if o == nil || o.AutoConfirmPartialMatch == nil {
		return nil, false
	}
	return o.AutoConfirmPartialMatch, true
}

// HasAutoConfirmPartialMatch returns a boolean if a field has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasAutoConfirmPartialMatch() bool {
	if o != nil && o.AutoConfirmPartialMatch != nil {
		return true
	}

	return false
}

// SetAutoConfirmPartialMatch gets a reference to the given bool and assigns it to the AutoConfirmPartialMatch field.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) SetAutoConfirmPartialMatch(v bool) {
	o.AutoConfirmPartialMatch = &v
}

// GetExactMatchCriteria returns the ExactMatchCriteria field value if set, zero value otherwise.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetExactMatchCriteria() string {
	if o == nil || o.ExactMatchCriteria == nil {
		var ret string
		return ret
	}
	return *o.ExactMatchCriteria
}

// GetExactMatchCriteriaOk returns a tuple with the ExactMatchCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetExactMatchCriteriaOk() (*string, bool) {
	if o == nil || o.ExactMatchCriteria == nil {
		return nil, false
	}
	return o.ExactMatchCriteria, true
}

// HasExactMatchCriteria returns a boolean if a field has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasExactMatchCriteria() bool {
	if o != nil && o.ExactMatchCriteria != nil {
		return true
	}

	return false
}

// SetExactMatchCriteria gets a reference to the given string and assigns it to the ExactMatchCriteria field.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) SetExactMatchCriteria(v string) {
	o.ExactMatchCriteria = &v
}

func (o CapabilitiesImportRulesUserCreateAndMatchObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowPartialMatch != nil {
		toSerialize["allowPartialMatch"] = o.AllowPartialMatch
	}
	if o.AutoActivateNewUsers != nil {
		toSerialize["autoActivateNewUsers"] = o.AutoActivateNewUsers
	}
	if o.AutoConfirmExactMatch != nil {
		toSerialize["autoConfirmExactMatch"] = o.AutoConfirmExactMatch
	}
	if o.AutoConfirmNewUsers != nil {
		toSerialize["autoConfirmNewUsers"] = o.AutoConfirmNewUsers
	}
	if o.AutoConfirmPartialMatch != nil {
		toSerialize["autoConfirmPartialMatch"] = o.AutoConfirmPartialMatch
	}
	if o.ExactMatchCriteria != nil {
		toSerialize["exactMatchCriteria"] = o.ExactMatchCriteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitiesImportRulesUserCreateAndMatchObject) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilitiesImportRulesUserCreateAndMatchObject := _CapabilitiesImportRulesUserCreateAndMatchObject{}

	err = json.Unmarshal(bytes, &varCapabilitiesImportRulesUserCreateAndMatchObject)
	if err == nil {
		*o = CapabilitiesImportRulesUserCreateAndMatchObject(varCapabilitiesImportRulesUserCreateAndMatchObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "allowPartialMatch")
		delete(additionalProperties, "autoActivateNewUsers")
		delete(additionalProperties, "autoConfirmExactMatch")
		delete(additionalProperties, "autoConfirmNewUsers")
		delete(additionalProperties, "autoConfirmPartialMatch")
		delete(additionalProperties, "exactMatchCriteria")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCapabilitiesImportRulesUserCreateAndMatchObject struct {
	value *CapabilitiesImportRulesUserCreateAndMatchObject
	isSet bool
}

func (v NullableCapabilitiesImportRulesUserCreateAndMatchObject) Get() *CapabilitiesImportRulesUserCreateAndMatchObject {
	return v.value
}

func (v *NullableCapabilitiesImportRulesUserCreateAndMatchObject) Set(val *CapabilitiesImportRulesUserCreateAndMatchObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitiesImportRulesUserCreateAndMatchObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitiesImportRulesUserCreateAndMatchObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitiesImportRulesUserCreateAndMatchObject(val *CapabilitiesImportRulesUserCreateAndMatchObject) *NullableCapabilitiesImportRulesUserCreateAndMatchObject {
	return &NullableCapabilitiesImportRulesUserCreateAndMatchObject{value: val, isSet: true}
}

func (v NullableCapabilitiesImportRulesUserCreateAndMatchObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitiesImportRulesUserCreateAndMatchObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

