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

// OAuth2ScopeConsentGrantEmbeddedScope struct for OAuth2ScopeConsentGrantEmbeddedScope
type OAuth2ScopeConsentGrantEmbeddedScope struct {
	// The name of the Okta scope for which consent is granted
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ScopeConsentGrantEmbeddedScope OAuth2ScopeConsentGrantEmbeddedScope

// NewOAuth2ScopeConsentGrantEmbeddedScope instantiates a new OAuth2ScopeConsentGrantEmbeddedScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ScopeConsentGrantEmbeddedScope() *OAuth2ScopeConsentGrantEmbeddedScope {
	this := OAuth2ScopeConsentGrantEmbeddedScope{}
	return &this
}

// NewOAuth2ScopeConsentGrantEmbeddedScopeWithDefaults instantiates a new OAuth2ScopeConsentGrantEmbeddedScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ScopeConsentGrantEmbeddedScopeWithDefaults() *OAuth2ScopeConsentGrantEmbeddedScope {
	this := OAuth2ScopeConsentGrantEmbeddedScope{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrantEmbeddedScope) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrantEmbeddedScope) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrantEmbeddedScope) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OAuth2ScopeConsentGrantEmbeddedScope) SetId(v string) {
	o.Id = &v
}

func (o OAuth2ScopeConsentGrantEmbeddedScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2ScopeConsentGrantEmbeddedScope) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2ScopeConsentGrantEmbeddedScope := _OAuth2ScopeConsentGrantEmbeddedScope{}

	err = json.Unmarshal(bytes, &varOAuth2ScopeConsentGrantEmbeddedScope)
	if err == nil {
		*o = OAuth2ScopeConsentGrantEmbeddedScope(varOAuth2ScopeConsentGrantEmbeddedScope)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOAuth2ScopeConsentGrantEmbeddedScope struct {
	value *OAuth2ScopeConsentGrantEmbeddedScope
	isSet bool
}

func (v NullableOAuth2ScopeConsentGrantEmbeddedScope) Get() *OAuth2ScopeConsentGrantEmbeddedScope {
	return v.value
}

func (v *NullableOAuth2ScopeConsentGrantEmbeddedScope) Set(val *OAuth2ScopeConsentGrantEmbeddedScope) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ScopeConsentGrantEmbeddedScope) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ScopeConsentGrantEmbeddedScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ScopeConsentGrantEmbeddedScope(val *OAuth2ScopeConsentGrantEmbeddedScope) *NullableOAuth2ScopeConsentGrantEmbeddedScope {
	return &NullableOAuth2ScopeConsentGrantEmbeddedScope{value: val, isSet: true}
}

func (v NullableOAuth2ScopeConsentGrantEmbeddedScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ScopeConsentGrantEmbeddedScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

