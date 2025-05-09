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

// OAuth2ScopeConsentGrant Grant object that represents an app consent scope grant
type OAuth2ScopeConsentGrant struct {
	// Client ID of the app integration
	ClientId *string `json:"clientId,omitempty"`
	// Timestamp when the object was created
	Created *time.Time `json:"created,omitempty"`
	CreatedBy *OAuth2Actor `json:"createdBy,omitempty"`
	// ID of the Grant object
	Id *string `json:"id,omitempty"`
	// The issuer of your org authorization server. This is typically your Okta domain.
	Issuer string `json:"issuer"`
	// Timestamp when the object was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The name of the [Okta scope](https://developer.okta.com/docs/api/oauth2/#oauth-20-scopes) for which consent is granted
	ScopeId string `json:"scopeId"`
	// User type source that granted consent
	Source *string `json:"source,omitempty"`
	// Status
	Status *string `json:"status,omitempty"`
	// User ID that granted consent (if `source` is `END_USER`)
	UserId *string `json:"userId,omitempty"`
	Embedded *OAuth2ScopeConsentGrantEmbedded `json:"_embedded,omitempty"`
	Links *OAuth2ScopeConsentGrantLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ScopeConsentGrant OAuth2ScopeConsentGrant

// NewOAuth2ScopeConsentGrant instantiates a new OAuth2ScopeConsentGrant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ScopeConsentGrant(issuer string, scopeId string) *OAuth2ScopeConsentGrant {
	this := OAuth2ScopeConsentGrant{}
	this.Issuer = issuer
	this.ScopeId = scopeId
	return &this
}

// NewOAuth2ScopeConsentGrantWithDefaults instantiates a new OAuth2ScopeConsentGrant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ScopeConsentGrantWithDefaults() *OAuth2ScopeConsentGrant {
	this := OAuth2ScopeConsentGrant{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrant) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrant) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrant) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OAuth2ScopeConsentGrant) SetClientId(v string) {
	o.ClientId = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrant) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrant) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrant) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *OAuth2ScopeConsentGrant) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrant) GetCreatedBy() OAuth2Actor {
	if o == nil || o.CreatedBy == nil {
		var ret OAuth2Actor
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrant) GetCreatedByOk() (*OAuth2Actor, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrant) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given OAuth2Actor and assigns it to the CreatedBy field.
func (o *OAuth2ScopeConsentGrant) SetCreatedBy(v OAuth2Actor) {
	o.CreatedBy = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrant) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrant) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrant) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OAuth2ScopeConsentGrant) SetId(v string) {
	o.Id = &v
}

// GetIssuer returns the Issuer field value
func (o *OAuth2ScopeConsentGrant) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrant) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *OAuth2ScopeConsentGrant) SetIssuer(v string) {
	o.Issuer = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrant) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrant) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrant) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *OAuth2ScopeConsentGrant) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetScopeId returns the ScopeId field value
func (o *OAuth2ScopeConsentGrant) GetScopeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScopeId
}

// GetScopeIdOk returns a tuple with the ScopeId field value
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrant) GetScopeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScopeId, true
}

// SetScopeId sets field value
func (o *OAuth2ScopeConsentGrant) SetScopeId(v string) {
	o.ScopeId = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrant) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrant) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrant) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *OAuth2ScopeConsentGrant) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrant) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrant) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrant) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OAuth2ScopeConsentGrant) SetStatus(v string) {
	o.Status = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrant) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrant) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrant) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *OAuth2ScopeConsentGrant) SetUserId(v string) {
	o.UserId = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrant) GetEmbedded() OAuth2ScopeConsentGrantEmbedded {
	if o == nil || o.Embedded == nil {
		var ret OAuth2ScopeConsentGrantEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrant) GetEmbeddedOk() (*OAuth2ScopeConsentGrantEmbedded, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrant) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given OAuth2ScopeConsentGrantEmbedded and assigns it to the Embedded field.
func (o *OAuth2ScopeConsentGrant) SetEmbedded(v OAuth2ScopeConsentGrantEmbedded) {
	o.Embedded = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrant) GetLinks() OAuth2ScopeConsentGrantLinks {
	if o == nil || o.Links == nil {
		var ret OAuth2ScopeConsentGrantLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrant) GetLinksOk() (*OAuth2ScopeConsentGrantLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrant) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OAuth2ScopeConsentGrantLinks and assigns it to the Links field.
func (o *OAuth2ScopeConsentGrant) SetLinks(v OAuth2ScopeConsentGrantLinks) {
	o.Links = &v
}

func (o OAuth2ScopeConsentGrant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["issuer"] = o.Issuer
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["scopeId"] = o.ScopeId
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
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

func (o *OAuth2ScopeConsentGrant) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2ScopeConsentGrant := _OAuth2ScopeConsentGrant{}

	err = json.Unmarshal(bytes, &varOAuth2ScopeConsentGrant)
	if err == nil {
		*o = OAuth2ScopeConsentGrant(varOAuth2ScopeConsentGrant)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "created")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "id")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "scopeId")
		delete(additionalProperties, "source")
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

type NullableOAuth2ScopeConsentGrant struct {
	value *OAuth2ScopeConsentGrant
	isSet bool
}

func (v NullableOAuth2ScopeConsentGrant) Get() *OAuth2ScopeConsentGrant {
	return v.value
}

func (v *NullableOAuth2ScopeConsentGrant) Set(val *OAuth2ScopeConsentGrant) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ScopeConsentGrant) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ScopeConsentGrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ScopeConsentGrant(val *OAuth2ScopeConsentGrant) *NullableOAuth2ScopeConsentGrant {
	return &NullableOAuth2ScopeConsentGrant{value: val, isSet: true}
}

func (v NullableOAuth2ScopeConsentGrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ScopeConsentGrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

