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

// EmailServerListResponse struct for EmailServerListResponse
type EmailServerListResponse struct {
	EmailServers []EmailServerResponse `json:"email-servers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailServerListResponse EmailServerListResponse

// NewEmailServerListResponse instantiates a new EmailServerListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailServerListResponse() *EmailServerListResponse {
	this := EmailServerListResponse{}
	return &this
}

// NewEmailServerListResponseWithDefaults instantiates a new EmailServerListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailServerListResponseWithDefaults() *EmailServerListResponse {
	this := EmailServerListResponse{}
	return &this
}

// GetEmailServers returns the EmailServers field value if set, zero value otherwise.
func (o *EmailServerListResponse) GetEmailServers() []EmailServerResponse {
	if o == nil || o.EmailServers == nil {
		var ret []EmailServerResponse
		return ret
	}
	return o.EmailServers
}

// GetEmailServersOk returns a tuple with the EmailServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerListResponse) GetEmailServersOk() ([]EmailServerResponse, bool) {
	if o == nil || o.EmailServers == nil {
		return nil, false
	}
	return o.EmailServers, true
}

// HasEmailServers returns a boolean if a field has been set.
func (o *EmailServerListResponse) HasEmailServers() bool {
	if o != nil && o.EmailServers != nil {
		return true
	}

	return false
}

// SetEmailServers gets a reference to the given []EmailServerResponse and assigns it to the EmailServers field.
func (o *EmailServerListResponse) SetEmailServers(v []EmailServerResponse) {
	o.EmailServers = v
}

func (o EmailServerListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmailServers != nil {
		toSerialize["email-servers"] = o.EmailServers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailServerListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varEmailServerListResponse := _EmailServerListResponse{}

	err = json.Unmarshal(bytes, &varEmailServerListResponse)
	if err == nil {
		*o = EmailServerListResponse(varEmailServerListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "email-servers")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailServerListResponse struct {
	value *EmailServerListResponse
	isSet bool
}

func (v NullableEmailServerListResponse) Get() *EmailServerListResponse {
	return v.value
}

func (v *NullableEmailServerListResponse) Set(val *EmailServerListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailServerListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailServerListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailServerListResponse(val *EmailServerListResponse) *NullableEmailServerListResponse {
	return &NullableEmailServerListResponse{value: val, isSet: true}
}

func (v NullableEmailServerListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailServerListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

