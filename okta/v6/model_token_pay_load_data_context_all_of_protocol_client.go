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

// TokenPayLoadDataContextAllOfProtocolClient The client making the token request
type TokenPayLoadDataContextAllOfProtocolClient struct {
	// The unique identifier of the client
	Id *string `json:"id,omitempty"`
	// The name of the client
	Name *string `json:"name,omitempty"`
	// The type of client
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenPayLoadDataContextAllOfProtocolClient TokenPayLoadDataContextAllOfProtocolClient

// NewTokenPayLoadDataContextAllOfProtocolClient instantiates a new TokenPayLoadDataContextAllOfProtocolClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenPayLoadDataContextAllOfProtocolClient() *TokenPayLoadDataContextAllOfProtocolClient {
	this := TokenPayLoadDataContextAllOfProtocolClient{}
	return &this
}

// NewTokenPayLoadDataContextAllOfProtocolClientWithDefaults instantiates a new TokenPayLoadDataContextAllOfProtocolClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenPayLoadDataContextAllOfProtocolClientWithDefaults() *TokenPayLoadDataContextAllOfProtocolClient {
	this := TokenPayLoadDataContextAllOfProtocolClient{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TokenPayLoadDataContextAllOfProtocolClient) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfProtocolClient) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfProtocolClient) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TokenPayLoadDataContextAllOfProtocolClient) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TokenPayLoadDataContextAllOfProtocolClient) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfProtocolClient) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfProtocolClient) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TokenPayLoadDataContextAllOfProtocolClient) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TokenPayLoadDataContextAllOfProtocolClient) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfProtocolClient) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfProtocolClient) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TokenPayLoadDataContextAllOfProtocolClient) SetType(v string) {
	o.Type = &v
}

func (o TokenPayLoadDataContextAllOfProtocolClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TokenPayLoadDataContextAllOfProtocolClient) UnmarshalJSON(bytes []byte) (err error) {
	varTokenPayLoadDataContextAllOfProtocolClient := _TokenPayLoadDataContextAllOfProtocolClient{}

	err = json.Unmarshal(bytes, &varTokenPayLoadDataContextAllOfProtocolClient)
	if err == nil {
		*o = TokenPayLoadDataContextAllOfProtocolClient(varTokenPayLoadDataContextAllOfProtocolClient)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTokenPayLoadDataContextAllOfProtocolClient struct {
	value *TokenPayLoadDataContextAllOfProtocolClient
	isSet bool
}

func (v NullableTokenPayLoadDataContextAllOfProtocolClient) Get() *TokenPayLoadDataContextAllOfProtocolClient {
	return v.value
}

func (v *NullableTokenPayLoadDataContextAllOfProtocolClient) Set(val *TokenPayLoadDataContextAllOfProtocolClient) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenPayLoadDataContextAllOfProtocolClient) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenPayLoadDataContextAllOfProtocolClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenPayLoadDataContextAllOfProtocolClient(val *TokenPayLoadDataContextAllOfProtocolClient) *NullableTokenPayLoadDataContextAllOfProtocolClient {
	return &NullableTokenPayLoadDataContextAllOfProtocolClient{value: val, isSet: true}
}

func (v NullableTokenPayLoadDataContextAllOfProtocolClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenPayLoadDataContextAllOfProtocolClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

