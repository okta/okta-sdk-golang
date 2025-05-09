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

// AppAndInstanceConditionEvaluatorAppOrInstance struct for AppAndInstanceConditionEvaluatorAppOrInstance
type AppAndInstanceConditionEvaluatorAppOrInstance struct {
	// ID of the app
	Id *string `json:"id,omitempty"`
	// Name of the app type
	Name *string `json:"name,omitempty"`
	// Type of app
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppAndInstanceConditionEvaluatorAppOrInstance AppAndInstanceConditionEvaluatorAppOrInstance

// NewAppAndInstanceConditionEvaluatorAppOrInstance instantiates a new AppAndInstanceConditionEvaluatorAppOrInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAndInstanceConditionEvaluatorAppOrInstance() *AppAndInstanceConditionEvaluatorAppOrInstance {
	this := AppAndInstanceConditionEvaluatorAppOrInstance{}
	return &this
}

// NewAppAndInstanceConditionEvaluatorAppOrInstanceWithDefaults instantiates a new AppAndInstanceConditionEvaluatorAppOrInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAndInstanceConditionEvaluatorAppOrInstanceWithDefaults() *AppAndInstanceConditionEvaluatorAppOrInstance {
	this := AppAndInstanceConditionEvaluatorAppOrInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppAndInstanceConditionEvaluatorAppOrInstance) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAndInstanceConditionEvaluatorAppOrInstance) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppAndInstanceConditionEvaluatorAppOrInstance) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppAndInstanceConditionEvaluatorAppOrInstance) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppAndInstanceConditionEvaluatorAppOrInstance) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAndInstanceConditionEvaluatorAppOrInstance) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppAndInstanceConditionEvaluatorAppOrInstance) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppAndInstanceConditionEvaluatorAppOrInstance) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AppAndInstanceConditionEvaluatorAppOrInstance) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAndInstanceConditionEvaluatorAppOrInstance) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AppAndInstanceConditionEvaluatorAppOrInstance) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AppAndInstanceConditionEvaluatorAppOrInstance) SetType(v string) {
	o.Type = &v
}

func (o AppAndInstanceConditionEvaluatorAppOrInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppAndInstanceConditionEvaluatorAppOrInstance) UnmarshalJSON(bytes []byte) (err error) {
	varAppAndInstanceConditionEvaluatorAppOrInstance := _AppAndInstanceConditionEvaluatorAppOrInstance{}

	err = json.Unmarshal(bytes, &varAppAndInstanceConditionEvaluatorAppOrInstance)
	if err == nil {
		*o = AppAndInstanceConditionEvaluatorAppOrInstance(varAppAndInstanceConditionEvaluatorAppOrInstance)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAppAndInstanceConditionEvaluatorAppOrInstance struct {
	value *AppAndInstanceConditionEvaluatorAppOrInstance
	isSet bool
}

func (v NullableAppAndInstanceConditionEvaluatorAppOrInstance) Get() *AppAndInstanceConditionEvaluatorAppOrInstance {
	return v.value
}

func (v *NullableAppAndInstanceConditionEvaluatorAppOrInstance) Set(val *AppAndInstanceConditionEvaluatorAppOrInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAndInstanceConditionEvaluatorAppOrInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAndInstanceConditionEvaluatorAppOrInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAndInstanceConditionEvaluatorAppOrInstance(val *AppAndInstanceConditionEvaluatorAppOrInstance) *NullableAppAndInstanceConditionEvaluatorAppOrInstance {
	return &NullableAppAndInstanceConditionEvaluatorAppOrInstance{value: val, isSet: true}
}

func (v NullableAppAndInstanceConditionEvaluatorAppOrInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAndInstanceConditionEvaluatorAppOrInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

