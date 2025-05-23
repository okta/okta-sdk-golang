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

// KeyRequest struct for KeyRequest
type KeyRequest struct {
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeyRequest KeyRequest

// NewKeyRequest instantiates a new KeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyRequest() *KeyRequest {
	this := KeyRequest{}
	return &this
}

// NewKeyRequestWithDefaults instantiates a new KeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyRequestWithDefaults() *KeyRequest {
	this := KeyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KeyRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KeyRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KeyRequest) SetName(v string) {
	o.Name = &v
}

func (o KeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KeyRequest) UnmarshalJSON(bytes []byte) (err error) {
	varKeyRequest := _KeyRequest{}

	err = json.Unmarshal(bytes, &varKeyRequest)
	if err == nil {
		*o = KeyRequest(varKeyRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableKeyRequest struct {
	value *KeyRequest
	isSet bool
}

func (v NullableKeyRequest) Get() *KeyRequest {
	return v.value
}

func (v *NullableKeyRequest) Set(val *KeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyRequest(val *KeyRequest) *NullableKeyRequest {
	return &NullableKeyRequest{value: val, isSet: true}
}

func (v NullableKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

