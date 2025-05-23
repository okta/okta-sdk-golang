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
	"fmt"
)


//model_oneof.mustache
// ListNetworkZones200ResponseInner - struct for ListNetworkZones200ResponseInner
type ListNetworkZones200ResponseInner struct {
	DynamicNetworkZone *DynamicNetworkZone
	EnhancedDynamicNetworkZone *EnhancedDynamicNetworkZone
	IPNetworkZone *IPNetworkZone
}

// DynamicNetworkZoneAsListNetworkZones200ResponseInner is a convenience function that returns DynamicNetworkZone wrapped in ListNetworkZones200ResponseInner
func DynamicNetworkZoneAsListNetworkZones200ResponseInner(v *DynamicNetworkZone) ListNetworkZones200ResponseInner {
	return ListNetworkZones200ResponseInner{
		DynamicNetworkZone: v,
	}
}

// EnhancedDynamicNetworkZoneAsListNetworkZones200ResponseInner is a convenience function that returns EnhancedDynamicNetworkZone wrapped in ListNetworkZones200ResponseInner
func EnhancedDynamicNetworkZoneAsListNetworkZones200ResponseInner(v *EnhancedDynamicNetworkZone) ListNetworkZones200ResponseInner {
	return ListNetworkZones200ResponseInner{
		EnhancedDynamicNetworkZone: v,
	}
}

// IPNetworkZoneAsListNetworkZones200ResponseInner is a convenience function that returns IPNetworkZone wrapped in ListNetworkZones200ResponseInner
func IPNetworkZoneAsListNetworkZones200ResponseInner(v *IPNetworkZone) ListNetworkZones200ResponseInner {
	return ListNetworkZones200ResponseInner{
		IPNetworkZone: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *ListNetworkZones200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'DYNAMIC'
	if jsonDict["type"] == "DYNAMIC" {
		// try to unmarshal JSON data into DynamicNetworkZone
		err = json.Unmarshal(data, &dst.DynamicNetworkZone)
		if err == nil {
			return nil // data stored in dst.DynamicNetworkZone, return on the first match
		} else {
			dst.DynamicNetworkZone = nil
			return fmt.Errorf("Failed to unmarshal ListNetworkZones200ResponseInner as DynamicNetworkZone: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DYNAMIC_V2'
	if jsonDict["type"] == "DYNAMIC_V2" {
		// try to unmarshal JSON data into EnhancedDynamicNetworkZone
		err = json.Unmarshal(data, &dst.EnhancedDynamicNetworkZone)
		if err == nil {
			return nil // data stored in dst.EnhancedDynamicNetworkZone, return on the first match
		} else {
			dst.EnhancedDynamicNetworkZone = nil
			return fmt.Errorf("Failed to unmarshal ListNetworkZones200ResponseInner as EnhancedDynamicNetworkZone: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DynamicNetworkZone'
	if jsonDict["type"] == "DynamicNetworkZone" {
		// try to unmarshal JSON data into DynamicNetworkZone
		err = json.Unmarshal(data, &dst.DynamicNetworkZone)
		if err == nil {
			return nil // data stored in dst.DynamicNetworkZone, return on the first match
		} else {
			dst.DynamicNetworkZone = nil
			return fmt.Errorf("Failed to unmarshal ListNetworkZones200ResponseInner as DynamicNetworkZone: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EnhancedDynamicNetworkZone'
	if jsonDict["type"] == "EnhancedDynamicNetworkZone" {
		// try to unmarshal JSON data into EnhancedDynamicNetworkZone
		err = json.Unmarshal(data, &dst.EnhancedDynamicNetworkZone)
		if err == nil {
			return nil // data stored in dst.EnhancedDynamicNetworkZone, return on the first match
		} else {
			dst.EnhancedDynamicNetworkZone = nil
			return fmt.Errorf("Failed to unmarshal ListNetworkZones200ResponseInner as EnhancedDynamicNetworkZone: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IP'
	if jsonDict["type"] == "IP" {
		// try to unmarshal JSON data into IPNetworkZone
		err = json.Unmarshal(data, &dst.IPNetworkZone)
		if err == nil {
			return nil // data stored in dst.IPNetworkZone, return on the first match
		} else {
			dst.IPNetworkZone = nil
			return fmt.Errorf("Failed to unmarshal ListNetworkZones200ResponseInner as IPNetworkZone: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IPNetworkZone'
	if jsonDict["type"] == "IPNetworkZone" {
		// try to unmarshal JSON data into IPNetworkZone
		err = json.Unmarshal(data, &dst.IPNetworkZone)
		if err == nil {
			return nil // data stored in dst.IPNetworkZone, return on the first match
		} else {
			dst.IPNetworkZone = nil
			return fmt.Errorf("Failed to unmarshal ListNetworkZones200ResponseInner as IPNetworkZone: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListNetworkZones200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.DynamicNetworkZone != nil {
		return json.Marshal(&src.DynamicNetworkZone)
	}

	if src.EnhancedDynamicNetworkZone != nil {
		return json.Marshal(&src.EnhancedDynamicNetworkZone)
	}

	if src.IPNetworkZone != nil {
		return json.Marshal(&src.IPNetworkZone)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListNetworkZones200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DynamicNetworkZone != nil {
		return obj.DynamicNetworkZone
	}

	if obj.EnhancedDynamicNetworkZone != nil {
		return obj.EnhancedDynamicNetworkZone
	}

	if obj.IPNetworkZone != nil {
		return obj.IPNetworkZone
	}

	// all schemas are nil
	return nil
}

type NullableListNetworkZones200ResponseInner struct {
	value *ListNetworkZones200ResponseInner
	isSet bool
}

func (v NullableListNetworkZones200ResponseInner) Get() *ListNetworkZones200ResponseInner {
	return v.value
}

func (v *NullableListNetworkZones200ResponseInner) Set(val *ListNetworkZones200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListNetworkZones200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListNetworkZones200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListNetworkZones200ResponseInner(val *ListNetworkZones200ResponseInner) *NullableListNetworkZones200ResponseInner {
	return &NullableListNetworkZones200ResponseInner{value: val, isSet: true}
}

func (v NullableListNetworkZones200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListNetworkZones200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


