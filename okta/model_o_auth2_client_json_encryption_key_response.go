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
)

// checks if the OAuth2ClientJsonEncryptionKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ClientJsonEncryptionKeyResponse{}

// OAuth2ClientJsonEncryptionKeyResponse <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>A [JSON Web Key (JWK)](https://tools.ietf.org/html/rfc7517) is a JSON representation of a cryptographic key. Okta uses an encryption key to encrypt an ID token JWT minted by the org authorization server or custom authorization server. Okta supports only RSA keys for encrypting tokens.
type OAuth2ClientJsonEncryptionKeyResponse struct {
	// RSA key value (exponent) for key binding
	E *string `json:"e,omitempty"`
	// Cryptographic algorithm family for the certificate's key pair
	Kty *string `json:"kty,omitempty"`
	// RSA key value (modulus) for key binding
	N *string `json:"n,omitempty"`
	// Acceptable use of the JSON Web Key
	Use *string `json:"use,omitempty"`
	// Unique identifier of the JSON Web Key in the OAUth 2.0 client's JWKS
	Kid NullableString `json:"kid,omitempty"`
	// Status of the OAuth 2.0 client JSON Web Key
	Status *string `json:"status,omitempty"`
	// Timestamp when the OAuth 2.0 client JSON Web Key was created
	Created *string `json:"created,omitempty"`
	// The unique ID of the OAuth Client JSON Web Key
	Id *string `json:"id,omitempty"`
	// Timestamp when the OAuth 2.0 client JSON Web Key was updated
	LastUpdated          *string                 `json:"lastUpdated,omitempty"`
	Links                *OAuthClientSecretLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientJsonEncryptionKeyResponse OAuth2ClientJsonEncryptionKeyResponse

// NewOAuth2ClientJsonEncryptionKeyResponse instantiates a new OAuth2ClientJsonEncryptionKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientJsonEncryptionKeyResponse() *OAuth2ClientJsonEncryptionKeyResponse {
	this := OAuth2ClientJsonEncryptionKeyResponse{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// NewOAuth2ClientJsonEncryptionKeyResponseWithDefaults instantiates a new OAuth2ClientJsonEncryptionKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientJsonEncryptionKeyResponseWithDefaults() *OAuth2ClientJsonEncryptionKeyResponse {
	this := OAuth2ClientJsonEncryptionKeyResponse{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetE() string {
	if o == nil || IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetEOk() (*string, bool) {
	if o == nil || IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) HasE() bool {
	if o != nil && !IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *OAuth2ClientJsonEncryptionKeyResponse) SetE(v string) {
	o.E = &v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetKty() string {
	if o == nil || IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetKtyOk() (*string, bool) {
	if o == nil || IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) HasKty() bool {
	if o != nil && !IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *OAuth2ClientJsonEncryptionKeyResponse) SetKty(v string) {
	o.Kty = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetN() string {
	if o == nil || IsNil(o.N) {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetNOk() (*string, bool) {
	if o == nil || IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) HasN() bool {
	if o != nil && !IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *OAuth2ClientJsonEncryptionKeyResponse) SetN(v string) {
	o.N = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *OAuth2ClientJsonEncryptionKeyResponse) SetUse(v string) {
	o.Use = &v
}

// GetKid returns the Kid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetKid() string {
	if o == nil || IsNil(o.Kid.Get()) {
		var ret string
		return ret
	}
	return *o.Kid.Get()
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kid.Get(), o.Kid.IsSet()
}

// HasKid returns a boolean if a field has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) HasKid() bool {
	if o != nil && o.Kid.IsSet() {
		return true
	}

	return false
}

// SetKid gets a reference to the given NullableString and assigns it to the Kid field.
func (o *OAuth2ClientJsonEncryptionKeyResponse) SetKid(v string) {
	o.Kid.Set(&v)
}

// SetKidNil sets the value for Kid to be an explicit nil
func (o *OAuth2ClientJsonEncryptionKeyResponse) SetKidNil() {
	o.Kid.Set(nil)
}

// UnsetKid ensures that no value is present for Kid, not even an explicit nil
func (o *OAuth2ClientJsonEncryptionKeyResponse) UnsetKid() {
	o.Kid.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OAuth2ClientJsonEncryptionKeyResponse) SetStatus(v string) {
	o.Status = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *OAuth2ClientJsonEncryptionKeyResponse) SetCreated(v string) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OAuth2ClientJsonEncryptionKeyResponse) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *OAuth2ClientJsonEncryptionKeyResponse) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetLinks() OAuthClientSecretLinks {
	if o == nil || IsNil(o.Links) {
		var ret OAuthClientSecretLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) GetLinksOk() (*OAuthClientSecretLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OAuth2ClientJsonEncryptionKeyResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OAuthClientSecretLinks and assigns it to the Links field.
func (o *OAuth2ClientJsonEncryptionKeyResponse) SetLinks(v OAuthClientSecretLinks) {
	o.Links = &v
}

func (o OAuth2ClientJsonEncryptionKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ClientJsonEncryptionKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	if !IsNil(o.Kty) {
		toSerialize["kty"] = o.Kty
	}
	if !IsNil(o.N) {
		toSerialize["n"] = o.N
	}
	if !IsNil(o.Use) {
		toSerialize["use"] = o.Use
	}
	if o.Kid.IsSet() {
		toSerialize["kid"] = o.Kid.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2ClientJsonEncryptionKeyResponse) UnmarshalJSON(data []byte) (err error) {
	varOAuth2ClientJsonEncryptionKeyResponse := _OAuth2ClientJsonEncryptionKeyResponse{}

	err = json.Unmarshal(data, &varOAuth2ClientJsonEncryptionKeyResponse)

	if err != nil {
		return err
	}

	*o = OAuth2ClientJsonEncryptionKeyResponse(varOAuth2ClientJsonEncryptionKeyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "n")
		delete(additionalProperties, "use")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "status")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2ClientJsonEncryptionKeyResponse struct {
	value *OAuth2ClientJsonEncryptionKeyResponse
	isSet bool
}

func (v NullableOAuth2ClientJsonEncryptionKeyResponse) Get() *OAuth2ClientJsonEncryptionKeyResponse {
	return v.value
}

func (v *NullableOAuth2ClientJsonEncryptionKeyResponse) Set(val *OAuth2ClientJsonEncryptionKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientJsonEncryptionKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientJsonEncryptionKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientJsonEncryptionKeyResponse(val *OAuth2ClientJsonEncryptionKeyResponse) *NullableOAuth2ClientJsonEncryptionKeyResponse {
	return &NullableOAuth2ClientJsonEncryptionKeyResponse{value: val, isSet: true}
}

func (v NullableOAuth2ClientJsonEncryptionKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientJsonEncryptionKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
