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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the OktaSupportCaseSelfAssigned type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OktaSupportCaseSelfAssigned{}

// OktaSupportCaseSelfAssigned Customer allows Okta Support access to self-assigned cases. Support cases are self-assigned when an Okta Support team member creates and assigns the case to themselves.
type OktaSupportCaseSelfAssigned struct {
	// Okta Support access approval status for self-assigned cases
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaSupportCaseSelfAssigned OktaSupportCaseSelfAssigned

// NewOktaSupportCaseSelfAssigned instantiates a new OktaSupportCaseSelfAssigned object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaSupportCaseSelfAssigned() *OktaSupportCaseSelfAssigned {
	this := OktaSupportCaseSelfAssigned{}
	return &this
}

// NewOktaSupportCaseSelfAssignedWithDefaults instantiates a new OktaSupportCaseSelfAssigned object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaSupportCaseSelfAssignedWithDefaults() *OktaSupportCaseSelfAssigned {
	this := OktaSupportCaseSelfAssigned{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OktaSupportCaseSelfAssigned) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSupportCaseSelfAssigned) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OktaSupportCaseSelfAssigned) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OktaSupportCaseSelfAssigned) SetStatus(v string) {
	o.Status = &v
}

func (o OktaSupportCaseSelfAssigned) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OktaSupportCaseSelfAssigned) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OktaSupportCaseSelfAssigned) UnmarshalJSON(data []byte) (err error) {
	varOktaSupportCaseSelfAssigned := _OktaSupportCaseSelfAssigned{}

	err = json.Unmarshal(data, &varOktaSupportCaseSelfAssigned)

	if err != nil {
		return err
	}

	*o = OktaSupportCaseSelfAssigned(varOktaSupportCaseSelfAssigned)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOktaSupportCaseSelfAssigned struct {
	value *OktaSupportCaseSelfAssigned
	isSet bool
}

func (v NullableOktaSupportCaseSelfAssigned) Get() *OktaSupportCaseSelfAssigned {
	return v.value
}

func (v *NullableOktaSupportCaseSelfAssigned) Set(val *OktaSupportCaseSelfAssigned) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaSupportCaseSelfAssigned) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaSupportCaseSelfAssigned) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaSupportCaseSelfAssigned(val *OktaSupportCaseSelfAssigned) *NullableOktaSupportCaseSelfAssigned {
	return &NullableOktaSupportCaseSelfAssigned{value: val, isSet: true}
}

func (v NullableOktaSupportCaseSelfAssigned) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaSupportCaseSelfAssigned) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
