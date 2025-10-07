/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the PasswordImportRequestDataContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordImportRequestDataContext{}

// PasswordImportRequestDataContext struct for PasswordImportRequestDataContext
type PasswordImportRequestDataContext struct {
	Request              *InlineHookRequestObject                    `json:"request,omitempty"`
	Credential           *PasswordImportRequestDataContextCredential `json:"credential,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordImportRequestDataContext PasswordImportRequestDataContext

// NewPasswordImportRequestDataContext instantiates a new PasswordImportRequestDataContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordImportRequestDataContext() *PasswordImportRequestDataContext {
	this := PasswordImportRequestDataContext{}
	return &this
}

// NewPasswordImportRequestDataContextWithDefaults instantiates a new PasswordImportRequestDataContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordImportRequestDataContextWithDefaults() *PasswordImportRequestDataContext {
	this := PasswordImportRequestDataContext{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *PasswordImportRequestDataContext) GetRequest() InlineHookRequestObject {
	if o == nil || IsNil(o.Request) {
		var ret InlineHookRequestObject
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequestDataContext) GetRequestOk() (*InlineHookRequestObject, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *PasswordImportRequestDataContext) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineHookRequestObject and assigns it to the Request field.
func (o *PasswordImportRequestDataContext) SetRequest(v InlineHookRequestObject) {
	o.Request = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *PasswordImportRequestDataContext) GetCredential() PasswordImportRequestDataContextCredential {
	if o == nil || IsNil(o.Credential) {
		var ret PasswordImportRequestDataContextCredential
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequestDataContext) GetCredentialOk() (*PasswordImportRequestDataContextCredential, bool) {
	if o == nil || IsNil(o.Credential) {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *PasswordImportRequestDataContext) HasCredential() bool {
	if o != nil && !IsNil(o.Credential) {
		return true
	}

	return false
}

// SetCredential gets a reference to the given PasswordImportRequestDataContextCredential and assigns it to the Credential field.
func (o *PasswordImportRequestDataContext) SetCredential(v PasswordImportRequestDataContextCredential) {
	o.Credential = &v
}

func (o PasswordImportRequestDataContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordImportRequestDataContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.Credential) {
		toSerialize["credential"] = o.Credential
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordImportRequestDataContext) UnmarshalJSON(data []byte) (err error) {
	varPasswordImportRequestDataContext := _PasswordImportRequestDataContext{}

	err = json.Unmarshal(data, &varPasswordImportRequestDataContext)

	if err != nil {
		return err
	}

	*o = PasswordImportRequestDataContext(varPasswordImportRequestDataContext)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "request")
		delete(additionalProperties, "credential")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordImportRequestDataContext struct {
	value *PasswordImportRequestDataContext
	isSet bool
}

func (v NullablePasswordImportRequestDataContext) Get() *PasswordImportRequestDataContext {
	return v.value
}

func (v *NullablePasswordImportRequestDataContext) Set(val *PasswordImportRequestDataContext) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordImportRequestDataContext) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordImportRequestDataContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordImportRequestDataContext(val *PasswordImportRequestDataContext) *NullablePasswordImportRequestDataContext {
	return &NullablePasswordImportRequestDataContext{value: val, isSet: true}
}

func (v NullablePasswordImportRequestDataContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordImportRequestDataContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
