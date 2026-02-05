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
)

// checks if the OAuth2ClientJsonWebKeyECResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ClientJsonWebKeyECResponse{}

// OAuth2ClientJsonWebKeyECResponse An EC signing key
type OAuth2ClientJsonWebKeyECResponse struct {
	// Algorithm used in the key
	Alg *string `json:"alg,omitempty"`
	// Timestamp when the OAuth 2.0 client JSON Web Key was created
	Created *string `json:"created,omitempty"`
	// Identifies the cryptographic curve used with the key
	Crv *string `json:"crv,omitempty"`
	// The unique ID of the OAuth Client JSON Web Key
	Id *string `json:"id,omitempty"`
	// Cryptographic algorithm family for the certificate's key pair
	Kty *string `json:"kty,omitempty"`
	// Timestamp when the OAuth 2.0 client JSON Web Key was updated
	LastUpdated *string `json:"lastUpdated,omitempty"`
	// The public x coordinate for the elliptic curve point
	X *string `json:"x,omitempty"`
	// The public y coordinate for the elliptic curve point
	Y *string `json:"y,omitempty"`
	// Unique identifier of the JSON Web Key in the OAuth 2.0 client's JWKS
	Kid NullableString `json:"kid,omitempty"`
	// Status of the OAuth 2.0 client JSON Web Key
	Status *string `json:"status,omitempty"`
	// Acceptable use of the JSON Web Key
	Use                  *string                 `json:"use,omitempty"`
	Links                *OAuthClientSecretLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientJsonWebKeyECResponse OAuth2ClientJsonWebKeyECResponse

// NewOAuth2ClientJsonWebKeyECResponse instantiates a new OAuth2ClientJsonWebKeyECResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientJsonWebKeyECResponse() *OAuth2ClientJsonWebKeyECResponse {
	this := OAuth2ClientJsonWebKeyECResponse{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// NewOAuth2ClientJsonWebKeyECResponseWithDefaults instantiates a new OAuth2ClientJsonWebKeyECResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientJsonWebKeyECResponseWithDefaults() *OAuth2ClientJsonWebKeyECResponse {
	this := OAuth2ClientJsonWebKeyECResponse{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetAlg() string {
	if o == nil || IsNil(o.Alg) {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetAlgOk() (*string, bool) {
	if o == nil || IsNil(o.Alg) {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasAlg() bool {
	if o != nil && !IsNil(o.Alg) {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetAlg(v string) {
	o.Alg = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetCreated(v string) {
	o.Created = &v
}

// GetCrv returns the Crv field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetCrv() string {
	if o == nil || IsNil(o.Crv) {
		var ret string
		return ret
	}
	return *o.Crv
}

// GetCrvOk returns a tuple with the Crv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetCrvOk() (*string, bool) {
	if o == nil || IsNil(o.Crv) {
		return nil, false
	}
	return o.Crv, true
}

// HasCrv returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasCrv() bool {
	if o != nil && !IsNil(o.Crv) {
		return true
	}

	return false
}

// SetCrv gets a reference to the given string and assigns it to the Crv field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetCrv(v string) {
	o.Crv = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetId(v string) {
	o.Id = &v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetKty() string {
	if o == nil || IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetKtyOk() (*string, bool) {
	if o == nil || IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasKty() bool {
	if o != nil && !IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetKty(v string) {
	o.Kty = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetX() string {
	if o == nil || IsNil(o.X) {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetXOk() (*string, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasX() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetX(v string) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetY() string {
	if o == nil || IsNil(o.Y) {
		var ret string
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetYOk() (*string, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given string and assigns it to the Y field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetY(v string) {
	o.Y = &v
}

// GetKid returns the Kid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientJsonWebKeyECResponse) GetKid() string {
	if o == nil || IsNil(o.Kid.Get()) {
		var ret string
		return ret
	}
	return *o.Kid.Get()
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientJsonWebKeyECResponse) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kid.Get(), o.Kid.IsSet()
}

// HasKid returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasKid() bool {
	if o != nil && o.Kid.IsSet() {
		return true
	}

	return false
}

// SetKid gets a reference to the given NullableString and assigns it to the Kid field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetKid(v string) {
	o.Kid.Set(&v)
}

// SetKidNil sets the value for Kid to be an explicit nil
func (o *OAuth2ClientJsonWebKeyECResponse) SetKidNil() {
	o.Kid.Set(nil)
}

// UnsetKid ensures that no value is present for Kid, not even an explicit nil
func (o *OAuth2ClientJsonWebKeyECResponse) UnsetKid() {
	o.Kid.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetUse(v string) {
	o.Use = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetLinks() OAuthClientSecretLinks {
	if o == nil || IsNil(o.Links) {
		var ret OAuthClientSecretLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetLinksOk() (*OAuthClientSecretLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OAuthClientSecretLinks and assigns it to the Links field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetLinks(v OAuthClientSecretLinks) {
	o.Links = &v
}

func (o OAuth2ClientJsonWebKeyECResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ClientJsonWebKeyECResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alg) {
		toSerialize["alg"] = o.Alg
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Crv) {
		toSerialize["crv"] = o.Crv
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Kty) {
		toSerialize["kty"] = o.Kty
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}
	if o.Kid.IsSet() {
		toSerialize["kid"] = o.Kid.Get()
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

func (o *OAuth2ClientJsonWebKeyECResponse) UnmarshalJSON(data []byte) (err error) {
	varOAuth2ClientJsonWebKeyECResponse := _OAuth2ClientJsonWebKeyECResponse{}

	err = json.Unmarshal(data, &varOAuth2ClientJsonWebKeyECResponse)

	if err != nil {
		return err
	}

	*o = OAuth2ClientJsonWebKeyECResponse(varOAuth2ClientJsonWebKeyECResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alg")
		delete(additionalProperties, "created")
		delete(additionalProperties, "crv")
		delete(additionalProperties, "id")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "status")
		delete(additionalProperties, "use")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2ClientJsonWebKeyECResponse struct {
	value *OAuth2ClientJsonWebKeyECResponse
	isSet bool
}

func (v NullableOAuth2ClientJsonWebKeyECResponse) Get() *OAuth2ClientJsonWebKeyECResponse {
	return v.value
}

func (v *NullableOAuth2ClientJsonWebKeyECResponse) Set(val *OAuth2ClientJsonWebKeyECResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientJsonWebKeyECResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientJsonWebKeyECResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientJsonWebKeyECResponse(val *OAuth2ClientJsonWebKeyECResponse) *NullableOAuth2ClientJsonWebKeyECResponse {
	return &NullableOAuth2ClientJsonWebKeyECResponse{value: val, isSet: true}
}

func (v NullableOAuth2ClientJsonWebKeyECResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientJsonWebKeyECResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
