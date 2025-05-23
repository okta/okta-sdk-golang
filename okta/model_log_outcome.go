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

// LogOutcome struct for LogOutcome
type LogOutcome struct {
	Reason *string `json:"reason,omitempty"`
	Result *string `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogOutcome LogOutcome

// NewLogOutcome instantiates a new LogOutcome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogOutcome() *LogOutcome {
	this := LogOutcome{}
	return &this
}

// NewLogOutcomeWithDefaults instantiates a new LogOutcome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogOutcomeWithDefaults() *LogOutcome {
	this := LogOutcome{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *LogOutcome) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogOutcome) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *LogOutcome) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *LogOutcome) SetReason(v string) {
	o.Reason = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *LogOutcome) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogOutcome) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *LogOutcome) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *LogOutcome) SetResult(v string) {
	o.Result = &v
}

func (o LogOutcome) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogOutcome) UnmarshalJSON(bytes []byte) (err error) {
	varLogOutcome := _LogOutcome{}

	err = json.Unmarshal(bytes, &varLogOutcome)
	if err == nil {
		*o = LogOutcome(varLogOutcome)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "reason")
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogOutcome struct {
	value *LogOutcome
	isSet bool
}

func (v NullableLogOutcome) Get() *LogOutcome {
	return v.value
}

func (v *NullableLogOutcome) Set(val *LogOutcome) {
	v.value = val
	v.isSet = true
}

func (v NullableLogOutcome) IsSet() bool {
	return v.isSet
}

func (v *NullableLogOutcome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogOutcome(val *LogOutcome) *NullableLogOutcome {
	return &NullableLogOutcome{value: val, isSet: true}
}

func (v NullableLogOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogOutcome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

