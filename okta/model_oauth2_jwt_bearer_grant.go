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

// checks if the OAUTH2JWTBEARERGRANT type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAUTH2JWTBEARERGRANT{}

// OAUTH2JWTBEARERGRANT struct for OAUTH2JWTBEARERGRANT
type OAUTH2JWTBEARERGRANT struct {
	BaseEmailServer
	// The URI of the authorization server that verifies the token. This is typically the token URI of your JWT.
	Audience *string `json:"audience,omitempty"`
	// The client ID that's used to access your SMTP server. This client ID is obtained when you create an OAuth 2.0 app with your email provider.
	ClientId *string `json:"clientId,omitempty"`
	// The unique ID of the entity that creates the JWT. This can sometimes be the email address of the user who creates the JWT. Check with your email provider for the correct value.
	Issuer *string `json:"issuer,omitempty"`
	// The ID of the private key that's used to sign the JWT
	KeyId *string `json:"keyId,omitempty"`
	// The secret RSA key that's used to cryptographically sign the JWT
	PrivateKey *string `json:"privateKey,omitempty"`
	// List of OAuth 2.0 scopes for your SMTP server. You must provide a scope that allows your email server to send emails.
	Scopes []string `json:"scopes,omitempty"`
	// The signing algorithm that's used to sign the JWT
	SigningAlgorithm *string `json:"signingAlgorithm,omitempty"`
	// The email address of the user account that the OAuth 2.0 app impersonates to send emails
	Subject *string `json:"subject,omitempty"`
	// The email provider's specific URL where the OAuth 2.0 app sends its credentials (or signed JWT) to exchange them for an access token
	TokenEndpoint        *string `json:"tokenEndpoint,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAUTH2JWTBEARERGRANT OAUTH2JWTBEARERGRANT

// NewOAUTH2JWTBEARERGRANT instantiates a new OAUTH2JWTBEARERGRANT object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAUTH2JWTBEARERGRANT() *OAUTH2JWTBEARERGRANT {
	this := OAUTH2JWTBEARERGRANT{}
	return &this
}

// NewOAUTH2JWTBEARERGRANTWithDefaults instantiates a new OAUTH2JWTBEARERGRANT object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAUTH2JWTBEARERGRANTWithDefaults() *OAUTH2JWTBEARERGRANT {
	this := OAUTH2JWTBEARERGRANT{}
	return &this
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *OAUTH2JWTBEARERGRANT) GetAudience() string {
	if o == nil || IsNil(o.Audience) {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANT) GetAudienceOk() (*string, bool) {
	if o == nil || IsNil(o.Audience) {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *OAUTH2JWTBEARERGRANT) HasAudience() bool {
	if o != nil && !IsNil(o.Audience) {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *OAUTH2JWTBEARERGRANT) SetAudience(v string) {
	o.Audience = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OAUTH2JWTBEARERGRANT) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANT) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OAUTH2JWTBEARERGRANT) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OAUTH2JWTBEARERGRANT) SetClientId(v string) {
	o.ClientId = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *OAUTH2JWTBEARERGRANT) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANT) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *OAUTH2JWTBEARERGRANT) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *OAUTH2JWTBEARERGRANT) SetIssuer(v string) {
	o.Issuer = &v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *OAUTH2JWTBEARERGRANT) GetKeyId() string {
	if o == nil || IsNil(o.KeyId) {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANT) GetKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeyId) {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *OAUTH2JWTBEARERGRANT) HasKeyId() bool {
	if o != nil && !IsNil(o.KeyId) {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *OAUTH2JWTBEARERGRANT) SetKeyId(v string) {
	o.KeyId = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *OAUTH2JWTBEARERGRANT) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANT) GetPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *OAUTH2JWTBEARERGRANT) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *OAUTH2JWTBEARERGRANT) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *OAUTH2JWTBEARERGRANT) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANT) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *OAUTH2JWTBEARERGRANT) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *OAUTH2JWTBEARERGRANT) SetScopes(v []string) {
	o.Scopes = v
}

// GetSigningAlgorithm returns the SigningAlgorithm field value if set, zero value otherwise.
func (o *OAUTH2JWTBEARERGRANT) GetSigningAlgorithm() string {
	if o == nil || IsNil(o.SigningAlgorithm) {
		var ret string
		return ret
	}
	return *o.SigningAlgorithm
}

// GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANT) GetSigningAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.SigningAlgorithm) {
		return nil, false
	}
	return o.SigningAlgorithm, true
}

// HasSigningAlgorithm returns a boolean if a field has been set.
func (o *OAUTH2JWTBEARERGRANT) HasSigningAlgorithm() bool {
	if o != nil && !IsNil(o.SigningAlgorithm) {
		return true
	}

	return false
}

// SetSigningAlgorithm gets a reference to the given string and assigns it to the SigningAlgorithm field.
func (o *OAUTH2JWTBEARERGRANT) SetSigningAlgorithm(v string) {
	o.SigningAlgorithm = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *OAUTH2JWTBEARERGRANT) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANT) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *OAUTH2JWTBEARERGRANT) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *OAUTH2JWTBEARERGRANT) SetSubject(v string) {
	o.Subject = &v
}

// GetTokenEndpoint returns the TokenEndpoint field value if set, zero value otherwise.
func (o *OAUTH2JWTBEARERGRANT) GetTokenEndpoint() string {
	if o == nil || IsNil(o.TokenEndpoint) {
		var ret string
		return ret
	}
	return *o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAUTH2JWTBEARERGRANT) GetTokenEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.TokenEndpoint) {
		return nil, false
	}
	return o.TokenEndpoint, true
}

// HasTokenEndpoint returns a boolean if a field has been set.
func (o *OAUTH2JWTBEARERGRANT) HasTokenEndpoint() bool {
	if o != nil && !IsNil(o.TokenEndpoint) {
		return true
	}

	return false
}

// SetTokenEndpoint gets a reference to the given string and assigns it to the TokenEndpoint field.
func (o *OAUTH2JWTBEARERGRANT) SetTokenEndpoint(v string) {
	o.TokenEndpoint = &v
}

func (o OAUTH2JWTBEARERGRANT) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAUTH2JWTBEARERGRANT) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseEmailServer, errBaseEmailServer := json.Marshal(o.BaseEmailServer)
	if errBaseEmailServer != nil {
		return map[string]interface{}{}, errBaseEmailServer
	}
	errBaseEmailServer = json.Unmarshal([]byte(serializedBaseEmailServer), &toSerialize)
	if errBaseEmailServer != nil {
		return map[string]interface{}{}, errBaseEmailServer
	}
	if !IsNil(o.Audience) {
		toSerialize["audience"] = o.Audience
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.KeyId) {
		toSerialize["keyId"] = o.KeyId
	}
	if !IsNil(o.PrivateKey) {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.SigningAlgorithm) {
		toSerialize["signingAlgorithm"] = o.SigningAlgorithm
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.TokenEndpoint) {
		toSerialize["tokenEndpoint"] = o.TokenEndpoint
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAUTH2JWTBEARERGRANT) UnmarshalJSON(data []byte) (err error) {
	type OAUTH2JWTBEARERGRANTWithoutEmbeddedStruct struct {
		// The URI of the authorization server that verifies the token. This is typically the token URI of your JWT.
		Audience *string `json:"audience,omitempty"`
		// The client ID that's used to access your SMTP server. This client ID is obtained when you create an OAuth 2.0 app with your email provider.
		ClientId *string `json:"clientId,omitempty"`
		// The unique ID of the entity that creates the JWT. This can sometimes be the email address of the user who creates the JWT. Check with your email provider for the correct value.
		Issuer *string `json:"issuer,omitempty"`
		// The ID of the private key that's used to sign the JWT
		KeyId *string `json:"keyId,omitempty"`
		// The secret RSA key that's used to cryptographically sign the JWT
		PrivateKey *string `json:"privateKey,omitempty"`
		// List of OAuth 2.0 scopes for your SMTP server. You must provide a scope that allows your email server to send emails.
		Scopes []string `json:"scopes,omitempty"`
		// The signing algorithm that's used to sign the JWT
		SigningAlgorithm *string `json:"signingAlgorithm,omitempty"`
		// The email address of the user account that the OAuth 2.0 app impersonates to send emails
		Subject *string `json:"subject,omitempty"`
		// The email provider's specific URL where the OAuth 2.0 app sends its credentials (or signed JWT) to exchange them for an access token
		TokenEndpoint *string `json:"tokenEndpoint,omitempty"`
	}

	varOAUTH2JWTBEARERGRANTWithoutEmbeddedStruct := OAUTH2JWTBEARERGRANTWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOAUTH2JWTBEARERGRANTWithoutEmbeddedStruct)
	if err == nil {
		varOAUTH2JWTBEARERGRANT := _OAUTH2JWTBEARERGRANT{}
		varOAUTH2JWTBEARERGRANT.Audience = varOAUTH2JWTBEARERGRANTWithoutEmbeddedStruct.Audience
		varOAUTH2JWTBEARERGRANT.ClientId = varOAUTH2JWTBEARERGRANTWithoutEmbeddedStruct.ClientId
		varOAUTH2JWTBEARERGRANT.Issuer = varOAUTH2JWTBEARERGRANTWithoutEmbeddedStruct.Issuer
		varOAUTH2JWTBEARERGRANT.KeyId = varOAUTH2JWTBEARERGRANTWithoutEmbeddedStruct.KeyId
		varOAUTH2JWTBEARERGRANT.PrivateKey = varOAUTH2JWTBEARERGRANTWithoutEmbeddedStruct.PrivateKey
		varOAUTH2JWTBEARERGRANT.Scopes = varOAUTH2JWTBEARERGRANTWithoutEmbeddedStruct.Scopes
		varOAUTH2JWTBEARERGRANT.SigningAlgorithm = varOAUTH2JWTBEARERGRANTWithoutEmbeddedStruct.SigningAlgorithm
		varOAUTH2JWTBEARERGRANT.Subject = varOAUTH2JWTBEARERGRANTWithoutEmbeddedStruct.Subject
		varOAUTH2JWTBEARERGRANT.TokenEndpoint = varOAUTH2JWTBEARERGRANTWithoutEmbeddedStruct.TokenEndpoint
		*o = OAUTH2JWTBEARERGRANT(varOAUTH2JWTBEARERGRANT)
	} else {
		return err
	}

	varOAUTH2JWTBEARERGRANT := _OAUTH2JWTBEARERGRANT{}

	err = json.Unmarshal(data, &varOAUTH2JWTBEARERGRANT)
	if err == nil {
		o.BaseEmailServer = varOAUTH2JWTBEARERGRANT.BaseEmailServer
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

type NullableOAUTH2JWTBEARERGRANT struct {
	value *OAUTH2JWTBEARERGRANT
	isSet bool
}

func (v NullableOAUTH2JWTBEARERGRANT) Get() *OAUTH2JWTBEARERGRANT {
	return v.value
}

func (v *NullableOAUTH2JWTBEARERGRANT) Set(val *OAUTH2JWTBEARERGRANT) {
	v.value = val
	v.isSet = true
}

func (v NullableOAUTH2JWTBEARERGRANT) IsSet() bool {
	return v.isSet
}

func (v *NullableOAUTH2JWTBEARERGRANT) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAUTH2JWTBEARERGRANT(val *OAUTH2JWTBEARERGRANT) *NullableOAUTH2JWTBEARERGRANT {
	return &NullableOAUTH2JWTBEARERGRANT{value: val, isSet: true}
}

func (v NullableOAUTH2JWTBEARERGRANT) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAUTH2JWTBEARERGRANT) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
