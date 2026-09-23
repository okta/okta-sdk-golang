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

// checks if the BASICSMTPAUTHCREATE type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BASICSMTPAUTHCREATE{}

// BASICSMTPAUTHCREATE struct for BASICSMTPAUTHCREATE
type BASICSMTPAUTHCREATE struct {
	BaseEmailServerCreate
	// The password of the user account that's used to sign in to your SMTP server
	Password             string `json:"password"`
	AdditionalProperties map[string]interface{}
}

type _BASICSMTPAUTHCREATE BASICSMTPAUTHCREATE

// NewBASICSMTPAUTHCREATE instantiates a new BASICSMTPAUTHCREATE object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBASICSMTPAUTHCREATE(password string, alias string, enabled bool, host string, port int32, username string) *BASICSMTPAUTHCREATE {
	this := BASICSMTPAUTHCREATE{}
	this.Alias = alias
	this.Enabled = enabled
	this.Host = host
	this.Port = port
	this.Username = username
	this.Password = password
	return &this
}

// NewBASICSMTPAUTHCREATEWithDefaults instantiates a new BASICSMTPAUTHCREATE object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBASICSMTPAUTHCREATEWithDefaults() *BASICSMTPAUTHCREATE {
	this := BASICSMTPAUTHCREATE{}
	return &this
}

// GetPassword returns the Password field value
func (o *BASICSMTPAUTHCREATE) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *BASICSMTPAUTHCREATE) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *BASICSMTPAUTHCREATE) SetPassword(v string) {
	o.Password = v
}

func (o BASICSMTPAUTHCREATE) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BASICSMTPAUTHCREATE) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseEmailServerCreate, errBaseEmailServerCreate := json.Marshal(o.BaseEmailServerCreate)
	if errBaseEmailServerCreate != nil {
		return map[string]interface{}{}, errBaseEmailServerCreate
	}
	errBaseEmailServerCreate = json.Unmarshal([]byte(serializedBaseEmailServerCreate), &toSerialize)
	if errBaseEmailServerCreate != nil {
		return map[string]interface{}{}, errBaseEmailServerCreate
	}
	toSerialize["password"] = o.Password

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BASICSMTPAUTHCREATE) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
		"alias",
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

	type BASICSMTPAUTHCREATEWithoutEmbeddedStruct struct {
		// The password of the user account that's used to sign in to your SMTP server
		Password string `json:"password"`
	}

	varBASICSMTPAUTHCREATEWithoutEmbeddedStruct := BASICSMTPAUTHCREATEWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBASICSMTPAUTHCREATEWithoutEmbeddedStruct)
	if err == nil {
		varBASICSMTPAUTHCREATE := _BASICSMTPAUTHCREATE{}
		varBASICSMTPAUTHCREATE.Password = varBASICSMTPAUTHCREATEWithoutEmbeddedStruct.Password
		*o = BASICSMTPAUTHCREATE(varBASICSMTPAUTHCREATE)
	} else {
		return err
	}

	varBASICSMTPAUTHCREATE := _BASICSMTPAUTHCREATE{}

	err = json.Unmarshal(data, &varBASICSMTPAUTHCREATE)
	if err == nil {
		o.BaseEmailServerCreate = varBASICSMTPAUTHCREATE.BaseEmailServerCreate
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "password")

		// remove fields from embedded structs
		reflectBaseEmailServerCreate := reflect.ValueOf(o.BaseEmailServerCreate)
		for i := 0; i < reflectBaseEmailServerCreate.Type().NumField(); i++ {
			t := reflectBaseEmailServerCreate.Type().Field(i)

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

type NullableBASICSMTPAUTHCREATE struct {
	value *BASICSMTPAUTHCREATE
	isSet bool
}

func (v NullableBASICSMTPAUTHCREATE) Get() *BASICSMTPAUTHCREATE {
	return v.value
}

func (v *NullableBASICSMTPAUTHCREATE) Set(val *BASICSMTPAUTHCREATE) {
	v.value = val
	v.isSet = true
}

func (v NullableBASICSMTPAUTHCREATE) IsSet() bool {
	return v.isSet
}

func (v *NullableBASICSMTPAUTHCREATE) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBASICSMTPAUTHCREATE(val *BASICSMTPAUTHCREATE) *NullableBASICSMTPAUTHCREATE {
	return &NullableBASICSMTPAUTHCREATE{value: val, isSet: true}
}

func (v NullableBASICSMTPAUTHCREATE) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBASICSMTPAUTHCREATE) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
