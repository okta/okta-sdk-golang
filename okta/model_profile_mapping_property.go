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

// ProfileMappingProperty A target property, in string form, that maps to a valid [JSON Schema Draft](https://tools.ietf.org/html/draft-zyp-json-schema-04) document.
type ProfileMappingProperty struct {
	// Combination or single source properties that are mapped to the target property
	Expression *string `json:"expression,omitempty"`
	// Indicates whether to update target properties for user create and update or just for user create.  Having a pushStatus of `PUSH` causes properties in the target to be updated on create and update. Having a pushStatus of `DONT_PUSH` causes properties in the target to be updated only on create.
	PushStatus *string `json:"pushStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProfileMappingProperty ProfileMappingProperty

// NewProfileMappingProperty instantiates a new ProfileMappingProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileMappingProperty() *ProfileMappingProperty {
	this := ProfileMappingProperty{}
	return &this
}

// NewProfileMappingPropertyWithDefaults instantiates a new ProfileMappingProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileMappingPropertyWithDefaults() *ProfileMappingProperty {
	this := ProfileMappingProperty{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *ProfileMappingProperty) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMappingProperty) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *ProfileMappingProperty) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *ProfileMappingProperty) SetExpression(v string) {
	o.Expression = &v
}

// GetPushStatus returns the PushStatus field value if set, zero value otherwise.
func (o *ProfileMappingProperty) GetPushStatus() string {
	if o == nil || o.PushStatus == nil {
		var ret string
		return ret
	}
	return *o.PushStatus
}

// GetPushStatusOk returns a tuple with the PushStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMappingProperty) GetPushStatusOk() (*string, bool) {
	if o == nil || o.PushStatus == nil {
		return nil, false
	}
	return o.PushStatus, true
}

// HasPushStatus returns a boolean if a field has been set.
func (o *ProfileMappingProperty) HasPushStatus() bool {
	if o != nil && o.PushStatus != nil {
		return true
	}

	return false
}

// SetPushStatus gets a reference to the given string and assigns it to the PushStatus field.
func (o *ProfileMappingProperty) SetPushStatus(v string) {
	o.PushStatus = &v
}

func (o ProfileMappingProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.PushStatus != nil {
		toSerialize["pushStatus"] = o.PushStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProfileMappingProperty) UnmarshalJSON(bytes []byte) (err error) {
	varProfileMappingProperty := _ProfileMappingProperty{}

	err = json.Unmarshal(bytes, &varProfileMappingProperty)
	if err == nil {
		*o = ProfileMappingProperty(varProfileMappingProperty)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "expression")
		delete(additionalProperties, "pushStatus")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProfileMappingProperty struct {
	value *ProfileMappingProperty
	isSet bool
}

func (v NullableProfileMappingProperty) Get() *ProfileMappingProperty {
	return v.value
}

func (v *NullableProfileMappingProperty) Set(val *ProfileMappingProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileMappingProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileMappingProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileMappingProperty(val *ProfileMappingProperty) *NullableProfileMappingProperty {
	return &NullableProfileMappingProperty{value: val, isSet: true}
}

func (v NullableProfileMappingProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileMappingProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

