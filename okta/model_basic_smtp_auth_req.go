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
	"reflect"
	"strings"
)

// checks if the BASICSMTPAUTHREQ type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BASICSMTPAUTHREQ{}

// BASICSMTPAUTHREQ struct for BASICSMTPAUTHREQ
type BASICSMTPAUTHREQ struct {
	BaseEmailServer
	// The password of the user account that's used to sign in to your SMTP server
	Password             string `json:"password"`
	AdditionalProperties map[string]interface{}
}

type _BASICSMTPAUTHREQ BASICSMTPAUTHREQ

// NewBASICSMTPAUTHREQ instantiates a new BASICSMTPAUTHREQ object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBASICSMTPAUTHREQ(password string, alias string, authType string, enabled bool, host string, port int32, username string) *BASICSMTPAUTHREQ {
	this := BASICSMTPAUTHREQ{}
	this.Alias = alias
	this.AuthType = authType
	this.Enabled = enabled
	this.Host = host
	this.Port = port
	this.Username = username
	this.Password = password
	return &this
}

// NewBASICSMTPAUTHREQWithDefaults instantiates a new BASICSMTPAUTHREQ object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBASICSMTPAUTHREQWithDefaults() *BASICSMTPAUTHREQ {
	this := BASICSMTPAUTHREQ{}
	return &this
}

// GetPassword returns the Password field value
func (o *BASICSMTPAUTHREQ) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *BASICSMTPAUTHREQ) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *BASICSMTPAUTHREQ) SetPassword(v string) {
	o.Password = v
}

func (o BASICSMTPAUTHREQ) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BASICSMTPAUTHREQ) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseEmailServer, errBaseEmailServer := json.Marshal(o.BaseEmailServer)
	if errBaseEmailServer != nil {
		return map[string]interface{}{}, errBaseEmailServer
	}
	errBaseEmailServer = json.Unmarshal([]byte(serializedBaseEmailServer), &toSerialize)
	if errBaseEmailServer != nil {
		return map[string]interface{}{}, errBaseEmailServer
	}
	toSerialize["password"] = o.Password

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BASICSMTPAUTHREQ) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
		"alias",
		"authType",
		"enabled",
		"host",
		"port",
		"username",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	type BASICSMTPAUTHREQWithoutEmbeddedStruct struct {
		// The password of the user account that's used to sign in to your SMTP server
		Password string `json:"password"`
	}

	varBASICSMTPAUTHREQWithoutEmbeddedStruct := BASICSMTPAUTHREQWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBASICSMTPAUTHREQWithoutEmbeddedStruct)
	if err == nil {
		varBASICSMTPAUTHREQ := _BASICSMTPAUTHREQ{}
		varBASICSMTPAUTHREQ.Password = varBASICSMTPAUTHREQWithoutEmbeddedStruct.Password
		*o = BASICSMTPAUTHREQ(varBASICSMTPAUTHREQ)
	} else {
		return err
	}

	varBASICSMTPAUTHREQ := _BASICSMTPAUTHREQ{}

	err = json.Unmarshal(data, &varBASICSMTPAUTHREQ)
	if err == nil {
		o.BaseEmailServer = varBASICSMTPAUTHREQ.BaseEmailServer
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

type NullableBASICSMTPAUTHREQ struct {
	value *BASICSMTPAUTHREQ
	isSet bool
}

func (v NullableBASICSMTPAUTHREQ) Get() *BASICSMTPAUTHREQ {
	return v.value
}

func (v *NullableBASICSMTPAUTHREQ) Set(val *BASICSMTPAUTHREQ) {
	v.value = val
	v.isSet = true
}

func (v NullableBASICSMTPAUTHREQ) IsSet() bool {
	return v.isSet
}

func (v *NullableBASICSMTPAUTHREQ) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBASICSMTPAUTHREQ(val *BASICSMTPAUTHREQ) *NullableBASICSMTPAUTHREQ {
	return &NullableBASICSMTPAUTHREQ{value: val, isSet: true}
}

func (v NullableBASICSMTPAUTHREQ) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBASICSMTPAUTHREQ) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
