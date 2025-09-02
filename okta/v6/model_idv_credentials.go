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
)

// IDVCredentials Credentials for verifying requests to the IDV
type IDVCredentials struct {
	Bearer *IDVCredentialsBearer `json:"bearer,omitempty"`
	Client *IDVCredentialsClient `json:"client,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IDVCredentials IDVCredentials

// NewIDVCredentials instantiates a new IDVCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIDVCredentials() *IDVCredentials {
	this := IDVCredentials{}
	return &this
}

// NewIDVCredentialsWithDefaults instantiates a new IDVCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIDVCredentialsWithDefaults() *IDVCredentials {
	this := IDVCredentials{}
	return &this
}

// GetBearer returns the Bearer field value if set, zero value otherwise.
func (o *IDVCredentials) GetBearer() IDVCredentialsBearer {
	if o == nil || o.Bearer == nil {
		var ret IDVCredentialsBearer
		return ret
	}
	return *o.Bearer
}

// GetBearerOk returns a tuple with the Bearer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDVCredentials) GetBearerOk() (*IDVCredentialsBearer, bool) {
	if o == nil || o.Bearer == nil {
		return nil, false
	}
	return o.Bearer, true
}

// HasBearer returns a boolean if a field has been set.
func (o *IDVCredentials) HasBearer() bool {
	if o != nil && o.Bearer != nil {
		return true
	}

	return false
}

// SetBearer gets a reference to the given IDVCredentialsBearer and assigns it to the Bearer field.
func (o *IDVCredentials) SetBearer(v IDVCredentialsBearer) {
	o.Bearer = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *IDVCredentials) GetClient() IDVCredentialsClient {
	if o == nil || o.Client == nil {
		var ret IDVCredentialsClient
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDVCredentials) GetClientOk() (*IDVCredentialsClient, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *IDVCredentials) HasClient() bool {
	if o != nil && o.Client != nil {
		return true
	}

	return false
}

// SetClient gets a reference to the given IDVCredentialsClient and assigns it to the Client field.
func (o *IDVCredentials) SetClient(v IDVCredentialsClient) {
	o.Client = &v
}

func (o IDVCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bearer != nil {
		toSerialize["bearer"] = o.Bearer
	}
	if o.Client != nil {
		toSerialize["client"] = o.Client
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IDVCredentials) UnmarshalJSON(bytes []byte) (err error) {
	varIDVCredentials := _IDVCredentials{}

	err = json.Unmarshal(bytes, &varIDVCredentials)
	if err == nil {
		*o = IDVCredentials(varIDVCredentials)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "bearer")
		delete(additionalProperties, "client")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIDVCredentials struct {
	value *IDVCredentials
	isSet bool
}

func (v NullableIDVCredentials) Get() *IDVCredentials {
	return v.value
}

func (v *NullableIDVCredentials) Set(val *IDVCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableIDVCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableIDVCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIDVCredentials(val *IDVCredentials) *NullableIDVCredentials {
	return &NullableIDVCredentials{value: val, isSet: true}
}

func (v NullableIDVCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIDVCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

