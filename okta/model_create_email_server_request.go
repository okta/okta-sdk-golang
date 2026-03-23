/*
Okta Admin Management API

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

// CreateEmailServerRequest - struct for CreateEmailServerRequest
type CreateEmailServerRequest struct {
	BASICSMTPAUTHREQ           *BASICSMTPAUTHREQ
	OAUTH2CLIENTCREDENTIALSREQ *OAUTH2CLIENTCREDENTIALSREQ
	OAUTH2JWTBEARERGRANTREQ    *OAUTH2JWTBEARERGRANTREQ
}

// BASICSMTPAUTHREQAsCreateEmailServerRequest is a convenience function that returns BASICSMTPAUTHREQ wrapped in CreateEmailServerRequest
func BASICSMTPAUTHREQAsCreateEmailServerRequest(v *BASICSMTPAUTHREQ) CreateEmailServerRequest {
	return CreateEmailServerRequest{
		BASICSMTPAUTHREQ: v,
	}
}

// OAUTH2CLIENTCREDENTIALSREQAsCreateEmailServerRequest is a convenience function that returns OAUTH2CLIENTCREDENTIALSREQ wrapped in CreateEmailServerRequest
func OAUTH2CLIENTCREDENTIALSREQAsCreateEmailServerRequest(v *OAUTH2CLIENTCREDENTIALSREQ) CreateEmailServerRequest {
	return CreateEmailServerRequest{
		OAUTH2CLIENTCREDENTIALSREQ: v,
	}
}

// OAUTH2JWTBEARERGRANTREQAsCreateEmailServerRequest is a convenience function that returns OAUTH2JWTBEARERGRANTREQ wrapped in CreateEmailServerRequest
func OAUTH2JWTBEARERGRANTREQAsCreateEmailServerRequest(v *OAUTH2JWTBEARERGRANTREQ) CreateEmailServerRequest {
	return CreateEmailServerRequest{
		OAUTH2JWTBEARERGRANTREQ: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateEmailServerRequest) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// Get discriminator value, treating nil/missing as empty string for comparison
	discriminatorValue, _ := jsonDict["authType"].(string)

	// check if the discriminator value is 'BASIC_SMTP_AUTH'
	if discriminatorValue == "BASIC_SMTP_AUTH" {
		// try to unmarshal JSON data into BASICSMTPAUTHREQ
		err = json.Unmarshal(data, &dst.BASICSMTPAUTHREQ)
		if err == nil {
			return nil // data stored in dst.BASICSMTPAUTHREQ, return on the first match
		} else {
			dst.BASICSMTPAUTHREQ = nil
			return fmt.Errorf("failed to unmarshal CreateEmailServerRequest as BASICSMTPAUTHREQ: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OAUTH2_CLIENT_CREDENTIALS'
	if discriminatorValue == "OAUTH2_CLIENT_CREDENTIALS" {
		// try to unmarshal JSON data into OAUTH2CLIENTCREDENTIALSREQ
		err = json.Unmarshal(data, &dst.OAUTH2CLIENTCREDENTIALSREQ)
		if err == nil {
			return nil // data stored in dst.OAUTH2CLIENTCREDENTIALSREQ, return on the first match
		} else {
			dst.OAUTH2CLIENTCREDENTIALSREQ = nil
			return fmt.Errorf("failed to unmarshal CreateEmailServerRequest as OAUTH2CLIENTCREDENTIALSREQ: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OAUTH2_JWT_BEARER_GRANT'
	if discriminatorValue == "OAUTH2_JWT_BEARER_GRANT" {
		// try to unmarshal JSON data into OAUTH2JWTBEARERGRANTREQ
		err = json.Unmarshal(data, &dst.OAUTH2JWTBEARERGRANTREQ)
		if err == nil {
			return nil // data stored in dst.OAUTH2JWTBEARERGRANTREQ, return on the first match
		} else {
			dst.OAUTH2JWTBEARERGRANTREQ = nil
			return fmt.Errorf("failed to unmarshal CreateEmailServerRequest as OAUTH2JWTBEARERGRANTREQ: %s", err.Error())
		}
	}

	// If discriminator value is empty/missing, default to the last mapped model (typically the most common type)
	if discriminatorValue == "" {
		err = json.Unmarshal(data, &dst.OAUTH2JWTBEARERGRANTREQ)
		if err == nil {
			return nil
		}
		dst.OAUTH2JWTBEARERGRANTREQ = nil
	}

	// No match found or unmarshal failed - return nil to allow partial unmarshalling
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateEmailServerRequest) MarshalJSON() ([]byte, error) {
	if src.BASICSMTPAUTHREQ != nil {
		return json.Marshal(&src.BASICSMTPAUTHREQ)
	}

	if src.OAUTH2CLIENTCREDENTIALSREQ != nil {
		return json.Marshal(&src.OAUTH2CLIENTCREDENTIALSREQ)
	}

	if src.OAUTH2JWTBEARERGRANTREQ != nil {
		return json.Marshal(&src.OAUTH2JWTBEARERGRANTREQ)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateEmailServerRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BASICSMTPAUTHREQ != nil {
		return obj.BASICSMTPAUTHREQ
	}

	if obj.OAUTH2CLIENTCREDENTIALSREQ != nil {
		return obj.OAUTH2CLIENTCREDENTIALSREQ
	}

	if obj.OAUTH2JWTBEARERGRANTREQ != nil {
		return obj.OAUTH2JWTBEARERGRANTREQ
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj CreateEmailServerRequest) GetActualInstanceValue() interface{} {
	if obj.BASICSMTPAUTHREQ != nil {
		return *obj.BASICSMTPAUTHREQ
	}

	if obj.OAUTH2CLIENTCREDENTIALSREQ != nil {
		return *obj.OAUTH2CLIENTCREDENTIALSREQ
	}

	if obj.OAUTH2JWTBEARERGRANTREQ != nil {
		return *obj.OAUTH2JWTBEARERGRANTREQ
	}

	// all schemas are nil
	return nil
}

type NullableCreateEmailServerRequest struct {
	value *CreateEmailServerRequest
	isSet bool
}

func (v NullableCreateEmailServerRequest) Get() *CreateEmailServerRequest {
	return v.value
}

func (v *NullableCreateEmailServerRequest) Set(val *CreateEmailServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEmailServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEmailServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEmailServerRequest(val *CreateEmailServerRequest) *NullableCreateEmailServerRequest {
	return &NullableCreateEmailServerRequest{value: val, isSet: true}
}

func (v NullableCreateEmailServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEmailServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
