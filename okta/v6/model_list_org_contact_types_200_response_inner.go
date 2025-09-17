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
	"fmt"
)

// ListOrgContactTypes200ResponseInner - struct for ListOrgContactTypes200ResponseInner
type ListOrgContactTypes200ResponseInner struct {
	OrgBillingContactType   *OrgBillingContactType
	OrgTechnicalContactType *OrgTechnicalContactType
}

// OrgBillingContactTypeAsListOrgContactTypes200ResponseInner is a convenience function that returns OrgBillingContactType wrapped in ListOrgContactTypes200ResponseInner
func OrgBillingContactTypeAsListOrgContactTypes200ResponseInner(v *OrgBillingContactType) ListOrgContactTypes200ResponseInner {
	return ListOrgContactTypes200ResponseInner{
		OrgBillingContactType: v,
	}
}

// OrgTechnicalContactTypeAsListOrgContactTypes200ResponseInner is a convenience function that returns OrgTechnicalContactType wrapped in ListOrgContactTypes200ResponseInner
func OrgTechnicalContactTypeAsListOrgContactTypes200ResponseInner(v *OrgTechnicalContactType) ListOrgContactTypes200ResponseInner {
	return ListOrgContactTypes200ResponseInner{
		OrgTechnicalContactType: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListOrgContactTypes200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'BILLING'
	if jsonDict["contactType"] == "BILLING" {
		// try to unmarshal JSON data into OrgBillingContactType
		err = json.Unmarshal(data, &dst.OrgBillingContactType)
		if err == nil {
			return nil // data stored in dst.OrgBillingContactType, return on the first match
		} else {
			dst.OrgBillingContactType = nil
			return fmt.Errorf("failed to unmarshal ListOrgContactTypes200ResponseInner as OrgBillingContactType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TECHNICAL'
	if jsonDict["contactType"] == "TECHNICAL" {
		// try to unmarshal JSON data into OrgTechnicalContactType
		err = json.Unmarshal(data, &dst.OrgTechnicalContactType)
		if err == nil {
			return nil // data stored in dst.OrgTechnicalContactType, return on the first match
		} else {
			dst.OrgTechnicalContactType = nil
			return fmt.Errorf("failed to unmarshal ListOrgContactTypes200ResponseInner as OrgTechnicalContactType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'orgBillingContactType'
	if jsonDict["contactType"] == "orgBillingContactType" {
		// try to unmarshal JSON data into OrgBillingContactType
		err = json.Unmarshal(data, &dst.OrgBillingContactType)
		if err == nil {
			return nil // data stored in dst.OrgBillingContactType, return on the first match
		} else {
			dst.OrgBillingContactType = nil
			return fmt.Errorf("failed to unmarshal ListOrgContactTypes200ResponseInner as OrgBillingContactType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'orgTechnicalContactType'
	if jsonDict["contactType"] == "orgTechnicalContactType" {
		// try to unmarshal JSON data into OrgTechnicalContactType
		err = json.Unmarshal(data, &dst.OrgTechnicalContactType)
		if err == nil {
			return nil // data stored in dst.OrgTechnicalContactType, return on the first match
		} else {
			dst.OrgTechnicalContactType = nil
			return fmt.Errorf("failed to unmarshal ListOrgContactTypes200ResponseInner as OrgTechnicalContactType: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListOrgContactTypes200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.OrgBillingContactType != nil {
		return json.Marshal(&src.OrgBillingContactType)
	}

	if src.OrgTechnicalContactType != nil {
		return json.Marshal(&src.OrgTechnicalContactType)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListOrgContactTypes200ResponseInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.OrgBillingContactType != nil {
		return obj.OrgBillingContactType
	}

	if obj.OrgTechnicalContactType != nil {
		return obj.OrgTechnicalContactType
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ListOrgContactTypes200ResponseInner) GetActualInstanceValue() interface{} {
	if obj.OrgBillingContactType != nil {
		return *obj.OrgBillingContactType
	}

	if obj.OrgTechnicalContactType != nil {
		return *obj.OrgTechnicalContactType
	}

	// all schemas are nil
	return nil
}

type NullableListOrgContactTypes200ResponseInner struct {
	value *ListOrgContactTypes200ResponseInner
	isSet bool
}

func (v NullableListOrgContactTypes200ResponseInner) Get() *ListOrgContactTypes200ResponseInner {
	return v.value
}

func (v *NullableListOrgContactTypes200ResponseInner) Set(val *ListOrgContactTypes200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrgContactTypes200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrgContactTypes200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrgContactTypes200ResponseInner(val *ListOrgContactTypes200ResponseInner) *NullableListOrgContactTypes200ResponseInner {
	return &NullableListOrgContactTypes200ResponseInner{value: val, isSet: true}
}

func (v NullableListOrgContactTypes200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrgContactTypes200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
