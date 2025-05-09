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
	"time"
)

// UserFactorVerifyResponse struct for UserFactorVerifyResponse
type UserFactorVerifyResponse struct {
	// Timestamp when the verification expires
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Optional display message for Factor verification
	FactorMessage NullableString `json:"factorMessage,omitempty"`
	// Result of a Factor verification
	FactorResult *string `json:"factorResult,omitempty"`
	Embedded map[string]map[string]interface{} `json:"_embedded,omitempty"`
	Links *UserFactorLinks `json:"_links,omitempty"`
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
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorVerifyResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *UserFactorVerifyResponse) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
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
	if o == nil || o.FactorMessage.Get() == nil {
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
	if o == nil || o.FactorResult == nil {
		var ret string
		return ret
	}
	return *o.FactorResult
}

// GetFactorResultOk returns a tuple with the FactorResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorVerifyResponse) GetFactorResultOk() (*string, bool) {
	if o == nil || o.FactorResult == nil {
		return nil, false
	}
	return o.FactorResult, true
}

// HasFactorResult returns a boolean if a field has been set.
func (o *UserFactorVerifyResponse) HasFactorResult() bool {
	if o != nil && o.FactorResult != nil {
		return true
	}

	return false
}

// SetFactorResult gets a reference to the given string and assigns it to the FactorResult field.
func (o *UserFactorVerifyResponse) SetFactorResult(v string) {
	o.FactorResult = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *UserFactorVerifyResponse) GetEmbedded() map[string]map[string]interface{} {
	if o == nil || o.Embedded == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorVerifyResponse) GetEmbeddedOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *UserFactorVerifyResponse) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
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
	if o == nil || o.Links == nil {
		var ret UserFactorLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorVerifyResponse) GetLinksOk() (*UserFactorLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserFactorVerifyResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given UserFactorLinks and assigns it to the Links field.
func (o *UserFactorVerifyResponse) SetLinks(v UserFactorLinks) {
	o.Links = &v
}

func (o UserFactorVerifyResponse) MarshalJSON() ([]byte, error) {
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

func (o *UserFactorVerifyResponse) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorVerifyResponse := _UserFactorVerifyResponse{}

	err = json.Unmarshal(bytes, &varUserFactorVerifyResponse)
	if err == nil {
		*o = UserFactorVerifyResponse(varUserFactorVerifyResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "factorMessage")
		delete(additionalProperties, "factorResult")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

