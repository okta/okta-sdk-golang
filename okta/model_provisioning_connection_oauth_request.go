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
)

// ProvisioningConnectionOauthRequest struct for ProvisioningConnectionOauthRequest
type ProvisioningConnectionOauthRequest struct {
	Profile ProvisioningConnectionOauthRequestProfile `json:"profile"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningConnectionOauthRequest ProvisioningConnectionOauthRequest

// NewProvisioningConnectionOauthRequest instantiates a new ProvisioningConnectionOauthRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningConnectionOauthRequest(profile ProvisioningConnectionOauthRequestProfile) *ProvisioningConnectionOauthRequest {
	this := ProvisioningConnectionOauthRequest{}
	this.Profile = profile
	return &this
}

// NewProvisioningConnectionOauthRequestWithDefaults instantiates a new ProvisioningConnectionOauthRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningConnectionOauthRequestWithDefaults() *ProvisioningConnectionOauthRequest {
	this := ProvisioningConnectionOauthRequest{}
	return &this
}

// GetProfile returns the Profile field value
func (o *ProvisioningConnectionOauthRequest) GetProfile() ProvisioningConnectionOauthRequestProfile {
	if o == nil {
		var ret ProvisioningConnectionOauthRequestProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionOauthRequest) GetProfileOk() (*ProvisioningConnectionOauthRequestProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *ProvisioningConnectionOauthRequest) SetProfile(v ProvisioningConnectionOauthRequestProfile) {
	o.Profile = v
}

func (o ProvisioningConnectionOauthRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProvisioningConnectionOauthRequest) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningConnectionOauthRequest := _ProvisioningConnectionOauthRequest{}

	err = json.Unmarshal(bytes, &varProvisioningConnectionOauthRequest)
	if err == nil {
		*o = ProvisioningConnectionOauthRequest(varProvisioningConnectionOauthRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProvisioningConnectionOauthRequest struct {
	value *ProvisioningConnectionOauthRequest
	isSet bool
}

func (v NullableProvisioningConnectionOauthRequest) Get() *ProvisioningConnectionOauthRequest {
	return v.value
}

func (v *NullableProvisioningConnectionOauthRequest) Set(val *ProvisioningConnectionOauthRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConnectionOauthRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConnectionOauthRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConnectionOauthRequest(val *ProvisioningConnectionOauthRequest) *NullableProvisioningConnectionOauthRequest {
	return &NullableProvisioningConnectionOauthRequest{value: val, isSet: true}
}

func (v NullableProvisioningConnectionOauthRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConnectionOauthRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

