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

// ProtocolEndpoint struct for ProtocolEndpoint
type ProtocolEndpoint struct {
	Binding *string `json:"binding,omitempty"`
	Destination *string `json:"destination,omitempty"`
	Type *string `json:"type,omitempty"`
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProtocolEndpoint ProtocolEndpoint

// NewProtocolEndpoint instantiates a new ProtocolEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocolEndpoint() *ProtocolEndpoint {
	this := ProtocolEndpoint{}
	return &this
}

// NewProtocolEndpointWithDefaults instantiates a new ProtocolEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolEndpointWithDefaults() *ProtocolEndpoint {
	this := ProtocolEndpoint{}
	return &this
}

// GetBinding returns the Binding field value if set, zero value otherwise.
func (o *ProtocolEndpoint) GetBinding() string {
	if o == nil || o.Binding == nil {
		var ret string
		return ret
	}
	return *o.Binding
}

// GetBindingOk returns a tuple with the Binding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolEndpoint) GetBindingOk() (*string, bool) {
	if o == nil || o.Binding == nil {
		return nil, false
	}
	return o.Binding, true
}

// HasBinding returns a boolean if a field has been set.
func (o *ProtocolEndpoint) HasBinding() bool {
	if o != nil && o.Binding != nil {
		return true
	}

	return false
}

// SetBinding gets a reference to the given string and assigns it to the Binding field.
func (o *ProtocolEndpoint) SetBinding(v string) {
	o.Binding = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *ProtocolEndpoint) GetDestination() string {
	if o == nil || o.Destination == nil {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolEndpoint) GetDestinationOk() (*string, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *ProtocolEndpoint) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *ProtocolEndpoint) SetDestination(v string) {
	o.Destination = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProtocolEndpoint) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolEndpoint) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProtocolEndpoint) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProtocolEndpoint) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ProtocolEndpoint) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolEndpoint) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ProtocolEndpoint) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ProtocolEndpoint) SetUrl(v string) {
	o.Url = &v
}

func (o ProtocolEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Binding != nil {
		toSerialize["binding"] = o.Binding
	}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProtocolEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	varProtocolEndpoint := _ProtocolEndpoint{}

	err = json.Unmarshal(bytes, &varProtocolEndpoint)
	if err == nil {
		*o = ProtocolEndpoint(varProtocolEndpoint)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "binding")
		delete(additionalProperties, "destination")
		delete(additionalProperties, "type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProtocolEndpoint struct {
	value *ProtocolEndpoint
	isSet bool
}

func (v NullableProtocolEndpoint) Get() *ProtocolEndpoint {
	return v.value
}

func (v *NullableProtocolEndpoint) Set(val *ProtocolEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolEndpoint(val *ProtocolEndpoint) *NullableProtocolEndpoint {
	return &NullableProtocolEndpoint{value: val, isSet: true}
}

func (v NullableProtocolEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

