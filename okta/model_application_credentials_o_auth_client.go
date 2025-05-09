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

// ApplicationCredentialsOAuthClient struct for ApplicationCredentialsOAuthClient
type ApplicationCredentialsOAuthClient struct {
	AutoKeyRotation *bool `json:"autoKeyRotation,omitempty"`
	ClientId *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`
	// Require Proof Key for Code Exchange (PKCE) for additional verification
	PkceRequired *bool `json:"pkce_required,omitempty"`
	TokenEndpointAuthMethod *string `json:"token_endpoint_auth_method,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationCredentialsOAuthClient ApplicationCredentialsOAuthClient

// NewApplicationCredentialsOAuthClient instantiates a new ApplicationCredentialsOAuthClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCredentialsOAuthClient() *ApplicationCredentialsOAuthClient {
	this := ApplicationCredentialsOAuthClient{}
	return &this
}

// NewApplicationCredentialsOAuthClientWithDefaults instantiates a new ApplicationCredentialsOAuthClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCredentialsOAuthClientWithDefaults() *ApplicationCredentialsOAuthClient {
	this := ApplicationCredentialsOAuthClient{}
	return &this
}

// GetAutoKeyRotation returns the AutoKeyRotation field value if set, zero value otherwise.
func (o *ApplicationCredentialsOAuthClient) GetAutoKeyRotation() bool {
	if o == nil || o.AutoKeyRotation == nil {
		var ret bool
		return ret
	}
	return *o.AutoKeyRotation
}

// GetAutoKeyRotationOk returns a tuple with the AutoKeyRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsOAuthClient) GetAutoKeyRotationOk() (*bool, bool) {
	if o == nil || o.AutoKeyRotation == nil {
		return nil, false
	}
	return o.AutoKeyRotation, true
}

// HasAutoKeyRotation returns a boolean if a field has been set.
func (o *ApplicationCredentialsOAuthClient) HasAutoKeyRotation() bool {
	if o != nil && o.AutoKeyRotation != nil {
		return true
	}

	return false
}

// SetAutoKeyRotation gets a reference to the given bool and assigns it to the AutoKeyRotation field.
func (o *ApplicationCredentialsOAuthClient) SetAutoKeyRotation(v bool) {
	o.AutoKeyRotation = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ApplicationCredentialsOAuthClient) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsOAuthClient) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ApplicationCredentialsOAuthClient) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ApplicationCredentialsOAuthClient) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ApplicationCredentialsOAuthClient) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsOAuthClient) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ApplicationCredentialsOAuthClient) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ApplicationCredentialsOAuthClient) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetPkceRequired returns the PkceRequired field value if set, zero value otherwise.
func (o *ApplicationCredentialsOAuthClient) GetPkceRequired() bool {
	if o == nil || o.PkceRequired == nil {
		var ret bool
		return ret
	}
	return *o.PkceRequired
}

// GetPkceRequiredOk returns a tuple with the PkceRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsOAuthClient) GetPkceRequiredOk() (*bool, bool) {
	if o == nil || o.PkceRequired == nil {
		return nil, false
	}
	return o.PkceRequired, true
}

// HasPkceRequired returns a boolean if a field has been set.
func (o *ApplicationCredentialsOAuthClient) HasPkceRequired() bool {
	if o != nil && o.PkceRequired != nil {
		return true
	}

	return false
}

// SetPkceRequired gets a reference to the given bool and assigns it to the PkceRequired field.
func (o *ApplicationCredentialsOAuthClient) SetPkceRequired(v bool) {
	o.PkceRequired = &v
}

// GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field value if set, zero value otherwise.
func (o *ApplicationCredentialsOAuthClient) GetTokenEndpointAuthMethod() string {
	if o == nil || o.TokenEndpointAuthMethod == nil {
		var ret string
		return ret
	}
	return *o.TokenEndpointAuthMethod
}

// GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsOAuthClient) GetTokenEndpointAuthMethodOk() (*string, bool) {
	if o == nil || o.TokenEndpointAuthMethod == nil {
		return nil, false
	}
	return o.TokenEndpointAuthMethod, true
}

// HasTokenEndpointAuthMethod returns a boolean if a field has been set.
func (o *ApplicationCredentialsOAuthClient) HasTokenEndpointAuthMethod() bool {
	if o != nil && o.TokenEndpointAuthMethod != nil {
		return true
	}

	return false
}

// SetTokenEndpointAuthMethod gets a reference to the given string and assigns it to the TokenEndpointAuthMethod field.
func (o *ApplicationCredentialsOAuthClient) SetTokenEndpointAuthMethod(v string) {
	o.TokenEndpointAuthMethod = &v
}

func (o ApplicationCredentialsOAuthClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoKeyRotation != nil {
		toSerialize["autoKeyRotation"] = o.AutoKeyRotation
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.PkceRequired != nil {
		toSerialize["pkce_required"] = o.PkceRequired
	}
	if o.TokenEndpointAuthMethod != nil {
		toSerialize["token_endpoint_auth_method"] = o.TokenEndpointAuthMethod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationCredentialsOAuthClient) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationCredentialsOAuthClient := _ApplicationCredentialsOAuthClient{}

	err = json.Unmarshal(bytes, &varApplicationCredentialsOAuthClient)
	if err == nil {
		*o = ApplicationCredentialsOAuthClient(varApplicationCredentialsOAuthClient)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "autoKeyRotation")
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "client_secret")
		delete(additionalProperties, "pkce_required")
		delete(additionalProperties, "token_endpoint_auth_method")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableApplicationCredentialsOAuthClient struct {
	value *ApplicationCredentialsOAuthClient
	isSet bool
}

func (v NullableApplicationCredentialsOAuthClient) Get() *ApplicationCredentialsOAuthClient {
	return v.value
}

func (v *NullableApplicationCredentialsOAuthClient) Set(val *ApplicationCredentialsOAuthClient) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCredentialsOAuthClient) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCredentialsOAuthClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCredentialsOAuthClient(val *ApplicationCredentialsOAuthClient) *NullableApplicationCredentialsOAuthClient {
	return &NullableApplicationCredentialsOAuthClient{value: val, isSet: true}
}

func (v NullableApplicationCredentialsOAuthClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCredentialsOAuthClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

