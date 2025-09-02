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
	"fmt"
)

// IDVCredentialsClient <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Client credentials for `IDV_CLEAR` and `IDV_INCODE` IdP types
type IDVCredentialsClient struct {
	// The client ID that you generate in your IDV
	ClientId string `json:"client_id"`
	// The client secret that you generate in your IDV
	ClientSecret string `json:"client_secret"`
	AdditionalProperties map[string]interface{}
}

type _IDVCredentialsClient IDVCredentialsClient

// NewIDVCredentialsClient instantiates a new IDVCredentialsClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIDVCredentialsClient(clientId string, clientSecret string) *IDVCredentialsClient {
	this := IDVCredentialsClient{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	return &this
}

// NewIDVCredentialsClientWithDefaults instantiates a new IDVCredentialsClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIDVCredentialsClientWithDefaults() *IDVCredentialsClient {
	this := IDVCredentialsClient{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *IDVCredentialsClient) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *IDVCredentialsClient) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *IDVCredentialsClient) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *IDVCredentialsClient) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *IDVCredentialsClient) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *IDVCredentialsClient) SetClientSecret(v string) {
	o.ClientSecret = v
}

func (o IDVCredentialsClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["client_secret"] = o.ClientSecret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IDVCredentialsClient) UnmarshalJSON(bytes []byte) (err error) {
	varIDVCredentialsClient := _IDVCredentialsClient{}

	err = json.Unmarshal(bytes, &varIDVCredentialsClient)
	if err == nil {
		*o = IDVCredentialsClient(varIDVCredentialsClient)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "client_secret")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIDVCredentialsClient struct {
	value *IDVCredentialsClient
	isSet bool
}

func (v NullableIDVCredentialsClient) Get() *IDVCredentialsClient {
	return v.value
}

func (v *NullableIDVCredentialsClient) Set(val *IDVCredentialsClient) {
	v.value = val
	v.isSet = true
}

func (v NullableIDVCredentialsClient) IsSet() bool {
	return v.isSet
}

func (v *NullableIDVCredentialsClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIDVCredentialsClient(val *IDVCredentialsClient) *NullableIDVCredentialsClient {
	return &NullableIDVCredentialsClient{value: val, isSet: true}
}

func (v NullableIDVCredentialsClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIDVCredentialsClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

