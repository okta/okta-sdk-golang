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
	"fmt"
)

// checks if the DeviceSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceSession{}

// DeviceSession struct for DeviceSession
type DeviceSession struct {
	// Authenticator methods that are verified in this session, with the most recently verified authenticator method first. This array always has at least one item.
	Factors []DeviceSessionFactor `json:"factors"`
	// The session's external identifier. The same value is recorded in the System Log as `authenticationContext.deviceSessionId`. Use that event to correlate this session with its events.
	Id string `json:"id"`
	// The unique identifier of the user that this session belongs to
	UserId               string              `json:"userId"`
	Links                *DeviceSessionLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceSession DeviceSession

// NewDeviceSession instantiates a new DeviceSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceSession(factors []DeviceSessionFactor, id string, userId string) *DeviceSession {
	this := DeviceSession{}
	this.Factors = factors
	this.Id = id
	this.UserId = userId
	return &this
}

// NewDeviceSessionWithDefaults instantiates a new DeviceSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceSessionWithDefaults() *DeviceSession {
	this := DeviceSession{}
	return &this
}

// GetFactors returns the Factors field value
func (o *DeviceSession) GetFactors() []DeviceSessionFactor {
	if o == nil {
		var ret []DeviceSessionFactor
		return ret
	}

	return o.Factors
}

// GetFactorsOk returns a tuple with the Factors field value
// and a boolean to check if the value has been set.
func (o *DeviceSession) GetFactorsOk() ([]DeviceSessionFactor, bool) {
	if o == nil {
		return nil, false
	}
	return o.Factors, true
}

// SetFactors sets field value
func (o *DeviceSession) SetFactors(v []DeviceSessionFactor) {
	o.Factors = v
}

// GetId returns the Id field value
func (o *DeviceSession) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeviceSession) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeviceSession) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *DeviceSession) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *DeviceSession) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *DeviceSession) SetUserId(v string) {
	o.UserId = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DeviceSession) GetLinks() DeviceSessionLinks {
	if o == nil || IsNil(o.Links) {
		var ret DeviceSessionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceSession) GetLinksOk() (*DeviceSessionLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DeviceSession) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given DeviceSessionLinks and assigns it to the Links field.
func (o *DeviceSession) SetLinks(v DeviceSessionLinks) {
	o.Links = &v
}

func (o DeviceSession) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["factors"] = o.Factors
	toSerialize["id"] = o.Id
	toSerialize["userId"] = o.UserId
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceSession) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"factors",
		"id",
		"userId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDeviceSession := _DeviceSession{}

	err = json.Unmarshal(data, &varDeviceSession)

	if err != nil {
		return err
	}

	*o = DeviceSession(varDeviceSession)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "factors")
		delete(additionalProperties, "id")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceSession struct {
	value *DeviceSession
	isSet bool
}

func (v NullableDeviceSession) Get() *DeviceSession {
	return v.value
}

func (v *NullableDeviceSession) Set(val *DeviceSession) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceSession) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceSession(val *DeviceSession) *NullableDeviceSession {
	return &NullableDeviceSession{value: val, isSet: true}
}

func (v NullableDeviceSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
