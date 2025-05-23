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
)

// APIServiceIntegrationInstanceSecret struct for APIServiceIntegrationInstanceSecret
type APIServiceIntegrationInstanceSecret struct {
	// The OAuth 2.0 client secret string. The client secret string is returned in the response of a Secret creation request. In other responses (such as list, activate, or deactivate requests), the client secret is returned as an undisclosed hashed value.
	ClientSecret string `json:"client_secret"`
	// Timestamp when the API Service Integration instance Secret was created
	Created string `json:"created"`
	// The ID of the API Service Integration instance Secret
	Id string `json:"id"`
	// Timestamp when the API Service Integration instance Secret was updated
	LastUpdated string `json:"lastUpdated"`
	// OAuth 2.0 client secret string hash
	SecretHash string `json:"secret_hash"`
	// Status of the API Service Integration instance Secret
	Status string `json:"status"`
	Links APIServiceIntegrationSecretLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _APIServiceIntegrationInstanceSecret APIServiceIntegrationInstanceSecret

// NewAPIServiceIntegrationInstanceSecret instantiates a new APIServiceIntegrationInstanceSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServiceIntegrationInstanceSecret(clientSecret string, created string, id string, lastUpdated string, secretHash string, status string, links APIServiceIntegrationSecretLinks) *APIServiceIntegrationInstanceSecret {
	this := APIServiceIntegrationInstanceSecret{}
	this.ClientSecret = clientSecret
	this.Created = created
	this.Id = id
	this.LastUpdated = lastUpdated
	this.SecretHash = secretHash
	this.Status = status
	this.Links = links
	return &this
}

// NewAPIServiceIntegrationInstanceSecretWithDefaults instantiates a new APIServiceIntegrationInstanceSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServiceIntegrationInstanceSecretWithDefaults() *APIServiceIntegrationInstanceSecret {
	this := APIServiceIntegrationInstanceSecret{}
	return &this
}

// GetClientSecret returns the ClientSecret field value
func (o *APIServiceIntegrationInstanceSecret) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationInstanceSecret) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *APIServiceIntegrationInstanceSecret) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetCreated returns the Created field value
func (o *APIServiceIntegrationInstanceSecret) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationInstanceSecret) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *APIServiceIntegrationInstanceSecret) SetCreated(v string) {
	o.Created = v
}

// GetId returns the Id field value
func (o *APIServiceIntegrationInstanceSecret) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationInstanceSecret) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *APIServiceIntegrationInstanceSecret) SetId(v string) {
	o.Id = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *APIServiceIntegrationInstanceSecret) GetLastUpdated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationInstanceSecret) GetLastUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *APIServiceIntegrationInstanceSecret) SetLastUpdated(v string) {
	o.LastUpdated = v
}

// GetSecretHash returns the SecretHash field value
func (o *APIServiceIntegrationInstanceSecret) GetSecretHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretHash
}

// GetSecretHashOk returns a tuple with the SecretHash field value
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationInstanceSecret) GetSecretHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretHash, true
}

// SetSecretHash sets field value
func (o *APIServiceIntegrationInstanceSecret) SetSecretHash(v string) {
	o.SecretHash = v
}

// GetStatus returns the Status field value
func (o *APIServiceIntegrationInstanceSecret) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationInstanceSecret) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *APIServiceIntegrationInstanceSecret) SetStatus(v string) {
	o.Status = v
}

// GetLinks returns the Links field value
func (o *APIServiceIntegrationInstanceSecret) GetLinks() APIServiceIntegrationSecretLinks {
	if o == nil {
		var ret APIServiceIntegrationSecretLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationInstanceSecret) GetLinksOk() (*APIServiceIntegrationSecretLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *APIServiceIntegrationInstanceSecret) SetLinks(v APIServiceIntegrationSecretLinks) {
	o.Links = v
}

func (o APIServiceIntegrationInstanceSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["secret_hash"] = o.SecretHash
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *APIServiceIntegrationInstanceSecret) UnmarshalJSON(bytes []byte) (err error) {
	varAPIServiceIntegrationInstanceSecret := _APIServiceIntegrationInstanceSecret{}

	err = json.Unmarshal(bytes, &varAPIServiceIntegrationInstanceSecret)
	if err == nil {
		*o = APIServiceIntegrationInstanceSecret(varAPIServiceIntegrationInstanceSecret)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "client_secret")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "secret_hash")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAPIServiceIntegrationInstanceSecret struct {
	value *APIServiceIntegrationInstanceSecret
	isSet bool
}

func (v NullableAPIServiceIntegrationInstanceSecret) Get() *APIServiceIntegrationInstanceSecret {
	return v.value
}

func (v *NullableAPIServiceIntegrationInstanceSecret) Set(val *APIServiceIntegrationInstanceSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServiceIntegrationInstanceSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServiceIntegrationInstanceSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServiceIntegrationInstanceSecret(val *APIServiceIntegrationInstanceSecret) *NullableAPIServiceIntegrationInstanceSecret {
	return &NullableAPIServiceIntegrationInstanceSecret{value: val, isSet: true}
}

func (v NullableAPIServiceIntegrationInstanceSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServiceIntegrationInstanceSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

