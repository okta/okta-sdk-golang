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

// LogStreamSelfLink struct for LogStreamSelfLink
type LogStreamSelfLink struct {
	// The URI of the resource
	Href string `json:"href"`
	// HTTP method allowed for the resource
	Method *string `json:"method,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogStreamSelfLink LogStreamSelfLink

// NewLogStreamSelfLink instantiates a new LogStreamSelfLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStreamSelfLink(href string) *LogStreamSelfLink {
	this := LogStreamSelfLink{}
	this.Href = href
	return &this
}

// NewLogStreamSelfLinkWithDefaults instantiates a new LogStreamSelfLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamSelfLinkWithDefaults() *LogStreamSelfLink {
	this := LogStreamSelfLink{}
	return &this
}

// GetHref returns the Href field value
func (o *LogStreamSelfLink) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *LogStreamSelfLink) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *LogStreamSelfLink) SetHref(v string) {
	o.Href = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *LogStreamSelfLink) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSelfLink) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *LogStreamSelfLink) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *LogStreamSelfLink) SetMethod(v string) {
	o.Method = &v
}

func (o LogStreamSelfLink) MarshalJSON() ([]byte, error) {
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

func (o *LogStreamSelfLink) UnmarshalJSON(bytes []byte) (err error) {
	varLogStreamSelfLink := _LogStreamSelfLink{}

	err = json.Unmarshal(bytes, &varLogStreamSelfLink)
	if err == nil {
		*o = LogStreamSelfLink(varLogStreamSelfLink)
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

type NullableLogStreamSelfLink struct {
	value *LogStreamSelfLink
	isSet bool
}

func (v NullableLogStreamSelfLink) Get() *LogStreamSelfLink {
	return v.value
}

func (v *NullableLogStreamSelfLink) Set(val *LogStreamSelfLink) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStreamSelfLink) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStreamSelfLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStreamSelfLink(val *LogStreamSelfLink) *NullableLogStreamSelfLink {
	return &NullableLogStreamSelfLink{value: val, isSet: true}
}

func (v NullableLogStreamSelfLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStreamSelfLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

