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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// LogDebugContext For some kinds of events (for example, OLM provisioning, sign-in request, second factor SMS, and so on), the fields that are provided in other response objects aren't sufficient to adequately describe the operations that the event has performed. In such cases, the `debugContext` object provides a way to store additional information.  For example, an event where a second factor SMS token is sent to a user may have a `debugContext` that looks like the following: ``` {     \"debugData\": {         \"requestUri\": \"/api/v1/users/00u3gjksoiRGRAZHLSYV/factors/smsf8luacpZJAva10x45/verify\",         \"smsProvider\": \"TELESIGN\",         \"transactionId\": \"268632458E3C100F5F5F594C6DC689D4\"     } } ``` By inspecting the debugData field, you can find the URI that is used to trigger the second factor SMS (`/api/v1/users/00u3gjksoiRGRAZHLSYV/factors/smsf8luacpZJAva10x45/verify`), the SMS provider (`TELESIGN`), and the ID used by Telesign to identify this transaction (`268632458E3C100F5F5F594C6DC689D4`).  If for some reason the information that is needed to implement a feature isn't provided in other response objects, you should scan the `debugContext.debugData` field for potentially useful fields. > **Important:** The information contained in `debugContext.debugData` is intended to add context when troubleshooting customer platform issues. Both key names and values may change from release to release and aren't guaranteed to be stable. Therefore, they shouldn't be viewed as a data contract but as a debugging aid instead.
type LogDebugContext struct {
	// A dynamic field that contains miscellaneous information that is dependent on the event type.
	DebugData map[string]interface{} `json:"debugData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogDebugContext LogDebugContext

// NewLogDebugContext instantiates a new LogDebugContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogDebugContext() *LogDebugContext {
	this := LogDebugContext{}
	return &this
}

// NewLogDebugContextWithDefaults instantiates a new LogDebugContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogDebugContextWithDefaults() *LogDebugContext {
	this := LogDebugContext{}
	return &this
}

// GetDebugData returns the DebugData field value if set, zero value otherwise.
func (o *LogDebugContext) GetDebugData() map[string]interface{} {
	if o == nil || o.DebugData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DebugData
}

// GetDebugDataOk returns a tuple with the DebugData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDebugContext) GetDebugDataOk() (map[string]interface{}, bool) {
	if o == nil || o.DebugData == nil {
		return nil, false
	}
	return o.DebugData, true
}

// HasDebugData returns a boolean if a field has been set.
func (o *LogDebugContext) HasDebugData() bool {
	if o != nil && o.DebugData != nil {
		return true
	}

	return false
}

// SetDebugData gets a reference to the given map[string]interface{} and assigns it to the DebugData field.
func (o *LogDebugContext) SetDebugData(v map[string]interface{}) {
	o.DebugData = v
}

func (o LogDebugContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DebugData != nil {
		toSerialize["debugData"] = o.DebugData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogDebugContext) UnmarshalJSON(bytes []byte) (err error) {
	varLogDebugContext := _LogDebugContext{}

	err = json.Unmarshal(bytes, &varLogDebugContext)
	if err == nil {
		*o = LogDebugContext(varLogDebugContext)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "debugData")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogDebugContext struct {
	value *LogDebugContext
	isSet bool
}

func (v NullableLogDebugContext) Get() *LogDebugContext {
	return v.value
}

func (v *NullableLogDebugContext) Set(val *LogDebugContext) {
	v.value = val
	v.isSet = true
}

func (v NullableLogDebugContext) IsSet() bool {
	return v.isSet
}

func (v *NullableLogDebugContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogDebugContext(val *LogDebugContext) *NullableLogDebugContext {
	return &NullableLogDebugContext{value: val, isSet: true}
}

func (v NullableLogDebugContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogDebugContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

