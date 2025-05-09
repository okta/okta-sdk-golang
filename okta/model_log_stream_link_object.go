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

// LogStreamLinkObject struct for LogStreamLinkObject
type LogStreamLinkObject struct {
	// The URI of the resource
	Href string `json:"href"`
	// HTTP method allowed for the resource
	Method *string `json:"method,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogStreamLinkObject LogStreamLinkObject

// NewLogStreamLinkObject instantiates a new LogStreamLinkObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStreamLinkObject(href string) *LogStreamLinkObject {
	this := LogStreamLinkObject{}
	this.Href = href
	return &this
}

// NewLogStreamLinkObjectWithDefaults instantiates a new LogStreamLinkObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamLinkObjectWithDefaults() *LogStreamLinkObject {
	this := LogStreamLinkObject{}
	return &this
}

// GetHref returns the Href field value
func (o *LogStreamLinkObject) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *LogStreamLinkObject) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *LogStreamLinkObject) SetHref(v string) {
	o.Href = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *LogStreamLinkObject) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamLinkObject) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *LogStreamLinkObject) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *LogStreamLinkObject) SetMethod(v string) {
	o.Method = &v
}

func (o LogStreamLinkObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["href"] = o.Href
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogStreamLinkObject) UnmarshalJSON(bytes []byte) (err error) {
	varLogStreamLinkObject := _LogStreamLinkObject{}

	err = json.Unmarshal(bytes, &varLogStreamLinkObject)
	if err == nil {
		*o = LogStreamLinkObject(varLogStreamLinkObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "method")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogStreamLinkObject struct {
	value *LogStreamLinkObject
	isSet bool
}

func (v NullableLogStreamLinkObject) Get() *LogStreamLinkObject {
	return v.value
}

func (v *NullableLogStreamLinkObject) Set(val *LogStreamLinkObject) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStreamLinkObject) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStreamLinkObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStreamLinkObject(val *LogStreamLinkObject) *NullableLogStreamLinkObject {
	return &NullableLogStreamLinkObject{value: val, isSet: true}
}

func (v NullableLogStreamLinkObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStreamLinkObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

