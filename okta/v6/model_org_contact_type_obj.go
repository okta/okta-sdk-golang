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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the OrgContactTypeObj type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgContactTypeObj{}

// OrgContactTypeObj struct for OrgContactTypeObj
type OrgContactTypeObj struct {
	OrgBillingContactType *OrgBillingContactType
	OrgTechnicalContactType *OrgTechnicalContactType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *OrgContactTypeObj) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'BILLING'
	if jsonDict["contactType"] == "BILLING" {
		// try to unmarshal JSON data into OrgBillingContactType
		err = json.Unmarshal(data, &dst.OrgBillingContactType);
		if err == nil {
			jsonOrgBillingContactType, _ := json.Marshal(dst.OrgBillingContactType)
			if string(jsonOrgBillingContactType) == "{}" { // empty struct
				dst.OrgBillingContactType = nil
			} else {
				return nil // data stored in dst.OrgBillingContactType, return on the first match
			}
		} else {
			dst.OrgBillingContactType = nil
		}
	}

	// check if the discriminator value is 'TECHNICAL'
	if jsonDict["contactType"] == "TECHNICAL" {
		// try to unmarshal JSON data into OrgTechnicalContactType
		err = json.Unmarshal(data, &dst.OrgTechnicalContactType);
		if err == nil {
			jsonOrgTechnicalContactType, _ := json.Marshal(dst.OrgTechnicalContactType)
			if string(jsonOrgTechnicalContactType) == "{}" { // empty struct
				dst.OrgTechnicalContactType = nil
			} else {
				return nil // data stored in dst.OrgTechnicalContactType, return on the first match
			}
		} else {
			dst.OrgTechnicalContactType = nil
		}
	}

	// check if the discriminator value is 'orgBillingContactType'
	if jsonDict["contactType"] == "orgBillingContactType" {
		// try to unmarshal JSON data into OrgBillingContactType
		err = json.Unmarshal(data, &dst.OrgBillingContactType);
		if err == nil {
			jsonOrgBillingContactType, _ := json.Marshal(dst.OrgBillingContactType)
			if string(jsonOrgBillingContactType) == "{}" { // empty struct
				dst.OrgBillingContactType = nil
			} else {
				return nil // data stored in dst.OrgBillingContactType, return on the first match
			}
		} else {
			dst.OrgBillingContactType = nil
		}
	}

	// check if the discriminator value is 'orgTechnicalContactType'
	if jsonDict["contactType"] == "orgTechnicalContactType" {
		// try to unmarshal JSON data into OrgTechnicalContactType
		err = json.Unmarshal(data, &dst.OrgTechnicalContactType);
		if err == nil {
			jsonOrgTechnicalContactType, _ := json.Marshal(dst.OrgTechnicalContactType)
			if string(jsonOrgTechnicalContactType) == "{}" { // empty struct
				dst.OrgTechnicalContactType = nil
			} else {
				return nil // data stored in dst.OrgTechnicalContactType, return on the first match
			}
		} else {
			dst.OrgTechnicalContactType = nil
		}
	}

	// try to unmarshal JSON data into OrgBillingContactType
	err = json.Unmarshal(data, &dst.OrgBillingContactType);
	if err == nil {
		jsonOrgBillingContactType, _ := json.Marshal(dst.OrgBillingContactType)
		if string(jsonOrgBillingContactType) == "{}" { // empty struct
			dst.OrgBillingContactType = nil
		} else {
			return nil // data stored in dst.OrgBillingContactType, return on the first match
		}
	} else {
		dst.OrgBillingContactType = nil
	}

	// try to unmarshal JSON data into OrgTechnicalContactType
	err = json.Unmarshal(data, &dst.OrgTechnicalContactType);
	if err == nil {
		jsonOrgTechnicalContactType, _ := json.Marshal(dst.OrgTechnicalContactType)
		if string(jsonOrgTechnicalContactType) == "{}" { // empty struct
			dst.OrgTechnicalContactType = nil
		} else {
			return nil // data stored in dst.OrgTechnicalContactType, return on the first match
		}
	} else {
		dst.OrgTechnicalContactType = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(OrgContactTypeObj)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OrgContactTypeObj) MarshalJSON() ([]byte, error) {
	if src.OrgBillingContactType != nil {
		return json.Marshal(&src.OrgBillingContactType)
	}

	if src.OrgTechnicalContactType != nil {
		return json.Marshal(&src.OrgTechnicalContactType)
	}

	return nil, nil // no data in anyOf schemas
}

func (src OrgContactTypeObj) ToMap() (map[string]interface{}, error) {
	if src.OrgBillingContactType != nil {
		return src.OrgBillingContactType.ToMap()
	}

	if src.OrgTechnicalContactType != nil {
		return src.OrgTechnicalContactType.ToMap()
	}

    return nil, nil // no data in anyOf schemas
}

type NullableOrgContactTypeObj struct {
	value *OrgContactTypeObj
	isSet bool
}

func (v NullableOrgContactTypeObj) Get() *OrgContactTypeObj {
	return v.value
}

func (v *NullableOrgContactTypeObj) Set(val *OrgContactTypeObj) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgContactTypeObj) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgContactTypeObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgContactTypeObj(val *OrgContactTypeObj) *NullableOrgContactTypeObj {
	return &NullableOrgContactTypeObj{value: val, isSet: true}
}

func (v NullableOrgContactTypeObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgContactTypeObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


