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
)

// checks if the StartOrgFailover200ResponseResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartOrgFailover200ResponseResultsInner{}

// StartOrgFailover200ResponseResultsInner struct for StartOrgFailover200ResponseResultsInner
type StartOrgFailover200ResponseResultsInner struct {
	// Domain for your org
	Domain *string `json:"domain,omitempty"`
	// Domain failover message
	Message              *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StartOrgFailover200ResponseResultsInner StartOrgFailover200ResponseResultsInner

// NewStartOrgFailover200ResponseResultsInner instantiates a new StartOrgFailover200ResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartOrgFailover200ResponseResultsInner() *StartOrgFailover200ResponseResultsInner {
	this := StartOrgFailover200ResponseResultsInner{}
	return &this
}

// NewStartOrgFailover200ResponseResultsInnerWithDefaults instantiates a new StartOrgFailover200ResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartOrgFailover200ResponseResultsInnerWithDefaults() *StartOrgFailover200ResponseResultsInner {
	this := StartOrgFailover200ResponseResultsInner{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *StartOrgFailover200ResponseResultsInner) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartOrgFailover200ResponseResultsInner) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *StartOrgFailover200ResponseResultsInner) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *StartOrgFailover200ResponseResultsInner) SetDomain(v string) {
	o.Domain = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *StartOrgFailover200ResponseResultsInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartOrgFailover200ResponseResultsInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *StartOrgFailover200ResponseResultsInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *StartOrgFailover200ResponseResultsInner) SetMessage(v string) {
	o.Message = &v
}

func (o StartOrgFailover200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartOrgFailover200ResponseResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StartOrgFailover200ResponseResultsInner) UnmarshalJSON(data []byte) (err error) {
	varStartOrgFailover200ResponseResultsInner := _StartOrgFailover200ResponseResultsInner{}

	err = json.Unmarshal(data, &varStartOrgFailover200ResponseResultsInner)

	if err != nil {
		return err
	}

	*o = StartOrgFailover200ResponseResultsInner(varStartOrgFailover200ResponseResultsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "domain")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStartOrgFailover200ResponseResultsInner struct {
	value *StartOrgFailover200ResponseResultsInner
	isSet bool
}

func (v NullableStartOrgFailover200ResponseResultsInner) Get() *StartOrgFailover200ResponseResultsInner {
	return v.value
}

func (v *NullableStartOrgFailover200ResponseResultsInner) Set(val *StartOrgFailover200ResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableStartOrgFailover200ResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableStartOrgFailover200ResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartOrgFailover200ResponseResultsInner(val *StartOrgFailover200ResponseResultsInner) *NullableStartOrgFailover200ResponseResultsInner {
	return &NullableStartOrgFailover200ResponseResultsInner{value: val, isSet: true}
}

func (v NullableStartOrgFailover200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartOrgFailover200ResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
