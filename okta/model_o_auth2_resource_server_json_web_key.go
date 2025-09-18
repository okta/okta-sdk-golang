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

// checks if the OAuth2ResourceServerJsonWebKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ResourceServerJsonWebKey{}

// OAuth2ResourceServerJsonWebKey struct for OAuth2ResourceServerJsonWebKey
type OAuth2ResourceServerJsonWebKey struct {
	// Timestamp when the JSON Web Key was created
	Created *string `json:"created,omitempty"`
	// RSA key value (exponent) for key binding
	E *string `json:"e,omitempty"`
	// The unique ID of the JSON Web Key
	Id *string `json:"id,omitempty"`
	// Unique identifier of the JSON Web Key in the Custom Authorization Server's Public JWKS
	Kid NullableString `json:"kid,omitempty"`
	// Cryptographic algorithm family for the certificate's key pair
	Kty *string `json:"kty,omitempty"`
	// Timestamp when the JSON Web Key was updated
	LastUpdated *string `json:"lastUpdated,omitempty"`
	// RSA key value (modulus) for key binding
	N *string `json:"n,omitempty"`
	// The status of the encryption key. You can use only an `ACTIVE` key to encrypt tokens issued by the authorization server.
	Status *string `json:"status,omitempty"`
	// Acceptable use of the JSON Web Key
	Use                  *string                      `json:"use,omitempty"`
	Links                *OAuthResourceServerKeyLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ResourceServerJsonWebKey OAuth2ResourceServerJsonWebKey

// NewOAuth2ResourceServerJsonWebKey instantiates a new OAuth2ResourceServerJsonWebKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ResourceServerJsonWebKey() *OAuth2ResourceServerJsonWebKey {
	this := OAuth2ResourceServerJsonWebKey{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// NewOAuth2ResourceServerJsonWebKeyWithDefaults instantiates a new OAuth2ResourceServerJsonWebKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ResourceServerJsonWebKeyWithDefaults() *OAuth2ResourceServerJsonWebKey {
	this := OAuth2ResourceServerJsonWebKey{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OAuth2ResourceServerJsonWebKey) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ResourceServerJsonWebKey) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OAuth2ResourceServerJsonWebKey) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *OAuth2ResourceServerJsonWebKey) SetCreated(v string) {
	o.Created = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OAuth2ResourceServerJsonWebKey) GetE() string {
	if o == nil || IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ResourceServerJsonWebKey) GetEOk() (*string, bool) {
	if o == nil || IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *OAuth2ResourceServerJsonWebKey) HasE() bool {
	if o != nil && !IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *OAuth2ResourceServerJsonWebKey) SetE(v string) {
	o.E = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OAuth2ResourceServerJsonWebKey) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ResourceServerJsonWebKey) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OAuth2ResourceServerJsonWebKey) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OAuth2ResourceServerJsonWebKey) SetId(v string) {
	o.Id = &v
}

// GetKid returns the Kid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ResourceServerJsonWebKey) GetKid() string {
	if o == nil || IsNil(o.Kid.Get()) {
		var ret string
		return ret
	}
	return *o.Kid.Get()
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ResourceServerJsonWebKey) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kid.Get(), o.Kid.IsSet()
}

// HasKid returns a boolean if a field has been set.
func (o *OAuth2ResourceServerJsonWebKey) HasKid() bool {
	if o != nil && o.Kid.IsSet() {
		return true
	}

	return false
}

// SetKid gets a reference to the given NullableString and assigns it to the Kid field.
func (o *OAuth2ResourceServerJsonWebKey) SetKid(v string) {
	o.Kid.Set(&v)
}

// SetKidNil sets the value for Kid to be an explicit nil
func (o *OAuth2ResourceServerJsonWebKey) SetKidNil() {
	o.Kid.Set(nil)
}

// UnsetKid ensures that no value is present for Kid, not even an explicit nil
func (o *OAuth2ResourceServerJsonWebKey) UnsetKid() {
	o.Kid.Unset()
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *OAuth2ResourceServerJsonWebKey) GetKty() string {
	if o == nil || IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ResourceServerJsonWebKey) GetKtyOk() (*string, bool) {
	if o == nil || IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *OAuth2ResourceServerJsonWebKey) HasKty() bool {
	if o != nil && !IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *OAuth2ResourceServerJsonWebKey) SetKty(v string) {
	o.Kty = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *OAuth2ResourceServerJsonWebKey) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ResourceServerJsonWebKey) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *OAuth2ResourceServerJsonWebKey) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *OAuth2ResourceServerJsonWebKey) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *OAuth2ResourceServerJsonWebKey) GetN() string {
	if o == nil || IsNil(o.N) {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ResourceServerJsonWebKey) GetNOk() (*string, bool) {
	if o == nil || IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *OAuth2ResourceServerJsonWebKey) HasN() bool {
	if o != nil && !IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *OAuth2ResourceServerJsonWebKey) SetN(v string) {
	o.N = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OAuth2ResourceServerJsonWebKey) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ResourceServerJsonWebKey) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OAuth2ResourceServerJsonWebKey) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OAuth2ResourceServerJsonWebKey) SetStatus(v string) {
	o.Status = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *OAuth2ResourceServerJsonWebKey) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ResourceServerJsonWebKey) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *OAuth2ResourceServerJsonWebKey) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *OAuth2ResourceServerJsonWebKey) SetUse(v string) {
	o.Use = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OAuth2ResourceServerJsonWebKey) GetLinks() OAuthResourceServerKeyLinks {
	if o == nil || IsNil(o.Links) {
		var ret OAuthResourceServerKeyLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ResourceServerJsonWebKey) GetLinksOk() (*OAuthResourceServerKeyLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OAuth2ResourceServerJsonWebKey) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OAuthResourceServerKeyLinks and assigns it to the Links field.
func (o *OAuth2ResourceServerJsonWebKey) SetLinks(v OAuthResourceServerKeyLinks) {
	o.Links = &v
}

func (o OAuth2ResourceServerJsonWebKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ResourceServerJsonWebKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Kid.IsSet() {
		toSerialize["kid"] = o.Kid.Get()
	}
	if !IsNil(o.Kty) {
		toSerialize["kty"] = o.Kty
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.N) {
		toSerialize["n"] = o.N
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Use) {
		toSerialize["use"] = o.Use
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2ResourceServerJsonWebKey) UnmarshalJSON(data []byte) (err error) {
	varOAuth2ResourceServerJsonWebKey := _OAuth2ResourceServerJsonWebKey{}

	err = json.Unmarshal(data, &varOAuth2ResourceServerJsonWebKey)

	if err != nil {
		return err
	}

	*o = OAuth2ResourceServerJsonWebKey(varOAuth2ResourceServerJsonWebKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "e")
		delete(additionalProperties, "id")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "n")
		delete(additionalProperties, "status")
		delete(additionalProperties, "use")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2ResourceServerJsonWebKey struct {
	value *OAuth2ResourceServerJsonWebKey
	isSet bool
}

func (v NullableOAuth2ResourceServerJsonWebKey) Get() *OAuth2ResourceServerJsonWebKey {
	return v.value
}

func (v *NullableOAuth2ResourceServerJsonWebKey) Set(val *OAuth2ResourceServerJsonWebKey) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ResourceServerJsonWebKey) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ResourceServerJsonWebKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ResourceServerJsonWebKey(val *OAuth2ResourceServerJsonWebKey) *NullableOAuth2ResourceServerJsonWebKey {
	return &NullableOAuth2ResourceServerJsonWebKey{value: val, isSet: true}
}

func (v NullableOAuth2ResourceServerJsonWebKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ResourceServerJsonWebKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
