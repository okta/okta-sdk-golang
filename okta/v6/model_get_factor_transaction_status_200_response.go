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

// GetFactorTransactionStatus200Response - struct for GetFactorTransactionStatus200Response
type GetFactorTransactionStatus200Response struct {
	UserFactorPushTransaction             *UserFactorPushTransaction
	UserFactorPushTransactionRejected     *UserFactorPushTransactionRejected
	UserFactorPushTransactionTimeout      *UserFactorPushTransactionTimeout
	UserFactorPushTransactionWaitingNMC   *UserFactorPushTransactionWaitingNMC
	UserFactorPushTransactionWaitingNoNMC *UserFactorPushTransactionWaitingNoNMC
}

// UserFactorPushTransactionAsGetFactorTransactionStatus200Response is a convenience function that returns UserFactorPushTransaction wrapped in GetFactorTransactionStatus200Response
func UserFactorPushTransactionAsGetFactorTransactionStatus200Response(v *UserFactorPushTransaction) GetFactorTransactionStatus200Response {
	return GetFactorTransactionStatus200Response{
		UserFactorPushTransaction: v,
	}
}

// UserFactorPushTransactionRejectedAsGetFactorTransactionStatus200Response is a convenience function that returns UserFactorPushTransactionRejected wrapped in GetFactorTransactionStatus200Response
func UserFactorPushTransactionRejectedAsGetFactorTransactionStatus200Response(v *UserFactorPushTransactionRejected) GetFactorTransactionStatus200Response {
	return GetFactorTransactionStatus200Response{
		UserFactorPushTransactionRejected: v,
	}
}

// UserFactorPushTransactionTimeoutAsGetFactorTransactionStatus200Response is a convenience function that returns UserFactorPushTransactionTimeout wrapped in GetFactorTransactionStatus200Response
func UserFactorPushTransactionTimeoutAsGetFactorTransactionStatus200Response(v *UserFactorPushTransactionTimeout) GetFactorTransactionStatus200Response {
	return GetFactorTransactionStatus200Response{
		UserFactorPushTransactionTimeout: v,
	}
}

// UserFactorPushTransactionWaitingNMCAsGetFactorTransactionStatus200Response is a convenience function that returns UserFactorPushTransactionWaitingNMC wrapped in GetFactorTransactionStatus200Response
func UserFactorPushTransactionWaitingNMCAsGetFactorTransactionStatus200Response(v *UserFactorPushTransactionWaitingNMC) GetFactorTransactionStatus200Response {
	return GetFactorTransactionStatus200Response{
		UserFactorPushTransactionWaitingNMC: v,
	}
}

// UserFactorPushTransactionWaitingNoNMCAsGetFactorTransactionStatus200Response is a convenience function that returns UserFactorPushTransactionWaitingNoNMC wrapped in GetFactorTransactionStatus200Response
func UserFactorPushTransactionWaitingNoNMCAsGetFactorTransactionStatus200Response(v *UserFactorPushTransactionWaitingNoNMC) GetFactorTransactionStatus200Response {
	return GetFactorTransactionStatus200Response{
		UserFactorPushTransactionWaitingNoNMC: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetFactorTransactionStatus200Response) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'REJECTED'
	if jsonDict["factorResult"] == "REJECTED" {
		// try to unmarshal JSON data into UserFactorPushTransactionRejected
		err = json.Unmarshal(data, &dst.UserFactorPushTransactionRejected)
		if err == nil {
			return nil // data stored in dst.UserFactorPushTransactionRejected, return on the first match
		} else {
			dst.UserFactorPushTransactionRejected = nil
			return fmt.Errorf("failed to unmarshal GetFactorTransactionStatus200Response as UserFactorPushTransactionRejected: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SUCCESS'
	if jsonDict["factorResult"] == "SUCCESS" {
		// try to unmarshal JSON data into UserFactorPushTransaction
		err = json.Unmarshal(data, &dst.UserFactorPushTransaction)
		if err == nil {
			return nil // data stored in dst.UserFactorPushTransaction, return on the first match
		} else {
			dst.UserFactorPushTransaction = nil
			return fmt.Errorf("failed to unmarshal GetFactorTransactionStatus200Response as UserFactorPushTransaction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TIMEOUT'
	if jsonDict["factorResult"] == "TIMEOUT" {
		// try to unmarshal JSON data into UserFactorPushTransactionTimeout
		err = json.Unmarshal(data, &dst.UserFactorPushTransactionTimeout)
		if err == nil {
			return nil // data stored in dst.UserFactorPushTransactionTimeout, return on the first match
		} else {
			dst.UserFactorPushTransactionTimeout = nil
			return fmt.Errorf("failed to unmarshal GetFactorTransactionStatus200Response as UserFactorPushTransactionTimeout: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WAITING'
	if jsonDict["factorResult"] == "WAITING" {
		// try to unmarshal JSON data into UserFactorPushTransactionWaitingNoNMC
		err = json.Unmarshal(data, &dst.UserFactorPushTransactionWaitingNoNMC)
		if err == nil {
			return nil // data stored in dst.UserFactorPushTransactionWaitingNoNMC, return on the first match
		} else {
			dst.UserFactorPushTransactionWaitingNoNMC = nil
			return fmt.Errorf("failed to unmarshal GetFactorTransactionStatus200Response as UserFactorPushTransactionWaitingNoNMC: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WAITING (with number matching challenge)'
	if jsonDict["factorResult"] == "WAITING (with number matching challenge)" {
		// try to unmarshal JSON data into UserFactorPushTransactionWaitingNMC
		err = json.Unmarshal(data, &dst.UserFactorPushTransactionWaitingNMC)
		if err == nil {
			return nil // data stored in dst.UserFactorPushTransactionWaitingNMC, return on the first match
		} else {
			dst.UserFactorPushTransactionWaitingNMC = nil
			return fmt.Errorf("failed to unmarshal GetFactorTransactionStatus200Response as UserFactorPushTransactionWaitingNMC: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetFactorTransactionStatus200Response) MarshalJSON() ([]byte, error) {
	if src.UserFactorPushTransaction != nil {
		return json.Marshal(&src.UserFactorPushTransaction)
	}

	if src.UserFactorPushTransactionRejected != nil {
		return json.Marshal(&src.UserFactorPushTransactionRejected)
	}

	if src.UserFactorPushTransactionTimeout != nil {
		return json.Marshal(&src.UserFactorPushTransactionTimeout)
	}

	if src.UserFactorPushTransactionWaitingNMC != nil {
		return json.Marshal(&src.UserFactorPushTransactionWaitingNMC)
	}

	if src.UserFactorPushTransactionWaitingNoNMC != nil {
		return json.Marshal(&src.UserFactorPushTransactionWaitingNoNMC)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetFactorTransactionStatus200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UserFactorPushTransaction != nil {
		return obj.UserFactorPushTransaction
	}

	if obj.UserFactorPushTransactionRejected != nil {
		return obj.UserFactorPushTransactionRejected
	}

	if obj.UserFactorPushTransactionTimeout != nil {
		return obj.UserFactorPushTransactionTimeout
	}

	if obj.UserFactorPushTransactionWaitingNMC != nil {
		return obj.UserFactorPushTransactionWaitingNMC
	}

	if obj.UserFactorPushTransactionWaitingNoNMC != nil {
		return obj.UserFactorPushTransactionWaitingNoNMC
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj GetFactorTransactionStatus200Response) GetActualInstanceValue() interface{} {
	if obj.UserFactorPushTransaction != nil {
		return *obj.UserFactorPushTransaction
	}

	if obj.UserFactorPushTransactionRejected != nil {
		return *obj.UserFactorPushTransactionRejected
	}

	if obj.UserFactorPushTransactionTimeout != nil {
		return *obj.UserFactorPushTransactionTimeout
	}

	if obj.UserFactorPushTransactionWaitingNMC != nil {
		return *obj.UserFactorPushTransactionWaitingNMC
	}

	if obj.UserFactorPushTransactionWaitingNoNMC != nil {
		return *obj.UserFactorPushTransactionWaitingNoNMC
	}

	// all schemas are nil
	return nil
}

type NullableGetFactorTransactionStatus200Response struct {
	value *GetFactorTransactionStatus200Response
	isSet bool
}

func (v NullableGetFactorTransactionStatus200Response) Get() *GetFactorTransactionStatus200Response {
	return v.value
}

func (v *NullableGetFactorTransactionStatus200Response) Set(val *GetFactorTransactionStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFactorTransactionStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFactorTransactionStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFactorTransactionStatus200Response(val *GetFactorTransactionStatus200Response) *NullableGetFactorTransactionStatus200Response {
	return &NullableGetFactorTransactionStatus200Response{value: val, isSet: true}
}

func (v NullableGetFactorTransactionStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFactorTransactionStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
