/*
Okta Admin Management API

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

// checks if the OktaVerifyPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OktaVerifyPayload{}

// OktaVerifyPayload Payload data for the OktaVerify provider
type OktaVerifyPayload struct {
	// The version of the Okta Verify client
	ClientVersion *string `json:"clientVersion,omitempty"`
	// Timestamp when the client instance was last updated
	LastUpdated          *time.Time `json:"lastUpdated,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaVerifyPayload OktaVerifyPayload

// NewOktaVerifyPayload instantiates a new OktaVerifyPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaVerifyPayload() *OktaVerifyPayload {
	this := OktaVerifyPayload{}
	return &this
}

// NewOktaVerifyPayloadWithDefaults instantiates a new OktaVerifyPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaVerifyPayloadWithDefaults() *OktaVerifyPayload {
	this := OktaVerifyPayload{}
	return &this
}

// GetClientVersion returns the ClientVersion field value if set, zero value otherwise.
func (o *OktaVerifyPayload) GetClientVersion() string {
	if o == nil || IsNil(o.ClientVersion) {
		var ret string
		return ret
	}
	return *o.ClientVersion
}

// GetClientVersionOk returns a tuple with the ClientVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaVerifyPayload) GetClientVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ClientVersion) {
		return nil, false
	}
	return o.ClientVersion, true
}

// HasClientVersion returns a boolean if a field has been set.
func (o *OktaVerifyPayload) HasClientVersion() bool {
	if o != nil && !IsNil(o.ClientVersion) {
		return true
	}

	return false
}

// SetClientVersion gets a reference to the given string and assigns it to the ClientVersion field.
func (o *OktaVerifyPayload) SetClientVersion(v string) {
	o.ClientVersion = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *OktaVerifyPayload) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaVerifyPayload) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *OktaVerifyPayload) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *OktaVerifyPayload) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o OktaVerifyPayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OktaVerifyPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientVersion) {
		toSerialize["clientVersion"] = o.ClientVersion
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OktaVerifyPayload) UnmarshalJSON(data []byte) (err error) {
	varOktaVerifyPayload := _OktaVerifyPayload{}

	err = json.Unmarshal(data, &varOktaVerifyPayload)

	if err != nil {
		return err
	}

	*o = OktaVerifyPayload(varOktaVerifyPayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "clientVersion")
		delete(additionalProperties, "lastUpdated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOktaVerifyPayload struct {
	value *OktaVerifyPayload
	isSet bool
}

func (v NullableOktaVerifyPayload) Get() *OktaVerifyPayload {
	return v.value
}

func (v *NullableOktaVerifyPayload) Set(val *OktaVerifyPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaVerifyPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaVerifyPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaVerifyPayload(val *OktaVerifyPayload) *NullableOktaVerifyPayload {
	return &NullableOktaVerifyPayload{value: val, isSet: true}
}

func (v NullableOktaVerifyPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaVerifyPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
