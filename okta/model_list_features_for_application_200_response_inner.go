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

// ListFeaturesForApplication200ResponseInner - struct for ListFeaturesForApplication200ResponseInner
type ListFeaturesForApplication200ResponseInner struct {
	InboundProvisioningApplicationFeature *InboundProvisioningApplicationFeature
	UserProvisioningApplicationFeature    *UserProvisioningApplicationFeature
}

// InboundProvisioningApplicationFeatureAsListFeaturesForApplication200ResponseInner is a convenience function that returns InboundProvisioningApplicationFeature wrapped in ListFeaturesForApplication200ResponseInner
func InboundProvisioningApplicationFeatureAsListFeaturesForApplication200ResponseInner(v *InboundProvisioningApplicationFeature) ListFeaturesForApplication200ResponseInner {
	return ListFeaturesForApplication200ResponseInner{
		InboundProvisioningApplicationFeature: v,
	}
}

// UserProvisioningApplicationFeatureAsListFeaturesForApplication200ResponseInner is a convenience function that returns UserProvisioningApplicationFeature wrapped in ListFeaturesForApplication200ResponseInner
func UserProvisioningApplicationFeatureAsListFeaturesForApplication200ResponseInner(v *UserProvisioningApplicationFeature) ListFeaturesForApplication200ResponseInner {
	return ListFeaturesForApplication200ResponseInner{
		UserProvisioningApplicationFeature: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListFeaturesForApplication200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// Get discriminator value, treating nil/missing as empty string for comparison
	discriminatorValue, _ := jsonDict["name"].(string)

	// check if the discriminator value is 'INBOUND_PROVISIONING'
	if discriminatorValue == "INBOUND_PROVISIONING" {
		// try to unmarshal JSON data into InboundProvisioningApplicationFeature
		err = json.Unmarshal(data, &dst.InboundProvisioningApplicationFeature)
		if err == nil {
			return nil // data stored in dst.InboundProvisioningApplicationFeature, return on the first match
		} else {
			dst.InboundProvisioningApplicationFeature = nil
			return fmt.Errorf("failed to unmarshal ListFeaturesForApplication200ResponseInner as InboundProvisioningApplicationFeature: %s", err.Error())
		}
	}

	// check if the discriminator value is 'USER_PROVISIONING'
	if discriminatorValue == "USER_PROVISIONING" {
		// try to unmarshal JSON data into UserProvisioningApplicationFeature
		err = json.Unmarshal(data, &dst.UserProvisioningApplicationFeature)
		if err == nil {
			return nil // data stored in dst.UserProvisioningApplicationFeature, return on the first match
		} else {
			dst.UserProvisioningApplicationFeature = nil
			return fmt.Errorf("failed to unmarshal ListFeaturesForApplication200ResponseInner as UserProvisioningApplicationFeature: %s", err.Error())
		}
	}

	// If discriminator value is empty/missing, default to the last mapped model (typically the most common type)
	if discriminatorValue == "" {
		err = json.Unmarshal(data, &dst.UserProvisioningApplicationFeature)
		if err == nil {
			return nil
		}
		dst.UserProvisioningApplicationFeature = nil
	}

	// No match found or unmarshal failed - return nil to allow partial unmarshalling
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListFeaturesForApplication200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.InboundProvisioningApplicationFeature != nil {
		return json.Marshal(&src.InboundProvisioningApplicationFeature)
	}

	if src.UserProvisioningApplicationFeature != nil {
		return json.Marshal(&src.UserProvisioningApplicationFeature)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListFeaturesForApplication200ResponseInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.InboundProvisioningApplicationFeature != nil {
		return obj.InboundProvisioningApplicationFeature
	}

	if obj.UserProvisioningApplicationFeature != nil {
		return obj.UserProvisioningApplicationFeature
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ListFeaturesForApplication200ResponseInner) GetActualInstanceValue() interface{} {
	if obj.InboundProvisioningApplicationFeature != nil {
		return *obj.InboundProvisioningApplicationFeature
	}

	if obj.UserProvisioningApplicationFeature != nil {
		return *obj.UserProvisioningApplicationFeature
	}

	// all schemas are nil
	return nil
}

type NullableListFeaturesForApplication200ResponseInner struct {
	value *ListFeaturesForApplication200ResponseInner
	isSet bool
}

func (v NullableListFeaturesForApplication200ResponseInner) Get() *ListFeaturesForApplication200ResponseInner {
	return v.value
}

func (v *NullableListFeaturesForApplication200ResponseInner) Set(val *ListFeaturesForApplication200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListFeaturesForApplication200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListFeaturesForApplication200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFeaturesForApplication200ResponseInner(val *ListFeaturesForApplication200ResponseInner) *NullableListFeaturesForApplication200ResponseInner {
	return &NullableListFeaturesForApplication200ResponseInner{value: val, isSet: true}
}

func (v NullableListFeaturesForApplication200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFeaturesForApplication200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
