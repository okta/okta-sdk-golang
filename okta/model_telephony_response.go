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

// TelephonyResponse struct for TelephonyResponse
type TelephonyResponse struct {
	// The `commands` object specifies whether Okta accepts the end user's sign-in credentials as valid or not. For the Telephony inline hook, you typically only return one `commands` object with one array element in it.
	Commands []TelephonyResponseCommandsInner `json:"commands,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelephonyResponse TelephonyResponse

// NewTelephonyResponse instantiates a new TelephonyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelephonyResponse() *TelephonyResponse {
	this := TelephonyResponse{}
	return &this
}

// NewTelephonyResponseWithDefaults instantiates a new TelephonyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelephonyResponseWithDefaults() *TelephonyResponse {
	this := TelephonyResponse{}
	return &this
}

// GetCommands returns the Commands field value if set, zero value otherwise.
func (o *TelephonyResponse) GetCommands() []TelephonyResponseCommandsInner {
	if o == nil || o.Commands == nil {
		var ret []TelephonyResponseCommandsInner
		return ret
	}
	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyResponse) GetCommandsOk() ([]TelephonyResponseCommandsInner, bool) {
	if o == nil || o.Commands == nil {
		return nil, false
	}
	return o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *TelephonyResponse) HasCommands() bool {
	if o != nil && o.Commands != nil {
		return true
	}

	return false
}

// SetCommands gets a reference to the given []TelephonyResponseCommandsInner and assigns it to the Commands field.
func (o *TelephonyResponse) SetCommands(v []TelephonyResponseCommandsInner) {
	o.Commands = v
}

func (o TelephonyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Commands != nil {
		toSerialize["commands"] = o.Commands
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelephonyResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTelephonyResponse := _TelephonyResponse{}

	err = json.Unmarshal(bytes, &varTelephonyResponse)
	if err == nil {
		*o = TelephonyResponse(varTelephonyResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "commands")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTelephonyResponse struct {
	value *TelephonyResponse
	isSet bool
}

func (v NullableTelephonyResponse) Get() *TelephonyResponse {
	return v.value
}

func (v *NullableTelephonyResponse) Set(val *TelephonyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTelephonyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTelephonyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelephonyResponse(val *TelephonyResponse) *NullableTelephonyResponse {
	return &NullableTelephonyResponse{value: val, isSet: true}
}

func (v NullableTelephonyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelephonyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

