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

// OAuth2Client struct for OAuth2Client
type OAuth2Client struct {
	// Unique key for the client application. The `client_id` is immutable
	ClientId *string `json:"client_id,omitempty"`
	// Human-readable string name of the client application
	ClientName *string `json:"client_name,omitempty"`
	ClientUri *string `json:"client_uri,omitempty"`
	// URL string that references a logo for the client consent dialog (not the sign-in dialog)
	LogoUri *string `json:"logo_uri,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2Client OAuth2Client

// NewOAuth2Client instantiates a new OAuth2Client object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2Client() *OAuth2Client {
	this := OAuth2Client{}
	return &this
}

// NewOAuth2ClientWithDefaults instantiates a new OAuth2Client object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientWithDefaults() *OAuth2Client {
	this := OAuth2Client{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OAuth2Client) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Client) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OAuth2Client) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OAuth2Client) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *OAuth2Client) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Client) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *OAuth2Client) HasClientName() bool {
	if o != nil && o.ClientName != nil {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *OAuth2Client) SetClientName(v string) {
	o.ClientName = &v
}

// GetClientUri returns the ClientUri field value if set, zero value otherwise.
func (o *OAuth2Client) GetClientUri() string {
	if o == nil || o.ClientUri == nil {
		var ret string
		return ret
	}
	return *o.ClientUri
}

// GetClientUriOk returns a tuple with the ClientUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Client) GetClientUriOk() (*string, bool) {
	if o == nil || o.ClientUri == nil {
		return nil, false
	}
	return o.ClientUri, true
}

// HasClientUri returns a boolean if a field has been set.
func (o *OAuth2Client) HasClientUri() bool {
	if o != nil && o.ClientUri != nil {
		return true
	}

	return false
}

// SetClientUri gets a reference to the given string and assigns it to the ClientUri field.
func (o *OAuth2Client) SetClientUri(v string) {
	o.ClientUri = &v
}

// GetLogoUri returns the LogoUri field value if set, zero value otherwise.
func (o *OAuth2Client) GetLogoUri() string {
	if o == nil || o.LogoUri == nil {
		var ret string
		return ret
	}
	return *o.LogoUri
}

// GetLogoUriOk returns a tuple with the LogoUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Client) GetLogoUriOk() (*string, bool) {
	if o == nil || o.LogoUri == nil {
		return nil, false
	}
	return o.LogoUri, true
}

// HasLogoUri returns a boolean if a field has been set.
func (o *OAuth2Client) HasLogoUri() bool {
	if o != nil && o.LogoUri != nil {
		return true
	}

	return false
}

// SetLogoUri gets a reference to the given string and assigns it to the LogoUri field.
func (o *OAuth2Client) SetLogoUri(v string) {
	o.LogoUri = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OAuth2Client) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Client) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OAuth2Client) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *OAuth2Client) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o OAuth2Client) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientName != nil {
		toSerialize["client_name"] = o.ClientName
	}
	if o.ClientUri != nil {
		toSerialize["client_uri"] = o.ClientUri
	}
	if o.LogoUri != nil {
		toSerialize["logo_uri"] = o.LogoUri
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2Client) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2Client := _OAuth2Client{}

	err = json.Unmarshal(bytes, &varOAuth2Client)
	if err == nil {
		*o = OAuth2Client(varOAuth2Client)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "client_name")
		delete(additionalProperties, "client_uri")
		delete(additionalProperties, "logo_uri")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOAuth2Client struct {
	value *OAuth2Client
	isSet bool
}

func (v NullableOAuth2Client) Get() *OAuth2Client {
	return v.value
}

func (v *NullableOAuth2Client) Set(val *OAuth2Client) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2Client) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2Client) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2Client(val *OAuth2Client) *NullableOAuth2Client {
	return &NullableOAuth2Client{value: val, isSet: true}
}

func (v NullableOAuth2Client) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2Client) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

