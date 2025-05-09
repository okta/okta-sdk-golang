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

// LogTransaction struct for LogTransaction
type LogTransaction struct {
	Detail map[string]interface{} `json:"detail,omitempty"`
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogTransaction LogTransaction

// NewLogTransaction instantiates a new LogTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogTransaction() *LogTransaction {
	this := LogTransaction{}
	return &this
}

// NewLogTransactionWithDefaults instantiates a new LogTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogTransactionWithDefaults() *LogTransaction {
	this := LogTransaction{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *LogTransaction) GetDetail() map[string]interface{} {
	if o == nil || o.Detail == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogTransaction) GetDetailOk() (map[string]interface{}, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *LogTransaction) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given map[string]interface{} and assigns it to the Detail field.
func (o *LogTransaction) SetDetail(v map[string]interface{}) {
	o.Detail = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LogTransaction) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogTransaction) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogTransaction) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LogTransaction) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LogTransaction) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogTransaction) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogTransaction) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LogTransaction) SetType(v string) {
	o.Type = &v
}

func (o LogTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
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

func (o *LogTransaction) UnmarshalJSON(bytes []byte) (err error) {
	varLogTransaction := _LogTransaction{}

	err = json.Unmarshal(bytes, &varLogTransaction)
	if err == nil {
		*o = LogTransaction(varLogTransaction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "detail")
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogTransaction struct {
	value *LogTransaction
	isSet bool
}

func (v NullableLogTransaction) Get() *LogTransaction {
	return v.value
}

func (v *NullableLogTransaction) Set(val *LogTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableLogTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableLogTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogTransaction(val *LogTransaction) *NullableLogTransaction {
	return &NullableLogTransaction{value: val, isSet: true}
}

func (v NullableLogTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

