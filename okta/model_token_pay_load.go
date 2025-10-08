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

// checks if the TokenPayLoad type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenPayLoad{}

// TokenPayLoad struct for TokenPayLoad
type TokenPayLoad struct {
	Data *TokenPayLoadData `json:"data,omitempty"`
	// The type of inline hook. The token inline hook type is `com.okta.oauth2.tokens.transform`.
	EventType *string `json:"eventType,omitempty"`
	// The URL of the token inline hook
	Source               *string `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenPayLoad TokenPayLoad

// NewTokenPayLoad instantiates a new TokenPayLoad object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenPayLoad() *TokenPayLoad {
	this := TokenPayLoad{}
	return &this
}

// NewTokenPayLoadWithDefaults instantiates a new TokenPayLoad object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenPayLoadWithDefaults() *TokenPayLoad {
	this := TokenPayLoad{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TokenPayLoad) GetData() TokenPayLoadData {
	if o == nil || IsNil(o.Data) {
		var ret TokenPayLoadData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoad) GetDataOk() (*TokenPayLoadData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TokenPayLoad) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given TokenPayLoadData and assigns it to the Data field.
func (o *TokenPayLoad) SetData(v TokenPayLoadData) {
	o.Data = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *TokenPayLoad) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoad) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *TokenPayLoad) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *TokenPayLoad) SetEventType(v string) {
	o.EventType = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *TokenPayLoad) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoad) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *TokenPayLoad) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *TokenPayLoad) SetSource(v string) {
	o.Source = &v
}

func (o TokenPayLoad) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenPayLoad) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenPayLoad) UnmarshalJSON(data []byte) (err error) {
	varTokenPayLoad := _TokenPayLoad{}

	err = json.Unmarshal(data, &varTokenPayLoad)

	if err != nil {
		return err
	}

	*o = TokenPayLoad(varTokenPayLoad)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "eventType")
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenPayLoad struct {
	value *TokenPayLoad
	isSet bool
}

func (v NullableTokenPayLoad) Get() *TokenPayLoad {
	return v.value
}

func (v *NullableTokenPayLoad) Set(val *TokenPayLoad) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenPayLoad) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenPayLoad) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenPayLoad(val *TokenPayLoad) *NullableTokenPayLoad {
	return &NullableTokenPayLoad{value: val, isSet: true}
}

func (v NullableTokenPayLoad) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenPayLoad) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
