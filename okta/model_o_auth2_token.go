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

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"time"
)

// OAuth2Token struct for OAuth2Token
type OAuth2Token struct {
	// Client ID
	ClientId *string `json:"clientId,omitempty"`
	// Timestamp when the object was created
	Created *time.Time `json:"created,omitempty"`
	// Expiration time of the OAuth 2.0 Token
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// ID of the Token object
	Id *string `json:"id,omitempty"`
	// The complete URL of the authorization server that issued the Token
	Issuer *string `json:"issuer,omitempty"`
	// Timestamp when the object was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Name of scopes attached to the Token
	Scopes []string `json:"scopes,omitempty"`
	// Status
	Status *string `json:"status,omitempty"`
	UserId *string `json:"userId,omitempty"`
	// Embedded resources related to the object if the `expand` query parameter is specified
	Embedded map[string]map[string]interface{} `json:"_embedded,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2Token OAuth2Token

// NewOAuth2Token instantiates a new OAuth2Token object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2Token() *OAuth2Token {
	this := OAuth2Token{}
	return &this
}

// NewOAuth2TokenWithDefaults instantiates a new OAuth2Token object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2TokenWithDefaults() *OAuth2Token {
	this := OAuth2Token{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OAuth2Token) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Token) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OAuth2Token) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OAuth2Token) SetClientId(v string) {
	o.ClientId = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OAuth2Token) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Token) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OAuth2Token) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *OAuth2Token) SetCreated(v time.Time) {
	o.Created = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *OAuth2Token) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Token) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *OAuth2Token) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *OAuth2Token) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OAuth2Token) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Token) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OAuth2Token) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OAuth2Token) SetId(v string) {
	o.Id = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *OAuth2Token) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Token) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *OAuth2Token) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *OAuth2Token) SetIssuer(v string) {
	o.Issuer = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *OAuth2Token) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Token) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *OAuth2Token) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *OAuth2Token) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *OAuth2Token) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Token) GetScopesOk() ([]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *OAuth2Token) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *OAuth2Token) SetScopes(v []string) {
	o.Scopes = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OAuth2Token) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Token) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OAuth2Token) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OAuth2Token) SetStatus(v string) {
	o.Status = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *OAuth2Token) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Token) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *OAuth2Token) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *OAuth2Token) SetUserId(v string) {
	o.UserId = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *OAuth2Token) GetEmbedded() map[string]map[string]interface{} {
	if o == nil || o.Embedded == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Token) GetEmbeddedOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *OAuth2Token) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]map[string]interface{} and assigns it to the Embedded field.
func (o *OAuth2Token) SetEmbedded(v map[string]map[string]interface{}) {
	o.Embedded = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OAuth2Token) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Token) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OAuth2Token) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *OAuth2Token) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o OAuth2Token) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.Embedded != nil {
		toSerialize["_embedded"] = o.Embedded
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2Token) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2Token := _OAuth2Token{}

	err = json.Unmarshal(bytes, &varOAuth2Token)
	if err == nil {
		*o = OAuth2Token(varOAuth2Token)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "created")
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "id")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "status")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOAuth2Token struct {
	value *OAuth2Token
	isSet bool
}

func (v NullableOAuth2Token) Get() *OAuth2Token {
	return v.value
}

func (v *NullableOAuth2Token) Set(val *OAuth2Token) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2Token) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2Token) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2Token(val *OAuth2Token) *NullableOAuth2Token {
	return &NullableOAuth2Token{value: val, isSet: true}
}

func (v NullableOAuth2Token) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2Token) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

