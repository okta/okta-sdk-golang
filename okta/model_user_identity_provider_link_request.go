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

// UserIdentityProviderLinkRequest struct for UserIdentityProviderLinkRequest
type UserIdentityProviderLinkRequest struct {
	ExternalId *string `json:"externalId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserIdentityProviderLinkRequest UserIdentityProviderLinkRequest

// NewUserIdentityProviderLinkRequest instantiates a new UserIdentityProviderLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentityProviderLinkRequest() *UserIdentityProviderLinkRequest {
	this := UserIdentityProviderLinkRequest{}
	return &this
}

// NewUserIdentityProviderLinkRequestWithDefaults instantiates a new UserIdentityProviderLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentityProviderLinkRequestWithDefaults() *UserIdentityProviderLinkRequest {
	this := UserIdentityProviderLinkRequest{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *UserIdentityProviderLinkRequest) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentityProviderLinkRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *UserIdentityProviderLinkRequest) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *UserIdentityProviderLinkRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

func (o UserIdentityProviderLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserIdentityProviderLinkRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUserIdentityProviderLinkRequest := _UserIdentityProviderLinkRequest{}

	err = json.Unmarshal(bytes, &varUserIdentityProviderLinkRequest)
	if err == nil {
		*o = UserIdentityProviderLinkRequest(varUserIdentityProviderLinkRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "externalId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserIdentityProviderLinkRequest struct {
	value *UserIdentityProviderLinkRequest
	isSet bool
}

func (v NullableUserIdentityProviderLinkRequest) Get() *UserIdentityProviderLinkRequest {
	return v.value
}

func (v *NullableUserIdentityProviderLinkRequest) Set(val *UserIdentityProviderLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentityProviderLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentityProviderLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentityProviderLinkRequest(val *UserIdentityProviderLinkRequest) *NullableUserIdentityProviderLinkRequest {
	return &NullableUserIdentityProviderLinkRequest{value: val, isSet: true}
}

func (v NullableUserIdentityProviderLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentityProviderLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

