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

// PrivilegedResourceAccountAppRequest struct for PrivilegedResourceAccountAppRequest
type PrivilegedResourceAccountAppRequest struct {
	PrivilegedResource
	ContainerDetails *AppAccountContainerDetails `json:"containerDetails,omitempty"`
	Credentials *PrivilegedResourceCredentials `json:"credentials,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrivilegedResourceAccountAppRequest PrivilegedResourceAccountAppRequest

// NewPrivilegedResourceAccountAppRequest instantiates a new PrivilegedResourceAccountAppRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegedResourceAccountAppRequest(resourceType string) *PrivilegedResourceAccountAppRequest {
	this := PrivilegedResourceAccountAppRequest{}
	this.ResourceType = resourceType
	return &this
}

// NewPrivilegedResourceAccountAppRequestWithDefaults instantiates a new PrivilegedResourceAccountAppRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegedResourceAccountAppRequestWithDefaults() *PrivilegedResourceAccountAppRequest {
	this := PrivilegedResourceAccountAppRequest{}
	return &this
}

// GetContainerDetails returns the ContainerDetails field value if set, zero value otherwise.
func (o *PrivilegedResourceAccountAppRequest) GetContainerDetails() AppAccountContainerDetails {
	if o == nil || o.ContainerDetails == nil {
		var ret AppAccountContainerDetails
		return ret
	}
	return *o.ContainerDetails
}

// GetContainerDetailsOk returns a tuple with the ContainerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceAccountAppRequest) GetContainerDetailsOk() (*AppAccountContainerDetails, bool) {
	if o == nil || o.ContainerDetails == nil {
		return nil, false
	}
	return o.ContainerDetails, true
}

// HasContainerDetails returns a boolean if a field has been set.
func (o *PrivilegedResourceAccountAppRequest) HasContainerDetails() bool {
	if o != nil && o.ContainerDetails != nil {
		return true
	}

	return false
}

// SetContainerDetails gets a reference to the given AppAccountContainerDetails and assigns it to the ContainerDetails field.
func (o *PrivilegedResourceAccountAppRequest) SetContainerDetails(v AppAccountContainerDetails) {
	o.ContainerDetails = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *PrivilegedResourceAccountAppRequest) GetCredentials() PrivilegedResourceCredentials {
	if o == nil || o.Credentials == nil {
		var ret PrivilegedResourceCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceAccountAppRequest) GetCredentialsOk() (*PrivilegedResourceCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *PrivilegedResourceAccountAppRequest) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given PrivilegedResourceCredentials and assigns it to the Credentials field.
func (o *PrivilegedResourceAccountAppRequest) SetCredentials(v PrivilegedResourceCredentials) {
	o.Credentials = &v
}

func (o PrivilegedResourceAccountAppRequest) MarshalJSON() ([]byte, error) {
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
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PrivilegedResourceAccountAppRequest) UnmarshalJSON(bytes []byte) (err error) {
	type PrivilegedResourceAccountAppRequestWithoutEmbeddedStruct struct {
		ContainerDetails *AppAccountContainerDetails `json:"containerDetails,omitempty"`
		Credentials *PrivilegedResourceCredentials `json:"credentials,omitempty"`
	}

	varPrivilegedResourceAccountAppRequestWithoutEmbeddedStruct := PrivilegedResourceAccountAppRequestWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPrivilegedResourceAccountAppRequestWithoutEmbeddedStruct)
	if err == nil {
		varPrivilegedResourceAccountAppRequest := _PrivilegedResourceAccountAppRequest{}
		varPrivilegedResourceAccountAppRequest.ContainerDetails = varPrivilegedResourceAccountAppRequestWithoutEmbeddedStruct.ContainerDetails
		varPrivilegedResourceAccountAppRequest.Credentials = varPrivilegedResourceAccountAppRequestWithoutEmbeddedStruct.Credentials
		*o = PrivilegedResourceAccountAppRequest(varPrivilegedResourceAccountAppRequest)
	} else {
		return err
	}

	varPrivilegedResourceAccountAppRequest := _PrivilegedResourceAccountAppRequest{}

	err = json.Unmarshal(bytes, &varPrivilegedResourceAccountAppRequest)
	if err == nil {
		o.PrivilegedResource = varPrivilegedResourceAccountAppRequest.PrivilegedResource
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

type NullablePrivilegedResourceAccountAppRequest struct {
	value *PrivilegedResourceAccountAppRequest
	isSet bool
}

func (v NullablePrivilegedResourceAccountAppRequest) Get() *PrivilegedResourceAccountAppRequest {
	return v.value
}

func (v *NullablePrivilegedResourceAccountAppRequest) Set(val *PrivilegedResourceAccountAppRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegedResourceAccountAppRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegedResourceAccountAppRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegedResourceAccountAppRequest(val *PrivilegedResourceAccountAppRequest) *NullablePrivilegedResourceAccountAppRequest {
	return &NullablePrivilegedResourceAccountAppRequest{value: val, isSet: true}
}

func (v NullablePrivilegedResourceAccountAppRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegedResourceAccountAppRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

