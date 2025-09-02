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

// OidcJwksEndpoint Endpoint for the JSON Web Key Set (JWKS) document. This document contains signing keys that are used to validate the signatures from the provider. For more information on JWKS, see [JSON Web Key](https://tools.ietf.org/html/rfc7517).
type OidcJwksEndpoint struct {
	Binding *string `json:"binding,omitempty"`
	// URL of the endpoint to the JWK Set
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OidcJwksEndpoint OidcJwksEndpoint

// NewOidcJwksEndpoint instantiates a new OidcJwksEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcJwksEndpoint() *OidcJwksEndpoint {
	this := OidcJwksEndpoint{}
	return &this
}

// NewOidcJwksEndpointWithDefaults instantiates a new OidcJwksEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcJwksEndpointWithDefaults() *OidcJwksEndpoint {
	this := OidcJwksEndpoint{}
	return &this
}

// GetBinding returns the Binding field value if set, zero value otherwise.
func (o *OidcJwksEndpoint) GetBinding() string {
	if o == nil || o.Binding == nil {
		var ret string
		return ret
	}
	return *o.Binding
}

// GetBindingOk returns a tuple with the Binding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcJwksEndpoint) GetBindingOk() (*string, bool) {
	if o == nil || o.Binding == nil {
		return nil, false
	}
	return o.Binding, true
}

// HasBinding returns a boolean if a field has been set.
func (o *OidcJwksEndpoint) HasBinding() bool {
	if o != nil && o.Binding != nil {
		return true
	}

	return false
}

// SetBinding gets a reference to the given string and assigns it to the Binding field.
func (o *OidcJwksEndpoint) SetBinding(v string) {
	o.Binding = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OidcJwksEndpoint) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcJwksEndpoint) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OidcJwksEndpoint) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OidcJwksEndpoint) SetUrl(v string) {
	o.Url = &v
}

func (o OidcJwksEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Binding != nil {
		toSerialize["binding"] = o.Binding
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OidcJwksEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	varOidcJwksEndpoint := _OidcJwksEndpoint{}

	err = json.Unmarshal(bytes, &varOidcJwksEndpoint)
	if err == nil {
		*o = OidcJwksEndpoint(varOidcJwksEndpoint)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "binding")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOidcJwksEndpoint struct {
	value *OidcJwksEndpoint
	isSet bool
}

func (v NullableOidcJwksEndpoint) Get() *OidcJwksEndpoint {
	return v.value
}

func (v *NullableOidcJwksEndpoint) Set(val *OidcJwksEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcJwksEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcJwksEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcJwksEndpoint(val *OidcJwksEndpoint) *NullableOidcJwksEndpoint {
	return &NullableOidcJwksEndpoint{value: val, isSet: true}
}

func (v NullableOidcJwksEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcJwksEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

