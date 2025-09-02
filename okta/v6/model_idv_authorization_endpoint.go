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

// IDVAuthorizationEndpoint IDV authorization endpoint
type IDVAuthorizationEndpoint struct {
	Binding *string `json:"binding,omitempty"`
	// URL of the IDV `authorization` endpoint
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IDVAuthorizationEndpoint IDVAuthorizationEndpoint

// NewIDVAuthorizationEndpoint instantiates a new IDVAuthorizationEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIDVAuthorizationEndpoint() *IDVAuthorizationEndpoint {
	this := IDVAuthorizationEndpoint{}
	return &this
}

// NewIDVAuthorizationEndpointWithDefaults instantiates a new IDVAuthorizationEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIDVAuthorizationEndpointWithDefaults() *IDVAuthorizationEndpoint {
	this := IDVAuthorizationEndpoint{}
	return &this
}

// GetBinding returns the Binding field value if set, zero value otherwise.
func (o *IDVAuthorizationEndpoint) GetBinding() string {
	if o == nil || o.Binding == nil {
		var ret string
		return ret
	}
	return *o.Binding
}

// GetBindingOk returns a tuple with the Binding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDVAuthorizationEndpoint) GetBindingOk() (*string, bool) {
	if o == nil || o.Binding == nil {
		return nil, false
	}
	return o.Binding, true
}

// HasBinding returns a boolean if a field has been set.
func (o *IDVAuthorizationEndpoint) HasBinding() bool {
	if o != nil && o.Binding != nil {
		return true
	}

	return false
}

// SetBinding gets a reference to the given string and assigns it to the Binding field.
func (o *IDVAuthorizationEndpoint) SetBinding(v string) {
	o.Binding = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *IDVAuthorizationEndpoint) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDVAuthorizationEndpoint) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *IDVAuthorizationEndpoint) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *IDVAuthorizationEndpoint) SetUrl(v string) {
	o.Url = &v
}

func (o IDVAuthorizationEndpoint) MarshalJSON() ([]byte, error) {
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

func (o *IDVAuthorizationEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	varIDVAuthorizationEndpoint := _IDVAuthorizationEndpoint{}

	err = json.Unmarshal(bytes, &varIDVAuthorizationEndpoint)
	if err == nil {
		*o = IDVAuthorizationEndpoint(varIDVAuthorizationEndpoint)
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

type NullableIDVAuthorizationEndpoint struct {
	value *IDVAuthorizationEndpoint
	isSet bool
}

func (v NullableIDVAuthorizationEndpoint) Get() *IDVAuthorizationEndpoint {
	return v.value
}

func (v *NullableIDVAuthorizationEndpoint) Set(val *IDVAuthorizationEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableIDVAuthorizationEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableIDVAuthorizationEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIDVAuthorizationEndpoint(val *IDVAuthorizationEndpoint) *NullableIDVAuthorizationEndpoint {
	return &NullableIDVAuthorizationEndpoint{value: val, isSet: true}
}

func (v NullableIDVAuthorizationEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIDVAuthorizationEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

