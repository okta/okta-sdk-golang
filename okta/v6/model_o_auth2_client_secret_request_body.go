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

// OAuth2ClientSecretRequestBody struct for OAuth2ClientSecretRequestBody
type OAuth2ClientSecretRequestBody struct {
	// The OAuth 2.0 client secret string
	ClientSecret *string `json:"client_secret,omitempty"`
	// Status of the OAuth 2.0 Client Secret
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientSecretRequestBody OAuth2ClientSecretRequestBody

// NewOAuth2ClientSecretRequestBody instantiates a new OAuth2ClientSecretRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientSecretRequestBody() *OAuth2ClientSecretRequestBody {
	this := OAuth2ClientSecretRequestBody{}
	return &this
}

// NewOAuth2ClientSecretRequestBodyWithDefaults instantiates a new OAuth2ClientSecretRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientSecretRequestBodyWithDefaults() *OAuth2ClientSecretRequestBody {
	this := OAuth2ClientSecretRequestBody{}
	return &this
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *OAuth2ClientSecretRequestBody) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientSecretRequestBody) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *OAuth2ClientSecretRequestBody) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *OAuth2ClientSecretRequestBody) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OAuth2ClientSecretRequestBody) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientSecretRequestBody) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OAuth2ClientSecretRequestBody) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OAuth2ClientSecretRequestBody) SetStatus(v string) {
	o.Status = &v
}

func (o OAuth2ClientSecretRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2ClientSecretRequestBody) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2ClientSecretRequestBody := _OAuth2ClientSecretRequestBody{}

	err = json.Unmarshal(bytes, &varOAuth2ClientSecretRequestBody)
	if err == nil {
		*o = OAuth2ClientSecretRequestBody(varOAuth2ClientSecretRequestBody)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "client_secret")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOAuth2ClientSecretRequestBody struct {
	value *OAuth2ClientSecretRequestBody
	isSet bool
}

func (v NullableOAuth2ClientSecretRequestBody) Get() *OAuth2ClientSecretRequestBody {
	return v.value
}

func (v *NullableOAuth2ClientSecretRequestBody) Set(val *OAuth2ClientSecretRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientSecretRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientSecretRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientSecretRequestBody(val *OAuth2ClientSecretRequestBody) *NullableOAuth2ClientSecretRequestBody {
	return &NullableOAuth2ClientSecretRequestBody{value: val, isSet: true}
}

func (v NullableOAuth2ClientSecretRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientSecretRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

