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
	"fmt"
)

// checks if the OAuth2ClientJsonSigningKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ClientJsonSigningKeyResponse{}

// OAuth2ClientJsonSigningKeyResponse A [JSON Web Key (JWK)](https://tools.ietf.org/html/rfc7517) is a JSON representation of a cryptographic key. Okta uses signing keys to verify the signature of a JWT when provided for the `private_key_jwt` client authentication method or for a signed authorize request object. Okta supports both RSA and Elliptic Curve (EC) keys for signing tokens.
type OAuth2ClientJsonSigningKeyResponse struct {
	// Algorithm used in the key
	Alg string `json:"alg"`
	// Timestamp when the OAuth 2.0 client JSON Web Key was created
	Created string `json:"created"`
	// The unique ID of the OAuth Client JSON Web Key
	Id string `json:"id"`
	// Unique identifier of the JSON Web Key in the OAuth 2.0 client's JWKS
	Kid NullableString `json:"kid,omitempty"`
	// Cryptographic algorithm family for the certificate's key pair
	Kty string `json:"kty"`
	// Timestamp when the OAuth 2.0 client JSON Web Key was updated
	LastUpdated string `json:"lastUpdated"`
	// Status of the OAuth 2.0 client JSON Web Key
	Status *string `json:"status,omitempty"`
	// Acceptable use of the JSON Web Key
	Use                  string                  `json:"use"`
	Links                *OAuthClientSecretLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientJsonSigningKeyResponse OAuth2ClientJsonSigningKeyResponse

// NewOAuth2ClientJsonSigningKeyResponse instantiates a new OAuth2ClientJsonSigningKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientJsonSigningKeyResponse(alg string, created string, id string, kty string, lastUpdated string, use string) *OAuth2ClientJsonSigningKeyResponse {
	this := OAuth2ClientJsonSigningKeyResponse{}
	this.Alg = alg
	this.Created = created
	this.Id = id
	this.Kty = kty
	this.LastUpdated = lastUpdated
	this.Use = use
	return &this
}

// NewOAuth2ClientJsonSigningKeyResponseWithDefaults instantiates a new OAuth2ClientJsonSigningKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientJsonSigningKeyResponseWithDefaults() *OAuth2ClientJsonSigningKeyResponse {
	this := OAuth2ClientJsonSigningKeyResponse{}
	return &this
}

// GetAlg returns the Alg field value
func (o *OAuth2ClientJsonSigningKeyResponse) GetAlg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alg
}

// GetAlgOk returns a tuple with the Alg field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonSigningKeyResponse) GetAlgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alg, true
}

// SetAlg sets field value
func (o *OAuth2ClientJsonSigningKeyResponse) SetAlg(v string) {
	o.Alg = v
}

// GetCreated returns the Created field value
func (o *OAuth2ClientJsonSigningKeyResponse) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonSigningKeyResponse) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *OAuth2ClientJsonSigningKeyResponse) SetCreated(v string) {
	o.Created = v
}

// GetId returns the Id field value
func (o *OAuth2ClientJsonSigningKeyResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonSigningKeyResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OAuth2ClientJsonSigningKeyResponse) SetId(v string) {
	o.Id = v
}

// GetKid returns the Kid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientJsonSigningKeyResponse) GetKid() string {
	if o == nil || IsNil(o.Kid.Get()) {
		var ret string
		return ret
	}
	return *o.Kid.Get()
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientJsonSigningKeyResponse) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kid.Get(), o.Kid.IsSet()
}

// HasKid returns a boolean if a field has been set.
func (o *OAuth2ClientJsonSigningKeyResponse) HasKid() bool {
	if o != nil && o.Kid.IsSet() {
		return true
	}

	return false
}

// SetKid gets a reference to the given NullableString and assigns it to the Kid field.
func (o *OAuth2ClientJsonSigningKeyResponse) SetKid(v string) {
	o.Kid.Set(&v)
}

// SetKidNil sets the value for Kid to be an explicit nil
func (o *OAuth2ClientJsonSigningKeyResponse) SetKidNil() {
	o.Kid.Set(nil)
}

// UnsetKid ensures that no value is present for Kid, not even an explicit nil
func (o *OAuth2ClientJsonSigningKeyResponse) UnsetKid() {
	o.Kid.Unset()
}

// GetKty returns the Kty field value
func (o *OAuth2ClientJsonSigningKeyResponse) GetKty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kty
}

// GetKtyOk returns a tuple with the Kty field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonSigningKeyResponse) GetKtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kty, true
}

// SetKty sets field value
func (o *OAuth2ClientJsonSigningKeyResponse) SetKty(v string) {
	o.Kty = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *OAuth2ClientJsonSigningKeyResponse) GetLastUpdated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonSigningKeyResponse) GetLastUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *OAuth2ClientJsonSigningKeyResponse) SetLastUpdated(v string) {
	o.LastUpdated = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OAuth2ClientJsonSigningKeyResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonSigningKeyResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OAuth2ClientJsonSigningKeyResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OAuth2ClientJsonSigningKeyResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUse returns the Use field value
func (o *OAuth2ClientJsonSigningKeyResponse) GetUse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Use
}

// GetUseOk returns a tuple with the Use field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonSigningKeyResponse) GetUseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Use, true
}

// SetUse sets field value
func (o *OAuth2ClientJsonSigningKeyResponse) SetUse(v string) {
	o.Use = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OAuth2ClientJsonSigningKeyResponse) GetLinks() OAuthClientSecretLinks {
	if o == nil || IsNil(o.Links) {
		var ret OAuthClientSecretLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonSigningKeyResponse) GetLinksOk() (*OAuthClientSecretLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OAuth2ClientJsonSigningKeyResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OAuthClientSecretLinks and assigns it to the Links field.
func (o *OAuth2ClientJsonSigningKeyResponse) SetLinks(v OAuthClientSecretLinks) {
	o.Links = &v
}

func (o OAuth2ClientJsonSigningKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ClientJsonSigningKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alg"] = o.Alg
	toSerialize["created"] = o.Created
	toSerialize["id"] = o.Id
	if o.Kid.IsSet() {
		toSerialize["kid"] = o.Kid.Get()
	}
	toSerialize["kty"] = o.Kty
	toSerialize["lastUpdated"] = o.LastUpdated
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["use"] = o.Use
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2ClientJsonSigningKeyResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alg",
		"created",
		"id",
		"kty",
		"lastUpdated",
		"use",
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

	varOAuth2ClientJsonSigningKeyResponse := _OAuth2ClientJsonSigningKeyResponse{}

	err = json.Unmarshal(data, &varOAuth2ClientJsonSigningKeyResponse)

	if err != nil {
		return err
	}

	*o = OAuth2ClientJsonSigningKeyResponse(varOAuth2ClientJsonSigningKeyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alg")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "status")
		delete(additionalProperties, "use")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2ClientJsonSigningKeyResponse struct {
	value *OAuth2ClientJsonSigningKeyResponse
	isSet bool
}

func (v NullableOAuth2ClientJsonSigningKeyResponse) Get() *OAuth2ClientJsonSigningKeyResponse {
	return v.value
}

func (v *NullableOAuth2ClientJsonSigningKeyResponse) Set(val *OAuth2ClientJsonSigningKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientJsonSigningKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientJsonSigningKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientJsonSigningKeyResponse(val *OAuth2ClientJsonSigningKeyResponse) *NullableOAuth2ClientJsonSigningKeyResponse {
	return &NullableOAuth2ClientJsonSigningKeyResponse{value: val, isSet: true}
}

func (v NullableOAuth2ClientJsonSigningKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientJsonSigningKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
