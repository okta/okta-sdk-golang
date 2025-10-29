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

// checks if the ManagedConnectionAuthorizationServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedConnectionAuthorizationServer{}

// ManagedConnectionAuthorizationServer struct for ManagedConnectionAuthorizationServer
type ManagedConnectionAuthorizationServer struct {
	// Issuer URL for the authorization server
	IssuerUrl string `json:"issuerUrl"`
	// Audience value to use when requesting tokens
	ResourceIndicator    *string `json:"resourceIndicator,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManagedConnectionAuthorizationServer ManagedConnectionAuthorizationServer

// NewManagedConnectionAuthorizationServer instantiates a new ManagedConnectionAuthorizationServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedConnectionAuthorizationServer(issuerUrl string) *ManagedConnectionAuthorizationServer {
	this := ManagedConnectionAuthorizationServer{}
	this.IssuerUrl = issuerUrl
	return &this
}

// NewManagedConnectionAuthorizationServerWithDefaults instantiates a new ManagedConnectionAuthorizationServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedConnectionAuthorizationServerWithDefaults() *ManagedConnectionAuthorizationServer {
	this := ManagedConnectionAuthorizationServer{}
	return &this
}

// GetIssuerUrl returns the IssuerUrl field value
func (o *ManagedConnectionAuthorizationServer) GetIssuerUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuerUrl
}

// GetIssuerUrlOk returns a tuple with the IssuerUrl field value
// and a boolean to check if the value has been set.
func (o *ManagedConnectionAuthorizationServer) GetIssuerUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuerUrl, true
}

// SetIssuerUrl sets field value
func (o *ManagedConnectionAuthorizationServer) SetIssuerUrl(v string) {
	o.IssuerUrl = v
}

// GetResourceIndicator returns the ResourceIndicator field value if set, zero value otherwise.
func (o *ManagedConnectionAuthorizationServer) GetResourceIndicator() string {
	if o == nil || IsNil(o.ResourceIndicator) {
		var ret string
		return ret
	}
	return *o.ResourceIndicator
}

// GetResourceIndicatorOk returns a tuple with the ResourceIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedConnectionAuthorizationServer) GetResourceIndicatorOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceIndicator) {
		return nil, false
	}
	return o.ResourceIndicator, true
}

// HasResourceIndicator returns a boolean if a field has been set.
func (o *ManagedConnectionAuthorizationServer) HasResourceIndicator() bool {
	if o != nil && !IsNil(o.ResourceIndicator) {
		return true
	}

	return false
}

// SetResourceIndicator gets a reference to the given string and assigns it to the ResourceIndicator field.
func (o *ManagedConnectionAuthorizationServer) SetResourceIndicator(v string) {
	o.ResourceIndicator = &v
}

func (o ManagedConnectionAuthorizationServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedConnectionAuthorizationServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["issuerUrl"] = o.IssuerUrl
	if !IsNil(o.ResourceIndicator) {
		toSerialize["resourceIndicator"] = o.ResourceIndicator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ManagedConnectionAuthorizationServer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"issuerUrl",
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

	varManagedConnectionAuthorizationServer := _ManagedConnectionAuthorizationServer{}

	err = json.Unmarshal(data, &varManagedConnectionAuthorizationServer)

	if err != nil {
		return err
	}

	*o = ManagedConnectionAuthorizationServer(varManagedConnectionAuthorizationServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "issuerUrl")
		delete(additionalProperties, "resourceIndicator")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManagedConnectionAuthorizationServer struct {
	value *ManagedConnectionAuthorizationServer
	isSet bool
}

func (v NullableManagedConnectionAuthorizationServer) Get() *ManagedConnectionAuthorizationServer {
	return v.value
}

func (v *NullableManagedConnectionAuthorizationServer) Set(val *ManagedConnectionAuthorizationServer) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedConnectionAuthorizationServer) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedConnectionAuthorizationServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedConnectionAuthorizationServer(val *ManagedConnectionAuthorizationServer) *NullableManagedConnectionAuthorizationServer {
	return &NullableManagedConnectionAuthorizationServer{value: val, isSet: true}
}

func (v NullableManagedConnectionAuthorizationServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedConnectionAuthorizationServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
