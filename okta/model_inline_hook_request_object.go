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

// checks if the InlineHookRequestObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineHookRequestObject{}

// InlineHookRequestObject The API request that triggered the inline hook
type InlineHookRequestObject struct {
	// The unique identifier that Okta assigned to the API request
	Id *string `json:"id,omitempty"`
	// The IP address of the client that made the API request
	IpAddress *string `json:"ipAddress,omitempty"`
	// The HTTP request method of the API request
	Method               *string                     `json:"method,omitempty"`
	Url                  *InlineHookRequestObjectUrl `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookRequestObject InlineHookRequestObject

// NewInlineHookRequestObject instantiates a new InlineHookRequestObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookRequestObject() *InlineHookRequestObject {
	this := InlineHookRequestObject{}
	return &this
}

// NewInlineHookRequestObjectWithDefaults instantiates a new InlineHookRequestObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookRequestObjectWithDefaults() *InlineHookRequestObject {
	this := InlineHookRequestObject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineHookRequestObject) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookRequestObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineHookRequestObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineHookRequestObject) SetId(v string) {
	o.Id = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *InlineHookRequestObject) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookRequestObject) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *InlineHookRequestObject) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *InlineHookRequestObject) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *InlineHookRequestObject) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookRequestObject) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *InlineHookRequestObject) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *InlineHookRequestObject) SetMethod(v string) {
	o.Method = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineHookRequestObject) GetUrl() InlineHookRequestObjectUrl {
	if o == nil || IsNil(o.Url) {
		var ret InlineHookRequestObjectUrl
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookRequestObject) GetUrlOk() (*InlineHookRequestObjectUrl, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineHookRequestObject) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given InlineHookRequestObjectUrl and assigns it to the Url field.
func (o *InlineHookRequestObject) SetUrl(v InlineHookRequestObjectUrl) {
	o.Url = &v
}

func (o InlineHookRequestObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineHookRequestObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineHookRequestObject) UnmarshalJSON(data []byte) (err error) {
	varInlineHookRequestObject := _InlineHookRequestObject{}

	err = json.Unmarshal(data, &varInlineHookRequestObject)

	if err != nil {
		return err
	}

	*o = InlineHookRequestObject(varInlineHookRequestObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "ipAddress")
		delete(additionalProperties, "method")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineHookRequestObject struct {
	value *InlineHookRequestObject
	isSet bool
}

func (v NullableInlineHookRequestObject) Get() *InlineHookRequestObject {
	return v.value
}

func (v *NullableInlineHookRequestObject) Set(val *InlineHookRequestObject) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookRequestObject) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookRequestObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookRequestObject(val *InlineHookRequestObject) *NullableInlineHookRequestObject {
	return &NullableInlineHookRequestObject{value: val, isSet: true}
}

func (v NullableInlineHookRequestObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookRequestObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
