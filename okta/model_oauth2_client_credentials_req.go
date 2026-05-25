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
	"fmt"
	"reflect"
	"strings"
)

// checks if the OAUTH2CLIENTCREDENTIALSREQ type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAUTH2CLIENTCREDENTIALSREQ{}

// OAUTH2CLIENTCREDENTIALSREQ struct for OAUTH2CLIENTCREDENTIALSREQ
type OAUTH2CLIENTCREDENTIALSREQ struct {
	BaseEmailServer
	// The client ID that's used to access your SMTP server. This client ID is obtained when you create an OAuth 2.0 app with your email provider.
	ClientId string `json:"clientId"`
	// The client secret that's used to access your SMTP server. This client secret is obtained when you create an OAuth 2.0 app with your email provider.
	ClientSecret string `json:"clientSecret"`
	// List of OAuth 2.0 scopes for your SMTP server. You must provide a scope that allows your email server to send emails.
	Scopes []string `json:"scopes"`
	// The email provider's specific URL where the OAuth 2.0 app sends its credentials (or signed JWT) to exchange them for an access token
	TokenEndpoint string `json:"tokenEndpoint"`
	// This method determines how your OAuth 2.0 app sends its credentials (`client_id` and `client_secret`) to the provider's server when requesting an access token.
	TokenEndpointAuthMethod string `json:"tokenEndpointAuthMethod"`
	AdditionalProperties    map[string]interface{}
}

type _OAUTH2CLIENTCREDENTIALSREQ OAUTH2CLIENTCREDENTIALSREQ

// NewOAUTH2CLIENTCREDENTIALSREQ instantiates a new OAUTH2CLIENTCREDENTIALSREQ object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAUTH2CLIENTCREDENTIALSREQ(clientId string, clientSecret string, scopes []string, tokenEndpoint string, tokenEndpointAuthMethod string, alias string, authType string, enabled bool, host string, port int32, username string) *OAUTH2CLIENTCREDENTIALSREQ {
	this := OAUTH2CLIENTCREDENTIALSREQ{}
	this.Alias = alias
	this.AuthType = authType
	this.Enabled = enabled
	this.Host = host
	this.Port = port
	this.Username = username
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.Scopes = scopes
	this.TokenEndpoint = tokenEndpoint
	this.TokenEndpointAuthMethod = tokenEndpointAuthMethod
	return &this
}

// NewOAUTH2CLIENTCREDENTIALSREQWithDefaults instantiates a new OAUTH2CLIENTCREDENTIALSREQ object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAUTH2CLIENTCREDENTIALSREQWithDefaults() *OAUTH2CLIENTCREDENTIALSREQ {
	this := OAUTH2CLIENTCREDENTIALSREQ{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *OAUTH2CLIENTCREDENTIALSREQ) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *OAUTH2CLIENTCREDENTIALSREQ) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *OAUTH2CLIENTCREDENTIALSREQ) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *OAUTH2CLIENTCREDENTIALSREQ) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *OAUTH2CLIENTCREDENTIALSREQ) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *OAUTH2CLIENTCREDENTIALSREQ) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetScopes returns the Scopes field value
func (o *OAUTH2CLIENTCREDENTIALSREQ) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *OAUTH2CLIENTCREDENTIALSREQ) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *OAUTH2CLIENTCREDENTIALSREQ) SetScopes(v []string) {
	o.Scopes = v
}

// GetTokenEndpoint returns the TokenEndpoint field value
func (o *OAUTH2CLIENTCREDENTIALSREQ) GetTokenEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value
// and a boolean to check if the value has been set.
func (o *OAUTH2CLIENTCREDENTIALSREQ) GetTokenEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenEndpoint, true
}

// SetTokenEndpoint sets field value
func (o *OAUTH2CLIENTCREDENTIALSREQ) SetTokenEndpoint(v string) {
	o.TokenEndpoint = v
}

// GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field value
func (o *OAUTH2CLIENTCREDENTIALSREQ) GetTokenEndpointAuthMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenEndpointAuthMethod
}

// GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field value
// and a boolean to check if the value has been set.
func (o *OAUTH2CLIENTCREDENTIALSREQ) GetTokenEndpointAuthMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenEndpointAuthMethod, true
}

// SetTokenEndpointAuthMethod sets field value
func (o *OAUTH2CLIENTCREDENTIALSREQ) SetTokenEndpointAuthMethod(v string) {
	o.TokenEndpointAuthMethod = v
}

func (o OAUTH2CLIENTCREDENTIALSREQ) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAUTH2CLIENTCREDENTIALSREQ) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseEmailServer, errBaseEmailServer := json.Marshal(o.BaseEmailServer)
	if errBaseEmailServer != nil {
		return map[string]interface{}{}, errBaseEmailServer
	}
	errBaseEmailServer = json.Unmarshal([]byte(serializedBaseEmailServer), &toSerialize)
	if errBaseEmailServer != nil {
		return map[string]interface{}{}, errBaseEmailServer
	}
	toSerialize["clientId"] = o.ClientId
	toSerialize["clientSecret"] = o.ClientSecret
	toSerialize["scopes"] = o.Scopes
	toSerialize["tokenEndpoint"] = o.TokenEndpoint
	toSerialize["tokenEndpointAuthMethod"] = o.TokenEndpointAuthMethod

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAUTH2CLIENTCREDENTIALSREQ) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientId",
		"clientSecret",
		"scopes",
		"tokenEndpoint",
		"tokenEndpointAuthMethod",
		"alias",
		"authType",
		"enabled",
		"host",
		"port",
		"username",
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

	type OAUTH2CLIENTCREDENTIALSREQWithoutEmbeddedStruct struct {
		// The client ID that's used to access your SMTP server. This client ID is obtained when you create an OAuth 2.0 app with your email provider.
		ClientId string `json:"clientId"`
		// The client secret that's used to access your SMTP server. This client secret is obtained when you create an OAuth 2.0 app with your email provider.
		ClientSecret string `json:"clientSecret"`
		// List of OAuth 2.0 scopes for your SMTP server. You must provide a scope that allows your email server to send emails.
		Scopes []string `json:"scopes"`
		// The email provider's specific URL where the OAuth 2.0 app sends its credentials (or signed JWT) to exchange them for an access token
		TokenEndpoint string `json:"tokenEndpoint"`
		// This method determines how your OAuth 2.0 app sends its credentials (`client_id` and `client_secret`) to the provider's server when requesting an access token.
		TokenEndpointAuthMethod string `json:"tokenEndpointAuthMethod"`
	}

	varOAUTH2CLIENTCREDENTIALSREQWithoutEmbeddedStruct := OAUTH2CLIENTCREDENTIALSREQWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOAUTH2CLIENTCREDENTIALSREQWithoutEmbeddedStruct)
	if err == nil {
		varOAUTH2CLIENTCREDENTIALSREQ := _OAUTH2CLIENTCREDENTIALSREQ{}
		varOAUTH2CLIENTCREDENTIALSREQ.ClientId = varOAUTH2CLIENTCREDENTIALSREQWithoutEmbeddedStruct.ClientId
		varOAUTH2CLIENTCREDENTIALSREQ.ClientSecret = varOAUTH2CLIENTCREDENTIALSREQWithoutEmbeddedStruct.ClientSecret
		varOAUTH2CLIENTCREDENTIALSREQ.Scopes = varOAUTH2CLIENTCREDENTIALSREQWithoutEmbeddedStruct.Scopes
		varOAUTH2CLIENTCREDENTIALSREQ.TokenEndpoint = varOAUTH2CLIENTCREDENTIALSREQWithoutEmbeddedStruct.TokenEndpoint
		varOAUTH2CLIENTCREDENTIALSREQ.TokenEndpointAuthMethod = varOAUTH2CLIENTCREDENTIALSREQWithoutEmbeddedStruct.TokenEndpointAuthMethod
		*o = OAUTH2CLIENTCREDENTIALSREQ(varOAUTH2CLIENTCREDENTIALSREQ)
	} else {
		return err
	}

	varOAUTH2CLIENTCREDENTIALSREQ := _OAUTH2CLIENTCREDENTIALSREQ{}

	err = json.Unmarshal(data, &varOAUTH2CLIENTCREDENTIALSREQ)
	if err == nil {
		o.BaseEmailServer = varOAUTH2CLIENTCREDENTIALSREQ.BaseEmailServer
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

type NullableOAUTH2CLIENTCREDENTIALSREQ struct {
	value *OAUTH2CLIENTCREDENTIALSREQ
	isSet bool
}

func (v NullableOAUTH2CLIENTCREDENTIALSREQ) Get() *OAUTH2CLIENTCREDENTIALSREQ {
	return v.value
}

func (v *NullableOAUTH2CLIENTCREDENTIALSREQ) Set(val *OAUTH2CLIENTCREDENTIALSREQ) {
	v.value = val
	v.isSet = true
}

func (v NullableOAUTH2CLIENTCREDENTIALSREQ) IsSet() bool {
	return v.isSet
}

func (v *NullableOAUTH2CLIENTCREDENTIALSREQ) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAUTH2CLIENTCREDENTIALSREQ(val *OAUTH2CLIENTCREDENTIALSREQ) *NullableOAUTH2CLIENTCREDENTIALSREQ {
	return &NullableOAUTH2CLIENTCREDENTIALSREQ{value: val, isSet: true}
}

func (v NullableOAUTH2CLIENTCREDENTIALSREQ) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAUTH2CLIENTCREDENTIALSREQ) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
