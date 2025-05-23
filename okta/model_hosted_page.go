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

// HostedPage struct for HostedPage
type HostedPage struct {
	Type string `json:"type"`
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HostedPage HostedPage

// NewHostedPage instantiates a new HostedPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostedPage(type_ string) *HostedPage {
	this := HostedPage{}
	this.Type = type_
	return &this
}

// NewHostedPageWithDefaults instantiates a new HostedPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostedPageWithDefaults() *HostedPage {
	this := HostedPage{}
	return &this
}

// GetType returns the Type field value
func (o *HostedPage) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HostedPage) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HostedPage) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *HostedPage) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostedPage) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *HostedPage) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *HostedPage) SetUrl(v string) {
	o.Url = &v
}

func (o HostedPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HostedPage) UnmarshalJSON(bytes []byte) (err error) {
	varHostedPage := _HostedPage{}

	err = json.Unmarshal(bytes, &varHostedPage)
	if err == nil {
		*o = HostedPage(varHostedPage)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableHostedPage struct {
	value *HostedPage
	isSet bool
}

func (v NullableHostedPage) Get() *HostedPage {
	return v.value
}

func (v *NullableHostedPage) Set(val *HostedPage) {
	v.value = val
	v.isSet = true
}

func (v NullableHostedPage) IsSet() bool {
	return v.isSet
}

func (v *NullableHostedPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostedPage(val *HostedPage) *NullableHostedPage {
	return &NullableHostedPage{value: val, isSet: true}
}

func (v NullableHostedPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostedPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

