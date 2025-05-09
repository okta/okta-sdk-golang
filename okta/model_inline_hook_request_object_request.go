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

// InlineHookRequestObjectRequest The API request that triggered the inline hook
type InlineHookRequestObjectRequest struct {
	// The unique identifier that Okta assigned to the API request
	Id *string `json:"id,omitempty"`
	// The HTTP request method of the API request
	Method *string `json:"method,omitempty"`
	Url *InlineHookRequestObjectRequestUrl `json:"url,omitempty"`
	// The IP address of the client that made the API request
	IpAddress *string `json:"ipAddress,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookRequestObjectRequest InlineHookRequestObjectRequest

// NewInlineHookRequestObjectRequest instantiates a new InlineHookRequestObjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookRequestObjectRequest() *InlineHookRequestObjectRequest {
	this := InlineHookRequestObjectRequest{}
	return &this
}

// NewInlineHookRequestObjectRequestWithDefaults instantiates a new InlineHookRequestObjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookRequestObjectRequestWithDefaults() *InlineHookRequestObjectRequest {
	this := InlineHookRequestObjectRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineHookRequestObjectRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookRequestObjectRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineHookRequestObjectRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineHookRequestObjectRequest) SetId(v string) {
	o.Id = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *InlineHookRequestObjectRequest) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookRequestObjectRequest) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *InlineHookRequestObjectRequest) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *InlineHookRequestObjectRequest) SetMethod(v string) {
	o.Method = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineHookRequestObjectRequest) GetUrl() InlineHookRequestObjectRequestUrl {
	if o == nil || o.Url == nil {
		var ret InlineHookRequestObjectRequestUrl
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookRequestObjectRequest) GetUrlOk() (*InlineHookRequestObjectRequestUrl, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineHookRequestObjectRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given InlineHookRequestObjectRequestUrl and assigns it to the Url field.
func (o *InlineHookRequestObjectRequest) SetUrl(v InlineHookRequestObjectRequestUrl) {
	o.Url = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *InlineHookRequestObjectRequest) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookRequestObjectRequest) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *InlineHookRequestObjectRequest) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *InlineHookRequestObjectRequest) SetIpAddress(v string) {
	o.IpAddress = &v
}

func (o InlineHookRequestObjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.IpAddress != nil {
		toSerialize["ipAddress"] = o.IpAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineHookRequestObjectRequest) UnmarshalJSON(bytes []byte) (err error) {
	varInlineHookRequestObjectRequest := _InlineHookRequestObjectRequest{}

	err = json.Unmarshal(bytes, &varInlineHookRequestObjectRequest)
	if err == nil {
		*o = InlineHookRequestObjectRequest(varInlineHookRequestObjectRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "method")
		delete(additionalProperties, "url")
		delete(additionalProperties, "ipAddress")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableInlineHookRequestObjectRequest struct {
	value *InlineHookRequestObjectRequest
	isSet bool
}

func (v NullableInlineHookRequestObjectRequest) Get() *InlineHookRequestObjectRequest {
	return v.value
}

func (v *NullableInlineHookRequestObjectRequest) Set(val *InlineHookRequestObjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookRequestObjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookRequestObjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookRequestObjectRequest(val *InlineHookRequestObjectRequest) *NullableInlineHookRequestObjectRequest {
	return &NullableInlineHookRequestObjectRequest{value: val, isSet: true}
}

func (v NullableInlineHookRequestObjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookRequestObjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

