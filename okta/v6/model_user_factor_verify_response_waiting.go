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
	"time"
)

// UserFactorVerifyResponseWaiting struct for UserFactorVerifyResponseWaiting
type UserFactorVerifyResponseWaiting struct {
	// Timestamp when the verification expires
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Optional display message for factor verification
	FactorMessage NullableString `json:"factorMessage,omitempty"`
	// Result of a factor verification
	FactorResult *string `json:"factorResult,omitempty"`
	Profile map[string]map[string]interface{} `json:"profile,omitempty"`
	Embedded *UserFactorVerifyResponseWaitingEmbedded `json:"_embedded,omitempty"`
	Links *UserFactorLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorVerifyResponseWaiting UserFactorVerifyResponseWaiting

// NewUserFactorVerifyResponseWaiting instantiates a new UserFactorVerifyResponseWaiting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorVerifyResponseWaiting() *UserFactorVerifyResponseWaiting {
	this := UserFactorVerifyResponseWaiting{}
	return &this
}

// NewUserFactorVerifyResponseWaitingWithDefaults instantiates a new UserFactorVerifyResponseWaiting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorVerifyResponseWaitingWithDefaults() *UserFactorVerifyResponseWaiting {
	this := UserFactorVerifyResponseWaiting{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *UserFactorVerifyResponseWaiting) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorVerifyResponseWaiting) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *UserFactorVerifyResponseWaiting) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *UserFactorVerifyResponseWaiting) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetFactorMessage returns the FactorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserFactorVerifyResponseWaiting) GetFactorMessage() string {
	if o == nil || o.FactorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.FactorMessage.Get()
}

// GetFactorMessageOk returns a tuple with the FactorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserFactorVerifyResponseWaiting) GetFactorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FactorMessage.Get(), o.FactorMessage.IsSet()
}

// HasFactorMessage returns a boolean if a field has been set.
func (o *UserFactorVerifyResponseWaiting) HasFactorMessage() bool {
	if o != nil && o.FactorMessage.IsSet() {
		return true
	}

	return false
}

// SetFactorMessage gets a reference to the given NullableString and assigns it to the FactorMessage field.
func (o *UserFactorVerifyResponseWaiting) SetFactorMessage(v string) {
	o.FactorMessage.Set(&v)
}
// SetFactorMessageNil sets the value for FactorMessage to be an explicit nil
func (o *UserFactorVerifyResponseWaiting) SetFactorMessageNil() {
	o.FactorMessage.Set(nil)
}

// UnsetFactorMessage ensures that no value is present for FactorMessage, not even an explicit nil
func (o *UserFactorVerifyResponseWaiting) UnsetFactorMessage() {
	o.FactorMessage.Unset()
}

// GetFactorResult returns the FactorResult field value if set, zero value otherwise.
func (o *UserFactorVerifyResponseWaiting) GetFactorResult() string {
	if o == nil || o.FactorResult == nil {
		var ret string
		return ret
	}
	return *o.FactorResult
}

// GetFactorResultOk returns a tuple with the FactorResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorVerifyResponseWaiting) GetFactorResultOk() (*string, bool) {
	if o == nil || o.FactorResult == nil {
		return nil, false
	}
	return o.FactorResult, true
}

// HasFactorResult returns a boolean if a field has been set.
func (o *UserFactorVerifyResponseWaiting) HasFactorResult() bool {
	if o != nil && o.FactorResult != nil {
		return true
	}

	return false
}

// SetFactorResult gets a reference to the given string and assigns it to the FactorResult field.
func (o *UserFactorVerifyResponseWaiting) SetFactorResult(v string) {
	o.FactorResult = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactorVerifyResponseWaiting) GetProfile() map[string]map[string]interface{} {
	if o == nil || o.Profile == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorVerifyResponseWaiting) GetProfileOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactorVerifyResponseWaiting) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]map[string]interface{} and assigns it to the Profile field.
func (o *UserFactorVerifyResponseWaiting) SetProfile(v map[string]map[string]interface{}) {
	o.Profile = v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *UserFactorVerifyResponseWaiting) GetEmbedded() UserFactorVerifyResponseWaitingEmbedded {
	if o == nil || o.Embedded == nil {
		var ret UserFactorVerifyResponseWaitingEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorVerifyResponseWaiting) GetEmbeddedOk() (*UserFactorVerifyResponseWaitingEmbedded, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *UserFactorVerifyResponseWaiting) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given UserFactorVerifyResponseWaitingEmbedded and assigns it to the Embedded field.
func (o *UserFactorVerifyResponseWaiting) SetEmbedded(v UserFactorVerifyResponseWaitingEmbedded) {
	o.Embedded = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserFactorVerifyResponseWaiting) GetLinks() UserFactorLinks {
	if o == nil || o.Links == nil {
		var ret UserFactorLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorVerifyResponseWaiting) GetLinksOk() (*UserFactorLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserFactorVerifyResponseWaiting) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given UserFactorLinks and assigns it to the Links field.
func (o *UserFactorVerifyResponseWaiting) SetLinks(v UserFactorLinks) {
	o.Links = &v
}

func (o UserFactorVerifyResponseWaiting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if o.FactorMessage.IsSet() {
		toSerialize["factorMessage"] = o.FactorMessage.Get()
	}
	if o.FactorResult != nil {
		toSerialize["factorResult"] = o.FactorResult
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.Embedded != nil {
		toSerialize["_embedded"] = o.Embedded
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorVerifyResponseWaiting) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorVerifyResponseWaiting := _UserFactorVerifyResponseWaiting{}

	err = json.Unmarshal(bytes, &varUserFactorVerifyResponseWaiting)
	if err == nil {
		*o = UserFactorVerifyResponseWaiting(varUserFactorVerifyResponseWaiting)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "factorMessage")
		delete(additionalProperties, "factorResult")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorVerifyResponseWaiting struct {
	value *UserFactorVerifyResponseWaiting
	isSet bool
}

func (v NullableUserFactorVerifyResponseWaiting) Get() *UserFactorVerifyResponseWaiting {
	return v.value
}

func (v *NullableUserFactorVerifyResponseWaiting) Set(val *UserFactorVerifyResponseWaiting) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorVerifyResponseWaiting) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorVerifyResponseWaiting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorVerifyResponseWaiting(val *UserFactorVerifyResponseWaiting) *NullableUserFactorVerifyResponseWaiting {
	return &NullableUserFactorVerifyResponseWaiting{value: val, isSet: true}
}

func (v NullableUserFactorVerifyResponseWaiting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorVerifyResponseWaiting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

