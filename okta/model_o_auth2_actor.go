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

// OAuth2Actor User that created the object
type OAuth2Actor struct {
	// User ID
	Id *string `json:"id,omitempty"`
	// Type of user
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2Actor OAuth2Actor

// NewOAuth2Actor instantiates a new OAuth2Actor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2Actor() *OAuth2Actor {
	this := OAuth2Actor{}
	return &this
}

// NewOAuth2ActorWithDefaults instantiates a new OAuth2Actor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ActorWithDefaults() *OAuth2Actor {
	this := OAuth2Actor{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OAuth2Actor) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Actor) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OAuth2Actor) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OAuth2Actor) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OAuth2Actor) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Actor) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OAuth2Actor) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OAuth2Actor) SetType(v string) {
	o.Type = &v
}

func (o OAuth2Actor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2Actor) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2Actor := _OAuth2Actor{}

	err = json.Unmarshal(bytes, &varOAuth2Actor)
	if err == nil {
		*o = OAuth2Actor(varOAuth2Actor)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOAuth2Actor struct {
	value *OAuth2Actor
	isSet bool
}

func (v NullableOAuth2Actor) Get() *OAuth2Actor {
	return v.value
}

func (v *NullableOAuth2Actor) Set(val *OAuth2Actor) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2Actor) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2Actor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2Actor(val *OAuth2Actor) *NullableOAuth2Actor {
	return &NullableOAuth2Actor{value: val, isSet: true}
}

func (v NullableOAuth2Actor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2Actor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

