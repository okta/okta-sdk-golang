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
)

// checks if the OptInStatusResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptInStatusResponseLinks{}

// OptInStatusResponseLinks Link relations available
type OptInStatusResponseLinks struct {
	// Link to the opt-in status resource
	OptInStatus          *HrefObject `json:"optInStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OptInStatusResponseLinks OptInStatusResponseLinks

// NewOptInStatusResponseLinks instantiates a new OptInStatusResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptInStatusResponseLinks() *OptInStatusResponseLinks {
	this := OptInStatusResponseLinks{}
	return &this
}

// NewOptInStatusResponseLinksWithDefaults instantiates a new OptInStatusResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptInStatusResponseLinksWithDefaults() *OptInStatusResponseLinks {
	this := OptInStatusResponseLinks{}
	return &this
}

// GetOptInStatus returns the OptInStatus field value if set, zero value otherwise.
func (o *OptInStatusResponseLinks) GetOptInStatus() HrefObject {
	if o == nil || IsNil(o.OptInStatus) {
		var ret HrefObject
		return ret
	}
	return *o.OptInStatus
}

// GetOptInStatusOk returns a tuple with the OptInStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptInStatusResponseLinks) GetOptInStatusOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.OptInStatus) {
		return nil, false
	}
	return o.OptInStatus, true
}

// HasOptInStatus returns a boolean if a field has been set.
func (o *OptInStatusResponseLinks) HasOptInStatus() bool {
	if o != nil && !IsNil(o.OptInStatus) {
		return true
	}

	return false
}

// SetOptInStatus gets a reference to the given HrefObject and assigns it to the OptInStatus field.
func (o *OptInStatusResponseLinks) SetOptInStatus(v HrefObject) {
	o.OptInStatus = &v
}

func (o OptInStatusResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptInStatusResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptInStatus) {
		toSerialize["optInStatus"] = o.OptInStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OptInStatusResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varOptInStatusResponseLinks := _OptInStatusResponseLinks{}

	err = json.Unmarshal(data, &varOptInStatusResponseLinks)

	if err != nil {
		return err
	}

	*o = OptInStatusResponseLinks(varOptInStatusResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "optInStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOptInStatusResponseLinks struct {
	value *OptInStatusResponseLinks
	isSet bool
}

func (v NullableOptInStatusResponseLinks) Get() *OptInStatusResponseLinks {
	return v.value
}

func (v *NullableOptInStatusResponseLinks) Set(val *OptInStatusResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOptInStatusResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOptInStatusResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptInStatusResponseLinks(val *OptInStatusResponseLinks) *NullableOptInStatusResponseLinks {
	return &NullableOptInStatusResponseLinks{value: val, isSet: true}
}

func (v NullableOptInStatusResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptInStatusResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
