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

// checks if the OAUTH2JWTBEARERGRANTCREATE type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAUTH2JWTBEARERGRANTCREATE{}

// OAUTH2JWTBEARERGRANTCREATE struct for OAUTH2JWTBEARERGRANTCREATE
type OAUTH2JWTBEARERGRANTCREATE struct {
	BaseEmailServerCreate
	// The URI of the authorization server that verifies the token. This is typically the token URI of your JWT.
	Audience string `json:"audience"`
	// The client ID that's used to access your SMTP server. This client ID is obtained when you create an OAuth 2.0 app with your email provider.
	ClientId string `json:"clientId"`
	// The unique ID of the entity that creates the JWT. This can sometimes be the email address of the user who creates the JWT. Check with your email provider for the correct value.
	Issuer string `json:"issuer"`
	// The ID of the private key that's used to sign the JWT
	KeyId string `json:"keyId"`
	// The secret RSA key that's used to cryptographically sign the JWT
	PrivateKey string `json:"privateKey"`
	// List of OAuth 2.0 scopes for your SMTP server. You must provide a scope that allows your email server to send emails.
	Scopes []string `json:"scopes"`
	// The signing algorithm that's used to sign the JWT
	SigningAlgorithm string `json:"signingAlgorithm"`
	// The email address of the user account that the OAuth 2.0 app impersonates to send emails
	Subject string `json:"subject"`
	// The email provider's specific URL where the OAuth 2.0 app sends its credentials (or signed JWT) to exchange them for an access token
	TokenEndpoint        string `json:"tokenEndpoint"`
	AdditionalProperties map[string]interface{}
}

type _OAUTH2JWTBEARERGRANTCREATE OAUTH2JWTBEARERGRANTCREATE

// NewOAUTH2JWTBEARERGRANTCREATE instantiates a new OAUTH2JWTBEARERGRANTCREATE object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAUTH2JWTBEARERGRANTCREATE(audience string, clientId string, issuer string, keyId string, privateKey string, scopes []string, signingAlgorithm string, subject string, tokenEndpoint string, alias string, enabled bool, host string, port int32, username string) *OAUTH2JWTBEARERGRANTCREATE {
	this := OAUTH2JWTBEARERGRANTCREATE{}
	this.Alias = alias
	this.Enabled = enabled
	this.Host = host
	this.Port = port
	this.Username = username
	this.Audience = audience
	this.ClientId = clientId
	this.Issuer = issuer
	this.KeyId = keyId
	this.PrivateKey = privateKey
	this.Scopes = scopes
	this.SigningAlgorithm = signingAlgorithm
	this.Subject = subject
	this.TokenEndpoint = tokenEndpoint
	return &this
}

// NewOAUTH2JWTBEARERGRANTCREATEWithDefaults instantiates a new OAUTH2JWTBEARERGRANTCREATE object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAUTH2JWTBEARERGRANTCREATEWithDefaults() *OAUTH2JWTBEARERGRANTCREATE {
	this := OAUTH2JWTBEARERGRANTCREATE{}
	return &this
}

// GetAudience returns the Audience field value
func (o *OAUTH2JWTBEARERGRANTCREATE) GetAudience() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANTCREATE) GetAudienceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audience, true
}

// SetAudience sets field value
func (o *OAUTH2JWTBEARERGRANTCREATE) SetAudience(v string) {
	o.Audience = v
}

// GetClientId returns the ClientId field value
func (o *OAUTH2JWTBEARERGRANTCREATE) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANTCREATE) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *OAUTH2JWTBEARERGRANTCREATE) SetClientId(v string) {
	o.ClientId = v
}

// GetIssuer returns the Issuer field value
func (o *OAUTH2JWTBEARERGRANTCREATE) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANTCREATE) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *OAUTH2JWTBEARERGRANTCREATE) SetIssuer(v string) {
	o.Issuer = v
}

// GetKeyId returns the KeyId field value
func (o *OAUTH2JWTBEARERGRANTCREATE) GetKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANTCREATE) GetKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyId, true
}

// SetKeyId sets field value
func (o *OAUTH2JWTBEARERGRANTCREATE) SetKeyId(v string) {
	o.KeyId = v
}

// GetPrivateKey returns the PrivateKey field value
func (o *OAUTH2JWTBEARERGRANTCREATE) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANTCREATE) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *OAUTH2JWTBEARERGRANTCREATE) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetScopes returns the Scopes field value
func (o *OAUTH2JWTBEARERGRANTCREATE) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANTCREATE) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *OAUTH2JWTBEARERGRANTCREATE) SetScopes(v []string) {
	o.Scopes = v
}

// GetSigningAlgorithm returns the SigningAlgorithm field value
func (o *OAUTH2JWTBEARERGRANTCREATE) GetSigningAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SigningAlgorithm
}

// GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field value
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANTCREATE) GetSigningAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SigningAlgorithm, true
}

// SetSigningAlgorithm sets field value
func (o *OAUTH2JWTBEARERGRANTCREATE) SetSigningAlgorithm(v string) {
	o.SigningAlgorithm = v
}

// GetSubject returns the Subject field value
func (o *OAUTH2JWTBEARERGRANTCREATE) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANTCREATE) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *OAUTH2JWTBEARERGRANTCREATE) SetSubject(v string) {
	o.Subject = v
}

// GetTokenEndpoint returns the TokenEndpoint field value
func (o *OAUTH2JWTBEARERGRANTCREATE) GetTokenEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANTCREATE) GetTokenEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenEndpoint, true
}

// SetTokenEndpoint sets field value
func (o *OAUTH2JWTBEARERGRANTCREATE) SetTokenEndpoint(v string) {
	o.TokenEndpoint = v
}

func (o OAUTH2JWTBEARERGRANTCREATE) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAUTH2JWTBEARERGRANTCREATE) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseEmailServerCreate, errBaseEmailServerCreate := json.Marshal(o.BaseEmailServerCreate)
	if errBaseEmailServerCreate != nil {
		return map[string]interface{}{}, errBaseEmailServerCreate
	}
	errBaseEmailServerCreate = json.Unmarshal([]byte(serializedBaseEmailServerCreate), &toSerialize)
	if errBaseEmailServerCreate != nil {
		return map[string]interface{}{}, errBaseEmailServerCreate
	}
	toSerialize["audience"] = o.Audience
	toSerialize["clientId"] = o.ClientId
	toSerialize["issuer"] = o.Issuer
	toSerialize["keyId"] = o.KeyId
	toSerialize["privateKey"] = o.PrivateKey
	toSerialize["scopes"] = o.Scopes
	toSerialize["signingAlgorithm"] = o.SigningAlgorithm
	toSerialize["subject"] = o.Subject
	toSerialize["tokenEndpoint"] = o.TokenEndpoint

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAUTH2JWTBEARERGRANTCREATE) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"audience",
		"clientId",
		"issuer",
		"keyId",
		"privateKey",
		"scopes",
		"signingAlgorithm",
		"subject",
		"tokenEndpoint",
		"alias",
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

	type OAUTH2JWTBEARERGRANTCREATEWithoutEmbeddedStruct struct {
		// The URI of the authorization server that verifies the token. This is typically the token URI of your JWT.
		Audience string `json:"audience"`
		// The client ID that's used to access your SMTP server. This client ID is obtained when you create an OAuth 2.0 app with your email provider.
		ClientId string `json:"clientId"`
		// The unique ID of the entity that creates the JWT. This can sometimes be the email address of the user who creates the JWT. Check with your email provider for the correct value.
		Issuer string `json:"issuer"`
		// The ID of the private key that's used to sign the JWT
		KeyId string `json:"keyId"`
		// The secret RSA key that's used to cryptographically sign the JWT
		PrivateKey string `json:"privateKey"`
		// List of OAuth 2.0 scopes for your SMTP server. You must provide a scope that allows your email server to send emails.
		Scopes []string `json:"scopes"`
		// The signing algorithm that's used to sign the JWT
		SigningAlgorithm string `json:"signingAlgorithm"`
		// The email address of the user account that the OAuth 2.0 app impersonates to send emails
		Subject string `json:"subject"`
		// The email provider's specific URL where the OAuth 2.0 app sends its credentials (or signed JWT) to exchange them for an access token
		TokenEndpoint string `json:"tokenEndpoint"`
	}

	varOAUTH2JWTBEARERGRANTCREATEWithoutEmbeddedStruct := OAUTH2JWTBEARERGRANTCREATEWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOAUTH2JWTBEARERGRANTCREATEWithoutEmbeddedStruct)
	if err == nil {
		varOAUTH2JWTBEARERGRANTCREATE := _OAUTH2JWTBEARERGRANTCREATE{}
		varOAUTH2JWTBEARERGRANTCREATE.Audience = varOAUTH2JWTBEARERGRANTCREATEWithoutEmbeddedStruct.Audience
		varOAUTH2JWTBEARERGRANTCREATE.ClientId = varOAUTH2JWTBEARERGRANTCREATEWithoutEmbeddedStruct.ClientId
		varOAUTH2JWTBEARERGRANTCREATE.Issuer = varOAUTH2JWTBEARERGRANTCREATEWithoutEmbeddedStruct.Issuer
		varOAUTH2JWTBEARERGRANTCREATE.KeyId = varOAUTH2JWTBEARERGRANTCREATEWithoutEmbeddedStruct.KeyId
		varOAUTH2JWTBEARERGRANTCREATE.PrivateKey = varOAUTH2JWTBEARERGRANTCREATEWithoutEmbeddedStruct.PrivateKey
		varOAUTH2JWTBEARERGRANTCREATE.Scopes = varOAUTH2JWTBEARERGRANTCREATEWithoutEmbeddedStruct.Scopes
		varOAUTH2JWTBEARERGRANTCREATE.SigningAlgorithm = varOAUTH2JWTBEARERGRANTCREATEWithoutEmbeddedStruct.SigningAlgorithm
		varOAUTH2JWTBEARERGRANTCREATE.Subject = varOAUTH2JWTBEARERGRANTCREATEWithoutEmbeddedStruct.Subject
		varOAUTH2JWTBEARERGRANTCREATE.TokenEndpoint = varOAUTH2JWTBEARERGRANTCREATEWithoutEmbeddedStruct.TokenEndpoint
		*o = OAUTH2JWTBEARERGRANTCREATE(varOAUTH2JWTBEARERGRANTCREATE)
	} else {
		return err
	}

	varOAUTH2JWTBEARERGRANTCREATE := _OAUTH2JWTBEARERGRANTCREATE{}

	err = json.Unmarshal(data, &varOAUTH2JWTBEARERGRANTCREATE)
	if err == nil {
		o.BaseEmailServerCreate = varOAUTH2JWTBEARERGRANTCREATE.BaseEmailServerCreate
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "audience")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "keyId")
		delete(additionalProperties, "privateKey")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "signingAlgorithm")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "tokenEndpoint")

		// remove fields from embedded structs
		reflectBaseEmailServerCreate := reflect.ValueOf(o.BaseEmailServerCreate)
		for i := 0; i < reflectBaseEmailServerCreate.Type().NumField(); i++ {
			t := reflectBaseEmailServerCreate.Type().Field(i)

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

type NullableOAUTH2JWTBEARERGRANTCREATE struct {
	value *OAUTH2JWTBEARERGRANTCREATE
	isSet bool
}

func (v NullableOAUTH2JWTBEARERGRANTCREATE) Get() *OAUTH2JWTBEARERGRANTCREATE {
	return v.value
}

func (v *NullableOAUTH2JWTBEARERGRANTCREATE) Set(val *OAUTH2JWTBEARERGRANTCREATE) {
	v.value = val
	v.isSet = true
}

func (v NullableOAUTH2JWTBEARERGRANTCREATE) IsSet() bool {
	return v.isSet
}

func (v *NullableOAUTH2JWTBEARERGRANTCREATE) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAUTH2JWTBEARERGRANTCREATE(val *OAUTH2JWTBEARERGRANTCREATE) *NullableOAUTH2JWTBEARERGRANTCREATE {
	return &NullableOAUTH2JWTBEARERGRANTCREATE{value: val, isSet: true}
}

func (v NullableOAUTH2JWTBEARERGRANTCREATE) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAUTH2JWTBEARERGRANTCREATE) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
