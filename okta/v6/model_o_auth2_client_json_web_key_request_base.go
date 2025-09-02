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

// OAuth2ClientJsonWebKeyRequestBase struct for OAuth2ClientJsonWebKeyRequestBase
type OAuth2ClientJsonWebKeyRequestBase struct {
	// Unique identifier of the JSON Web Key in the OAUth 2.0 client's JWKS
	Kid NullableString `json:"kid,omitempty"`
	// Status of the OAuth 2.0 client JSON Web Key
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientJsonWebKeyRequestBase OAuth2ClientJsonWebKeyRequestBase

// NewOAuth2ClientJsonWebKeyRequestBase instantiates a new OAuth2ClientJsonWebKeyRequestBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientJsonWebKeyRequestBase() *OAuth2ClientJsonWebKeyRequestBase {
	this := OAuth2ClientJsonWebKeyRequestBase{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// NewOAuth2ClientJsonWebKeyRequestBaseWithDefaults instantiates a new OAuth2ClientJsonWebKeyRequestBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientJsonWebKeyRequestBaseWithDefaults() *OAuth2ClientJsonWebKeyRequestBase {
	this := OAuth2ClientJsonWebKeyRequestBase{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// GetKid returns the Kid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientJsonWebKeyRequestBase) GetKid() string {
	if o == nil || o.Kid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Kid.Get()
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientJsonWebKeyRequestBase) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kid.Get(), o.Kid.IsSet()
}

// HasKid returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyRequestBase) HasKid() bool {
	if o != nil && o.Kid.IsSet() {
		return true
	}

	return false
}

// SetKid gets a reference to the given NullableString and assigns it to the Kid field.
func (o *OAuth2ClientJsonWebKeyRequestBase) SetKid(v string) {
	o.Kid.Set(&v)
}
// SetKidNil sets the value for Kid to be an explicit nil
func (o *OAuth2ClientJsonWebKeyRequestBase) SetKidNil() {
	o.Kid.Set(nil)
}

// UnsetKid ensures that no value is present for Kid, not even an explicit nil
func (o *OAuth2ClientJsonWebKeyRequestBase) UnsetKid() {
	o.Kid.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyRequestBase) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyRequestBase) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyRequestBase) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OAuth2ClientJsonWebKeyRequestBase) SetStatus(v string) {
	o.Status = &v
}

func (o OAuth2ClientJsonWebKeyRequestBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kid.IsSet() {
		toSerialize["kid"] = o.Kid.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2ClientJsonWebKeyRequestBase) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2ClientJsonWebKeyRequestBase := _OAuth2ClientJsonWebKeyRequestBase{}

	err = json.Unmarshal(bytes, &varOAuth2ClientJsonWebKeyRequestBase)
	if err == nil {
		*o = OAuth2ClientJsonWebKeyRequestBase(varOAuth2ClientJsonWebKeyRequestBase)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "kid")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOAuth2ClientJsonWebKeyRequestBase struct {
	value *OAuth2ClientJsonWebKeyRequestBase
	isSet bool
}

func (v NullableOAuth2ClientJsonWebKeyRequestBase) Get() *OAuth2ClientJsonWebKeyRequestBase {
	return v.value
}

func (v *NullableOAuth2ClientJsonWebKeyRequestBase) Set(val *OAuth2ClientJsonWebKeyRequestBase) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientJsonWebKeyRequestBase) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientJsonWebKeyRequestBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientJsonWebKeyRequestBase(val *OAuth2ClientJsonWebKeyRequestBase) *NullableOAuth2ClientJsonWebKeyRequestBase {
	return &NullableOAuth2ClientJsonWebKeyRequestBase{value: val, isSet: true}
}

func (v NullableOAuth2ClientJsonWebKeyRequestBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientJsonWebKeyRequestBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

