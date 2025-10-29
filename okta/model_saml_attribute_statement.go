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
)

// checks if the SamlAttributeStatement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlAttributeStatement{}

// SamlAttributeStatement struct for SamlAttributeStatement
type SamlAttributeStatement struct {
	SamlAttributeStatementExpression *SamlAttributeStatementExpression
	SamlAttributeStatementGroup      *SamlAttributeStatementGroup
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SamlAttributeStatement) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'EXPRESSION'
	if jsonDict["type"] == "EXPRESSION" {
		// try to unmarshal JSON data into SamlAttributeStatementExpression
		err = json.Unmarshal(data, &dst.SamlAttributeStatementExpression)
		if err == nil {
			jsonSamlAttributeStatementExpression, _ := json.Marshal(dst.SamlAttributeStatementExpression)
			if string(jsonSamlAttributeStatementExpression) == "{}" { // empty struct
				dst.SamlAttributeStatementExpression = nil
			} else {
				return nil // data stored in dst.SamlAttributeStatementExpression, return on the first match
			}
		} else {
			dst.SamlAttributeStatementExpression = nil
		}
	}

	// check if the discriminator value is 'GROUP'
	if jsonDict["type"] == "GROUP" {
		// try to unmarshal JSON data into SamlAttributeStatementGroup
		err = json.Unmarshal(data, &dst.SamlAttributeStatementGroup)
		if err == nil {
			jsonSamlAttributeStatementGroup, _ := json.Marshal(dst.SamlAttributeStatementGroup)
			if string(jsonSamlAttributeStatementGroup) == "{}" { // empty struct
				dst.SamlAttributeStatementGroup = nil
			} else {
				return nil // data stored in dst.SamlAttributeStatementGroup, return on the first match
			}
		} else {
			dst.SamlAttributeStatementGroup = nil
		}
	}

	// try to unmarshal JSON data into SamlAttributeStatementExpression
	err = json.Unmarshal(data, &dst.SamlAttributeStatementExpression)
	if err == nil {
		jsonSamlAttributeStatementExpression, _ := json.Marshal(dst.SamlAttributeStatementExpression)
		if string(jsonSamlAttributeStatementExpression) == "{}" { // empty struct
			dst.SamlAttributeStatementExpression = nil
		} else {
			return nil // data stored in dst.SamlAttributeStatementExpression, return on the first match
		}
	} else {
		dst.SamlAttributeStatementExpression = nil
	}

	// try to unmarshal JSON data into SamlAttributeStatementGroup
	err = json.Unmarshal(data, &dst.SamlAttributeStatementGroup)
	if err == nil {
		jsonSamlAttributeStatementGroup, _ := json.Marshal(dst.SamlAttributeStatementGroup)
		if string(jsonSamlAttributeStatementGroup) == "{}" { // empty struct
			dst.SamlAttributeStatementGroup = nil
		} else {
			return nil // data stored in dst.SamlAttributeStatementGroup, return on the first match
		}
	} else {
		dst.SamlAttributeStatementGroup = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SamlAttributeStatement)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SamlAttributeStatement) MarshalJSON() ([]byte, error) {
	if src.SamlAttributeStatementExpression != nil {
		return json.Marshal(&src.SamlAttributeStatementExpression)
	}

	if src.SamlAttributeStatementGroup != nil {
		return json.Marshal(&src.SamlAttributeStatementGroup)
	}

	return nil, nil // no data in anyOf schemas
}

func (src SamlAttributeStatement) ToMap() (map[string]interface{}, error) {
	if src.SamlAttributeStatementExpression != nil {
		return src.SamlAttributeStatementExpression.ToMap()
	}

	if src.SamlAttributeStatementGroup != nil {
		return src.SamlAttributeStatementGroup.ToMap()
	}

	return nil, nil // no data in anyOf schemas
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
