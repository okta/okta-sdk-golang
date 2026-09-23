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
	"reflect"
	"strings"
)

// checks if the BASICSMTPAUTH type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BASICSMTPAUTH{}

// BASICSMTPAUTH struct for BASICSMTPAUTH
type BASICSMTPAUTH struct {
	BaseEmailServer
	// The password of the user account that's used to sign in to your SMTP server
	Password             *string `json:"password,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BASICSMTPAUTH BASICSMTPAUTH

// NewBASICSMTPAUTH instantiates a new BASICSMTPAUTH object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBASICSMTPAUTH() *BASICSMTPAUTH {
	this := BASICSMTPAUTH{}
	return &this
}

// NewBASICSMTPAUTHWithDefaults instantiates a new BASICSMTPAUTH object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBASICSMTPAUTHWithDefaults() *BASICSMTPAUTH {
	this := BASICSMTPAUTH{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *BASICSMTPAUTH) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BASICSMTPAUTH) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *BASICSMTPAUTH) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *BASICSMTPAUTH) SetPassword(v string) {
	o.Password = &v
}

func (o BASICSMTPAUTH) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BASICSMTPAUTH) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseEmailServer, errBaseEmailServer := json.Marshal(o.BaseEmailServer)
	if errBaseEmailServer != nil {
		return map[string]interface{}{}, errBaseEmailServer
	}
	errBaseEmailServer = json.Unmarshal([]byte(serializedBaseEmailServer), &toSerialize)
	if errBaseEmailServer != nil {
		return map[string]interface{}{}, errBaseEmailServer
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BASICSMTPAUTH) UnmarshalJSON(data []byte) (err error) {
	type BASICSMTPAUTHWithoutEmbeddedStruct struct {
		// The password of the user account that's used to sign in to your SMTP server
		Password *string `json:"password,omitempty"`
	}

	varBASICSMTPAUTHWithoutEmbeddedStruct := BASICSMTPAUTHWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBASICSMTPAUTHWithoutEmbeddedStruct)
	if err == nil {
		varBASICSMTPAUTH := _BASICSMTPAUTH{}
		varBASICSMTPAUTH.Password = varBASICSMTPAUTHWithoutEmbeddedStruct.Password
		*o = BASICSMTPAUTH(varBASICSMTPAUTH)
	} else {
		return err
	}

	varBASICSMTPAUTH := _BASICSMTPAUTH{}

	err = json.Unmarshal(data, &varBASICSMTPAUTH)
	if err == nil {
		o.BaseEmailServer = varBASICSMTPAUTH.BaseEmailServer
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "password")

		// remove fields from embedded structs
		reflectBaseEmailServer := reflect.ValueOf(o.BaseEmailServer)
		for i := 0; i < reflectBaseEmailServer.Type().NumField(); i++ {
			t := reflectBaseEmailServer.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBASICSMTPAUTH struct {
	value *BASICSMTPAUTH
	isSet bool
}

func (v NullableBASICSMTPAUTH) Get() *BASICSMTPAUTH {
	return v.value
}

func (v *NullableBASICSMTPAUTH) Set(val *BASICSMTPAUTH) {
	v.value = val
	v.isSet = true
}

func (v NullableBASICSMTPAUTH) IsSet() bool {
	return v.isSet
}

func (v *NullableBASICSMTPAUTH) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBASICSMTPAUTH(val *BASICSMTPAUTH) *NullableBASICSMTPAUTH {
	return &NullableBASICSMTPAUTH{value: val, isSet: true}
}

func (v NullableBASICSMTPAUTH) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBASICSMTPAUTH) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
