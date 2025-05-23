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

// ContextPolicyRuleCondition struct for ContextPolicyRuleCondition
type ContextPolicyRuleCondition struct {
	Migrated *bool `json:"migrated,omitempty"`
	Platform *DevicePolicyRuleConditionPlatform `json:"platform,omitempty"`
	Rooted *bool `json:"rooted,omitempty"`
	TrustLevel *string `json:"trustLevel,omitempty"`
	Expression *string `json:"expression,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContextPolicyRuleCondition ContextPolicyRuleCondition

// NewContextPolicyRuleCondition instantiates a new ContextPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextPolicyRuleCondition() *ContextPolicyRuleCondition {
	this := ContextPolicyRuleCondition{}
	return &this
}

// NewContextPolicyRuleConditionWithDefaults instantiates a new ContextPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextPolicyRuleConditionWithDefaults() *ContextPolicyRuleCondition {
	this := ContextPolicyRuleCondition{}
	return &this
}

// GetMigrated returns the Migrated field value if set, zero value otherwise.
func (o *ContextPolicyRuleCondition) GetMigrated() bool {
	if o == nil || o.Migrated == nil {
		var ret bool
		return ret
	}
	return *o.Migrated
}

// GetMigratedOk returns a tuple with the Migrated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextPolicyRuleCondition) GetMigratedOk() (*bool, bool) {
	if o == nil || o.Migrated == nil {
		return nil, false
	}
	return o.Migrated, true
}

// HasMigrated returns a boolean if a field has been set.
func (o *ContextPolicyRuleCondition) HasMigrated() bool {
	if o != nil && o.Migrated != nil {
		return true
	}

	return false
}

// SetMigrated gets a reference to the given bool and assigns it to the Migrated field.
func (o *ContextPolicyRuleCondition) SetMigrated(v bool) {
	o.Migrated = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *ContextPolicyRuleCondition) GetPlatform() DevicePolicyRuleConditionPlatform {
	if o == nil || o.Platform == nil {
		var ret DevicePolicyRuleConditionPlatform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextPolicyRuleCondition) GetPlatformOk() (*DevicePolicyRuleConditionPlatform, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *ContextPolicyRuleCondition) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given DevicePolicyRuleConditionPlatform and assigns it to the Platform field.
func (o *ContextPolicyRuleCondition) SetPlatform(v DevicePolicyRuleConditionPlatform) {
	o.Platform = &v
}

// GetRooted returns the Rooted field value if set, zero value otherwise.
func (o *ContextPolicyRuleCondition) GetRooted() bool {
	if o == nil || o.Rooted == nil {
		var ret bool
		return ret
	}
	return *o.Rooted
}

// GetRootedOk returns a tuple with the Rooted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextPolicyRuleCondition) GetRootedOk() (*bool, bool) {
	if o == nil || o.Rooted == nil {
		return nil, false
	}
	return o.Rooted, true
}

// HasRooted returns a boolean if a field has been set.
func (o *ContextPolicyRuleCondition) HasRooted() bool {
	if o != nil && o.Rooted != nil {
		return true
	}

	return false
}

// SetRooted gets a reference to the given bool and assigns it to the Rooted field.
func (o *ContextPolicyRuleCondition) SetRooted(v bool) {
	o.Rooted = &v
}

// GetTrustLevel returns the TrustLevel field value if set, zero value otherwise.
func (o *ContextPolicyRuleCondition) GetTrustLevel() string {
	if o == nil || o.TrustLevel == nil {
		var ret string
		return ret
	}
	return *o.TrustLevel
}

// GetTrustLevelOk returns a tuple with the TrustLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextPolicyRuleCondition) GetTrustLevelOk() (*string, bool) {
	if o == nil || o.TrustLevel == nil {
		return nil, false
	}
	return o.TrustLevel, true
}

// HasTrustLevel returns a boolean if a field has been set.
func (o *ContextPolicyRuleCondition) HasTrustLevel() bool {
	if o != nil && o.TrustLevel != nil {
		return true
	}

	return false
}

// SetTrustLevel gets a reference to the given string and assigns it to the TrustLevel field.
func (o *ContextPolicyRuleCondition) SetTrustLevel(v string) {
	o.TrustLevel = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *ContextPolicyRuleCondition) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextPolicyRuleCondition) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *ContextPolicyRuleCondition) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *ContextPolicyRuleCondition) SetExpression(v string) {
	o.Expression = &v
}

func (o ContextPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Migrated != nil {
		toSerialize["migrated"] = o.Migrated
	}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if o.Rooted != nil {
		toSerialize["rooted"] = o.Rooted
	}
	if o.TrustLevel != nil {
		toSerialize["trustLevel"] = o.TrustLevel
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ContextPolicyRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varContextPolicyRuleCondition := _ContextPolicyRuleCondition{}

	err = json.Unmarshal(bytes, &varContextPolicyRuleCondition)
	if err == nil {
		*o = ContextPolicyRuleCondition(varContextPolicyRuleCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "migrated")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "rooted")
		delete(additionalProperties, "trustLevel")
		delete(additionalProperties, "expression")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableContextPolicyRuleCondition struct {
	value *ContextPolicyRuleCondition
	isSet bool
}

func (v NullableContextPolicyRuleCondition) Get() *ContextPolicyRuleCondition {
	return v.value
}

func (v *NullableContextPolicyRuleCondition) Set(val *ContextPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableContextPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableContextPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextPolicyRuleCondition(val *ContextPolicyRuleCondition) *NullableContextPolicyRuleCondition {
	return &NullableContextPolicyRuleCondition{value: val, isSet: true}
}

func (v NullableContextPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

