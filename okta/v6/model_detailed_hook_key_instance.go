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
	"time"
)

// DetailedHookKeyInstance A key object with public key details
type DetailedHookKeyInstance struct {
	// Timestamp when the key was created
	Created NullableTime `json:"created,omitempty"`
	// The unique Okta ID of this key record
	Id *string `json:"id,omitempty"`
	// Whether this key is currently in use by other applications
	IsUsed *bool `json:"isUsed,omitempty"`
	// The alias of the public key
	KeyId *string `json:"keyId,omitempty"`
	// Timestamp when the key was updated
	LastUpdated NullableTime `json:"lastUpdated,omitempty"`
	// Display name of the key
	Name *string `json:"name,omitempty"`
	Embedded *Embedded `json:"_embedded,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DetailedHookKeyInstance DetailedHookKeyInstance

// NewDetailedHookKeyInstance instantiates a new DetailedHookKeyInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedHookKeyInstance() *DetailedHookKeyInstance {
	this := DetailedHookKeyInstance{}
	return &this
}

// NewDetailedHookKeyInstanceWithDefaults instantiates a new DetailedHookKeyInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedHookKeyInstanceWithDefaults() *DetailedHookKeyInstance {
	this := DetailedHookKeyInstance{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetailedHookKeyInstance) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetailedHookKeyInstance) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *DetailedHookKeyInstance) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *DetailedHookKeyInstance) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *DetailedHookKeyInstance) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *DetailedHookKeyInstance) UnsetCreated() {
	o.Created.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DetailedHookKeyInstance) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedHookKeyInstance) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DetailedHookKeyInstance) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DetailedHookKeyInstance) SetId(v string) {
	o.Id = &v
}

// GetIsUsed returns the IsUsed field value if set, zero value otherwise.
func (o *DetailedHookKeyInstance) GetIsUsed() bool {
	if o == nil || o.IsUsed == nil {
		var ret bool
		return ret
	}
	return *o.IsUsed
}

// GetIsUsedOk returns a tuple with the IsUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedHookKeyInstance) GetIsUsedOk() (*bool, bool) {
	if o == nil || o.IsUsed == nil {
		return nil, false
	}
	return o.IsUsed, true
}

// HasIsUsed returns a boolean if a field has been set.
func (o *DetailedHookKeyInstance) HasIsUsed() bool {
	if o != nil && o.IsUsed != nil {
		return true
	}

	return false
}

// SetIsUsed gets a reference to the given bool and assigns it to the IsUsed field.
func (o *DetailedHookKeyInstance) SetIsUsed(v bool) {
	o.IsUsed = &v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *DetailedHookKeyInstance) GetKeyId() string {
	if o == nil || o.KeyId == nil {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedHookKeyInstance) GetKeyIdOk() (*string, bool) {
	if o == nil || o.KeyId == nil {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *DetailedHookKeyInstance) HasKeyId() bool {
	if o != nil && o.KeyId != nil {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *DetailedHookKeyInstance) SetKeyId(v string) {
	o.KeyId = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetailedHookKeyInstance) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetailedHookKeyInstance) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *DetailedHookKeyInstance) HasLastUpdated() bool {
	if o != nil && o.LastUpdated.IsSet() {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given NullableTime and assigns it to the LastUpdated field.
func (o *DetailedHookKeyInstance) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}
// SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil
func (o *DetailedHookKeyInstance) SetLastUpdatedNil() {
	o.LastUpdated.Set(nil)
}

// UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
func (o *DetailedHookKeyInstance) UnsetLastUpdated() {
	o.LastUpdated.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DetailedHookKeyInstance) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedHookKeyInstance) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DetailedHookKeyInstance) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DetailedHookKeyInstance) SetName(v string) {
	o.Name = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *DetailedHookKeyInstance) GetEmbedded() Embedded {
	if o == nil || o.Embedded == nil {
		var ret Embedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedHookKeyInstance) GetEmbeddedOk() (*Embedded, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *DetailedHookKeyInstance) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given Embedded and assigns it to the Embedded field.
func (o *DetailedHookKeyInstance) SetEmbedded(v Embedded) {
	o.Embedded = &v
}

func (o DetailedHookKeyInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
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
	if o.LastUpdated.IsSet() {
		toSerialize["lastUpdated"] = o.LastUpdated.Get()
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

func (o *DetailedHookKeyInstance) UnmarshalJSON(bytes []byte) (err error) {
	varDetailedHookKeyInstance := _DetailedHookKeyInstance{}

	err = json.Unmarshal(bytes, &varDetailedHookKeyInstance)
	if err == nil {
		*o = DetailedHookKeyInstance(varDetailedHookKeyInstance)
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

type NullableDetailedHookKeyInstance struct {
	value *DetailedHookKeyInstance
	isSet bool
}

func (v NullableDetailedHookKeyInstance) Get() *DetailedHookKeyInstance {
	return v.value
}

func (v *NullableDetailedHookKeyInstance) Set(val *DetailedHookKeyInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedHookKeyInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedHookKeyInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedHookKeyInstance(val *DetailedHookKeyInstance) *NullableDetailedHookKeyInstance {
	return &NullableDetailedHookKeyInstance{value: val, isSet: true}
}

func (v NullableDetailedHookKeyInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedHookKeyInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

