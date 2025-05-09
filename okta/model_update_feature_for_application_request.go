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
// UpdateFeatureForApplicationRequest - struct for UpdateFeatureForApplicationRequest
type UpdateFeatureForApplicationRequest struct {
	CapabilitiesInboundProvisioningObject *CapabilitiesInboundProvisioningObject
	CapabilitiesObject *CapabilitiesObject
}

// CapabilitiesInboundProvisioningObjectAsUpdateFeatureForApplicationRequest is a convenience function that returns CapabilitiesInboundProvisioningObject wrapped in UpdateFeatureForApplicationRequest
func CapabilitiesInboundProvisioningObjectAsUpdateFeatureForApplicationRequest(v *CapabilitiesInboundProvisioningObject) UpdateFeatureForApplicationRequest {
	return UpdateFeatureForApplicationRequest{
		CapabilitiesInboundProvisioningObject: v,
	}
}

// CapabilitiesObjectAsUpdateFeatureForApplicationRequest is a convenience function that returns CapabilitiesObject wrapped in UpdateFeatureForApplicationRequest
func CapabilitiesObjectAsUpdateFeatureForApplicationRequest(v *CapabilitiesObject) UpdateFeatureForApplicationRequest {
	return UpdateFeatureForApplicationRequest{
		CapabilitiesObject: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *UpdateFeatureForApplicationRequest) UnmarshalJSON(data []byte) error {
	var err error
        match := 0
        // try to unmarshal data into CapabilitiesInboundProvisioningObject
        err = json.Unmarshal(data, &dst.CapabilitiesInboundProvisioningObject)
        if err == nil {
                jsonCapabilitiesInboundProvisioningObject, _ := json.Marshal(dst.CapabilitiesInboundProvisioningObject)
                if string(jsonCapabilitiesInboundProvisioningObject) == "{}" { // empty struct
                        dst.CapabilitiesInboundProvisioningObject = nil
                } else {
                        match++
                }
        } else {
                dst.CapabilitiesInboundProvisioningObject = nil
        }

        // try to unmarshal data into CapabilitiesObject
        err = json.Unmarshal(data, &dst.CapabilitiesObject)
        if err == nil {
                jsonCapabilitiesObject, _ := json.Marshal(dst.CapabilitiesObject)
                if string(jsonCapabilitiesObject) == "{}" { // empty struct
                        dst.CapabilitiesObject = nil
                } else {
                        match++
                }
        } else {
                dst.CapabilitiesObject = nil
        }

        if match > 1 { // more than 1 match
                // reset to nil
                dst.CapabilitiesInboundProvisioningObject = nil
                dst.CapabilitiesObject = nil

                return fmt.Errorf("Data matches more than one schema in oneOf(UpdateFeatureForApplicationRequest)")
        } else if match == 1 {
                return nil // exactly one match
        } else { // no match
                return fmt.Errorf("Data failed to match schemas in oneOf(UpdateFeatureForApplicationRequest)")
        }
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateFeatureForApplicationRequest) MarshalJSON() ([]byte, error) {
	if src.CapabilitiesInboundProvisioningObject != nil {
		return json.Marshal(&src.CapabilitiesInboundProvisioningObject)
	}

	if src.CapabilitiesObject != nil {
		return json.Marshal(&src.CapabilitiesObject)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateFeatureForApplicationRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CapabilitiesInboundProvisioningObject != nil {
		return obj.CapabilitiesInboundProvisioningObject
	}

	if obj.CapabilitiesObject != nil {
		return obj.CapabilitiesObject
	}

	// all schemas are nil
	return nil
}

type NullableUpdateFeatureForApplicationRequest struct {
	value *UpdateFeatureForApplicationRequest
	isSet bool
}

func (v NullableUpdateFeatureForApplicationRequest) Get() *UpdateFeatureForApplicationRequest {
	return v.value
}

func (v *NullableUpdateFeatureForApplicationRequest) Set(val *UpdateFeatureForApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFeatureForApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFeatureForApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFeatureForApplicationRequest(val *UpdateFeatureForApplicationRequest) *NullableUpdateFeatureForApplicationRequest {
	return &NullableUpdateFeatureForApplicationRequest{value: val, isSet: true}
}

func (v NullableUpdateFeatureForApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFeatureForApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


