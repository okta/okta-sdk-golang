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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the SamlAttributeStatementExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlAttributeStatementExpression{}

// SamlAttributeStatementExpression Generic `EXPRESSION` attribute statements
type SamlAttributeStatementExpression struct {
	// The name of the attribute in your app. The attribute name must be unique across all user and group attribute statements.
	Name *string `json:"name,omitempty"`
	// The name format of the attribute. Supported values:
	Namespace *string `json:"namespace,omitempty"`
	// The type of attribute statements object
	Type *string `json:"type,omitempty"`
	// The attribute values (supports [Okta Expression Language](https://developer.okta.com/docs/reference/okta-expression-language/))
	Values               []string `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlAttributeStatementExpression SamlAttributeStatementExpression

// NewSamlAttributeStatementExpression instantiates a new SamlAttributeStatementExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlAttributeStatementExpression() *SamlAttributeStatementExpression {
	this := SamlAttributeStatementExpression{}
	return &this
}

// NewSamlAttributeStatementExpressionWithDefaults instantiates a new SamlAttributeStatementExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlAttributeStatementExpressionWithDefaults() *SamlAttributeStatementExpression {
	this := SamlAttributeStatementExpression{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SamlAttributeStatementExpression) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAttributeStatementExpression) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SamlAttributeStatementExpression) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SamlAttributeStatementExpression) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *SamlAttributeStatementExpression) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAttributeStatementExpression) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *SamlAttributeStatementExpression) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *SamlAttributeStatementExpression) SetNamespace(v string) {
	o.Namespace = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SamlAttributeStatementExpression) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAttributeStatementExpression) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SamlAttributeStatementExpression) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SamlAttributeStatementExpression) SetType(v string) {
	o.Type = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *SamlAttributeStatementExpression) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAttributeStatementExpression) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *SamlAttributeStatementExpression) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *SamlAttributeStatementExpression) SetValues(v []string) {
	o.Values = v
}

func (o SamlAttributeStatementExpression) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlAttributeStatementExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SamlAttributeStatementExpression) UnmarshalJSON(data []byte) (err error) {
	varSamlAttributeStatementExpression := _SamlAttributeStatementExpression{}

	err = json.Unmarshal(data, &varSamlAttributeStatementExpression)

	if err != nil {
		return err
	}

	*o = SamlAttributeStatementExpression(varSamlAttributeStatementExpression)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "namespace")
		delete(additionalProperties, "type")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSamlAttributeStatementExpression struct {
	value *SamlAttributeStatementExpression
	isSet bool
}

func (v NullableSamlAttributeStatementExpression) Get() *SamlAttributeStatementExpression {
	return v.value
}

func (v *NullableSamlAttributeStatementExpression) Set(val *SamlAttributeStatementExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlAttributeStatementExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlAttributeStatementExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlAttributeStatementExpression(val *SamlAttributeStatementExpression) *NullableSamlAttributeStatementExpression {
	return &NullableSamlAttributeStatementExpression{value: val, isSet: true}
}

func (v NullableSamlAttributeStatementExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlAttributeStatementExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
