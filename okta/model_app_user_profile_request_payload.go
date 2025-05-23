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

// AppUserProfileRequestPayload Updates the assigned user profile > **Note:** The Okta API currently doesn't support entity tags for conditional updates. As long as you're the only user updating the the user profile, Okta recommends you fetch the most recent profile with [Retrieve an Application User](/openapi/okta-management/management/tag/ApplicationUsers/#tag/ApplicationUsers/operation/getApplicationUser), apply your profile update, and then `POST` back the updated profile.
type AppUserProfileRequestPayload struct {
	// Specifies the default and custom profile properties for a user. Properties that are visible in the Admin Console for an app assignment can also be assigned through the API. Some properties are reference properties that are imported from the target app and can't be configured. See [profile](/openapi/okta-management/management/tag/User/#tag/User/operation/getUser!c=200&path=profile&t=response). 
	Profile map[string]interface{} `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppUserProfileRequestPayload AppUserProfileRequestPayload

// NewAppUserProfileRequestPayload instantiates a new AppUserProfileRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUserProfileRequestPayload() *AppUserProfileRequestPayload {
	this := AppUserProfileRequestPayload{}
	return &this
}

// NewAppUserProfileRequestPayloadWithDefaults instantiates a new AppUserProfileRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserProfileRequestPayloadWithDefaults() *AppUserProfileRequestPayload {
	this := AppUserProfileRequestPayload{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *AppUserProfileRequestPayload) GetProfile() map[string]interface{} {
	if o == nil || o.Profile == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserProfileRequestPayload) GetProfileOk() (map[string]interface{}, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *AppUserProfileRequestPayload) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]interface{} and assigns it to the Profile field.
func (o *AppUserProfileRequestPayload) SetProfile(v map[string]interface{}) {
	o.Profile = v
}

func (o AppUserProfileRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppUserProfileRequestPayload) UnmarshalJSON(bytes []byte) (err error) {
	varAppUserProfileRequestPayload := _AppUserProfileRequestPayload{}

	err = json.Unmarshal(bytes, &varAppUserProfileRequestPayload)
	if err == nil {
		*o = AppUserProfileRequestPayload(varAppUserProfileRequestPayload)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAppUserProfileRequestPayload struct {
	value *AppUserProfileRequestPayload
	isSet bool
}

func (v NullableAppUserProfileRequestPayload) Get() *AppUserProfileRequestPayload {
	return v.value
}

func (v *NullableAppUserProfileRequestPayload) Set(val *AppUserProfileRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserProfileRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserProfileRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserProfileRequestPayload(val *AppUserProfileRequestPayload) *NullableAppUserProfileRequestPayload {
	return &NullableAppUserProfileRequestPayload{value: val, isSet: true}
}

func (v NullableAppUserProfileRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserProfileRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

