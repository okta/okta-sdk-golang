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
	"reflect"
	"strings"
)

// checks if the ServiceAccountDetailsAppAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountDetailsAppAccount{}

// ServiceAccountDetailsAppAccount Details for a SaaS app account, which will be managed as a service account
type ServiceAccountDetailsAppAccount struct {
	ServiceAccount
	Details              ServiceAccountDetailsAppAccountSub `json:"details"`
	AdditionalProperties map[string]interface{}
}

type _ServiceAccountDetailsAppAccount ServiceAccountDetailsAppAccount

// NewServiceAccountDetailsAppAccount instantiates a new ServiceAccountDetailsAppAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountDetailsAppAccount(details ServiceAccountDetailsAppAccountSub, accountType string, name string) *ServiceAccountDetailsAppAccount {
	this := ServiceAccountDetailsAppAccount{}
	this.AccountType = accountType
	this.Name = name
	this.Details = details
	return &this
}

// NewServiceAccountDetailsAppAccountWithDefaults instantiates a new ServiceAccountDetailsAppAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountDetailsAppAccountWithDefaults() *ServiceAccountDetailsAppAccount {
	this := ServiceAccountDetailsAppAccount{}
	return &this
}

// GetDetails returns the Details field value
func (o *ServiceAccountDetailsAppAccount) GetDetails() ServiceAccountDetailsAppAccountSub {
	if o == nil {
		var ret ServiceAccountDetailsAppAccountSub
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetailsAppAccount) GetDetailsOk() (*ServiceAccountDetailsAppAccountSub, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *ServiceAccountDetailsAppAccount) SetDetails(v ServiceAccountDetailsAppAccountSub) {
	o.Details = v
}

func (o ServiceAccountDetailsAppAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountDetailsAppAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedServiceAccount, errServiceAccount := json.Marshal(o.ServiceAccount)
	if errServiceAccount != nil {
		return map[string]interface{}{}, errServiceAccount
	}
	errServiceAccount = json.Unmarshal([]byte(serializedServiceAccount), &toSerialize)
	if errServiceAccount != nil {
		return map[string]interface{}{}, errServiceAccount
	}
	toSerialize["details"] = o.Details

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceAccountDetailsAppAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"details",
		"accountType",
		"name",
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

	type ServiceAccountDetailsAppAccountWithoutEmbeddedStruct struct {
		Details ServiceAccountDetailsAppAccountSub `json:"details"`
	}

	varServiceAccountDetailsAppAccountWithoutEmbeddedStruct := ServiceAccountDetailsAppAccountWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varServiceAccountDetailsAppAccountWithoutEmbeddedStruct)
	if err == nil {
		varServiceAccountDetailsAppAccount := _ServiceAccountDetailsAppAccount{}
		varServiceAccountDetailsAppAccount.Details = varServiceAccountDetailsAppAccountWithoutEmbeddedStruct.Details
		*o = ServiceAccountDetailsAppAccount(varServiceAccountDetailsAppAccount)
	} else {
		return err
	}

	varServiceAccountDetailsAppAccount := _ServiceAccountDetailsAppAccount{}

	err = json.Unmarshal(data, &varServiceAccountDetailsAppAccount)
	if err == nil {
		o.ServiceAccount = varServiceAccountDetailsAppAccount.ServiceAccount
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "details")

		// remove fields from embedded structs
		reflectServiceAccount := reflect.ValueOf(o.ServiceAccount)
		for i := 0; i < reflectServiceAccount.Type().NumField(); i++ {
			t := reflectServiceAccount.Type().Field(i)

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

type NullableServiceAccountDetailsAppAccount struct {
	value *ServiceAccountDetailsAppAccount
	isSet bool
}

func (v NullableServiceAccountDetailsAppAccount) Get() *ServiceAccountDetailsAppAccount {
	return v.value
}

func (v *NullableServiceAccountDetailsAppAccount) Set(val *ServiceAccountDetailsAppAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountDetailsAppAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountDetailsAppAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountDetailsAppAccount(val *ServiceAccountDetailsAppAccount) *NullableServiceAccountDetailsAppAccount {
	return &NullableServiceAccountDetailsAppAccount{value: val, isSet: true}
}

func (v NullableServiceAccountDetailsAppAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountDetailsAppAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
