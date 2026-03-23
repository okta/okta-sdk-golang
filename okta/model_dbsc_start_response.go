/*
Okta Admin Management API

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
	"fmt"
)

// checks if the DbscStartResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbscStartResponse{}

// DbscStartResponse struct for DbscStartResponse
type DbscStartResponse struct {
	Credentials []DbscCredential `json:"credentials"`
	// URL to call for cookie refresh
	RefreshUrl string    `json:"refresh_url"`
	Scope      DbscScope `json:"scope"`
	// The session identifier for this DBSC binding. Use this value in the `Sec-Secure-Session-Id` header for refresh requests.
	SessionIdentifier    string `json:"session_identifier"`
	AdditionalProperties map[string]interface{}
}

type _DbscStartResponse DbscStartResponse

// NewDbscStartResponse instantiates a new DbscStartResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbscStartResponse(credentials []DbscCredential, refreshUrl string, scope DbscScope, sessionIdentifier string) *DbscStartResponse {
	this := DbscStartResponse{}
	this.Credentials = credentials
	this.RefreshUrl = refreshUrl
	this.Scope = scope
	this.SessionIdentifier = sessionIdentifier
	return &this
}

// NewDbscStartResponseWithDefaults instantiates a new DbscStartResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbscStartResponseWithDefaults() *DbscStartResponse {
	this := DbscStartResponse{}
	return &this
}

// GetCredentials returns the Credentials field value
func (o *DbscStartResponse) GetCredentials() []DbscCredential {
	if o == nil {
		var ret []DbscCredential
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *DbscStartResponse) GetCredentialsOk() ([]DbscCredential, bool) {
	if o == nil {
		return nil, false
	}
	return o.Credentials, true
}

// SetCredentials sets field value
func (o *DbscStartResponse) SetCredentials(v []DbscCredential) {
	o.Credentials = v
}

// GetRefreshUrl returns the RefreshUrl field value
func (o *DbscStartResponse) GetRefreshUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefreshUrl
}

// GetRefreshUrlOk returns a tuple with the RefreshUrl field value
// and a boolean to check if the value has been set.
func (o *DbscStartResponse) GetRefreshUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshUrl, true
}

// SetRefreshUrl sets field value
func (o *DbscStartResponse) SetRefreshUrl(v string) {
	o.RefreshUrl = v
}

// GetScope returns the Scope field value
func (o *DbscStartResponse) GetScope() DbscScope {
	if o == nil {
		var ret DbscScope
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *DbscStartResponse) GetScopeOk() (*DbscScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *DbscStartResponse) SetScope(v DbscScope) {
	o.Scope = v
}

// GetSessionIdentifier returns the SessionIdentifier field value
func (o *DbscStartResponse) GetSessionIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionIdentifier
}

// GetSessionIdentifierOk returns a tuple with the SessionIdentifier field value
// and a boolean to check if the value has been set.
func (o *DbscStartResponse) GetSessionIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionIdentifier, true
}

// SetSessionIdentifier sets field value
func (o *DbscStartResponse) SetSessionIdentifier(v string) {
	o.SessionIdentifier = v
}

func (o DbscStartResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbscStartResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["credentials"] = o.Credentials
	toSerialize["refresh_url"] = o.RefreshUrl
	toSerialize["scope"] = o.Scope
	toSerialize["session_identifier"] = o.SessionIdentifier

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DbscStartResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"credentials",
		"refresh_url",
		"scope",
		"session_identifier",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDbscStartResponse := _DbscStartResponse{}

	err = json.Unmarshal(data, &varDbscStartResponse)

	if err != nil {
		return err
	}

	*o = DbscStartResponse(varDbscStartResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "refresh_url")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "session_identifier")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDbscStartResponse struct {
	value *DbscStartResponse
	isSet bool
}

func (v NullableDbscStartResponse) Get() *DbscStartResponse {
	return v.value
}

func (v *NullableDbscStartResponse) Set(val *DbscStartResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDbscStartResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDbscStartResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbscStartResponse(val *DbscStartResponse) *NullableDbscStartResponse {
	return &NullableDbscStartResponse{value: val, isSet: true}
}

func (v NullableDbscStartResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbscStartResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
