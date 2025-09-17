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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the ApplicationCredentialsOAuthClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCredentialsOAuthClient{}

// ApplicationCredentialsOAuthClient struct for ApplicationCredentialsOAuthClient
type ApplicationCredentialsOAuthClient struct {
	// Requested key rotation mode
	AutoKeyRotation *bool `json:"autoKeyRotation,omitempty"`
	// Unique identifier for the OAuth 2.0 client app  > **Notes:** > * If you don't specify the `client_id`, this immutable property is populated with the [Application instance ID](/openapi/okta-management/management/tag/Application/#tag/Application/operation/getApplication!c=200&path=4/id&t=response). > * The `client_id` must consist of alphanumeric characters or the following special characters: `$-_.+!*'(),`. > * You can't use the reserved word `ALL_CLIENTS`.
	ClientId *string `json:"client_id,omitempty"`
	// OAuth 2.0 client secret string (used for confidential clients)  > **Notes:** If a `client_secret` isn't provided on creation, and the `token_endpoint_auth_method` requires one, Okta generates a random `client_secret` for the client app. > The `client_secret` is only shown when an OAuth 2.0 client app is created or updated (and only if the `token_endpoint_auth_method` requires a client secret).
	ClientSecret *string `json:"client_secret,omitempty"`
	// Requires Proof Key for Code Exchange (PKCE) for additional verification. If `token_endpoint_auth_method` is `none`, then `pkce_required` must be `true`. The default is `true` for browser and native app types.
	PkceRequired *bool `json:"pkce_required,omitempty"`
	// Requested authentication method for the token endpoint
	TokenEndpointAuthMethod *string `json:"token_endpoint_auth_method,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _ApplicationCredentialsOAuthClient ApplicationCredentialsOAuthClient

// NewApplicationCredentialsOAuthClient instantiates a new ApplicationCredentialsOAuthClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCredentialsOAuthClient() *ApplicationCredentialsOAuthClient {
	this := ApplicationCredentialsOAuthClient{}
	var autoKeyRotation bool = true
	this.AutoKeyRotation = &autoKeyRotation
	var pkceRequired bool = true
	this.PkceRequired = &pkceRequired
	var tokenEndpointAuthMethod string = "client_secret_basic"
	this.TokenEndpointAuthMethod = &tokenEndpointAuthMethod
	return &this
}

// NewApplicationCredentialsOAuthClientWithDefaults instantiates a new ApplicationCredentialsOAuthClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCredentialsOAuthClientWithDefaults() *ApplicationCredentialsOAuthClient {
	this := ApplicationCredentialsOAuthClient{}
	var autoKeyRotation bool = true
	this.AutoKeyRotation = &autoKeyRotation
	var pkceRequired bool = true
	this.PkceRequired = &pkceRequired
	var tokenEndpointAuthMethod string = "client_secret_basic"
	this.TokenEndpointAuthMethod = &tokenEndpointAuthMethod
	return &this
}

// GetAutoKeyRotation returns the AutoKeyRotation field value if set, zero value otherwise.
func (o *ApplicationCredentialsOAuthClient) GetAutoKeyRotation() bool {
	if o == nil || IsNil(o.AutoKeyRotation) {
		var ret bool
		return ret
	}
	return *o.AutoKeyRotation
}

// GetAutoKeyRotationOk returns a tuple with the AutoKeyRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsOAuthClient) GetAutoKeyRotationOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoKeyRotation) {
		return nil, false
	}
	return o.AutoKeyRotation, true
}

// HasAutoKeyRotation returns a boolean if a field has been set.
func (o *ApplicationCredentialsOAuthClient) HasAutoKeyRotation() bool {
	if o != nil && !IsNil(o.AutoKeyRotation) {
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
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsOAuthClient) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ApplicationCredentialsOAuthClient) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
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
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsOAuthClient) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ApplicationCredentialsOAuthClient) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
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
	if o == nil || IsNil(o.PkceRequired) {
		var ret bool
		return ret
	}
	return *o.PkceRequired
}

// GetPkceRequiredOk returns a tuple with the PkceRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsOAuthClient) GetPkceRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.PkceRequired) {
		return nil, false
	}
	return o.PkceRequired, true
}

// HasPkceRequired returns a boolean if a field has been set.
func (o *ApplicationCredentialsOAuthClient) HasPkceRequired() bool {
	if o != nil && !IsNil(o.PkceRequired) {
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
	if o == nil || IsNil(o.TokenEndpointAuthMethod) {
		var ret string
		return ret
	}
	return *o.TokenEndpointAuthMethod
}

// GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsOAuthClient) GetTokenEndpointAuthMethodOk() (*string, bool) {
	if o == nil || IsNil(o.TokenEndpointAuthMethod) {
		return nil, false
	}
	return o.TokenEndpointAuthMethod, true
}

// HasTokenEndpointAuthMethod returns a boolean if a field has been set.
func (o *ApplicationCredentialsOAuthClient) HasTokenEndpointAuthMethod() bool {
	if o != nil && !IsNil(o.TokenEndpointAuthMethod) {
		return true
	}

	return false
}

// SetTokenEndpointAuthMethod gets a reference to the given string and assigns it to the TokenEndpointAuthMethod field.
func (o *ApplicationCredentialsOAuthClient) SetTokenEndpointAuthMethod(v string) {
	o.TokenEndpointAuthMethod = &v
}

func (o ApplicationCredentialsOAuthClient) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCredentialsOAuthClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoKeyRotation) {
		toSerialize["autoKeyRotation"] = o.AutoKeyRotation
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if !IsNil(o.PkceRequired) {
		toSerialize["pkce_required"] = o.PkceRequired
	}
	if !IsNil(o.TokenEndpointAuthMethod) {
		toSerialize["token_endpoint_auth_method"] = o.TokenEndpointAuthMethod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationCredentialsOAuthClient) UnmarshalJSON(data []byte) (err error) {
	varApplicationCredentialsOAuthClient := _ApplicationCredentialsOAuthClient{}

	err = json.Unmarshal(data, &varApplicationCredentialsOAuthClient)

	if err != nil {
		return err
	}

	*o = ApplicationCredentialsOAuthClient(varApplicationCredentialsOAuthClient)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "autoKeyRotation")
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "client_secret")
		delete(additionalProperties, "pkce_required")
		delete(additionalProperties, "token_endpoint_auth_method")
		o.AdditionalProperties = additionalProperties
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
