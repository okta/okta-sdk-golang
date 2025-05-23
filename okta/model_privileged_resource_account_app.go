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
	"reflect"
	"strings"
)

// PrivilegedResourceAccountApp struct for PrivilegedResourceAccountApp
type PrivilegedResourceAccountApp struct {
	PrivilegedResource
	ContainerDetails *AppAccountContainerDetails `json:"containerDetails,omitempty"`
	Credentials PrivilegedResourceCredentials `json:"credentials"`
	AdditionalProperties map[string]interface{}
}

type _PrivilegedResourceAccountApp PrivilegedResourceAccountApp

// NewPrivilegedResourceAccountApp instantiates a new PrivilegedResourceAccountApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegedResourceAccountApp(credentials PrivilegedResourceCredentials) *PrivilegedResourceAccountApp {
	this := PrivilegedResourceAccountApp{}
	this.Credentials = credentials
	return &this
}

// NewPrivilegedResourceAccountAppWithDefaults instantiates a new PrivilegedResourceAccountApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegedResourceAccountAppWithDefaults() *PrivilegedResourceAccountApp {
	this := PrivilegedResourceAccountApp{}
	return &this
}

// GetContainerDetails returns the ContainerDetails field value if set, zero value otherwise.
func (o *PrivilegedResourceAccountApp) GetContainerDetails() AppAccountContainerDetails {
	if o == nil || o.ContainerDetails == nil {
		var ret AppAccountContainerDetails
		return ret
	}
	return *o.ContainerDetails
}

// GetContainerDetailsOk returns a tuple with the ContainerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceAccountApp) GetContainerDetailsOk() (*AppAccountContainerDetails, bool) {
	if o == nil || o.ContainerDetails == nil {
		return nil, false
	}
	return o.ContainerDetails, true
}

// HasContainerDetails returns a boolean if a field has been set.
func (o *PrivilegedResourceAccountApp) HasContainerDetails() bool {
	if o != nil && o.ContainerDetails != nil {
		return true
	}

	return false
}

// SetContainerDetails gets a reference to the given AppAccountContainerDetails and assigns it to the ContainerDetails field.
func (o *PrivilegedResourceAccountApp) SetContainerDetails(v AppAccountContainerDetails) {
	o.ContainerDetails = &v
}

// GetCredentials returns the Credentials field value
func (o *PrivilegedResourceAccountApp) GetCredentials() PrivilegedResourceCredentials {
	if o == nil {
		var ret PrivilegedResourceCredentials
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceAccountApp) GetCredentialsOk() (*PrivilegedResourceCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *PrivilegedResourceAccountApp) SetCredentials(v PrivilegedResourceCredentials) {
	o.Credentials = v
}

func (o PrivilegedResourceAccountApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPrivilegedResource, errPrivilegedResource := json.Marshal(o.PrivilegedResource)
	if errPrivilegedResource != nil {
		return []byte{}, errPrivilegedResource
	}
	errPrivilegedResource = json.Unmarshal([]byte(serializedPrivilegedResource), &toSerialize)
	if errPrivilegedResource != nil {
		return []byte{}, errPrivilegedResource
	}
	if o.ContainerDetails != nil {
		toSerialize["containerDetails"] = o.ContainerDetails
	}
	if true {
		toSerialize["credentials"] = o.Credentials
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PrivilegedResourceAccountApp) UnmarshalJSON(bytes []byte) (err error) {
	type PrivilegedResourceAccountAppWithoutEmbeddedStruct struct {
		ContainerDetails *AppAccountContainerDetails `json:"containerDetails,omitempty"`
		Credentials PrivilegedResourceCredentials `json:"credentials"`
	}

	varPrivilegedResourceAccountAppWithoutEmbeddedStruct := PrivilegedResourceAccountAppWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPrivilegedResourceAccountAppWithoutEmbeddedStruct)
	if err == nil {
		varPrivilegedResourceAccountApp := _PrivilegedResourceAccountApp{}
		varPrivilegedResourceAccountApp.ContainerDetails = varPrivilegedResourceAccountAppWithoutEmbeddedStruct.ContainerDetails
		varPrivilegedResourceAccountApp.Credentials = varPrivilegedResourceAccountAppWithoutEmbeddedStruct.Credentials
		*o = PrivilegedResourceAccountApp(varPrivilegedResourceAccountApp)
	} else {
		return err
	}

	varPrivilegedResourceAccountApp := _PrivilegedResourceAccountApp{}

	err = json.Unmarshal(bytes, &varPrivilegedResourceAccountApp)
	if err == nil {
		o.PrivilegedResource = varPrivilegedResourceAccountApp.PrivilegedResource
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "containerDetails")
		delete(additionalProperties, "credentials")

		// remove fields from embedded structs
		reflectPrivilegedResource := reflect.ValueOf(o.PrivilegedResource)
		for i := 0; i < reflectPrivilegedResource.Type().NumField(); i++ {
			t := reflectPrivilegedResource.Type().Field(i)

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
	} else {
		return err
	}

	return err
}

type NullablePrivilegedResourceAccountApp struct {
	value *PrivilegedResourceAccountApp
	isSet bool
}

func (v NullablePrivilegedResourceAccountApp) Get() *PrivilegedResourceAccountApp {
	return v.value
}

func (v *NullablePrivilegedResourceAccountApp) Set(val *PrivilegedResourceAccountApp) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegedResourceAccountApp) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegedResourceAccountApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegedResourceAccountApp(val *PrivilegedResourceAccountApp) *NullablePrivilegedResourceAccountApp {
	return &NullablePrivilegedResourceAccountApp{value: val, isSet: true}
}

func (v NullablePrivilegedResourceAccountApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegedResourceAccountApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

