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
	"time"
)

// HookKey struct for HookKey
type HookKey struct {
	// Timestamp when the key was created.
	Created *time.Time `json:"created,omitempty"`
	// The unique identifier for the key.
	Id *string `json:"id,omitempty"`
	// Whether this key is currently in use by other hooks.
	IsUsed *bool `json:"isUsed,omitempty"`
	// The alias of the public key.
	KeyId *string `json:"keyId,omitempty"`
	// Timestamp when the key was updated.
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Display name of the key.
	Name *string `json:"name,omitempty"`
	Embedded *JsonWebKey `json:"_embedded,omitempty"`
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

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *HookKey) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookKey) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *HookKey) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *HookKey) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HookKey) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookKey) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HookKey) HasId() bool {
	if o != nil && o.Id != nil {
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
	if o == nil || o.IsUsed == nil {
		var ret bool
		return ret
	}
	return *o.IsUsed
}

// GetIsUsedOk returns a tuple with the IsUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookKey) GetIsUsedOk() (*bool, bool) {
	if o == nil || o.IsUsed == nil {
		return nil, false
	}
	return o.IsUsed, true
}

// HasIsUsed returns a boolean if a field has been set.
func (o *HookKey) HasIsUsed() bool {
	if o != nil && o.IsUsed != nil {
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
	if o == nil || o.KeyId == nil {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookKey) GetKeyIdOk() (*string, bool) {
	if o == nil || o.KeyId == nil {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *HookKey) HasKeyId() bool {
	if o != nil && o.KeyId != nil {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *HookKey) SetKeyId(v string) {
	o.KeyId = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *HookKey) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookKey) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *HookKey) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *HookKey) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HookKey) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookKey) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HookKey) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HookKey) SetName(v string) {
	o.Name = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *HookKey) GetEmbedded() JsonWebKey {
	if o == nil || o.Embedded == nil {
		var ret JsonWebKey
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookKey) GetEmbeddedOk() (*JsonWebKey, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *HookKey) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given JsonWebKey and assigns it to the Embedded field.
func (o *HookKey) SetEmbedded(v JsonWebKey) {
	o.Embedded = &v
}

func (o HookKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsUsed != nil {
		toSerialize["isUsed"] = o.IsUsed
	}
	if o.KeyId != nil {
		toSerialize["keyId"] = o.KeyId
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Embedded != nil {
		toSerialize["_embedded"] = o.Embedded
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HookKey) UnmarshalJSON(bytes []byte) (err error) {
	varHookKey := _HookKey{}

	err = json.Unmarshal(bytes, &varHookKey)
	if err == nil {
		*o = HookKey(varHookKey)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "isUsed")
		delete(additionalProperties, "keyId")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "_embedded")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

