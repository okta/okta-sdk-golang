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

// InlineHookOAuthChannelConfigCreate struct for InlineHookOAuthChannelConfigCreate
type InlineHookOAuthChannelConfigCreate struct {
	// The authentication method for the token endpoint
	AuthType *string `json:"authType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookOAuthChannelConfigCreate InlineHookOAuthChannelConfigCreate

// NewInlineHookOAuthChannelConfigCreate instantiates a new InlineHookOAuthChannelConfigCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookOAuthChannelConfigCreate() *InlineHookOAuthChannelConfigCreate {
	this := InlineHookOAuthChannelConfigCreate{}
	return &this
}

// NewInlineHookOAuthChannelConfigCreateWithDefaults instantiates a new InlineHookOAuthChannelConfigCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookOAuthChannelConfigCreateWithDefaults() *InlineHookOAuthChannelConfigCreate {
	this := InlineHookOAuthChannelConfigCreate{}
	return &this
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *InlineHookOAuthChannelConfigCreate) GetAuthType() string {
	if o == nil || o.AuthType == nil {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthChannelConfigCreate) GetAuthTypeOk() (*string, bool) {
	if o == nil || o.AuthType == nil {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *InlineHookOAuthChannelConfigCreate) HasAuthType() bool {
	if o != nil && o.AuthType != nil {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *InlineHookOAuthChannelConfigCreate) SetAuthType(v string) {
	o.AuthType = &v
}

func (o InlineHookOAuthChannelConfigCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthType != nil {
		toSerialize["authType"] = o.AuthType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineHookOAuthChannelConfigCreate) UnmarshalJSON(bytes []byte) (err error) {
	varInlineHookOAuthChannelConfigCreate := _InlineHookOAuthChannelConfigCreate{}

	err = json.Unmarshal(bytes, &varInlineHookOAuthChannelConfigCreate)
	if err == nil {
		*o = InlineHookOAuthChannelConfigCreate(varInlineHookOAuthChannelConfigCreate)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableInlineHookOAuthChannelConfigCreate struct {
	value *InlineHookOAuthChannelConfigCreate
	isSet bool
}

func (v NullableInlineHookOAuthChannelConfigCreate) Get() *InlineHookOAuthChannelConfigCreate {
	return v.value
}

func (v *NullableInlineHookOAuthChannelConfigCreate) Set(val *InlineHookOAuthChannelConfigCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookOAuthChannelConfigCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookOAuthChannelConfigCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookOAuthChannelConfigCreate(val *InlineHookOAuthChannelConfigCreate) *NullableInlineHookOAuthChannelConfigCreate {
	return &NullableInlineHookOAuthChannelConfigCreate{value: val, isSet: true}
}

func (v NullableInlineHookOAuthChannelConfigCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookOAuthChannelConfigCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

