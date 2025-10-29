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
	"reflect"
	"strings"
)

// checks if the PrivilegedResourceAccountOkta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivilegedResourceAccountOkta{}

// PrivilegedResourceAccountOkta struct for PrivilegedResourceAccountOkta
type PrivilegedResourceAccountOkta struct {
	PrivilegedResource
	// The user ID associated with the Okta privileged resource
	ResourceId  *string                        `json:"resourceId,omitempty"`
	Credentials *PrivilegedResourceCredentials `json:"credentials,omitempty"`
	// Specific profile properties for the privileged resource
	Profile              map[string]interface{} `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrivilegedResourceAccountOkta PrivilegedResourceAccountOkta

// NewPrivilegedResourceAccountOkta instantiates a new PrivilegedResourceAccountOkta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegedResourceAccountOkta(resourceType string) *PrivilegedResourceAccountOkta {
	this := PrivilegedResourceAccountOkta{}
	this.ResourceType = resourceType
	return &this
}

// NewPrivilegedResourceAccountOktaWithDefaults instantiates a new PrivilegedResourceAccountOkta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegedResourceAccountOktaWithDefaults() *PrivilegedResourceAccountOkta {
	this := PrivilegedResourceAccountOkta{}
	return &this
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *PrivilegedResourceAccountOkta) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceAccountOkta) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *PrivilegedResourceAccountOkta) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *PrivilegedResourceAccountOkta) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *PrivilegedResourceAccountOkta) GetCredentials() PrivilegedResourceCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret PrivilegedResourceCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceAccountOkta) GetCredentialsOk() (*PrivilegedResourceCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *PrivilegedResourceAccountOkta) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given PrivilegedResourceCredentials and assigns it to the Credentials field.
func (o *PrivilegedResourceAccountOkta) SetCredentials(v PrivilegedResourceCredentials) {
	o.Credentials = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *PrivilegedResourceAccountOkta) GetProfile() map[string]interface{} {
	if o == nil || IsNil(o.Profile) {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceAccountOkta) GetProfileOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Profile) {
		return map[string]interface{}{}, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *PrivilegedResourceAccountOkta) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]interface{} and assigns it to the Profile field.
func (o *PrivilegedResourceAccountOkta) SetProfile(v map[string]interface{}) {
	o.Profile = v
}

func (o PrivilegedResourceAccountOkta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivilegedResourceAccountOkta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPrivilegedResource, errPrivilegedResource := json.Marshal(o.PrivilegedResource)
	if errPrivilegedResource != nil {
		return map[string]interface{}{}, errPrivilegedResource
	}
	errPrivilegedResource = json.Unmarshal([]byte(serializedPrivilegedResource), &toSerialize)
	if errPrivilegedResource != nil {
		return map[string]interface{}{}, errPrivilegedResource
	}
	if !IsNil(o.ResourceId) {
		toSerialize["resourceId"] = o.ResourceId
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

func (o *PrivilegedResourceAccountOkta) UnmarshalJSON(data []byte) (err error) {
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

	type PrivilegedResourceAccountOktaWithoutEmbeddedStruct struct {
		// The user ID associated with the Okta privileged resource
		ResourceId  *string                        `json:"resourceId,omitempty"`
		Credentials *PrivilegedResourceCredentials `json:"credentials,omitempty"`
		// Specific profile properties for the privileged resource
		Profile map[string]interface{} `json:"profile,omitempty"`
	}

	varPrivilegedResourceAccountOktaWithoutEmbeddedStruct := PrivilegedResourceAccountOktaWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPrivilegedResourceAccountOktaWithoutEmbeddedStruct)
	if err == nil {
		varPrivilegedResourceAccountOkta := _PrivilegedResourceAccountOkta{}
		varPrivilegedResourceAccountOkta.ResourceId = varPrivilegedResourceAccountOktaWithoutEmbeddedStruct.ResourceId
		varPrivilegedResourceAccountOkta.Credentials = varPrivilegedResourceAccountOktaWithoutEmbeddedStruct.Credentials
		varPrivilegedResourceAccountOkta.Profile = varPrivilegedResourceAccountOktaWithoutEmbeddedStruct.Profile
		*o = PrivilegedResourceAccountOkta(varPrivilegedResourceAccountOkta)
	} else {
		return err
	}

	varPrivilegedResourceAccountOkta := _PrivilegedResourceAccountOkta{}

	err = json.Unmarshal(data, &varPrivilegedResourceAccountOkta)
	if err == nil {
		o.PrivilegedResource = varPrivilegedResourceAccountOkta.PrivilegedResource
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceId")
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

type NullablePrivilegedResourceAccountOkta struct {
	value *PrivilegedResourceAccountOkta
	isSet bool
}

func (v NullablePrivilegedResourceAccountOkta) Get() *PrivilegedResourceAccountOkta {
	return v.value
}

func (v *NullablePrivilegedResourceAccountOkta) Set(val *PrivilegedResourceAccountOkta) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegedResourceAccountOkta) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegedResourceAccountOkta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegedResourceAccountOkta(val *PrivilegedResourceAccountOkta) *NullablePrivilegedResourceAccountOkta {
	return &NullablePrivilegedResourceAccountOkta{value: val, isSet: true}
}

func (v NullablePrivilegedResourceAccountOkta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegedResourceAccountOkta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
