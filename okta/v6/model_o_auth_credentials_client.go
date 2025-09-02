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
)

// OAuthCredentialsClient OAuth 2.0 and OpenID Connect Client object > **Note:** You must complete client registration with the IdP Authorization Server for your Okta IdP instance to obtain client credentials.
type OAuthCredentialsClient struct {
	// The [Unique identifier](https://tools.ietf.org/html/rfc6749#section-2.2) issued by the AS for the Okta IdP instance
	ClientId *string `json:"client_id,omitempty"`
	// The [Client secret](https://tools.ietf.org/html/rfc6749#section-2.3.1) issued by the AS for the Okta IdP instance
	ClientSecret *string `json:"client_secret,omitempty"`
	// Require Proof Key for Code Exchange (PKCE) for additional verification
	PkceRequired *bool `json:"pkce_required,omitempty"`
	// Client authentication methods supported by the token endpoint
	TokenEndpointAuthMethod *string `json:"token_endpoint_auth_method,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuthCredentialsClient OAuthCredentialsClient

// NewOAuthCredentialsClient instantiates a new OAuthCredentialsClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthCredentialsClient() *OAuthCredentialsClient {
	this := OAuthCredentialsClient{}
	return &this
}

// NewOAuthCredentialsClientWithDefaults instantiates a new OAuthCredentialsClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthCredentialsClientWithDefaults() *OAuthCredentialsClient {
	this := OAuthCredentialsClient{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OAuthCredentialsClient) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthCredentialsClient) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OAuthCredentialsClient) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OAuthCredentialsClient) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *OAuthCredentialsClient) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthCredentialsClient) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *OAuthCredentialsClient) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *OAuthCredentialsClient) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetPkceRequired returns the PkceRequired field value if set, zero value otherwise.
func (o *OAuthCredentialsClient) GetPkceRequired() bool {
	if o == nil || o.PkceRequired == nil {
		var ret bool
		return ret
	}
	return *o.PkceRequired
}

// GetPkceRequiredOk returns a tuple with the PkceRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthCredentialsClient) GetPkceRequiredOk() (*bool, bool) {
	if o == nil || o.PkceRequired == nil {
		return nil, false
	}
	return o.PkceRequired, true
}

// HasPkceRequired returns a boolean if a field has been set.
func (o *OAuthCredentialsClient) HasPkceRequired() bool {
	if o != nil && o.PkceRequired != nil {
		return true
	}

	return false
}

// SetPkceRequired gets a reference to the given bool and assigns it to the PkceRequired field.
func (o *OAuthCredentialsClient) SetPkceRequired(v bool) {
	o.PkceRequired = &v
}

// GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field value if set, zero value otherwise.
func (o *OAuthCredentialsClient) GetTokenEndpointAuthMethod() string {
	if o == nil || o.TokenEndpointAuthMethod == nil {
		var ret string
		return ret
	}
	return *o.TokenEndpointAuthMethod
}

// GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthCredentialsClient) GetTokenEndpointAuthMethodOk() (*string, bool) {
	if o == nil || o.TokenEndpointAuthMethod == nil {
		return nil, false
	}
	return o.TokenEndpointAuthMethod, true
}

// HasTokenEndpointAuthMethod returns a boolean if a field has been set.
func (o *OAuthCredentialsClient) HasTokenEndpointAuthMethod() bool {
	if o != nil && o.TokenEndpointAuthMethod != nil {
		return true
	}

	return false
}

// SetTokenEndpointAuthMethod gets a reference to the given string and assigns it to the TokenEndpointAuthMethod field.
func (o *OAuthCredentialsClient) SetTokenEndpointAuthMethod(v string) {
	o.TokenEndpointAuthMethod = &v
}

func (o OAuthCredentialsClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *OAuthCredentialsClient) UnmarshalJSON(bytes []byte) (err error) {
	varOAuthCredentialsClient := _OAuthCredentialsClient{}

	err = json.Unmarshal(bytes, &varOAuthCredentialsClient)
	if err == nil {
		*o = OAuthCredentialsClient(varOAuthCredentialsClient)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
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

type NullableOAuthCredentialsClient struct {
	value *OAuthCredentialsClient
	isSet bool
}

func (v NullableOAuthCredentialsClient) Get() *OAuthCredentialsClient {
	return v.value
}

func (v *NullableOAuthCredentialsClient) Set(val *OAuthCredentialsClient) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthCredentialsClient) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthCredentialsClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthCredentialsClient(val *OAuthCredentialsClient) *NullableOAuthCredentialsClient {
	return &NullableOAuthCredentialsClient{value: val, isSet: true}
}

func (v NullableOAuthCredentialsClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthCredentialsClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

