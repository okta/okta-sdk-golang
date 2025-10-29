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
)

// checks if the ProtocolSaml type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProtocolSaml{}

// ProtocolSaml Protocol settings for the [SAML 2.0 Authentication Request Protocol](http://docs.oasis-open.org/security/saml/v2.0/saml-core-2.0-os.pdf)
type ProtocolSaml struct {
	Algorithms  *SamlAlgorithms  `json:"algorithms,omitempty"`
	Credentials *SamlCredentials `json:"credentials,omitempty"`
	Endpoints   *SamlEndpoints   `json:"endpoints,omitempty"`
	RelayState  *SamlRelayState  `json:"relayState,omitempty"`
	Settings    *SamlSettings    `json:"settings,omitempty"`
	// SAML 2.0 protocol
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProtocolSaml ProtocolSaml

// NewProtocolSaml instantiates a new ProtocolSaml object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocolSaml() *ProtocolSaml {
	this := ProtocolSaml{}
	return &this
}

// NewProtocolSamlWithDefaults instantiates a new ProtocolSaml object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolSamlWithDefaults() *ProtocolSaml {
	this := ProtocolSaml{}
	return &this
}

// GetAlgorithms returns the Algorithms field value if set, zero value otherwise.
func (o *ProtocolSaml) GetAlgorithms() SamlAlgorithms {
	if o == nil || IsNil(o.Algorithms) {
		var ret SamlAlgorithms
		return ret
	}
	return *o.Algorithms
}

// GetAlgorithmsOk returns a tuple with the Algorithms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolSaml) GetAlgorithmsOk() (*SamlAlgorithms, bool) {
	if o == nil || IsNil(o.Algorithms) {
		return nil, false
	}
	return o.Algorithms, true
}

// HasAlgorithms returns a boolean if a field has been set.
func (o *ProtocolSaml) HasAlgorithms() bool {
	if o != nil && !IsNil(o.Algorithms) {
		return true
	}

	return false
}

// SetAlgorithms gets a reference to the given SamlAlgorithms and assigns it to the Algorithms field.
func (o *ProtocolSaml) SetAlgorithms(v SamlAlgorithms) {
	o.Algorithms = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *ProtocolSaml) GetCredentials() SamlCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret SamlCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolSaml) GetCredentialsOk() (*SamlCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *ProtocolSaml) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given SamlCredentials and assigns it to the Credentials field.
func (o *ProtocolSaml) SetCredentials(v SamlCredentials) {
	o.Credentials = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *ProtocolSaml) GetEndpoints() SamlEndpoints {
	if o == nil || IsNil(o.Endpoints) {
		var ret SamlEndpoints
		return ret
	}
	return *o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolSaml) GetEndpointsOk() (*SamlEndpoints, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *ProtocolSaml) HasEndpoints() bool {
	if o != nil && !IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given SamlEndpoints and assigns it to the Endpoints field.
func (o *ProtocolSaml) SetEndpoints(v SamlEndpoints) {
	o.Endpoints = &v
}

// GetRelayState returns the RelayState field value if set, zero value otherwise.
func (o *ProtocolSaml) GetRelayState() SamlRelayState {
	if o == nil || IsNil(o.RelayState) {
		var ret SamlRelayState
		return ret
	}
	return *o.RelayState
}

// GetRelayStateOk returns a tuple with the RelayState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolSaml) GetRelayStateOk() (*SamlRelayState, bool) {
	if o == nil || IsNil(o.RelayState) {
		return nil, false
	}
	return o.RelayState, true
}

// HasRelayState returns a boolean if a field has been set.
func (o *ProtocolSaml) HasRelayState() bool {
	if o != nil && !IsNil(o.RelayState) {
		return true
	}

	return false
}

// SetRelayState gets a reference to the given SamlRelayState and assigns it to the RelayState field.
func (o *ProtocolSaml) SetRelayState(v SamlRelayState) {
	o.RelayState = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ProtocolSaml) GetSettings() SamlSettings {
	if o == nil || IsNil(o.Settings) {
		var ret SamlSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolSaml) GetSettingsOk() (*SamlSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ProtocolSaml) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given SamlSettings and assigns it to the Settings field.
func (o *ProtocolSaml) SetSettings(v SamlSettings) {
	o.Settings = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProtocolSaml) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolSaml) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProtocolSaml) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProtocolSaml) SetType(v string) {
	o.Type = &v
}

func (o ProtocolSaml) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProtocolSaml) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Algorithms) {
		toSerialize["algorithms"] = o.Algorithms
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.Endpoints) {
		toSerialize["endpoints"] = o.Endpoints
	}
	if !IsNil(o.RelayState) {
		toSerialize["relayState"] = o.RelayState
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProtocolSaml) UnmarshalJSON(data []byte) (err error) {
	varProtocolSaml := _ProtocolSaml{}

	err = json.Unmarshal(data, &varProtocolSaml)

	if err != nil {
		return err
	}

	*o = ProtocolSaml(varProtocolSaml)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "algorithms")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "endpoints")
		delete(additionalProperties, "relayState")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProtocolSaml struct {
	value *ProtocolSaml
	isSet bool
}

func (v NullableProtocolSaml) Get() *ProtocolSaml {
	return v.value
}

func (v *NullableProtocolSaml) Set(val *ProtocolSaml) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolSaml) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolSaml) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolSaml(val *ProtocolSaml) *NullableProtocolSaml {
	return &NullableProtocolSaml{value: val, isSet: true}
}

func (v NullableProtocolSaml) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolSaml) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
