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
)

// checks if the LogActor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogActor{}

// LogActor Describes the user, app, client, or other entity (actor) who performs an action on a target. The actor is dependent on the action that is performed. All events have actors.
type LogActor struct {
	// Alternative ID of the actor
	AlternateId *string `json:"alternateId,omitempty"`
	// Further details about the actor
	DetailEntry map[string]interface{} `json:"detailEntry,omitempty"`
	// Display name of the actor
	DisplayName *string `json:"displayName,omitempty"`
	// ID of the actor
	Id *string `json:"id,omitempty"`
	// Type of actor
	Type                 *string `json:"type,omitempty"`
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
	if o == nil || IsNil(o.AlternateId) {
		var ret string
		return ret
	}
	return *o.AlternateId
}

// GetAlternateIdOk returns a tuple with the AlternateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogActor) GetAlternateIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlternateId) {
		return nil, false
	}
	return o.AlternateId, true
}

// HasAlternateId returns a boolean if a field has been set.
func (o *LogActor) HasAlternateId() bool {
	if o != nil && !IsNil(o.AlternateId) {
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
	if o == nil || IsNil(o.DetailEntry) {
		var ret map[string]interface{}
		return ret
	}
	return o.DetailEntry
}

// GetDetailEntryOk returns a tuple with the DetailEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogActor) GetDetailEntryOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DetailEntry) {
		return map[string]interface{}{}, false
	}
	return o.DetailEntry, true
}

// HasDetailEntry returns a boolean if a field has been set.
func (o *LogActor) HasDetailEntry() bool {
	if o != nil && !IsNil(o.DetailEntry) {
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
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogActor) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *LogActor) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
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
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogActor) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogActor) HasId() bool {
	if o != nil && !IsNil(o.Id) {
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
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogActor) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogActor) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LogActor) SetType(v string) {
	o.Type = &v
}

func (o LogActor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogActor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlternateId) {
		toSerialize["alternateId"] = o.AlternateId
	}
	if !IsNil(o.DetailEntry) {
		toSerialize["detailEntry"] = o.DetailEntry
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogActor) UnmarshalJSON(data []byte) (err error) {
	varLogActor := _LogActor{}

	err = json.Unmarshal(data, &varLogActor)

	if err != nil {
		return err
	}

	*o = LogActor(varLogActor)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alternateId")
		delete(additionalProperties, "detailEntry")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
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
