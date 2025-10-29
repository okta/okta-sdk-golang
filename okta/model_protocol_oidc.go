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

// checks if the ProtocolOidc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProtocolOidc{}

// ProtocolOidc Protocol settings for authentication using the [OpenID Connect Protocol](http://openid.net/specs/openid-connect-core-1_0.html#CodeFlowAuth)
type ProtocolOidc struct {
	Algorithms  *OidcAlgorithms   `json:"algorithms,omitempty"`
	Credentials *OAuthCredentials `json:"credentials,omitempty"`
	Endpoints   *OAuthEndpoints   `json:"endpoints,omitempty"`
	// URL of the IdP org
	OktaIdpOrgUrl *string `json:"oktaIdpOrgUrl,omitempty"`
	// OpenID Connect and IdP-defined permission bundles to request delegated access from the user > **Note:** The [IdP type](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/IdentityProvider/#tag/IdentityProvider/operation/createIdentityProvider!path=type&t=request) table lists the scopes that are supported for each IdP.
	Scopes   []string      `json:"scopes,omitempty"`
	Settings *OidcSettings `json:"settings,omitempty"`
	// OpenID Connect Authorization Code flow
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProtocolOidc ProtocolOidc

// NewProtocolOidc instantiates a new ProtocolOidc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocolOidc() *ProtocolOidc {
	this := ProtocolOidc{}
	return &this
}

// NewProtocolOidcWithDefaults instantiates a new ProtocolOidc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolOidcWithDefaults() *ProtocolOidc {
	this := ProtocolOidc{}
	return &this
}

// GetAlgorithms returns the Algorithms field value if set, zero value otherwise.
func (o *ProtocolOidc) GetAlgorithms() OidcAlgorithms {
	if o == nil || IsNil(o.Algorithms) {
		var ret OidcAlgorithms
		return ret
	}
	return *o.Algorithms
}

// GetAlgorithmsOk returns a tuple with the Algorithms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolOidc) GetAlgorithmsOk() (*OidcAlgorithms, bool) {
	if o == nil || IsNil(o.Algorithms) {
		return nil, false
	}
	return o.Algorithms, true
}

// HasAlgorithms returns a boolean if a field has been set.
func (o *ProtocolOidc) HasAlgorithms() bool {
	if o != nil && !IsNil(o.Algorithms) {
		return true
	}

	return false
}

// SetAlgorithms gets a reference to the given OidcAlgorithms and assigns it to the Algorithms field.
func (o *ProtocolOidc) SetAlgorithms(v OidcAlgorithms) {
	o.Algorithms = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *ProtocolOidc) GetCredentials() OAuthCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret OAuthCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolOidc) GetCredentialsOk() (*OAuthCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *ProtocolOidc) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given OAuthCredentials and assigns it to the Credentials field.
func (o *ProtocolOidc) SetCredentials(v OAuthCredentials) {
	o.Credentials = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *ProtocolOidc) GetEndpoints() OAuthEndpoints {
	if o == nil || IsNil(o.Endpoints) {
		var ret OAuthEndpoints
		return ret
	}
	return *o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolOidc) GetEndpointsOk() (*OAuthEndpoints, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *ProtocolOidc) HasEndpoints() bool {
	if o != nil && !IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given OAuthEndpoints and assigns it to the Endpoints field.
func (o *ProtocolOidc) SetEndpoints(v OAuthEndpoints) {
	o.Endpoints = &v
}

// GetOktaIdpOrgUrl returns the OktaIdpOrgUrl field value if set, zero value otherwise.
func (o *ProtocolOidc) GetOktaIdpOrgUrl() string {
	if o == nil || IsNil(o.OktaIdpOrgUrl) {
		var ret string
		return ret
	}
	return *o.OktaIdpOrgUrl
}

// GetOktaIdpOrgUrlOk returns a tuple with the OktaIdpOrgUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolOidc) GetOktaIdpOrgUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OktaIdpOrgUrl) {
		return nil, false
	}
	return o.OktaIdpOrgUrl, true
}

// HasOktaIdpOrgUrl returns a boolean if a field has been set.
func (o *ProtocolOidc) HasOktaIdpOrgUrl() bool {
	if o != nil && !IsNil(o.OktaIdpOrgUrl) {
		return true
	}

	return false
}

// SetOktaIdpOrgUrl gets a reference to the given string and assigns it to the OktaIdpOrgUrl field.
func (o *ProtocolOidc) SetOktaIdpOrgUrl(v string) {
	o.OktaIdpOrgUrl = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ProtocolOidc) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolOidc) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ProtocolOidc) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ProtocolOidc) SetScopes(v []string) {
	o.Scopes = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ProtocolOidc) GetSettings() OidcSettings {
	if o == nil || IsNil(o.Settings) {
		var ret OidcSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolOidc) GetSettingsOk() (*OidcSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ProtocolOidc) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given OidcSettings and assigns it to the Settings field.
func (o *ProtocolOidc) SetSettings(v OidcSettings) {
	o.Settings = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProtocolOidc) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolOidc) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProtocolOidc) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProtocolOidc) SetType(v string) {
	o.Type = &v
}

func (o ProtocolOidc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProtocolOidc) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.OktaIdpOrgUrl) {
		toSerialize["oktaIdpOrgUrl"] = o.OktaIdpOrgUrl
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
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

func (o *ProtocolOidc) UnmarshalJSON(data []byte) (err error) {
	varProtocolOidc := _ProtocolOidc{}

	err = json.Unmarshal(data, &varProtocolOidc)

	if err != nil {
		return err
	}

	*o = ProtocolOidc(varProtocolOidc)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "algorithms")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "endpoints")
		delete(additionalProperties, "oktaIdpOrgUrl")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProtocolOidc struct {
	value *ProtocolOidc
	isSet bool
}

func (v NullableProtocolOidc) Get() *ProtocolOidc {
	return v.value
}

func (v *NullableProtocolOidc) Set(val *ProtocolOidc) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolOidc) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolOidc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolOidc(val *ProtocolOidc) *NullableProtocolOidc {
	return &NullableProtocolOidc{value: val, isSet: true}
}

func (v NullableProtocolOidc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolOidc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
