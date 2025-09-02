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

// OAuth2ClientJsonWebKeyECResponse An EC signing key
type OAuth2ClientJsonWebKeyECResponse struct {
	// Cryptographic algorithm family for the certificate's key pair
	Kty *string `json:"kty,omitempty"`
	// The public x coordinate for the elliptic curve point
	X *string `json:"x,omitempty"`
	// The public y coordinate for the elliptic curve point
	Y *string `json:"y,omitempty"`
	// Unique identifier of the JSON Web Key in the OAUth 2.0 client's JWKS
	Kid NullableString `json:"kid,omitempty"`
	// Status of the OAuth 2.0 client JSON Web Key
	Status *string `json:"status,omitempty"`
	// Timestamp when the OAuth 2.0 client JSON Web Key was created
	Created *string `json:"created,omitempty"`
	// The unique ID of the OAuth Client JSON Web Key
	Id *string `json:"id,omitempty"`
	// Timestamp when the OAuth 2.0 client JSON Web Key was updated
	LastUpdated *string `json:"lastUpdated,omitempty"`
	Links *OAuthClientSecretLinks `json:"_links,omitempty"`
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

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetKty() string {
	if o == nil || o.Kty == nil {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetKtyOk() (*string, bool) {
	if o == nil || o.Kty == nil {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasKty() bool {
	if o != nil && o.Kty != nil {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetKty(v string) {
	o.Kty = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetX() string {
	if o == nil || o.X == nil {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetXOk() (*string, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasX() bool {
	if o != nil && o.X != nil {
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
	if o == nil || o.Y == nil {
		var ret string
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetYOk() (*string, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasY() bool {
	if o != nil && o.Y != nil {
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
	if o == nil || o.Kid.Get() == nil {
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
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetStatus(v string) {
	o.Status = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetCreated(v string) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetLastUpdated() string {
	if o == nil || o.LastUpdated == nil {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetLastUpdatedOk() (*string, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetLinks() OAuthClientSecretLinks {
	if o == nil || o.Links == nil {
		var ret OAuthClientSecretLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetLinksOk() (*OAuthClientSecretLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OAuthClientSecretLinks and assigns it to the Links field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetLinks(v OAuthClientSecretLinks) {
	o.Links = &v
}

func (o OAuth2ClientJsonWebKeyECResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kty != nil {
		toSerialize["kty"] = o.Kty
	}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Y != nil {
		toSerialize["y"] = o.Y
	}
	if o.Kid.IsSet() {
		toSerialize["kid"] = o.Kid.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2ClientJsonWebKeyECResponse) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2ClientJsonWebKeyECResponse := _OAuth2ClientJsonWebKeyECResponse{}

	err = json.Unmarshal(bytes, &varOAuth2ClientJsonWebKeyECResponse)
	if err == nil {
		*o = OAuth2ClientJsonWebKeyECResponse(varOAuth2ClientJsonWebKeyECResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "kty")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "status")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

