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

// LogActor struct for LogActor
type LogActor struct {
	AlternateId *string `json:"alternateId,omitempty"`
	DetailEntry map[string]interface{} `json:"detailEntry,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogActor LogActor

// NewLogActor instantiates a new LogActor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogActor() *LogActor {
	this := LogActor{}
	return &this
}

// NewLogActorWithDefaults instantiates a new LogActor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogActorWithDefaults() *LogActor {
	this := LogActor{}
	return &this
}

// GetAlternateId returns the AlternateId field value if set, zero value otherwise.
func (o *LogActor) GetAlternateId() string {
	if o == nil || o.AlternateId == nil {
		var ret string
		return ret
	}
	return *o.AlternateId
}

// GetAlternateIdOk returns a tuple with the AlternateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogActor) GetAlternateIdOk() (*string, bool) {
	if o == nil || o.AlternateId == nil {
		return nil, false
	}
	return o.AlternateId, true
}

// HasAlternateId returns a boolean if a field has been set.
func (o *LogActor) HasAlternateId() bool {
	if o != nil && o.AlternateId != nil {
		return true
	}

	return false
}

// SetAlternateId gets a reference to the given string and assigns it to the AlternateId field.
func (o *LogActor) SetAlternateId(v string) {
	o.AlternateId = &v
}

// GetDetailEntry returns the DetailEntry field value if set, zero value otherwise.
func (o *LogActor) GetDetailEntry() map[string]interface{} {
	if o == nil || o.DetailEntry == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DetailEntry
}

// GetDetailEntryOk returns a tuple with the DetailEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogActor) GetDetailEntryOk() (map[string]interface{}, bool) {
	if o == nil || o.DetailEntry == nil {
		return nil, false
	}
	return o.DetailEntry, true
}

// HasDetailEntry returns a boolean if a field has been set.
func (o *LogActor) HasDetailEntry() bool {
	if o != nil && o.DetailEntry != nil {
		return true
	}

	return false
}

// SetDetailEntry gets a reference to the given map[string]interface{} and assigns it to the DetailEntry field.
func (o *LogActor) SetDetailEntry(v map[string]interface{}) {
	o.DetailEntry = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *LogActor) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogActor) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *LogActor) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *LogActor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LogActor) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogActor) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogActor) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LogActor) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LogActor) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogActor) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogActor) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LogActor) SetType(v string) {
	o.Type = &v
}

func (o LogActor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AlternateId != nil {
		toSerialize["alternateId"] = o.AlternateId
	}
	if o.DetailEntry != nil {
		toSerialize["detailEntry"] = o.DetailEntry
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogActor) UnmarshalJSON(bytes []byte) (err error) {
	varLogActor := _LogActor{}

	err = json.Unmarshal(bytes, &varLogActor)
	if err == nil {
		*o = LogActor(varLogActor)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "alternateId")
		delete(additionalProperties, "detailEntry")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogActor struct {
	value *LogActor
	isSet bool
}

func (v NullableLogActor) Get() *LogActor {
	return v.value
}

func (v *NullableLogActor) Set(val *LogActor) {
	v.value = val
	v.isSet = true
}

func (v NullableLogActor) IsSet() bool {
	return v.isSet
}

func (v *NullableLogActor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogActor(val *LogActor) *NullableLogActor {
	return &NullableLogActor{value: val, isSet: true}
}

func (v NullableLogActor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogActor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

