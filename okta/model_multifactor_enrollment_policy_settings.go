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

// MultifactorEnrollmentPolicySettings struct for MultifactorEnrollmentPolicySettings
type MultifactorEnrollmentPolicySettings struct {
	Authenticators []MultifactorEnrollmentPolicyAuthenticatorSettings `json:"authenticators,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MultifactorEnrollmentPolicySettings MultifactorEnrollmentPolicySettings

// NewMultifactorEnrollmentPolicySettings instantiates a new MultifactorEnrollmentPolicySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultifactorEnrollmentPolicySettings() *MultifactorEnrollmentPolicySettings {
	this := MultifactorEnrollmentPolicySettings{}
	return &this
}

// NewMultifactorEnrollmentPolicySettingsWithDefaults instantiates a new MultifactorEnrollmentPolicySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultifactorEnrollmentPolicySettingsWithDefaults() *MultifactorEnrollmentPolicySettings {
	this := MultifactorEnrollmentPolicySettings{}
	return &this
}

// GetAuthenticators returns the Authenticators field value if set, zero value otherwise.
func (o *MultifactorEnrollmentPolicySettings) GetAuthenticators() []MultifactorEnrollmentPolicyAuthenticatorSettings {
	if o == nil || o.Authenticators == nil {
		var ret []MultifactorEnrollmentPolicyAuthenticatorSettings
		return ret
	}
	return o.Authenticators
}

// GetAuthenticatorsOk returns a tuple with the Authenticators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultifactorEnrollmentPolicySettings) GetAuthenticatorsOk() ([]MultifactorEnrollmentPolicyAuthenticatorSettings, bool) {
	if o == nil || o.Authenticators == nil {
		return nil, false
	}
	return o.Authenticators, true
}

// HasAuthenticators returns a boolean if a field has been set.
func (o *MultifactorEnrollmentPolicySettings) HasAuthenticators() bool {
	if o != nil && o.Authenticators != nil {
		return true
	}

	return false
}

// SetAuthenticators gets a reference to the given []MultifactorEnrollmentPolicyAuthenticatorSettings and assigns it to the Authenticators field.
func (o *MultifactorEnrollmentPolicySettings) SetAuthenticators(v []MultifactorEnrollmentPolicyAuthenticatorSettings) {
	o.Authenticators = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MultifactorEnrollmentPolicySettings) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultifactorEnrollmentPolicySettings) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MultifactorEnrollmentPolicySettings) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MultifactorEnrollmentPolicySettings) SetType(v string) {
	o.Type = &v
}

func (o MultifactorEnrollmentPolicySettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Authenticators != nil {
		toSerialize["authenticators"] = o.Authenticators
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MultifactorEnrollmentPolicySettings) UnmarshalJSON(bytes []byte) (err error) {
	varMultifactorEnrollmentPolicySettings := _MultifactorEnrollmentPolicySettings{}

	err = json.Unmarshal(bytes, &varMultifactorEnrollmentPolicySettings)
	if err == nil {
		*o = MultifactorEnrollmentPolicySettings(varMultifactorEnrollmentPolicySettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authenticators")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableMultifactorEnrollmentPolicySettings struct {
	value *MultifactorEnrollmentPolicySettings
	isSet bool
}

func (v NullableMultifactorEnrollmentPolicySettings) Get() *MultifactorEnrollmentPolicySettings {
	return v.value
}

func (v *NullableMultifactorEnrollmentPolicySettings) Set(val *MultifactorEnrollmentPolicySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMultifactorEnrollmentPolicySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMultifactorEnrollmentPolicySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultifactorEnrollmentPolicySettings(val *MultifactorEnrollmentPolicySettings) *NullableMultifactorEnrollmentPolicySettings {
	return &NullableMultifactorEnrollmentPolicySettings{value: val, isSet: true}
}

func (v NullableMultifactorEnrollmentPolicySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultifactorEnrollmentPolicySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

