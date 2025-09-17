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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"time"
)

// checks if the UserClassification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserClassification{}

// UserClassification struct for UserClassification
type UserClassification struct {
	// The timestamp when the user classification was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The type of user classification
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserClassification UserClassification

// NewUserClassification instantiates a new UserClassification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserClassification() *UserClassification {
	this := UserClassification{}
	return &this
}

// NewUserClassificationWithDefaults instantiates a new UserClassification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserClassificationWithDefaults() *UserClassification {
	this := UserClassification{}
	return &this
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *UserClassification) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserClassification) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *UserClassification) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *UserClassification) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserClassification) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserClassification) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserClassification) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserClassification) SetType(v string) {
	o.Type = &v
}

func (o UserClassification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserClassification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserClassification) UnmarshalJSON(data []byte) (err error) {
	varUserClassification := _UserClassification{}

	err = json.Unmarshal(data, &varUserClassification)

	if err != nil {
		return err
	}

	*o = UserClassification(varUserClassification)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserClassification struct {
	value *UserClassification
	isSet bool
}

func (v NullableUserClassification) Get() *UserClassification {
	return v.value
}

func (v *NullableUserClassification) Set(val *UserClassification) {
	v.value = val
	v.isSet = true
}

func (v NullableUserClassification) IsSet() bool {
	return v.isSet
}

func (v *NullableUserClassification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserClassification(val *UserClassification) *NullableUserClassification {
	return &NullableUserClassification{value: val, isSet: true}
}

func (v NullableUserClassification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserClassification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
