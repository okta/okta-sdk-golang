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

// OidcSloEndpoint OIDC IdP logout endpoint
type OidcSloEndpoint struct {
	// IdP logout endpoint URL
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OidcSloEndpoint OidcSloEndpoint

// NewOidcSloEndpoint instantiates a new OidcSloEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcSloEndpoint() *OidcSloEndpoint {
	this := OidcSloEndpoint{}
	return &this
}

// NewOidcSloEndpointWithDefaults instantiates a new OidcSloEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcSloEndpointWithDefaults() *OidcSloEndpoint {
	this := OidcSloEndpoint{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OidcSloEndpoint) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcSloEndpoint) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OidcSloEndpoint) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OidcSloEndpoint) SetUrl(v string) {
	o.Url = &v
}

func (o OidcSloEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OidcSloEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	varOidcSloEndpoint := _OidcSloEndpoint{}

	err = json.Unmarshal(bytes, &varOidcSloEndpoint)
	if err == nil {
		*o = OidcSloEndpoint(varOidcSloEndpoint)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOidcSloEndpoint struct {
	value *OidcSloEndpoint
	isSet bool
}

func (v NullableOidcSloEndpoint) Get() *OidcSloEndpoint {
	return v.value
}

func (v *NullableOidcSloEndpoint) Set(val *OidcSloEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcSloEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcSloEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcSloEndpoint(val *OidcSloEndpoint) *NullableOidcSloEndpoint {
	return &NullableOidcSloEndpoint{value: val, isSet: true}
}

func (v NullableOidcSloEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcSloEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

