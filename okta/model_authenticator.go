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

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
	"time"
)

// Authenticator struct for Authenticator
type Authenticator struct {
	Created              *time.Time             `json:"created,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Key                  *string                `json:"key,omitempty"`
	LastUpdated          *time.Time             `json:"lastUpdated,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	Provider             *AuthenticatorProvider `json:"provider,omitempty"`
	Settings             *AuthenticatorSettings `json:"settings,omitempty"`
	Status               *string                `json:"status,omitempty"`
	Type                 *string                `json:"type,omitempty"`
	Embedded             *AuthenticatorEmbedded `json:"_embedded,omitempty"`
	Links                *AuthenticatorLinks    `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Authenticator Authenticator

// NewAuthenticator instantiates a new Authenticator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticator() *Authenticator {
	this := Authenticator{}
	return &this
}

// NewAuthenticatorWithDefaults instantiates a new Authenticator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorWithDefaults() *Authenticator {
	this := Authenticator{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Authenticator) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authenticator) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Authenticator) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Authenticator) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Authenticator) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authenticator) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Authenticator) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Authenticator) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *Authenticator) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authenticator) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *Authenticator) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *Authenticator) SetKey(v string) {
	o.Key = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Authenticator) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authenticator) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Authenticator) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Authenticator) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Authenticator) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authenticator) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Authenticator) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Authenticator) SetName(v string) {
	o.Name = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *Authenticator) GetProvider() AuthenticatorProvider {
	if o == nil || o.Provider == nil {
		var ret AuthenticatorProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authenticator) GetProviderOk() (*AuthenticatorProvider, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *Authenticator) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given AuthenticatorProvider and assigns it to the Provider field.
func (o *Authenticator) SetProvider(v AuthenticatorProvider) {
	o.Provider = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *Authenticator) GetSettings() AuthenticatorSettings {
	if o == nil || o.Settings == nil {
		var ret AuthenticatorSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authenticator) GetSettingsOk() (*AuthenticatorSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Authenticator) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given AuthenticatorSettings and assigns it to the Settings field.
func (o *Authenticator) SetSettings(v AuthenticatorSettings) {
	o.Settings = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Authenticator) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authenticator) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Authenticator) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Authenticator) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Authenticator) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authenticator) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Authenticator) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Authenticator) SetType(v string) {
	o.Type = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *Authenticator) GetEmbedded() AuthenticatorEmbedded {
	if o == nil || o.Embedded == nil {
		var ret AuthenticatorEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authenticator) GetEmbeddedOk() (*AuthenticatorEmbedded, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *Authenticator) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given AuthenticatorEmbedded and assigns it to the Embedded field.
func (o *Authenticator) SetEmbedded(v AuthenticatorEmbedded) {
	o.Embedded = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Authenticator) GetLinks() AuthenticatorLinks {
	if o == nil || o.Links == nil {
		var ret AuthenticatorLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authenticator) GetLinksOk() (*AuthenticatorLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Authenticator) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AuthenticatorLinks and assigns it to the Links field.
func (o *Authenticator) SetLinks(v AuthenticatorLinks) {
	o.Links = &v
}

func (o Authenticator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Embedded != nil {
		toSerialize["_embedded"] = o.Embedded
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Authenticator) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticator := _Authenticator{}

	err = json.Unmarshal(bytes, &varAuthenticator)
	if err == nil {
		*o = Authenticator(varAuthenticator)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticator struct {
	value *Authenticator
	isSet bool
}

func (v NullableAuthenticator) Get() *Authenticator {
	return v.value
}

func (v *NullableAuthenticator) Set(val *Authenticator) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticator) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticator(val *Authenticator) *NullableAuthenticator {
	return &NullableAuthenticator{value: val, isSet: true}
}

func (v NullableAuthenticator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
