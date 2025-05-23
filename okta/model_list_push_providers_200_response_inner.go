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
// ListPushProviders200ResponseInner - struct for ListPushProviders200ResponseInner
type ListPushProviders200ResponseInner struct {
	APNSPushProvider *APNSPushProvider
	FCMPushProvider *FCMPushProvider
}

// APNSPushProviderAsListPushProviders200ResponseInner is a convenience function that returns APNSPushProvider wrapped in ListPushProviders200ResponseInner
func APNSPushProviderAsListPushProviders200ResponseInner(v *APNSPushProvider) ListPushProviders200ResponseInner {
	return ListPushProviders200ResponseInner{
		APNSPushProvider: v,
	}
}

// FCMPushProviderAsListPushProviders200ResponseInner is a convenience function that returns FCMPushProvider wrapped in ListPushProviders200ResponseInner
func FCMPushProviderAsListPushProviders200ResponseInner(v *FCMPushProvider) ListPushProviders200ResponseInner {
	return ListPushProviders200ResponseInner{
		FCMPushProvider: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *ListPushProviders200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'APNS'
	if jsonDict["providerType"] == "APNS" {
		// try to unmarshal JSON data into APNSPushProvider
		err = json.Unmarshal(data, &dst.APNSPushProvider)
		if err == nil {
			return nil // data stored in dst.APNSPushProvider, return on the first match
		} else {
			dst.APNSPushProvider = nil
			return fmt.Errorf("Failed to unmarshal ListPushProviders200ResponseInner as APNSPushProvider: %s", err.Error())
		}
	}

	// check if the discriminator value is 'APNSPushProvider'
	if jsonDict["providerType"] == "APNSPushProvider" {
		// try to unmarshal JSON data into APNSPushProvider
		err = json.Unmarshal(data, &dst.APNSPushProvider)
		if err == nil {
			return nil // data stored in dst.APNSPushProvider, return on the first match
		} else {
			dst.APNSPushProvider = nil
			return fmt.Errorf("Failed to unmarshal ListPushProviders200ResponseInner as APNSPushProvider: %s", err.Error())
		}
	}

	// check if the discriminator value is 'FCM'
	if jsonDict["providerType"] == "FCM" {
		// try to unmarshal JSON data into FCMPushProvider
		err = json.Unmarshal(data, &dst.FCMPushProvider)
		if err == nil {
			return nil // data stored in dst.FCMPushProvider, return on the first match
		} else {
			dst.FCMPushProvider = nil
			return fmt.Errorf("Failed to unmarshal ListPushProviders200ResponseInner as FCMPushProvider: %s", err.Error())
		}
	}

	// check if the discriminator value is 'FCMPushProvider'
	if jsonDict["providerType"] == "FCMPushProvider" {
		// try to unmarshal JSON data into FCMPushProvider
		err = json.Unmarshal(data, &dst.FCMPushProvider)
		if err == nil {
			return nil // data stored in dst.FCMPushProvider, return on the first match
		} else {
			dst.FCMPushProvider = nil
			return fmt.Errorf("Failed to unmarshal ListPushProviders200ResponseInner as FCMPushProvider: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListPushProviders200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.APNSPushProvider != nil {
		return json.Marshal(&src.APNSPushProvider)
	}

	if src.FCMPushProvider != nil {
		return json.Marshal(&src.FCMPushProvider)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListPushProviders200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.APNSPushProvider != nil {
		return obj.APNSPushProvider
	}

	if obj.FCMPushProvider != nil {
		return obj.FCMPushProvider
	}

	// all schemas are nil
	return nil
}

type NullableListPushProviders200ResponseInner struct {
	value *ListPushProviders200ResponseInner
	isSet bool
}

func (v NullableListPushProviders200ResponseInner) Get() *ListPushProviders200ResponseInner {
	return v.value
}

func (v *NullableListPushProviders200ResponseInner) Set(val *ListPushProviders200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListPushProviders200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListPushProviders200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPushProviders200ResponseInner(val *ListPushProviders200ResponseInner) *NullableListPushProviders200ResponseInner {
	return &NullableListPushProviders200ResponseInner{value: val, isSet: true}
}

func (v NullableListPushProviders200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPushProviders200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


