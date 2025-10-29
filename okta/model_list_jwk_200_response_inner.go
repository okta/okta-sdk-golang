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
	"fmt"
)

// ListJwk200ResponseInner - struct for ListJwk200ResponseInner
type ListJwk200ResponseInner struct {
	OAuth2ClientJsonEncryptionKeyResponse *OAuth2ClientJsonEncryptionKeyResponse
	OAuth2ClientJsonSigningKeyResponse    *OAuth2ClientJsonSigningKeyResponse
}

// OAuth2ClientJsonEncryptionKeyResponseAsListJwk200ResponseInner is a convenience function that returns OAuth2ClientJsonEncryptionKeyResponse wrapped in ListJwk200ResponseInner
func OAuth2ClientJsonEncryptionKeyResponseAsListJwk200ResponseInner(v *OAuth2ClientJsonEncryptionKeyResponse) ListJwk200ResponseInner {
	return ListJwk200ResponseInner{
		OAuth2ClientJsonEncryptionKeyResponse: v,
	}
}

// OAuth2ClientJsonSigningKeyResponseAsListJwk200ResponseInner is a convenience function that returns OAuth2ClientJsonSigningKeyResponse wrapped in ListJwk200ResponseInner
func OAuth2ClientJsonSigningKeyResponseAsListJwk200ResponseInner(v *OAuth2ClientJsonSigningKeyResponse) ListJwk200ResponseInner {
	return ListJwk200ResponseInner{
		OAuth2ClientJsonSigningKeyResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListJwk200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OAuth2ClientJsonEncryptionKeyResponse
	err = json.Unmarshal(data, &dst.OAuth2ClientJsonEncryptionKeyResponse)
	if err == nil {
		jsonOAuth2ClientJsonEncryptionKeyResponse, _ := json.Marshal(dst.OAuth2ClientJsonEncryptionKeyResponse)
		if string(jsonOAuth2ClientJsonEncryptionKeyResponse) == "{}" { // empty struct
			dst.OAuth2ClientJsonEncryptionKeyResponse = nil
		} else {
			match++
		}
	} else {
		dst.OAuth2ClientJsonEncryptionKeyResponse = nil
	}

	// try to unmarshal data into OAuth2ClientJsonSigningKeyResponse
	err = json.Unmarshal(data, &dst.OAuth2ClientJsonSigningKeyResponse)
	if err == nil {
		jsonOAuth2ClientJsonSigningKeyResponse, _ := json.Marshal(dst.OAuth2ClientJsonSigningKeyResponse)
		if string(jsonOAuth2ClientJsonSigningKeyResponse) == "{}" { // empty struct
			dst.OAuth2ClientJsonSigningKeyResponse = nil
		} else {
			match++
		}
	} else {
		dst.OAuth2ClientJsonSigningKeyResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.OAuth2ClientJsonEncryptionKeyResponse = nil
		dst.OAuth2ClientJsonSigningKeyResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListJwk200ResponseInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListJwk200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListJwk200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.OAuth2ClientJsonEncryptionKeyResponse != nil {
		return json.Marshal(&src.OAuth2ClientJsonEncryptionKeyResponse)
	}

	if src.OAuth2ClientJsonSigningKeyResponse != nil {
		return json.Marshal(&src.OAuth2ClientJsonSigningKeyResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListJwk200ResponseInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.OAuth2ClientJsonEncryptionKeyResponse != nil {
		return obj.OAuth2ClientJsonEncryptionKeyResponse
	}

	if obj.OAuth2ClientJsonSigningKeyResponse != nil {
		return obj.OAuth2ClientJsonSigningKeyResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ListJwk200ResponseInner) GetActualInstanceValue() interface{} {
	if obj.OAuth2ClientJsonEncryptionKeyResponse != nil {
		return *obj.OAuth2ClientJsonEncryptionKeyResponse
	}

	if obj.OAuth2ClientJsonSigningKeyResponse != nil {
		return *obj.OAuth2ClientJsonSigningKeyResponse
	}

	// all schemas are nil
	return nil
}

type NullableListJwk200ResponseInner struct {
	value *ListJwk200ResponseInner
	isSet bool
}

func (v NullableListJwk200ResponseInner) Get() *ListJwk200ResponseInner {
	return v.value
}

func (v *NullableListJwk200ResponseInner) Set(val *ListJwk200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListJwk200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListJwk200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListJwk200ResponseInner(val *ListJwk200ResponseInner) *NullableListJwk200ResponseInner {
	return &NullableListJwk200ResponseInner{value: val, isSet: true}
}

func (v NullableListJwk200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListJwk200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
