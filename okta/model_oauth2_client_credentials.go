/*
Okta Admin Management API

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
	"reflect"
	"strings"
)

// checks if the OAUTH2CLIENTCREDENTIALS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAUTH2CLIENTCREDENTIALS{}

// OAUTH2CLIENTCREDENTIALS struct for OAUTH2CLIENTCREDENTIALS
type OAUTH2CLIENTCREDENTIALS struct {
	BaseEmailServer
	// The client ID that's used to access your SMTP server. This client ID is obtained when you create an OAuth 2.0 app with your email provider.
	ClientId *string `json:"clientId,omitempty"`
	// The client secret that's used to access your SMTP server. This client secret is obtained when you create an OAuth 2.0 app with your email provider.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// List of OAuth 2.0 scopes for your SMTP server. You must provide a scope that allows your email server to send emails.
	Scopes []string `json:"scopes,omitempty"`
	// The email provider's specific URL where the OAuth 2.0 app sends its credentials (or signed JWT) to exchange them for an access token
	TokenEndpoint *string `json:"tokenEndpoint,omitempty"`
	// This method determines how your OAuth 2.0 app sends its credentials (`client_id` and `client_secret`) to the provider's server when requesting an access token
	TokenEndpointAuthMethod *string `json:"tokenEndpointAuthMethod,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _OAUTH2CLIENTCREDENTIALS OAUTH2CLIENTCREDENTIALS

// NewOAUTH2CLIENTCREDENTIALS instantiates a new OAUTH2CLIENTCREDENTIALS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAUTH2CLIENTCREDENTIALS() *OAUTH2CLIENTCREDENTIALS {
	this := OAUTH2CLIENTCREDENTIALS{}
	return &this
}

// NewOAUTH2CLIENTCREDENTIALSWithDefaults instantiates a new OAUTH2CLIENTCREDENTIALS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAUTH2CLIENTCREDENTIALSWithDefaults() *OAUTH2CLIENTCREDENTIALS {
	this := OAUTH2CLIENTCREDENTIALS{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OAUTH2CLIENTCREDENTIALS) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAUTH2CLIENTCREDENTIALS) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OAUTH2CLIENTCREDENTIALS) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OAUTH2CLIENTCREDENTIALS) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *OAUTH2CLIENTCREDENTIALS) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAUTH2CLIENTCREDENTIALS) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *OAUTH2CLIENTCREDENTIALS) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *OAUTH2CLIENTCREDENTIALS) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *OAUTH2CLIENTCREDENTIALS) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAUTH2CLIENTCREDENTIALS) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *OAUTH2CLIENTCREDENTIALS) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *OAUTH2CLIENTCREDENTIALS) SetScopes(v []string) {
	o.Scopes = v
}

// GetTokenEndpoint returns the TokenEndpoint field value if set, zero value otherwise.
func (o *OAUTH2CLIENTCREDENTIALS) GetTokenEndpoint() string {
	if o == nil || IsNil(o.TokenEndpoint) {
		var ret string
		return ret
	}
	return *o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAUTH2CLIENTCREDENTIALS) GetTokenEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.TokenEndpoint) {
		return nil, false
	}
	return o.TokenEndpoint, true
}

// HasTokenEndpoint returns a boolean if a field has been set.
func (o *OAUTH2CLIENTCREDENTIALS) HasTokenEndpoint() bool {
	if o != nil && !IsNil(o.TokenEndpoint) {
		return true
	}

	return false
}

// SetTokenEndpoint gets a reference to the given string and assigns it to the TokenEndpoint field.
func (o *OAUTH2CLIENTCREDENTIALS) SetTokenEndpoint(v string) {
	o.TokenEndpoint = &v
}

// GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field value if set, zero value otherwise.
func (o *OAUTH2CLIENTCREDENTIALS) GetTokenEndpointAuthMethod() string {
	if o == nil || IsNil(o.TokenEndpointAuthMethod) {
		var ret string
		return ret
	}
	return *o.TokenEndpointAuthMethod
}

// GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAUTH2CLIENTCREDENTIALS) GetTokenEndpointAuthMethodOk() (*string, bool) {
	if o == nil || IsNil(o.TokenEndpointAuthMethod) {
		return nil, false
	}
	return o.TokenEndpointAuthMethod, true
}

// HasTokenEndpointAuthMethod returns a boolean if a field has been set.
func (o *OAUTH2CLIENTCREDENTIALS) HasTokenEndpointAuthMethod() bool {
	if o != nil && !IsNil(o.TokenEndpointAuthMethod) {
		return true
	}

	return false
}

// SetTokenEndpointAuthMethod gets a reference to the given string and assigns it to the TokenEndpointAuthMethod field.
func (o *OAUTH2CLIENTCREDENTIALS) SetTokenEndpointAuthMethod(v string) {
	o.TokenEndpointAuthMethod = &v
}

func (o OAUTH2CLIENTCREDENTIALS) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAUTH2CLIENTCREDENTIALS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseEmailServer, errBaseEmailServer := json.Marshal(o.BaseEmailServer)
	if errBaseEmailServer != nil {
		return map[string]interface{}{}, errBaseEmailServer
	}
	errBaseEmailServer = json.Unmarshal([]byte(serializedBaseEmailServer), &toSerialize)
	if errBaseEmailServer != nil {
		return map[string]interface{}{}, errBaseEmailServer
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.TokenEndpoint) {
		toSerialize["tokenEndpoint"] = o.TokenEndpoint
	}
	if !IsNil(o.TokenEndpointAuthMethod) {
		toSerialize["tokenEndpointAuthMethod"] = o.TokenEndpointAuthMethod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAUTH2CLIENTCREDENTIALS) UnmarshalJSON(data []byte) (err error) {
	type OAUTH2CLIENTCREDENTIALSWithoutEmbeddedStruct struct {
		// The client ID that's used to access your SMTP server. This client ID is obtained when you create an OAuth 2.0 app with your email provider.
		ClientId *string `json:"clientId,omitempty"`
		// The client secret that's used to access your SMTP server. This client secret is obtained when you create an OAuth 2.0 app with your email provider.
		ClientSecret *string `json:"clientSecret,omitempty"`
		// List of OAuth 2.0 scopes for your SMTP server. You must provide a scope that allows your email server to send emails.
		Scopes []string `json:"scopes,omitempty"`
		// The email provider's specific URL where the OAuth 2.0 app sends its credentials (or signed JWT) to exchange them for an access token
		TokenEndpoint *string `json:"tokenEndpoint,omitempty"`
		// This method determines how your OAuth 2.0 app sends its credentials (`client_id` and `client_secret`) to the provider's server when requesting an access token
		TokenEndpointAuthMethod *string `json:"tokenEndpointAuthMethod,omitempty"`
	}

	varOAUTH2CLIENTCREDENTIALSWithoutEmbeddedStruct := OAUTH2CLIENTCREDENTIALSWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOAUTH2CLIENTCREDENTIALSWithoutEmbeddedStruct)
	if err == nil {
		varOAUTH2CLIENTCREDENTIALS := _OAUTH2CLIENTCREDENTIALS{}
		varOAUTH2CLIENTCREDENTIALS.ClientId = varOAUTH2CLIENTCREDENTIALSWithoutEmbeddedStruct.ClientId
		varOAUTH2CLIENTCREDENTIALS.ClientSecret = varOAUTH2CLIENTCREDENTIALSWithoutEmbeddedStruct.ClientSecret
		varOAUTH2CLIENTCREDENTIALS.Scopes = varOAUTH2CLIENTCREDENTIALSWithoutEmbeddedStruct.Scopes
		varOAUTH2CLIENTCREDENTIALS.TokenEndpoint = varOAUTH2CLIENTCREDENTIALSWithoutEmbeddedStruct.TokenEndpoint
		varOAUTH2CLIENTCREDENTIALS.TokenEndpointAuthMethod = varOAUTH2CLIENTCREDENTIALSWithoutEmbeddedStruct.TokenEndpointAuthMethod
		*o = OAUTH2CLIENTCREDENTIALS(varOAUTH2CLIENTCREDENTIALS)
	} else {
		return err
	}

	varOAUTH2CLIENTCREDENTIALS := _OAUTH2CLIENTCREDENTIALS{}

	err = json.Unmarshal(data, &varOAUTH2CLIENTCREDENTIALS)
	if err == nil {
		o.BaseEmailServer = varOAUTH2CLIENTCREDENTIALS.BaseEmailServer
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "clientSecret")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "tokenEndpoint")
		delete(additionalProperties, "tokenEndpointAuthMethod")

		// remove fields from embedded structs
		reflectBaseEmailServer := reflect.ValueOf(o.BaseEmailServer)
		for i := 0; i < reflectBaseEmailServer.Type().NumField(); i++ {
			t := reflectBaseEmailServer.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAUTH2CLIENTCREDENTIALS struct {
	value *OAUTH2CLIENTCREDENTIALS
	isSet bool
}

func (v NullableOAUTH2CLIENTCREDENTIALS) Get() *OAUTH2CLIENTCREDENTIALS {
	return v.value
}

func (v *NullableOAUTH2CLIENTCREDENTIALS) Set(val *OAUTH2CLIENTCREDENTIALS) {
	v.value = val
	v.isSet = true
}

func (v NullableOAUTH2CLIENTCREDENTIALS) IsSet() bool {
	return v.isSet
}

func (v *NullableOAUTH2CLIENTCREDENTIALS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAUTH2CLIENTCREDENTIALS(val *OAUTH2CLIENTCREDENTIALS) *NullableOAUTH2CLIENTCREDENTIALS {
	return &NullableOAUTH2CLIENTCREDENTIALS{value: val, isSet: true}
}

func (v NullableOAUTH2CLIENTCREDENTIALS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAUTH2CLIENTCREDENTIALS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
