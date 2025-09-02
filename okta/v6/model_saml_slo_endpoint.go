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

// SamlSloEndpoint IdP's `SingleLogoutService` endpoint where Okta sends a `<LogoutRequest>` message
type SamlSloEndpoint struct {
	Binding *string `json:"binding,omitempty"`
	// URL of the binding-specific IdP endpoint where Okta sends a `<LogoutRequest>`
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlSloEndpoint SamlSloEndpoint

// NewSamlSloEndpoint instantiates a new SamlSloEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlSloEndpoint() *SamlSloEndpoint {
	this := SamlSloEndpoint{}
	return &this
}

// NewSamlSloEndpointWithDefaults instantiates a new SamlSloEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlSloEndpointWithDefaults() *SamlSloEndpoint {
	this := SamlSloEndpoint{}
	return &this
}

// GetBinding returns the Binding field value if set, zero value otherwise.
func (o *SamlSloEndpoint) GetBinding() string {
	if o == nil || o.Binding == nil {
		var ret string
		return ret
	}
	return *o.Binding
}

// GetBindingOk returns a tuple with the Binding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSloEndpoint) GetBindingOk() (*string, bool) {
	if o == nil || o.Binding == nil {
		return nil, false
	}
	return o.Binding, true
}

// HasBinding returns a boolean if a field has been set.
func (o *SamlSloEndpoint) HasBinding() bool {
	if o != nil && o.Binding != nil {
		return true
	}

	return false
}

// SetBinding gets a reference to the given string and assigns it to the Binding field.
func (o *SamlSloEndpoint) SetBinding(v string) {
	o.Binding = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SamlSloEndpoint) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSloEndpoint) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SamlSloEndpoint) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SamlSloEndpoint) SetUrl(v string) {
	o.Url = &v
}

func (o SamlSloEndpoint) MarshalJSON() ([]byte, error) {
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

func (o *SamlSloEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	varSamlSloEndpoint := _SamlSloEndpoint{}

	err = json.Unmarshal(bytes, &varSamlSloEndpoint)
	if err == nil {
		*o = SamlSloEndpoint(varSamlSloEndpoint)
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

type NullableSamlSloEndpoint struct {
	value *SamlSloEndpoint
	isSet bool
}

func (v NullableSamlSloEndpoint) Get() *SamlSloEndpoint {
	return v.value
}

func (v *NullableSamlSloEndpoint) Set(val *SamlSloEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlSloEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlSloEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlSloEndpoint(val *SamlSloEndpoint) *NullableSamlSloEndpoint {
	return &NullableSamlSloEndpoint{value: val, isSet: true}
}

func (v NullableSamlSloEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlSloEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

