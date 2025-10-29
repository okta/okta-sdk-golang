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
	"time"
)

// checks if the UserFactorVerifyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFactorVerifyResponse{}

// UserFactorVerifyResponse struct for UserFactorVerifyResponse
type UserFactorVerifyResponse struct {
	// Timestamp when the verification expires
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Optional display message for factor verification
	FactorMessage NullableString `json:"factorMessage,omitempty"`
	// Result of a factor verification
	FactorResult         *string                           `json:"factorResult,omitempty"`
	Profile              map[string]map[string]interface{} `json:"profile,omitempty"`
	Embedded             map[string]map[string]interface{} `json:"_embedded,omitempty"`
	Links                *UserFactorLinks                  `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorVerifyResponse UserFactorVerifyResponse

// NewUserFactorVerifyResponse instantiates a new UserFactorVerifyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorVerifyResponse() *UserFactorVerifyResponse {
	this := UserFactorVerifyResponse{}
	return &this
}

// NewUserFactorVerifyResponseWithDefaults instantiates a new UserFactorVerifyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorVerifyResponseWithDefaults() *UserFactorVerifyResponse {
	this := UserFactorVerifyResponse{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *UserFactorVerifyResponse) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorVerifyResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *UserFactorVerifyResponse) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *UserFactorVerifyResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetFactorMessage returns the FactorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserFactorVerifyResponse) GetFactorMessage() string {
	if o == nil || IsNil(o.FactorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.FactorMessage.Get()
}

// GetFactorMessageOk returns a tuple with the FactorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserFactorVerifyResponse) GetFactorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FactorMessage.Get(), o.FactorMessage.IsSet()
}

// HasFactorMessage returns a boolean if a field has been set.
func (o *UserFactorVerifyResponse) HasFactorMessage() bool {
	if o != nil && o.FactorMessage.IsSet() {
		return true
	}

	return false
}

// SetFactorMessage gets a reference to the given NullableString and assigns it to the FactorMessage field.
func (o *UserFactorVerifyResponse) SetFactorMessage(v string) {
	o.FactorMessage.Set(&v)
}

// SetFactorMessageNil sets the value for FactorMessage to be an explicit nil
func (o *UserFactorVerifyResponse) SetFactorMessageNil() {
	o.FactorMessage.Set(nil)
}

// UnsetFactorMessage ensures that no value is present for FactorMessage, not even an explicit nil
func (o *UserFactorVerifyResponse) UnsetFactorMessage() {
	o.FactorMessage.Unset()
}

// GetFactorResult returns the FactorResult field value if set, zero value otherwise.
func (o *UserFactorVerifyResponse) GetFactorResult() string {
	if o == nil || IsNil(o.FactorResult) {
		var ret string
		return ret
	}
	return *o.FactorResult
}

// GetFactorResultOk returns a tuple with the FactorResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorVerifyResponse) GetFactorResultOk() (*string, bool) {
	if o == nil || IsNil(o.FactorResult) {
		return nil, false
	}
	return o.FactorResult, true
}

// HasFactorResult returns a boolean if a field has been set.
func (o *UserFactorVerifyResponse) HasFactorResult() bool {
	if o != nil && !IsNil(o.FactorResult) {
		return true
	}

	return false
}

// SetFactorResult gets a reference to the given string and assigns it to the FactorResult field.
func (o *UserFactorVerifyResponse) SetFactorResult(v string) {
	o.FactorResult = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactorVerifyResponse) GetProfile() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Profile) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorVerifyResponse) GetProfileOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Profile) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactorVerifyResponse) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]map[string]interface{} and assigns it to the Profile field.
func (o *UserFactorVerifyResponse) SetProfile(v map[string]map[string]interface{}) {
	o.Profile = v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *UserFactorVerifyResponse) GetEmbedded() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Embedded) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorVerifyResponse) GetEmbeddedOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Embedded) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *UserFactorVerifyResponse) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]map[string]interface{} and assigns it to the Embedded field.
func (o *UserFactorVerifyResponse) SetEmbedded(v map[string]map[string]interface{}) {
	o.Embedded = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserFactorVerifyResponse) GetLinks() UserFactorLinks {
	if o == nil || IsNil(o.Links) {
		var ret UserFactorLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorVerifyResponse) GetLinksOk() (*UserFactorLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserFactorVerifyResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given UserFactorLinks and assigns it to the Links field.
func (o *UserFactorVerifyResponse) SetLinks(v UserFactorLinks) {
	o.Links = &v
}

func (o UserFactorVerifyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFactorVerifyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if o.FactorMessage.IsSet() {
		toSerialize["factorMessage"] = o.FactorMessage.Get()
	}
	if !IsNil(o.FactorResult) {
		toSerialize["factorResult"] = o.FactorResult
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserFactorVerifyResponse) UnmarshalJSON(data []byte) (err error) {
	varUserFactorVerifyResponse := _UserFactorVerifyResponse{}

	err = json.Unmarshal(data, &varUserFactorVerifyResponse)

	if err != nil {
		return err
	}

	*o = UserFactorVerifyResponse(varUserFactorVerifyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "factorMessage")
		delete(additionalProperties, "factorResult")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserFactorVerifyResponse struct {
	value *UserFactorVerifyResponse
	isSet bool
}

func (v NullableUserFactorVerifyResponse) Get() *UserFactorVerifyResponse {
	return v.value
}

func (v *NullableUserFactorVerifyResponse) Set(val *UserFactorVerifyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorVerifyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorVerifyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorVerifyResponse(val *UserFactorVerifyResponse) *NullableUserFactorVerifyResponse {
	return &NullableUserFactorVerifyResponse{value: val, isSet: true}
}

func (v NullableUserFactorVerifyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorVerifyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
