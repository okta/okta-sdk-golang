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
)

// ProvisioningDetails Supported provisioning configurations for your integration
type ProvisioningDetails struct {
	// List of provisioning features supported in this integration
	Features []string `json:"features"`
	Scim Scim `json:"scim"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningDetails ProvisioningDetails

// NewProvisioningDetails instantiates a new ProvisioningDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningDetails(features []string, scim Scim) *ProvisioningDetails {
	this := ProvisioningDetails{}
	this.Features = features
	this.Scim = scim
	return &this
}

// NewProvisioningDetailsWithDefaults instantiates a new ProvisioningDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningDetailsWithDefaults() *ProvisioningDetails {
	this := ProvisioningDetails{}
	return &this
}

// GetFeatures returns the Features field value
func (o *ProvisioningDetails) GetFeatures() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *ProvisioningDetails) GetFeaturesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *ProvisioningDetails) SetFeatures(v []string) {
	o.Features = v
}

// GetScim returns the Scim field value
func (o *ProvisioningDetails) GetScim() Scim {
	if o == nil {
		var ret Scim
		return ret
	}

	return o.Scim
}

// GetScimOk returns a tuple with the Scim field value
// and a boolean to check if the value has been set.
func (o *ProvisioningDetails) GetScimOk() (*Scim, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scim, true
}

// SetScim sets field value
func (o *ProvisioningDetails) SetScim(v Scim) {
	o.Scim = v
}

func (o ProvisioningDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["features"] = o.Features
	}
	if true {
		toSerialize["scim"] = o.Scim
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProvisioningDetails) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningDetails := _ProvisioningDetails{}

	err = json.Unmarshal(bytes, &varProvisioningDetails)
	if err == nil {
		*o = ProvisioningDetails(varProvisioningDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "features")
		delete(additionalProperties, "scim")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProvisioningDetails struct {
	value *ProvisioningDetails
	isSet bool
}

func (v NullableProvisioningDetails) Get() *ProvisioningDetails {
	return v.value
}

func (v *NullableProvisioningDetails) Set(val *ProvisioningDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningDetails(val *ProvisioningDetails) *NullableProvisioningDetails {
	return &NullableProvisioningDetails{value: val, isSet: true}
}

func (v NullableProvisioningDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

