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
	"time"
)

// checks if the AuthorizationServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationServer{}

// AuthorizationServer struct for AuthorizationServer
type AuthorizationServer struct {
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>The algorithm for encrypting access tokens issued by this authorization server. If this is requested, the response is signed, and then encrypted. The result is a nested JWT. The default, if omitted, is that no encryption is performed.
	AccessTokenEncryptedResponseAlgorithm *string `json:"accessTokenEncryptedResponseAlgorithm,omitempty"`
	// The recipients that the tokens are intended for. This becomes the `aud` claim in an access token. Okta currently supports only one audience.
	Audiences   []string                        `json:"audiences,omitempty"`
	Created     *time.Time                      `json:"created,omitempty"`
	Credentials *AuthorizationServerCredentials `json:"credentials,omitempty"`
	// The description of the custom authorization server
	Description *string `json:"description,omitempty"`
	// The ID of the custom authorization server
	Id *string `json:"id,omitempty"`
	// The complete URL for the custom authorization server. This becomes the `iss` claim in an access token.
	Issuer *string `json:"issuer,omitempty"`
	// Indicates which value is specified in the issuer of the tokens that a custom authorization server returns: the Okta org domain URL or a custom domain URL.  `issuerMode` is visible if you have a custom URL domain configured or the Dynamic Issuer Mode feature enabled. If you have a custom URL domain configured, you can set a custom domain URL in a custom authorization server, and this property is returned in the appropriate responses.  When set to `ORG_URL`, then in responses, `issuer` is the Okta org domain URL: `https://${yourOktaDomain}`.  When set to `CUSTOM_URL`, then in responses, `issuer` is the custom domain URL configured in the administration user interface.  When set to `DYNAMIC`, then in responses, `issuer` is the custom domain URL if the OAuth 2.0 request was sent to the custom domain, or is the Okta org's domain URL if the OAuth 2.0 request was sent to the original Okta org domain.  After you configure a custom URL domain, all new custom authorization servers use `CUSTOM_URL` by default. If the Dynamic Issuer Mode feature is enabled, then all new custom authorization servers use `DYNAMIC` by default. All existing custom authorization servers continue to use the original value until they're changed using the Admin Console or the API. This way, existing integrations with the client and resource server continue to work after the feature is enabled.
	IssuerMode *string                    `json:"issuerMode,omitempty"`
	Jwks       *ResourceServerJsonWebKeys `json:"jwks,omitempty"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>URL string that references a JSON Web Key Set for encrypting JWTs minted by the custom authorization server
	JwksUri     *string    `json:"jwks_uri,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The name of the custom authorization server
	Name                 *string          `json:"name,omitempty"`
	Status               *string          `json:"status,omitempty"`
	Links                *AuthServerLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServer AuthorizationServer

// NewAuthorizationServer instantiates a new AuthorizationServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServer() *AuthorizationServer {
	this := AuthorizationServer{}
	return &this
}

// NewAuthorizationServerWithDefaults instantiates a new AuthorizationServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerWithDefaults() *AuthorizationServer {
	this := AuthorizationServer{}
	return &this
}

// GetAccessTokenEncryptedResponseAlgorithm returns the AccessTokenEncryptedResponseAlgorithm field value if set, zero value otherwise.
func (o *AuthorizationServer) GetAccessTokenEncryptedResponseAlgorithm() string {
	if o == nil || IsNil(o.AccessTokenEncryptedResponseAlgorithm) {
		var ret string
		return ret
	}
	return *o.AccessTokenEncryptedResponseAlgorithm
}

// GetAccessTokenEncryptedResponseAlgorithmOk returns a tuple with the AccessTokenEncryptedResponseAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetAccessTokenEncryptedResponseAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.AccessTokenEncryptedResponseAlgorithm) {
		return nil, false
	}
	return o.AccessTokenEncryptedResponseAlgorithm, true
}

// HasAccessTokenEncryptedResponseAlgorithm returns a boolean if a field has been set.
func (o *AuthorizationServer) HasAccessTokenEncryptedResponseAlgorithm() bool {
	if o != nil && !IsNil(o.AccessTokenEncryptedResponseAlgorithm) {
		return true
	}

	return false
}

// SetAccessTokenEncryptedResponseAlgorithm gets a reference to the given string and assigns it to the AccessTokenEncryptedResponseAlgorithm field.
func (o *AuthorizationServer) SetAccessTokenEncryptedResponseAlgorithm(v string) {
	o.AccessTokenEncryptedResponseAlgorithm = &v
}

// GetAudiences returns the Audiences field value if set, zero value otherwise.
func (o *AuthorizationServer) GetAudiences() []string {
	if o == nil || IsNil(o.Audiences) {
		var ret []string
		return ret
	}
	return o.Audiences
}

// GetAudiencesOk returns a tuple with the Audiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetAudiencesOk() ([]string, bool) {
	if o == nil || IsNil(o.Audiences) {
		return nil, false
	}
	return o.Audiences, true
}

// HasAudiences returns a boolean if a field has been set.
func (o *AuthorizationServer) HasAudiences() bool {
	if o != nil && !IsNil(o.Audiences) {
		return true
	}

	return false
}

// SetAudiences gets a reference to the given []string and assigns it to the Audiences field.
func (o *AuthorizationServer) SetAudiences(v []string) {
	o.Audiences = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AuthorizationServer) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AuthorizationServer) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AuthorizationServer) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *AuthorizationServer) GetCredentials() AuthorizationServerCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret AuthorizationServerCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetCredentialsOk() (*AuthorizationServerCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *AuthorizationServer) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given AuthorizationServerCredentials and assigns it to the Credentials field.
func (o *AuthorizationServer) SetCredentials(v AuthorizationServerCredentials) {
	o.Credentials = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthorizationServer) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthorizationServer) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthorizationServer) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthorizationServer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthorizationServer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthorizationServer) SetId(v string) {
	o.Id = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *AuthorizationServer) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *AuthorizationServer) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *AuthorizationServer) SetIssuer(v string) {
	o.Issuer = &v
}

// GetIssuerMode returns the IssuerMode field value if set, zero value otherwise.
func (o *AuthorizationServer) GetIssuerMode() string {
	if o == nil || IsNil(o.IssuerMode) {
		var ret string
		return ret
	}
	return *o.IssuerMode
}

// GetIssuerModeOk returns a tuple with the IssuerMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetIssuerModeOk() (*string, bool) {
	if o == nil || IsNil(o.IssuerMode) {
		return nil, false
	}
	return o.IssuerMode, true
}

// HasIssuerMode returns a boolean if a field has been set.
func (o *AuthorizationServer) HasIssuerMode() bool {
	if o != nil && !IsNil(o.IssuerMode) {
		return true
	}

	return false
}

// SetIssuerMode gets a reference to the given string and assigns it to the IssuerMode field.
func (o *AuthorizationServer) SetIssuerMode(v string) {
	o.IssuerMode = &v
}

// GetJwks returns the Jwks field value if set, zero value otherwise.
func (o *AuthorizationServer) GetJwks() ResourceServerJsonWebKeys {
	if o == nil || IsNil(o.Jwks) {
		var ret ResourceServerJsonWebKeys
		return ret
	}
	return *o.Jwks
}

// GetJwksOk returns a tuple with the Jwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetJwksOk() (*ResourceServerJsonWebKeys, bool) {
	if o == nil || IsNil(o.Jwks) {
		return nil, false
	}
	return o.Jwks, true
}

// HasJwks returns a boolean if a field has been set.
func (o *AuthorizationServer) HasJwks() bool {
	if o != nil && !IsNil(o.Jwks) {
		return true
	}

	return false
}

// SetJwks gets a reference to the given ResourceServerJsonWebKeys and assigns it to the Jwks field.
func (o *AuthorizationServer) SetJwks(v ResourceServerJsonWebKeys) {
	o.Jwks = &v
}

// GetJwksUri returns the JwksUri field value if set, zero value otherwise.
func (o *AuthorizationServer) GetJwksUri() string {
	if o == nil || IsNil(o.JwksUri) {
		var ret string
		return ret
	}
	return *o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetJwksUriOk() (*string, bool) {
	if o == nil || IsNil(o.JwksUri) {
		return nil, false
	}
	return o.JwksUri, true
}

// HasJwksUri returns a boolean if a field has been set.
func (o *AuthorizationServer) HasJwksUri() bool {
	if o != nil && !IsNil(o.JwksUri) {
		return true
	}

	return false
}

// SetJwksUri gets a reference to the given string and assigns it to the JwksUri field.
func (o *AuthorizationServer) SetJwksUri(v string) {
	o.JwksUri = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AuthorizationServer) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AuthorizationServer) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AuthorizationServer) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthorizationServer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthorizationServer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthorizationServer) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AuthorizationServer) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AuthorizationServer) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AuthorizationServer) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AuthorizationServer) GetLinks() AuthServerLinks {
	if o == nil || IsNil(o.Links) {
		var ret AuthServerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetLinksOk() (*AuthServerLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AuthorizationServer) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AuthServerLinks and assigns it to the Links field.
func (o *AuthorizationServer) SetLinks(v AuthServerLinks) {
	o.Links = &v
}

func (o AuthorizationServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessTokenEncryptedResponseAlgorithm) {
		toSerialize["accessTokenEncryptedResponseAlgorithm"] = o.AccessTokenEncryptedResponseAlgorithm
	}
	if !IsNil(o.Audiences) {
		toSerialize["audiences"] = o.Audiences
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.IssuerMode) {
		toSerialize["issuerMode"] = o.IssuerMode
	}
	if !IsNil(o.Jwks) {
		toSerialize["jwks"] = o.Jwks
	}
	if !IsNil(o.JwksUri) {
		toSerialize["jwks_uri"] = o.JwksUri
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthorizationServer) UnmarshalJSON(data []byte) (err error) {
	varAuthorizationServer := _AuthorizationServer{}

	err = json.Unmarshal(data, &varAuthorizationServer)

	if err != nil {
		return err
	}

	*o = AuthorizationServer(varAuthorizationServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessTokenEncryptedResponseAlgorithm")
		delete(additionalProperties, "audiences")
		delete(additionalProperties, "created")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "issuerMode")
		delete(additionalProperties, "jwks")
		delete(additionalProperties, "jwks_uri")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthorizationServer struct {
	value *AuthorizationServer
	isSet bool
}

func (v NullableAuthorizationServer) Get() *AuthorizationServer {
	return v.value
}

func (v *NullableAuthorizationServer) Set(val *AuthorizationServer) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServer) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServer(val *AuthorizationServer) *NullableAuthorizationServer {
	return &NullableAuthorizationServer{value: val, isSet: true}
}

func (v NullableAuthorizationServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
