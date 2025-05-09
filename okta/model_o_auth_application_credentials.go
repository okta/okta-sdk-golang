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

// OAuthApplicationCredentials struct for OAuthApplicationCredentials
type OAuthApplicationCredentials struct {
	Signing *ApplicationCredentialsSigning `json:"signing,omitempty"`
	UserNameTemplate *ApplicationCredentialsUsernameTemplate `json:"userNameTemplate,omitempty"`
	OauthClient *ApplicationCredentialsOAuthClient `json:"oauthClient,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuthApplicationCredentials OAuthApplicationCredentials

// NewOAuthApplicationCredentials instantiates a new OAuthApplicationCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthApplicationCredentials() *OAuthApplicationCredentials {
	this := OAuthApplicationCredentials{}
	return &this
}

// NewOAuthApplicationCredentialsWithDefaults instantiates a new OAuthApplicationCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthApplicationCredentialsWithDefaults() *OAuthApplicationCredentials {
	this := OAuthApplicationCredentials{}
	return &this
}

// GetSigning returns the Signing field value if set, zero value otherwise.
func (o *OAuthApplicationCredentials) GetSigning() ApplicationCredentialsSigning {
	if o == nil || o.Signing == nil {
		var ret ApplicationCredentialsSigning
		return ret
	}
	return *o.Signing
}

// GetSigningOk returns a tuple with the Signing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthApplicationCredentials) GetSigningOk() (*ApplicationCredentialsSigning, bool) {
	if o == nil || o.Signing == nil {
		return nil, false
	}
	return o.Signing, true
}

// HasSigning returns a boolean if a field has been set.
func (o *OAuthApplicationCredentials) HasSigning() bool {
	if o != nil && o.Signing != nil {
		return true
	}

	return false
}

// SetSigning gets a reference to the given ApplicationCredentialsSigning and assigns it to the Signing field.
func (o *OAuthApplicationCredentials) SetSigning(v ApplicationCredentialsSigning) {
	o.Signing = &v
}

// GetUserNameTemplate returns the UserNameTemplate field value if set, zero value otherwise.
func (o *OAuthApplicationCredentials) GetUserNameTemplate() ApplicationCredentialsUsernameTemplate {
	if o == nil || o.UserNameTemplate == nil {
		var ret ApplicationCredentialsUsernameTemplate
		return ret
	}
	return *o.UserNameTemplate
}

// GetUserNameTemplateOk returns a tuple with the UserNameTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthApplicationCredentials) GetUserNameTemplateOk() (*ApplicationCredentialsUsernameTemplate, bool) {
	if o == nil || o.UserNameTemplate == nil {
		return nil, false
	}
	return o.UserNameTemplate, true
}

// HasUserNameTemplate returns a boolean if a field has been set.
func (o *OAuthApplicationCredentials) HasUserNameTemplate() bool {
	if o != nil && o.UserNameTemplate != nil {
		return true
	}

	return false
}

// SetUserNameTemplate gets a reference to the given ApplicationCredentialsUsernameTemplate and assigns it to the UserNameTemplate field.
func (o *OAuthApplicationCredentials) SetUserNameTemplate(v ApplicationCredentialsUsernameTemplate) {
	o.UserNameTemplate = &v
}

// GetOauthClient returns the OauthClient field value if set, zero value otherwise.
func (o *OAuthApplicationCredentials) GetOauthClient() ApplicationCredentialsOAuthClient {
	if o == nil || o.OauthClient == nil {
		var ret ApplicationCredentialsOAuthClient
		return ret
	}
	return *o.OauthClient
}

// GetOauthClientOk returns a tuple with the OauthClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthApplicationCredentials) GetOauthClientOk() (*ApplicationCredentialsOAuthClient, bool) {
	if o == nil || o.OauthClient == nil {
		return nil, false
	}
	return o.OauthClient, true
}

// HasOauthClient returns a boolean if a field has been set.
func (o *OAuthApplicationCredentials) HasOauthClient() bool {
	if o != nil && o.OauthClient != nil {
		return true
	}

	return false
}

// SetOauthClient gets a reference to the given ApplicationCredentialsOAuthClient and assigns it to the OauthClient field.
func (o *OAuthApplicationCredentials) SetOauthClient(v ApplicationCredentialsOAuthClient) {
	o.OauthClient = &v
}

func (o OAuthApplicationCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Signing != nil {
		toSerialize["signing"] = o.Signing
	}
	if o.UserNameTemplate != nil {
		toSerialize["userNameTemplate"] = o.UserNameTemplate
	}
	if o.OauthClient != nil {
		toSerialize["oauthClient"] = o.OauthClient
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuthApplicationCredentials) UnmarshalJSON(bytes []byte) (err error) {
	varOAuthApplicationCredentials := _OAuthApplicationCredentials{}

	err = json.Unmarshal(bytes, &varOAuthApplicationCredentials)
	if err == nil {
		*o = OAuthApplicationCredentials(varOAuthApplicationCredentials)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "signing")
		delete(additionalProperties, "userNameTemplate")
		delete(additionalProperties, "oauthClient")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOAuthApplicationCredentials struct {
	value *OAuthApplicationCredentials
	isSet bool
}

func (v NullableOAuthApplicationCredentials) Get() *OAuthApplicationCredentials {
	return v.value
}

func (v *NullableOAuthApplicationCredentials) Set(val *OAuthApplicationCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthApplicationCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthApplicationCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthApplicationCredentials(val *OAuthApplicationCredentials) *NullableOAuthApplicationCredentials {
	return &NullableOAuthApplicationCredentials{value: val, isSet: true}
}

func (v NullableOAuthApplicationCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthApplicationCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

