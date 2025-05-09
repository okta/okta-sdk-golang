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

// StreamConfigurationDelivery Contains information about the intended SET delivery method by the receiver
type StreamConfigurationDelivery struct {
	// The HTTP Authorization header that is included for each HTTP POST request
	AuthorizationHeader NullableString `json:"authorization_header,omitempty"`
	// The target endpoint URL where the transmitter delivers the SET using HTTP POST requests
	EndpointUrl string `json:"endpoint_url"`
	// The delivery method that the transmitter uses for delivering a SET
	Method string `json:"method"`
	AdditionalProperties map[string]interface{}
}

type _StreamConfigurationDelivery StreamConfigurationDelivery

// NewStreamConfigurationDelivery instantiates a new StreamConfigurationDelivery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamConfigurationDelivery(endpointUrl string, method string) *StreamConfigurationDelivery {
	this := StreamConfigurationDelivery{}
	this.EndpointUrl = endpointUrl
	this.Method = method
	return &this
}

// NewStreamConfigurationDeliveryWithDefaults instantiates a new StreamConfigurationDelivery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamConfigurationDeliveryWithDefaults() *StreamConfigurationDelivery {
	this := StreamConfigurationDelivery{}
	return &this
}

// GetAuthorizationHeader returns the AuthorizationHeader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StreamConfigurationDelivery) GetAuthorizationHeader() string {
	if o == nil || o.AuthorizationHeader.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationHeader.Get()
}

// GetAuthorizationHeaderOk returns a tuple with the AuthorizationHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StreamConfigurationDelivery) GetAuthorizationHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationHeader.Get(), o.AuthorizationHeader.IsSet()
}

// HasAuthorizationHeader returns a boolean if a field has been set.
func (o *StreamConfigurationDelivery) HasAuthorizationHeader() bool {
	if o != nil && o.AuthorizationHeader.IsSet() {
		return true
	}

	return false
}

// SetAuthorizationHeader gets a reference to the given NullableString and assigns it to the AuthorizationHeader field.
func (o *StreamConfigurationDelivery) SetAuthorizationHeader(v string) {
	o.AuthorizationHeader.Set(&v)
}
// SetAuthorizationHeaderNil sets the value for AuthorizationHeader to be an explicit nil
func (o *StreamConfigurationDelivery) SetAuthorizationHeaderNil() {
	o.AuthorizationHeader.Set(nil)
}

// UnsetAuthorizationHeader ensures that no value is present for AuthorizationHeader, not even an explicit nil
func (o *StreamConfigurationDelivery) UnsetAuthorizationHeader() {
	o.AuthorizationHeader.Unset()
}

// GetEndpointUrl returns the EndpointUrl field value
func (o *StreamConfigurationDelivery) GetEndpointUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointUrl
}

// GetEndpointUrlOk returns a tuple with the EndpointUrl field value
// and a boolean to check if the value has been set.
func (o *StreamConfigurationDelivery) GetEndpointUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointUrl, true
}

// SetEndpointUrl sets field value
func (o *StreamConfigurationDelivery) SetEndpointUrl(v string) {
	o.EndpointUrl = v
}

// GetMethod returns the Method field value
func (o *StreamConfigurationDelivery) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *StreamConfigurationDelivery) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *StreamConfigurationDelivery) SetMethod(v string) {
	o.Method = v
}

func (o StreamConfigurationDelivery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorizationHeader.IsSet() {
		toSerialize["authorization_header"] = o.AuthorizationHeader.Get()
	}
	if true {
		toSerialize["endpoint_url"] = o.EndpointUrl
	}
	if true {
		toSerialize["method"] = o.Method
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StreamConfigurationDelivery) UnmarshalJSON(bytes []byte) (err error) {
	varStreamConfigurationDelivery := _StreamConfigurationDelivery{}

	err = json.Unmarshal(bytes, &varStreamConfigurationDelivery)
	if err == nil {
		*o = StreamConfigurationDelivery(varStreamConfigurationDelivery)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authorization_header")
		delete(additionalProperties, "endpoint_url")
		delete(additionalProperties, "method")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableStreamConfigurationDelivery struct {
	value *StreamConfigurationDelivery
	isSet bool
}

func (v NullableStreamConfigurationDelivery) Get() *StreamConfigurationDelivery {
	return v.value
}

func (v *NullableStreamConfigurationDelivery) Set(val *StreamConfigurationDelivery) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamConfigurationDelivery) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamConfigurationDelivery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamConfigurationDelivery(val *StreamConfigurationDelivery) *NullableStreamConfigurationDelivery {
	return &NullableStreamConfigurationDelivery{value: val, isSet: true}
}

func (v NullableStreamConfigurationDelivery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamConfigurationDelivery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

