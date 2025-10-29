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

// checks if the Scim type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Scim{}

// Scim SCIM configuration details
type Scim struct {
	// The authentication mode for requests to your SCIM server  | authMode | Description | | -------- | ----------- | | `header` | Uses authorization header with a customer-provided token value in the following format: `Authorization: {API token}` | | `bearer` | Uses authorization header with a customer-provided bearer token in the following format: `Authorization: Bearer {API token}` | | {authModeId} | The ID of the auth mode object that contains OAuth 2.0 credentials. <br> **Note:** Use the `/integrations/api/v1/internal/authModes` endpoint to create the auth mode object. |
	AuthMode string `json:"authMode"`
	// The base URL that Okta uses to send outbound calls to your SCIM server. Only the HTTPS protocol is supported. You can use the app-level variables defined in the `config` array for the base URL. For example, if you have a `subdomain` variable defined in the `config` array and the URL to retrieve SCIM users for your integration is `https://${subdomain}.example.com/scim/v2/Users`, then specify the following base URL: `'https://' + app.subdomain + '.example.com/scim/v2'`.
	BaseUri string `json:"baseUri"`
	// List of supported entitlement types
	EntitlementTypes []EntitlementTypesInner `json:"entitlementTypes,omitempty"`
	ScimServerConfig ScimScimServerConfig    `json:"scimServerConfig"`
	// The URL to your customer-facing instructions for configuring your SCIM integration. See [Customer configuration document guidelines](https://developer.okta.com/docs/guides/submit-app-prereq/main/#customer-configuration-document-guidelines).
	SetupInstructionsUri string `json:"setupInstructionsUri"`
	AdditionalProperties map[string]interface{}
}

type _Scim Scim

// NewScim instantiates a new Scim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScim(authMode string, baseUri string, scimServerConfig ScimScimServerConfig, setupInstructionsUri string) *Scim {
	this := Scim{}
	this.AuthMode = authMode
	this.BaseUri = baseUri
	this.ScimServerConfig = scimServerConfig
	this.SetupInstructionsUri = setupInstructionsUri
	return &this
}

// NewScimWithDefaults instantiates a new Scim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimWithDefaults() *Scim {
	this := Scim{}
	return &this
}

// GetAuthMode returns the AuthMode field value
func (o *Scim) GetAuthMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthMode
}

// GetAuthModeOk returns a tuple with the AuthMode field value
// and a boolean to check if the value has been set.
func (o *Scim) GetAuthModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthMode, true
}

// SetAuthMode sets field value
func (o *Scim) SetAuthMode(v string) {
	o.AuthMode = v
}

// GetBaseUri returns the BaseUri field value
func (o *Scim) GetBaseUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUri
}

// GetBaseUriOk returns a tuple with the BaseUri field value
// and a boolean to check if the value has been set.
func (o *Scim) GetBaseUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUri, true
}

// SetBaseUri sets field value
func (o *Scim) SetBaseUri(v string) {
	o.BaseUri = v
}

// GetEntitlementTypes returns the EntitlementTypes field value if set, zero value otherwise.
func (o *Scim) GetEntitlementTypes() []EntitlementTypesInner {
	if o == nil || IsNil(o.EntitlementTypes) {
		var ret []EntitlementTypesInner
		return ret
	}
	return o.EntitlementTypes
}

// GetEntitlementTypesOk returns a tuple with the EntitlementTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scim) GetEntitlementTypesOk() ([]EntitlementTypesInner, bool) {
	if o == nil || IsNil(o.EntitlementTypes) {
		return nil, false
	}
	return o.EntitlementTypes, true
}

// HasEntitlementTypes returns a boolean if a field has been set.
func (o *Scim) HasEntitlementTypes() bool {
	if o != nil && !IsNil(o.EntitlementTypes) {
		return true
	}

	return false
}

// SetEntitlementTypes gets a reference to the given []EntitlementTypesInner and assigns it to the EntitlementTypes field.
func (o *Scim) SetEntitlementTypes(v []EntitlementTypesInner) {
	o.EntitlementTypes = v
}

// GetScimServerConfig returns the ScimServerConfig field value
func (o *Scim) GetScimServerConfig() ScimScimServerConfig {
	if o == nil {
		var ret ScimScimServerConfig
		return ret
	}

	return o.ScimServerConfig
}

// GetScimServerConfigOk returns a tuple with the ScimServerConfig field value
// and a boolean to check if the value has been set.
func (o *Scim) GetScimServerConfigOk() (*ScimScimServerConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScimServerConfig, true
}

// SetScimServerConfig sets field value
func (o *Scim) SetScimServerConfig(v ScimScimServerConfig) {
	o.ScimServerConfig = v
}

// GetSetupInstructionsUri returns the SetupInstructionsUri field value
func (o *Scim) GetSetupInstructionsUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SetupInstructionsUri
}

// GetSetupInstructionsUriOk returns a tuple with the SetupInstructionsUri field value
// and a boolean to check if the value has been set.
func (o *Scim) GetSetupInstructionsUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SetupInstructionsUri, true
}

// SetSetupInstructionsUri sets field value
func (o *Scim) SetSetupInstructionsUri(v string) {
	o.SetupInstructionsUri = v
}

func (o Scim) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Scim) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authMode"] = o.AuthMode
	toSerialize["baseUri"] = o.BaseUri
	if !IsNil(o.EntitlementTypes) {
		toSerialize["entitlementTypes"] = o.EntitlementTypes
	}
	toSerialize["scimServerConfig"] = o.ScimServerConfig
	toSerialize["setupInstructionsUri"] = o.SetupInstructionsUri

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Scim) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authMode",
		"baseUri",
		"scimServerConfig",
		"setupInstructionsUri",
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

	varScim := _Scim{}

	err = json.Unmarshal(data, &varScim)

	if err != nil {
		return err
	}

	*o = Scim(varScim)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authMode")
		delete(additionalProperties, "baseUri")
		delete(additionalProperties, "entitlementTypes")
		delete(additionalProperties, "scimServerConfig")
		delete(additionalProperties, "setupInstructionsUri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScim struct {
	value *Scim
	isSet bool
}

func (v NullableScim) Get() *Scim {
	return v.value
}

func (v *NullableScim) Set(val *Scim) {
	v.value = val
	v.isSet = true
}

func (v NullableScim) IsSet() bool {
	return v.isSet
}

func (v *NullableScim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScim(val *Scim) *NullableScim {
	return &NullableScim{value: val, isSet: true}
}

func (v NullableScim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
