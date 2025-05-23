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

// LogStreamLinksSelfAndLifecycle Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the current status of an application using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification. This object is used for dynamic discovery of related resources and lifecycle operations.
type LogStreamLinksSelfAndLifecycle struct {
	Activate *LogStreamActivateLink `json:"activate,omitempty"`
	Deactivate *LogStreamDeactivateLink `json:"deactivate,omitempty"`
	Self LogStreamSelfLink `json:"self"`
	AdditionalProperties map[string]interface{}
}

type _LogStreamLinksSelfAndLifecycle LogStreamLinksSelfAndLifecycle

// NewLogStreamLinksSelfAndLifecycle instantiates a new LogStreamLinksSelfAndLifecycle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStreamLinksSelfAndLifecycle(self LogStreamSelfLink) *LogStreamLinksSelfAndLifecycle {
	this := LogStreamLinksSelfAndLifecycle{}
	this.Self = self
	return &this
}

// NewLogStreamLinksSelfAndLifecycleWithDefaults instantiates a new LogStreamLinksSelfAndLifecycle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamLinksSelfAndLifecycleWithDefaults() *LogStreamLinksSelfAndLifecycle {
	this := LogStreamLinksSelfAndLifecycle{}
	return &this
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *LogStreamLinksSelfAndLifecycle) GetActivate() LogStreamActivateLink {
	if o == nil || o.Activate == nil {
		var ret LogStreamActivateLink
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamLinksSelfAndLifecycle) GetActivateOk() (*LogStreamActivateLink, bool) {
	if o == nil || o.Activate == nil {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *LogStreamLinksSelfAndLifecycle) HasActivate() bool {
	if o != nil && o.Activate != nil {
		return true
	}

	return false
}

// SetActivate gets a reference to the given LogStreamActivateLink and assigns it to the Activate field.
func (o *LogStreamLinksSelfAndLifecycle) SetActivate(v LogStreamActivateLink) {
	o.Activate = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *LogStreamLinksSelfAndLifecycle) GetDeactivate() LogStreamDeactivateLink {
	if o == nil || o.Deactivate == nil {
		var ret LogStreamDeactivateLink
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamLinksSelfAndLifecycle) GetDeactivateOk() (*LogStreamDeactivateLink, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *LogStreamLinksSelfAndLifecycle) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given LogStreamDeactivateLink and assigns it to the Deactivate field.
func (o *LogStreamLinksSelfAndLifecycle) SetDeactivate(v LogStreamDeactivateLink) {
	o.Deactivate = &v
}

// GetSelf returns the Self field value
func (o *LogStreamLinksSelfAndLifecycle) GetSelf() LogStreamSelfLink {
	if o == nil {
		var ret LogStreamSelfLink
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *LogStreamLinksSelfAndLifecycle) GetSelfOk() (*LogStreamSelfLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *LogStreamLinksSelfAndLifecycle) SetSelf(v LogStreamSelfLink) {
	o.Self = v
}

func (o LogStreamLinksSelfAndLifecycle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Activate != nil {
		toSerialize["activate"] = o.Activate
	}
	if o.Deactivate != nil {
		toSerialize["deactivate"] = o.Deactivate
	}
	if true {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogStreamLinksSelfAndLifecycle) UnmarshalJSON(bytes []byte) (err error) {
	varLogStreamLinksSelfAndLifecycle := _LogStreamLinksSelfAndLifecycle{}

	err = json.Unmarshal(bytes, &varLogStreamLinksSelfAndLifecycle)
	if err == nil {
		*o = LogStreamLinksSelfAndLifecycle(varLogStreamLinksSelfAndLifecycle)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "activate")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogStreamLinksSelfAndLifecycle struct {
	value *LogStreamLinksSelfAndLifecycle
	isSet bool
}

func (v NullableLogStreamLinksSelfAndLifecycle) Get() *LogStreamLinksSelfAndLifecycle {
	return v.value
}

func (v *NullableLogStreamLinksSelfAndLifecycle) Set(val *LogStreamLinksSelfAndLifecycle) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStreamLinksSelfAndLifecycle) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStreamLinksSelfAndLifecycle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStreamLinksSelfAndLifecycle(val *LogStreamLinksSelfAndLifecycle) *NullableLogStreamLinksSelfAndLifecycle {
	return &NullableLogStreamLinksSelfAndLifecycle{value: val, isSet: true}
}

func (v NullableLogStreamLinksSelfAndLifecycle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStreamLinksSelfAndLifecycle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

