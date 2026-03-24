/*
Okta Admin Management API

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

// checks if the ApiService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiService{}

// ApiService struct for ApiService
type ApiService struct {
	// Specifies the authentication method used by the API service application when requesting an access token<br> **Note:** Only the `client_secret_basic` method is supported
	AuthenticationMethod *string `json:"authenticationMethod,omitempty"`
	// A list of Okta OAuth 2.0 scopes required for the API service app to function
	Scopes []string `json:"scopes,omitempty"`
	// The URL for the API service integration configuration document
	SetupInstructionsUri *string `json:"setupInstructionsUri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApiService ApiService

// NewApiService instantiates a new ApiService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiService() *ApiService {
	this := ApiService{}
	return &this
}

// NewApiServiceWithDefaults instantiates a new ApiService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiServiceWithDefaults() *ApiService {
	this := ApiService{}
	return &this
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *ApiService) GetAuthenticationMethod() string {
	if o == nil || IsNil(o.AuthenticationMethod) {
		var ret string
		return ret
	}
	return *o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiService) GetAuthenticationMethodOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationMethod) {
		return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *ApiService) HasAuthenticationMethod() bool {
	if o != nil && !IsNil(o.AuthenticationMethod) {
		return true
	}

	return false
}

// SetAuthenticationMethod gets a reference to the given string and assigns it to the AuthenticationMethod field.
func (o *ApiService) SetAuthenticationMethod(v string) {
	o.AuthenticationMethod = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ApiService) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiService) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ApiService) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ApiService) SetScopes(v []string) {
	o.Scopes = v
}

// GetSetupInstructionsUri returns the SetupInstructionsUri field value if set, zero value otherwise.
func (o *ApiService) GetSetupInstructionsUri() string {
	if o == nil || IsNil(o.SetupInstructionsUri) {
		var ret string
		return ret
	}
	return *o.SetupInstructionsUri
}

// GetSetupInstructionsUriOk returns a tuple with the SetupInstructionsUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiService) GetSetupInstructionsUriOk() (*string, bool) {
	if o == nil || IsNil(o.SetupInstructionsUri) {
		return nil, false
	}
	return o.SetupInstructionsUri, true
}

// HasSetupInstructionsUri returns a boolean if a field has been set.
func (o *ApiService) HasSetupInstructionsUri() bool {
	if o != nil && !IsNil(o.SetupInstructionsUri) {
		return true
	}

	return false
}

// SetSetupInstructionsUri gets a reference to the given string and assigns it to the SetupInstructionsUri field.
func (o *ApiService) SetSetupInstructionsUri(v string) {
	o.SetupInstructionsUri = &v
}

func (o ApiService) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenticationMethod) {
		toSerialize["authenticationMethod"] = o.AuthenticationMethod
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.SetupInstructionsUri) {
		toSerialize["setupInstructionsUri"] = o.SetupInstructionsUri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApiService) UnmarshalJSON(data []byte) (err error) {
	varApiService := _ApiService{}

	err = json.Unmarshal(data, &varApiService)

	if err != nil {
		return err
	}

	*o = ApiService(varApiService)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticationMethod")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "setupInstructionsUri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApiService struct {
	value *ApiService
	isSet bool
}

func (v NullableApiService) Get() *ApiService {
	return v.value
}

func (v *NullableApiService) Set(val *ApiService) {
	v.value = val
	v.isSet = true
}

func (v NullableApiService) IsSet() bool {
	return v.isSet
}

func (v *NullableApiService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiService(val *ApiService) *NullableApiService {
	return &NullableApiService{value: val, isSet: true}
}

func (v NullableApiService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
