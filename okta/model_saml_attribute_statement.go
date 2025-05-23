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

// SamlAttributeStatement Define custom attribute statements for the integration. These statements are inserted into the SAML assertions shared with your app
type SamlAttributeStatement struct {
	FilterType *string `json:"filterType,omitempty"`
	FilterValue *string `json:"filterValue,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Type *string `json:"type,omitempty"`
	Values []string `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlAttributeStatement SamlAttributeStatement

// NewSamlAttributeStatement instantiates a new SamlAttributeStatement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlAttributeStatement() *SamlAttributeStatement {
	this := SamlAttributeStatement{}
	return &this
}

// NewSamlAttributeStatementWithDefaults instantiates a new SamlAttributeStatement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlAttributeStatementWithDefaults() *SamlAttributeStatement {
	this := SamlAttributeStatement{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *SamlAttributeStatement) GetFilterType() string {
	if o == nil || o.FilterType == nil {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAttributeStatement) GetFilterTypeOk() (*string, bool) {
	if o == nil || o.FilterType == nil {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *SamlAttributeStatement) HasFilterType() bool {
	if o != nil && o.FilterType != nil {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *SamlAttributeStatement) SetFilterType(v string) {
	o.FilterType = &v
}

// GetFilterValue returns the FilterValue field value if set, zero value otherwise.
func (o *SamlAttributeStatement) GetFilterValue() string {
	if o == nil || o.FilterValue == nil {
		var ret string
		return ret
	}
	return *o.FilterValue
}

// GetFilterValueOk returns a tuple with the FilterValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAttributeStatement) GetFilterValueOk() (*string, bool) {
	if o == nil || o.FilterValue == nil {
		return nil, false
	}
	return o.FilterValue, true
}

// HasFilterValue returns a boolean if a field has been set.
func (o *SamlAttributeStatement) HasFilterValue() bool {
	if o != nil && o.FilterValue != nil {
		return true
	}

	return false
}

// SetFilterValue gets a reference to the given string and assigns it to the FilterValue field.
func (o *SamlAttributeStatement) SetFilterValue(v string) {
	o.FilterValue = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SamlAttributeStatement) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAttributeStatement) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SamlAttributeStatement) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SamlAttributeStatement) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *SamlAttributeStatement) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAttributeStatement) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *SamlAttributeStatement) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *SamlAttributeStatement) SetNamespace(v string) {
	o.Namespace = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SamlAttributeStatement) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAttributeStatement) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SamlAttributeStatement) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SamlAttributeStatement) SetType(v string) {
	o.Type = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *SamlAttributeStatement) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAttributeStatement) GetValuesOk() ([]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *SamlAttributeStatement) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *SamlAttributeStatement) SetValues(v []string) {
	o.Values = v
}

func (o SamlAttributeStatement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FilterType != nil {
		toSerialize["filterType"] = o.FilterType
	}
	if o.FilterValue != nil {
		toSerialize["filterValue"] = o.FilterValue
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SamlAttributeStatement) UnmarshalJSON(bytes []byte) (err error) {
	varSamlAttributeStatement := _SamlAttributeStatement{}

	err = json.Unmarshal(bytes, &varSamlAttributeStatement)
	if err == nil {
		*o = SamlAttributeStatement(varSamlAttributeStatement)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "filterValue")
		delete(additionalProperties, "name")
		delete(additionalProperties, "namespace")
		delete(additionalProperties, "type")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSamlAttributeStatement struct {
	value *SamlAttributeStatement
	isSet bool
}

func (v NullableSamlAttributeStatement) Get() *SamlAttributeStatement {
	return v.value
}

func (v *NullableSamlAttributeStatement) Set(val *SamlAttributeStatement) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlAttributeStatement) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlAttributeStatement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlAttributeStatement(val *SamlAttributeStatement) *NullableSamlAttributeStatement {
	return &NullableSamlAttributeStatement{value: val, isSet: true}
}

func (v NullableSamlAttributeStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlAttributeStatement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

