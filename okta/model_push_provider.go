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

// PushProvider struct for PushProvider
type PushProvider struct {
	// Unique key for the Push Provider
	Id *string `json:"id,omitempty"`
	// Timestamp when the Push Provider was last modified
	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty"`
	// Display name of the push provider
	Name *string `json:"name,omitempty"`
	ProviderType *string `json:"providerType,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PushProvider PushProvider

// NewPushProvider instantiates a new PushProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushProvider() *PushProvider {
	this := PushProvider{}
	return &this
}

// NewPushProviderWithDefaults instantiates a new PushProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushProviderWithDefaults() *PushProvider {
	this := PushProvider{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PushProvider) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushProvider) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PushProvider) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PushProvider) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *PushProvider) GetLastUpdatedDate() string {
	if o == nil || o.LastUpdatedDate == nil {
		var ret string
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushProvider) GetLastUpdatedDateOk() (*string, bool) {
	if o == nil || o.LastUpdatedDate == nil {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *PushProvider) HasLastUpdatedDate() bool {
	if o != nil && o.LastUpdatedDate != nil {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given string and assigns it to the LastUpdatedDate field.
func (o *PushProvider) SetLastUpdatedDate(v string) {
	o.LastUpdatedDate = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PushProvider) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushProvider) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PushProvider) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PushProvider) SetName(v string) {
	o.Name = &v
}

// GetProviderType returns the ProviderType field value if set, zero value otherwise.
func (o *PushProvider) GetProviderType() string {
	if o == nil || o.ProviderType == nil {
		var ret string
		return ret
	}
	return *o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushProvider) GetProviderTypeOk() (*string, bool) {
	if o == nil || o.ProviderType == nil {
		return nil, false
	}
	return o.ProviderType, true
}

// HasProviderType returns a boolean if a field has been set.
func (o *PushProvider) HasProviderType() bool {
	if o != nil && o.ProviderType != nil {
		return true
	}

	return false
}

// SetProviderType gets a reference to the given string and assigns it to the ProviderType field.
func (o *PushProvider) SetProviderType(v string) {
	o.ProviderType = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PushProvider) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushProvider) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PushProvider) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *PushProvider) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o PushProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdatedDate != nil {
		toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProviderType != nil {
		toSerialize["providerType"] = o.ProviderType
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PushProvider) UnmarshalJSON(bytes []byte) (err error) {
	varPushProvider := _PushProvider{}

	err = json.Unmarshal(bytes, &varPushProvider)
	if err == nil {
		*o = PushProvider(varPushProvider)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdatedDate")
		delete(additionalProperties, "name")
		delete(additionalProperties, "providerType")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePushProvider struct {
	value *PushProvider
	isSet bool
}

func (v NullablePushProvider) Get() *PushProvider {
	return v.value
}

func (v *NullablePushProvider) Set(val *PushProvider) {
	v.value = val
	v.isSet = true
}

func (v NullablePushProvider) IsSet() bool {
	return v.isSet
}

func (v *NullablePushProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushProvider(val *PushProvider) *NullablePushProvider {
	return &NullablePushProvider{value: val, isSet: true}
}

func (v NullablePushProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

