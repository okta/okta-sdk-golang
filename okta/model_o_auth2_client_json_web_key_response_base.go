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

// checks if the OAuth2ClientJsonWebKeyResponseBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ClientJsonWebKeyResponseBase{}

// OAuth2ClientJsonWebKeyResponseBase struct for OAuth2ClientJsonWebKeyResponseBase
type OAuth2ClientJsonWebKeyResponseBase struct {
	// Timestamp when the OAuth 2.0 client JSON Web Key was created
	Created *string `json:"created,omitempty"`
	// The unique ID of the OAuth Client JSON Web Key
	Id *string `json:"id,omitempty"`
	// Timestamp when the OAuth 2.0 client JSON Web Key was updated
	LastUpdated          *string                 `json:"lastUpdated,omitempty"`
	Links                *OAuthClientSecretLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientJsonWebKeyResponseBase OAuth2ClientJsonWebKeyResponseBase

// NewOAuth2ClientJsonWebKeyResponseBase instantiates a new OAuth2ClientJsonWebKeyResponseBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientJsonWebKeyResponseBase() *OAuth2ClientJsonWebKeyResponseBase {
	this := OAuth2ClientJsonWebKeyResponseBase{}
	return &this
}

// NewOAuth2ClientJsonWebKeyResponseBaseWithDefaults instantiates a new OAuth2ClientJsonWebKeyResponseBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientJsonWebKeyResponseBaseWithDefaults() *OAuth2ClientJsonWebKeyResponseBase {
	this := OAuth2ClientJsonWebKeyResponseBase{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyResponseBase) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyResponseBase) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyResponseBase) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *OAuth2ClientJsonWebKeyResponseBase) SetCreated(v string) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyResponseBase) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyResponseBase) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyResponseBase) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OAuth2ClientJsonWebKeyResponseBase) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyResponseBase) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyResponseBase) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyResponseBase) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *OAuth2ClientJsonWebKeyResponseBase) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyResponseBase) GetLinks() OAuthClientSecretLinks {
	if o == nil || IsNil(o.Links) {
		var ret OAuthClientSecretLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyResponseBase) GetLinksOk() (*OAuthClientSecretLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyResponseBase) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OAuthClientSecretLinks and assigns it to the Links field.
func (o *OAuth2ClientJsonWebKeyResponseBase) SetLinks(v OAuthClientSecretLinks) {
	o.Links = &v
}

func (o OAuth2ClientJsonWebKeyResponseBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ClientJsonWebKeyResponseBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *OAuth2ClientJsonWebKeyResponseBase) UnmarshalJSON(data []byte) (err error) {
	varOAuth2ClientJsonWebKeyResponseBase := _OAuth2ClientJsonWebKeyResponseBase{}

	err = json.Unmarshal(data, &varOAuth2ClientJsonWebKeyResponseBase)

	if err != nil {
		return err
	}

	*o = OAuth2ClientJsonWebKeyResponseBase(varOAuth2ClientJsonWebKeyResponseBase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2ClientJsonWebKeyResponseBase struct {
	value *OAuth2ClientJsonWebKeyResponseBase
	isSet bool
}

func (v NullableOAuth2ClientJsonWebKeyResponseBase) Get() *OAuth2ClientJsonWebKeyResponseBase {
	return v.value
}

func (v *NullableOAuth2ClientJsonWebKeyResponseBase) Set(val *OAuth2ClientJsonWebKeyResponseBase) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientJsonWebKeyResponseBase) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientJsonWebKeyResponseBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientJsonWebKeyResponseBase(val *OAuth2ClientJsonWebKeyResponseBase) *NullableOAuth2ClientJsonWebKeyResponseBase {
	return &NullableOAuth2ClientJsonWebKeyResponseBase{value: val, isSet: true}
}

func (v NullableOAuth2ClientJsonWebKeyResponseBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientJsonWebKeyResponseBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
