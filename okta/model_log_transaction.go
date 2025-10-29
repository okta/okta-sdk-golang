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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the LogTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogTransaction{}

// LogTransaction A `transaction` object comprises contextual information associated with its respective event. This information is useful for understanding sequences of correlated events. For example, a `transaction` object such as the following: ``` {   \"id\": \"Wn4f-0RQ8D8lTSLkAmkKdQAADqo\",   \"type\": \"WEB\",   \"detail\": null } ``` indicates that a `WEB` request with `id` `Wn4f-0RQ8D8lTSLkAmkKdQAADqo` has created this event.  A `transaction` object with a `requestApiTokenId` in the `detail` object, for example : ``` {   \"id\": \"YjSlblAAqnKY7CdyCkXNBgAAAIU\",   \"type\": \"WEB\",   \"detail\": {     \"requestApiTokenId\": \"00T94e3cn9kSEO3c51s5\"   } } ``` indicates that this event was the result of an action performed through an API using the token identified by 00T94e3cn9kSEO3c51s5. The token ID is visible in the Admin Console, **Security** > **API**. See [API token management](https://help.okta.com/okta_help.htm?id=Security_API). For more information on API tokens, see [Create an API token](https://developer.okta.com/docs/guides/create-an-api-token/).
type LogTransaction struct {
	// Details for this transaction.
	Detail map[string]interface{} `json:"detail,omitempty"`
	// Unique identifier for this transaction.
	Id *string `json:"id,omitempty"`
	// Describes the kind of transaction. `WEB` indicates a web request. `JOB` indicates an asynchronous task.
	Type                 *string `json:"type,omitempty"`
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
	if o == nil || IsNil(o.Detail) {
		var ret map[string]interface{}
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogTransaction) GetDetailOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Detail) {
		return map[string]interface{}{}, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *LogTransaction) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
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
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogTransaction) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogTransaction) HasId() bool {
	if o != nil && !IsNil(o.Id) {
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
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogTransaction) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogTransaction) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LogTransaction) SetType(v string) {
	o.Type = &v
}

func (o LogTransaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
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

func (o *LogTransaction) UnmarshalJSON(data []byte) (err error) {
	varLogTransaction := _LogTransaction{}

	err = json.Unmarshal(data, &varLogTransaction)

	if err != nil {
		return err
	}

	*o = LogTransaction(varLogTransaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "detail")
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
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
