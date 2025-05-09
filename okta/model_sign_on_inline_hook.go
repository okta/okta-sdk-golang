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

// SignOnInlineHook struct for SignOnInlineHook
type SignOnInlineHook struct {
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SignOnInlineHook SignOnInlineHook

// NewSignOnInlineHook instantiates a new SignOnInlineHook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnInlineHook() *SignOnInlineHook {
	this := SignOnInlineHook{}
	return &this
}

// NewSignOnInlineHookWithDefaults instantiates a new SignOnInlineHook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnInlineHookWithDefaults() *SignOnInlineHook {
	this := SignOnInlineHook{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SignOnInlineHook) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnInlineHook) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SignOnInlineHook) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SignOnInlineHook) SetId(v string) {
	o.Id = &v
}

func (o SignOnInlineHook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SignOnInlineHook) UnmarshalJSON(bytes []byte) (err error) {
	varSignOnInlineHook := _SignOnInlineHook{}

	err = json.Unmarshal(bytes, &varSignOnInlineHook)
	if err == nil {
		*o = SignOnInlineHook(varSignOnInlineHook)
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

type NullableSignOnInlineHook struct {
	value *SignOnInlineHook
	isSet bool
}

func (v NullableSignOnInlineHook) Get() *SignOnInlineHook {
	return v.value
}

func (v *NullableSignOnInlineHook) Set(val *SignOnInlineHook) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnInlineHook) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnInlineHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnInlineHook(val *SignOnInlineHook) *NullableSignOnInlineHook {
	return &NullableSignOnInlineHook{value: val, isSet: true}
}

func (v NullableSignOnInlineHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnInlineHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

