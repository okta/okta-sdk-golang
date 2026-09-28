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

// checks if the ScriptInvokeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScriptInvokeRequest{}

// ScriptInvokeRequest Request to invoke a remote script on an AD agent.
type ScriptInvokeRequest struct {
	// Optional JSON payload forwarded to the script. Maximum size is 4 KB.
	Payload map[string]interface{} `json:"payload,omitempty"`
	// The type of script to invoke
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _ScriptInvokeRequest ScriptInvokeRequest

// NewScriptInvokeRequest instantiates a new ScriptInvokeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptInvokeRequest(type_ string) *ScriptInvokeRequest {
	this := ScriptInvokeRequest{}
	this.Type = type_
	return &this
}

// NewScriptInvokeRequestWithDefaults instantiates a new ScriptInvokeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptInvokeRequestWithDefaults() *ScriptInvokeRequest {
	this := ScriptInvokeRequest{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ScriptInvokeRequest) GetPayload() map[string]interface{} {
	if o == nil || IsNil(o.Payload) {
		var ret map[string]interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptInvokeRequest) GetPayloadOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Payload) {
		return map[string]interface{}{}, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ScriptInvokeRequest) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.
func (o *ScriptInvokeRequest) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetType returns the Type field value
func (o *ScriptInvokeRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScriptInvokeRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScriptInvokeRequest) SetType(v string) {
	o.Type = v
}

func (o ScriptInvokeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScriptInvokeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScriptInvokeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varScriptInvokeRequest := _ScriptInvokeRequest{}

	err = json.Unmarshal(data, &varScriptInvokeRequest)

	if err != nil {
		return err
	}

	*o = ScriptInvokeRequest(varScriptInvokeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "payload")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScriptInvokeRequest struct {
	value *ScriptInvokeRequest
	isSet bool
}

func (v NullableScriptInvokeRequest) Get() *ScriptInvokeRequest {
	return v.value
}

func (v *NullableScriptInvokeRequest) Set(val *ScriptInvokeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptInvokeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptInvokeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptInvokeRequest(val *ScriptInvokeRequest) *NullableScriptInvokeRequest {
	return &NullableScriptInvokeRequest{value: val, isSet: true}
}

func (v NullableScriptInvokeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptInvokeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
