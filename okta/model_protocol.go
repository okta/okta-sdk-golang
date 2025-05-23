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

// Protocol struct for Protocol
type Protocol struct {
	Algorithms *ProtocolAlgorithms `json:"algorithms,omitempty"`
	Credentials *IdentityProviderCredentials `json:"credentials,omitempty"`
	Endpoints *ProtocolEndpoints `json:"endpoints,omitempty"`
	Issuer *ProtocolEndpoint `json:"issuer,omitempty"`
	RelayState *ProtocolRelayState `json:"relayState,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
	Settings *ProtocolSettings `json:"settings,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Protocol Protocol

// NewProtocol instantiates a new Protocol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocol() *Protocol {
	this := Protocol{}
	return &this
}

// NewProtocolWithDefaults instantiates a new Protocol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolWithDefaults() *Protocol {
	this := Protocol{}
	return &this
}

// GetAlgorithms returns the Algorithms field value if set, zero value otherwise.
func (o *Protocol) GetAlgorithms() ProtocolAlgorithms {
	if o == nil || o.Algorithms == nil {
		var ret ProtocolAlgorithms
		return ret
	}
	return *o.Algorithms
}

// GetAlgorithmsOk returns a tuple with the Algorithms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Protocol) GetAlgorithmsOk() (*ProtocolAlgorithms, bool) {
	if o == nil || o.Algorithms == nil {
		return nil, false
	}
	return o.Algorithms, true
}

// HasAlgorithms returns a boolean if a field has been set.
func (o *Protocol) HasAlgorithms() bool {
	if o != nil && o.Algorithms != nil {
		return true
	}

	return false
}

// SetAlgorithms gets a reference to the given ProtocolAlgorithms and assigns it to the Algorithms field.
func (o *Protocol) SetAlgorithms(v ProtocolAlgorithms) {
	o.Algorithms = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *Protocol) GetCredentials() IdentityProviderCredentials {
	if o == nil || o.Credentials == nil {
		var ret IdentityProviderCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Protocol) GetCredentialsOk() (*IdentityProviderCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *Protocol) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given IdentityProviderCredentials and assigns it to the Credentials field.
func (o *Protocol) SetCredentials(v IdentityProviderCredentials) {
	o.Credentials = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *Protocol) GetEndpoints() ProtocolEndpoints {
	if o == nil || o.Endpoints == nil {
		var ret ProtocolEndpoints
		return ret
	}
	return *o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Protocol) GetEndpointsOk() (*ProtocolEndpoints, bool) {
	if o == nil || o.Endpoints == nil {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *Protocol) HasEndpoints() bool {
	if o != nil && o.Endpoints != nil {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given ProtocolEndpoints and assigns it to the Endpoints field.
func (o *Protocol) SetEndpoints(v ProtocolEndpoints) {
	o.Endpoints = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *Protocol) GetIssuer() ProtocolEndpoint {
	if o == nil || o.Issuer == nil {
		var ret ProtocolEndpoint
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Protocol) GetIssuerOk() (*ProtocolEndpoint, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *Protocol) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given ProtocolEndpoint and assigns it to the Issuer field.
func (o *Protocol) SetIssuer(v ProtocolEndpoint) {
	o.Issuer = &v
}

// GetRelayState returns the RelayState field value if set, zero value otherwise.
func (o *Protocol) GetRelayState() ProtocolRelayState {
	if o == nil || o.RelayState == nil {
		var ret ProtocolRelayState
		return ret
	}
	return *o.RelayState
}

// GetRelayStateOk returns a tuple with the RelayState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Protocol) GetRelayStateOk() (*ProtocolRelayState, bool) {
	if o == nil || o.RelayState == nil {
		return nil, false
	}
	return o.RelayState, true
}

// HasRelayState returns a boolean if a field has been set.
func (o *Protocol) HasRelayState() bool {
	if o != nil && o.RelayState != nil {
		return true
	}

	return false
}

// SetRelayState gets a reference to the given ProtocolRelayState and assigns it to the RelayState field.
func (o *Protocol) SetRelayState(v ProtocolRelayState) {
	o.RelayState = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *Protocol) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Protocol) GetScopesOk() ([]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *Protocol) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *Protocol) SetScopes(v []string) {
	o.Scopes = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *Protocol) GetSettings() ProtocolSettings {
	if o == nil || o.Settings == nil {
		var ret ProtocolSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Protocol) GetSettingsOk() (*ProtocolSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Protocol) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given ProtocolSettings and assigns it to the Settings field.
func (o *Protocol) SetSettings(v ProtocolSettings) {
	o.Settings = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Protocol) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Protocol) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Protocol) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Protocol) SetType(v string) {
	o.Type = &v
}

func (o Protocol) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Algorithms != nil {
		toSerialize["algorithms"] = o.Algorithms
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Endpoints != nil {
		toSerialize["endpoints"] = o.Endpoints
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.RelayState != nil {
		toSerialize["relayState"] = o.RelayState
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Protocol) UnmarshalJSON(bytes []byte) (err error) {
	varProtocol := _Protocol{}

	err = json.Unmarshal(bytes, &varProtocol)
	if err == nil {
		*o = Protocol(varProtocol)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "algorithms")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "endpoints")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "relayState")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProtocol struct {
	value *Protocol
	isSet bool
}

func (v NullableProtocol) Get() *Protocol {
	return v.value
}

func (v *NullableProtocol) Set(val *Protocol) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocol(val *Protocol) *NullableProtocol {
	return &NullableProtocol{value: val, isSet: true}
}

func (v NullableProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

