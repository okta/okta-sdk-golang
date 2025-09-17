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
)

// checks if the ProvisioningConnectionTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningConnectionTokenRequest{}

// ProvisioningConnectionTokenRequest struct for ProvisioningConnectionTokenRequest
type ProvisioningConnectionTokenRequest struct {
	// Only used for the Zscaler 2.0 (`zscalerbyz`) app. The base URL for the Zscaler 2.0 target app, which also contains the Zscaler ID.
	BaseUrl              *string                                   `json:"baseUrl,omitempty"`
	Profile              ProvisioningConnectionTokenRequestProfile `json:"profile"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningConnectionTokenRequest ProvisioningConnectionTokenRequest

// NewProvisioningConnectionTokenRequest instantiates a new ProvisioningConnectionTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningConnectionTokenRequest(profile ProvisioningConnectionTokenRequestProfile) *ProvisioningConnectionTokenRequest {
	this := ProvisioningConnectionTokenRequest{}
	this.Profile = profile
	return &this
}

// NewProvisioningConnectionTokenRequestWithDefaults instantiates a new ProvisioningConnectionTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningConnectionTokenRequestWithDefaults() *ProvisioningConnectionTokenRequest {
	this := ProvisioningConnectionTokenRequest{}
	return &this
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *ProvisioningConnectionTokenRequest) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl) {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionTokenRequest) GetBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BaseUrl) {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *ProvisioningConnectionTokenRequest) HasBaseUrl() bool {
	if o != nil && !IsNil(o.BaseUrl) {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *ProvisioningConnectionTokenRequest) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetProfile returns the Profile field value
func (o *ProvisioningConnectionTokenRequest) GetProfile() ProvisioningConnectionTokenRequestProfile {
	if o == nil {
		var ret ProvisioningConnectionTokenRequestProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionTokenRequest) GetProfileOk() (*ProvisioningConnectionTokenRequestProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *ProvisioningConnectionTokenRequest) SetProfile(v ProvisioningConnectionTokenRequestProfile) {
	o.Profile = v
}

func (o ProvisioningConnectionTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningConnectionTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BaseUrl) {
		toSerialize["baseUrl"] = o.BaseUrl
	}
	toSerialize["profile"] = o.Profile

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningConnectionTokenRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"profile",
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

	varProvisioningConnectionTokenRequest := _ProvisioningConnectionTokenRequest{}

	err = json.Unmarshal(data, &varProvisioningConnectionTokenRequest)

	if err != nil {
		return err
	}

	*o = ProvisioningConnectionTokenRequest(varProvisioningConnectionTokenRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "baseUrl")
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningConnectionTokenRequest struct {
	value *ProvisioningConnectionTokenRequest
	isSet bool
}

func (v NullableProvisioningConnectionTokenRequest) Get() *ProvisioningConnectionTokenRequest {
	return v.value
}

func (v *NullableProvisioningConnectionTokenRequest) Set(val *ProvisioningConnectionTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConnectionTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConnectionTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConnectionTokenRequest(val *ProvisioningConnectionTokenRequest) *NullableProvisioningConnectionTokenRequest {
	return &NullableProvisioningConnectionTokenRequest{value: val, isSet: true}
}

func (v NullableProvisioningConnectionTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConnectionTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
