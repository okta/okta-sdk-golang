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
)

// checks if the SamlAttributeStatementGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlAttributeStatementGroup{}

// SamlAttributeStatementGroup `GROUP` attribute statements
type SamlAttributeStatementGroup struct {
	// The operation to filter groups based on `filterValue`
	FilterType *string `json:"filterType,omitempty"`
	// Filter the groups based on a specific value.
	FilterValue *string `json:"filterValue,omitempty"`
	// The name of the group attribute in your app. The attribute name must be unique across all user and group attribute statements.
	Name *string `json:"name,omitempty"`
	// The name format of the group attribute. Supported values:
	Namespace *string `json:"namespace,omitempty"`
	// The type of attribute statements object
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlAttributeStatementGroup SamlAttributeStatementGroup

// NewSamlAttributeStatementGroup instantiates a new SamlAttributeStatementGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlAttributeStatementGroup() *SamlAttributeStatementGroup {
	this := SamlAttributeStatementGroup{}
	return &this
}

// NewSamlAttributeStatementGroupWithDefaults instantiates a new SamlAttributeStatementGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlAttributeStatementGroupWithDefaults() *SamlAttributeStatementGroup {
	this := SamlAttributeStatementGroup{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *SamlAttributeStatementGroup) GetFilterType() string {
	if o == nil || IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAttributeStatementGroup) GetFilterTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *SamlAttributeStatementGroup) HasFilterType() bool {
	if o != nil && !IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *SamlAttributeStatementGroup) SetFilterType(v string) {
	o.FilterType = &v
}

// GetFilterValue returns the FilterValue field value if set, zero value otherwise.
func (o *SamlAttributeStatementGroup) GetFilterValue() string {
	if o == nil || IsNil(o.FilterValue) {
		var ret string
		return ret
	}
	return *o.FilterValue
}

// GetFilterValueOk returns a tuple with the FilterValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAttributeStatementGroup) GetFilterValueOk() (*string, bool) {
	if o == nil || IsNil(o.FilterValue) {
		return nil, false
	}
	return o.FilterValue, true
}

// HasFilterValue returns a boolean if a field has been set.
func (o *SamlAttributeStatementGroup) HasFilterValue() bool {
	if o != nil && !IsNil(o.FilterValue) {
		return true
	}

	return false
}

// SetFilterValue gets a reference to the given string and assigns it to the FilterValue field.
func (o *SamlAttributeStatementGroup) SetFilterValue(v string) {
	o.FilterValue = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SamlAttributeStatementGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAttributeStatementGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SamlAttributeStatementGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SamlAttributeStatementGroup) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *SamlAttributeStatementGroup) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAttributeStatementGroup) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *SamlAttributeStatementGroup) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *SamlAttributeStatementGroup) SetNamespace(v string) {
	o.Namespace = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SamlAttributeStatementGroup) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAttributeStatementGroup) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SamlAttributeStatementGroup) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SamlAttributeStatementGroup) SetType(v string) {
	o.Type = &v
}

func (o SamlAttributeStatementGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlAttributeStatementGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !IsNil(o.FilterValue) {
		toSerialize["filterValue"] = o.FilterValue
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SamlAttributeStatementGroup) UnmarshalJSON(data []byte) (err error) {
	varSamlAttributeStatementGroup := _SamlAttributeStatementGroup{}

	err = json.Unmarshal(data, &varSamlAttributeStatementGroup)

	if err != nil {
		return err
	}

	*o = SamlAttributeStatementGroup(varSamlAttributeStatementGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "filterValue")
		delete(additionalProperties, "name")
		delete(additionalProperties, "namespace")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSamlAttributeStatementGroup struct {
	value *SamlAttributeStatementGroup
	isSet bool
}

func (v NullableSamlAttributeStatementGroup) Get() *SamlAttributeStatementGroup {
	return v.value
}

func (v *NullableSamlAttributeStatementGroup) Set(val *SamlAttributeStatementGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlAttributeStatementGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlAttributeStatementGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlAttributeStatementGroup(val *SamlAttributeStatementGroup) *NullableSamlAttributeStatementGroup {
	return &NullableSamlAttributeStatementGroup{value: val, isSet: true}
}

func (v NullableSamlAttributeStatementGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlAttributeStatementGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
