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
)

// checks if the IdentityAssertionCustomASConnectionCreatableAuthorizationServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityAssertionCustomASConnectionCreatableAuthorizationServer{}

// IdentityAssertionCustomASConnectionCreatableAuthorizationServer Reference to a custom authorization server and its configuration
type IdentityAssertionCustomASConnectionCreatableAuthorizationServer struct {
	// The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the authorization server
	Orn string `json:"orn"`
	// Optional resource indicator (audience) used when requesting tokens
	ResourceIndicator    *string `json:"resourceIndicator,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityAssertionCustomASConnectionCreatableAuthorizationServer IdentityAssertionCustomASConnectionCreatableAuthorizationServer

// NewIdentityAssertionCustomASConnectionCreatableAuthorizationServer instantiates a new IdentityAssertionCustomASConnectionCreatableAuthorizationServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityAssertionCustomASConnectionCreatableAuthorizationServer(orn string) *IdentityAssertionCustomASConnectionCreatableAuthorizationServer {
	this := IdentityAssertionCustomASConnectionCreatableAuthorizationServer{}
	this.Orn = orn
	return &this
}

// NewIdentityAssertionCustomASConnectionCreatableAuthorizationServerWithDefaults instantiates a new IdentityAssertionCustomASConnectionCreatableAuthorizationServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityAssertionCustomASConnectionCreatableAuthorizationServerWithDefaults() *IdentityAssertionCustomASConnectionCreatableAuthorizationServer {
	this := IdentityAssertionCustomASConnectionCreatableAuthorizationServer{}
	return &this
}

// GetOrn returns the Orn field value
func (o *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) GetOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Orn
}

// GetOrnOk returns a tuple with the Orn field value
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) GetOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orn, true
}

// SetOrn sets field value
func (o *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) SetOrn(v string) {
	o.Orn = v
}

// GetResourceIndicator returns the ResourceIndicator field value if set, zero value otherwise.
func (o *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) GetResourceIndicator() string {
	if o == nil || IsNil(o.ResourceIndicator) {
		var ret string
		return ret
	}
	return *o.ResourceIndicator
}

// GetResourceIndicatorOk returns a tuple with the ResourceIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) GetResourceIndicatorOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceIndicator) {
		return nil, false
	}
	return o.ResourceIndicator, true
}

// HasResourceIndicator returns a boolean if a field has been set.
func (o *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) HasResourceIndicator() bool {
	if o != nil && !IsNil(o.ResourceIndicator) {
		return true
	}

	return false
}

// SetResourceIndicator gets a reference to the given string and assigns it to the ResourceIndicator field.
func (o *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) SetResourceIndicator(v string) {
	o.ResourceIndicator = &v
}

func (o IdentityAssertionCustomASConnectionCreatableAuthorizationServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityAssertionCustomASConnectionCreatableAuthorizationServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orn"] = o.Orn
	if !IsNil(o.ResourceIndicator) {
		toSerialize["resourceIndicator"] = o.ResourceIndicator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"orn",
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

	varIdentityAssertionCustomASConnectionCreatableAuthorizationServer := _IdentityAssertionCustomASConnectionCreatableAuthorizationServer{}

	err = json.Unmarshal(data, &varIdentityAssertionCustomASConnectionCreatableAuthorizationServer)

	if err != nil {
		return err
	}

	*o = IdentityAssertionCustomASConnectionCreatableAuthorizationServer(varIdentityAssertionCustomASConnectionCreatableAuthorizationServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orn")
		delete(additionalProperties, "resourceIndicator")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityAssertionCustomASConnectionCreatableAuthorizationServer struct {
	value *IdentityAssertionCustomASConnectionCreatableAuthorizationServer
	isSet bool
}

func (v NullableIdentityAssertionCustomASConnectionCreatableAuthorizationServer) Get() *IdentityAssertionCustomASConnectionCreatableAuthorizationServer {
	return v.value
}

func (v *NullableIdentityAssertionCustomASConnectionCreatableAuthorizationServer) Set(val *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAssertionCustomASConnectionCreatableAuthorizationServer) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAssertionCustomASConnectionCreatableAuthorizationServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAssertionCustomASConnectionCreatableAuthorizationServer(val *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) *NullableIdentityAssertionCustomASConnectionCreatableAuthorizationServer {
	return &NullableIdentityAssertionCustomASConnectionCreatableAuthorizationServer{value: val, isSet: true}
}

func (v NullableIdentityAssertionCustomASConnectionCreatableAuthorizationServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAssertionCustomASConnectionCreatableAuthorizationServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
