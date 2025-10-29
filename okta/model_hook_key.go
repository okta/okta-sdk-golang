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
	"time"
)

// checks if the HookKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HookKey{}

// HookKey The `id` property in the response as `id` serves as the unique ID for the key, which you can specify when invoking other CRUD operations.  The `keyId` provided in the response is the alias of the public key that you can use to get details of the public key data in a separate call.
type HookKey struct {
	// Timestamp when the key was created
	Created NullableTime `json:"created,omitempty"`
	// The unique identifier for the key
	Id *string `json:"id,omitempty"`
	// Whether this key is currently in use by other applications
	IsUsed *bool `json:"isUsed,omitempty"`
	// The alias of the public key
	KeyId *string `json:"keyId,omitempty"`
	// Timestamp when the key was updated
	LastUpdated NullableTime `json:"lastUpdated,omitempty"`
	// Display name of the key
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HookKey HookKey

// NewHookKey instantiates a new HookKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHookKey() *HookKey {
	this := HookKey{}
	return &this
}

// NewHookKeyWithDefaults instantiates a new HookKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHookKeyWithDefaults() *HookKey {
	this := HookKey{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HookKey) GetCreated() time.Time {
	if o == nil || IsNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HookKey) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *HookKey) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *HookKey) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// SetCreatedNil sets the value for Created to be an explicit nil
func (o *HookKey) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *HookKey) UnsetCreated() {
	o.Created.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HookKey) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookKey) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HookKey) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HookKey) SetId(v string) {
	o.Id = &v
}

// GetIsUsed returns the IsUsed field value if set, zero value otherwise.
func (o *HookKey) GetIsUsed() bool {
	if o == nil || IsNil(o.IsUsed) {
		var ret bool
		return ret
	}
	return *o.IsUsed
}

// GetIsUsedOk returns a tuple with the IsUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookKey) GetIsUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUsed) {
		return nil, false
	}
	return o.IsUsed, true
}

// HasIsUsed returns a boolean if a field has been set.
func (o *HookKey) HasIsUsed() bool {
	if o != nil && !IsNil(o.IsUsed) {
		return true
	}

	return false
}

// SetIsUsed gets a reference to the given bool and assigns it to the IsUsed field.
func (o *HookKey) SetIsUsed(v bool) {
	o.IsUsed = &v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *HookKey) GetKeyId() string {
	if o == nil || IsNil(o.KeyId) {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookKey) GetKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeyId) {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *HookKey) HasKeyId() bool {
	if o != nil && !IsNil(o.KeyId) {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *HookKey) SetKeyId(v string) {
	o.KeyId = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HookKey) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HookKey) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *HookKey) HasLastUpdated() bool {
	if o != nil && o.LastUpdated.IsSet() {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given NullableTime and assigns it to the LastUpdated field.
func (o *HookKey) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil
func (o *HookKey) SetLastUpdatedNil() {
	o.LastUpdated.Set(nil)
}

// UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
func (o *HookKey) UnsetLastUpdated() {
	o.LastUpdated.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HookKey) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookKey) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HookKey) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HookKey) SetName(v string) {
	o.Name = &v
}

func (o HookKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HookKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsUsed) {
		toSerialize["isUsed"] = o.IsUsed
	}
	if !IsNil(o.KeyId) {
		toSerialize["keyId"] = o.KeyId
	}
	if o.LastUpdated.IsSet() {
		toSerialize["lastUpdated"] = o.LastUpdated.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HookKey) UnmarshalJSON(data []byte) (err error) {
	varHookKey := _HookKey{}

	err = json.Unmarshal(data, &varHookKey)

	if err != nil {
		return err
	}

	*o = HookKey(varHookKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "isUsed")
		delete(additionalProperties, "keyId")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHookKey struct {
	value *HookKey
	isSet bool
}

func (v NullableHookKey) Get() *HookKey {
	return v.value
}

func (v *NullableHookKey) Set(val *HookKey) {
	v.value = val
	v.isSet = true
}

func (v NullableHookKey) IsSet() bool {
	return v.isSet
}

func (v *NullableHookKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHookKey(val *HookKey) *NullableHookKey {
	return &NullableHookKey{value: val, isSet: true}
}

func (v NullableHookKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHookKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
