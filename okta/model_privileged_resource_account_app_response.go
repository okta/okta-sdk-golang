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

// checks if the PrivilegedResourceAccountAppResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivilegedResourceAccountAppResponse{}

// PrivilegedResourceAccountAppResponse struct for PrivilegedResourceAccountAppResponse
type PrivilegedResourceAccountAppResponse struct {
	PrivilegedResource
	Credentials *PrivilegedResourceCredentials `json:"credentials,omitempty"`
	// Specific profile properties for the privileged resource
	Profile              map[string]interface{} `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrivilegedResourceAccountAppResponse PrivilegedResourceAccountAppResponse

// NewPrivilegedResourceAccountAppResponse instantiates a new PrivilegedResourceAccountAppResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegedResourceAccountAppResponse(resourceType string) *PrivilegedResourceAccountAppResponse {
	this := PrivilegedResourceAccountAppResponse{}
	this.ResourceType = resourceType
	return &this
}

// NewPrivilegedResourceAccountAppResponseWithDefaults instantiates a new PrivilegedResourceAccountAppResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegedResourceAccountAppResponseWithDefaults() *PrivilegedResourceAccountAppResponse {
	this := PrivilegedResourceAccountAppResponse{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *PrivilegedResourceAccountAppResponse) GetCredentials() PrivilegedResourceCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret PrivilegedResourceCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceAccountAppResponse) GetCredentialsOk() (*PrivilegedResourceCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *PrivilegedResourceAccountAppResponse) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given PrivilegedResourceCredentials and assigns it to the Credentials field.
func (o *PrivilegedResourceAccountAppResponse) SetCredentials(v PrivilegedResourceCredentials) {
	o.Credentials = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *PrivilegedResourceAccountAppResponse) GetProfile() map[string]interface{} {
	if o == nil || IsNil(o.Profile) {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceAccountAppResponse) GetProfileOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Profile) {
		return map[string]interface{}{}, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *PrivilegedResourceAccountAppResponse) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]interface{} and assigns it to the Profile field.
func (o *PrivilegedResourceAccountAppResponse) SetProfile(v map[string]interface{}) {
	o.Profile = v
}

func (o PrivilegedResourceAccountAppResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivilegedResourceAccountAppResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPrivilegedResource, errPrivilegedResource := json.Marshal(o.PrivilegedResource)
	if errPrivilegedResource != nil {
		return map[string]interface{}{}, errPrivilegedResource
	}
	errPrivilegedResource = json.Unmarshal([]byte(serializedPrivilegedResource), &toSerialize)
	if errPrivilegedResource != nil {
		return map[string]interface{}{}, errPrivilegedResource
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrivilegedResourceAccountAppResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
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

	type PrivilegedResourceAccountAppResponseWithoutEmbeddedStruct struct {
		Credentials *PrivilegedResourceCredentials `json:"credentials,omitempty"`
		// Specific profile properties for the privileged resource
		Profile map[string]interface{} `json:"profile,omitempty"`
	}

	varPrivilegedResourceAccountAppResponseWithoutEmbeddedStruct := PrivilegedResourceAccountAppResponseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPrivilegedResourceAccountAppResponseWithoutEmbeddedStruct)
	if err == nil {
		varPrivilegedResourceAccountAppResponse := _PrivilegedResourceAccountAppResponse{}
		varPrivilegedResourceAccountAppResponse.Credentials = varPrivilegedResourceAccountAppResponseWithoutEmbeddedStruct.Credentials
		varPrivilegedResourceAccountAppResponse.Profile = varPrivilegedResourceAccountAppResponseWithoutEmbeddedStruct.Profile
		*o = PrivilegedResourceAccountAppResponse(varPrivilegedResourceAccountAppResponse)
	} else {
		return err
	}

	varPrivilegedResourceAccountAppResponse := _PrivilegedResourceAccountAppResponse{}

	err = json.Unmarshal(data, &varPrivilegedResourceAccountAppResponse)
	if err == nil {
		o.PrivilegedResource = varPrivilegedResourceAccountAppResponse.PrivilegedResource
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "profile")

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
	}

	return err
}

type NullablePrivilegedResourceAccountAppResponse struct {
	value *PrivilegedResourceAccountAppResponse
	isSet bool
}

func (v NullablePrivilegedResourceAccountAppResponse) Get() *PrivilegedResourceAccountAppResponse {
	return v.value
}

func (v *NullablePrivilegedResourceAccountAppResponse) Set(val *PrivilegedResourceAccountAppResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegedResourceAccountAppResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegedResourceAccountAppResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegedResourceAccountAppResponse(val *PrivilegedResourceAccountAppResponse) *NullablePrivilegedResourceAccountAppResponse {
	return &NullablePrivilegedResourceAccountAppResponse{value: val, isSet: true}
}

func (v NullablePrivilegedResourceAccountAppResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegedResourceAccountAppResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
