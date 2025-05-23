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

// APNSConfiguration struct for APNSConfiguration
type APNSConfiguration struct {
	// (Optional) File name for Admin Console display
	FileName *string `json:"fileName,omitempty"`
	// 10-character Key ID obtained from the Apple developer account
	KeyId *string `json:"keyId,omitempty"`
	// 10-character Team ID used to develop the iOS app
	TeamId *string `json:"teamId,omitempty"`
	// APNs private authentication token signing key
	TokenSigningKey *string `json:"tokenSigningKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _APNSConfiguration APNSConfiguration

// NewAPNSConfiguration instantiates a new APNSConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPNSConfiguration() *APNSConfiguration {
	this := APNSConfiguration{}
	return &this
}

// NewAPNSConfigurationWithDefaults instantiates a new APNSConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPNSConfigurationWithDefaults() *APNSConfiguration {
	this := APNSConfiguration{}
	return &this
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *APNSConfiguration) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APNSConfiguration) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *APNSConfiguration) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *APNSConfiguration) SetFileName(v string) {
	o.FileName = &v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *APNSConfiguration) GetKeyId() string {
	if o == nil || o.KeyId == nil {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APNSConfiguration) GetKeyIdOk() (*string, bool) {
	if o == nil || o.KeyId == nil {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *APNSConfiguration) HasKeyId() bool {
	if o != nil && o.KeyId != nil {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *APNSConfiguration) SetKeyId(v string) {
	o.KeyId = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *APNSConfiguration) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APNSConfiguration) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *APNSConfiguration) HasTeamId() bool {
	if o != nil && o.TeamId != nil {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *APNSConfiguration) SetTeamId(v string) {
	o.TeamId = &v
}

// GetTokenSigningKey returns the TokenSigningKey field value if set, zero value otherwise.
func (o *APNSConfiguration) GetTokenSigningKey() string {
	if o == nil || o.TokenSigningKey == nil {
		var ret string
		return ret
	}
	return *o.TokenSigningKey
}

// GetTokenSigningKeyOk returns a tuple with the TokenSigningKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APNSConfiguration) GetTokenSigningKeyOk() (*string, bool) {
	if o == nil || o.TokenSigningKey == nil {
		return nil, false
	}
	return o.TokenSigningKey, true
}

// HasTokenSigningKey returns a boolean if a field has been set.
func (o *APNSConfiguration) HasTokenSigningKey() bool {
	if o != nil && o.TokenSigningKey != nil {
		return true
	}

	return false
}

// SetTokenSigningKey gets a reference to the given string and assigns it to the TokenSigningKey field.
func (o *APNSConfiguration) SetTokenSigningKey(v string) {
	o.TokenSigningKey = &v
}

func (o APNSConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FileName != nil {
		toSerialize["fileName"] = o.FileName
	}
	if o.KeyId != nil {
		toSerialize["keyId"] = o.KeyId
	}
	if o.TeamId != nil {
		toSerialize["teamId"] = o.TeamId
	}
	if o.TokenSigningKey != nil {
		toSerialize["tokenSigningKey"] = o.TokenSigningKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *APNSConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varAPNSConfiguration := _APNSConfiguration{}

	err = json.Unmarshal(bytes, &varAPNSConfiguration)
	if err == nil {
		*o = APNSConfiguration(varAPNSConfiguration)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "fileName")
		delete(additionalProperties, "keyId")
		delete(additionalProperties, "teamId")
		delete(additionalProperties, "tokenSigningKey")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAPNSConfiguration struct {
	value *APNSConfiguration
	isSet bool
}

func (v NullableAPNSConfiguration) Get() *APNSConfiguration {
	return v.value
}

func (v *NullableAPNSConfiguration) Set(val *APNSConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAPNSConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAPNSConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPNSConfiguration(val *APNSConfiguration) *NullableAPNSConfiguration {
	return &NullableAPNSConfiguration{value: val, isSet: true}
}

func (v NullableAPNSConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPNSConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

