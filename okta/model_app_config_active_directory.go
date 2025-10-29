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
	"fmt"
	"reflect"
	"strings"
)

// checks if the AppConfigActiveDirectory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppConfigActiveDirectory{}

// AppConfigActiveDirectory struct for AppConfigActiveDirectory
type AppConfigActiveDirectory struct {
	AppConfig
	// The distinguished name of the group in Active Directory
	DistinguishedName string `json:"distinguishedName"`
	// The scope of the group in Active Directory
	GroupScope string `json:"groupScope"`
	// The type of the group in Active Directory
	GroupType string `json:"groupType"`
	// The SAM account name of the group in Active Directory
	SamAccountName       string `json:"samAccountName"`
	AdditionalProperties map[string]interface{}
}

type _AppConfigActiveDirectory AppConfigActiveDirectory

// NewAppConfigActiveDirectory instantiates a new AppConfigActiveDirectory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfigActiveDirectory(distinguishedName string, groupScope string, groupType string, samAccountName string, type_ string) *AppConfigActiveDirectory {
	this := AppConfigActiveDirectory{}
	this.Type = type_
	return &this
}

// NewAppConfigActiveDirectoryWithDefaults instantiates a new AppConfigActiveDirectory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigActiveDirectoryWithDefaults() *AppConfigActiveDirectory {
	this := AppConfigActiveDirectory{}
	return &this
}

// GetDistinguishedName returns the DistinguishedName field value
func (o *AppConfigActiveDirectory) GetDistinguishedName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DistinguishedName
}

// GetDistinguishedNameOk returns a tuple with the DistinguishedName field value
// and a boolean to check if the value has been set.
func (o *AppConfigActiveDirectory) GetDistinguishedNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DistinguishedName, true
}

// SetDistinguishedName sets field value
func (o *AppConfigActiveDirectory) SetDistinguishedName(v string) {
	o.DistinguishedName = v
}

// GetGroupScope returns the GroupScope field value
func (o *AppConfigActiveDirectory) GetGroupScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupScope
}

// GetGroupScopeOk returns a tuple with the GroupScope field value
// and a boolean to check if the value has been set.
func (o *AppConfigActiveDirectory) GetGroupScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupScope, true
}

// SetGroupScope sets field value
func (o *AppConfigActiveDirectory) SetGroupScope(v string) {
	o.GroupScope = v
}

// GetGroupType returns the GroupType field value
func (o *AppConfigActiveDirectory) GetGroupType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupType
}

// GetGroupTypeOk returns a tuple with the GroupType field value
// and a boolean to check if the value has been set.
func (o *AppConfigActiveDirectory) GetGroupTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupType, true
}

// SetGroupType sets field value
func (o *AppConfigActiveDirectory) SetGroupType(v string) {
	o.GroupType = v
}

// GetSamAccountName returns the SamAccountName field value
func (o *AppConfigActiveDirectory) GetSamAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SamAccountName
}

// GetSamAccountNameOk returns a tuple with the SamAccountName field value
// and a boolean to check if the value has been set.
func (o *AppConfigActiveDirectory) GetSamAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SamAccountName, true
}

// SetSamAccountName sets field value
func (o *AppConfigActiveDirectory) SetSamAccountName(v string) {
	o.SamAccountName = v
}

func (o AppConfigActiveDirectory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppConfigActiveDirectory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAppConfig, errAppConfig := json.Marshal(o.AppConfig)
	if errAppConfig != nil {
		return map[string]interface{}{}, errAppConfig
	}
	errAppConfig = json.Unmarshal([]byte(serializedAppConfig), &toSerialize)
	if errAppConfig != nil {
		return map[string]interface{}{}, errAppConfig
	}
	toSerialize["distinguishedName"] = o.DistinguishedName
	toSerialize["groupScope"] = o.GroupScope
	toSerialize["groupType"] = o.GroupType
	toSerialize["samAccountName"] = o.SamAccountName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppConfigActiveDirectory) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"distinguishedName",
		"groupScope",
		"groupType",
		"samAccountName",
		"type",
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

	type AppConfigActiveDirectoryWithoutEmbeddedStruct struct {
		// The distinguished name of the group in Active Directory
		DistinguishedName string `json:"distinguishedName"`
		// The scope of the group in Active Directory
		GroupScope string `json:"groupScope"`
		// The type of the group in Active Directory
		GroupType string `json:"groupType"`
		// The SAM account name of the group in Active Directory
		SamAccountName string `json:"samAccountName"`
	}

	varAppConfigActiveDirectoryWithoutEmbeddedStruct := AppConfigActiveDirectoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAppConfigActiveDirectoryWithoutEmbeddedStruct)
	if err == nil {
		varAppConfigActiveDirectory := _AppConfigActiveDirectory{}
		varAppConfigActiveDirectory.DistinguishedName = varAppConfigActiveDirectoryWithoutEmbeddedStruct.DistinguishedName
		varAppConfigActiveDirectory.GroupScope = varAppConfigActiveDirectoryWithoutEmbeddedStruct.GroupScope
		varAppConfigActiveDirectory.GroupType = varAppConfigActiveDirectoryWithoutEmbeddedStruct.GroupType
		varAppConfigActiveDirectory.SamAccountName = varAppConfigActiveDirectoryWithoutEmbeddedStruct.SamAccountName
		*o = AppConfigActiveDirectory(varAppConfigActiveDirectory)
	} else {
		return err
	}

	varAppConfigActiveDirectory := _AppConfigActiveDirectory{}

	err = json.Unmarshal(data, &varAppConfigActiveDirectory)
	if err == nil {
		o.AppConfig = varAppConfigActiveDirectory.AppConfig
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "distinguishedName")
		delete(additionalProperties, "groupScope")
		delete(additionalProperties, "groupType")
		delete(additionalProperties, "samAccountName")

		// remove fields from embedded structs
		reflectAppConfig := reflect.ValueOf(o.AppConfig)
		for i := 0; i < reflectAppConfig.Type().NumField(); i++ {
			t := reflectAppConfig.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppConfigActiveDirectory struct {
	value *AppConfigActiveDirectory
	isSet bool
}

func (v NullableAppConfigActiveDirectory) Get() *AppConfigActiveDirectory {
	return v.value
}

func (v *NullableAppConfigActiveDirectory) Set(val *AppConfigActiveDirectory) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfigActiveDirectory) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfigActiveDirectory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfigActiveDirectory(val *AppConfigActiveDirectory) *NullableAppConfigActiveDirectory {
	return &NullableAppConfigActiveDirectory{value: val, isSet: true}
}

func (v NullableAppConfigActiveDirectory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfigActiveDirectory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
