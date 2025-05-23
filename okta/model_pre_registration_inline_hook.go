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

// PreRegistrationInlineHook struct for PreRegistrationInlineHook
type PreRegistrationInlineHook struct {
	InlineHookId *string `json:"inlineHookId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PreRegistrationInlineHook PreRegistrationInlineHook

// NewPreRegistrationInlineHook instantiates a new PreRegistrationInlineHook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreRegistrationInlineHook() *PreRegistrationInlineHook {
	this := PreRegistrationInlineHook{}
	return &this
}

// NewPreRegistrationInlineHookWithDefaults instantiates a new PreRegistrationInlineHook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreRegistrationInlineHookWithDefaults() *PreRegistrationInlineHook {
	this := PreRegistrationInlineHook{}
	return &this
}

// GetInlineHookId returns the InlineHookId field value if set, zero value otherwise.
func (o *PreRegistrationInlineHook) GetInlineHookId() string {
	if o == nil || o.InlineHookId == nil {
		var ret string
		return ret
	}
	return *o.InlineHookId
}

// GetInlineHookIdOk returns a tuple with the InlineHookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegistrationInlineHook) GetInlineHookIdOk() (*string, bool) {
	if o == nil || o.InlineHookId == nil {
		return nil, false
	}
	return o.InlineHookId, true
}

// HasInlineHookId returns a boolean if a field has been set.
func (o *PreRegistrationInlineHook) HasInlineHookId() bool {
	if o != nil && o.InlineHookId != nil {
		return true
	}

	return false
}

// SetInlineHookId gets a reference to the given string and assigns it to the InlineHookId field.
func (o *PreRegistrationInlineHook) SetInlineHookId(v string) {
	o.InlineHookId = &v
}

func (o PreRegistrationInlineHook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InlineHookId != nil {
		toSerialize["inlineHookId"] = o.InlineHookId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PreRegistrationInlineHook) UnmarshalJSON(bytes []byte) (err error) {
	varPreRegistrationInlineHook := _PreRegistrationInlineHook{}

	err = json.Unmarshal(bytes, &varPreRegistrationInlineHook)
	if err == nil {
		*o = PreRegistrationInlineHook(varPreRegistrationInlineHook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "inlineHookId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePreRegistrationInlineHook struct {
	value *PreRegistrationInlineHook
	isSet bool
}

func (v NullablePreRegistrationInlineHook) Get() *PreRegistrationInlineHook {
	return v.value
}

func (v *NullablePreRegistrationInlineHook) Set(val *PreRegistrationInlineHook) {
	v.value = val
	v.isSet = true
}

func (v NullablePreRegistrationInlineHook) IsSet() bool {
	return v.isSet
}

func (v *NullablePreRegistrationInlineHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreRegistrationInlineHook(val *PreRegistrationInlineHook) *NullablePreRegistrationInlineHook {
	return &NullablePreRegistrationInlineHook{value: val, isSet: true}
}

func (v NullablePreRegistrationInlineHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreRegistrationInlineHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

