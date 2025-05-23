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

// CreateSessionRequest struct for CreateSessionRequest
type CreateSessionRequest struct {
	// The session token obtained during authentication
	SessionToken *string `json:"sessionToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateSessionRequest CreateSessionRequest

// NewCreateSessionRequest instantiates a new CreateSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSessionRequest() *CreateSessionRequest {
	this := CreateSessionRequest{}
	return &this
}

// NewCreateSessionRequestWithDefaults instantiates a new CreateSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSessionRequestWithDefaults() *CreateSessionRequest {
	this := CreateSessionRequest{}
	return &this
}

// GetSessionToken returns the SessionToken field value if set, zero value otherwise.
func (o *CreateSessionRequest) GetSessionToken() string {
	if o == nil || o.SessionToken == nil {
		var ret string
		return ret
	}
	return *o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSessionRequest) GetSessionTokenOk() (*string, bool) {
	if o == nil || o.SessionToken == nil {
		return nil, false
	}
	return o.SessionToken, true
}

// HasSessionToken returns a boolean if a field has been set.
func (o *CreateSessionRequest) HasSessionToken() bool {
	if o != nil && o.SessionToken != nil {
		return true
	}

	return false
}

// SetSessionToken gets a reference to the given string and assigns it to the SessionToken field.
func (o *CreateSessionRequest) SetSessionToken(v string) {
	o.SessionToken = &v
}

func (o CreateSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SessionToken != nil {
		toSerialize["sessionToken"] = o.SessionToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateSessionRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateSessionRequest := _CreateSessionRequest{}

	err = json.Unmarshal(bytes, &varCreateSessionRequest)
	if err == nil {
		*o = CreateSessionRequest(varCreateSessionRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "sessionToken")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCreateSessionRequest struct {
	value *CreateSessionRequest
	isSet bool
}

func (v NullableCreateSessionRequest) Get() *CreateSessionRequest {
	return v.value
}

func (v *NullableCreateSessionRequest) Set(val *CreateSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSessionRequest(val *CreateSessionRequest) *NullableCreateSessionRequest {
	return &NullableCreateSessionRequest{value: val, isSet: true}
}

func (v NullableCreateSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

