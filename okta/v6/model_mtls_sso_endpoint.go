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

// MtlsSsoEndpoint The Single Sign-On (SSO) endpoint is the IdP's `SingleSignOnService` endpoint
type MtlsSsoEndpoint struct {
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MtlsSsoEndpoint MtlsSsoEndpoint

// NewMtlsSsoEndpoint instantiates a new MtlsSsoEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtlsSsoEndpoint() *MtlsSsoEndpoint {
	this := MtlsSsoEndpoint{}
	return &this
}

// NewMtlsSsoEndpointWithDefaults instantiates a new MtlsSsoEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtlsSsoEndpointWithDefaults() *MtlsSsoEndpoint {
	this := MtlsSsoEndpoint{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *MtlsSsoEndpoint) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtlsSsoEndpoint) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MtlsSsoEndpoint) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MtlsSsoEndpoint) SetUrl(v string) {
	o.Url = &v
}

func (o MtlsSsoEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MtlsSsoEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	varMtlsSsoEndpoint := _MtlsSsoEndpoint{}

	err = json.Unmarshal(bytes, &varMtlsSsoEndpoint)
	if err == nil {
		*o = MtlsSsoEndpoint(varMtlsSsoEndpoint)
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

type NullableMtlsSsoEndpoint struct {
	value *MtlsSsoEndpoint
	isSet bool
}

func (v NullableMtlsSsoEndpoint) Get() *MtlsSsoEndpoint {
	return v.value
}

func (v *NullableMtlsSsoEndpoint) Set(val *MtlsSsoEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableMtlsSsoEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableMtlsSsoEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtlsSsoEndpoint(val *MtlsSsoEndpoint) *NullableMtlsSsoEndpoint {
	return &NullableMtlsSsoEndpoint{value: val, isSet: true}
}

func (v NullableMtlsSsoEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtlsSsoEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

