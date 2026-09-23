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

// GetEmailServer200Response - struct for GetEmailServer200Response
type GetEmailServer200Response struct {
	BASICSMTPAUTH           *BASICSMTPAUTH
	OAUTH2CLIENTCREDENTIALS *OAUTH2CLIENTCREDENTIALS
	OAUTH2JWTBEARERGRANT    *OAUTH2JWTBEARERGRANT
}

// BASICSMTPAUTHAsGetEmailServer200Response is a convenience function that returns BASICSMTPAUTH wrapped in GetEmailServer200Response
func BASICSMTPAUTHAsGetEmailServer200Response(v *BASICSMTPAUTH) GetEmailServer200Response {
	return GetEmailServer200Response{
		BASICSMTPAUTH: v,
	}
}

// OAUTH2CLIENTCREDENTIALSAsGetEmailServer200Response is a convenience function that returns OAUTH2CLIENTCREDENTIALS wrapped in GetEmailServer200Response
func OAUTH2CLIENTCREDENTIALSAsGetEmailServer200Response(v *OAUTH2CLIENTCREDENTIALS) GetEmailServer200Response {
	return GetEmailServer200Response{
		OAUTH2CLIENTCREDENTIALS: v,
	}
}

// OAUTH2JWTBEARERGRANTAsGetEmailServer200Response is a convenience function that returns OAUTH2JWTBEARERGRANT wrapped in GetEmailServer200Response
func OAUTH2JWTBEARERGRANTAsGetEmailServer200Response(v *OAUTH2JWTBEARERGRANT) GetEmailServer200Response {
	return GetEmailServer200Response{
		OAUTH2JWTBEARERGRANT: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetEmailServer200Response) UnmarshalJSON(data []byte) error {
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
		// try to unmarshal JSON data into BASICSMTPAUTH
		err = json.Unmarshal(data, &dst.BASICSMTPAUTH)
		if err == nil {
			return nil // data stored in dst.BASICSMTPAUTH, return on the first match
		} else {
			dst.BASICSMTPAUTH = nil
			return fmt.Errorf("failed to unmarshal GetEmailServer200Response as BASICSMTPAUTH: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OAUTH2_CLIENT_CREDENTIALS'
	if discriminatorValue == "OAUTH2_CLIENT_CREDENTIALS" {
		// try to unmarshal JSON data into OAUTH2CLIENTCREDENTIALS
		err = json.Unmarshal(data, &dst.OAUTH2CLIENTCREDENTIALS)
		if err == nil {
			return nil // data stored in dst.OAUTH2CLIENTCREDENTIALS, return on the first match
		} else {
			dst.OAUTH2CLIENTCREDENTIALS = nil
			return fmt.Errorf("failed to unmarshal GetEmailServer200Response as OAUTH2CLIENTCREDENTIALS: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OAUTH2_JWT_BEARER_GRANT'
	if discriminatorValue == "OAUTH2_JWT_BEARER_GRANT" {
		// try to unmarshal JSON data into OAUTH2JWTBEARERGRANT
		err = json.Unmarshal(data, &dst.OAUTH2JWTBEARERGRANT)
		if err == nil {
			return nil // data stored in dst.OAUTH2JWTBEARERGRANT, return on the first match
		} else {
			dst.OAUTH2JWTBEARERGRANT = nil
			return fmt.Errorf("failed to unmarshal GetEmailServer200Response as OAUTH2JWTBEARERGRANT: %s", err.Error())
		}
	}

	// If discriminator value is empty/missing, default to the last mapped model (typically the most common type)
	if discriminatorValue == "" {
		err = json.Unmarshal(data, &dst.OAUTH2JWTBEARERGRANT)
		if err == nil {
			return nil
		}
		dst.OAUTH2JWTBEARERGRANT = nil
	}

	// No match found or unmarshal failed - return nil to allow partial unmarshalling
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetEmailServer200Response) MarshalJSON() ([]byte, error) {
	if src.BASICSMTPAUTH != nil {
		return json.Marshal(&src.BASICSMTPAUTH)
	}

	if src.OAUTH2CLIENTCREDENTIALS != nil {
		return json.Marshal(&src.OAUTH2CLIENTCREDENTIALS)
	}

	if src.OAUTH2JWTBEARERGRANT != nil {
		return json.Marshal(&src.OAUTH2JWTBEARERGRANT)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetEmailServer200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BASICSMTPAUTH != nil {
		return obj.BASICSMTPAUTH
	}

	if obj.OAUTH2CLIENTCREDENTIALS != nil {
		return obj.OAUTH2CLIENTCREDENTIALS
	}

	if obj.OAUTH2JWTBEARERGRANT != nil {
		return obj.OAUTH2JWTBEARERGRANT
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj GetEmailServer200Response) GetActualInstanceValue() interface{} {
	if obj.BASICSMTPAUTH != nil {
		return *obj.BASICSMTPAUTH
	}

	if obj.OAUTH2CLIENTCREDENTIALS != nil {
		return *obj.OAUTH2CLIENTCREDENTIALS
	}

	if obj.OAUTH2JWTBEARERGRANT != nil {
		return *obj.OAUTH2JWTBEARERGRANT
	}

	// all schemas are nil
	return nil
}

type NullableGetEmailServer200Response struct {
	value *GetEmailServer200Response
	isSet bool
}

func (v NullableGetEmailServer200Response) Get() *GetEmailServer200Response {
	return v.value
}

func (v *NullableGetEmailServer200Response) Set(val *GetEmailServer200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEmailServer200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEmailServer200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEmailServer200Response(val *GetEmailServer200Response) *NullableGetEmailServer200Response {
	return &NullableGetEmailServer200Response{value: val, isSet: true}
}

func (v NullableGetEmailServer200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEmailServer200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
