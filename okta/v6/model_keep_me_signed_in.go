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

// KeepMeSignedIn <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Controls how often the post-authentication prompt is presented to users
type KeepMeSignedIn struct {
	// Whether the post-authentication [Keep Me Signed In (KMSI)](https://help.okta.com/oie/en-us/content/topics/security/stay-signed-in.htm) flow is allowed
	PostAuth *string `json:"postAuth,omitempty"`
	// A time duration specified as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations).
	PostAuthPromptFrequency *string `json:"postAuthPromptFrequency,omitempty" validate:"regexp=^P(?:$)(\\\\d+Y)?(\\\\d+M)?(\\\\d+W)?(\\\\d+D)?(T(?:\\\\d)(\\\\d+H)?(\\\\d+M)?(\\\\d+S)?)?$"`
	AdditionalProperties map[string]interface{}
}

type _KeepMeSignedIn KeepMeSignedIn

// NewKeepMeSignedIn instantiates a new KeepMeSignedIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeepMeSignedIn() *KeepMeSignedIn {
	this := KeepMeSignedIn{}
	return &this
}

// NewKeepMeSignedInWithDefaults instantiates a new KeepMeSignedIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeepMeSignedInWithDefaults() *KeepMeSignedIn {
	this := KeepMeSignedIn{}
	return &this
}

// GetPostAuth returns the PostAuth field value if set, zero value otherwise.
func (o *KeepMeSignedIn) GetPostAuth() string {
	if o == nil || o.PostAuth == nil {
		var ret string
		return ret
	}
	return *o.PostAuth
}

// GetPostAuthOk returns a tuple with the PostAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeepMeSignedIn) GetPostAuthOk() (*string, bool) {
	if o == nil || o.PostAuth == nil {
		return nil, false
	}
	return o.PostAuth, true
}

// HasPostAuth returns a boolean if a field has been set.
func (o *KeepMeSignedIn) HasPostAuth() bool {
	if o != nil && o.PostAuth != nil {
		return true
	}

	return false
}

// SetPostAuth gets a reference to the given string and assigns it to the PostAuth field.
func (o *KeepMeSignedIn) SetPostAuth(v string) {
	o.PostAuth = &v
}

// GetPostAuthPromptFrequency returns the PostAuthPromptFrequency field value if set, zero value otherwise.
func (o *KeepMeSignedIn) GetPostAuthPromptFrequency() string {
	if o == nil || o.PostAuthPromptFrequency == nil {
		var ret string
		return ret
	}
	return *o.PostAuthPromptFrequency
}

// GetPostAuthPromptFrequencyOk returns a tuple with the PostAuthPromptFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeepMeSignedIn) GetPostAuthPromptFrequencyOk() (*string, bool) {
	if o == nil || o.PostAuthPromptFrequency == nil {
		return nil, false
	}
	return o.PostAuthPromptFrequency, true
}

// HasPostAuthPromptFrequency returns a boolean if a field has been set.
func (o *KeepMeSignedIn) HasPostAuthPromptFrequency() bool {
	if o != nil && o.PostAuthPromptFrequency != nil {
		return true
	}

	return false
}

// SetPostAuthPromptFrequency gets a reference to the given string and assigns it to the PostAuthPromptFrequency field.
func (o *KeepMeSignedIn) SetPostAuthPromptFrequency(v string) {
	o.PostAuthPromptFrequency = &v
}

func (o KeepMeSignedIn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PostAuth != nil {
		toSerialize["postAuth"] = o.PostAuth
	}
	if o.PostAuthPromptFrequency != nil {
		toSerialize["postAuthPromptFrequency"] = o.PostAuthPromptFrequency
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KeepMeSignedIn) UnmarshalJSON(bytes []byte) (err error) {
	varKeepMeSignedIn := _KeepMeSignedIn{}

	err = json.Unmarshal(bytes, &varKeepMeSignedIn)
	if err == nil {
		*o = KeepMeSignedIn(varKeepMeSignedIn)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "postAuth")
		delete(additionalProperties, "postAuthPromptFrequency")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableKeepMeSignedIn struct {
	value *KeepMeSignedIn
	isSet bool
}

func (v NullableKeepMeSignedIn) Get() *KeepMeSignedIn {
	return v.value
}

func (v *NullableKeepMeSignedIn) Set(val *KeepMeSignedIn) {
	v.value = val
	v.isSet = true
}

func (v NullableKeepMeSignedIn) IsSet() bool {
	return v.isSet
}

func (v *NullableKeepMeSignedIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeepMeSignedIn(val *KeepMeSignedIn) *NullableKeepMeSignedIn {
	return &NullableKeepMeSignedIn{value: val, isSet: true}
}

func (v NullableKeepMeSignedIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeepMeSignedIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

