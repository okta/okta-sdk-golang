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

// InlineHookResponse struct for InlineHookResponse
type InlineHookResponse struct {
	Commands []InlineHookResponseCommands `json:"commands,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookResponse InlineHookResponse

// NewInlineHookResponse instantiates a new InlineHookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookResponse() *InlineHookResponse {
	this := InlineHookResponse{}
	return &this
}

// NewInlineHookResponseWithDefaults instantiates a new InlineHookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookResponseWithDefaults() *InlineHookResponse {
	this := InlineHookResponse{}
	return &this
}

// GetCommands returns the Commands field value if set, zero value otherwise.
func (o *InlineHookResponse) GetCommands() []InlineHookResponseCommands {
	if o == nil || o.Commands == nil {
		var ret []InlineHookResponseCommands
		return ret
	}
	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookResponse) GetCommandsOk() ([]InlineHookResponseCommands, bool) {
	if o == nil || o.Commands == nil {
		return nil, false
	}
	return o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *InlineHookResponse) HasCommands() bool {
	if o != nil && o.Commands != nil {
		return true
	}

	return false
}

// SetCommands gets a reference to the given []InlineHookResponseCommands and assigns it to the Commands field.
func (o *InlineHookResponse) SetCommands(v []InlineHookResponseCommands) {
	o.Commands = v
}

func (o InlineHookResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Commands != nil {
		toSerialize["commands"] = o.Commands
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineHookResponse) UnmarshalJSON(bytes []byte) (err error) {
	varInlineHookResponse := _InlineHookResponse{}

	err = json.Unmarshal(bytes, &varInlineHookResponse)
	if err == nil {
		*o = InlineHookResponse(varInlineHookResponse)
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

type NullableInlineHookResponse struct {
	value *InlineHookResponse
	isSet bool
}

func (v NullableInlineHookResponse) Get() *InlineHookResponse {
	return v.value
}

func (v *NullableInlineHookResponse) Set(val *InlineHookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookResponse(val *InlineHookResponse) *NullableInlineHookResponse {
	return &NullableInlineHookResponse{value: val, isSet: true}
}

func (v NullableInlineHookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

