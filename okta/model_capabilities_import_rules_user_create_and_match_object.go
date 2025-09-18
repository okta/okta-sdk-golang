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
)

// checks if the CapabilitiesImportRulesUserCreateAndMatchObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilitiesImportRulesUserCreateAndMatchObject{}

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
	ExactMatchCriteria   *string `json:"exactMatchCriteria,omitempty"`
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
	if o == nil || IsNil(o.AllowPartialMatch) {
		var ret bool
		return ret
	}
	return *o.AllowPartialMatch
}

// GetAllowPartialMatchOk returns a tuple with the AllowPartialMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAllowPartialMatchOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPartialMatch) {
		return nil, false
	}
	return o.AllowPartialMatch, true
}

// HasAllowPartialMatch returns a boolean if a field has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasAllowPartialMatch() bool {
	if o != nil && !IsNil(o.AllowPartialMatch) {
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
	if o == nil || IsNil(o.AutoActivateNewUsers) {
		var ret bool
		return ret
	}
	return *o.AutoActivateNewUsers
}

// GetAutoActivateNewUsersOk returns a tuple with the AutoActivateNewUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoActivateNewUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoActivateNewUsers) {
		return nil, false
	}
	return o.AutoActivateNewUsers, true
}

// HasAutoActivateNewUsers returns a boolean if a field has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasAutoActivateNewUsers() bool {
	if o != nil && !IsNil(o.AutoActivateNewUsers) {
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
	if o == nil || IsNil(o.AutoConfirmExactMatch) {
		var ret bool
		return ret
	}
	return *o.AutoConfirmExactMatch
}

// GetAutoConfirmExactMatchOk returns a tuple with the AutoConfirmExactMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoConfirmExactMatchOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoConfirmExactMatch) {
		return nil, false
	}
	return o.AutoConfirmExactMatch, true
}

// HasAutoConfirmExactMatch returns a boolean if a field has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasAutoConfirmExactMatch() bool {
	if o != nil && !IsNil(o.AutoConfirmExactMatch) {
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
	if o == nil || IsNil(o.AutoConfirmNewUsers) {
		var ret bool
		return ret
	}
	return *o.AutoConfirmNewUsers
}

// GetAutoConfirmNewUsersOk returns a tuple with the AutoConfirmNewUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoConfirmNewUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoConfirmNewUsers) {
		return nil, false
	}
	return o.AutoConfirmNewUsers, true
}

// HasAutoConfirmNewUsers returns a boolean if a field has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasAutoConfirmNewUsers() bool {
	if o != nil && !IsNil(o.AutoConfirmNewUsers) {
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
	if o == nil || IsNil(o.AutoConfirmPartialMatch) {
		var ret bool
		return ret
	}
	return *o.AutoConfirmPartialMatch
}

// GetAutoConfirmPartialMatchOk returns a tuple with the AutoConfirmPartialMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoConfirmPartialMatchOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoConfirmPartialMatch) {
		return nil, false
	}
	return o.AutoConfirmPartialMatch, true
}

// HasAutoConfirmPartialMatch returns a boolean if a field has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasAutoConfirmPartialMatch() bool {
	if o != nil && !IsNil(o.AutoConfirmPartialMatch) {
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
	if o == nil || IsNil(o.ExactMatchCriteria) {
		var ret string
		return ret
	}
	return *o.ExactMatchCriteria
}

// GetExactMatchCriteriaOk returns a tuple with the ExactMatchCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetExactMatchCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.ExactMatchCriteria) {
		return nil, false
	}
	return o.ExactMatchCriteria, true
}

// HasExactMatchCriteria returns a boolean if a field has been set.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasExactMatchCriteria() bool {
	if o != nil && !IsNil(o.ExactMatchCriteria) {
		return true
	}

	return false
}

// SetExactMatchCriteria gets a reference to the given string and assigns it to the ExactMatchCriteria field.
func (o *CapabilitiesImportRulesUserCreateAndMatchObject) SetExactMatchCriteria(v string) {
	o.ExactMatchCriteria = &v
}

func (o CapabilitiesImportRulesUserCreateAndMatchObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilitiesImportRulesUserCreateAndMatchObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowPartialMatch) {
		toSerialize["allowPartialMatch"] = o.AllowPartialMatch
	}
	if !IsNil(o.AutoActivateNewUsers) {
		toSerialize["autoActivateNewUsers"] = o.AutoActivateNewUsers
	}
	if !IsNil(o.AutoConfirmExactMatch) {
		toSerialize["autoConfirmExactMatch"] = o.AutoConfirmExactMatch
	}
	if !IsNil(o.AutoConfirmNewUsers) {
		toSerialize["autoConfirmNewUsers"] = o.AutoConfirmNewUsers
	}
	if !IsNil(o.AutoConfirmPartialMatch) {
		toSerialize["autoConfirmPartialMatch"] = o.AutoConfirmPartialMatch
	}
	if !IsNil(o.ExactMatchCriteria) {
		toSerialize["exactMatchCriteria"] = o.ExactMatchCriteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilitiesImportRulesUserCreateAndMatchObject) UnmarshalJSON(data []byte) (err error) {
	varCapabilitiesImportRulesUserCreateAndMatchObject := _CapabilitiesImportRulesUserCreateAndMatchObject{}

	err = json.Unmarshal(data, &varCapabilitiesImportRulesUserCreateAndMatchObject)

	if err != nil {
		return err
	}

	*o = CapabilitiesImportRulesUserCreateAndMatchObject(varCapabilitiesImportRulesUserCreateAndMatchObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allowPartialMatch")
		delete(additionalProperties, "autoActivateNewUsers")
		delete(additionalProperties, "autoConfirmExactMatch")
		delete(additionalProperties, "autoConfirmNewUsers")
		delete(additionalProperties, "autoConfirmPartialMatch")
		delete(additionalProperties, "exactMatchCriteria")
		o.AdditionalProperties = additionalProperties
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
