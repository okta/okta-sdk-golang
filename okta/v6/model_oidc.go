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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the Oidc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Oidc{}

// Oidc OIDC configuration details
type Oidc struct {
	// The URL to your customer-facing instructions for configuring your OIDC integration. See [Customer configuration document guidelines](https://developer.okta.com/docs/guides/submit-app-prereq/main/#customer-configuration-document-guidelines).
	Doc string `json:"doc"`
	// The URL to redirect users when they click on your app from their Okta End-User Dashboard
	InitiateLoginUri *string `json:"initiateLoginUri,omitempty"`
	// The sign-out redirect URIs for your app. You can send a request to `/v1/logout` to sign the user out and redirect them to one of these URIs.
	PostLogoutUris []string `json:"postLogoutUris,omitempty"`
	// List of sign-in redirect URIs
	RedirectUris         []string `json:"redirectUris"`
	AdditionalProperties map[string]interface{}
}

type _Oidc Oidc

// NewOidc instantiates a new Oidc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidc(doc string, redirectUris []string) *Oidc {
	this := Oidc{}
	this.Doc = doc
	this.RedirectUris = redirectUris
	return &this
}

// NewOidcWithDefaults instantiates a new Oidc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcWithDefaults() *Oidc {
	this := Oidc{}
	return &this
}

// GetDoc returns the Doc field value
func (o *Oidc) GetDoc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Doc
}

// GetDocOk returns a tuple with the Doc field value
// and a boolean to check if the value has been set.
func (o *Oidc) GetDocOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Doc, true
}

// SetDoc sets field value
func (o *Oidc) SetDoc(v string) {
	o.Doc = v
}

// GetInitiateLoginUri returns the InitiateLoginUri field value if set, zero value otherwise.
func (o *Oidc) GetInitiateLoginUri() string {
	if o == nil || IsNil(o.InitiateLoginUri) {
		var ret string
		return ret
	}
	return *o.InitiateLoginUri
}

// GetInitiateLoginUriOk returns a tuple with the InitiateLoginUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Oidc) GetInitiateLoginUriOk() (*string, bool) {
	if o == nil || IsNil(o.InitiateLoginUri) {
		return nil, false
	}
	return o.InitiateLoginUri, true
}

// HasInitiateLoginUri returns a boolean if a field has been set.
func (o *Oidc) HasInitiateLoginUri() bool {
	if o != nil && !IsNil(o.InitiateLoginUri) {
		return true
	}

	return false
}

// SetInitiateLoginUri gets a reference to the given string and assigns it to the InitiateLoginUri field.
func (o *Oidc) SetInitiateLoginUri(v string) {
	o.InitiateLoginUri = &v
}

// GetPostLogoutUris returns the PostLogoutUris field value if set, zero value otherwise.
func (o *Oidc) GetPostLogoutUris() []string {
	if o == nil || IsNil(o.PostLogoutUris) {
		var ret []string
		return ret
	}
	return o.PostLogoutUris
}

// GetPostLogoutUrisOk returns a tuple with the PostLogoutUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Oidc) GetPostLogoutUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.PostLogoutUris) {
		return nil, false
	}
	return o.PostLogoutUris, true
}

// HasPostLogoutUris returns a boolean if a field has been set.
func (o *Oidc) HasPostLogoutUris() bool {
	if o != nil && !IsNil(o.PostLogoutUris) {
		return true
	}

	return false
}

// SetPostLogoutUris gets a reference to the given []string and assigns it to the PostLogoutUris field.
func (o *Oidc) SetPostLogoutUris(v []string) {
	o.PostLogoutUris = v
}

// GetRedirectUris returns the RedirectUris field value
func (o *Oidc) GetRedirectUris() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value
// and a boolean to check if the value has been set.
func (o *Oidc) GetRedirectUrisOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// SetRedirectUris sets field value
func (o *Oidc) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

func (o Oidc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Oidc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["doc"] = o.Doc
	if !IsNil(o.InitiateLoginUri) {
		toSerialize["initiateLoginUri"] = o.InitiateLoginUri
	}
	if !IsNil(o.PostLogoutUris) {
		toSerialize["postLogoutUris"] = o.PostLogoutUris
	}
	toSerialize["redirectUris"] = o.RedirectUris

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Oidc) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"doc",
		"redirectUris",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOidc := _Oidc{}

	err = json.Unmarshal(data, &varOidc)

	if err != nil {
		return err
	}

	*o = Oidc(varOidc)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "doc")
		delete(additionalProperties, "initiateLoginUri")
		delete(additionalProperties, "postLogoutUris")
		delete(additionalProperties, "redirectUris")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOidc struct {
	value *Oidc
	isSet bool
}

func (v NullableOidc) Get() *Oidc {
	return v.value
}

func (v *NullableOidc) Set(val *Oidc) {
	v.value = val
	v.isSet = true
}

func (v NullableOidc) IsSet() bool {
	return v.isSet
}

func (v *NullableOidc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidc(val *Oidc) *NullableOidc {
	return &NullableOidc{value: val, isSet: true}
}

func (v NullableOidc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
