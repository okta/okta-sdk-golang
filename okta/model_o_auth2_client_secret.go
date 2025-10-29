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

// checks if the OAuth2ClientSecret type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ClientSecret{}

// OAuth2ClientSecret struct for OAuth2ClientSecret
type OAuth2ClientSecret struct {
	// The OAuth 2.0 client secret string
	ClientSecret *string `json:"client_secret,omitempty"`
	// Timestamp when the OAuth Client 2.0 Secret was created
	Created *string `json:"created,omitempty"`
	// The unique ID of the OAuth Client Secret
	Id *string `json:"id,omitempty"`
	// Timestamp when the OAuth Client 2.0 Secret was updated
	LastUpdated *string `json:"lastUpdated,omitempty"`
	// OAuth 2.0 client secret string hash
	SecretHash *string `json:"secret_hash,omitempty"`
	// Status of the OAuth 2.0 Client Secret
	Status               *string                 `json:"status,omitempty"`
	Links                *OAuthClientSecretLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientSecret OAuth2ClientSecret

// NewOAuth2ClientSecret instantiates a new OAuth2ClientSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientSecret() *OAuth2ClientSecret {
	this := OAuth2ClientSecret{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// NewOAuth2ClientSecretWithDefaults instantiates a new OAuth2ClientSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientSecretWithDefaults() *OAuth2ClientSecret {
	this := OAuth2ClientSecret{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *OAuth2ClientSecret) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientSecret) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *OAuth2ClientSecret) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *OAuth2ClientSecret) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OAuth2ClientSecret) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientSecret) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OAuth2ClientSecret) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *OAuth2ClientSecret) SetCreated(v string) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OAuth2ClientSecret) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientSecret) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OAuth2ClientSecret) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OAuth2ClientSecret) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *OAuth2ClientSecret) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientSecret) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *OAuth2ClientSecret) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *OAuth2ClientSecret) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetSecretHash returns the SecretHash field value if set, zero value otherwise.
func (o *OAuth2ClientSecret) GetSecretHash() string {
	if o == nil || IsNil(o.SecretHash) {
		var ret string
		return ret
	}
	return *o.SecretHash
}

// GetSecretHashOk returns a tuple with the SecretHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientSecret) GetSecretHashOk() (*string, bool) {
	if o == nil || IsNil(o.SecretHash) {
		return nil, false
	}
	return o.SecretHash, true
}

// HasSecretHash returns a boolean if a field has been set.
func (o *OAuth2ClientSecret) HasSecretHash() bool {
	if o != nil && !IsNil(o.SecretHash) {
		return true
	}

	return false
}

// SetSecretHash gets a reference to the given string and assigns it to the SecretHash field.
func (o *OAuth2ClientSecret) SetSecretHash(v string) {
	o.SecretHash = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OAuth2ClientSecret) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientSecret) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OAuth2ClientSecret) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OAuth2ClientSecret) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OAuth2ClientSecret) GetLinks() OAuthClientSecretLinks {
	if o == nil || IsNil(o.Links) {
		var ret OAuthClientSecretLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientSecret) GetLinksOk() (*OAuthClientSecretLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OAuth2ClientSecret) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OAuthClientSecretLinks and assigns it to the Links field.
func (o *OAuth2ClientSecret) SetLinks(v OAuthClientSecretLinks) {
	o.Links = &v
}

func (o OAuth2ClientSecret) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ClientSecret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientSecret) {
		toSerialize["client_secret"] = o.ClientSecret
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
	if !IsNil(o.SecretHash) {
		toSerialize["secret_hash"] = o.SecretHash
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

func (o *OAuth2ClientSecret) UnmarshalJSON(data []byte) (err error) {
	varOAuth2ClientSecret := _OAuth2ClientSecret{}

	err = json.Unmarshal(data, &varOAuth2ClientSecret)

	if err != nil {
		return err
	}

	*o = OAuth2ClientSecret(varOAuth2ClientSecret)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "client_secret")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "secret_hash")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2ClientSecret struct {
	value *OAuth2ClientSecret
	isSet bool
}

func (v NullableOAuth2ClientSecret) Get() *OAuth2ClientSecret {
	return v.value
}

func (v *NullableOAuth2ClientSecret) Set(val *OAuth2ClientSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientSecret(val *OAuth2ClientSecret) *NullableOAuth2ClientSecret {
	return &NullableOAuth2ClientSecret{value: val, isSet: true}
}

func (v NullableOAuth2ClientSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
