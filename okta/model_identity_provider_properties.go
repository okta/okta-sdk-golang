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

// IdentityProviderProperties struct for IdentityProviderProperties
type IdentityProviderProperties struct {
	AdditionalAmr []string `json:"additionalAmr,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProviderProperties IdentityProviderProperties

// NewIdentityProviderProperties instantiates a new IdentityProviderProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderProperties() *IdentityProviderProperties {
	this := IdentityProviderProperties{}
	return &this
}

// NewIdentityProviderPropertiesWithDefaults instantiates a new IdentityProviderProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderPropertiesWithDefaults() *IdentityProviderProperties {
	this := IdentityProviderProperties{}
	return &this
}

// GetAdditionalAmr returns the AdditionalAmr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityProviderProperties) GetAdditionalAmr() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AdditionalAmr
}

// GetAdditionalAmrOk returns a tuple with the AdditionalAmr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityProviderProperties) GetAdditionalAmrOk() ([]string, bool) {
	if o == nil || o.AdditionalAmr == nil {
		return nil, false
	}
	return o.AdditionalAmr, true
}

// HasAdditionalAmr returns a boolean if a field has been set.
func (o *IdentityProviderProperties) HasAdditionalAmr() bool {
	if o != nil && o.AdditionalAmr != nil {
		return true
	}

	return false
}

// SetAdditionalAmr gets a reference to the given []string and assigns it to the AdditionalAmr field.
func (o *IdentityProviderProperties) SetAdditionalAmr(v []string) {
	o.AdditionalAmr = v
}

func (o IdentityProviderProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalAmr != nil {
		toSerialize["additionalAmr"] = o.AdditionalAmr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityProviderProperties) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityProviderProperties := _IdentityProviderProperties{}

	err = json.Unmarshal(bytes, &varIdentityProviderProperties)
	if err == nil {
		*o = IdentityProviderProperties(varIdentityProviderProperties)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "additionalAmr")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdentityProviderProperties struct {
	value *IdentityProviderProperties
	isSet bool
}

func (v NullableIdentityProviderProperties) Get() *IdentityProviderProperties {
	return v.value
}

func (v *NullableIdentityProviderProperties) Set(val *IdentityProviderProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderProperties(val *IdentityProviderProperties) *NullableIdentityProviderProperties {
	return &NullableIdentityProviderProperties{value: val, isSet: true}
}

func (v NullableIdentityProviderProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

