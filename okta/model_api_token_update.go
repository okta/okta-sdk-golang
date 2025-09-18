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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"time"
)

// checks if the ApiTokenUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTokenUpdate{}

// ApiTokenUpdate An API Token Update Object for an Okta user. This token is NOT scoped any further and can be used for any API that the user has permissions to call.
type ApiTokenUpdate struct {
	// The client name associated with the API Token
	ClientName *string `json:"clientName,omitempty"`
	// The creation date of the API Token
	Created *time.Time `json:"created,omitempty"`
	// The name associated with the API Token
	Name    *string          `json:"name,omitempty"`
	Network *ApiTokenNetwork `json:"network,omitempty"`
	// The userId of the user who created the API Token
	UserId               *string `json:"userId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApiTokenUpdate ApiTokenUpdate

// NewApiTokenUpdate instantiates a new ApiTokenUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTokenUpdate() *ApiTokenUpdate {
	this := ApiTokenUpdate{}
	return &this
}

// NewApiTokenUpdateWithDefaults instantiates a new ApiTokenUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTokenUpdateWithDefaults() *ApiTokenUpdate {
	this := ApiTokenUpdate{}
	return &this
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *ApiTokenUpdate) GetClientName() string {
	if o == nil || IsNil(o.ClientName) {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTokenUpdate) GetClientNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClientName) {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *ApiTokenUpdate) HasClientName() bool {
	if o != nil && !IsNil(o.ClientName) {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *ApiTokenUpdate) SetClientName(v string) {
	o.ClientName = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ApiTokenUpdate) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTokenUpdate) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ApiTokenUpdate) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ApiTokenUpdate) SetCreated(v time.Time) {
	o.Created = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiTokenUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTokenUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiTokenUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiTokenUpdate) SetName(v string) {
	o.Name = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *ApiTokenUpdate) GetNetwork() ApiTokenNetwork {
	if o == nil || IsNil(o.Network) {
		var ret ApiTokenNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTokenUpdate) GetNetworkOk() (*ApiTokenNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *ApiTokenUpdate) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given ApiTokenNetwork and assigns it to the Network field.
func (o *ApiTokenUpdate) SetNetwork(v ApiTokenNetwork) {
	o.Network = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ApiTokenUpdate) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTokenUpdate) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ApiTokenUpdate) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ApiTokenUpdate) SetUserId(v string) {
	o.UserId = &v
}

func (o ApiTokenUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTokenUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientName) {
		toSerialize["clientName"] = o.ClientName
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApiTokenUpdate) UnmarshalJSON(data []byte) (err error) {
	varApiTokenUpdate := _ApiTokenUpdate{}

	err = json.Unmarshal(data, &varApiTokenUpdate)

	if err != nil {
		return err
	}

	*o = ApiTokenUpdate(varApiTokenUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "clientName")
		delete(additionalProperties, "created")
		delete(additionalProperties, "name")
		delete(additionalProperties, "network")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApiTokenUpdate struct {
	value *ApiTokenUpdate
	isSet bool
}

func (v NullableApiTokenUpdate) Get() *ApiTokenUpdate {
	return v.value
}

func (v *NullableApiTokenUpdate) Set(val *ApiTokenUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTokenUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTokenUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTokenUpdate(val *ApiTokenUpdate) *NullableApiTokenUpdate {
	return &NullableApiTokenUpdate{value: val, isSet: true}
}

func (v NullableApiTokenUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTokenUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
