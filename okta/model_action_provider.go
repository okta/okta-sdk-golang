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
	"fmt"
)

// checks if the ActionProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionProvider{}

// ActionProvider struct for ActionProvider
type ActionProvider struct {
	// The unique identifier of the action flow in the provider system
	ExternalId string `json:"externalId"`
	// Type of action provider
	Type string `json:"type"`
	// The URL to the action flow
	Url                  string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _ActionProvider ActionProvider

// NewActionProvider instantiates a new ActionProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionProvider(externalId string, type_ string, url string) *ActionProvider {
	this := ActionProvider{}
	this.ExternalId = externalId
	this.Type = type_
	this.Url = url
	return &this
}

// NewActionProviderWithDefaults instantiates a new ActionProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionProviderWithDefaults() *ActionProvider {
	this := ActionProvider{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *ActionProvider) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *ActionProvider) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *ActionProvider) SetExternalId(v string) {
	o.ExternalId = v
}

// GetType returns the Type field value
func (o *ActionProvider) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ActionProvider) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ActionProvider) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *ActionProvider) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ActionProvider) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ActionProvider) SetUrl(v string) {
	o.Url = v
}

func (o ActionProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalId"] = o.ExternalId
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ActionProvider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"externalId",
		"type",
		"url",
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

	varActionProvider := _ActionProvider{}

	err = json.Unmarshal(data, &varActionProvider)

	if err != nil {
		return err
	}

	*o = ActionProvider(varActionProvider)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableActionProvider struct {
	value *ActionProvider
	isSet bool
}

func (v NullableActionProvider) Get() *ActionProvider {
	return v.value
}

func (v *NullableActionProvider) Set(val *ActionProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableActionProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableActionProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionProvider(val *ActionProvider) *NullableActionProvider {
	return &NullableActionProvider{value: val, isSet: true}
}

func (v NullableActionProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
